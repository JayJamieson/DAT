package main

import "github.com/hashicorp/consul/api"

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	session := client.Session()

	session.CreateNoChecks(&api.SessionEntry{
		Name: "worker1",
	}, &api.WriteOptions{})
}
