// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/async_files/v3/async_file_manager.proto

package async_filesv3

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

// Validate checks the field values on AsyncFileManagerConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AsyncFileManagerConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AsyncFileManagerConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AsyncFileManagerConfigMultiError, or nil if none found.
func (m *AsyncFileManagerConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *AsyncFileManagerConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	switch m.ManagerType.(type) {

	case *AsyncFileManagerConfig_ThreadPool_:

		if all {
			switch v := interface{}(m.GetThreadPool()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AsyncFileManagerConfigValidationError{
						field:  "ThreadPool",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AsyncFileManagerConfigValidationError{
						field:  "ThreadPool",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetThreadPool()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AsyncFileManagerConfigValidationError{
					field:  "ThreadPool",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := AsyncFileManagerConfigValidationError{
			field:  "ManagerType",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return AsyncFileManagerConfigMultiError(errors)
	}

	return nil
}

// AsyncFileManagerConfigMultiError is an error wrapping multiple validation
// errors returned by AsyncFileManagerConfig.ValidateAll() if the designated
// constraints aren't met.
type AsyncFileManagerConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AsyncFileManagerConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AsyncFileManagerConfigMultiError) AllErrors() []error { return m }

// AsyncFileManagerConfigValidationError is the validation error returned by
// AsyncFileManagerConfig.Validate if the designated constraints aren't met.
type AsyncFileManagerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AsyncFileManagerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AsyncFileManagerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AsyncFileManagerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AsyncFileManagerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AsyncFileManagerConfigValidationError) ErrorName() string {
	return "AsyncFileManagerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e AsyncFileManagerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAsyncFileManagerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AsyncFileManagerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AsyncFileManagerConfigValidationError{}

// Validate checks the field values on AsyncFileManagerConfig_ThreadPool with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *AsyncFileManagerConfig_ThreadPool) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AsyncFileManagerConfig_ThreadPool
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// AsyncFileManagerConfig_ThreadPoolMultiError, or nil if none found.
func (m *AsyncFileManagerConfig_ThreadPool) ValidateAll() error {
	return m.validate(true)
}

func (m *AsyncFileManagerConfig_ThreadPool) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ThreadCount

	if len(errors) > 0 {
		return AsyncFileManagerConfig_ThreadPoolMultiError(errors)
	}

	return nil
}

// AsyncFileManagerConfig_ThreadPoolMultiError is an error wrapping multiple
// validation errors returned by
// AsyncFileManagerConfig_ThreadPool.ValidateAll() if the designated
// constraints aren't met.
type AsyncFileManagerConfig_ThreadPoolMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AsyncFileManagerConfig_ThreadPoolMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AsyncFileManagerConfig_ThreadPoolMultiError) AllErrors() []error { return m }

// AsyncFileManagerConfig_ThreadPoolValidationError is the validation error
// returned by AsyncFileManagerConfig_ThreadPool.Validate if the designated
// constraints aren't met.
type AsyncFileManagerConfig_ThreadPoolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AsyncFileManagerConfig_ThreadPoolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AsyncFileManagerConfig_ThreadPoolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AsyncFileManagerConfig_ThreadPoolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AsyncFileManagerConfig_ThreadPoolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AsyncFileManagerConfig_ThreadPoolValidationError) ErrorName() string {
	return "AsyncFileManagerConfig_ThreadPoolValidationError"
}

// Error satisfies the builtin error interface
func (e AsyncFileManagerConfig_ThreadPoolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAsyncFileManagerConfig_ThreadPool.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AsyncFileManagerConfig_ThreadPoolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AsyncFileManagerConfig_ThreadPoolValidationError{}
