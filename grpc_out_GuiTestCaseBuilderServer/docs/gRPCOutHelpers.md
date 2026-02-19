# gRPCOutHelpers.go

## File Overview
- Path: `grpc_out_GuiTestCaseBuilderServer/gRPCOutHelpers.go`
- Package: `grpc_out_GuiTestCaseBuilderServer`
- Functions/Methods: `6`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`
- `SetDialAddressString`
- `SetLogger`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpcurl`
- `context`
- `crypto/tls`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
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
### setConnectionToFenixGuiTestCaseBuilderServer_new (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) setConnectionToFenixGuiTestCaseBuilderServer_new(ctx context.Context) (_ context.Context, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: SetConnectionToFenixExecutionWorkerServer - Set upp connection and Dial to FenixExecutionServer
- Internal calls: `dialFromGrpcurl`
- Selector calls: `grpc.Dial`, `grpc.WithInsecure`, `fenixTestCaseBuilderServerGrpcApi.NewFenixTestCaseBuilderServerGrpcServicesClient`, `time.Sleep`, `time.Duration`

### dialFromGrpcurl
- Signature: `func dialFromGrpcurl(ctx context.Context, target string) (context.Context, *grpc.ClientConn)`
- Exported: `false`
- Control-flow features: `if, defer`
- Internal calls: `cancel`, `isUnixSocket`
- Selector calls: `context.WithTimeout`, `credentials.NewTLS`, `grpc.WithUserAgent`, `grpcurl.BlockingDial`, `log.Panicln`, `err.Error`

### setConnectionToFenixGuiTestCaseBuilderServer (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) setConnectionToFenixGuiTestCaseBuilderServer() returnMessage *fenixTestCaseBuilderServerGrpcApi.AckNackResponse`
- Exported: `false`
- Control-flow features: `if`
- Doc: Set upp connection and Dial to FenixGuiTestCaseBuilderServer
- Selector calls: `credentials.NewTLS`, `grpc.WithTransportCredentials`, `grpc.Dial`, `grpc.WithInsecure`, `fenixTestCaseBuilderServerGrpcApi.NewFenixTestCaseBuilderServerGrpcServicesClient`

### GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion() int32`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion ******************************************************************************************************************** Get the highest FenixProtoFileVersionEnumeration

### SetLogger (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same Logger reference as is used by central part of system

### SetDialAddressString (method on `*GRPCOutGuiTestCaseBuilderServerStruct`)
- Signature: `func (*GRPCOutGuiTestCaseBuilderServerStruct) SetDialAddressString(dialAddress string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetDialAddressString Set the Dial Address, which was received from environment variables

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
