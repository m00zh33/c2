package api

import (
	"bytes"
	"context"
	"errors"
	"net"

	"gitlab.com/teserakt/c2/internal/config"
	"gitlab.com/teserakt/c2/internal/services"
	e4 "gitlab.com/teserakt/e4common"

	"github.com/go-kit/kit/log"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GRPCServer defines available endpoints on a GRPC server
type GRPCServer interface {
	e4.C2Server
	ListenAndServe(ctx context.Context) error
}

type grpcServer struct {
	logger    log.Logger
	cfg       config.ServerCfg
	e4Service services.E4
}

var _ GRPCServer = &grpcServer{}

// NewGRPCServer creates a new server over GRPC
func NewGRPCServer(scfg config.ServerCfg, e4Service services.E4, logger log.Logger) GRPCServer {
	return &grpcServer{
		cfg:       scfg,
		logger:    logger,
		e4Service: e4Service,
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
	e4.RegisterC2Server(srv, s)

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

// C2Command processes a command received over gRPC by the CLI tool.
func (s *grpcServer) C2Command(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	ctx, span := trace.StartSpan(ctx, e4.C2Request_Command_name[int32(in.Command)])
	defer span.End()

	s.logger.Log("msg", "received gRPC request", "request", e4.C2Request_Command_name[int32(in.Command)])

	switch in.Command {
	case e4.C2Request_NEW_CLIENT:
		return s.gRPCnewClient(ctx, in)
	case e4.C2Request_REMOVE_CLIENT:
		return s.gRPCremoveClient(ctx, in)
	case e4.C2Request_NEW_TOPIC_CLIENT:
		return s.gRPCnewTopicClient(ctx, in)
	case e4.C2Request_REMOVE_TOPIC_CLIENT:
		return s.gRPCremoveTopicClient(ctx, in)
	case e4.C2Request_RESET_CLIENT:
		return s.gRPCresetClient(ctx, in)
	case e4.C2Request_NEW_TOPIC:
		return s.gRPCnewTopic(ctx, in)
	case e4.C2Request_REMOVE_TOPIC:
		return s.gRPCremoveTopic(ctx, in)
	case e4.C2Request_NEW_CLIENT_KEY:
		return s.gRPCnewClientKey(ctx, in)
	case e4.C2Request_SEND_MESSAGE:
		return s.gRPCsendMessage(ctx, in)
	case e4.C2Request_GET_CLIENTS:
		return s.gRPCgetClients(ctx, in)
	case e4.C2Request_GET_TOPICS:
		return s.gRPCgetTopics(ctx, in)
	case e4.C2Request_GET_CLIENT_TOPIC_COUNT:
		return s.gRPCgetClientTopicCount(ctx, in)
	case e4.C2Request_GET_CLIENT_TOPICS:
		return s.gRPCgetClientTopics(ctx, in)
	case e4.C2Request_GET_TOPIC_CLIENT_COUNT:
		return s.gRPCgetTopicClientCount(ctx, in)
	case e4.C2Request_GET_TOPIC_CLIENTS:
		return s.gRPCgetTopicClients(ctx, in)
	}
	return &e4.C2Response{Success: false, Err: "unknown command"}, nil
}

func (s *grpcServer) gRPCnewClient(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, true, false, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.NewClient(ctx, in.Name, nil, in.Key)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCremoveClient(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, false, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.RemoveClientByName(ctx, in.Name)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCnewTopicClient(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.NewTopicClient(ctx, in.Name, in.Id, in.Topic)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCremoveTopicClient(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.RemoveTopicClientByName(ctx, in.Name, in.Topic)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCresetClient(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, false, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}
	err = s.e4Service.ResetClientByName(ctx, in.Name)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCnewTopic(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, false, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.NewTopic(ctx, in.Topic)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCremoveTopic(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, false, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.RemoveTopic(ctx, in.Topic)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCnewClientKey(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, false, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.NewClientKey(ctx, in.Name, in.Id)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

func (s *grpcServer) gRPCgetClients(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	names, err := s.e4Service.GetAllClientsAsNames(ctx)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: "", Names: names}, nil
}

func (s *grpcServer) gRPCgetTopics(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	topics, err := s.e4Service.GetAllTopics(ctx)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: "", Topics: topics}, nil
}

func (s *grpcServer) gRPCgetClientTopicCount(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, false, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	count, err := s.e4Service.CountTopicsForClientByName(ctx, in.Name)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: "", Count: uint64(count)}, nil
}

func (s *grpcServer) gRPCgetClientTopics(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, true, false, false, false, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	topics, err := s.e4Service.GetTopicsForClientByID(ctx, in.Id, int(in.Offset), int(in.Count))
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: "", Topics: topics}, nil
}

func (s *grpcServer) gRPCgetTopicClientCount(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, false, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	count, err := s.e4Service.CountClientsForTopic(ctx, in.Topic)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: "", Count: uint64(count)}, nil
}

func (s *grpcServer) gRPCgetTopicClients(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, false, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	clients, err := s.e4Service.GetClientsByNameForTopic(ctx, in.Topic, int(in.Offset), int(in.Count))
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: "", Ids: clients}, nil
}

func (s *grpcServer) gRPCsendMessage(ctx context.Context, in *e4.C2Request) (*e4.C2Response, error) {
	err := checkRequest(ctx, in, false, false, false, true, false)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	err = s.e4Service.SendMessage(ctx, in.Topic, in.Msg)
	if err != nil {
		return &e4.C2Response{Success: false, Err: err.Error()}, nil
	}

	return &e4.C2Response{Success: true, Err: ""}, nil
}

// helper to check inputs' sanity
func checkRequest(ctx context.Context, in *e4.C2Request, needName, needID, needKey, needTopic, needOffsetCount bool) error {
	ctx, span := trace.StartSpan(ctx, "checkRequest")
	defer span.End()

	if needName {
		if err := e4.IsValidName(in.Name); err != nil {
			return err
		}
	} else {
		if in.Name != "" {
			return errors.New("unexpected name")
		}
	}
	if needID {
		if err := e4.IsValidID(in.Id); err != nil {
			return err
		}
		if needName {
			if !bytes.Equal(in.Id, e4.HashIDAlias(in.Name)) {
				return errors.New("name and e4id are not consistent")
			}
		}
	} else {
		// we might send the ID along with the name; if ID is optional,
		// don't worry if it is sent.
		if needName == false && in.Id != nil {
			return errors.New("unexpected id")
		}
	}
	if needKey {
		if err := e4.IsValidKey(in.Key); err != nil {
			return err
		}
	} else {
		if in.Key != nil {
			return errors.New("unexpected key")
		}
	}
	if needTopic {
		if err := e4.IsValidTopic(in.Topic); err != nil {
			return err
		}
	} else {
		if in.Topic != "" {
			return errors.New("unexpected topic")
		}
	}
	if needOffsetCount {
		if in.Count == 0 {
			return errors.New("No data to return with zero count")
		}
	}

	return nil
}
