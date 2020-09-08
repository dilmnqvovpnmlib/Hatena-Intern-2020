package renderer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var flagtests = []struct {
	in  string
	out string
}{
	// 記法のテスト
	{`# 見出し`, "<h1>見出し</h1>\n"},
	// リストのテスト
	{"- リスト", "<ul>\n<li>リスト</li>\n</ul>\n"},
	// リンクのテスト
	{"[localhost](localhost:8000)", `<p><a href="localhost:8000">localhost</a></p>` + "\n"},
	// 独自記法
	{"``` bash\n" + "echo Hello" + "\n" + "```", `<pre style="background-color:#fff"><span style="color:#0086b3">echo</span> Hello` + "\n" + "</pre>"},
}

func Test_Render(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			html, err := Render(context.Background(), tt.in)
			assert.NoError(t, err)
			if html != tt.out {
				t.Errorf("got %q, want %q", html, tt.out)
			}
		})
	}
}
