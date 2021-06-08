package operationalinsightsapi

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
	"github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2020-08-01/operationalinsights"
	"github.com/Azure/go-autorest/autorest"
)

// DataExportsClientAPI contains the set of methods on the DataExportsClient type.
type DataExportsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, parameters operationalinsights.DataExport) (result operationalinsights.DataExport, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string) (result operationalinsights.DataExport, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.DataExportListResult, err error)
}

var _ DataExportsClientAPI = (*operationalinsights.DataExportsClient)(nil)

// DataSourcesClientAPI contains the set of methods on the DataSourcesClient type.
type DataSourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, parameters operationalinsights.DataSource) (result operationalinsights.DataSource, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string) (result operationalinsights.DataSource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, filter string, skiptoken string) (result operationalinsights.DataSourceListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, skiptoken string) (result operationalinsights.DataSourceListResultIterator, err error)
}

var _ DataSourcesClientAPI = (*operationalinsights.DataSourcesClient)(nil)

// IntelligencePacksClientAPI contains the set of methods on the IntelligencePacksClient type.
type IntelligencePacksClientAPI interface {
	Disable(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string) (result autorest.Response, err error)
	Enable(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string) (result autorest.Response, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.ListIntelligencePack, err error)
}

var _ IntelligencePacksClientAPI = (*operationalinsights.IntelligencePacksClient)(nil)

// LinkedServicesClientAPI contains the set of methods on the LinkedServicesClient type.
type LinkedServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters operationalinsights.LinkedService) (result operationalinsights.LinkedServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (result operationalinsights.LinkedServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (result operationalinsights.LinkedService, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.LinkedServiceListResult, err error)
}

var _ LinkedServicesClientAPI = (*operationalinsights.LinkedServicesClient)(nil)

// LinkedStorageAccountsClientAPI contains the set of methods on the LinkedStorageAccountsClient type.
type LinkedStorageAccountsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType operationalinsights.DataSourceType, parameters operationalinsights.LinkedStorageAccountsResource) (result operationalinsights.LinkedStorageAccountsResource, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType operationalinsights.DataSourceType) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType operationalinsights.DataSourceType) (result operationalinsights.LinkedStorageAccountsResource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.LinkedStorageAccountsListResult, err error)
}

var _ LinkedStorageAccountsClientAPI = (*operationalinsights.LinkedStorageAccountsClient)(nil)

// ManagementGroupsClientAPI contains the set of methods on the ManagementGroupsClient type.
type ManagementGroupsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.WorkspaceListManagementGroupsResult, err error)
}

var _ ManagementGroupsClientAPI = (*operationalinsights.ManagementGroupsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result operationalinsights.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*operationalinsights.OperationsClient)(nil)

// OperationStatusesClientAPI contains the set of methods on the OperationStatusesClient type.
type OperationStatusesClientAPI interface {
	Get(ctx context.Context, location string, asyncOperationID string) (result operationalinsights.OperationStatus, err error)
}

var _ OperationStatusesClientAPI = (*operationalinsights.OperationStatusesClient)(nil)

// SharedKeysClientAPI contains the set of methods on the SharedKeysClient type.
type SharedKeysClientAPI interface {
	GetSharedKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SharedKeys, err error)
	Regenerate(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SharedKeys, err error)
}

var _ SharedKeysClientAPI = (*operationalinsights.SharedKeysClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.WorkspaceListUsagesResult, err error)
}

var _ UsagesClientAPI = (*operationalinsights.UsagesClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (result operationalinsights.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, force *bool) (result operationalinsights.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.Workspace, err error)
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.WorkspacePatch) (result operationalinsights.Workspace, err error)
}

var _ WorkspacesClientAPI = (*operationalinsights.WorkspacesClient)(nil)

// DeletedWorkspacesClientAPI contains the set of methods on the DeletedWorkspacesClient type.
type DeletedWorkspacesClientAPI interface {
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
}

var _ DeletedWorkspacesClientAPI = (*operationalinsights.DeletedWorkspacesClient)(nil)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters operationalinsights.Cluster) (result operationalinsights.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result operationalinsights.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result operationalinsights.Cluster, err error)
	List(ctx context.Context) (result operationalinsights.ClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.ClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.ClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result operationalinsights.ClusterListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters operationalinsights.ClusterPatch) (result operationalinsights.Cluster, err error)
}

var _ ClustersClientAPI = (*operationalinsights.ClustersClient)(nil)

// StorageInsightConfigsClientAPI contains the set of methods on the StorageInsightConfigsClient type.
type StorageInsightConfigsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters operationalinsights.StorageInsight) (result operationalinsights.StorageInsight, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string) (result operationalinsights.StorageInsight, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.StorageInsightListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.StorageInsightListResultIterator, err error)
}

var _ StorageInsightConfigsClientAPI = (*operationalinsights.StorageInsightConfigsClient)(nil)

// SavedSearchesClientAPI contains the set of methods on the SavedSearchesClient type.
type SavedSearchesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters operationalinsights.SavedSearch) (result operationalinsights.SavedSearch, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result operationalinsights.SavedSearch, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SavedSearchesListResult, err error)
}

var _ SavedSearchesClientAPI = (*operationalinsights.SavedSearchesClient)(nil)

// AvailableServiceTiersClientAPI contains the set of methods on the AvailableServiceTiersClient type.
type AvailableServiceTiersClientAPI interface {
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.ListAvailableServiceTier, err error)
}

var _ AvailableServiceTiersClientAPI = (*operationalinsights.AvailableServiceTiersClient)(nil)

// GatewaysClientAPI contains the set of methods on the GatewaysClient type.
type GatewaysClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, gatewayID string) (result autorest.Response, err error)
}

var _ GatewaysClientAPI = (*operationalinsights.GatewaysClient)(nil)

// SchemaClientAPI contains the set of methods on the SchemaClient type.
type SchemaClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SearchGetSchemaResponse, err error)
}

var _ SchemaClientAPI = (*operationalinsights.SchemaClient)(nil)

// WorkspacePurgeClientAPI contains the set of methods on the WorkspacePurgeClient type.
type WorkspacePurgeClientAPI interface {
	GetPurgeStatus(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string) (result operationalinsights.WorkspacePurgeStatusResponse, err error)
	Purge(ctx context.Context, resourceGroupName string, workspaceName string, body operationalinsights.WorkspacePurgeBody) (result operationalinsights.WorkspacePurgeResponse, err error)
}

var _ WorkspacePurgeClientAPI = (*operationalinsights.WorkspacePurgeClient)(nil)

// TablesClientAPI contains the set of methods on the TablesClient type.
type TablesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, tableName string) (result operationalinsights.Table, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.TablesListResult, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, tableName string, parameters operationalinsights.Table) (result operationalinsights.Table, err error)
}

var _ TablesClientAPI = (*operationalinsights.TablesClient)(nil)
