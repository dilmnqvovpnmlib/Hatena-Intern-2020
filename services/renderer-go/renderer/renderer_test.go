package renderer

import (
    "fmt"
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_Render(t *testing.T) {
    // src := `foo https://google.com/ bar`
    // html, err := Render(context.Background(), src)
    // assert.NoError(t, err)
	// assert.Equal(t, `foo <a href="https://google.com/">https://google.com/</a> bar`, html)
	fmt.Println("---------Test----------")
    // 記法のテスト
    headlineSrc := `# 見出し`
    headlineHtml, headlineErr := Render(context.Background(), headlineSrc)
    assert.NoError(t, headlineErr)
	assert.Equal(t, "<h1>見出し</h1>\n", headlineHtml)
    // リストのテスト
    listSrc := "- リスト"
    listHtml, listErr := Render(context.Background(), listSrc)
    assert.NoError(t, listErr)
	assert.Equal(t, "<ul>\n<li>リスト</li>\n</ul>\n", listHtml)
    // リンクのテスト
    linkSrc := "[localhost](localhost:8000)"
    linkHtml, linkErr := Render(context.Background(), linkSrc)
    assert.NoError(t, linkErr)
	assert.Equal(t, `<p><a href="localhost:8000">localhost</a></p>` + "\n", linkHtml)
	// 独自記法
	fmt.Println("---------Test----------")
}
