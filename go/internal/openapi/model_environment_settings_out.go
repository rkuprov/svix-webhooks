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

// EnvironmentSettingsOut struct for EnvironmentSettingsOut
type EnvironmentSettingsOut struct {
	ColorPaletteDark *CustomColorPalette `json:"colorPaletteDark,omitempty"`
	ColorPaletteLight *CustomColorPalette `json:"colorPaletteLight,omitempty"`
	CustomColor NullableString `json:"customColor,omitempty"`
	CustomFontFamily NullableString `json:"customFontFamily,omitempty"`
	CustomLogoUrl NullableString `json:"customLogoUrl,omitempty"`
	CustomThemeOverride *CustomThemeOverride `json:"customThemeOverride,omitempty"`
	EnableChannels *bool `json:"enableChannels,omitempty"`
	EnableIntegrationManagement *bool `json:"enableIntegrationManagement,omitempty"`
	EnableTransformations *bool `json:"enableTransformations,omitempty"`
}

// NewEnvironmentSettingsOut instantiates a new EnvironmentSettingsOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSettingsOut() *EnvironmentSettingsOut {
	this := EnvironmentSettingsOut{}
	var enableChannels bool = false
	this.EnableChannels = &enableChannels
	var enableIntegrationManagement bool = false
	this.EnableIntegrationManagement = &enableIntegrationManagement
	var enableTransformations bool = false
	this.EnableTransformations = &enableTransformations
	return &this
}

// NewEnvironmentSettingsOutWithDefaults instantiates a new EnvironmentSettingsOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSettingsOutWithDefaults() *EnvironmentSettingsOut {
	this := EnvironmentSettingsOut{}
	var enableChannels bool = false
	this.EnableChannels = &enableChannels
	var enableIntegrationManagement bool = false
	this.EnableIntegrationManagement = &enableIntegrationManagement
	var enableTransformations bool = false
	this.EnableTransformations = &enableTransformations
	return &this
}

// GetColorPaletteDark returns the ColorPaletteDark field value if set, zero value otherwise.
func (o *EnvironmentSettingsOut) GetColorPaletteDark() CustomColorPalette {
	if o == nil || o.ColorPaletteDark == nil {
		var ret CustomColorPalette
		return ret
	}
	return *o.ColorPaletteDark
}

// GetColorPaletteDarkOk returns a tuple with the ColorPaletteDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSettingsOut) GetColorPaletteDarkOk() (*CustomColorPalette, bool) {
	if o == nil || o.ColorPaletteDark == nil {
		return nil, false
	}
	return o.ColorPaletteDark, true
}

// HasColorPaletteDark returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasColorPaletteDark() bool {
	if o != nil && o.ColorPaletteDark != nil {
		return true
	}

	return false
}

// SetColorPaletteDark gets a reference to the given CustomColorPalette and assigns it to the ColorPaletteDark field.
func (o *EnvironmentSettingsOut) SetColorPaletteDark(v CustomColorPalette) {
	o.ColorPaletteDark = &v
}

// GetColorPaletteLight returns the ColorPaletteLight field value if set, zero value otherwise.
func (o *EnvironmentSettingsOut) GetColorPaletteLight() CustomColorPalette {
	if o == nil || o.ColorPaletteLight == nil {
		var ret CustomColorPalette
		return ret
	}
	return *o.ColorPaletteLight
}

// GetColorPaletteLightOk returns a tuple with the ColorPaletteLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSettingsOut) GetColorPaletteLightOk() (*CustomColorPalette, bool) {
	if o == nil || o.ColorPaletteLight == nil {
		return nil, false
	}
	return o.ColorPaletteLight, true
}

// HasColorPaletteLight returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasColorPaletteLight() bool {
	if o != nil && o.ColorPaletteLight != nil {
		return true
	}

	return false
}

// SetColorPaletteLight gets a reference to the given CustomColorPalette and assigns it to the ColorPaletteLight field.
func (o *EnvironmentSettingsOut) SetColorPaletteLight(v CustomColorPalette) {
	o.ColorPaletteLight = &v
}

// GetCustomColor returns the CustomColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentSettingsOut) GetCustomColor() string {
	if o == nil || o.CustomColor.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomColor.Get()
}

// GetCustomColorOk returns a tuple with the CustomColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentSettingsOut) GetCustomColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomColor.Get(), o.CustomColor.IsSet()
}

// HasCustomColor returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasCustomColor() bool {
	if o != nil && o.CustomColor.IsSet() {
		return true
	}

	return false
}

