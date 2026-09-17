# http-client

`http-client` builds instrumented HTTP clients for cleartext HTTP/2
communication on an internal service mesh.

## Installation

```bash
go get github.com/pbrpc/http-client
```

## Client

`FromEnv` wraps a supplied `http.RoundTripper` with OpenTelemetry trace
propagation. A nil transport selects the standard cleartext HTTP/2 transport
configured through these environment variables:

| Variable                  | Default | Meaning                                  |
| ------------------------- | ------- | ---------------------------------------- |
| `HTTP2_SEND_PING_TIMEOUT` | `2m`    | Quiet time before sending an HTTP/2 ping |
| `HTTP2_PING_TIMEOUT`      | `20s`   | Wait for a ping response before closing  |

```go
client, err := httpclient.FromEnv(nil)
if err != nil {
	return err
}
```

An injected transport is placed beneath the tracing transport:

```go
client, err := httpclient.FromEnv(discoveryTransport)
if err != nil {
	return err
}
```
