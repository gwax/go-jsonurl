package jsonurl_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/gwax/go-jsonurl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func urlMust(t *testing.T) func(*url.URL, error) url.URL {
	return func(u *url.URL, e error) url.URL {
		require.NoError(t, e)
		return *u
	}
}

func TestJsonURLRoundtrip(t *testing.T) {
	testURLs := []string{
		"https://github.com/gwax/go-jsonurl",
		"github.com",
		"https://www.foo.bar.co.baz/thingstuff?hello=thing&stuff=other",
	}

	for _, u := range testURLs {
		u := u // capture loop variable: // https://golang.org/doc/faq#closures_and_goroutines
		qu := fmt.Sprintf("%q", u)
		t.Run(u, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				ju := &jsonurl.URL{
					URL: urlMust(t)(url.Parse(u)),
				}
				b, err := json.Marshal(ju)
				require.NoError(t, err)
				assert.Equal(t, qu, string(b))
			})
			t.Run("unmarshal", func(t *testing.T) {
				var ju jsonurl.URL
				err := json.Unmarshal([]byte(qu), &ju)
				require.NoError(t, err)
				assert.Equal(t, u, ju.String())
			})
		})
	}
}

func TestUnmarshalJsonURL(t *testing.T) {
	testCases := []struct {
		value        []byte
		expected     *jsonurl.URL
		errAssertion require.ErrorAssertionFunc
	}{
		{
			value:        []byte(`github.com`),
			errAssertion: require.Error,
		},
		{
			value:        []byte(`"https://www.foo.bar.co.baz/thingstuff?hello=thing&stuff=other"`),
			expected:     &jsonurl.URL{},
			errAssertion: require.NoError,
		},
	}
	for _, tc := range testCases {
		tc := tc // capture loop variable: // https://golang.org/doc/faq#closures_and_goroutines
		t.Run(string(tc.value), func(t *testing.T) {
			var ju jsonurl.URL
			err := json.Unmarshal(tc.value, &ju)
			tc.errAssertion(t, err)
			if tc.expected != nil {
				assert.Equal(t, tc.expected, &ju)
			}
		})
	}
}
