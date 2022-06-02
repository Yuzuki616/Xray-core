//go:build !linux && !freebsd
// +build !linux,!freebsd

package tcp

import (
	"github.com/Yuzuki616/xray-core/common/net"
	"github.com/Yuzuki616/xray-core/transport/internet/stat"
)

func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
