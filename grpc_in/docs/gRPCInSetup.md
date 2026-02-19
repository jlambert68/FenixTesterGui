# gRPCInSetup.go

## File Overview
- Path: `grpc_in/gRPCInSetup.go`
- Package: `grpc_in`
- Functions/Methods: `2`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitGrpcServer`
- `StopGrpcServer`

## Imports
- `FenixTesterGui/common_code`
- `github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `google.golang.org/grpc`
- `google.golang.org/grpc/reflection`
- `net`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### InitGrpcServer (method on `*GRPCInStruct`)
- Signature: `func (*GRPCInStruct) InitGrpcServer()`
- Exported: `true`
- Control-flow features: `if`
- Doc: InitGrpcServer - Set up and start Backend gRPC-server
- Selector calls: `net.Listen`, `strconv.Itoa`, `grpc.NewServer`, `fenixUserGuiGrpcApi.RegisterFenixUserGuiGrpcServicesServer`, `reflection.Register`, `registerFenixUserGuiServer.Serve`

### StopGrpcServer (method on `*GRPCInStruct`)
- Signature: `func (*GRPCInStruct) StopGrpcServer()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: StopGrpcServer - Stop Backend gRPC-server
- Selector calls: `registerFenixUserGuiServer.GracefulStop`, `lis.Close`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
