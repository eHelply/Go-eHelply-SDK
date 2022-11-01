/*
eHelply SDK - 1.1.121

eHelply SDK for SuperStack Services

API version: 1.1.121
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Pagination Pagination state
type Pagination struct {
	CurrentPage int32 `json:"current_page"`
	PageSize int32 `json:"page_size"`
	TotalItems int32 `json:"total_items"`
	TotalPages int32 `json:"total_pages"`
	HasPreviousPage bool `json:"has_previous_page"`
	HasNextPage bool `json:"has_next_page"`
	PreviousPage *int32 `json:"previous_page,omitempty"`
	NextPage *int32 `json:"next_page,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination(currentPage int32, pageSize int32, totalItems int32, totalPages int32, hasPreviousPage bool, hasNextPage bool) *Pagination {
	this := Pagination{}
	this.CurrentPage = currentPage
	this.PageSize = pageSize
	this.TotalItems = totalItems
	this.TotalPages = totalPages
	this.HasPreviousPage = hasPreviousPage
	this.HasNextPage = hasNextPage
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value
func (o *Pagination) GetCurrentPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetCurrentPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *Pagination) SetCurrentPage(v int32) {
	o.CurrentPage = v
}

// GetPageSize returns the PageSize field value
func (o *Pagination) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *Pagination) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotalItems returns the TotalItems field value
func (o *Pagination) GetTotalItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalItems, true
}

// SetTotalItems sets field value
func (o *Pagination) SetTotalItems(v int32) {
	o.TotalItems = v
}

// GetTotalPages returns the TotalPages field value
func (o *Pagination) GetTotalPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *Pagination) SetTotalPages(v int32) {
	o.TotalPages = v
}

// GetHasPreviousPage returns the HasPreviousPage field value
func (o *Pagination) GetHasPreviousPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPreviousPage
}

// GetHasPreviousPageOk returns a tuple with the HasPreviousPage field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetHasPreviousPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPreviousPage, true
}

// SetHasPreviousPage sets field value
func (o *Pagination) SetHasPreviousPage(v bool) {
	o.HasPreviousPage = v
}

// GetHasNextPage returns the HasNextPage field value
func (o *Pagination) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *Pagination) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *Pagination) GetPreviousPage() int32 {
	if o == nil || o.PreviousPage == nil {
		var ret int32
		return ret
	}
	return *o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetPreviousPageOk() (*int32, bool) {
	if o == nil || o.PreviousPage == nil {
		return nil, false
	}
	return o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *Pagination) HasPreviousPage() bool {
	if o != nil && o.PreviousPage != nil {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given int32 and assigns it to the PreviousPage field.
func (o *Pagination) SetPreviousPage(v int32) {
	o.PreviousPage = &v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *Pagination) GetNextPage() int32 {
	if o == nil || o.NextPage == nil {
		var ret int32
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetNextPageOk() (*int32, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *Pagination) HasNextPage() bool {
	if o != nil && o.NextPage != nil {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given int32 and assigns it to the NextPage field.
func (o *Pagination) SetNextPage(v int32) {
	o.NextPage = &v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_page"] = o.CurrentPage
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if true {
		toSerialize["total_items"] = o.TotalItems
	}
	if true {
		toSerialize["total_pages"] = o.TotalPages
	}
	if true {
		toSerialize["has_previous_page"] = o.HasPreviousPage
	}
	if true {
		toSerialize["has_next_page"] = o.HasNextPage
	}
	if o.PreviousPage != nil {
		toSerialize["previous_page"] = o.PreviousPage
	}
	if o.NextPage != nil {
		toSerialize["next_page"] = o.NextPage
	}
	return json.Marshal(toSerialize)
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


