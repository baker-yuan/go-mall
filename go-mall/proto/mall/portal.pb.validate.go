// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: portal/portal.proto

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

// Validate checks the field values on EmptyRsp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyRsp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyRspMultiError, or nil
// if none found.
func (m *EmptyRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyRspMultiError(errors)
	}

	return nil
}

// EmptyRspMultiError is an error wrapping multiple validation errors returned
// by EmptyRsp.ValidateAll() if the designated constraints aren't met.
type EmptyRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyRspMultiError) AllErrors() []error { return m }

// EmptyRspValidationError is the validation error returned by
// EmptyRsp.Validate if the designated constraints aren't met.
type EmptyRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyRspValidationError) ErrorName() string { return "EmptyRspValidationError" }

// Error satisfies the builtin error interface
func (e EmptyRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyRspValidationError{}
