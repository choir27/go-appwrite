package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/client"
)

func main() {
	appwriteClient := client.NewClient()
	appwriteClient.SetEndpoint("https://v16.appwrite.org/v1")
	appwriteClient.SetProject("669c026b0002eaad620f")

	service := account.NewAccount(appwriteClient)

	res, err := service.ListLogs(
		account.WithListLogsQueries([]interface{}{}),
	)

	fmt.Println(res)

	if err != nil {
		fmt.Println(err)
		return
	}

	response, error := service.GetSession(
		"SESSION_ID",
	)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(response)
}
