package renderer

import (
	"context"
	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"
	"testing"

	pb "github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

type inputType struct {
	html  string
	title string
}

var flagtests = []struct {
	in  inputType
	out string
}{
	// 記法のテスト
	{
		inputType{html: `# 見出し`, title: ""},
		"<h1>見出し</h1>\n",
	},
	// リストのテスト
	{
		inputType{html: "- リスト", title: ""},
		"<ul>\n<li>リスト</li>\n</ul>\n",
	},
	// リンクのテスト
	{
		inputType{html: "[localhost](localhost:8000)", title: ""},
		`<p><a href="localhost:8000">localhost</a></p>` + "\n",
	},
	// 独自記法
	{
		inputType{html: "``` bash\n" + "echo Hello" + "\n" + "```", title: ""},
		`<pre style="background-color:#fff"><span style="color:#0086b3">echo</span> Hello` + "\n" + "</pre>",
	},
	// タイトルの自動取得
	{
		inputType{html: "[](https://example.com)", title: "Example Domain"},
		`<p><a href="https://example.com">Example Domain</a></p>` + "\n",
	},
}

type fakeFetcherClient struct {
	pb.FetcherClient
	FakeRender func(ctx context.Context, in *pb.FetcherRequest, opts ...grpc.CallOption) (*pb.FetcherReply, error)
}

func (c *fakeFetcherClient) Render(ctx context.Context, in *pb.FetcherRequest, opts ...grpc.CallOption) (*pb.FetcherReply, error) {
	return c.FakeRender(ctx, in, opts...)
}

func Test_Render(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in.html, func(t *testing.T) {
			fakeFetcherClient := &fakeFetcherClient{
				FakeRender: func(ctx context.Context, in *pb.FetcherRequest, opts ...grpc.CallOption) (*pb.FetcherReply, error) {
					return &pb.FetcherReply{Title: tt.in.title}, nil
				},
			}
			html, err := Render(context.Background(), tt.in.html, fakeFetcherClient)
			assert.NoError(t, err)
			if html != tt.out {
				t.Errorf("got %q, want %q", html, tt.out)
			}
		})
	}
}
