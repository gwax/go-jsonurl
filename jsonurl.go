package jsonurl

import (
	"encoding/json"
	"net/url"
)

type URL struct {
	url.URL
}

var _ interface {
	json.Marshaler
	json.Unmarshaler
} = (*URL)(nil)

func (u *URL) MarshalJSON() ([]byte, error) {
	if u == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(u.String())
}

func (u *URL) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	*u = URL{
		URL: *parsed,
	}
	return nil
}
