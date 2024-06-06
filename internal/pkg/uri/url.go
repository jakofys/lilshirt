package uri

import "net/url"

type URL struct {
	url.URL
}

func Parse(raw string) (*URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	return &URL{*u}, nil
}

func (u *URL) MarshalText() ([]byte, error) {
	return u.MarshalBinary()
}
