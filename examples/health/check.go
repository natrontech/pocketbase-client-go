package examples_health

import "github.com/natrontech/pocketbase-client-go/examples"

func DoHealthCheck() {
	ar, err := examples.Client.HealthCheck()
	if err != nil {
		panic(err)
	}

	println(ar.Message)
}
