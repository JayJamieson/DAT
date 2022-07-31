# Leader Election

## Running

### On Local machine

- Start consul container using command found at [Run consul container](##Run-consul-container)
- `go build -o node .` to build binary

- `./node -port 9000 -consul http://localhost:8500`

  - `-port` is used as a basic identifier for session/lease value
  - `-consul` is address of consul node to use for leader election, can be ommited if running locally

### With Docker compose

- `docker build -t node .` to build a docker container for a node.
- `docker compose up` to start consul and three nodes in a *basic* cluster.

## Run consul container

Run consule outside of docker compose.

```bash
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=badger \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

## References

- <https://martinfowler.com/articles/patterns-of-distributed-systems/leader-follower.html>
- <https://martinfowler.com/articles/patterns-of-distributed-systems/consistent-core.html>
- <https://martinfowler.com/articles/patterns-of-distributed-systems/time-bound-lease.html>
