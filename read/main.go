package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
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

	// Retrieve a document
	indexName := "dev-article"
	documentID := "1"
	retrievedDocument, err := getDocument(client, indexName, documentID)
	if err != nil {
		fmt.Println("Error retrieving document:", err)
		return
	}
	fmt.Println("Retrieved Document:", retrievedDocument)
}

func getDocument(client *opensearch.Client, indexName string, documentID string) (interface{}, error) {

	response, err := client.Get(indexName, documentID)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data := make(map[string]interface{})

	resp, errResp := io.ReadAll(response.Body)
	if errResp != nil {
		return nil, errResp
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}