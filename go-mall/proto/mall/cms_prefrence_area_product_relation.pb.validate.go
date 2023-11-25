// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/cms_prefrence_area_product_relation.proto

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

// Validate checks the field values on PrefrenceAreaProductRelation with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PrefrenceAreaProductRelation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PrefrenceAreaProductRelation with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PrefrenceAreaProductRelationMultiError, or nil if none found.
func (m *PrefrenceAreaProductRelation) ValidateAll() error {
	return m.validate(true)
}

func (m *PrefrenceAreaProductRelation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for PrefrenceAreaId

	// no validation rules for ProductId

	if len(errors) > 0 {
		return PrefrenceAreaProductRelationMultiError(errors)
	}

	return nil
}

// PrefrenceAreaProductRelationMultiError is an error wrapping multiple
// validation errors returned by PrefrenceAreaProductRelation.ValidateAll() if
// the designated constraints aren't met.
type PrefrenceAreaProductRelationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PrefrenceAreaProductRelationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PrefrenceAreaProductRelationMultiError) AllErrors() []error { return m }

// PrefrenceAreaProductRelationValidationError is the validation error returned
// by PrefrenceAreaProductRelation.Validate if the designated constraints
// aren't met.
type PrefrenceAreaProductRelationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrefrenceAreaProductRelationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrefrenceAreaProductRelationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrefrenceAreaProductRelationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrefrenceAreaProductRelationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrefrenceAreaProductRelationValidationError) ErrorName() string {
	return "PrefrenceAreaProductRelationValidationError"
}

// Error satisfies the builtin error interface
func (e PrefrenceAreaProductRelationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrefrenceAreaProductRelation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrefrenceAreaProductRelationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrefrenceAreaProductRelationValidationError{}
