package fetcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Fetcher(t *testing.T) {
	url := "https://google.com/"
	title := GetTitle(url)
	assert.Equal(t, title, "Google")
}
