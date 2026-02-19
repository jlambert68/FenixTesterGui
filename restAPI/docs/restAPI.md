# restAPI.go

## File Overview
- Path: `restAPI/restAPI.go`
- Package: `restAPI`
- Generated: `2026-02-19T14:23:17+01:00`
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
### HealthCheck (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) HealthCheck(w http.ResponseWriter, _ *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- External calls: `fmt.Fprintf`, `log.Fatalln`, `w.WriteHeader`

### RestAPIServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestAPIServer()`
- Exported: `true`
- Control-flow features: `if`
- External calls: `http.Handle`, `http.ListenAndServe`, `log.Fatalln`, `log.Println`, `mux.NewRouter`, `router.HandleFunc`

### RestSendAreYouAliveToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendAreYouAliveToFenixGuiBuilderServer(w http.ResponseWriter, _ *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- External calls: `json.Marshal`, `log.Fatalln`, `w.Header`, `w.Write`, `w.WriteHeader`

### RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Internal calls: `extractAndValidateJsonBody`
- External calls: `err.Error`, `fmt.Fprintf`, `json.Marshal`, `log.Fatalln`, `w.Header`, `w.Write`, `w.WriteHeader`

### RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Internal calls: `extractAndValidateJsonBody`
- External calls: `err.Error`, `fmt.Fprintf`, `json.Marshal`, `log.Fatalln`, `w.Header`, `w.Write`, `w.WriteHeader`

### RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request)`
- Exported: `true`
- Control-flow features: `if, defer`
- Internal calls: `extractAndValidateJsonBody`
- External calls: `err.Error`, `fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum`, `fmt.Fprintf`, `grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion`, `json.Marshal`, `log.Fatalln`, `w.Header`, `w.Write`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
