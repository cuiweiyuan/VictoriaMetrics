//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/http/v3/cookie.proto

package httpv3

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

// Validate checks the field values on Cookie with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Cookie) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cookie with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CookieMultiError, or nil if none found.
func (m *Cookie) ValidateAll() error {
	return m.validate(true)
}

func (m *Cookie) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CookieValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetTtl(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = CookieValidationError{
				field:  "Ttl",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur < gte {
				err := CookieValidationError{
					field:  "Ttl",
					reason: "value must be greater than or equal to 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	// no validation rules for Path

	if len(errors) > 0 {
		return CookieMultiError(errors)
	}

	return nil
}

// CookieMultiError is an error wrapping multiple validation errors returned by
// Cookie.ValidateAll() if the designated constraints aren't met.
type CookieMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CookieMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CookieMultiError) AllErrors() []error { return m }

// CookieValidationError is the validation error returned by Cookie.Validate if
// the designated constraints aren't met.
type CookieValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CookieValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CookieValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CookieValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CookieValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CookieValidationError) ErrorName() string { return "CookieValidationError" }

// Error satisfies the builtin error interface
func (e CookieValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCookie.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CookieValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CookieValidationError{}