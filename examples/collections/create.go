package examples_collections

import (
	"github.com/natrontech/pocketbase-client-go/examples"
	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

func CreateCollection() {

	base := models.CollectionCreateRequest{
		Name: "Base",
		Type: "base",
		Schema: []models.Schema{
			models.Schema{
				Name:     "title",
				Type:     "text",
				Required: true,
				Options: struct {
					Min int `json:"min"`
				}{
					Min: 10,
				},
			},
		},
	}

	resp, err := examples.Client.CreateCollection(&base)
	if err != nil {
		panic(err)
	}

	println(resp.Id)
}
