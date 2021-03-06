package scheduler

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DayOfWeek enumerates the values for day of week.
type DayOfWeek string

const (
	// Friday specifies the friday state for day of week.
	Friday DayOfWeek = "Friday"
	// Monday specifies the monday state for day of week.
	Monday DayOfWeek = "Monday"
	// Saturday specifies the saturday state for day of week.
	Saturday DayOfWeek = "Saturday"
	// Sunday specifies the sunday state for day of week.
	Sunday DayOfWeek = "Sunday"
	// Thursday specifies the thursday state for day of week.
	Thursday DayOfWeek = "Thursday"
	// Tuesday specifies the tuesday state for day of week.
	Tuesday DayOfWeek = "Tuesday"
	// Wednesday specifies the wednesday state for day of week.
	Wednesday DayOfWeek = "Wednesday"
)

// HTTPAuthenticationType enumerates the values for http authentication type.
type HTTPAuthenticationType string

const (
	// ActiveDirectoryOAuth specifies the active directory o auth state for
	// http authentication type.
	ActiveDirectoryOAuth HTTPAuthenticationType = "ActiveDirectoryOAuth"
	// Basic specifies the basic state for http authentication type.
	Basic HTTPAuthenticationType = "Basic"
	// ClientCertificate specifies the client certificate state for http
	// authentication type.
	ClientCertificate HTTPAuthenticationType = "ClientCertificate"
	// NotSpecified specifies the not specified state for http authentication
	// type.
	NotSpecified HTTPAuthenticationType = "NotSpecified"
)

// JobActionType enumerates the values for job action type.
type JobActionType string

const (
	// HTTP specifies the http state for job action type.
	HTTP JobActionType = "Http"
	// HTTPS specifies the https state for job action type.
	HTTPS JobActionType = "Https"
	// ServiceBusQueue specifies the service bus queue state for job action
	// type.
	ServiceBusQueue JobActionType = "ServiceBusQueue"
	// ServiceBusTopic specifies the service bus topic state for job action
	// type.
	ServiceBusTopic JobActionType = "ServiceBusTopic"
	// StorageQueue specifies the storage queue state for job action type.
	StorageQueue JobActionType = "StorageQueue"
)

// JobCollectionState enumerates the values for job collection state.
type JobCollectionState string

const (
	// Deleted specifies the deleted state for job collection state.
	Deleted JobCollectionState = "Deleted"
	// Disabled specifies the disabled state for job collection state.
	Disabled JobCollectionState = "Disabled"
	// Enabled specifies the enabled state for job collection state.
	Enabled JobCollectionState = "Enabled"
	// Suspended specifies the suspended state for job collection state.
	Suspended JobCollectionState = "Suspended"
)

// JobExecutionStatus enumerates the values for job execution status.
type JobExecutionStatus string

const (
	// CallbackNotFound specifies the callback not found state for job
	// execution status.
	CallbackNotFound JobExecutionStatus = "CallbackNotFound"
	// Cancelled specifies the cancelled state for job execution status.
	Cancelled JobExecutionStatus = "Cancelled"
	// Completed specifies the completed state for job execution status.
	Completed JobExecutionStatus = "Completed"
	// Failed specifies the failed state for job execution status.
	Failed JobExecutionStatus = "Failed"
	// Postponed specifies the postponed state for job execution status.
	Postponed JobExecutionStatus = "Postponed"
)

// JobHistoryActionName enumerates the values for job history action name.
type JobHistoryActionName string

const (
	// ErrorAction specifies the error action state for job history action
	// name.
	ErrorAction JobHistoryActionName = "ErrorAction"
	// MainAction specifies the main action state for job history action name.
	MainAction JobHistoryActionName = "MainAction"
)

// JobScheduleDay enumerates the values for job schedule day.
type JobScheduleDay string

