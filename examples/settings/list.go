package examples_settings

import "github.com/natrontech/pocketbase-client-go/examples"

func ListSettings() {
	resp, err := examples.Client.ListSettings(nil)
	if err != nil {
		panic(err)
	}

	println(resp.Logs.MaxDays)
}
