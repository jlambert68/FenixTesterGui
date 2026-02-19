# dial.go

## File Overview
- Path: `grpcurl/dial.go`
- Package: `grpcurl`
- Functions/Methods: `5`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `BlockingDial`
- `ClientHandshake`
- `ClientTLSConfig`
- `ClientTransportCredentials`
- `ServerTransportCredentials`

## Imports
- `context`
- `crypto/tls`
- `crypto/x509`
- `errors`
- `fmt`
- `google.golang.org/grpc`
- `google.golang.org/grpc/credentials`
- `google.golang.org/grpc/credentials/insecure`
- `io/ioutil`
- `net`

## Declared Types
- `errSignalingCreds`
- `optionalBoolFlag`

## Declared Constants
- `no_version`

## Declared Variables
- `version`

## Functions and Methods
### BlockingDial
- Signature: `func BlockingDial(ctx context.Context, network, address string, creds credentials.TransportCredentials, opts ...grpc.DialOption) (*grpc.ClientConn, error)`
- Exported: `true`
- Control-flow features: `if, select, go, returns error`
- Doc: BlockingDial is a helper method to dial the given address, using optional TLS credentials, and blocking until the returned connection is ready. If the given credentials are nil, the
- Internal calls: `writeResult`
- Selector calls: `ctx.Done`, `ctx.Err`, `grpc.DialContext`, `grpc.FailOnNonTempDialError`, `grpc.WithBlock`, `grpc.WithContextDialer`, `grpc.WithTransportCredentials`, `insecure.NewCredentials`

### ClientHandshake (method on `*errSignalingCreds`)
- Signature: `func (*errSignalingCreds) ClientHandshake(ctx context.Context, addr string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Selector calls: `c.writeResult`

### ClientTLSConfig
- Signature: `func ClientTLSConfig(insecureSkipVerify bool, cacertFile, clientCertFile, clientKeyFile string) (*tls.Config, error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: ClientTLSConfig builds transport-layer config for a gRPC client using the given properties. If cacertFile is blank, only standard trusted certs are used to
- Selector calls: `certPool.AppendCertsFromPEM`, `errors.New`, `fmt.Errorf`, `ioutil.ReadFile`, `tls.LoadX509KeyPair`, `x509.NewCertPool`

### ClientTransportCredentials
- Signature: `func ClientTransportCredentials(insecureSkipVerify bool, cacertFile, clientCertFile, clientKeyFile string) (credentials.TransportCredentials, error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: ClientTransportCredentials is a helper function that constructs a TLS config with the given properties (see ClientTLSConfig) and then constructs and returns gRPC
- Internal calls: `ClientTLSConfig`
- Selector calls: `credentials.NewTLS`

### ServerTransportCredentials
- Signature: `func ServerTransportCredentials(cacertFile, serverCertFile, serverKeyFile string, requireClientCerts bool) (credentials.TransportCredentials, error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: ServerTransportCredentials builds transport credentials for a gRPC server using the given properties. If cacertFile is blank, the server will not request client certs
- Selector calls: `certPool.AppendCertsFromPEM`, `credentials.NewTLS`, `errors.New`, `fmt.Errorf`, `ioutil.ReadFile`, `tls.LoadX509KeyPair`, `x509.NewCertPool`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