const (
	// JobScheduleDayFriday specifies the job schedule day friday state for
	// job schedule day.
	JobScheduleDayFriday JobScheduleDay = "Friday"
	// JobScheduleDayMonday specifies the job schedule day monday state for
	// job schedule day.
	JobScheduleDayMonday JobScheduleDay = "Monday"
	// JobScheduleDaySaturday specifies the job schedule day saturday state
	// for job schedule day.
	JobScheduleDaySaturday JobScheduleDay = "Saturday"
	// JobScheduleDaySunday specifies the job schedule day sunday state for
	// job schedule day.
	JobScheduleDaySunday JobScheduleDay = "Sunday"
	// JobScheduleDayThursday specifies the job schedule day thursday state
	// for job schedule day.
	JobScheduleDayThursday JobScheduleDay = "Thursday"
	// JobScheduleDayTuesday specifies the job schedule day tuesday state for
	// job schedule day.
	JobScheduleDayTuesday JobScheduleDay = "Tuesday"
	// JobScheduleDayWednesday specifies the job schedule day wednesday state
	// for job schedule day.
	JobScheduleDayWednesday JobScheduleDay = "Wednesday"
)

// JobState enumerates the values for job state.
type JobState string

const (
	// JobStateCompleted specifies the job state completed state for job state.
	JobStateCompleted JobState = "Completed"
	// JobStateDisabled specifies the job state disabled state for job state.
	JobStateDisabled JobState = "Disabled"
	// JobStateEnabled specifies the job state enabled state for job state.
	JobStateEnabled JobState = "Enabled"
	// JobStateFaulted specifies the job state faulted state for job state.
	JobStateFaulted JobState = "Faulted"
)

// RecurrenceFrequency enumerates the values for recurrence frequency.
type RecurrenceFrequency string

const (
	// Day specifies the day state for recurrence frequency.
	Day RecurrenceFrequency = "Day"
	// Hour specifies the hour state for recurrence frequency.
	Hour RecurrenceFrequency = "Hour"
	// Minute specifies the minute state for recurrence frequency.
	Minute RecurrenceFrequency = "Minute"
	// Month specifies the month state for recurrence frequency.
	Month RecurrenceFrequency = "Month"
	// Week specifies the week state for recurrence frequency.
	Week RecurrenceFrequency = "Week"
)

// RetryType enumerates the values for retry type.
type RetryType string

const (
	// Fixed specifies the fixed state for retry type.
	Fixed RetryType = "Fixed"
	// None specifies the none state for retry type.
	None RetryType = "None"
)

// ServiceBusAuthenticationType enumerates the values for service bus
// authentication type.
type ServiceBusAuthenticationType string

const (
	// ServiceBusAuthenticationTypeNotSpecified specifies the service bus
	// authentication type not specified state for service bus authentication
	// type.
	ServiceBusAuthenticationTypeNotSpecified ServiceBusAuthenticationType = "NotSpecified"
	// ServiceBusAuthenticationTypeSharedAccessKey specifies the service bus
	// authentication type shared access key state for service bus
	// authentication type.
	ServiceBusAuthenticationTypeSharedAccessKey ServiceBusAuthenticationType = "SharedAccessKey"
)

// ServiceBusTransportType enumerates the values for service bus transport
// type.
type ServiceBusTransportType string

const (
	// ServiceBusTransportTypeAMQP specifies the service bus transport type
	// amqp state for service bus transport type.
	ServiceBusTransportTypeAMQP ServiceBusTransportType = "AMQP"
	// ServiceBusTransportTypeNetMessaging specifies the service bus transport
	// type net messaging state for service bus transport type.
	ServiceBusTransportTypeNetMessaging ServiceBusTransportType = "NetMessaging"
	// ServiceBusTransportTypeNotSpecified specifies the service bus transport
	// type not specified state for service bus transport type.
	ServiceBusTransportTypeNotSpecified ServiceBusTransportType = "NotSpecified"
)

// SkuDefinition enumerates the values for sku definition.
type SkuDefinition string

const (
	// Free specifies the free state for sku definition.
	Free SkuDefinition = "Free"
	// Premium specifies the premium state for sku definition.
	Premium SkuDefinition = "Premium"
	// Standard specifies the standard state for sku definition.
	Standard SkuDefinition = "Standard"
)

