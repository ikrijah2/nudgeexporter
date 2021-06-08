package dns

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns"

// AaaaRecord an AAAA record.
type AaaaRecord struct {
	// Ipv6Address - Gets or sets the IPv6 address of this AAAA record in string notation.
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}

// ARecord an A record.
type ARecord struct {
	// Ipv4Address - Gets or sets the IPv4 address of this A record in string notation.
	Ipv4Address *string `json:"ipv4Address,omitempty"`
}

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// CnameRecord a CNAME record.
type CnameRecord struct {
	// Cname - Gets or sets the canonical name for this record without a terminating dot.
	Cname *string `json:"cname,omitempty"`
}

// MxRecord an MX record.
type MxRecord struct {
	// Preference - Gets or sets the preference metric for this record.
	Preference *int32 `json:"preference,omitempty"`
	// Exchange - Gets or sets the domain name of the mail host, without a terminating dot.
	Exchange *string `json:"exchange,omitempty"`
}

// NsRecord an NS record.
type NsRecord struct {
	// Nsdname - Gets or sets the name server name for this record, without a terminating dot.
	Nsdname *string `json:"nsdname,omitempty"`
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// PtrRecord a PTR record.
type PtrRecord struct {
	// Ptrdname - Gets or sets the PTR target domain name for this record without a terminating dot.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// RecordSet describes a DNS RecordSet (a set of DNS records with the same name and type).
type RecordSet struct {
	autorest.Response `json:"-"`
	// Etag - Gets or sets the ETag of the RecordSet.
	Etag *string `json:"etag,omitempty"`
	// Properties - Gets or sets the properties of the RecordSet.
	Properties *RecordSetProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for RecordSet.
func (rs RecordSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rs.Etag != nil {
		objectMap["etag"] = rs.Etag
	}
	if rs.Properties != nil {
		objectMap["properties"] = rs.Properties
	}
	if rs.Tags != nil {
		objectMap["tags"] = rs.Tags
	}
	if rs.Location != nil {
		objectMap["location"] = rs.Location
	}
	return json.Marshal(objectMap)
}

// RecordSetListResult the response to a RecordSet List operation.
type RecordSetListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets information about the RecordSets in the response.
	Value *[]RecordSet `json:"value,omitempty"`
	// NextLink - Gets or sets the continuation token for the next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// RecordSetListResultIterator provides access to a complete listing of RecordSet values.
type RecordSetListResultIterator struct {
	i    int
	page RecordSetListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RecordSetListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *RecordSetListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RecordSetListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RecordSetListResultIterator) Response() RecordSetListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RecordSetListResultIterator) Value() RecordSet {
	if !iter.page.NotDone() {
		return RecordSet{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RecordSetListResultIterator type.
func NewRecordSetListResultIterator(page RecordSetListResultPage) RecordSetListResultIterator {
	return RecordSetListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rslr RecordSetListResult) IsEmpty() bool {
	return rslr.Value == nil || len(*rslr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rslr RecordSetListResult) hasNextLink() bool {
	return rslr.NextLink != nil && len(*rslr.NextLink) != 0
}

// recordSetListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rslr RecordSetListResult) recordSetListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !rslr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rslr.NextLink)))
}

// RecordSetListResultPage contains a page of RecordSet values.
type RecordSetListResultPage struct {
	fn   func(context.Context, RecordSetListResult) (RecordSetListResult, error)
	rslr RecordSetListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RecordSetListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rslr)
		if err != nil {
			return err
		}
		page.rslr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RecordSetListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RecordSetListResultPage) NotDone() bool {
	return !page.rslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RecordSetListResultPage) Response() RecordSetListResult {
	return page.rslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RecordSetListResultPage) Values() []RecordSet {
	if page.rslr.IsEmpty() {
		return nil
	}
	return *page.rslr.Value
}

// Creates a new instance of the RecordSetListResultPage type.
func NewRecordSetListResultPage(cur RecordSetListResult, getNextPage func(context.Context, RecordSetListResult) (RecordSetListResult, error)) RecordSetListResultPage {
	return RecordSetListResultPage{
		fn:   getNextPage,
		rslr: cur,
	}
}

// RecordSetProperties represents the properties of the records in the RecordSet.
type RecordSetProperties struct {
	// TTL - Gets or sets the TTL of the records in the RecordSet.
	TTL *int64 `json:"TTL,omitempty"`
	// ARecords - Gets or sets the list of A records in the RecordSet.
	ARecords *[]ARecord `json:"ARecords,omitempty"`
	// AAAARecords - Gets or sets the list of AAAA records in the RecordSet.
	AAAARecords *[]AaaaRecord `json:"AAAARecords,omitempty"`
	// MXRecords - Gets or sets the list of MX records in the RecordSet.
	MXRecords *[]MxRecord `json:"MXRecords,omitempty"`
	// NSRecords - Gets or sets the list of NS records in the RecordSet.
	NSRecords *[]NsRecord `json:"NSRecords,omitempty"`
	// PTRRecords - Gets or sets the list of PTR records in the RecordSet.
	PTRRecords *[]PtrRecord `json:"PTRRecords,omitempty"`
	// SRVRecords - Gets or sets the list of SRV records in the RecordSet.
	SRVRecords *[]SrvRecord `json:"SRVRecords,omitempty"`
	// TXTRecords - Gets or sets the list of TXT records in the RecordSet.
	TXTRecords *[]TxtRecord `json:"TXTRecords,omitempty"`
	// CNAMERecord - Gets or sets the CNAME record in the RecordSet.
	CNAMERecord *CnameRecord `json:"CNAMERecord,omitempty"`
	// SOARecord - Gets or sets the SOA record in the RecordSet.
	SOARecord *SoaRecord `json:"SOARecord,omitempty"`
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// SoaRecord an SOA record.
type SoaRecord struct {
	// Host - Gets or sets the domain name of the authoritative name server, without a terminating dot.
	Host *string `json:"host,omitempty"`
	// Email - Gets or sets the email for this record.
	Email *string `json:"email,omitempty"`
	// SerialNumber - Gets or sets the serial number for this record.
	SerialNumber *int64 `json:"serialNumber,omitempty"`
	// RefreshTime - Gets or sets the refresh value for this record.
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// RetryTime - Gets or sets the retry time for this record.
	RetryTime *int64 `json:"retryTime,omitempty"`
	// ExpireTime - Gets or sets the expire time for this record.
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// MinimumTTL - Gets or sets the minimum TTL value for this record.
	MinimumTTL *int64 `json:"minimumTTL,omitempty"`
}

// SrvRecord an SRV record.
type SrvRecord struct {
	// Priority - Gets or sets the priority metric for this record.
	Priority *int32 `json:"priority,omitempty"`
	// Weight - Gets or sets the weight metric for this record.
	Weight *int32 `json:"weight,omitempty"`
	// Port - Gets or sets the port of the service for this record.
	Port *int32 `json:"port,omitempty"`
	// Target - Gets or sets the domain name of the target for this record, without a terminating dot.
	Target *string `json:"target,omitempty"`
}

// SubResource ...
type SubResource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}

// TxtRecord a TXT record.
type TxtRecord struct {
	// Value - Gets or sets the text value of this record.
	Value *[]string `json:"value,omitempty"`
}

// Zone describes a DNS zone.
type Zone struct {
	autorest.Response `json:"-"`
	// Etag - Gets or sets the ETag of the zone that is being updated, as received from a Get operation.
	Etag *string `json:"etag,omitempty"`
	// Properties - Gets or sets the properties of the zone.
	Properties *ZoneProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Zone.
func (z Zone) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if z.Etag != nil {
		objectMap["etag"] = z.Etag
	}
	if z.Properties != nil {
		objectMap["properties"] = z.Properties
	}
	if z.Tags != nil {
		objectMap["tags"] = z.Tags
	}
	if z.Location != nil {
		objectMap["location"] = z.Location
	}
	return json.Marshal(objectMap)
}

// ZoneListResult the response to a Zone List or ListAll operation.
type ZoneListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets information about the zones in the response.
	Value *[]Zone `json:"value,omitempty"`
	// NextLink - Gets or sets the continuation token for the next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// ZoneListResultIterator provides access to a complete listing of Zone values.
type ZoneListResultIterator struct {
	i    int
	page ZoneListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ZoneListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ZoneListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ZoneListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ZoneListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ZoneListResultIterator) Response() ZoneListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ZoneListResultIterator) Value() Zone {
	if !iter.page.NotDone() {
		return Zone{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ZoneListResultIterator type.
func NewZoneListResultIterator(page ZoneListResultPage) ZoneListResultIterator {
	return ZoneListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (zlr ZoneListResult) IsEmpty() bool {
	return zlr.Value == nil || len(*zlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (zlr ZoneListResult) hasNextLink() bool {
	return zlr.NextLink != nil && len(*zlr.NextLink) != 0
}

// zoneListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (zlr ZoneListResult) zoneListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !zlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(zlr.NextLink)))
}

// ZoneListResultPage contains a page of Zone values.
type ZoneListResultPage struct {
	fn  func(context.Context, ZoneListResult) (ZoneListResult, error)
	zlr ZoneListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ZoneListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ZoneListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.zlr)
		if err != nil {
			return err
		}
		page.zlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ZoneListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ZoneListResultPage) NotDone() bool {
	return !page.zlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ZoneListResultPage) Response() ZoneListResult {
	return page.zlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ZoneListResultPage) Values() []Zone {
	if page.zlr.IsEmpty() {
		return nil
	}
	return *page.zlr.Value
}

// Creates a new instance of the ZoneListResultPage type.
func NewZoneListResultPage(cur ZoneListResult, getNextPage func(context.Context, ZoneListResult) (ZoneListResult, error)) ZoneListResultPage {
	return ZoneListResultPage{
		fn:  getNextPage,
		zlr: cur,
	}
}

// ZoneProperties represents the properties of the zone.
type ZoneProperties struct {
	// MaxNumberOfRecordSets - Gets or sets the maximum number of record sets that can be created in this zone.
	MaxNumberOfRecordSets *int64 `json:"maxNumberOfRecordSets,omitempty"`
	// NumberOfRecordSets - Gets or sets the current number of record sets in this zone.
	NumberOfRecordSets *int64 `json:"numberOfRecordSets,omitempty"`
}
