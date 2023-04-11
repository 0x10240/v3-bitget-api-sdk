/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MarginCrossVipResult struct for MarginCrossVipResult
type MarginCrossVipResult struct {
	DailyInterestRate    *string `json:"dailyInterestRate,omitempty"`
	DiscountRate         *string `json:"discountRate,omitempty"`
	Level                *string `json:"level,omitempty"`
	YearlyInterestRate   *string `json:"yearlyInterestRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCrossVipResult MarginCrossVipResult

// NewMarginCrossVipResult instantiates a new MarginCrossVipResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCrossVipResult() *MarginCrossVipResult {
	this := MarginCrossVipResult{}
	return &this
}

// NewMarginCrossVipResultWithDefaults instantiates a new MarginCrossVipResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCrossVipResultWithDefaults() *MarginCrossVipResult {
	this := MarginCrossVipResult{}
	return &this
}

// GetDailyInterestRate returns the DailyInterestRate field value if set, zero value otherwise.
func (o *MarginCrossVipResult) GetDailyInterestRate() string {
	if o == nil || isNil(o.DailyInterestRate) {
		var ret string
		return ret
	}
	return *o.DailyInterestRate
}

// GetDailyInterestRateOk returns a tuple with the DailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossVipResult) GetDailyInterestRateOk() (*string, bool) {
	if o == nil || isNil(o.DailyInterestRate) {
		return nil, false
	}
	return o.DailyInterestRate, true
}

// HasDailyInterestRate returns a boolean if a field has been set.
func (o *MarginCrossVipResult) HasDailyInterestRate() bool {
	if o != nil && !isNil(o.DailyInterestRate) {
		return true
	}

	return false
}

// SetDailyInterestRate gets a reference to the given string and assigns it to the DailyInterestRate field.
func (o *MarginCrossVipResult) SetDailyInterestRate(v string) {
	o.DailyInterestRate = &v
}

// GetDiscountRate returns the DiscountRate field value if set, zero value otherwise.
func (o *MarginCrossVipResult) GetDiscountRate() string {
	if o == nil || isNil(o.DiscountRate) {
		var ret string
		return ret
	}
	return *o.DiscountRate
}

// GetDiscountRateOk returns a tuple with the DiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossVipResult) GetDiscountRateOk() (*string, bool) {
	if o == nil || isNil(o.DiscountRate) {
		return nil, false
	}
	return o.DiscountRate, true
}

// HasDiscountRate returns a boolean if a field has been set.
func (o *MarginCrossVipResult) HasDiscountRate() bool {
	if o != nil && !isNil(o.DiscountRate) {
		return true
	}

	return false
}

// SetDiscountRate gets a reference to the given string and assigns it to the DiscountRate field.
func (o *MarginCrossVipResult) SetDiscountRate(v string) {
	o.DiscountRate = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *MarginCrossVipResult) GetLevel() string {
	if o == nil || isNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossVipResult) GetLevelOk() (*string, bool) {
	if o == nil || isNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *MarginCrossVipResult) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *MarginCrossVipResult) SetLevel(v string) {
	o.Level = &v
}

// GetYearlyInterestRate returns the YearlyInterestRate field value if set, zero value otherwise.
func (o *MarginCrossVipResult) GetYearlyInterestRate() string {
	if o == nil || isNil(o.YearlyInterestRate) {
		var ret string
		return ret
	}
	return *o.YearlyInterestRate
}

// GetYearlyInterestRateOk returns a tuple with the YearlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossVipResult) GetYearlyInterestRateOk() (*string, bool) {
	if o == nil || isNil(o.YearlyInterestRate) {
		return nil, false
	}
	return o.YearlyInterestRate, true
}

// HasYearlyInterestRate returns a boolean if a field has been set.
func (o *MarginCrossVipResult) HasYearlyInterestRate() bool {
	if o != nil && !isNil(o.YearlyInterestRate) {
		return true
	}

	return false
}

// SetYearlyInterestRate gets a reference to the given string and assigns it to the YearlyInterestRate field.
func (o *MarginCrossVipResult) SetYearlyInterestRate(v string) {
	o.YearlyInterestRate = &v
}

func (o MarginCrossVipResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DailyInterestRate) {
		toSerialize["dailyInterestRate"] = o.DailyInterestRate
	}
	if !isNil(o.DiscountRate) {
		toSerialize["discountRate"] = o.DiscountRate
	}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !isNil(o.YearlyInterestRate) {
		toSerialize["yearlyInterestRate"] = o.YearlyInterestRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginCrossVipResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginCrossVipResult := _MarginCrossVipResult{}

	if err = json.Unmarshal(bytes, &varMarginCrossVipResult); err == nil {
		*o = MarginCrossVipResult(varMarginCrossVipResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dailyInterestRate")
		delete(additionalProperties, "discountRate")
		delete(additionalProperties, "level")
		delete(additionalProperties, "yearlyInterestRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCrossVipResult struct {
	value *MarginCrossVipResult
	isSet bool
}

func (v NullableMarginCrossVipResult) Get() *MarginCrossVipResult {
	return v.value
}

func (v *NullableMarginCrossVipResult) Set(val *MarginCrossVipResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCrossVipResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCrossVipResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCrossVipResult(val *MarginCrossVipResult) *NullableMarginCrossVipResult {
	return &NullableMarginCrossVipResult{value: val, isSet: true}
}

func (v NullableMarginCrossVipResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCrossVipResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}