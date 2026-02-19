# guiExecutionServer_helpers.go

## File Overview
- Path: `messageStreamEngine/guiExecutionServer_helpers.go`
- Package: `messageStreamEngine`
- Functions/Methods: `5`
- Imports: `14`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateAndStartMessageStreamChannelReader`
- `MetadataFromHeaders`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `FenixTesterGui/grpcurl`
- `context`
- `crypto/tls`
- `encoding/base64`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `google.golang.org/grpc`
- `google.golang.org/grpc/credentials`
- `google.golang.org/grpc/metadata`
- `log`
- `strings`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `base64Codecs`
- `isUnixSocket`

## Functions and Methods
### InitiateAndStartMessageStreamChannelReader
- Signature: `func InitiateAndStartMessageStreamChannelReader()`
- Exported: `true`
- Control-flow features: `go`
- Doc: InitiateAndStartMessageStreamChannelReader Initiate the channel reader which is used for reading and processing messages that were received from GuiExecutionServer
- Internal calls: `initiatePubSubFunctionality`
- Selector calls: `messageStreamEngineObject.startCommandChannelReader`

### MetadataFromHeaders
- Signature: `func MetadataFromHeaders(headers []string) metadata.MD`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: MetadataFromHeaders converts a list of header strings (each string in "Header-Name: Header-Value" form) into metadata. If a string has a header
- Internal calls: `decode`
- Selector calls: `strings.HasSuffix`, `strings.SplitN`, `strings.ToLower`, `strings.TrimSpace`

### decode
- Signature: `func decode(val string) (string, error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Internal calls: `string`
- Selector calls: `d.DecodeString`

### dialFromGrpcurl
- Signature: `func dialFromGrpcurl(ctx context.Context, target string) (context.Context, *grpc.ClientConn)`
- Exported: `false`
- Control-flow features: `if, defer`
- Internal calls: `cancel`, `isUnixSocket`
- Selector calls: `context.WithTimeout`, `credentials.NewTLS`, `err.Error`, `grpc.WithUserAgent`, `grpcurl.BlockingDial`, `log.Panicln`

### setConnectionToFenixGuiExecutionMessageServer (method on `*MessageStreamEngineStruct`)
- Signature: `func (*MessageStreamEngineStruct) setConnectionToFenixGuiExecutionMessageServer(ctx context.Context) (_ context.Context, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: SetConnectionToFenixExecutionWorkerServer - Set upp connection and Dial to FenixExecutionServer
- Internal calls: `dialFromGrpcurl`
- Selector calls: `fenixExecutionServerGuiGrpcApi.NewFenixExecutionServerGuiGrpcServicesForGuiClientClient`, `grpc.Dial`, `grpc.WithInsecure`, `time.Duration`, `time.Sleep`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
