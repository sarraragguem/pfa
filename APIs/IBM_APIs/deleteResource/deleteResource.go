package main

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func deleteResource(instanceGUID string) {
	// Initialize the IAM authenticator using your API key.
	apiKey := GetApiKey()
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	// Create a new Resource Controller service instance.
	resourceController, err := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
		Authenticator: authenticator,
	})
	if err != nil {
		log.Fatalf("Failed to create Resource Controller service: %v", err)
	}

	// ID of the resource instance you want to delete.
	resourceInstanceID := instanceGUID

	// Delete the resource instance.
	deleteResourceInstanceOptions := resourceController.NewDeleteResourceInstanceOptions(resourceInstanceID)
	_, err = resourceController.DeleteResourceInstance(deleteResourceInstanceOptions)
	if err != nil {
		log.Fatalf("Failed to delete resource instance: %v", err)
	} else {
		fmt.Printf("Resource instance %s deleted successfully\n", resourceInstanceID)
	}
}
func main() {
	deleteResource("crn:v1:bluemix:public:cloud-object-storage:global:a/e81bb157bc654c579dee87e38bb1db3b:874e0f5d-af20-47af-a6e8-e8d63057aacc::")
}
