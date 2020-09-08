package renderer
import (
    "bytes"
    "context"
    "html/template"
    "regexp"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
)

var urlRE = regexp.MustCompile(`https?://[^\s]+`)
var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))

func ParseHtml(src string) (string, error) {
    markdown := goldmark.New(
        goldmark.WithExtensions(
            highlighting.Highlighting,
        ),
    )
    var buf bytes.Buffer
    if err := markdown.Convert([]byte(src), &buf); err != nil {
        panic(err)
    }
    html := buf.String()
    return html, nil
}

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
    html, err := ParseHtml(src) // buf.String()
    return html, err
}
