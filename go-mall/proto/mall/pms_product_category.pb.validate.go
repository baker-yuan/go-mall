// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/pms_product_category.proto

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

// Validate checks the field values on ProductCategory with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProductCategory) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductCategory with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductCategoryMultiError, or nil if none found.
func (m *ProductCategory) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductCategory) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ParentId

	// no validation rules for Name

	// no validation rules for Icon

	// no validation rules for ProductUnit

	// no validation rules for Sort

	// no validation rules for Description

	// no validation rules for Keywords

	// no validation rules for NavStatus

	// no validation rules for ShowStatus

	// no validation rules for Level

	// no validation rules for ProductCount

	if len(errors) > 0 {
		return ProductCategoryMultiError(errors)
	}

	return nil
}

// ProductCategoryMultiError is an error wrapping multiple validation errors
// returned by ProductCategory.ValidateAll() if the designated constraints
// aren't met.
type ProductCategoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductCategoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductCategoryMultiError) AllErrors() []error { return m }

// ProductCategoryValidationError is the validation error returned by
// ProductCategory.Validate if the designated constraints aren't met.
type ProductCategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductCategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductCategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductCategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductCategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductCategoryValidationError) ErrorName() string { return "ProductCategoryValidationError" }

// Error satisfies the builtin error interface
func (e ProductCategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductCategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductCategoryValidationError{}
