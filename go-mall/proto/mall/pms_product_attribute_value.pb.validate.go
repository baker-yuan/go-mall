// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/pms_product_attribute_value.proto

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

// Validate checks the field values on ProductAttributeValue with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProductAttributeValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductAttributeValue with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProductAttributeValueMultiError, or nil if none found.
func (m *ProductAttributeValue) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductAttributeValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ProductId

	// no validation rules for ProductAttributeId

	// no validation rules for Value

	if len(errors) > 0 {
		return ProductAttributeValueMultiError(errors)
	}

	return nil
}

// ProductAttributeValueMultiError is an error wrapping multiple validation
// errors returned by ProductAttributeValue.ValidateAll() if the designated
// constraints aren't met.
type ProductAttributeValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductAttributeValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductAttributeValueMultiError) AllErrors() []error { return m }

// ProductAttributeValueValidationError is the validation error returned by
// ProductAttributeValue.Validate if the designated constraints aren't met.
type ProductAttributeValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductAttributeValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductAttributeValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductAttributeValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductAttributeValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductAttributeValueValidationError) ErrorName() string {
	return "ProductAttributeValueValidationError"
}

// Error satisfies the builtin error interface
func (e ProductAttributeValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductAttributeValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductAttributeValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductAttributeValueValidationError{}
