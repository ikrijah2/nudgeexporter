// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package securityinsight

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/securityinsight/mgmt/v1.0/securityinsight"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlertRuleKind = original.AlertRuleKind

const (
	Fusion                            AlertRuleKind = original.Fusion
	MicrosoftSecurityIncidentCreation AlertRuleKind = original.MicrosoftSecurityIncidentCreation
	Scheduled                         AlertRuleKind = original.Scheduled
)

type AlertSeverity = original.AlertSeverity

const (
	High          AlertSeverity = original.High
	Informational AlertSeverity = original.Informational
	Low           AlertSeverity = original.Low
	Medium        AlertSeverity = original.Medium
)

type AttackTactic = original.AttackTactic

const (
	Collection          AttackTactic = original.Collection
	CommandAndControl   AttackTactic = original.CommandAndControl
	CredentialAccess    AttackTactic = original.CredentialAccess
	DefenseEvasion      AttackTactic = original.DefenseEvasion
	Discovery           AttackTactic = original.Discovery
	Execution           AttackTactic = original.Execution
	Exfiltration        AttackTactic = original.Exfiltration
	Impact              AttackTactic = original.Impact
	InitialAccess       AttackTactic = original.InitialAccess
	LateralMovement     AttackTactic = original.LateralMovement
	Persistence         AttackTactic = original.Persistence
	PrivilegeEscalation AttackTactic = original.PrivilegeEscalation
)

type CaseSeverity = original.CaseSeverity

const (
	CaseSeverityCritical      CaseSeverity = original.CaseSeverityCritical
	CaseSeverityHigh          CaseSeverity = original.CaseSeverityHigh
	CaseSeverityInformational CaseSeverity = original.CaseSeverityInformational
	CaseSeverityLow           CaseSeverity = original.CaseSeverityLow
	CaseSeverityMedium        CaseSeverity = original.CaseSeverityMedium
)

type DataConnectorKind = original.DataConnectorKind

const (
	DataConnectorKindAmazonWebServicesCloudTrail               DataConnectorKind = original.DataConnectorKindAmazonWebServicesCloudTrail
	DataConnectorKindAzureActiveDirectory                      DataConnectorKind = original.DataConnectorKindAzureActiveDirectory
	DataConnectorKindAzureAdvancedThreatProtection             DataConnectorKind = original.DataConnectorKindAzureAdvancedThreatProtection
	DataConnectorKindAzureSecurityCenter                       DataConnectorKind = original.DataConnectorKindAzureSecurityCenter
	DataConnectorKindMicrosoftCloudAppSecurity                 DataConnectorKind = original.DataConnectorKindMicrosoftCloudAppSecurity
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection DataConnectorKind = original.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection
	DataConnectorKindOffice365                                 DataConnectorKind = original.DataConnectorKindOffice365
	DataConnectorKindThreatIntelligence                        DataConnectorKind = original.DataConnectorKindThreatIntelligence
)

type DataTypeState = original.DataTypeState

const (
	Disabled DataTypeState = original.Disabled
	Enabled  DataTypeState = original.Enabled
)

type IncidentClassification = original.IncidentClassification

const (
	BenignPositive IncidentClassification = original.BenignPositive
	FalsePositive  IncidentClassification = original.FalsePositive
	TruePositive   IncidentClassification = original.TruePositive
	Undetermined   IncidentClassification = original.Undetermined
)

type IncidentClassificationReason = original.IncidentClassificationReason

const (
	InaccurateData        IncidentClassificationReason = original.InaccurateData
	IncorrectAlertLogic   IncidentClassificationReason = original.IncorrectAlertLogic
	SuspiciousActivity    IncidentClassificationReason = original.SuspiciousActivity
	SuspiciousButExpected IncidentClassificationReason = original.SuspiciousButExpected
)

type IncidentLabelType = original.IncidentLabelType

const (
	System IncidentLabelType = original.System
	User   IncidentLabelType = original.User
)

type IncidentSeverity = original.IncidentSeverity

const (
	IncidentSeverityHigh          IncidentSeverity = original.IncidentSeverityHigh
	IncidentSeverityInformational IncidentSeverity = original.IncidentSeverityInformational
	IncidentSeverityLow           IncidentSeverity = original.IncidentSeverityLow
	IncidentSeverityMedium        IncidentSeverity = original.IncidentSeverityMedium
)

type IncidentStatus = original.IncidentStatus

const (
	IncidentStatusActive IncidentStatus = original.IncidentStatusActive
	IncidentStatusClosed IncidentStatus = original.IncidentStatusClosed
	IncidentStatusNew    IncidentStatus = original.IncidentStatusNew
)

