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

func updateVMSize(cred *azidentity.ClientSecretCredential, subscriptionID, resourceGroupName, vmName, newSize string) {
	ctx := context.Background()
	vmClient, err := armcompute.NewVirtualMachinesClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create VM client: %v", err)
	}

	// Deallocate VM before resizing
	poller, err := vmClient.BeginDeallocate(ctx, resourceGroupName, vmName, nil)
	if err != nil {
		log.Fatalf("Failed to deallocate VM: %v", err)
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to complete VM deallocation: %v", err)
	}

	// Get the current VM
	vm, err := vmClient.Get(ctx, resourceGroupName, vmName, nil)
	if err != nil {
		log.Fatalf("Failed to get VM: %v", err)
	}

	// Updating VM size
	if newSize != "" {
		size := armcompute.VirtualMachineSizeTypes(newSize)
		vm.Properties.HardwareProfile.VMSize = &size
	}

	_, err = vmClient.BeginUpdate(ctx, resourceGroupName, vmName, armcompute.VirtualMachineUpdate{
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: vm.Properties.HardwareProfile.VMSize,
			},
		},
	}, nil)

	if err != nil {
		log.Fatalf("Failed to update VM size: %v", err)
	}

	// Start the VM after updating
	startPoller, err := vmClient.BeginStart(ctx, resourceGroupName, vmName, nil)
	if err != nil {
		log.Fatalf("Failed to start VM: %v", err)
	}

	_, err = startPoller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to start VM: %v", err)
	}

	log.Printf("Successfully updated and started VM: %s", vmName)
}

func updateVMTags(cred *azidentity.ClientSecretCredential, subscriptionID, resourceGroupName, vmName string, newTags map[string]*string) {
	ctx := context.Background()
	vmClient, err := armcompute.NewVirtualMachinesClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create VM client: %v", err)
	}

	// Update VM tags
	_, err = vmClient.BeginUpdate(ctx, resourceGroupName, vmName, armcompute.VirtualMachineUpdate{
		Tags: newTags,
	}, nil)

	if err != nil {
		log.Fatalf("Failed to update VM tags: %v", err)
	}

	log.Printf("Successfully updated VM tags for VM: %s", vmName)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	azureCred := connectAzure()
	subscriptionID := os.Getenv("SUBSCRIPTIONID")
	resourceGroupName := os.Getenv("RESOURCEGROUPNAME")
	vmName := os.Getenv("VMNAME")

	// Update Azure VM size
	newSize := os.Getenv("NEWSIZE")
	updateVMSize(azureCred, subscriptionID, resourceGroupName, vmName, newSize)

	// Update Azure VM tags
	environment := os.Getenv("ENVIRONMENT")
	department := os.Getenv("DEPARTMENT")

	// Initialize map with environment variable values
	newTags := map[string]*string{
		"Environment": &environment,
		"Department":  &department,
	}
	updateVMTags(azureCred, subscriptionID, resourceGroupName, vmName, newTags)
}
