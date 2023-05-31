package main

import (
	examples_collections "github.com/natrontech/pocketbase-client-go/examples/collections"
	examples_health "github.com/natrontech/pocketbase-client-go/examples/health"
)

func main() {
	examples_collections.CreateCollection()
	examples_health.DoHealthCheck()
}
