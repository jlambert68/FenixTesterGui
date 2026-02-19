# clickableTInTICNameInPreview.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/clickableTInTICNameInPreview.go`
- Package: `listTestSuiteExecutionsUI`
- Functions/Methods: `4`
- Imports: `16`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `image/color`
- `log`
- `strconv`
- `strings`
- `time`

## Declared Types
- `clickableTInTICNameLabelInPreviewStruct`
- `labelTypeType`

## Declared Constants
- `labelIsTestInstruction`
- `labelIsTestInstructionContainer`
- `notDefined`

## Declared Variables
- None

## Functions and Methods
### newClickableTestInstructionNameLabelInPreview
- Signature: `func newClickableTestInstructionNameLabelInPreview(tInTICName string, tCExecutionKey testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType, tInTICExecutionKey testCaseExecutionsModel.
	TCEoTICoTIEAttributesContainerMapKeyType, leftClicked func(), rightClicked func(), labelType labelTypeType, testCaseInstructionPreViewObject *TestSuiteInstructionPreViewStruct) *clickableTInTICNameLabelInPreviewStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new TestInstructionName label
- Selector calls: `time.Now`, `clickableTInTICNameLabelInPreview.ExtendBaseWidget`

### CreateRenderer (method on `*clickableTInTICNameLabelInPreviewStruct`)
- Signature: `func (*clickableTInTICNameLabelInPreviewStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: CreateRenderer Renderer (required by fyne.Widget)
- Selector calls: `widget.NewLabel`, `widget.NewSimpleRenderer`

### Tapped (method on `*clickableTInTICNameLabelInPreviewStruct`)
- Signature: `func (*clickableTInTICNameLabelInPreviewStruct) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if, for/range, switch`
- Doc: Tapped Tapped interface implementation
- Internal calls: `int32`
- Selector calls: `testSuiteExecutionAttributesForPreviewMapMutex.Lock`, `testSuiteExecutionAttributesForPreviewMapMutex.Unlock`, `attributesContainerPtr.Hide`, `attributesContainerPtr.Show`, `attributesContainerPtr.Refresh`, `testCaseExecutionsModel.TestInstructionExecutionLogPostMapKeyType`, `testCaseExecutionsModel.TestInstructionExecutionDetailsMapKeyType`, `container.New`

### TappedSecondary (method on `*clickableTInTICNameLabelInPreviewStruct`)
- Signature: `func (*clickableTInTICNameLabelInPreviewStruct) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TappedSecondary Optional: Handle secondary tap (right-click)
- Selector calls: `fenixMasterWindow.Clipboard`, `clipboard.SetContent`, `fyne.CurrentApp`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
