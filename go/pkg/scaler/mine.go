package scaler

import (
	"context"
	"github.com/AliyunContainerService/scaler/go/pkg/config"
	model2 "github.com/AliyunContainerService/scaler/go/pkg/model"

	pb "github.com/AliyunContainerService/scaler/proto"
	
)

type Mine struct {
}

func NewMine(metaData *model2.Meta, config *config.Config) Mine {
	return Mine{}
}

func (m *Mine) Assign(ctx context.Context, request *pb.AssignRequest) (*pb.AssignReply, error) {
}

func (m *Mine) Idle(ctx context.Context, request *pb.IdleRequest) (*pb.IdleReply, error) {
}

func (m *Mine) Stats() Stats {
	
}

