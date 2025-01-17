//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

// ConfigurationsClientCreateInResourceGroupResponse contains the response from method ConfigurationsClient.CreateInResourceGroup.
type ConfigurationsClientCreateInResourceGroupResponse struct {
	ConfigData
}

// ConfigurationsClientCreateInSubscriptionResponse contains the response from method ConfigurationsClient.CreateInSubscription.
type ConfigurationsClientCreateInSubscriptionResponse struct {
	ConfigData
}

// ConfigurationsClientListByResourceGroupResponse contains the response from method ConfigurationsClient.ListByResourceGroup.
type ConfigurationsClientListByResourceGroupResponse struct {
	ConfigurationListResult
}

// ConfigurationsClientListBySubscriptionResponse contains the response from method ConfigurationsClient.ListBySubscription.
type ConfigurationsClientListBySubscriptionResponse struct {
	ConfigurationListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationEntityListResult
}

// RecommendationMetadataClientGetResponse contains the response from method RecommendationMetadataClient.Get.
type RecommendationMetadataClientGetResponse struct {
	MetadataEntity
}

// RecommendationMetadataClientListResponse contains the response from method RecommendationMetadataClient.List.
type RecommendationMetadataClientListResponse struct {
	MetadataEntityListResult
}

// RecommendationsClientGenerateResponse contains the response from method RecommendationsClient.Generate.
type RecommendationsClientGenerateResponse struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *string
}

// RecommendationsClientGetGenerateStatusResponse contains the response from method RecommendationsClient.GetGenerateStatus.
type RecommendationsClientGetGenerateStatusResponse struct {
	// placeholder for future response values
}

// RecommendationsClientGetResponse contains the response from method RecommendationsClient.Get.
type RecommendationsClientGetResponse struct {
	ResourceRecommendationBase
}

// RecommendationsClientListResponse contains the response from method RecommendationsClient.List.
type RecommendationsClientListResponse struct {
	ResourceRecommendationBaseListResult
}

// SuppressionsClientCreateResponse contains the response from method SuppressionsClient.Create.
type SuppressionsClientCreateResponse struct {
	SuppressionContract
}

// SuppressionsClientDeleteResponse contains the response from method SuppressionsClient.Delete.
type SuppressionsClientDeleteResponse struct {
	// placeholder for future response values
}

// SuppressionsClientGetResponse contains the response from method SuppressionsClient.Get.
type SuppressionsClientGetResponse struct {
	SuppressionContract
}

// SuppressionsClientListResponse contains the response from method SuppressionsClient.List.
type SuppressionsClientListResponse struct {
	SuppressionContractListResult
}
