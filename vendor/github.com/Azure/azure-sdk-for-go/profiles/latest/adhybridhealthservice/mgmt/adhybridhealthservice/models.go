// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package adhybridhealthservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/adhybridhealthservice/mgmt/2014-01-01/adhybridhealthservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlgorithmStepType = original.AlgorithmStepType

const (
	ConnectorFilter         AlgorithmStepType = original.ConnectorFilter
	Deprovisioning          AlgorithmStepType = original.Deprovisioning
	ExportFlow              AlgorithmStepType = original.ExportFlow
	ImportFlow              AlgorithmStepType = original.ImportFlow
	Join                    AlgorithmStepType = original.Join
	MvDeletion              AlgorithmStepType = original.MvDeletion
	MvObjectTypeChange      AlgorithmStepType = original.MvObjectTypeChange
	Projection              AlgorithmStepType = original.Projection
	Provisioning            AlgorithmStepType = original.Provisioning
	Recall                  AlgorithmStepType = original.Recall
	Staging                 AlgorithmStepType = original.Staging
	Undefined               AlgorithmStepType = original.Undefined
	ValidateConnectorFilter AlgorithmStepType = original.ValidateConnectorFilter
)

type AttributeDeltaOperationType = original.AttributeDeltaOperationType

const (
	AttributeDeltaOperationTypeAdd       AttributeDeltaOperationType = original.AttributeDeltaOperationTypeAdd
	AttributeDeltaOperationTypeDelete    AttributeDeltaOperationType = original.AttributeDeltaOperationTypeDelete
	AttributeDeltaOperationTypeReplace   AttributeDeltaOperationType = original.AttributeDeltaOperationTypeReplace
	AttributeDeltaOperationTypeUndefined AttributeDeltaOperationType = original.AttributeDeltaOperationTypeUndefined
	AttributeDeltaOperationTypeUpdate    AttributeDeltaOperationType = original.AttributeDeltaOperationTypeUpdate
)

type AttributeMappingType = original.AttributeMappingType

const (
	Constant AttributeMappingType = original.Constant
	Direct   AttributeMappingType = original.Direct
	DnPart   AttributeMappingType = original.DnPart
	Script   AttributeMappingType = original.Script
)

type DeltaOperationType = original.DeltaOperationType

const (
	DeltaOperationTypeAdd       DeltaOperationType = original.DeltaOperationTypeAdd
	DeltaOperationTypeDelete    DeltaOperationType = original.DeltaOperationTypeDelete
	DeltaOperationTypeDeleteAdd DeltaOperationType = original.DeltaOperationTypeDeleteAdd
	DeltaOperationTypeNone      DeltaOperationType = original.DeltaOperationTypeNone
	DeltaOperationTypeObsolete  DeltaOperationType = original.DeltaOperationTypeObsolete
	DeltaOperationTypeReplace   DeltaOperationType = original.DeltaOperationTypeReplace
	DeltaOperationTypeUndefined DeltaOperationType = original.DeltaOperationTypeUndefined
	DeltaOperationTypeUpdate    DeltaOperationType = original.DeltaOperationTypeUpdate
)

type HealthStatus = original.HealthStatus

const (
	Error        HealthStatus = original.Error
	Healthy      HealthStatus = original.Healthy
	Missing      HealthStatus = original.Missing
	NotMonitored HealthStatus = original.NotMonitored
	Warning      HealthStatus = original.Warning
)

type Level = original.Level

const (
	LevelError      Level = original.LevelError
	LevelPreWarning Level = original.LevelPreWarning
	LevelWarning    Level = original.LevelWarning
)

type MonitoringLevel = original.MonitoringLevel

const (
	Full    MonitoringLevel = original.Full
	Off     MonitoringLevel = original.Off
	Partial MonitoringLevel = original.Partial
)

type PasswordOperationTypes = original.PasswordOperationTypes

const (
	PasswordOperationTypesChange    PasswordOperationTypes = original.PasswordOperationTypesChange
	PasswordOperationTypesSet       PasswordOperationTypes = original.PasswordOperationTypesSet
	PasswordOperationTypesUndefined PasswordOperationTypes = original.PasswordOperationTypesUndefined
)

