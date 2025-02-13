/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`. For more information on authentication, please refer to the [authentication token docs](https://docs.svix.com/api-keys).  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4.12
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndpointStats struct for EndpointStats
type EndpointStats struct {
	Fail int64 `json:"fail"`
	Pending int64 `json:"pending"`
	Sending int64 `json:"sending"`
	Success int64 `json:"success"`
}

// NewEndpointStats instantiates a new EndpointStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointStats(fail int64, pending int64, sending int64, success int64) *EndpointStats {
	this := EndpointStats{}
	this.Fail = fail
	this.Pending = pending
	this.Sending = sending
	this.Success = success
	return &this
}

// NewEndpointStatsWithDefaults instantiates a new EndpointStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointStatsWithDefaults() *EndpointStats {
	this := EndpointStats{}
	return &this
}

// GetFail returns the Fail field value
func (o *EndpointStats) GetFail() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fail
}

// GetFailOk returns a tuple with the Fail field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetFailOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fail, true
}

// SetFail sets field value
func (o *EndpointStats) SetFail(v int64) {
	o.Fail = v
}

// GetPending returns the Pending field value
func (o *EndpointStats) GetPending() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetPendingOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *EndpointStats) SetPending(v int64) {
	o.Pending = v
}

// GetSending returns the Sending field value
func (o *EndpointStats) GetSending() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sending
}

// GetSendingOk returns a tuple with the Sending field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetSendingOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sending, true
}

// SetSending sets field value
func (o *EndpointStats) SetSending(v int64) {
	o.Sending = v
}

// GetSuccess returns the Success field value
func (o *EndpointStats) GetSuccess() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *EndpointStats) GetSuccessOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *EndpointStats) SetSuccess(v int64) {
	o.Success = v
}

func (o EndpointStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fail"] = o.Fail
	}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if true {
		toSerialize["sending"] = o.Sending
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointStats struct {
	value *EndpointStats
	isSet bool
}

func (v NullableEndpointStats) Get() *EndpointStats {
	return v.value
}

func (v *NullableEndpointStats) Set(val *EndpointStats) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointStats) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointStats(val *EndpointStats) *NullableEndpointStats {
	return &NullableEndpointStats{value: val, isSet: true}
}

func (v NullableEndpointStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


