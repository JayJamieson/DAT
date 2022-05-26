# Go channels and go routines

Example here shows how easy it is to do concurrency

## Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

## Goroutines

The primary concurrency construct is the goroutine, a type of light-weight process. A function call prefixed with the go keyword starts a function in a new goroutine.

- "don't communicate by sharing memory; share memory by communicating"
