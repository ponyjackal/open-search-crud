package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
)

func main() {
	endpoint := "https://localhost:9200"
    username := "admin" // Leave empty if not using authentication
    password := "admin" // Leave empty if not using authentication

	// Create a client
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{endpoint},
		Username: username,
		Password: password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		fmt.Println("Error creating OpenSearch client:", err)
        return
	}

	// Create an index
	indexName := "dev-article"
	err = createIndex(client, indexName)
	if err != nil {
		fmt.Println("Error creating index: ", err)
		return
	}
	fmt.Println("Index created: ", indexName)
}

func createIndex(client *opensearch.Client, indexName string) error {
	_, err := client.Indices.Create(indexName)
	if err != nil {
		return err
	}
	return nil
}