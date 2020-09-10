package grpc

import (
	"context"
	"testing"

	pb "github.com/hatena/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Fetcher(t *testing.T) {
	s := NewServer()
	url := "https://google.com/"
	reply, err := s.Render(context.Background(), &pb.FetcherRequest{Url: url})
	assert.NoError(t, err)
	assert.Equal(t, `Google`, reply.Title)
}
