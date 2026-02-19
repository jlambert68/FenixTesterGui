# testSuitesModel_IsTestSuiteChanged.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_IsTestSuiteChanged.go`
- Package: `testSuitesModel`
- Functions/Methods: `3`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `IsTestSuiteChanged`

## Imports
- `FenixTesterGui/common_code`
- `fmt`
- `log`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### IsTestSuiteChanged (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) IsTestSuiteChanged() testSuiteIsChanged bool`
- Exported: `true`
- Control-flow features: `if`
- Doc: IsTestSuiteChanged Checks if the TestSuite content has been changed from last saved occasion
- Selector calls: `testSuiteModel.generateTesSuiteMetaDataHash`, `testSuiteModel.generateTesSuiteTestDataHash`

### generateTesSuiteMetaDataHash (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTesSuiteMetaDataHash() testSuiteMetaDataHash string`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Selector calls: `fmt.Sprintf`, `log.Fatalln`, `sharedCode.HashValues`

### generateTesSuiteTestDataHash (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateTesSuiteTestDataHash() testSuiteMetaDataHash string`
- Exported: `false`
- Control-flow features: `for/range`
- Selector calls: `fmt.Sprintf`, `sharedCode.HashValues`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
