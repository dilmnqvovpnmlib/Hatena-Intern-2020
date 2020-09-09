package renderer

import (
	"bytes"
	"context"
	"log"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"

	pb_fetcher "github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

type autoTitleLinker struct {
	ctx context.Context
	fetcherClient pb_fetcher.FetcherClient
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			node.AppendChild(node, ast.NewString([]byte(fetchTitle(l.ctx, string(node.Destination), l.fetcherClient))))
		}
		return ast.WalkContinue, nil
	})
}

func fetchTitle(ctx context.Context, url string, fetcherClient pb_fetcher.FetcherClient) string {
	reply, err := fetcherClient.Render(ctx, &pb_fetcher.FetcherRequest{Url: url})
	if err != nil {
		log.Fatal(err)
	}
	return reply.Title
}

func ParseHtml(ctx context.Context, src string, fetcherClient pb_fetcher.FetcherClient) (string, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.Highlighting,
		),
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&autoTitleLinker{ctx: ctx, fetcherClient: fetcherClient}, 999),
			),
		),
	)
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(src), &buf); err != nil {
		return "", err
	}
	html := buf.String()
	return html, nil
}

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string, fetcherClient pb_fetcher.FetcherClient) (string, error) {
	html, err := ParseHtml(ctx, src, fetcherClient)
	return html, err
}
