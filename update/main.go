package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchutil"
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

	// Update a document
	indexName := "dev-article"
	documentID := "1"
	updatedFields := map[string]interface{}{
		"doc": map[string]interface{}{
			"content": "Updated api content-- OpenSearch is a powerful open-source search",
		},
	}

	err = updateDocument(client, indexName, documentID, updatedFields)
	if err != nil {
		fmt.Println("Error updating document:", err)
		return
	}
	fmt.Println("Document updated:", documentID)
}

func updateDocument(client *opensearch.Client, indexName string, documentID string, updatedFields map[string]interface{}) error {
	res, err := client.Update(indexName, documentID, opensearchutil.NewJSONReader(updatedFields))

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("update document request failed: %s", res.String())
	}

	return nil
}