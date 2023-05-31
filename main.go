package main

import (
	"fmt"

	"github.com/natrontech/pocketbase-client-go/pkg/client"
	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

func main() {
	endpoint := "http://127.0.0.1:8090"
	identity := "admin@natron.io"
	password := "0123456789"

	c, err := client.NewClient(&endpoint, &identity, &password)
	if err != nil {
		panic(err)
	}

	newAdmin := models.AdminCreateRequest{
		Email:           "test@test.ch",
		Password:        "0123456789",
		PasswordConfirm: "0123456789",
	}
	ar, err := c.CreateAdmin(&newAdmin)
	if err != nil {
		panic(err)
	}

	fmt.Println(ar)
}
