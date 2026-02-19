# testSuitesModel_ExecutionTestSuiteWithAllTestDataSets.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_ExecutionTestSuiteWithAllTestDataSets.go`
- Package: `testSuitesModel`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ExecuteOneTestSuiteWithAllItsTestDataSets`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `FenixTesterGui/soundEngine`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ExecuteOneTestSuiteWithAllItsTestDataSets (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) ExecuteOneTestSuiteWithAllItsTestDataSets()`
- Exported: `true`
- Control-flow features: `if`
- Doc: ExecuteOneTestSuiteWithAllItsTestDataSets Execute one TestSuit's all TestCases with all its TestDataSets
- Selector calls: `errors.New`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fmt.Println`, `fmt.Sprintf`, `fyne.CurrentApp`, `grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion`, `testSuiteModel.GetTestSuiteUuid`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
