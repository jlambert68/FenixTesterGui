# common_helpers.go

## File Overview
- Path: `common_code/common_helpers.go`
- Package: `sharedCode`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateTestDataHeaderItemMessageHash`
- `HashSingleValue`

## Imports
- `crypto/sha256`
- `encoding/hex`
- `github.com/jlambert68/FenixGrpcApi/Fenix/fenixTestDataSyncServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### CreateTestDataHeaderItemMessageHash
- Signature: `func CreateTestDataHeaderItemMessageHash(testDataHeaderItemMessage *fenixTestDataSyncServerGrpcApi.TestDataHeaderItemMessage) testDataHeaderItemMessageHash string`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: Exctract Values, and create, for TestDataHeaderItemMessageHash
- Internal calls: `HashValues`
- External calls: `headerFilterValue.String`

### HashSingleValue
- Signature: `func HashSingleValue(valueToHash string) hashValue string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Hash a single value
- External calls: `hash.Sum`, `hash.Write`, `hex.EncodeToString`, `sha256.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
