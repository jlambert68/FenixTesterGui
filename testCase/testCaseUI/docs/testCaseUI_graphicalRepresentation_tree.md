# testCaseUI_graphicalRepresentation_tree.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_graphicalRepresentation_tree.go`
- Package: `testCaseUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `3`
- Imports: `14`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testUIDragNDropStatemachine`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `image/color`
- `strconv`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### convertRGBAHexStringIntoRGBAColor (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) convertRGBAHexStringIntoRGBAColor(rgbaHexString string) (rgbaValue color.RGBA, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Converts a colors in a hex-string into 'color.RGBA'-format. "#FF03AFFF"
- Internal calls: `uint8`
- External calls: `errors.New`, `fmt.Sprintf`, `strconv.ParseInt`

### makeTestCaseGraphicalUIObject (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) makeTestCaseGraphicalUIObject(testCaseUuid string) testCaseCanvasObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the graphical structure for the TestCase
- External calls: `container.NewVBox`, `err.Error`, `testCasesUiCanvasObject.recursiveMakeTestCaseGraphicalUIObject`, `widget.NewLabel`

### recursiveMakeTestCaseGraphicalUIObject (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) recursiveMakeTestCaseGraphicalUIObject(uuid string, testCaseModelForUITree *map[string][]testCaseModel.TestCaseModelAdaptedForUiTreeDataStruct, firstAccordion *clickableAccordion, nodeTreeLevel float32, testCaseUuid string, testcaseTreeContainer *fyne.Container, testCasesModel *testCaseModel.TestCaseModelStruct) testCaseCanvasObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Generates the graphical structure for the TestCase
- Internal calls: `float32`, `uint32`
- External calls: `canvas.NewRectangle`, `container.NewHBox`, `container.NewStack`, `container.NewVBox`, `err.Error`, `errors.New`, `fmt.Sprintf`, `fyne.NewSize`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
