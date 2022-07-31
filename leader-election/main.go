package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Node struct {
	mu        sync.Mutex
	ctx       context.Context
	cancel    func()
	g         errgroup.Group
	isPrimary bool
	Leaser    *Leaser
}

func NewNode(leaser *Leaser) *Node {
	n := &Node{
		Leaser: leaser,
	}

	n.ctx, n.cancel = context.WithCancel(context.Background())

	return n
}

// Close signals for the store to shut down.
func (n *Node) Close() error {
	n.cancel()
	return n.g.Wait()
}

func (n *Node) monitor(ctx context.Context) error {
	for {
		// Exit if store is closed.
		if err := ctx.Err(); err != nil {
			return nil
		}

		lease, primaryPort, err := n.acquireLeaseOrPrimaryPort(ctx)

		if err != nil {
			log.Printf("cannot acquire lease or find primary, retrying: %s", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// Monitor as primary if we have obtained a lease.
		if lease != nil {
			log.Printf("primary lease acquired, advertising as %s", n.Leaser.AdvertisePort())
			if err := n.monitorAsPrimary(ctx, lease); err != nil {
				log.Printf("primary lease lost, retrying: %s", err)
			}
			continue
		}

		// Monitor as replica if another primary already exists.
		log.Printf("existing primary found (%s), waiting as replica", primaryPort)
		if err := n.monitorChanges(ctx); err != nil {
			log.Printf("stopped monitoring changes: %s", err)
		}
		continue
	}
}

func (n *Node) acquireLeaseOrPrimaryPort(ctx context.Context) (*Lease, string, error) {
	// Attempt to find an existing primary first.
	primaryURL, err := n.Leaser.PrimaryURL(ctx)
	if err != nil && err != ErrNoPrimary {
		return nil, "", fmt.Errorf("fetch primary url: %w", err)
	} else if primaryURL != "" {
		return nil, primaryURL, nil
	}

	// If no primary, attempt to become primary.
	lease, err := n.Leaser.Acquire(ctx)
	if err != nil && err != ErrPrimaryExists {
		return nil, "", fmt.Errorf("acquire lease: %w", err)
	} else if lease != nil {
		return lease, "", nil
	}

	// If we raced to become primary and another node beat us, retry the fetch.
	primaryURL, err = n.Leaser.PrimaryURL(ctx)
	if err != nil {
		return nil, "", err
	}
	return nil, primaryURL, nil
}

func (n *Node) monitorAsPrimary(ctx context.Context, lease *Lease) error {
	const timeout = 1 * time.Second

	// Attempt to destroy lease when we exit this function.
	defer func() {
		log.Printf("exiting primary, destroying lease")
		if err := lease.Close(); err != nil {
			log.Printf("cannot remove lease: %s", err)
		}
	}()

	// Mark as the primary node while we're in this function.
	n.mu.Lock()
	n.isPrimary = true
	n.mu.Unlock()

	// Ensure that we are no longer marked as primary once we exit this function.
	defer func() {
		n.mu.Lock()
		defer n.mu.Unlock()
		n.isPrimary = false
	}()

	waitDur := lease.TTL() / 2

	for {
		select {
		case <-time.After(waitDur):
			// Attempt to renew the lease. If the lease is gone then we need to
			// just exit and we can start over or connect to the new primary.
			//
			// If we just have a connection error then we'll try to more
			// aggressively retry the renewal until we exceed TTL.
			if err := lease.Renew(ctx); err == ErrLeaseExpired {
				return err
			} else if err != nil {
				// If our next renewal will exceed TTL, exit now.
				if time.Since(lease.RenewedAt())+timeout > lease.TTL() {
					time.Sleep(timeout)
					return ErrLeaseExpired
				}

				// Otherwise log error and try again after a shorter period.
				log.Printf("lease renewal error, retrying: %s", err)
				waitDur = time.Second
				continue
			}

			// Renewal was successful, restart with low frequency.
			waitDur = lease.TTL() / 2

		case <-ctx.Done():
			return nil // release lease when we shut down
		}
	}
}

func (n *Node) monitorChanges(ctx context.Context) error {
	const wait = 1 * time.Second
	var modifyIndex uint64
	var current string

	current, modifyIndex, err := n.Leaser.WaitChanges(ctx, 0)

	if err != nil {
		log.Printf("Encounter error waiting for change with error: %v", err)
		return err
	}

	for {
		select {
		case <-time.After(wait):
			var newChange string
			newChange, modifyIndex, err = n.Leaser.WaitChanges(ctx, modifyIndex)

			if err == ErrNoPrimary {
				log.Printf("%v", err)
				return err
			}
			if current != newChange && newChange != "" {
				log.Printf("current state: %v, new state: %v", current, newChange)
			}

			current = newChange
			continue
		case <-ctx.Done():
			return nil // release lease when we shut down
		}
	}
}

func main() {
	port := flag.String("port", "", "")
	flag.Parse()

	leaser := NewLeaser("http://localhost:8500", *port)
	err := leaser.Open()

	if err != nil {
		panic("Unable to connect to consul")
	}

	node := NewNode(leaser)

	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	_, cancel := context.WithCancel(context.Background())

	node.g.Go(func() error { return node.monitor(node.ctx) })

	<-signalCh
	node.Close()
	cancel()
	log.Println("signal received, node shutting down")
}
