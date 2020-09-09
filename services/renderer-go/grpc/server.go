package grpc

import (
	"context"

	pb "github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/dilmnqvovpnmlib/Hatena-Intern-2020/services/renderer-go/app"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
// Server は Web サーバーを表す構造体
type Server struct {
	pb.UnimplementedRendererServer
	healthpb.UnimplementedHealthServer
	app *app.App
}

// NewServer は gRPC サーバーを作成する
func NewServer(app *app.App) (*Server, error) {
	return &Server{app: app}, nil
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb.RenderRequest) (*pb.RenderReply, error) {
	html, err := s.app.Render(ctx, in.Src)
	if err != nil {
		return nil, err
	}
	return &pb.RenderReply{Html: html}, nil
}
