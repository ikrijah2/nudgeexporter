package avsapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/avs/mgmt/2020-07-17-preview/avs"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result avs.OperationListPage, err error)
	ListComplete(ctx context.Context) (result avs.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*avs.OperationsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	CheckQuotaAvailability(ctx context.Context, location string) (result avs.Quota, err error)
	CheckTrialAvailability(ctx context.Context, location string) (result avs.Trial, err error)
}

var _ LocationsClientAPI = (*avs.LocationsClient)(nil)

// PrivateCloudsClientAPI contains the set of methods on the PrivateCloudsClient type.
type PrivateCloudsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud avs.PrivateCloud) (result avs.PrivateCloudsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.PrivateCloudsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.PrivateCloud, err error)
	List(ctx context.Context, resourceGroupName string) (result avs.PrivateCloudListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result avs.PrivateCloudListIterator, err error)
	ListAdminCredentials(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.AdminCredentials, err error)
	ListInSubscription(ctx context.Context) (result avs.PrivateCloudListPage, err error)
	ListInSubscriptionComplete(ctx context.Context) (result avs.PrivateCloudListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate avs.PrivateCloudUpdate) (result avs.PrivateCloudsUpdateFuture, err error)
}

var _ PrivateCloudsClientAPI = (*avs.PrivateCloudsClient)(nil)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, cluster avs.Cluster) (result avs.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string) (result avs.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string) (result avs.Cluster, err error)
	List(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.ClusterListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.ClusterListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, clusterUpdate avs.ClusterUpdate) (result avs.ClustersUpdateFuture, err error)
}

var _ ClustersClientAPI = (*avs.ClustersClient)(nil)

// HcxEnterpriseSitesClientAPI contains the set of methods on the HcxEnterpriseSitesClient type.
type HcxEnterpriseSitesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, hcxEnterpriseSite avs.HcxEnterpriseSite) (result avs.HcxEnterpriseSite, err error)
	Delete(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string) (result avs.HcxEnterpriseSite, err error)
	List(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.HcxEnterpriseSiteListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.HcxEnterpriseSiteListIterator, err error)
}

var _ HcxEnterpriseSitesClientAPI = (*avs.HcxEnterpriseSitesClient)(nil)

// AuthorizationsClientAPI contains the set of methods on the AuthorizationsClient type.
type AuthorizationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization avs.ExpressRouteAuthorization) (result avs.AuthorizationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string) (result avs.AuthorizationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string) (result avs.ExpressRouteAuthorization, err error)
	List(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.ExpressRouteAuthorizationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.ExpressRouteAuthorizationListIterator, err error)
}

var _ AuthorizationsClientAPI = (*avs.AuthorizationsClient)(nil)

// GlobalReachConnectionsClientAPI contains the set of methods on the GlobalReachConnectionsClient type.
type GlobalReachConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, globalReachConnection avs.GlobalReachConnection) (result avs.GlobalReachConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string) (result avs.GlobalReachConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string) (result avs.GlobalReachConnection, err error)
	List(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.GlobalReachConnectionListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.GlobalReachConnectionListIterator, err error)
}

var _ GlobalReachConnectionsClientAPI = (*avs.GlobalReachConnectionsClient)(nil)