type Kind = original.Kind

const (
	KindAlertRule                         Kind = original.KindAlertRule
	KindFusion                            Kind = original.KindFusion
	KindMicrosoftSecurityIncidentCreation Kind = original.KindMicrosoftSecurityIncidentCreation
	KindScheduled                         Kind = original.KindScheduled
)

type KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplate

const (
	KindBasicAlertRuleTemplateKindAlertRuleTemplate                 KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindAlertRuleTemplate
	KindBasicAlertRuleTemplateKindFusion                            KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindFusion
	KindBasicAlertRuleTemplateKindMicrosoftSecurityIncidentCreation KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindMicrosoftSecurityIncidentCreation
	KindBasicAlertRuleTemplateKindScheduled                         KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindScheduled
)

type KindBasicDataConnector = original.KindBasicDataConnector

const (
	KindAmazonWebServicesCloudTrail               KindBasicDataConnector = original.KindAmazonWebServicesCloudTrail
	KindAzureActiveDirectory                      KindBasicDataConnector = original.KindAzureActiveDirectory
	KindAzureAdvancedThreatProtection             KindBasicDataConnector = original.KindAzureAdvancedThreatProtection
	KindAzureSecurityCenter                       KindBasicDataConnector = original.KindAzureSecurityCenter
	KindDataConnector                             KindBasicDataConnector = original.KindDataConnector
	KindMicrosoftCloudAppSecurity                 KindBasicDataConnector = original.KindMicrosoftCloudAppSecurity
	KindMicrosoftDefenderAdvancedThreatProtection KindBasicDataConnector = original.KindMicrosoftDefenderAdvancedThreatProtection
	KindOffice365                                 KindBasicDataConnector = original.KindOffice365
	KindThreatIntelligence                        KindBasicDataConnector = original.KindThreatIntelligence
)

type KindBasicSettings = original.KindBasicSettings

const (
	KindSettings       KindBasicSettings = original.KindSettings
	KindToggleSettings KindBasicSettings = original.KindToggleSettings
	KindUebaSettings   KindBasicSettings = original.KindUebaSettings
)

type LicenseStatus = original.LicenseStatus

const (
	LicenseStatusDisabled LicenseStatus = original.LicenseStatusDisabled
	LicenseStatusEnabled  LicenseStatus = original.LicenseStatusEnabled
)

type MicrosoftSecurityProductName = original.MicrosoftSecurityProductName

const (
	AzureActiveDirectoryIdentityProtection MicrosoftSecurityProductName = original.AzureActiveDirectoryIdentityProtection
	AzureAdvancedThreatProtection          MicrosoftSecurityProductName = original.AzureAdvancedThreatProtection
	AzureSecurityCenter                    MicrosoftSecurityProductName = original.AzureSecurityCenter
	AzureSecurityCenterforIoT              MicrosoftSecurityProductName = original.AzureSecurityCenterforIoT
	MicrosoftCloudAppSecurity              MicrosoftSecurityProductName = original.MicrosoftCloudAppSecurity
)

type SettingKind = original.SettingKind

const (
	SettingKindToggleSettings SettingKind = original.SettingKindToggleSettings
	SettingKindUebaSettings   SettingKind = original.SettingKindUebaSettings
)

type StatusInMcas = original.StatusInMcas

const (
	StatusInMcasDisabled StatusInMcas = original.StatusInMcasDisabled
	StatusInMcasEnabled  StatusInMcas = original.StatusInMcasEnabled
)

type TemplateStatus = original.TemplateStatus

const (
	Available    TemplateStatus = original.Available
	Installed    TemplateStatus = original.Installed
	NotAvailable TemplateStatus = original.NotAvailable
)

type TriggerOperator = original.TriggerOperator

const (
	Equal       TriggerOperator = original.Equal
	GreaterThan TriggerOperator = original.GreaterThan
	LessThan    TriggerOperator = original.LessThan
	NotEqual    TriggerOperator = original.NotEqual
)

