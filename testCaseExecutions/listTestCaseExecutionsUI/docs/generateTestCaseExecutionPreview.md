# generateTestCaseExecutionPreview.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/generateTestCaseExecutionPreview.go`
- Package: `listTestCaseExecutionsUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `16`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ClearTestCaseExecutionPreviewContainer`
- `GenerateTestCaseExecutionPreviewContainer`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `bytes`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `image/color`
- `image/png`
- `log`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ClearTestCaseExecutionPreviewContainer (method on `*TestCaseInstructionPreViewStruct`)
- Signature: `func (*TestCaseInstructionPreViewStruct) ClearTestCaseExecutionPreviewContainer()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: ClearTestCaseExecutionPreviewContainer Clears the preview area for the TestCaseExecution
- External calls: `container.NewCenter`, `widget.NewLabel`

### GenerateTestCaseExecutionPreviewContainer (method on `*TestCaseInstructionPreViewStruct`)
- Signature: `func (*TestCaseInstructionPreViewStruct) GenerateTestCaseExecutionPreviewContainer(testCaseExecutionUuid string, testCaseExecutionVersion uint32, testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct, openedTestCaseExecutionFrom OpenedTestCaseExecutionOrTestSuiteExecutionFromType, currentWindowPtr *fyne.Window, testCasePreviewStructureMessageFromTestSuiteExecution *fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage, testTestInstructionsExecutionStatusPreviewValuesFromTestSuiteExecution *fenixExecutionServerGuiGrpcApi.
	TestInstructionsExecutionStatusPreviewValuesMessage)`
- Exported: `true`
- Control-flow features: `if, for/range, switch, defer`
- Internal calls: `float32`, `generatePreViewObject`, `int`, `int32`, `newClickableAttributeInPreview`, `newClickableTestInstructionNameLabelInPreview`, `newCopyableLabel`, `openedDetailedTestCaseExecutionsMapKeyType`
- External calls: `btn.Size`, `bytes.NewReader`, `canvas.NewImageFromImage`, `canvas.NewRectangle`, `canvas.NewText`, `container.New`, `container.NewBorder`, `container.NewCenter`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
