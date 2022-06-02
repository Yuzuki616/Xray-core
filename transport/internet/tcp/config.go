package tcp

import (
	"github.com/Yuzuki616/xray-core/common"
	"github.com/Yuzuki616/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
