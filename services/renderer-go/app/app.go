package app

import (
	"context"

	pb_fetcher "github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/renderer"
)

// App はアプリケーションを表す
type App struct {
	fetcherClient        pb_fetcher.FetcherClient
}

// NewApp は App を作成する
func NewApp(
	fetcherClient pb_fetcher.FetcherClient,
) *App {
	return &App{fetcherClient}
}

// 
func (a *App) Render(ctx context.Context, src string) (string, error) {
	html, err := renderer.Render(ctx, src, a.fetcherClient)
	if err != nil {
		return "", err
	}
	return html, nil
}
