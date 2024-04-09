package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getResourceGroupName(resourceID string) (string, error) {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if part == "resourceGroups" && i+1 < len(parts) {
			return parts[i+1], nil
		}
	}
	return "", fmt.Errorf("resource group name not found in resource ID: %s", resourceID)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	resourceID := os.Getenv("RESOURCEID")
	resourceGroupName, err := getResourceGroupName(resourceID)
	if err != nil {
		log.Printf("Error extracting resource group name: %v", err)
	} else {
		log.Printf("Resource Group Name: %s", resourceGroupName)
	}
}
