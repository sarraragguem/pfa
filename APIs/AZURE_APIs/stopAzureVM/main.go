// stopVM.go
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

func stopAzureVM(cred *azidentity.ClientSecretCredential, subscriptionID, resourceGroupName, vmName string) {
	vmClient, err := armcompute.NewVirtualMachinesClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create VM client: %v", err)
	}

	poller, err := vmClient.BeginDeallocate(context.Background(), resourceGroupName, vmName, nil)
	if err != nil {
		log.Fatalf("Failed to start VM deallocation: %v", err)
	}

	_, err = poller.PollUntilDone(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to deallocate VM: %v", err)
	}

	log.Printf("Successfully deallocated VM: %s", vmName)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	azureCred := connectAzure()
	subscriptionID := os.Getenv("SUBSCRIPTIONID")
	vmName := os.Getenv("VMNAME")
	resourceGroupName := os.Getenv("RESOURCEGROUPNAME")

	stopAzureVM(azureCred, subscriptionID, resourceGroupName, vmName)
}
