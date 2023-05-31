package examples_health

import "github.com/natrontech/pocketbase-client-go/examples"

func DoHealthCheck() {
	resp, err := examples.Client.HealthCheck()
	if err != nil {
		panic(err)
	}

	println(resp.Message)
}
