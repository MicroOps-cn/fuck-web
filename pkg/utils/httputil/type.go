package httputil

import (
	"net/url"

	w "github.com/MicroOps-cn/fuck/wrapper"
)

type Map[KT comparable, VT any] map[KT]VT

func (m Map[KT, VT]) String() string {
	return w.JSONStringer(m).String()
}

type URL url.URL

func (u URL) String() string {
	return (*url.URL)(&u).String()
}

func (u *URL) Set(s string) error {
	ou, err := url.Parse(s)
	if err != nil {
		return err
	}
	*u = URL(*ou)
	return nil
}

func (u *URL) Type() string {
	return "URL"
}
