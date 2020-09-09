package grpc

import (
	"context"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/hatena/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
	"github.com/hatena/Hatena-Intern-2020/services/fetcher-go/fetcher"
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
	title := fetcher.GetTitle(in.Url)
	return &pb.FetcherReply{Title: title}, nil
}
