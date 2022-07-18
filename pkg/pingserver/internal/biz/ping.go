package biz

import (
	"context"
	"time"

	"github.com/moyrne/delay-record/pkg/pingserver/params"
)

func Ping(_ context.Context) (*params.PingResponse, error) {
	return &params.PingResponse{Time: time.Now().Unix()}, nil
}
