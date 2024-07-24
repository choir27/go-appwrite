package main

import (
	"encoding/json"
	"fmt"

	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
)

func main() {
	appwriteClient := client.NewClient()
	appwriteClient.SetEndpoint("https://v16.appwrite.org/v1")
	appwriteClient.SetProject("669c026b0002eaad620f")
	appwriteClient.SetKey("standard_f8bdb93356c87ee8ed5284431e896cbb1c9bf40a57e6945c44a9b194901415506783d649c332a2d71daab0f3f9c2a790d3a7b4a94a882e9be81ace80a7114a0d3975fdf316f5c97059c7e7d82cfdf69a98ffaa163f8bfacf1aa484534614998e4077fcbd18f6357b9669f7385a740730ab7d71c68e0834b8cfbbe38c03073def")

	type Pokemon struct {
		Name        string   `json:"Name"`
		PokemonType string   `json:"PokemonType"`
		MoveSet     []string `json:"MoveSet"`
	}

	person1 := Pokemon{
		Name:        "Misdrevous",
		PokemonType: "Ghost",
		MoveSet:     []string{"ball", "rock", "ball", "pain"},
	}

	// Marshal the Pokemon struct into JSON
	person1JSON, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshaling Pokemon to JSON:", err)
		return
	}

	appwriteDatabases := databases.NewDatabases(appwriteClient)

	response, err := appwriteDatabases.CreateDocument("669c028b00237f0ac4bf", "669c02990025ab75e1d4", "unique()", json.RawMessage(person1JSON))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
