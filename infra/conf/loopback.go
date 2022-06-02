package conf

import (
	"github.com/Yuzuki616/xray-core/proxy/loopback"
	"github.com/golang/protobuf/proto"
)

type LoopbackConfig struct {
	InboundTag string `json:"inboundTag"`
}

func (l LoopbackConfig) Build() (proto.Message, error) {
	return &loopback.Config{InboundTag: l.InboundTag}, nil
}
