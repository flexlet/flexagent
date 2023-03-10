// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// QueryURL generates an URL for the query operation
type QueryURL struct {
	Urn string

	OutputLineLimit *int32
	OutputLineStart *int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QueryURL) WithBasePath(bp string) *QueryURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QueryURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *QueryURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/jobs/{urn}"

	urn := o.Urn
	if urn != "" {
		_path = strings.Replace(_path, "{urn}", urn, -1)
	} else {
		return nil, errors.New("urn is required on QueryURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var outputLineLimitQ string
	if o.OutputLineLimit != nil {
		outputLineLimitQ = swag.FormatInt32(*o.OutputLineLimit)
	}
	if outputLineLimitQ != "" {
		qs.Set("outputLineLimit", outputLineLimitQ)
	}

	var outputLineStartQ string
	if o.OutputLineStart != nil {
		outputLineStartQ = swag.FormatInt32(*o.OutputLineStart)
	}
	if outputLineStartQ != "" {
		qs.Set("outputLineStart", outputLineStartQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *QueryURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *QueryURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *QueryURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on QueryURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on QueryURL")
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
func (o *QueryURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