// BasicAuthentication is
type BasicAuthentication struct {
	Type     HTTPAuthenticationType `json:"type,omitempty"`
	Username *string                `json:"username,omitempty"`
	Password *string                `json:"password,omitempty"`
}

// ClientCertAuthentication is
type ClientCertAuthentication struct {
	Type                      HTTPAuthenticationType `json:"type,omitempty"`
	Password                  *string                `json:"password,omitempty"`
	Pfx                       *string                `json:"pfx,omitempty"`
	CertificateThumbprint     *string                `json:"certificateThumbprint,omitempty"`
	CertificateExpirationDate *date.Time             `json:"certificateExpirationDate,omitempty"`
	CertificateSubjectName    *string                `json:"certificateSubjectName,omitempty"`
}

// HTTPAuthentication is
type HTTPAuthentication struct {
	Type HTTPAuthenticationType `json:"type,omitempty"`
}

// HTTPRequest is
type HTTPRequest struct {
	HTTPAuthentication *HTTPAuthentication `json:"httpAuthentication,omitempty"`
	URI                *string             `json:"uri,omitempty"`
	Method             *string             `json:"method,omitempty"`
	Body               *string             `json:"body,omitempty"`
	Headers            *map[string]*string `json:"headers,omitempty"`
}

// JobAction is
type JobAction struct {
	Type                   JobActionType           `json:"type,omitempty"`
	Request                *HTTPRequest            `json:"request,omitempty"`
	QueueMessage           *StorageQueueMessage    `json:"queueMessage,omitempty"`
	ServiceBusQueueMessage *ServiceBusQueueMessage `json:"serviceBusQueueMessage,omitempty"`
	ServiceBusTopicMessage *ServiceBusTopicMessage `json:"serviceBusTopicMessage,omitempty"`
	RetryPolicy            *RetryPolicy            `json:"retryPolicy,omitempty"`
	ErrorAction            *JobErrorAction         `json:"errorAction,omitempty"`
}

// JobCollectionDefinition is
type JobCollectionDefinition struct {
	autorest.Response `json:"-"`
	ID                *string                  `json:"id,omitempty"`
	Type              *string                  `json:"type,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Location          *string                  `json:"location,omitempty"`
	Tags              *map[string]*string      `json:"tags,omitempty"`
	Properties        *JobCollectionProperties `json:"properties,omitempty"`
}

// JobCollectionListResult is
type JobCollectionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]JobCollectionDefinition `json:"value,omitempty"`
	NextLink          *string                    `json:"nextLink,omitempty"`
}

// JobCollectionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client JobCollectionListResult) JobCollectionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// JobCollectionProperties is
type JobCollectionProperties struct {
	Sku   *Sku                `json:"sku,omitempty"`
	State JobCollectionState  `json:"state,omitempty"`
	Quota *JobCollectionQuota `json:"quota,omitempty"`
}

// JobCollectionQuota is
type JobCollectionQuota struct {
	MaxJobCount      *int              `json:"maxJobCount,omitempty"`
	MaxJobOccurrence *int              `json:"maxJobOccurrence,omitempty"`
	MaxRecurrence    *JobMaxRecurrence `json:"maxRecurrence,omitempty"`
}

// JobDefinition is
type JobDefinition struct {
	autorest.Response `json:"-"`
	ID                *string        `json:"id,omitempty"`
	Type              *string        `json:"type,omitempty"`
	Name              *string        `json:"name,omitempty"`
	Properties        *JobProperties `json:"properties,omitempty"`
}

// JobErrorAction is
type JobErrorAction struct {
	Type                   JobActionType           `json:"type,omitempty"`
	Request                *HTTPRequest            `json:"request,omitempty"`
	QueueMessage           *StorageQueueMessage    `json:"queueMessage,omitempty"`
	ServiceBusQueueMessage *ServiceBusQueueMessage `json:"serviceBusQueueMessage,omitempty"`
	ServiceBusTopicMessage *ServiceBusTopicMessage `json:"serviceBusTopicMessage,omitempty"`
	RetryPolicy            *RetryPolicy            `json:"retryPolicy,omitempty"`
}

// JobHistoryDefinition is
type JobHistoryDefinition struct {
	ID         *string                         `json:"id,omitempty"`
	Type       *string                         `json:"type,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties *JobHistoryDefinitionProperties `json:"properties,omitempty"`
}