type AADDataConnector = original.AADDataConnector
type AADDataConnectorProperties = original.AADDataConnectorProperties
type AATPDataConnector = original.AATPDataConnector
type AATPDataConnectorProperties = original.AATPDataConnectorProperties
type ASCDataConnector = original.ASCDataConnector
type ASCDataConnectorProperties = original.ASCDataConnectorProperties
type ActionPropertiesBase = original.ActionPropertiesBase
type ActionRequest = original.ActionRequest
type ActionRequestProperties = original.ActionRequestProperties
type ActionResponse = original.ActionResponse
type ActionResponseProperties = original.ActionResponseProperties
type ActionsClient = original.ActionsClient
type ActionsList = original.ActionsList
type ActionsListIterator = original.ActionsListIterator
type ActionsListPage = original.ActionsListPage
type AlertRule = original.AlertRule
type AlertRuleModel = original.AlertRuleModel
type AlertRuleTemplate = original.AlertRuleTemplate
type AlertRuleTemplateDataSource = original.AlertRuleTemplateDataSource
type AlertRuleTemplateModel = original.AlertRuleTemplateModel
type AlertRuleTemplatesClient = original.AlertRuleTemplatesClient
type AlertRuleTemplatesList = original.AlertRuleTemplatesList
type AlertRuleTemplatesListIterator = original.AlertRuleTemplatesListIterator
type AlertRuleTemplatesListPage = original.AlertRuleTemplatesListPage
type AlertRulesClient = original.AlertRulesClient
type AlertRulesList = original.AlertRulesList
type AlertRulesListIterator = original.AlertRulesListIterator
type AlertRulesListPage = original.AlertRulesListPage
type AlertsDataTypeOfDataConnector = original.AlertsDataTypeOfDataConnector
type AwsCloudTrailDataConnector = original.AwsCloudTrailDataConnector
type AwsCloudTrailDataConnectorDataTypes = original.AwsCloudTrailDataConnectorDataTypes
type AwsCloudTrailDataConnectorDataTypesLogs = original.AwsCloudTrailDataConnectorDataTypesLogs
type AwsCloudTrailDataConnectorProperties = original.AwsCloudTrailDataConnectorProperties
type BaseClient = original.BaseClient
type BasicAlertRule = original.BasicAlertRule
type BasicAlertRuleTemplate = original.BasicAlertRuleTemplate
type BasicDataConnector = original.BasicDataConnector
type BasicSettings = original.BasicSettings
type Bookmark = original.Bookmark
type BookmarkList = original.BookmarkList
type BookmarkListIterator = original.BookmarkListIterator
type BookmarkListPage = original.BookmarkListPage
type BookmarkProperties = original.BookmarkProperties
type BookmarksClient = original.BookmarksClient
type ClientInfo = original.ClientInfo
type CloudError = original.CloudError
type DataConnector = original.DataConnector
type DataConnectorDataTypeCommon = original.DataConnectorDataTypeCommon
type DataConnectorList = original.DataConnectorList
type DataConnectorListIterator = original.DataConnectorListIterator
type DataConnectorListPage = original.DataConnectorListPage
type DataConnectorModel = original.DataConnectorModel
type DataConnectorTenantID = original.DataConnectorTenantID
type DataConnectorWithAlertsProperties = original.DataConnectorWithAlertsProperties
type DataConnectorsClient = original.DataConnectorsClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type FusionAlertRule = original.FusionAlertRule
type FusionAlertRuleProperties = original.FusionAlertRuleProperties
type FusionAlertRuleTemplate = original.FusionAlertRuleTemplate
type FusionAlertRuleTemplateProperties = original.FusionAlertRuleTemplateProperties
type Incident = original.Incident
type IncidentAdditionalData = original.IncidentAdditionalData
type IncidentComment = original.IncidentComment
type IncidentCommentList = original.IncidentCommentList
type IncidentCommentListIterator = original.IncidentCommentListIterator
type IncidentCommentListPage = original.IncidentCommentListPage
type IncidentCommentProperties = original.IncidentCommentProperties
type IncidentCommentsClient = original.IncidentCommentsClient
type IncidentInfo = original.IncidentInfo
type IncidentLabel = original.IncidentLabel
type IncidentList = original.IncidentList
type IncidentListIterator = original.IncidentListIterator
type IncidentListPage = original.IncidentListPage
type IncidentOwnerInfo = original.IncidentOwnerInfo
type IncidentProperties = original.IncidentProperties
type IncidentsClient = original.IncidentsClient
type MCASDataConnector = original.MCASDataConnector
type MCASDataConnectorDataTypes = original.MCASDataConnectorDataTypes
type MCASDataConnectorProperties = original.MCASDataConnectorProperties
type MDATPDataConnector = original.MDATPDataConnector
type MDATPDataConnectorProperties = original.MDATPDataConnectorProperties
type MicrosoftSecurityIncidentCreationAlertRule = original.MicrosoftSecurityIncidentCreationAlertRule
type MicrosoftSecurityIncidentCreationAlertRuleCommonProperties = original.MicrosoftSecurityIncidentCreationAlertRuleCommonProperties
type MicrosoftSecurityIncidentCreationAlertRuleProperties = original.MicrosoftSecurityIncidentCreationAlertRuleProperties
type MicrosoftSecurityIncidentCreationAlertRuleTemplate = original.MicrosoftSecurityIncidentCreationAlertRuleTemplate
type MicrosoftSecurityIncidentCreationAlertRuleTemplateProperties = original.MicrosoftSecurityIncidentCreationAlertRuleTemplateProperties
type OfficeConsent = original.OfficeConsent
type OfficeConsentList = original.OfficeConsentList
type OfficeConsentProperties = original.OfficeConsentProperties
type OfficeDataConnector = original.OfficeDataConnector
type OfficeDataConnectorDataTypes = original.OfficeDataConnectorDataTypes
type OfficeDataConnectorDataTypesExchange = original.OfficeDataConnectorDataTypesExchange
type OfficeDataConnectorDataTypesSharePoint = original.OfficeDataConnectorDataTypesSharePoint
type OfficeDataConnectorProperties = original.OfficeDataConnectorProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type Resource = original.Resource
type ResourceWithEtag = original.ResourceWithEtag
type ScheduledAlertRule = original.ScheduledAlertRule
type ScheduledAlertRuleCommonProperties = original.ScheduledAlertRuleCommonProperties
type ScheduledAlertRuleProperties = original.ScheduledAlertRuleProperties
type ScheduledAlertRuleTemplate = original.ScheduledAlertRuleTemplate
type ScheduledAlertRuleTemplateProperties = original.ScheduledAlertRuleTemplateProperties
type Settings = original.Settings
type TIDataConnector = original.TIDataConnector
type TIDataConnectorDataTypes = original.TIDataConnectorDataTypes
type TIDataConnectorDataTypesIndicators = original.TIDataConnectorDataTypesIndicators
type TIDataConnectorProperties = original.TIDataConnectorProperties
type ThreatIntelligence = original.ThreatIntelligence
type ToggleSettings = original.ToggleSettings
type ToggleSettingsProperties = original.ToggleSettingsProperties
type UebaSettings = original.UebaSettings
type UebaSettingsProperties = original.UebaSettingsProperties
type UserInfo = original.UserInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActionsClient(subscriptionID string) ActionsClient {
	return original.NewActionsClient(subscriptionID)
}
func NewActionsClientWithBaseURI(baseURI string, subscriptionID string) ActionsClient {
	return original.NewActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewActionsListIterator(page ActionsListPage) ActionsListIterator {
	return original.NewActionsListIterator(page)
}
func NewActionsListPage(cur ActionsList, getNextPage func(context.Context, ActionsList) (ActionsList, error)) ActionsListPage {
	return original.NewActionsListPage(cur, getNextPage)
}
func NewAlertRuleTemplatesClient(subscriptionID string) AlertRuleTemplatesClient {
	return original.NewAlertRuleTemplatesClient(subscriptionID)
}
func NewAlertRuleTemplatesClientWithBaseURI(baseURI string, subscriptionID string) AlertRuleTemplatesClient {
	return original.NewAlertRuleTemplatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRuleTemplatesListIterator(page AlertRuleTemplatesListPage) AlertRuleTemplatesListIterator {
	return original.NewAlertRuleTemplatesListIterator(page)
}
func NewAlertRuleTemplatesListPage(cur AlertRuleTemplatesList, getNextPage func(context.Context, AlertRuleTemplatesList) (AlertRuleTemplatesList, error)) AlertRuleTemplatesListPage {
	return original.NewAlertRuleTemplatesListPage(cur, getNextPage)
}
func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClient(subscriptionID)
}
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRulesListIterator(page AlertRulesListPage) AlertRulesListIterator {
	return original.NewAlertRulesListIterator(page)
}
func NewAlertRulesListPage(cur AlertRulesList, getNextPage func(context.Context, AlertRulesList) (AlertRulesList, error)) AlertRulesListPage {
	return original.NewAlertRulesListPage(cur, getNextPage)
}
func NewBookmarkListIterator(page BookmarkListPage) BookmarkListIterator {
	return original.NewBookmarkListIterator(page)
}
func NewBookmarkListPage(cur BookmarkList, getNextPage func(context.Context, BookmarkList) (BookmarkList, error)) BookmarkListPage {
	return original.NewBookmarkListPage(cur, getNextPage)
}
func NewBookmarksClient(subscriptionID string) BookmarksClient {
	return original.NewBookmarksClient(subscriptionID)
}
func NewBookmarksClientWithBaseURI(baseURI string, subscriptionID string) BookmarksClient {
	return original.NewBookmarksClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataConnectorListIterator(page DataConnectorListPage) DataConnectorListIterator {
	return original.NewDataConnectorListIterator(page)
}
func NewDataConnectorListPage(cur DataConnectorList, getNextPage func(context.Context, DataConnectorList) (DataConnectorList, error)) DataConnectorListPage {
	return original.NewDataConnectorListPage(cur, getNextPage)
}
func NewDataConnectorsClient(subscriptionID string) DataConnectorsClient {
	return original.NewDataConnectorsClient(subscriptionID)
}
func NewDataConnectorsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectorsClient {
	return original.NewDataConnectorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewIncidentCommentListIterator(page IncidentCommentListPage) IncidentCommentListIterator {
	return original.NewIncidentCommentListIterator(page)
}
func NewIncidentCommentListPage(cur IncidentCommentList, getNextPage func(context.Context, IncidentCommentList) (IncidentCommentList, error)) IncidentCommentListPage {
	return original.NewIncidentCommentListPage(cur, getNextPage)
}
func NewIncidentCommentsClient(subscriptionID string) IncidentCommentsClient {
	return original.NewIncidentCommentsClient(subscriptionID)
}
func NewIncidentCommentsClientWithBaseURI(baseURI string, subscriptionID string) IncidentCommentsClient {
	return original.NewIncidentCommentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewIncidentListIterator(page IncidentListPage) IncidentListIterator {
	return original.NewIncidentListIterator(page)
}
func NewIncidentListPage(cur IncidentList, getNextPage func(context.Context, IncidentList) (IncidentList, error)) IncidentListPage {
	return original.NewIncidentListPage(cur, getNextPage)
}
func NewIncidentsClient(subscriptionID string) IncidentsClient {
	return original.NewIncidentsClient(subscriptionID)
}
func NewIncidentsClientWithBaseURI(baseURI string, subscriptionID string) IncidentsClient {
	return original.NewIncidentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(cur OperationsList, getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAlertRuleKindValues() []AlertRuleKind {
	return original.PossibleAlertRuleKindValues()
}
func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}
func PossibleAttackTacticValues() []AttackTactic {
	return original.PossibleAttackTacticValues()
}
func PossibleCaseSeverityValues() []CaseSeverity {
	return original.PossibleCaseSeverityValues()
}
func PossibleDataConnectorKindValues() []DataConnectorKind {
	return original.PossibleDataConnectorKindValues()
}
func PossibleDataTypeStateValues() []DataTypeState {
	return original.PossibleDataTypeStateValues()
}
func PossibleIncidentClassificationReasonValues() []IncidentClassificationReason {
	return original.PossibleIncidentClassificationReasonValues()
}
func PossibleIncidentClassificationValues() []IncidentClassification {
	return original.PossibleIncidentClassificationValues()
}
func PossibleIncidentLabelTypeValues() []IncidentLabelType {
	return original.PossibleIncidentLabelTypeValues()
}
func PossibleIncidentSeverityValues() []IncidentSeverity {
	return original.PossibleIncidentSeverityValues()
}
func PossibleIncidentStatusValues() []IncidentStatus {
	return original.PossibleIncidentStatusValues()
}
func PossibleKindBasicAlertRuleTemplateValues() []KindBasicAlertRuleTemplate {
	return original.PossibleKindBasicAlertRuleTemplateValues()
}
func PossibleKindBasicDataConnectorValues() []KindBasicDataConnector {
	return original.PossibleKindBasicDataConnectorValues()
}
func PossibleKindBasicSettingsValues() []KindBasicSettings {
	return original.PossibleKindBasicSettingsValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLicenseStatusValues() []LicenseStatus {
	return original.PossibleLicenseStatusValues()
}
func PossibleMicrosoftSecurityProductNameValues() []MicrosoftSecurityProductName {
	return original.PossibleMicrosoftSecurityProductNameValues()
}
func PossibleSettingKindValues() []SettingKind {
	return original.PossibleSettingKindValues()
}
func PossibleStatusInMcasValues() []StatusInMcas {
	return original.PossibleStatusInMcasValues()
}
func PossibleTemplateStatusValues() []TemplateStatus {
	return original.PossibleTemplateStatusValues()
}
func PossibleTriggerOperatorValues() []TriggerOperator {
	return original.PossibleTriggerOperatorValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
