package api

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/codes"

	"gitlab.com/teserakt/c2/internal/config"
	"gitlab.com/teserakt/c2/internal/events"
	"gitlab.com/teserakt/c2/internal/services"
	"gitlab.com/teserakt/c2/pkg/pb"

	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/ptypes"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Request parameters validation errors
var (
	ErrClientRequired     = grpc.Errorf(codes.InvalidArgument, "a client is required.")
	ErrClientNameRequired = grpc.Errorf(codes.InvalidArgument, "a client name is required.")
)

// GRPCServer defines available endpoints on a GRPC server
type GRPCServer interface {
	pb.C2Server
	ListenAndServe(ctx context.Context) error
}

type grpcServer struct {
	logger          log.Logger
	cfg             config.ServerCfg
	e4Service       services.E4
	eventDispatcher events.Dispatcher
}

var _ GRPCServer = &grpcServer{}

// NewGRPCServer creates a new server over GRPC
func NewGRPCServer(scfg config.ServerCfg, e4Service services.E4, eventDispatcher events.Dispatcher, logger log.Logger) GRPCServer {
	return &grpcServer{
		cfg:             scfg,
		logger:          logger,
		e4Service:       e4Service,
		eventDispatcher: eventDispatcher,
	}
}

func (s *grpcServer) ListenAndServe(ctx context.Context) error {
	var lc net.ListenConfig
	lis, err := lc.Listen(ctx, "tcp", s.cfg.Addr)
	if err != nil {
		s.logger.Log("msg", "failed to listen", "error", err)

		return err
	}

	creds, err := credentials.NewServerTLSFromFile(s.cfg.Cert, s.cfg.Key)
	if err != nil {
		s.logger.Log("msg", "failed to get credentials", "cert", s.cfg.Cert, "key", s.cfg.Key, "error", err)
		return err
	}

	s.logger.Log("msg", "using TLS for gRPC", "cert", s.cfg.Cert, "key", s.cfg.Key)

	if err = view.Register(ocgrpc.DefaultServerViews...); err != nil {
		s.logger.Log("msg", "failed to register ocgrpc server views", "error", err)

		return err
	}

	srv := grpc.NewServer(grpc.Creds(creds), grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	pb.RegisterC2Server(srv, s)

	s.logger.Log("msg", "Starting api grpc server", "addr", s.cfg.Addr)

	errc := make(chan error)
	go func() {
		errc <- srv.Serve(lis)
	}()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *grpcServer) NewClient(ctx context.Context, req *pb.NewClientRequest) (*pb.NewClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	if len(req.Client.Name) == 0 {
		return nil, ErrClientNameRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	err = s.e4Service.NewClient(ctx, req.Client.Name, id, req.Key)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.NewClientResponse{}, nil
}

func (s *grpcServer) RemoveClient(ctx context.Context, req *pb.RemoveClientRequest) (*pb.RemoveClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	err = s.e4Service.RemoveClient(ctx, id)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.RemoveClientResponse{}, nil
}

func (s *grpcServer) ResetClient(ctx context.Context, req *pb.ResetClientRequest) (*pb.ResetClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	err = s.e4Service.ResetClient(ctx, id)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.ResetClientResponse{}, nil
}

func (s *grpcServer) NewClientKey(ctx context.Context, req *pb.NewClientKeyRequest) (*pb.NewClientKeyResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	err = s.e4Service.NewClientKey(ctx, id)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.NewClientKeyResponse{}, nil
}

func (s *grpcServer) NewTopic(ctx context.Context, req *pb.NewTopicRequest) (*pb.NewTopicResponse, error) {
	err := s.e4Service.NewTopic(ctx, req.Topic)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.NewTopicResponse{}, nil
}

func (s *grpcServer) RemoveTopic(ctx context.Context, req *pb.RemoveTopicRequest) (*pb.RemoveTopicResponse, error) {
	err := s.e4Service.RemoveTopic(ctx, req.Topic)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.RemoveTopicResponse{}, nil
}

func (s *grpcServer) NewTopicClient(ctx context.Context, req *pb.NewTopicClientRequest) (*pb.NewTopicClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	err = s.e4Service.NewTopicClient(ctx, id, req.Topic)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.NewTopicClientResponse{}, nil
}

func (s *grpcServer) RemoveTopicClient(ctx context.Context, req *pb.RemoveTopicClientRequest) (*pb.RemoveTopicClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	err = s.e4Service.RemoveTopicClient(ctx, id, req.Topic)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.RemoveTopicClientResponse{}, nil
}

func (s *grpcServer) CountTopicsForClient(ctx context.Context, req *pb.CountTopicsForClientRequest) (*pb.CountTopicsForClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	count, err := s.e4Service.CountTopicsForClient(ctx, id)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CountTopicsForClientResponse{Count: int64(count)}, nil
}

func (s *grpcServer) GetTopicsForClient(ctx context.Context, req *pb.GetTopicsForClientRequest) (*pb.GetTopicsForClientResponse, error) {
	if req.Client == nil {
		return nil, ErrClientRequired
	}

	id, err := validateE4NameOrIDPair(req.Client.Name, nil)
	if err != nil {
		return nil, grpcError(err)
	}

	topics, err := s.e4Service.GetTopicsRangeByClient(ctx, id, int(req.Offset), int(req.Count))
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.GetTopicsForClientResponse{Topics: topics}, nil
}

func (s *grpcServer) CountClientsForTopic(ctx context.Context, req *pb.CountClientsForTopicRequest) (*pb.CountClientsForTopicResponse, error) {
	count, err := s.e4Service.CountClientsForTopic(ctx, req.Topic)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CountClientsForTopicResponse{Count: int64(count)}, nil
}

func (s *grpcServer) GetClientsForTopic(ctx context.Context, req *pb.GetClientsForTopicRequest) (*pb.GetClientsForTopicResponse, error) {
	clients, err := s.e4Service.GetClientsRangeByTopic(ctx, req.Topic, int(req.Offset), int(req.Count))
	if err != nil {
		return nil, grpcError(err)
	}

	pbClients := make([]*pb.Client, 0, len(clients))
	for _, client := range clients {
		pbClients = append(pbClients, &pb.Client{Name: client.Name})
	}

	return &pb.GetClientsForTopicResponse{Clients: pbClients}, nil
}

func (s *grpcServer) GetClients(ctx context.Context, req *pb.GetClientsRequest) (*pb.GetClientsResponse, error) {
	clients, err := s.e4Service.GetClientsRange(ctx, int(req.Offset), int(req.Count))
	if err != nil {
		return nil, grpcError(err)
	}

	pbClients := make([]*pb.Client, 0, len(clients))
	for _, client := range clients {
		pbClients = append(pbClients, &pb.Client{Name: client.Name})
	}

	return &pb.GetClientsResponse{Clients: pbClients}, nil
}
func (s *grpcServer) GetTopics(ctx context.Context, req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	topics, err := s.e4Service.GetTopicsRange(ctx, int(req.Offset), int(req.Count))
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.GetTopicsResponse{Topics: topics}, nil
}

func (s *grpcServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	err := s.e4Service.SendMessage(ctx, req.Topic, req.Message)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.SendMessageResponse{}, nil
}

func (s *grpcServer) CountClients(ctx context.Context, req *pb.CountClientsRequest) (*pb.CountClientsResponse, error) {
	count, err := s.e4Service.CountClients(ctx)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CountClientsResponse{Count: int64(count)}, nil
}

func (s *grpcServer) CountTopics(ctx context.Context, req *pb.CountTopicsRequest) (*pb.CountTopicsResponse, error) {
	count, err := s.e4Service.CountTopics(ctx)
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.CountTopicsResponse{Count: int64(count)}, nil
}

func (s *grpcServer) SubscribeToEventStream(req *pb.SubscribeToEventStreamRequest, srv pb.C2_SubscribeToEventStreamServer) error {
	listener := events.NewListener(s.eventDispatcher)
	defer listener.Close()

	logger := log.With(s.logger, "listener", fmt.Sprintf("%p", listener))

	logger.Log("msg", "Started new event stream")

	for {
		ctx := srv.Context()
		select {
		case evt := <-listener.C():
			ts, err := ptypes.TimestampProto(evt.Timestamp)
			if err != nil {
				return fmt.Errorf("failed to convert time.Time to proto.Timestamp: %v", err)
			}

			pbEvt := &pb.Event{
				Type:      pb.EventType(evt.Type),
				Source:    evt.Source,
				Target:    evt.Target,
				Timestamp: ts,
			}

			if err := srv.Send(pbEvt); err != nil {
				logger.Log("msg", "failed to send event", "error", err)
				return err
			}

			logger.Log("msg", "Successfully sent event", "eventType", pbEvt.Type.String())
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// validateE4NamedOrIDPair wrap around services.ValidateE4NameOrIDPair but will
// conver the error to a suitable GRPC error
func validateE4NameOrIDPair(name string, id []byte) ([]byte, error) {
	id, err := services.ValidateE4NameOrIDPair(name, id)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return id, nil
}

// grpcError will convert the given error to a GRPC error with appropriate code
func grpcError(err error) error {
	var code codes.Code
	switch err.(type) {
	case services.ErrClientNotFound, services.ErrTopicNotFound:
		code = codes.NotFound
	case services.ErrValidation:
		code = codes.InvalidArgument
	default:
		code = codes.Internal
	}

	return grpc.Errorf(code, "%s", err.Error())
}
