//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_List.json
func ExamplePrivateCloudsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_ListInSubscription.json
func ExamplePrivateCloudsClient_NewListInSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListInSubscriptionPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Get.json
func ExamplePrivateCloudsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_CreateOrUpdate.json
func ExamplePrivateCloudsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		armavs.PrivateCloud{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Identity: &armavs.PrivateCloudIdentity{
				Type: to.Ptr(armavs.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armavs.PrivateCloudProperties{
				ManagementCluster: &armavs.ManagementCluster{
					ClusterSize: to.Ptr[int32](4),
				},
				NetworkBlock: to.Ptr("<network-block>"),
			},
			SKU: &armavs.SKU{
				Name: to.Ptr("<name>"),
			},
		},
		&armavs.PrivateCloudsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Update.json
func ExamplePrivateCloudsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		armavs.PrivateCloudUpdate{
			Identity: &armavs.PrivateCloudIdentity{
				Type: to.Ptr(armavs.ResourceIdentityTypeNone),
			},
			Properties: &armavs.PrivateCloudUpdateProperties{
				Encryption: &armavs.Encryption{
					KeyVaultProperties: &armavs.EncryptionKeyVaultProperties{
						KeyName:     to.Ptr("<key-name>"),
						KeyVaultURL: to.Ptr("<key-vault-url>"),
						KeyVersion:  to.Ptr("<key-version>"),
					},
					Status: to.Ptr(armavs.EncryptionStateEnabled),
				},
				ManagementCluster: &armavs.ManagementCluster{
					ClusterSize: to.Ptr[int32](4),
				},
			},
		},
		&armavs.PrivateCloudsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Delete.json
func ExamplePrivateCloudsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		&armavs.PrivateCloudsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_RotateVcenterPassword.json
func ExamplePrivateCloudsClient_BeginRotateVcenterPassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRotateVcenterPassword(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		&armavs.PrivateCloudsClientBeginRotateVcenterPasswordOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_RotateNsxtPassword.json
func ExamplePrivateCloudsClient_BeginRotateNsxtPassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRotateNsxtPassword(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		&armavs.PrivateCloudsClientBeginRotateNsxtPasswordOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_ListAdminCredentials.json
func ExamplePrivateCloudsClient_ListAdminCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewPrivateCloudsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListAdminCredentials(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
