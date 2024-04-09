package main

import (
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
)

func main() {
	apikey := GetApiKey()
	authenticator := &core.IamAuthenticator{
		ApiKey: apikey,
	}

	service, err := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
		Authenticator: authenticator,
	})

	if err != nil {
		fmt.Println("Service creation failed: ", err)
		return
	}

	// Attempt to list resource groups as a way to validate the API key
	listResourceGroupsOptions := service.NewListResourceGroupsOptions()
	resourceGroups, response, err := service.ListResourceGroups(listResourceGroupsOptions)

	if err != nil {
		fmt.Printf("Failed to list resource groups: %s\n", err)
		return
	}

	fmt.Printf(apikey)
	fmt.Printf("ListResourceGroups response: %s\n", response)
	fmt.Printf("Resource Groups: %s\n", resourceGroups)
}
