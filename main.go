package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Test struct {
	URL *URL
}

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

func main() {
	u, err := Parse("https://fluid.io/path/param")
	if err != nil {
		panic(err)
	}
	t := &Test{URL: u}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
