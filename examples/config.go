package examples

import "github.com/natrontech/pocketbase-client-go/pkg/client"

var (
	Endpoint = "http://127.0.0.1:8090"
	Identity = "admin@natron.io"
	Password = "0123456789"
	Client   *client.Client
)

func init() {
	var err error
	Client, err = client.NewClient(&Endpoint, &Identity, &Password)
	if err != nil {
		panic(err)
	}
}
