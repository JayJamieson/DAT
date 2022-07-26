package main

import (
	"io/ioutil"
	"log"

	"github.com/hashicorp/consul/api"
)

func getSession() (string, error) {
	sessionId, err := ioutil.ReadFile("session.id")
	return string(sessionId), err
}

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	session := client.Session()

	existingSessionId, err := getSession()

	var sessionId = existingSessionId

	if err != nil {
		log.Println("No existing session found")

		sessionId, _, err = session.CreateNoChecks(&api.SessionEntry{
			Name: "worker1",
		}, &api.WriteOptions{})

		if err != nil {
			panic(err)
		}

		log.Println("Created new session in session.id file")
		ioutil.WriteFile("session.id", []byte(sessionId), 0644)
	}

	log.Printf("Session ID %v", sessionId)

	// get a lock on the key using session
	// client.KV().Acquire()

	// renew the session to keep the lock active
	// session.RenewPeriodic()
}