// JobHistoryDefinitionProperties is
type JobHistoryDefinitionProperties struct {
	StartTime             *date.Time           `json:"startTime,omitempty"`
	EndTime               *date.Time           `json:"endTime,omitempty"`
	ExpectedExecutionTime *date.Time           `json:"expectedExecutionTime,omitempty"`
	ActionName            JobHistoryActionName `json:"actionName,omitempty"`
	Status                JobExecutionStatus   `json:"status,omitempty"`
	Message               *string              `json:"message,omitempty"`
	RetryCount            *int                 `json:"retryCount,omitempty"`
	RepeatCount           *int                 `json:"repeatCount,omitempty"`
}

// JobHistoryListResult is
type JobHistoryListResult struct {
	autorest.Response `json:"-"`
	Value             *[]JobHistoryDefinition `json:"value,omitempty"`
	NextLink          *string                 `json:"nextLink,omitempty"`
}

// JobHistoryListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client JobHistoryListResult) JobHistoryListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// JobListResult is
type JobListResult struct {
	autorest.Response `json:"-"`
	Value             *[]JobDefinition `json:"value,omitempty"`
	NextLink          *string          `json:"nextLink,omitempty"`
}

// JobListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client JobListResult) JobListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// JobMaxRecurrence is
type JobMaxRecurrence struct {
	Frequency RecurrenceFrequency `json:"frequency,omitempty"`
	Interval  *int                `json:"interval,omitempty"`
}

// JobProperties is
type JobProperties struct {
	StartTime  *date.Time     `json:"startTime,omitempty"`
	Action     *JobAction     `json:"action,omitempty"`
	Recurrence *JobRecurrence `json:"recurrence,omitempty"`
	State      JobState       `json:"state,omitempty"`
	Status     *JobStatus     `json:"status,omitempty"`
}

// JobRecurrence is
type JobRecurrence struct {
	Frequency RecurrenceFrequency    `json:"frequency,omitempty"`
	Interval  *int                   `json:"interval,omitempty"`
	Count     *int                   `json:"count,omitempty"`
	EndTime   *date.Time             `json:"endTime,omitempty"`
	Schedule  *JobRecurrenceSchedule `json:"schedule,omitempty"`
}

// JobRecurrenceSchedule is
type JobRecurrenceSchedule struct {
	WeekDays           *[]DayOfWeek                              `json:"weekDays,omitempty"`
	Hours              *[]int                                    `json:"hours,omitempty"`
	Minutes            *[]int                                    `json:"minutes,omitempty"`
	MonthDays          *[]int                                    `json:"monthDays,omitempty"`
	MonthlyOccurrences *[]JobRecurrenceScheduleMonthlyOccurrence `json:"monthlyOccurrences,omitempty"`
}

// JobRecurrenceScheduleMonthlyOccurrence is
type JobRecurrenceScheduleMonthlyOccurrence struct {
	Day        JobScheduleDay `json:"day,omitempty"`
	Occurrence *int           `json:"Occurrence,omitempty"`
}

// JobStatus is
type JobStatus struct {
	ExecutionCount    *int       `json:"executionCount,omitempty"`
	FailureCount      *int       `json:"failureCount,omitempty"`
	FaultedCount      *int       `json:"faultedCount,omitempty"`
	LastExecutionTime *date.Time `json:"lastExecutionTime,omitempty"`
	NextExecutionTime *date.Time `json:"nextExecutionTime,omitempty"`
}

