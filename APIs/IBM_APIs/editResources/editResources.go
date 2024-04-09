package main

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type ResourceUpdate struct {
	Name           *string
	Parameters     map[string]interface{}
	ResourcePlanID *string
}

func updateResourceInstance(service *resourcecontrollerv2.ResourceControllerV2, resourceID string, updates ResourceUpdate) {
	options := &resourcecontrollerv2.UpdateResourceInstanceOptions{
		ID: &resourceID,
	}

	if updates.Name != nil {
		options.Name = updates.Name
	}

	if updates.Parameters != nil {
		options.Parameters = updates.Parameters
	}

	if updates.ResourcePlanID != nil {
		options.ResourcePlanID = updates.ResourcePlanID
	}

	result, response, err := service.UpdateResourceInstance(options)
	if err != nil {
		log.Fatalf("Failed to update resource instance: %v", err)
	}

	fmt.Printf("-------------------------------- \n")
	fmt.Printf("Update response: %+v\n", response)
	fmt.Printf("--------------------------------\n")
	fmt.Printf("Update result: %+v\n", result)
}

func main() {
	authenticator := &core.IamAuthenticator{
		ApiKey: GetApiKey(),
	}

	resourceController, err := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
		Authenticator: authenticator,
	})
	if err != nil {
		log.Fatalf("Failed to create Resource Controller service: %v", err)
	}

	resourceID := "crn:v1:bluemix:public:cloud-object-storage:global:a/e81bb157bc654c579dee87e38bb1db3b:988141fb-3c50-4897-93a2-ee1e8a46bfe4::"
	updates := ResourceUpdate{
		Name: core.StringPtr("cloudObjectStorage"),
		// Populate other fields as needed
	}

	updateResourceInstance(resourceController, resourceID, updates)
}
