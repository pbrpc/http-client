//revive:disable:package-comments
package httpclient

import (
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	transport "github.com/pbrpc/http-transport"
)

type configuration struct {
	SendPingTimeout time.Duration `env:"HTTP2_SEND_PING_TIMEOUT" envDefault:"2m"`
	PingTimeout     time.Duration `env:"HTTP2_PING_TIMEOUT" envDefault:"20s"`
}

// FromEnv creates an HTTP client with OpenTelemetry trace propagation. A nil
// base selects the standard cleartext HTTP/2 transport configured by
// HTTP2_SEND_PING_TIMEOUT and HTTP2_PING_TIMEOUT.
func FromEnv(base http.RoundTripper) (*http.Client, error) {
	if base == nil {
		configured, err := configurationFromEnv()
		if err != nil {
			return nil, err
		}

		base = transport.New(&http.HTTP2Config{
			SendPingTimeout: configured.SendPingTimeout,
			PingTimeout:     configured.PingTimeout,
		})
	}

	return &http.Client{Transport: otelhttp.NewTransport(base)}, nil
}

func configurationFromEnv() (configuration, error) {
	return env.ParseAs[configuration]()
}