// OAuthAuthentication is
type OAuthAuthentication struct {
	Type     HTTPAuthenticationType `json:"type,omitempty"`
	Secret   *string                `json:"secret,omitempty"`
	Tenant   *string                `json:"tenant,omitempty"`
	Audience *string                `json:"audience,omitempty"`
	ClientID *string                `json:"clientId,omitempty"`
}

// RetryPolicy is
type RetryPolicy struct {
	RetryType     RetryType `json:"retryType,omitempty"`
	RetryInterval *string   `json:"retryInterval,omitempty"`
	RetryCount    *int      `json:"retryCount,omitempty"`
}

// ServiceBusAuthentication is
type ServiceBusAuthentication struct {
	SasKey     *string                      `json:"sasKey,omitempty"`
	SasKeyName *string                      `json:"sasKeyName,omitempty"`
	Type       ServiceBusAuthenticationType `json:"type,omitempty"`
}

// ServiceBusBrokeredMessageProperties is
type ServiceBusBrokeredMessageProperties struct {
	ContentType             *string    `json:"contentType,omitempty"`
	CorrelationID           *string    `json:"correlationId,omitempty"`
	ForcePersistence        *bool      `json:"forcePersistence,omitempty"`
	Label                   *string    `json:"label,omitempty"`
	MessageID               *string    `json:"messageId,omitempty"`
	PartitionKey            *string    `json:"partitionKey,omitempty"`
	ReplyTo                 *string    `json:"replyTo,omitempty"`
	ReplyToSessionID        *string    `json:"replyToSessionId,omitempty"`
	ScheduledEnqueueTimeUtc *date.Time `json:"scheduledEnqueueTimeUtc,omitempty"`
	SessionID               *string    `json:"sessionId,omitempty"`
	TimeToLive              *date.Time `json:"timeToLive,omitempty"`
	To                      *string    `json:"to,omitempty"`
	ViaPartitionKey         *string    `json:"viaPartitionKey,omitempty"`
}

// ServiceBusMessage is
type ServiceBusMessage struct {
	Authentication            *ServiceBusAuthentication            `json:"authentication,omitempty"`
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `json:"brokeredMessageProperties,omitempty"`
	CustomMessageProperties   *map[string]*string                  `json:"customMessageProperties,omitempty"`
	Message                   *string                              `json:"message,omitempty"`
	Namespace                 *string                              `json:"namespace,omitempty"`
	TransportType             ServiceBusTransportType              `json:"transportType,omitempty"`
}

// ServiceBusQueueMessage is
type ServiceBusQueueMessage struct {
	Authentication            *ServiceBusAuthentication            `json:"authentication,omitempty"`
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `json:"brokeredMessageProperties,omitempty"`
	CustomMessageProperties   *map[string]*string                  `json:"customMessageProperties,omitempty"`
	Message                   *string                              `json:"message,omitempty"`
	Namespace                 *string                              `json:"namespace,omitempty"`
	TransportType             ServiceBusTransportType              `json:"transportType,omitempty"`
	QueueName                 *string                              `json:"queueName,omitempty"`
}

// ServiceBusTopicMessage is
type ServiceBusTopicMessage struct {
	Authentication            *ServiceBusAuthentication            `json:"authentication,omitempty"`
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `json:"brokeredMessageProperties,omitempty"`
	CustomMessageProperties   *map[string]*string                  `json:"customMessageProperties,omitempty"`
	Message                   *string                              `json:"message,omitempty"`
	Namespace                 *string                              `json:"namespace,omitempty"`
	TransportType             ServiceBusTransportType              `json:"transportType,omitempty"`
	TopicPath                 *string                              `json:"topicPath,omitempty"`
}

// Sku is
type Sku struct {
	Name SkuDefinition `json:"name,omitempty"`
}

// StorageQueueMessage is
type StorageQueueMessage struct {
	StorageAccount *string `json:"storageAccount,omitempty"`
	QueueName      *string `json:"queueName,omitempty"`
	SasToken       *string `json:"sasToken,omitempty"`
	Message        *string `json:"message,omitempty"`
}
