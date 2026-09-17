//revive:disable:package-comments
package httpclient

import (
	"time"
)

type configuration struct {
	SendPingTimeout time.Duration `env:"HTTP2_SEND_PING_TIMEOUT" envDefault:"2m"`
	PingTimeout     time.Duration `env:"HTTP2_PING_TIMEOUT" envDefault:"20s"`
}
