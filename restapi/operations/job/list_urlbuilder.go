// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ListURL generates an URL for the list operation
type ListURL struct {
	Operation      *string
	Plugin         *string
	StartTimeBegin *int64
	StartTimeEnd   *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListURL) WithBasePath(bp string) *ListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/jobs"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var operationQ string
	if o.Operation != nil {
		operationQ = *o.Operation
	}
	if operationQ != "" {
		qs.Set("operation", operationQ)
	}

	var pluginQ string
	if o.Plugin != nil {
		pluginQ = *o.Plugin
	}
	if pluginQ != "" {
		qs.Set("plugin", pluginQ)
	}

	var startTimeBeginQ string
	if o.StartTimeBegin != nil {
		startTimeBeginQ = swag.FormatInt64(*o.StartTimeBegin)
	}
	if startTimeBeginQ != "" {
		qs.Set("startTimeBegin", startTimeBeginQ)
	}

	var startTimeEndQ string
	if o.StartTimeEnd != nil {
		startTimeEndQ = swag.FormatInt64(*o.StartTimeEnd)
	}
	if startTimeEndQ != "" {
		qs.Set("startTimeEnd", startTimeEndQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