// SetCustomColor gets a reference to the given NullableString and assigns it to the CustomColor field.
func (o *EnvironmentSettingsOut) SetCustomColor(v string) {
	o.CustomColor.Set(&v)
}
// SetCustomColorNil sets the value for CustomColor to be an explicit nil
func (o *EnvironmentSettingsOut) SetCustomColorNil() {
	o.CustomColor.Set(nil)
}

// UnsetCustomColor ensures that no value is present for CustomColor, not even an explicit nil
func (o *EnvironmentSettingsOut) UnsetCustomColor() {
	o.CustomColor.Unset()
}

// GetCustomFontFamily returns the CustomFontFamily field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentSettingsOut) GetCustomFontFamily() string {
	if o == nil || o.CustomFontFamily.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomFontFamily.Get()
}

// GetCustomFontFamilyOk returns a tuple with the CustomFontFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentSettingsOut) GetCustomFontFamilyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomFontFamily.Get(), o.CustomFontFamily.IsSet()
}

// HasCustomFontFamily returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasCustomFontFamily() bool {
	if o != nil && o.CustomFontFamily.IsSet() {
		return true
	}

	return false
}

// SetCustomFontFamily gets a reference to the given NullableString and assigns it to the CustomFontFamily field.
func (o *EnvironmentSettingsOut) SetCustomFontFamily(v string) {
	o.CustomFontFamily.Set(&v)
}
// SetCustomFontFamilyNil sets the value for CustomFontFamily to be an explicit nil
func (o *EnvironmentSettingsOut) SetCustomFontFamilyNil() {
	o.CustomFontFamily.Set(nil)
}

// UnsetCustomFontFamily ensures that no value is present for CustomFontFamily, not even an explicit nil
func (o *EnvironmentSettingsOut) UnsetCustomFontFamily() {
	o.CustomFontFamily.Unset()
}

// GetCustomLogoUrl returns the CustomLogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentSettingsOut) GetCustomLogoUrl() string {
	if o == nil || o.CustomLogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomLogoUrl.Get()
}

// GetCustomLogoUrlOk returns a tuple with the CustomLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentSettingsOut) GetCustomLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomLogoUrl.Get(), o.CustomLogoUrl.IsSet()
}

// HasCustomLogoUrl returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasCustomLogoUrl() bool {
	if o != nil && o.CustomLogoUrl.IsSet() {
		return true
	}

	return false
}

// SetCustomLogoUrl gets a reference to the given NullableString and assigns it to the CustomLogoUrl field.
func (o *EnvironmentSettingsOut) SetCustomLogoUrl(v string) {
	o.CustomLogoUrl.Set(&v)
}
// SetCustomLogoUrlNil sets the value for CustomLogoUrl to be an explicit nil
func (o *EnvironmentSettingsOut) SetCustomLogoUrlNil() {
	o.CustomLogoUrl.Set(nil)
}

// UnsetCustomLogoUrl ensures that no value is present for CustomLogoUrl, not even an explicit nil
func (o *EnvironmentSettingsOut) UnsetCustomLogoUrl() {
	o.CustomLogoUrl.Unset()
}

// GetCustomThemeOverride returns the CustomThemeOverride field value if set, zero value otherwise.
func (o *EnvironmentSettingsOut) GetCustomThemeOverride() CustomThemeOverride {
	if o == nil || o.CustomThemeOverride == nil {
		var ret CustomThemeOverride
		return ret
	}
	return *o.CustomThemeOverride
}

// GetCustomThemeOverrideOk returns a tuple with the CustomThemeOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSettingsOut) GetCustomThemeOverrideOk() (*CustomThemeOverride, bool) {
	if o == nil || o.CustomThemeOverride == nil {
		return nil, false
	}
	return o.CustomThemeOverride, true
}

// HasCustomThemeOverride returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasCustomThemeOverride() bool {
	if o != nil && o.CustomThemeOverride != nil {
		return true
	}

	return false
}

// SetCustomThemeOverride gets a reference to the given CustomThemeOverride and assigns it to the CustomThemeOverride field.
func (o *EnvironmentSettingsOut) SetCustomThemeOverride(v CustomThemeOverride) {
	o.CustomThemeOverride = &v
}

// GetEnableChannels returns the EnableChannels field value if set, zero value otherwise.
func (o *EnvironmentSettingsOut) GetEnableChannels() bool {
	if o == nil || o.EnableChannels == nil {
		var ret bool
		return ret
	}
	return *o.EnableChannels
}

