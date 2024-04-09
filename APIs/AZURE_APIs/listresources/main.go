package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// connectGCP establishes a connection to Google Cloud Storage (GCS) using the provided context.
func connectGCP(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Fatalf("Failed to create GCP client: %v", err)
	}
	log.Println("Successfully authenticated with GCP")
	return client
}

func connectAzure() *azidentity.ClientSecretCredential {
	clientID := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	tenantID := os.Getenv("AZURE_TENANT_ID")

	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		log.Fatalf("Failed to create Azure credentials: %v", err)
	}
	log.Println("Successfully connected to Azure")
	return cred
}

// listAllResources lists all the resources in the specified Azure subscription.
func listAllResources(cred *azidentity.ClientSecretCredential, subscriptionID string) {
	client, err := armresources.NewClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create resources client: %v", err)
	}

	pager := client.NewListPager(nil)

	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		if err != nil {
			log.Fatalf("Failed to get the next page of resources: %v", err)
		}

		for _, resource := range resp.ResourceListResult.Value {
			resourceJSON, err := json.MarshalIndent(resource, "", "  ")
			if err != nil {
				log.Printf("Failed to marshal resource: %v", err)
				continue
			}
			log.Printf("Resource Details:\n%s\n", resourceJSON)
		}
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Authenticate to Azure
	azureCred := connectAzure()

	// List all resources in the Azure subscription
	subscriptionID := os.Getenv("SUBSCRIPTIONID")
	listAllResources(azureCred, subscriptionID)
}
