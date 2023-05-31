package main

import (
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

	// base := models.CollectionCreateRequest{
	// 	Name: "Base",
	// 	Type: "base",
	// 	Schema: []models.Schema{
	// 		models.Schema{
	// 			Name:     "title",
	// 			Type:     "text",
	// 			Required: true,
	// 			Options: struct {
	// 				Min int `json:"min"`
	// 			}{
	// 				Min: 10,
	// 			},
	// 		},
	// 	},
	// }

	// ar, err := c.CreateCollection(&base)
	// if err != nil {
	// 	panic(err)
	// }

	// println(ar.Id)

	health, err := c.HealthCheck()
	if err != nil {
		panic(err)
	}

	println(health.Message)

}
