package main

import (
	"fmt"

	"github.com/natrontech/pocketbase-client-go/pkg/client"
)

func main() {
	endpoint := "http://127.0.0.1:8090"
	identity := "admin@natron.io"
	password := "0123456789"

	c, err := client.NewClient(&endpoint, &identity, &password)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", c)
}
