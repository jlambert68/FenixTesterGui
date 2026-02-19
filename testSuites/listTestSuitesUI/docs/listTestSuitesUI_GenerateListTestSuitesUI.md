# listTestSuitesUI_GenerateListTestSuitesUI.go

## File Overview
- Path: `testSuites/listTestSuitesUI/listTestSuitesUI_GenerateListTestSuitesUI.go`
- Package: `listTestSuitesUI`
- Functions/Methods: `2`
- Imports: `17`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateListTestSuitesUI`
- `InitiateListTestSuiteUIObject`

## Imports
- `FenixTesterGui/executions/detailedExecutionsUI`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/listTestSuitesModel`
- `embed`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `image`
- `image/color`
- `strconv`
- `sync`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `imageData_tic_parallellImage`
- `imageData_tic_serialImage`
- `sortImageAscendingAsByteArray`
- `sortImageAscendingAsImage`
- `sortImageDescendingAsByteArray`
- `sortImageDescendingAsImage`
- `sortImageUnspecifiedAsImage`
- `sortUnspecifiedImageAsByteArray`
- `tic_parallellImage`
- `tic_serialImage`

## Functions and Methods
### InitiateListTestSuiteUIObject
- Signature: `func InitiateListTestSuiteUIObject(tempHowShouldItBeUsed UsedForTestSuitesListType, selectedTestSuitesPtr *map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage) listTestCaseUIObject *ListTestSuiteUIStruct`
- Exported: `true`
- Control-flow features: `if`

### GenerateListTestSuitesUI (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) GenerateListTestSuitesUI(testCasesModel *testCaseModel.TestCasesModelsStruct, preViewAndFilterTabsUsedForCreateTestSuite *container.AppTabs) _ *fyne.Container`
- Exported: `true`
- Control-flow features: `if`
- Doc: Create the UI used for list all TestSuitesMapPtr that the User can edit
- Internal calls: `filterTestSuitesButtonFunction`, `NewHoverableRect`
- Selector calls: `fmt.Println`, `listTestSuitesModel.LoadtestSuiteThatCanBeEditedByUser`, `time.Now`, `widget.NewButton`, `listTestSuiteUIObject.loadTestSuiteListTableTable`, `listTestSuiteUIObject.calculateAndSetCorrectColumnWidths`, `listTestSuiteUIObject.updateTestSuitesListTable`, `strconv.Itoa`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
