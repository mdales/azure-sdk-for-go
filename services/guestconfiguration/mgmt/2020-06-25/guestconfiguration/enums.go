package guestconfiguration

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionAfterReboot enumerates the values for action after reboot.
type ActionAfterReboot string

const (
	// ContinueConfiguration ...
	ContinueConfiguration ActionAfterReboot = "ContinueConfiguration"
	// StopConfiguration ...
	StopConfiguration ActionAfterReboot = "StopConfiguration"
)

// PossibleActionAfterRebootValues returns an array of possible values for the ActionAfterReboot const type.
func PossibleActionAfterRebootValues() []ActionAfterReboot {
	return []ActionAfterReboot{ContinueConfiguration, StopConfiguration}
}

// AllowModuleOverwrite enumerates the values for allow module overwrite.
type AllowModuleOverwrite string

const (
	// False ...
	False AllowModuleOverwrite = "False"
	// True ...
	True AllowModuleOverwrite = "True"
)

// PossibleAllowModuleOverwriteValues returns an array of possible values for the AllowModuleOverwrite const type.
func PossibleAllowModuleOverwriteValues() []AllowModuleOverwrite {
	return []AllowModuleOverwrite{False, True}
}

// ComplianceStatus enumerates the values for compliance status.
type ComplianceStatus string

const (
	// Compliant ...
	Compliant ComplianceStatus = "Compliant"
	// NonCompliant ...
	NonCompliant ComplianceStatus = "NonCompliant"
	// Pending ...
	Pending ComplianceStatus = "Pending"
)

// PossibleComplianceStatusValues returns an array of possible values for the ComplianceStatus const type.
func PossibleComplianceStatusValues() []ComplianceStatus {
	return []ComplianceStatus{Compliant, NonCompliant, Pending}
}

// ConfigurationMode enumerates the values for configuration mode.
type ConfigurationMode string

const (
	// ApplyAndAutoCorrect ...
	ApplyAndAutoCorrect ConfigurationMode = "ApplyAndAutoCorrect"
	// ApplyAndMonitor ...
	ApplyAndMonitor ConfigurationMode = "ApplyAndMonitor"
	// ApplyOnly ...
	ApplyOnly ConfigurationMode = "ApplyOnly"
)

// PossibleConfigurationModeValues returns an array of possible values for the ConfigurationMode const type.
func PossibleConfigurationModeValues() []ConfigurationMode {
	return []ConfigurationMode{ApplyAndAutoCorrect, ApplyAndMonitor, ApplyOnly}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// DSC ...
	DSC Kind = "DSC"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{DSC}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Created ...
	Created ProvisioningState = "Created"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Created, Failed, Succeeded}
}

// RebootIfNeeded enumerates the values for reboot if needed.
type RebootIfNeeded string

const (
	// RebootIfNeededFalse ...
	RebootIfNeededFalse RebootIfNeeded = "False"
	// RebootIfNeededTrue ...
	RebootIfNeededTrue RebootIfNeeded = "True"
)

// PossibleRebootIfNeededValues returns an array of possible values for the RebootIfNeeded const type.
func PossibleRebootIfNeededValues() []RebootIfNeeded {
	return []RebootIfNeeded{RebootIfNeededFalse, RebootIfNeededTrue}
}

// Type enumerates the values for type.
type Type string

const (
	// Consistency ...
	Consistency Type = "Consistency"
	// Initial ...
	Initial Type = "Initial"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{Consistency, Initial}
}