type State = original.State

const (
	Active                   State = original.Active
	ResolvedByPositiveResult State = original.ResolvedByPositiveResult
	ResolvedByStateChange    State = original.ResolvedByStateChange
	ResolvedByTimer          State = original.ResolvedByTimer
	ResolvedManually         State = original.ResolvedManually
)

type ValueDeltaOperationType = original.ValueDeltaOperationType

const (
	ValueDeltaOperationTypeAdd       ValueDeltaOperationType = original.ValueDeltaOperationTypeAdd
	ValueDeltaOperationTypeDelete    ValueDeltaOperationType = original.ValueDeltaOperationTypeDelete
	ValueDeltaOperationTypeUndefined ValueDeltaOperationType = original.ValueDeltaOperationTypeUndefined
	ValueDeltaOperationTypeUpdate    ValueDeltaOperationType = original.ValueDeltaOperationTypeUpdate
)

type ValueType = original.ValueType

const (
	ValueTypeBinary    ValueType = original.ValueTypeBinary
	ValueTypeBoolean   ValueType = original.ValueTypeBoolean
	ValueTypeDn        ValueType = original.ValueTypeDn
	ValueTypeInteger   ValueType = original.ValueTypeInteger
	ValueTypeString    ValueType = original.ValueTypeString
	ValueTypeUndefined ValueType = original.ValueTypeUndefined
)

