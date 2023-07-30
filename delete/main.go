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
		Username:  username,
		Password:  password,
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

	// Delete a document
	indexName := "dev-article"
	documentID := "1"
	err = deleteDocument(client, indexName, documentID)
	if err != nil {
		fmt.Println("Error deleting document:", err)
		return
	}
	fmt.Println("Document deleted:", documentID)
}

func deleteDocument(client *opensearch.Client, indexName string, documentID string) error {

	_, err := client.Delete(indexName, documentID)
	if err != nil {
		return err
	}

	return nil
}