package main

import (
	"context"
	"fmt"
	"log"
	"os"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func listProjectIamPolicies(ctx context.Context, projectID string) {
	service, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		log.Fatalf("cloudresourcemanager.NewService: %v", err)
	}

	resp, err := service.Projects.GetIamPolicy(projectID, &cloudresourcemanager.GetIamPolicyRequest{}).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Projects.GetIamPolicy: %v", err)
	}

	fmt.Printf("IAM Policy for project %s:\n", projectID)
	for _, binding := range resp.Bindings {
		fmt.Printf("Role: %s\n", binding.Role)
		for _, member := range binding.Members {
			fmt.Printf(" - Member: %s\n", member)
		}
	}
}

func listInstances(ctx context.Context, projectID string) {
	c, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create instance client: %v", err)
	}
	defer c.Close()

	req := &computepb.AggregatedListInstancesRequest{Project: projectID}
	it := c.AggregatedList(ctx, req)
	fmt.Println("Compute Engine Instances:")
	for {
		pair, err := it.Next()
		if err != nil {
			break // End of list
		}
		if pair.Value.Instances != nil {
			for _, instance := range pair.Value.Instances {
				// Print various instance details (expand as needed)
				fmt.Printf("- %s (Zone: %s, Machine Type: %s, Status: %s, Creation Timestamp: %s)\n",
					instance.GetName(), instance.GetZone(), instance.GetMachineType(), instance.GetStatus(), instance.GetCreationTimestamp())
			}
		}
	}
	if err != nil {
		log.Fatalf("Failed to list instances: %v", err)
	}
}

func listStorageBuckets(ctx context.Context, projectID string) error {
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer storageClient.Close()

	it := storageClient.Buckets(ctx, projectID)
	fmt.Println("Cloud Storage Buckets:")
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		// Print various bucket details (expand as needed)
		fmt.Printf("- %s (Location: %s, Storage Class: %s, Created: %v)\n",
			attrs.Name, attrs.Location, attrs.StorageClass, attrs.Created)
	}
	return nil
}

func listVPCNetworks(ctx context.Context, projectID string) error {
	c, err := compute.NewNetworksRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Networks client: %v", err)
	}
	defer c.Close()

	req := &computepb.ListNetworksRequest{Project: projectID}
	it := c.List(ctx, req)
	fmt.Println("VPC Networks:")
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Failed to list networks: %v", err)
		}
		// Print various network details (expand as needed)
		fmt.Printf("- %s (Auto Create Subnetworks: %t, Subnetworks: %v, Peerings: %v, IPv4 Range: %s, Gateway IPv4: %s, Creation Timestamp: %s)\n",
			resp.GetName(), resp.GetAutoCreateSubnetworks(), resp.GetSubnetworks(), resp.GetPeerings(), resp.GetIPv4Range(), resp.GetGatewayIPv4(), resp.GetCreationTimestamp())
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")

	listInstances(ctx, projectID)

	if err := listStorageBuckets(ctx, projectID); err != nil {
		log.Printf("Error listing Cloud Storage buckets: %v", err)
	}

	if err := listVPCNetworks(ctx, projectID); err != nil {
		log.Printf("Error listing VPC Networks: %v", err)
	}

	listProjectIamPolicies(ctx, projectID)
}