type AdDomainServiceMembersClient = original.AdDomainServiceMembersClient
type AdditionalInformation = original.AdditionalInformation
type AddsConfiguration = original.AddsConfiguration
type AddsConfigurationIterator = original.AddsConfigurationIterator
type AddsConfigurationPage = original.AddsConfigurationPage
type AddsServiceClient = original.AddsServiceClient
type AddsServiceMember = original.AddsServiceMember
type AddsServiceMembers = original.AddsServiceMembers
type AddsServiceMembersClient = original.AddsServiceMembersClient
type AddsServiceMembersIterator = original.AddsServiceMembersIterator
type AddsServiceMembersPage = original.AddsServiceMembersPage
type AddsServicesClient = original.AddsServicesClient
type AddsServicesReplicationStatusClient = original.AddsServicesReplicationStatusClient
type AddsServicesServiceMembersClient = original.AddsServicesServiceMembersClient
type AddsServicesUserPreferenceClient = original.AddsServicesUserPreferenceClient
type Agent = original.Agent
type Alert = original.Alert
type AlertFeedback = original.AlertFeedback
type AlertFeedbacks = original.AlertFeedbacks
type Alerts = original.Alerts
type AlertsClient = original.AlertsClient
type AlertsIterator = original.AlertsIterator
type AlertsPage = original.AlertsPage
type AssociatedObject = original.AssociatedObject
type AttributeDelta = original.AttributeDelta
type AttributeMapping = original.AttributeMapping
type AttributeMppingSource = original.AttributeMppingSource
type BaseClient = original.BaseClient
type ChangeNotReimported = original.ChangeNotReimported
type ChangeNotReimportedDelta = original.ChangeNotReimportedDelta
type ChangeNotReimportedEntry = original.ChangeNotReimportedEntry
type ConfigurationClient = original.ConfigurationClient
type Connector = original.Connector
type ConnectorConnectionError = original.ConnectorConnectionError
type ConnectorConnectionErrors = original.ConnectorConnectionErrors
type ConnectorMetadata = original.ConnectorMetadata
type ConnectorMetadataDetails = original.ConnectorMetadataDetails
type ConnectorObjectError = original.ConnectorObjectError
type ConnectorObjectErrors = original.ConnectorObjectErrors
type Connectors = original.Connectors
type Credential = original.Credential
type Credentials = original.Credentials
type DataFreshnessDetails = original.DataFreshnessDetails
type Dimension = original.Dimension
type Dimensions = original.Dimensions
type DimensionsClient = original.DimensionsClient
type DimensionsIterator = original.DimensionsIterator
type DimensionsPage = original.DimensionsPage
type Display = original.Display
type ErrorCount = original.ErrorCount
type ErrorCounts = original.ErrorCounts
type ErrorDetail = original.ErrorDetail
type ErrorReportUsersEntries = original.ErrorReportUsersEntries
type ErrorReportUsersEntry = original.ErrorReportUsersEntry
type ExportError = original.ExportError
type ExportErrors = original.ExportErrors
type ExportStatus = original.ExportStatus
type ExportStatuses = original.ExportStatuses
type ExportStatusesIterator = original.ExportStatusesIterator
type ExportStatusesPage = original.ExportStatusesPage
type ExtensionErrorInfo = original.ExtensionErrorInfo
type ForestSummary = original.ForestSummary
type GlobalConfiguration = original.GlobalConfiguration
type GlobalConfigurations = original.GlobalConfigurations
type HelpLink = original.HelpLink
type Hotfix = original.Hotfix
type Hotfixes = original.Hotfixes
type ImportError = original.ImportError
type ImportErrors = original.ImportErrors
type InboundReplicationNeighbor = original.InboundReplicationNeighbor
type InboundReplicationNeighbors = original.InboundReplicationNeighbors
type Item = original.Item
type Items = original.Items
type MergedExportError = original.MergedExportError
type MergedExportErrors = original.MergedExportErrors
type MetricGroup = original.MetricGroup
type MetricMetadata = original.MetricMetadata
type MetricMetadataList = original.MetricMetadataList
type MetricMetadataListIterator = original.MetricMetadataListIterator
type MetricMetadataListPage = original.MetricMetadataListPage
type MetricSet = original.MetricSet
type MetricSets = original.MetricSets
type Metrics = original.Metrics
type MetricsIterator = original.MetricsIterator
type MetricsPage = original.MetricsPage
type ModuleConfiguration = original.ModuleConfiguration
type ModuleConfigurations = original.ModuleConfigurations
type ObjectWithSyncError = original.ObjectWithSyncError
type Operation = original.Operation
type OperationListResponse = original.OperationListResponse
type OperationListResponseIterator = original.OperationListResponseIterator
type OperationListResponsePage = original.OperationListResponsePage
type OperationsClient = original.OperationsClient
type Partition = original.Partition
type PartitionScope = original.PartitionScope
type PasswordHashSyncConfiguration = original.PasswordHashSyncConfiguration
type PasswordManagementSettings = original.PasswordManagementSettings
type ReplicationDetailsList = original.ReplicationDetailsList
type ReplicationStatus = original.ReplicationStatus
type ReplicationSummary = original.ReplicationSummary
type ReplicationSummaryList = original.ReplicationSummaryList
type ReportsClient = original.ReportsClient
type Result = original.Result
type RiskyIPBlobURI = original.RiskyIPBlobURI
type RiskyIPBlobUris = original.RiskyIPBlobUris
type RuleErrorInfo = original.RuleErrorInfo
type RunProfile = original.RunProfile
type RunProfiles = original.RunProfiles
type RunStep = original.RunStep
type ServiceClient = original.ServiceClient
type ServiceConfiguration = original.ServiceConfiguration
type ServiceMember = original.ServiceMember
type ServiceMembers = original.ServiceMembers
type ServiceMembersClient = original.ServiceMembersClient
type ServiceMembersIterator = original.ServiceMembersIterator
type ServiceMembersPage = original.ServiceMembersPage
type ServiceProperties = original.ServiceProperties
type Services = original.Services
type ServicesClient = original.ServicesClient
type ServicesIterator = original.ServicesIterator
type ServicesPage = original.ServicesPage
type TabularExportError = original.TabularExportError
type Tenant = original.Tenant
type TenantOnboardingDetails = original.TenantOnboardingDetails
type UserPreference = original.UserPreference
type ValueDelta = original.ValueDelta

