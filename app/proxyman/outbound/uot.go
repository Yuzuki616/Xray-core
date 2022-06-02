//go:build go1.18

package outbound

import (
	"context"
	"os"

	"github.com/Yuzuki616/xray-core/common/net"
	"github.com/Yuzuki616/xray-core/transport/internet"
	"github.com/Yuzuki616/xray-core/transport/internet/stat"
	"github.com/sagernet/sing/common/uot"
)

func (h *Handler) getUoTConnection(ctx context.Context, dest net.Destination) (stat.Connection, error) {
	if !dest.Address.Family().IsDomain() || dest.Address.Domain() != uot.UOTMagicAddress {
		return nil, os.ErrInvalid
	}
	packetConn, err := internet.ListenSystemPacket(ctx, &net.UDPAddr{IP: net.AnyIP.IP(), Port: 0}, h.streamSettings.SocketSettings)
	if err != nil {
		return nil, newError("unable to listen socket").Base(err)
	}
	conn := uot.NewServerConn(packetConn)
	return h.getStatCouterConnection(conn), nil
}
