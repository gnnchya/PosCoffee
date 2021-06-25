package implement

import (
	"google.golang.org/grpc"
)

type implementation struct {
	conn      *grpc.ClientConn
}
//
//func New(grpcRepo util.RepositoryGRPC) (service grpcClient.Service) {
//	conn, err := grpcRepo.NewClient()
//	if err != nil {
//		return nil
//	}
//	impl := &implementation{
//		conn:      conn,
//	}
//	return impl
//
//}