func New() BaseClient {
	return original.New()
}
func NewAdDomainServiceMembersClient() AdDomainServiceMembersClient {
	return original.NewAdDomainServiceMembersClient()
}
func NewAdDomainServiceMembersClientWithBaseURI(baseURI string) AdDomainServiceMembersClient {
	return original.NewAdDomainServiceMembersClientWithBaseURI(baseURI)
}
func NewAddsConfigurationIterator(page AddsConfigurationPage) AddsConfigurationIterator {
	return original.NewAddsConfigurationIterator(page)
}
func NewAddsConfigurationPage(getNextPage func(context.Context, AddsConfiguration) (AddsConfiguration, error)) AddsConfigurationPage {
	return original.NewAddsConfigurationPage(getNextPage)
}
func NewAddsServiceClient() AddsServiceClient {
	return original.NewAddsServiceClient()
}
func NewAddsServiceClientWithBaseURI(baseURI string) AddsServiceClient {
	return original.NewAddsServiceClientWithBaseURI(baseURI)
}
func NewAddsServiceMembersClient() AddsServiceMembersClient {
	return original.NewAddsServiceMembersClient()
}
func NewAddsServiceMembersClientWithBaseURI(baseURI string) AddsServiceMembersClient {
	return original.NewAddsServiceMembersClientWithBaseURI(baseURI)
}
func NewAddsServiceMembersIterator(page AddsServiceMembersPage) AddsServiceMembersIterator {
	return original.NewAddsServiceMembersIterator(page)
}
func NewAddsServiceMembersPage(getNextPage func(context.Context, AddsServiceMembers) (AddsServiceMembers, error)) AddsServiceMembersPage {
	return original.NewAddsServiceMembersPage(getNextPage)
}
func NewAddsServicesClient() AddsServicesClient {
	return original.NewAddsServicesClient()
}
func NewAddsServicesClientWithBaseURI(baseURI string) AddsServicesClient {
	return original.NewAddsServicesClientWithBaseURI(baseURI)
}
func NewAddsServicesReplicationStatusClient() AddsServicesReplicationStatusClient {
	return original.NewAddsServicesReplicationStatusClient()
}
func NewAddsServicesReplicationStatusClientWithBaseURI(baseURI string) AddsServicesReplicationStatusClient {
	return original.NewAddsServicesReplicationStatusClientWithBaseURI(baseURI)
}
func NewAddsServicesServiceMembersClient() AddsServicesServiceMembersClient {
	return original.NewAddsServicesServiceMembersClient()
}
func NewAddsServicesServiceMembersClientWithBaseURI(baseURI string) AddsServicesServiceMembersClient {
	return original.NewAddsServicesServiceMembersClientWithBaseURI(baseURI)
}
func NewAddsServicesUserPreferenceClient() AddsServicesUserPreferenceClient {
	return original.NewAddsServicesUserPreferenceClient()
}
func NewAddsServicesUserPreferenceClientWithBaseURI(baseURI string) AddsServicesUserPreferenceClient {
	return original.NewAddsServicesUserPreferenceClientWithBaseURI(baseURI)
}
func NewAlertsClient() AlertsClient {
	return original.NewAlertsClient()
}
func NewAlertsClientWithBaseURI(baseURI string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI)
}
func NewAlertsIterator(page AlertsPage) AlertsIterator {
	return original.NewAlertsIterator(page)
}
func NewAlertsPage(getNextPage func(context.Context, Alerts) (Alerts, error)) AlertsPage {
	return original.NewAlertsPage(getNextPage)
}
func NewConfigurationClient() ConfigurationClient {
	return original.NewConfigurationClient()
}
func NewConfigurationClientWithBaseURI(baseURI string) ConfigurationClient {
	return original.NewConfigurationClientWithBaseURI(baseURI)
}
func NewDimensionsClient() DimensionsClient {
	return original.NewDimensionsClient()
}
func NewDimensionsClientWithBaseURI(baseURI string) DimensionsClient {
	return original.NewDimensionsClientWithBaseURI(baseURI)
}
func NewDimensionsIterator(page DimensionsPage) DimensionsIterator {
	return original.NewDimensionsIterator(page)
}
func NewDimensionsPage(getNextPage func(context.Context, Dimensions) (Dimensions, error)) DimensionsPage {
	return original.NewDimensionsPage(getNextPage)
}
func NewExportStatusesIterator(page ExportStatusesPage) ExportStatusesIterator {
	return original.NewExportStatusesIterator(page)
}
func NewExportStatusesPage(getNextPage func(context.Context, ExportStatuses) (ExportStatuses, error)) ExportStatusesPage {
	return original.NewExportStatusesPage(getNextPage)
}
func NewMetricMetadataListIterator(page MetricMetadataListPage) MetricMetadataListIterator {
	return original.NewMetricMetadataListIterator(page)
}
func NewMetricMetadataListPage(getNextPage func(context.Context, MetricMetadataList) (MetricMetadataList, error)) MetricMetadataListPage {
	return original.NewMetricMetadataListPage(getNextPage)
}
func NewMetricsIterator(page MetricsPage) MetricsIterator {
	return original.NewMetricsIterator(page)
}
func NewMetricsPage(getNextPage func(context.Context, Metrics) (Metrics, error)) MetricsPage {
	return original.NewMetricsPage(getNextPage)
}
func NewOperationListResponseIterator(page OperationListResponsePage) OperationListResponseIterator {
	return original.NewOperationListResponseIterator(page)
}
func NewOperationListResponsePage(getNextPage func(context.Context, OperationListResponse) (OperationListResponse, error)) OperationListResponsePage {
	return original.NewOperationListResponsePage(getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewReportsClient() ReportsClient {
	return original.NewReportsClient()
}
func NewReportsClientWithBaseURI(baseURI string) ReportsClient {
	return original.NewReportsClientWithBaseURI(baseURI)
}
func NewServiceClient() ServiceClient {
	return original.NewServiceClient()
}
func NewServiceClientWithBaseURI(baseURI string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI)
}
func NewServiceMembersClient() ServiceMembersClient {
	return original.NewServiceMembersClient()
}
func NewServiceMembersClientWithBaseURI(baseURI string) ServiceMembersClient {
	return original.NewServiceMembersClientWithBaseURI(baseURI)
}
func NewServiceMembersIterator(page ServiceMembersPage) ServiceMembersIterator {
	return original.NewServiceMembersIterator(page)
}
func NewServiceMembersPage(getNextPage func(context.Context, ServiceMembers) (ServiceMembers, error)) ServiceMembersPage {
	return original.NewServiceMembersPage(getNextPage)
}
func NewServicesClient() ServicesClient {
	return original.NewServicesClient()
}
func NewServicesClientWithBaseURI(baseURI string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI)
}
func NewServicesIterator(page ServicesPage) ServicesIterator {
	return original.NewServicesIterator(page)
}
func NewServicesPage(getNextPage func(context.Context, Services) (Services, error)) ServicesPage {
	return original.NewServicesPage(getNextPage)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAlgorithmStepTypeValues() []AlgorithmStepType {
	return original.PossibleAlgorithmStepTypeValues()
}
func PossibleAttributeDeltaOperationTypeValues() []AttributeDeltaOperationType {
	return original.PossibleAttributeDeltaOperationTypeValues()
}
func PossibleAttributeMappingTypeValues() []AttributeMappingType {
	return original.PossibleAttributeMappingTypeValues()
}
func PossibleDeltaOperationTypeValues() []DeltaOperationType {
	return original.PossibleDeltaOperationTypeValues()
}
func PossibleHealthStatusValues() []HealthStatus {
	return original.PossibleHealthStatusValues()
}
func PossibleLevelValues() []Level {
	return original.PossibleLevelValues()
}
func PossibleMonitoringLevelValues() []MonitoringLevel {
	return original.PossibleMonitoringLevelValues()
}
func PossiblePasswordOperationTypesValues() []PasswordOperationTypes {
	return original.PossiblePasswordOperationTypesValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleValueDeltaOperationTypeValues() []ValueDeltaOperationType {
	return original.PossibleValueDeltaOperationTypeValues()
}
func PossibleValueTypeValues() []ValueType {
	return original.PossibleValueTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
