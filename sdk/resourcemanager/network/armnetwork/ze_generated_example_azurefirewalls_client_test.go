//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallDelete.json
func ExampleAzureFirewallsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<azure-firewall-name>",
		&armnetwork.AzureFirewallsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallGet.json
func ExampleAzureFirewallsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<azure-firewall-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallPut.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<azure-firewall-name>",
		armnetwork.AzureFirewall{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armnetwork.AzureFirewallPropertiesFormat{
				ApplicationRuleCollections: []*armnetwork.AzureFirewallApplicationRuleCollection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallApplicationRuleCollectionPropertiesFormat{
							Action: &armnetwork.AzureFirewallRCAction{
								Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
							},
							Priority: to.Ptr[int32](110),
							Rules: []*armnetwork.AzureFirewallApplicationRule{
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									Protocols: []*armnetwork.AzureFirewallApplicationRuleProtocol{
										{
											Port:         to.Ptr[int32](443),
											ProtocolType: to.Ptr(armnetwork.AzureFirewallApplicationRuleProtocolTypeHTTPS),
										}},
									SourceAddresses: []*string{
										to.Ptr("216.58.216.164"),
										to.Ptr("10.0.0.0/24")},
									TargetFqdns: []*string{
										to.Ptr("www.test.com")},
								}},
						},
					}},
				IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
							PublicIPAddress: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Subnet: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				NatRuleCollections: []*armnetwork.AzureFirewallNatRuleCollection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallNatRuleCollectionProperties{
							Action: &armnetwork.AzureFirewallNatRCAction{
								Type: to.Ptr(armnetwork.AzureFirewallNatRCActionTypeDnat),
							},
							Priority: to.Ptr[int32](112),
							Rules: []*armnetwork.AzureFirewallNatRule{
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationAddresses: []*string{
										to.Ptr("1.2.3.4")},
									DestinationPorts: []*string{
										to.Ptr("443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("*")},
									TranslatedAddress: to.Ptr("<translated-address>"),
									TranslatedPort:    to.Ptr("<translated-port>"),
								},
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationAddresses: []*string{
										to.Ptr("1.2.3.4")},
									DestinationPorts: []*string{
										to.Ptr("80")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("*")},
									TranslatedFqdn: to.Ptr("<translated-fqdn>"),
									TranslatedPort: to.Ptr("<translated-port>"),
								}},
						},
					}},
				NetworkRuleCollections: []*armnetwork.AzureFirewallNetworkRuleCollection{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.AzureFirewallNetworkRuleCollectionPropertiesFormat{
							Action: &armnetwork.AzureFirewallRCAction{
								Type: to.Ptr(armnetwork.AzureFirewallRCActionTypeDeny),
							},
							Priority: to.Ptr[int32](112),
							Rules: []*armnetwork.AzureFirewallNetworkRule{
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationAddresses: []*string{
										to.Ptr("*")},
									DestinationPorts: []*string{
										to.Ptr("443-444"),
										to.Ptr("8443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("192.168.1.1-192.168.1.12"),
										to.Ptr("10.1.4.12-10.1.4.255")},
								},
								{
									Name:        to.Ptr("<name>"),
									Description: to.Ptr("<description>"),
									DestinationFqdns: []*string{
										to.Ptr("www.amazon.com")},
									DestinationPorts: []*string{
										to.Ptr("443-444"),
										to.Ptr("8443")},
									Protocols: []*armnetwork.AzureFirewallNetworkRuleProtocol{
										to.Ptr(armnetwork.AzureFirewallNetworkRuleProtocolTCP)},
									SourceAddresses: []*string{
										to.Ptr("10.2.4.12-10.2.4.255")},
								}},
						},
					}},
				SKU: &armnetwork.AzureFirewallSKU{
					Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWVnet),
					Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
				},
				ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
			},
			Zones: []*string{},
		},
		&armnetwork.AzureFirewallsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallUpdateTags.json
func ExampleAzureFirewallsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateTags(ctx,
		"<resource-group-name>",
		"<azure-firewall-name>",
		armnetwork.TagsObject{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		&armnetwork.AzureFirewallsClientBeginUpdateTagsOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallListByResourceGroup.json
func ExampleAzureFirewallsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/AzureFirewallListBySubscription.json
func ExampleAzureFirewallsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewAzureFirewallsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAllPager(nil)
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
