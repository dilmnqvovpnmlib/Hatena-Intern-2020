package fetcher

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	pb_fetcher "github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

var markdown = goldmark.New(
	goldmark.WithParserOptions(
		parser.WithASTTransformers(
			util.Prioritized(&autoTitleLinker{}, 999),
		),
	),
)

func main() {
	src := []byte("# link samples\n" +
		"[normal link](https://example.com)\n" +
		"[](https://example.com)\n")
	var buf bytes.Buffer
	if err := markdown.Convert(src, &buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf.String())
}

type autoTitleLinker struct {
	fetcherCli pb_fetcher.FetcherClient
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			node.AppendChild(node, ast.NewString([]byte(fetchTitle(l.fetcherCli, string(node.Destination)))))
		}
		return ast.WalkContinue, nil
	})
}

// grpc で処理する部分に相当してて、 /app/app.go の Render のこと
func fetchTitle(fetcherCli pb_fetcher.FetcherClient, url string) string {
	reply, err := fetcherCli.Render(ctx, &pb_fetcher.FetcherRequest{Url: url})
	if err != nil {
		return nil
	}
	return reply.Title
}
