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
	err  bool
}

var flagtests = []struct {
	in  inputType
	out string
}{
	// タイトルと GetTitle 関数で取得した値が同じ時
	{
		inputType{url: "http://foo.com", html: "<title>foo</title>", err: false},
		"foo",
	},
	// Html が返ってこない時 (異常系)
	{
		inputType{url: "http://bar.com", html: "", err: true},
		"",
	},
	// TODO 他の異常系のケースも実装
}

func Test_Fetcher(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in.url, func(t *testing.T) {
			defer gock.Off()
			gock.New(tt.in.url).
				Reply(200).
				Body(strings.NewReader(tt.in.html))

			title, err := GetTitle(tt.in.url)

			if tt.in.err {
				if err != nil {
					assert.Error(t, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, title, tt.out)
			}
		})
	}
}
