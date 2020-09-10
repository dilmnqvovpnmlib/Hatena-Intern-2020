package fetcher

import (
	"context"
	"log"

	pb_fetcher "github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
)

func FetchTitle(ctx context.Context, url string, fetcherClient pb_fetcher.FetcherClient) string {
	reply, err := fetcherClient.Render(ctx, &pb_fetcher.FetcherRequest{Url: url})
	if err != nil {
		log.Fatal(err)
	}
	return reply.Title
}
