//revive:disable:package-comments
package httpclient

import (
	"net/http"
	"testing"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/pbrpc/connect-testing/mocks/roundtripper"
)

func TestFromEnv(t *testing.T) {
	t.Run("resolves the correct configuration", func(t *testing.T) {
		t.Setenv("HTTP2_SEND_PING_TIMEOUT", "3m")
		t.Setenv("HTTP2_PING_TIMEOUT", "30s")

		configured, err := configurationFromEnv()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if configured.SendPingTimeout != 3*time.Minute {
			t.Errorf("send ping timeout = %v, want 3m", configured.SendPingTimeout)
		}
		if configured.PingTimeout != 30*time.Second {
			t.Errorf("ping timeout = %v, want 30s", configured.PingTimeout)
		}
	})

	t.Run("builds the configured standard client", func(t *testing.T) {
		client, err := FromEnv(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, ok := client.Transport.(*otelhttp.Transport); !ok {
			t.Fatalf("transport = %T, want the tracing transport", client.Transport)
		}
	})

	t.Run("returns invalid configuration", func(t *testing.T) {
		t.Setenv("HTTP2_SEND_PING_TIMEOUT", "not-a-duration")

		if _, err := FromEnv(nil); err == nil {
			t.Fatal("error = nil, want the configuration error")
		}
	})

	t.Run("wraps an injected transport", func(t *testing.T) {
		t.Setenv("HTTP2_SEND_PING_TIMEOUT", "not-a-duration")

		base := roundtripper.Record(roundtripper.Respond(http.StatusNoContent, nil, ""))
		client, err := FromEnv(base)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		request, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://service/", nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, err = client.Do(request); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		sent := base.Sent()
		if len(sent) != 1 || sent[0].Request.URL.Host != "service" {
			t.Errorf("sent = %v, want one request to service", sent)
		}
	})
}
