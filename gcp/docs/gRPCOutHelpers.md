# gRPCOutHelpers.go

## File Overview
- Path: `gcp/gRPCOutHelpers.go`
- Package: `gcp`
- Functions/Methods: `2`
- Imports: `13`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GRPCDialer`

## Imports
- `FenixTesterGui/common_code`
- `bufio`
- `context`
- `crypto/tls`
- `fmt`
- `google.golang.org/grpc`
- `google.golang.org/grpc/credentials`
- `log`
- `net`
- `net/http`
- `net/url`
- `strconv`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### GRPCDialer
- Signature: `func GRPCDialer(bearerToken string) (*grpc.ClientConn, error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Internal calls: `proxyDialer`
- Selector calls: `credentials.NewTLS`, `grpc.Dial`, `grpc.WithContextDialer`, `grpc.WithTransportCredentials`, `log.Fatalf`, `strconv.Itoa`

### proxyDialer
- Signature: `func proxyDialer(proxyURL string, bearerToken string) (func(ctx context.Context, addr string) (net.Conn, error))`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bufio.NewReader`, `connectReq.Write`, `dialer.DialContext`, `fmt.Errorf`, `http.ReadResponse`, `proxyConn.Close`, `url.Parse`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
