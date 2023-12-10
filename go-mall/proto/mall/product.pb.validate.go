// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: portal/product.proto

package mall

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ProductItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProductItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProductItemMultiError, or
// nil if none found.
func (m *ProductItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ProductCategoryId

	// no validation rules for Name

	if len(errors) > 0 {
		return ProductItemMultiError(errors)
	}

	return nil
}

// ProductItemMultiError is an error wrapping multiple validation errors
// returned by ProductItem.ValidateAll() if the designated constraints aren't met.
type ProductItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductItemMultiError) AllErrors() []error { return m }

// ProductItemValidationError is the validation error returned by
// ProductItem.Validate if the designated constraints aren't met.
type ProductItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductItemValidationError) ErrorName() string { return "ProductItemValidationError" }

// Error satisfies the builtin error interface
func (e ProductItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductItemValidationError{}

// Validate checks the field values on SearchProductReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SearchProductReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchProductReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchProductReqMultiError, or nil if none found.
func (m *SearchProductReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchProductReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Keyword

	if all {
		switch v := interface{}(m.GetBrandId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SearchProductReqValidationError{
					field:  "BrandId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SearchProductReqValidationError{
					field:  "BrandId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBrandId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SearchProductReqValidationError{
				field:  "BrandId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetProductCategoryId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SearchProductReqValidationError{
					field:  "ProductCategoryId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SearchProductReqValidationError{
					field:  "ProductCategoryId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProductCategoryId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SearchProductReqValidationError{
				field:  "ProductCategoryId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PageNum

	// no validation rules for PageSize

	if _, ok := _SearchProductReq_Sort_InLookup[m.GetSort()]; !ok {
		err := SearchProductReqValidationError{
			field:  "Sort",
			reason: "value must be in list [0 1 2 3 4]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SearchProductReqMultiError(errors)
	}

	return nil
}

// SearchProductReqMultiError is an error wrapping multiple validation errors
// returned by SearchProductReq.ValidateAll() if the designated constraints
// aren't met.
type SearchProductReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchProductReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchProductReqMultiError) AllErrors() []error { return m }

// SearchProductReqValidationError is the validation error returned by
// SearchProductReq.Validate if the designated constraints aren't met.
type SearchProductReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchProductReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchProductReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchProductReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchProductReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchProductReqValidationError) ErrorName() string { return "SearchProductReqValidationError" }

// Error satisfies the builtin error interface
func (e SearchProductReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchProductReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchProductReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchProductReqValidationError{}

var _SearchProductReq_Sort_InLookup = map[uint32]struct{}{
	0: {},
	1: {},
	2: {},
	3: {},
	4: {},
}

// Validate checks the field values on ProductRspItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProductRspItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductRspItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProductRspItemMultiError,
// or nil if none found.
func (m *ProductRspItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductRspItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProductRspItemValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProductRspItemValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductRspItemValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	// no validation rules for PageTotal

	// no validation rules for PageSize

	// no validation rules for PageNum

	if len(errors) > 0 {
		return ProductRspItemMultiError(errors)
	}

	return nil
}

// ProductRspItemMultiError is an error wrapping multiple validation errors
// returned by ProductRspItem.ValidateAll() if the designated constraints
// aren't met.
type ProductRspItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductRspItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductRspItemMultiError) AllErrors() []error { return m }

// ProductRspItemValidationError is the validation error returned by
// ProductRspItem.Validate if the designated constraints aren't met.
type ProductRspItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductRspItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductRspItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductRspItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductRspItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductRspItemValidationError) ErrorName() string { return "ProductRspItemValidationError" }

// Error satisfies the builtin error interface
func (e ProductRspItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductRspItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductRspItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductRspItemValidationError{}

// Validate checks the field values on SearchProductRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SearchProductRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchProductRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchProductRspMultiError, or nil if none found.
func (m *SearchProductRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchProductRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SearchProductRspValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SearchProductRspValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SearchProductRspValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SearchProductRspMultiError(errors)
	}

	return nil
}

// SearchProductRspMultiError is an error wrapping multiple validation errors
// returned by SearchProductRsp.ValidateAll() if the designated constraints
// aren't met.
type SearchProductRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchProductRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchProductRspMultiError) AllErrors() []error { return m }

// SearchProductRspValidationError is the validation error returned by
// SearchProductRsp.Validate if the designated constraints aren't met.
type SearchProductRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchProductRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchProductRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchProductRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchProductRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchProductRspValidationError) ErrorName() string { return "SearchProductRspValidationError" }

// Error satisfies the builtin error interface
func (e SearchProductRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchProductRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchProductRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchProductRspValidationError{}
