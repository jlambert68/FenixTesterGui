# gRPCOutHelpers.go

## File Overview
- Path: `grpc_out_GuiExecutionServer/gRPCOutHelpers.go`
- Package: `grpc_out_GuiExecutionServer`
- Functions/Methods: `6`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GetHighestFenixGuiExecutionServerProtoFileVersion`
- `SetConnectionToFenixGuiExecutionServer`
- `SetDialAddressString`
- `SetLogger`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/gcp`
- `FenixTesterGui/grpcurl`
- `context`
- `crypto/tls`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `google.golang.org/grpc`
- `google.golang.org/grpc/credentials`
- `log`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `isUnixSocket`

## Functions and Methods
### setConnectionToFenixGuiExecutionMessageServer_new (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) setConnectionToFenixGuiExecutionMessageServer_new(ctx context.Context) (_ context.Context, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: SetConnectionToFenixExecutionWorkerServer - Set upp connection and Dial to FenixExecutionServer
- Internal calls: `dialFromGrpcurl`
- Selector calls: `gcp.GRPCDialer`, `grpc.Dial`, `grpc.WithInsecure`, `fenixExecutionServerGuiGrpcApi.NewFenixExecutionServerGuiGrpcServicesForGuiClientClient`, `time.Sleep`, `time.Duration`

### dialFromGrpcurl
- Signature: `func dialFromGrpcurl(ctx context.Context, target string) (context.Context, *grpc.ClientConn)`
- Exported: `false`
- Control-flow features: `if, defer`
- Internal calls: `cancel`, `isUnixSocket`
- Selector calls: `context.WithTimeout`, `credentials.NewTLS`, `grpc.WithUserAgent`, `grpcurl.BlockingDial`, `log.Panicln`, `err.Error`

### SetConnectionToFenixGuiExecutionServer (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SetConnectionToFenixGuiExecutionServer() returnMessage *fenixExecutionServerGuiGrpcApi.AckNackResponse`
- Exported: `true`
- Control-flow features: `if`
- Doc: Set upp connection and Dial to FenixGuiExecutionServer
- Internal calls: `GetHighestFenixGuiExecutionServerProtoFileVersion`
- Selector calls: `credentials.NewTLS`, `grpc.WithTransportCredentials`, `grpc.Dial`, `grpc.WithInsecure`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fenixExecutionServerGuiGrpcApi.NewFenixExecutionServerGuiGrpcServicesForGuiClientClient`

### GetHighestFenixGuiExecutionServerProtoFileVersion
- Signature: `func GetHighestFenixGuiExecutionServerProtoFileVersion() int32`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: GetHighestFenixGuiServerProtoFileVersion ******************************************************************************************************************** Get the highest FenixProtoFileVersionEnumeration

### SetLogger (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same Logger reference as is used by central part of system

### SetDialAddressString (method on `*GRPCOutGuiExecutionServerStruct`)
- Signature: `func (*GRPCOutGuiExecutionServerStruct) SetDialAddressString(dialAddress string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetDialAddressString Set the Dial Address, which was received from environment variables

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