// WorkloadNetworksClientAPI contains the set of methods on the WorkloadNetworksClient type.
type WorkloadNetworksClientAPI interface {
	CreateDhcp(ctx context.Context, resourceGroupName string, privateCloudName string, dhcpID string, workloadNetworkDhcp avs.WorkloadNetworkDhcp) (result avs.WorkloadNetworksCreateDhcpFuture, err error)
	CreatePortMirroring(ctx context.Context, resourceGroupName string, privateCloudName string, portMirroringID string, workloadNetworkPortMirroring avs.WorkloadNetworkPortMirroring) (result avs.WorkloadNetworksCreatePortMirroringFuture, err error)
	CreateSegments(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string, workloadNetworkSegment avs.WorkloadNetworkSegment) (result avs.WorkloadNetworksCreateSegmentsFuture, err error)
	CreateVMGroup(ctx context.Context, resourceGroupName string, privateCloudName string, VMGroupID string, workloadNetworkVMGroup avs.WorkloadNetworkVMGroup) (result avs.WorkloadNetworksCreateVMGroupFuture, err error)
	DeleteDhcp(ctx context.Context, resourceGroupName string, privateCloudName string, dhcpID string) (result avs.WorkloadNetworksDeleteDhcpFuture, err error)
	DeletePortMirroring(ctx context.Context, resourceGroupName string, portMirroringID string, privateCloudName string) (result avs.WorkloadNetworksDeletePortMirroringFuture, err error)
	DeleteSegment(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string) (result avs.WorkloadNetworksDeleteSegmentFuture, err error)
	DeleteVMGroup(ctx context.Context, resourceGroupName string, VMGroupID string, privateCloudName string) (result avs.WorkloadNetworksDeleteVMGroupFuture, err error)
	GetDhcp(ctx context.Context, resourceGroupName string, dhcpID string, privateCloudName string) (result avs.WorkloadNetworkDhcp, err error)
	GetGateway(ctx context.Context, resourceGroupName string, privateCloudName string, gatewayID string) (result avs.WorkloadNetworkGateway, err error)
	GetPortMirroring(ctx context.Context, resourceGroupName string, privateCloudName string, portMirroringID string) (result avs.WorkloadNetworkPortMirroring, err error)
	GetSegment(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string) (result avs.WorkloadNetworkSegment, err error)
	GetVirtualMachine(ctx context.Context, resourceGroupName string, privateCloudName string, virtualMachineID string) (result avs.WorkloadNetworkVirtualMachine, err error)
	GetVMGroup(ctx context.Context, resourceGroupName string, privateCloudName string, VMGroupID string) (result avs.WorkloadNetworkVMGroup, err error)
	ListDhcp(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkDhcpListPage, err error)
	ListDhcpComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkDhcpListIterator, err error)
	ListGateways(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkGatewayListPage, err error)
	ListGatewaysComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkGatewayListIterator, err error)
	ListPortMirroring(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkPortMirroringListPage, err error)
	ListPortMirroringComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkPortMirroringListIterator, err error)
	ListSegments(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkSegmentsListPage, err error)
	ListSegmentsComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkSegmentsListIterator, err error)
	ListVirtualMachines(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkVirtualMachinesListPage, err error)
	ListVirtualMachinesComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkVirtualMachinesListIterator, err error)
	ListVMGroups(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkVMGroupsListPage, err error)
	ListVMGroupsComplete(ctx context.Context, resourceGroupName string, privateCloudName string) (result avs.WorkloadNetworkVMGroupsListIterator, err error)
	UpdateDhcp(ctx context.Context, resourceGroupName string, privateCloudName string, dhcpID string, workloadNetworkDhcp avs.WorkloadNetworkDhcp) (result avs.WorkloadNetworksUpdateDhcpFuture, err error)
	UpdatePortMirroring(ctx context.Context, resourceGroupName string, privateCloudName string, portMirroringID string, workloadNetworkPortMirroring avs.WorkloadNetworkPortMirroring) (result avs.WorkloadNetworksUpdatePortMirroringFuture, err error)
	UpdateSegments(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string, workloadNetworkSegment avs.WorkloadNetworkSegment) (result avs.WorkloadNetworksUpdateSegmentsFuture, err error)
	UpdateVMGroup(ctx context.Context, resourceGroupName string, privateCloudName string, VMGroupID string, workloadNetworkVMGroup avs.WorkloadNetworkVMGroup) (result avs.WorkloadNetworksUpdateVMGroupFuture, err error)
}

var _ WorkloadNetworksClientAPI = (*avs.WorkloadNetworksClient)(nil)
