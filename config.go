//revive:disable:package-comments
package httpclient

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type configuration struct {
	SendPingTimeout time.Duration `env:"HTTP2_SEND_PING_TIMEOUT" envDefault:"2m"`
	PingTimeout     time.Duration `env:"HTTP2_PING_TIMEOUT" envDefault:"20s"`
}

func configurationFromEnv() (configuration, error) {
	return env.ParseAs[configuration]()
}
