// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/cms_subject.proto

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

// Validate checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Subject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SubjectMultiError, or nil if none found.
func (m *Subject) ValidateAll() error {
	return m.validate(true)
}

func (m *Subject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CategoryId

	// no validation rules for Title

	// no validation rules for Pic

	// no validation rules for AlbumPics

	// no validation rules for Description

	// no validation rules for Content

	// no validation rules for ShowStatus

	// no validation rules for RecommendStatus

	// no validation rules for CreateTime

	// no validation rules for ForwardCount

	// no validation rules for CollectCount

	// no validation rules for ReadCount

	// no validation rules for CommentCount

	// no validation rules for ProductCount

	// no validation rules for CategoryName

	if len(errors) > 0 {
		return SubjectMultiError(errors)
	}

	return nil
}

// SubjectMultiError is an error wrapping multiple validation errors returned
// by Subject.ValidateAll() if the designated constraints aren't met.
type SubjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectMultiError) AllErrors() []error { return m }

// SubjectValidationError is the validation error returned by Subject.Validate
// if the designated constraints aren't met.
type SubjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectValidationError) ErrorName() string { return "SubjectValidationError" }

// Error satisfies the builtin error interface
func (e SubjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectValidationError{}
