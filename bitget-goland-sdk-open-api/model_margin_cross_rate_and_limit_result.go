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

// MarginCrossRateAndLimitResult struct for MarginCrossRateAndLimitResult
type MarginCrossRateAndLimitResult struct {
	BorrowAble           *bool                  `json:"borrowAble,omitempty"`
	Coin                 *string                `json:"coin,omitempty"`
	DailyInterestRate    *string                `json:"dailyInterestRate,omitempty"`
	Leverage             *string                `json:"leverage,omitempty"`
	MaxBorrowableAmount  *string                `json:"maxBorrowableAmount,omitempty"`
	TransferInAble       *bool                  `json:"transferInAble,omitempty"`
	Vips                 []MarginCrossVipResult `json:"vips,omitempty"`
	YearlyInterestRate   *string                `json:"yearlyInterestRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCrossRateAndLimitResult MarginCrossRateAndLimitResult

// NewMarginCrossRateAndLimitResult instantiates a new MarginCrossRateAndLimitResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCrossRateAndLimitResult() *MarginCrossRateAndLimitResult {
	this := MarginCrossRateAndLimitResult{}
	return &this
}

// NewMarginCrossRateAndLimitResultWithDefaults instantiates a new MarginCrossRateAndLimitResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCrossRateAndLimitResultWithDefaults() *MarginCrossRateAndLimitResult {
	this := MarginCrossRateAndLimitResult{}
	return &this
}

// GetBorrowAble returns the BorrowAble field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetBorrowAble() bool {
	if o == nil || isNil(o.BorrowAble) {
		var ret bool
		return ret
	}
	return *o.BorrowAble
}

// GetBorrowAbleOk returns a tuple with the BorrowAble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetBorrowAbleOk() (*bool, bool) {
	if o == nil || isNil(o.BorrowAble) {
		return nil, false
	}
	return o.BorrowAble, true
}

// HasBorrowAble returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasBorrowAble() bool {
	if o != nil && !isNil(o.BorrowAble) {
		return true
	}

	return false
}

// SetBorrowAble gets a reference to the given bool and assigns it to the BorrowAble field.
func (o *MarginCrossRateAndLimitResult) SetBorrowAble(v bool) {
	o.BorrowAble = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MarginCrossRateAndLimitResult) SetCoin(v string) {
	o.Coin = &v
}

// GetDailyInterestRate returns the DailyInterestRate field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetDailyInterestRate() string {
	if o == nil || isNil(o.DailyInterestRate) {
		var ret string
		return ret
	}
	return *o.DailyInterestRate
}

// GetDailyInterestRateOk returns a tuple with the DailyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetDailyInterestRateOk() (*string, bool) {
	if o == nil || isNil(o.DailyInterestRate) {
		return nil, false
	}
	return o.DailyInterestRate, true
}

// HasDailyInterestRate returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasDailyInterestRate() bool {
	if o != nil && !isNil(o.DailyInterestRate) {
		return true
	}

	return false
}

// SetDailyInterestRate gets a reference to the given string and assigns it to the DailyInterestRate field.
func (o *MarginCrossRateAndLimitResult) SetDailyInterestRate(v string) {
	o.DailyInterestRate = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetLeverage() string {
	if o == nil || isNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetLeverageOk() (*string, bool) {
	if o == nil || isNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasLeverage() bool {
	if o != nil && !isNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *MarginCrossRateAndLimitResult) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMaxBorrowableAmount returns the MaxBorrowableAmount field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetMaxBorrowableAmount() string {
	if o == nil || isNil(o.MaxBorrowableAmount) {
		var ret string
		return ret
	}
	return *o.MaxBorrowableAmount
}

// GetMaxBorrowableAmountOk returns a tuple with the MaxBorrowableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetMaxBorrowableAmountOk() (*string, bool) {
	if o == nil || isNil(o.MaxBorrowableAmount) {
		return nil, false
	}
	return o.MaxBorrowableAmount, true
}

// HasMaxBorrowableAmount returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasMaxBorrowableAmount() bool {
	if o != nil && !isNil(o.MaxBorrowableAmount) {
		return true
	}

	return false
}

// SetMaxBorrowableAmount gets a reference to the given string and assigns it to the MaxBorrowableAmount field.
func (o *MarginCrossRateAndLimitResult) SetMaxBorrowableAmount(v string) {
	o.MaxBorrowableAmount = &v
}

// GetTransferInAble returns the TransferInAble field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetTransferInAble() bool {
	if o == nil || isNil(o.TransferInAble) {
		var ret bool
		return ret
	}
	return *o.TransferInAble
}

// GetTransferInAbleOk returns a tuple with the TransferInAble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetTransferInAbleOk() (*bool, bool) {
	if o == nil || isNil(o.TransferInAble) {
		return nil, false
	}
	return o.TransferInAble, true
}

// HasTransferInAble returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasTransferInAble() bool {
	if o != nil && !isNil(o.TransferInAble) {
		return true
	}

	return false
}

// SetTransferInAble gets a reference to the given bool and assigns it to the TransferInAble field.
func (o *MarginCrossRateAndLimitResult) SetTransferInAble(v bool) {
	o.TransferInAble = &v
}

// GetVips returns the Vips field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetVips() []MarginCrossVipResult {
	if o == nil || isNil(o.Vips) {
		var ret []MarginCrossVipResult
		return ret
	}
	return o.Vips
}

// GetVipsOk returns a tuple with the Vips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetVipsOk() ([]MarginCrossVipResult, bool) {
	if o == nil || isNil(o.Vips) {
		return nil, false
	}
	return o.Vips, true
}

// HasVips returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasVips() bool {
	if o != nil && !isNil(o.Vips) {
		return true
	}

	return false
}

// SetVips gets a reference to the given []MarginCrossVipResult and assigns it to the Vips field.
func (o *MarginCrossRateAndLimitResult) SetVips(v []MarginCrossVipResult) {
	o.Vips = v
}

// GetYearlyInterestRate returns the YearlyInterestRate field value if set, zero value otherwise.
func (o *MarginCrossRateAndLimitResult) GetYearlyInterestRate() string {
	if o == nil || isNil(o.YearlyInterestRate) {
		var ret string
		return ret
	}
	return *o.YearlyInterestRate
}

// GetYearlyInterestRateOk returns a tuple with the YearlyInterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCrossRateAndLimitResult) GetYearlyInterestRateOk() (*string, bool) {
	if o == nil || isNil(o.YearlyInterestRate) {
		return nil, false
	}
	return o.YearlyInterestRate, true
}

// HasYearlyInterestRate returns a boolean if a field has been set.
func (o *MarginCrossRateAndLimitResult) HasYearlyInterestRate() bool {
	if o != nil && !isNil(o.YearlyInterestRate) {
		return true
	}

	return false
}

// SetYearlyInterestRate gets a reference to the given string and assigns it to the YearlyInterestRate field.
func (o *MarginCrossRateAndLimitResult) SetYearlyInterestRate(v string) {
	o.YearlyInterestRate = &v
}

func (o MarginCrossRateAndLimitResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BorrowAble) {
		toSerialize["borrowAble"] = o.BorrowAble
	}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.DailyInterestRate) {
		toSerialize["dailyInterestRate"] = o.DailyInterestRate
	}
	if !isNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !isNil(o.MaxBorrowableAmount) {
		toSerialize["maxBorrowableAmount"] = o.MaxBorrowableAmount
	}
	if !isNil(o.TransferInAble) {
		toSerialize["transferInAble"] = o.TransferInAble
	}
	if !isNil(o.Vips) {
		toSerialize["vips"] = o.Vips
	}
	if !isNil(o.YearlyInterestRate) {
		toSerialize["yearlyInterestRate"] = o.YearlyInterestRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginCrossRateAndLimitResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginCrossRateAndLimitResult := _MarginCrossRateAndLimitResult{}

	if err = json.Unmarshal(bytes, &varMarginCrossRateAndLimitResult); err == nil {
		*o = MarginCrossRateAndLimitResult(varMarginCrossRateAndLimitResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "borrowAble")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "dailyInterestRate")
		delete(additionalProperties, "leverage")
		delete(additionalProperties, "maxBorrowableAmount")
		delete(additionalProperties, "transferInAble")
		delete(additionalProperties, "vips")
		delete(additionalProperties, "yearlyInterestRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCrossRateAndLimitResult struct {
	value *MarginCrossRateAndLimitResult
	isSet bool
}

func (v NullableMarginCrossRateAndLimitResult) Get() *MarginCrossRateAndLimitResult {
	return v.value
}

func (v *NullableMarginCrossRateAndLimitResult) Set(val *MarginCrossRateAndLimitResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCrossRateAndLimitResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCrossRateAndLimitResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCrossRateAndLimitResult(val *MarginCrossRateAndLimitResult) *NullableMarginCrossRateAndLimitResult {
	return &NullableMarginCrossRateAndLimitResult{value: val, isSet: true}
}

func (v NullableMarginCrossRateAndLimitResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCrossRateAndLimitResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}