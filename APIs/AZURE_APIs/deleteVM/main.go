package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/joho/godotenv"
)

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

// deleteAzureVM deletes the specified virtual machine in Azure.
func deleteAzureVM(cred *azidentity.ClientSecretCredential, subscriptionID, resourceGroupName, vmName string) {
	vmClient, err := armcompute.NewVirtualMachinesClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create VM client: %v", err)
	}

	// Start the deletion process
	poller, err := vmClient.BeginDelete(context.Background(), resourceGroupName, vmName, nil)
	if err != nil {
		log.Fatalf("Failed to start VM deletion: %v", err)
	}

	// Wait for the deletion process to complete
	_, err = poller.PollUntilDone(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to delete VM: %v", err)
	}

	log.Printf("Successfully deleted VM: %s", vmName)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	azureCred := connectAzure()
	subscriptionID := os.Getenv("SUBSCRIPTIONID")
	resourceGroupName := os.Getenv("RESOURCEGROUPNAME")
	vmName := os.Getenv("VMNAME")

	// Delete Azure VM
	deleteAzureVM(azureCred, subscriptionID, resourceGroupName, vmName)
}
