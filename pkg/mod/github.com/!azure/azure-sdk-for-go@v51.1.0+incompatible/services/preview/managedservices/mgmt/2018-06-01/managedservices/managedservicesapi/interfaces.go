package managedservicesapi

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

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/managedservices/mgmt/2018-06-01/managedservices"
	"github.com/Azure/go-autorest/autorest"
)

// RegistrationDefinitionsClientAPI contains the set of methods on the RegistrationDefinitionsClient type.
type RegistrationDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string, requestBody managedservices.RegistrationDefinition) (result managedservices.RegistrationDefinition, err error)
	Delete(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, registrationDefinitionID string, APIVersion string) (result managedservices.RegistrationDefinition, err error)
	List(ctx context.Context, scope string, APIVersion string) (result managedservices.RegistrationDefinitionListPage, err error)
	ListComplete(ctx context.Context, scope string, APIVersion string) (result managedservices.RegistrationDefinitionListIterator, err error)
}

var _ RegistrationDefinitionsClientAPI = (*managedservices.RegistrationDefinitionsClient)(nil)

// RegistrationAssignmentsClientAPI contains the set of methods on the RegistrationAssignmentsClient type.
type RegistrationAssignmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, requestBody managedservices.RegistrationAssignment) (result managedservices.RegistrationAssignment, err error)
	Delete(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, expandRegistrationDefinition *bool) (result managedservices.RegistrationAssignment, err error)
	List(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (result managedservices.RegistrationAssignmentListPage, err error)
	ListComplete(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (result managedservices.RegistrationAssignmentListIterator, err error)
}

var _ RegistrationAssignmentsClientAPI = (*managedservices.RegistrationAssignmentsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, APIVersion string) (result managedservices.OperationList, err error)
}

var _ OperationsClientAPI = (*managedservices.OperationsClient)(nil)
