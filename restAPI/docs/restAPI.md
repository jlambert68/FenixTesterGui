# restAPI.go

## File Overview
- Path: `restAPI/restAPI.go`
- Package: `restAPI`
- Functions/Methods: `6`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `HealthCheck`
- `RestAPIServer`
- `RestSendAreYouAliveToFenixGuiBuilderServer`
- `RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer`
- `RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer`
- `RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer`

## Imports
- `FenixTesterGui/grpc_out_GuiTestCaseBuilderServer`
- `encoding/json`
- `fmt`
- `github.com/gorilla/mux`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `log`
- `net/http`

## Declared Types
- `RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct`
- `RestUserMessageStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### RestAPIServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestAPIServer()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `log.Println`, `mux.NewRouter`, `router.HandleFunc`, `http.Handle`, `http.ListenAndServe`, `log.Fatalln`

### HealthCheck (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) HealthCheck(w http.ResponseWriter, _ *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Selector calls: `w.WriteHeader`, `fmt.Fprintf`, `log.Fatalln`

### RestSendAreYouAliveToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendAreYouAliveToFenixGuiBuilderServer(w http.ResponseWriter, _ *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Selector calls: `w.Header`, `w.WriteHeader`, `json.Marshal`, `w.Write`, `log.Fatalln`

### RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Internal calls: `extractAndValidateJsonBody`
- Selector calls: `w.Header`, `w.WriteHeader`, `json.Marshal`, `fmt.Fprintf`, `err.Error`, `log.Fatalln`, `w.Write`

### RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Internal calls: `extractAndValidateJsonBody`
- Selector calls: `w.Header`, `w.WriteHeader`, `json.Marshal`, `fmt.Fprintf`, `err.Error`, `log.Fatalln`, `w.Write`

### RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Internal calls: `extractAndValidateJsonBody`
- Selector calls: `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `w.Header`, `w.WriteHeader`, `json.Marshal`, `fmt.Fprintf`, `err.Error`, `log.Fatalln`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
