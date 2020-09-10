package fetcher

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"strings"
	"testing"
)

type inputType struct {
	url  string
	html string
}

var flagtests = []struct {
	in  inputType
	out string
}{
	// Title と GetTitle 関数で取得した値が違う時
	{
		inputType{url: "http://hoge.com", html: "<title>hoge</title>"},
		"fuga",
	},
	// タイトルと GetTitle 関数で取得した値が同じ時
	{
		inputType{url: "http://foo.com", html: "<title>foo</title>"},
		"foo",
	},
	// Html が返ってこない時
	{
		inputType{url: "http://bar.com", html: ""},
		"",
	},
}

func Test_Fetcher(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in.url, func(t *testing.T) {
			defer gock.Off()
			gock.New(tt.in.url).
				Reply(200).
				Body(strings.NewReader(tt.in.html))

			title, err := GetTitle(tt.in.url)
			if err != nil {
				assert.Error(t, err)
			} else if title != tt.out {
				assert.NotEqual(t, title, tt.out)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, title, tt.out)
			}
		})
	}
}
