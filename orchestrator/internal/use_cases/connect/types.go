package use_connect

import (
	"errors"

	pb "github.com/eonias189/calculationService/orchestrator/internal/proto"
)

var (
	ErrInChanClosed                 = errors.New("input chan closed")
	ErrMetadataInvalidOrNotProvided = errors.New("metadata is invalid or not provided")
)

type Connector interface {
	Connect(pb.Orchestrator_ConnectServer) error
}