// GetEnableChannelsOk returns a tuple with the EnableChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSettingsOut) GetEnableChannelsOk() (*bool, bool) {
	if o == nil || o.EnableChannels == nil {
		return nil, false
	}
	return o.EnableChannels, true
}

// HasEnableChannels returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasEnableChannels() bool {
	if o != nil && o.EnableChannels != nil {
		return true
	}

	return false
}

// SetEnableChannels gets a reference to the given bool and assigns it to the EnableChannels field.
func (o *EnvironmentSettingsOut) SetEnableChannels(v bool) {
	o.EnableChannels = &v
}

// GetEnableIntegrationManagement returns the EnableIntegrationManagement field value if set, zero value otherwise.
func (o *EnvironmentSettingsOut) GetEnableIntegrationManagement() bool {
	if o == nil || o.EnableIntegrationManagement == nil {
		var ret bool
		return ret
	}
	return *o.EnableIntegrationManagement
}

// GetEnableIntegrationManagementOk returns a tuple with the EnableIntegrationManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSettingsOut) GetEnableIntegrationManagementOk() (*bool, bool) {
	if o == nil || o.EnableIntegrationManagement == nil {
		return nil, false
	}
	return o.EnableIntegrationManagement, true
}

// HasEnableIntegrationManagement returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasEnableIntegrationManagement() bool {
	if o != nil && o.EnableIntegrationManagement != nil {
		return true
	}

	return false
}

// SetEnableIntegrationManagement gets a reference to the given bool and assigns it to the EnableIntegrationManagement field.
func (o *EnvironmentSettingsOut) SetEnableIntegrationManagement(v bool) {
	o.EnableIntegrationManagement = &v
}

// GetEnableTransformations returns the EnableTransformations field value if set, zero value otherwise.
func (o *EnvironmentSettingsOut) GetEnableTransformations() bool {
	if o == nil || o.EnableTransformations == nil {
		var ret bool
		return ret
	}
	return *o.EnableTransformations
}

// GetEnableTransformationsOk returns a tuple with the EnableTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSettingsOut) GetEnableTransformationsOk() (*bool, bool) {
	if o == nil || o.EnableTransformations == nil {
		return nil, false
	}
	return o.EnableTransformations, true
}

// HasEnableTransformations returns a boolean if a field has been set.
func (o *EnvironmentSettingsOut) HasEnableTransformations() bool {
	if o != nil && o.EnableTransformations != nil {
		return true
	}

	return false
}

// SetEnableTransformations gets a reference to the given bool and assigns it to the EnableTransformations field.
func (o *EnvironmentSettingsOut) SetEnableTransformations(v bool) {
	o.EnableTransformations = &v
}

func (o EnvironmentSettingsOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColorPaletteDark != nil {
		toSerialize["colorPaletteDark"] = o.ColorPaletteDark
	}
	if o.ColorPaletteLight != nil {
		toSerialize["colorPaletteLight"] = o.ColorPaletteLight
	}
	if o.CustomColor.IsSet() {
		toSerialize["customColor"] = o.CustomColor.Get()
	}
	if o.CustomFontFamily.IsSet() {
		toSerialize["customFontFamily"] = o.CustomFontFamily.Get()
	}
	if o.CustomLogoUrl.IsSet() {
		toSerialize["customLogoUrl"] = o.CustomLogoUrl.Get()
	}
	if o.CustomThemeOverride != nil {
		toSerialize["customThemeOverride"] = o.CustomThemeOverride
	}
	if o.EnableChannels != nil {
		toSerialize["enableChannels"] = o.EnableChannels
	}
	if o.EnableIntegrationManagement != nil {
		toSerialize["enableIntegrationManagement"] = o.EnableIntegrationManagement
	}
	if o.EnableTransformations != nil {
		toSerialize["enableTransformations"] = o.EnableTransformations
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentSettingsOut struct {
	value *EnvironmentSettingsOut
	isSet bool
}

func (v NullableEnvironmentSettingsOut) Get() *EnvironmentSettingsOut {
	return v.value
}

func (v *NullableEnvironmentSettingsOut) Set(val *EnvironmentSettingsOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSettingsOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSettingsOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSettingsOut(val *EnvironmentSettingsOut) *NullableEnvironmentSettingsOut {
	return &NullableEnvironmentSettingsOut{value: val, isSet: true}
}

func (v NullableEnvironmentSettingsOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSettingsOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


