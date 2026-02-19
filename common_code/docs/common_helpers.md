# common_helpers.go

## File Overview
- Path: `common_code/common_helpers.go`
- Package: `sharedCode`
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
- Selector calls: `headerFilterValue.String`

### HashSingleValue
- Signature: `func HashSingleValue(valueToHash string) hashValue string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: Hash a single value
- Selector calls: `sha256.New`, `hash.Write`, `hex.EncodeToString`, `hash.Sum`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
