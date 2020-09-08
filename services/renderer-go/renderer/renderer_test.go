package renderer

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Render(t *testing.T) {
	// 記法のテスト
	headlineSrc := `# 見出し`
	headlineHtml, headlineErr := Render(context.Background(), headlineSrc)
	assert.NoError(t, headlineErr)
	assert.Equal(t, "<h1>見出し</h1>\n", headlineHtml)

	fmt.Println("------------------------------")

	// リストのテスト
	listSrc := "- リスト"
	listHtml, listErr := Render(context.Background(), listSrc)
	assert.NoError(t, listErr)
	assert.Equal(t, "<ul>\n<li>リスト</li>\n</ul>\n", listHtml)

	fmt.Println("------------------------------")

	// リンクのテスト
	linkSrc := "[localhost](localhost:8000)"
	linkHtml, linkErr := Render(context.Background(), linkSrc)
	assert.NoError(t, linkErr)
	assert.Equal(t, `<p><a href="localhost:8000">localhost</a></p>`+"\n", linkHtml)

	fmt.Println("------------------------------")

	// 独自記法
	originalSrc := "``` bash\n" + "echo Hello" + "\n" + "```"
	originalHtml, originalErr := Render(context.Background(), originalSrc)
	ans := `<pre style="background-color:#fff"><span style="color:#0086b3">echo</span> Hello` + "\n" + "</pre>"
	assert.NoError(t, originalErr)
	assert.Equal(t, ans, originalHtml)
}
