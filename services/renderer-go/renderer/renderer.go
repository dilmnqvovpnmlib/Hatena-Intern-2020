package renderer
import (
    "fmt"
    "bytes"
    "context"
    "html/template"
    "regexp"
	"github.com/yuin/goldmark"
    // "github.com/yuin/goldmark/extension"
    // "github.com/yuin/goldmark/parser"
	// "github.com/yuin/goldmark/renderer/html"
	// "github.com/yuin/goldmark-highlighting"
)

var urlRE = regexp.MustCompile(`https?://[^\s]+`)
var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))

func ParseHtml(src string) (string, error) {
	// md := goldmark.New(
	// 	goldmark.WithExtensions(extension.GFM),
	// 	goldmark.WithParserOptions(
	// 		parser.WithAutoHeadingID(),
	// 	),
	// 	goldmark.WithRendererOptions(
	// 		html.WithHardWraps(),
	// 		html.WithXHTML(),
	// 	),
	// )
    fmt.Println("------------------markdown-----------------------------------")
	// markdown := goldmark.New(
	// 	goldmark.WithExtensions(
	// 		Highlighting,
	// 	),
	// )
    var buf bytes.Buffer
    if err := goldmark.Convert([]byte(src), &buf); err != nil {
        panic(err)
    }
    html := buf.String()
    return html, nil
}

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
    // TODO: これはサンプル実装 (URL の自動リンク) です
    // もとからあったコード
    // html := urlRE.ReplaceAllStringFunc(src, func(url string) string {
    //  var w bytes.Buffer
    //  err := linkTmpl.ExecuteTemplate(&w, "link", url)
    //  if err != nil {
    //      return url
    //  }
    //  return w.String()
    // })
    fmt.Println("------------------Hello0000-----------------------------------")
    // var buf bytes.Buffer
    // if err := goldmark.Convert([]byte(src), &buf); err != nil {
    //  panic(err)
    // }
    html, err := ParseHtml(src) // buf.String()
    return html, err
}
