//go:build !with_utls

package tls

import (
	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func newUTLSClient(router adapter.Router, serverAddress string, options option.OutboundTLSOptions) (Config, error) {
	return nil, E.New(`uTLS is not included in this build, rebuild with -tags with_utls`)
}
