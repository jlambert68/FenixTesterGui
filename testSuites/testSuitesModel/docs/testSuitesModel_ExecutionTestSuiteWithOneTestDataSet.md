# testSuitesModel_ExecutionTestSuiteWithOneTestDataSet.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_ExecutionTestSuiteWithOneTestDataSet.go`
- Package: `testSuitesModel`
- Functions/Methods: `1`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ExecuteOneTestSuiteWithOneTestDataSet`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `FenixTesterGui/soundEngine`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/dialog`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ExecuteOneTestSuiteWithOneTestDataSet (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) ExecuteOneTestSuiteWithOneTestDataSet()`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: ExecuteOneTestSuiteWithOneTestDataSet Execute one TestSuit's all TestCases with one TestDataSet
- Internal calls: `string`
- Selector calls: `dialog.ShowConfirm`, `errors.New`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `fmt.Println`, `fmt.Sprintf`, `fyne.CurrentApp`, `grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion`, `testDataEngine.TestDataAreaUuidType`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
