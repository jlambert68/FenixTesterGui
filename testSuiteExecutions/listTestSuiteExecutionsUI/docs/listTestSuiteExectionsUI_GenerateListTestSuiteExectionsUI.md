# listTestSuiteExectionsUI_GenerateListTestSuiteExectionsUI.go

## File Overview
- Path: `testSuiteExecutions/listTestSuiteExecutionsUI/listTestSuiteExectionsUI_GenerateListTestSuiteExectionsUI.go`
- Package: `listTestSuiteExecutionsUI`
- Functions/Methods: `3`
- Imports: `15`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateListTestSuiteExecutionsUI`
- `LoadOneTestSuiteExecutionPerTestSuiteFromDataBaseFunction`

## Imports
- `FenixTesterGui/executions/detailedExecutionsUI`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testSuiteExecutions/listTestSuiteExecutionsModel`
- `FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel`
- `embed`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `image`
- `image/color`
- `strconv`

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
### LoadOneTestSuiteExecutionPerTestSuiteFromDataBaseFunction
- Signature: `func LoadOneTestSuiteExecutionPerTestSuiteFromDataBaseFunction(testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct, updateGui bool, testSuiteInstructionPreViewObject *TestSuiteInstructionPreViewStruct)`
- Exported: `true`
- Control-flow features: `if, go`
- Doc: Define the function to be executed to load TestSuiteExecutions from that Database that the user can view Only loads one TestSuiteExecution per TestCase
- Internal calls: `SortGuiTableOnCurrentColumnAndSorting`
- Selector calls: `testSuiteInstructionPreViewObject.ClearTestSuiteExecutionPreviewContainer`, `loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Disable`, `listTestSuiteExecutionsModel.LoadTestSuiteExecutionsThatCanBeViewedByUser`

### GenerateListTestSuiteExecutionsUI
- Signature: `func GenerateListTestSuiteExecutionsUI(testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct, detailedTestSuiteExecutionsUITabObject *container.DocTabs, exitFunctionsForDetailedTestSuiteExecutionsUITabObject *map[*container.TabItem]func(), testCaseInstructionPreViewObject *TestSuiteInstructionPreViewStruct) listTestCasesUI fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `if, go`
- Doc: Create the UI used for list all TestCasesMapPtr that the User can edit
- Internal calls: `LoadOneTestSuiteExecutionPerTestSuiteFromDataBaseFunction`, `loadTestSuiteExecutionListTableTable`, `calculateAndSetCorrectColumnWidths`, `updateTestSuiteExecutionsListTable`, `sortGuiTableAscendingOnTestSuiteExecutionTimeStamp`, `generateTestSuiteExecutionsListTable`, `generatePreViewObject`, `NewHoverableRect`
- Selector calls: `widget.NewButton`, `fmt.Println`, `strconv.Itoa`, `numberOfTestSuiteExecutionsAfterLocalFilters.Set`, `fmt.Sprintf`, `numberOfTestSuiteExecutionsInTheDatabaseSearch.Set`, `fyne.CurrentApp`, `testCaseInstructionPreViewObject.ClearTestSuiteExecutionPreviewContainer`

### generatePreViewObject
- Signature: `func generatePreViewObject(testCaseInstructionPreViewObject *TestSuiteInstructionPreViewStruct) tempTestCasePreviewTestInstructionExecutionLogSplitContainer *container.Split`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Generates the Container structure for the PreView-container
- Internal calls: `NewHoverableRect`
- Selector calls: `container.NewCenter`, `widget.NewLabel`, `container.NewBorder`, `container.NewScroll`, `container.NewTabItem`, `container.NewAppTabs`, `testCaseTreePreViewOverlay.Hide`, `explorerTabOverlay.Hide`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
