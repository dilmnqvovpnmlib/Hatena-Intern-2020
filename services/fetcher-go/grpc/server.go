package grpc

import (
	"context"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hatena/Hatena-Intern-2020/services/fetcher-go/fetcher"
	pb "github.com/hatena/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
)

type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// 受け取った url からタイトルに変換する
func (s *Server) Render(ctx context.Context, in *pb.FetcherRequest) (*pb.FetcherReply, error) {
	title, err := fetcher.GetTitle(in.Url)
	if err != nil {
		return &pb.FetcherReply{}, err
	}
	return &pb.FetcherReply{Title: title}, nil
}
