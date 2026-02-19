# restApiHelpers.go

## File Overview
- Path: `restAPI/restApiHelpers.go`
- Package: `restAPI`
- Functions/Methods: `3`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SetDialAddressString`
- `SetLogger`

## Imports
- `encoding/json`
- `errors`
- `fmt`
- `github.com/golang/gddo/httputil/header`
- `github.com/sirupsen/logrus`
- `io`
- `log`
- `net/http`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### extractAndValidateJsonBody
- Signature: `func extractAndValidateJsonBody(responseWriterPointer *http.ResponseWriter, httpRequest *http.Request, myInputTypeVariable interface{}) err error`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: ******************************************************************************* extractAndValidateJsonBody
- Selector calls: `header.ParseValueAndParams`, `http.Error`, `http.MaxBytesReader`, `json.NewDecoder`, `dec.DisallowUnknownFields`, `dec.Decode`, `errors.As`, `fmt.Sprintf`

### SetLogger (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: ******************************************************************************* SetLogger

### SetDialAddressString (method on `*RestApiStruct`)
- Signature: `func (*RestApiStruct) SetDialAddressString(dialAddress string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: ******************************************************************************* SetDialAddressString

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
