package use_connect

import (
	pb "github.com/eonias189/calculationService/backend/internal/proto"
)

type Connector interface {
	Connect(pb.Orchestrator_ConnectServer) error
}
