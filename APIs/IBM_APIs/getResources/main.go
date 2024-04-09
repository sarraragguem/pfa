package main

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type ResourceDetail struct {
	ID              string
	Name            string
	State           string
	ResourcePlanID  string
	ResourceGroupID string
	RegionID        string
	AccountID       string
	CreatedAt       string
	UpdatedAt       string
}

func getResources() []ResourceDetail {

	apiKey := GetApiKey()

	// Create an IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	// Create the service client.
	resourceController, err := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
		Authenticator: authenticator,
	})
	if err != nil {
		log.Fatalf("Error creating Resource Controller service: %v", err)
	}

	// List resources.
	listResourceInstancesOptions := resourceController.NewListResourceInstancesOptions()
	resources, _, err := resourceController.ListResourceInstances(listResourceInstancesOptions)
	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	var resourceDetails []ResourceDetail

	// Populate the slice with resource details.
	for _, resource := range resources.Resources {
		detail := ResourceDetail{
			ID:              *resource.ID,
			Name:            *resource.Name,
			State:           *resource.State,
			ResourcePlanID:  *resource.ResourcePlanID,
			ResourceGroupID: *resource.ResourceGroupID,
			RegionID:        *resource.RegionID,
			AccountID:       *resource.AccountID,
			CreatedAt:       resource.CreatedAt.String(),
			UpdatedAt:       resource.UpdatedAt.String(),
		}
		resourceDetails = append(resourceDetails, detail)
	}

	fmt.Println(resourceDetails)
	return resourceDetails

}
func main() {

	var resourceDetails = getResources()

	for _, detail := range resourceDetails {
		fmt.Printf("ID: %s - The unique identifier of the resource.\n", detail.ID)
		fmt.Printf("Name: %s - The name of the resource.\n", detail.Name)
		fmt.Printf("State: %s - The current state of the resource (e.g., active, deleted, etc.).\n", detail.State)
		fmt.Printf("Resource Plan ID: %s - The identifier for the resource's plan (specifies the tier, features, etc.).\n", detail.ResourcePlanID)
		fmt.Printf("Resource Group ID: %s - The identifier of the resource group to which this resource belongs.\n", detail.ResourceGroupID)
		fmt.Printf("Region ID: %s - The region or location where the resource is provisioned.\n", detail.RegionID)
		fmt.Printf("Account ID: %s - The identifier of the IBM Cloud account that owns the resource.\n", detail.AccountID)
		fmt.Printf("Created At: %s - The timestamp when the resource was created.\n", detail.CreatedAt)
		fmt.Printf("Updated At: %s - The timestamp when the resource was last updated.\n", detail.UpdatedAt)
		fmt.Println("---------------")
	}

}
