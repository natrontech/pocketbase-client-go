package main

import examples_settings "github.com/natrontech/pocketbase-client-go/examples/settings"

// examples_collections "github.com/natrontech/pocketbase-client-go/examples/collections"
// examples_health "github.com/natrontech/pocketbase-client-go/examples/health"

func main() {
	// examples_collections.CreateCollection()
	// examples_health.DoHealthCheck()
	examples_settings.ListSettings()
	examples_settings.UpdateSettings()
}
