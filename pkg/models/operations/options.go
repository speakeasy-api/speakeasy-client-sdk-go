// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/retry"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

var ErrUnsupportedOption = errors.New("unsupported option")

const (
	SupportedOptionRetries              = "retries"
	SupportedOptionTimeout              = "timeout"
	SupportedOptionAcceptHeaderOverride = "acceptHeaderOverride"
	SupportedOptionURLOverride          = "urlOverride"
)

type AcceptHeaderEnum string

const (
	AcceptHeaderEnumApplicationJson                          AcceptHeaderEnum = "application/json"
	AcceptHeaderEnumApplicationXYaml                         AcceptHeaderEnum = "application/x-yaml"
	AcceptHeaderEnumApplicationOctetStream                   AcceptHeaderEnum = "application/octet-stream"
	AcceptHeaderEnumApplicationVndOciImageManifestV1PlusJson AcceptHeaderEnum = "application/vnd.oci.image.manifest.v1+json"
)

func (e AcceptHeaderEnum) ToPointer() *AcceptHeaderEnum {
	return &e
}

type Options struct {
	ServerURL            *string
	Retries              *retry.Config
	Timeout              *time.Duration
	AcceptHeaderOverride *AcceptHeaderEnum
	URLOverride          *string
	SetHeaders           map[string]string
}

type Option func(*Options, ...string) error

// WithServerURL allows providing an alternative server URL.
func WithServerURL(serverURL string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		opts.ServerURL = &serverURL
		return nil
	}
}

// WithTemplatedServerURL allows providing an alternative server URL with templated parameters.
func WithTemplatedServerURL(serverURL string, params map[string]string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		opts.ServerURL = &serverURL
		return nil
	}
}

// WithRetries allows customizing the default retry configuration.
func WithRetries(config retry.Config) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionRetries) {
			return ErrUnsupportedOption
		}

		opts.Retries = &config
		return nil
	}
}

// WithOperationTimeout allows setting the request timeout applied for an operation.
func WithOperationTimeout(timeout time.Duration) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionRetries) {
			return ErrUnsupportedOption
		}

		opts.Timeout = &timeout
		return nil
	}
}

func WithAcceptHeaderOverride(acceptHeaderOverride AcceptHeaderEnum) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionAcceptHeaderOverride) {
			return ErrUnsupportedOption
		}

		opts.AcceptHeaderOverride = &acceptHeaderOverride
		return nil
	}
}

// WithURLOverride allows overriding the URL.
func WithURLOverride(urlOverride string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionURLOverride) {
			return ErrUnsupportedOption
		}

		opts.URLOverride = &urlOverride
		return nil
	}
}

// WithSetHeaders takes a map of headers that will applied to a request. If the
// request contains headers that are in the map then they will be overwritten.
func WithSetHeaders(hdrs map[string]string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		opts.SetHeaders = hdrs
		return nil
	}
}
