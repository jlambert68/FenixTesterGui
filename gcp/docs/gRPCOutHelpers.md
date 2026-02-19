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
- Selector calls: `strconv.Itoa`, `credentials.NewTLS`, `grpc.Dial`, `grpc.WithTransportCredentials`, `grpc.WithContextDialer`, `log.Fatalf`

### proxyDialer
- Signature: `func proxyDialer(proxyURL string, bearerToken string) (func(ctx context.Context, addr string) (net.Conn, error))`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `url.Parse`, `fmt.Errorf`, `dialer.DialContext`, `connectReq.Write`, `proxyConn.Close`, `http.ReadResponse`, `bufio.NewReader`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
