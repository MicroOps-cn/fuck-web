package common

import (
	"context"
	"net/url"

	http2 "github.com/MicroOps-cn/fuck/http"

	"github.com/MicroOps-cn/fuck-web/pkg/errors"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
)

type getWebURLOptions struct {
	subPages []string
	query    url.Values
}

type WithGetWebURLOptions func(*getWebURLOptions)

func WithSubPages(subPages ...string) WithGetWebURLOptions {
	return func(o *getWebURLOptions) {
		o.subPages = subPages
	}
}

func WithQuery(query url.Values) WithGetWebURLOptions {
	return func(o *getWebURLOptions) {
		o.query = query
	}
}

func GetWebURL(ctx context.Context, o ...WithGetWebURLOptions) (string, error) {
	var opts getWebURLOptions
	for _, options := range o {
		options(&opts)
	}

	adminPrefix, ok := ctx.Value(global.HTTPWebPrefixKey).(string)
	if !ok {
		return "", errors.NewServerError(500, "adminPrefix is null")
	}
	externalURL, ok := ctx.Value(global.HTTPExternalURLKey).(string)
	if !ok {
		return "", errors.NewServerError(500, "externalURL is null")
	}
	extURL, err := url.Parse(externalURL)
	if err != nil {
		return "", errors.NewServerError(500, "")
	}
	p := []string{extURL.Path, adminPrefix}
	if len(opts.subPages) > 0 {
		p = append(p, opts.subPages...)
	}
	extURL.Path = http2.JoinPath(p...)
	if len(opts.query) > 0 {
		q := extURL.Query()
		for name, vals := range opts.query {
			for _, val := range vals {
				if len(val) != 0 {
					q.Add(name, val)
				}
			}
		}
		extURL.RawQuery = q.Encode()
	}
	return extURL.String(), nil
}
