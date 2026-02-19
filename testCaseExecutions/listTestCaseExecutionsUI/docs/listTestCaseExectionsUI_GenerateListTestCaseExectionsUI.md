# listTestCaseExectionsUI_GenerateListTestCaseExectionsUI.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/listTestCaseExectionsUI_GenerateListTestCaseExectionsUI.go`
- Package: `listTestCaseExecutionsUI`
- Functions/Methods: `3`
- Imports: `16`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateListTestCaseExecutionsUI`
- `LoadOneTestCaseExecutionPerTestCaseFromDataBaseFunction`

## Imports
- `FenixTesterGui/executions/detailedExecutionsUI`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCaseExecutions/listTestCaseExecutionsModel`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
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
### LoadOneTestCaseExecutionPerTestCaseFromDataBaseFunction
- Signature: `func LoadOneTestCaseExecutionPerTestCaseFromDataBaseFunction(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct, updateGui bool, testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct)`
- Exported: `true`
- Control-flow features: `if, go`
- Doc: Define the function to be executed to load TestCaseExecutions from that Database that the user can view Only loads one TestCaseExecution per TestCase
- Internal calls: `SortGuiTableOnCurrentColumnAndSorting`
- Selector calls: `testCaseInstructionPreViewObject.ClearTestCaseExecutionPreviewContainer`, `loadAllTestCaseExecutionsForOneTestCaseButtonReference.Disable`, `listTestCaseExecutionsModel.LoadTestCaseExecutionsThatCanBeViewedByUser`

### GenerateListTestCaseExecutionsUI
- Signature: `func GenerateListTestCaseExecutionsUI(testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct, detailedTestCaseExecutionsUITabObject *container.DocTabs, exitFunctionsForDetailedTestCaseExecutionsUITabObject *map[*container.TabItem]func(), testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) listTestCasesUI fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `if, for/range, go`
- Doc: Create the UI used for list all TestCasesMapPtr that the User can edit
- Internal calls: `LoadOneTestCaseExecutionPerTestCaseFromDataBaseFunction`, `loadTestCaseExecutionListTableTable`, `calculateAndSetCorrectColumnWidths`, `updateTestCaseExecutionsListTable`, `sortGuiTableAscendingOnTestCaseExecutionTimeStamp`, `generateTestCaseExecutionsListTable`, `generatePreViewObject`, `NewHoverableRect`
- Selector calls: `widget.NewButton`, `fmt.Println`, `strconv.Itoa`, `numberOfTestCaseExecutionsAfterLocalFilters.Set`, `fmt.Sprintf`, `numberOfTestCaseExecutionsInTheDatabaseSearch.Set`, `fyne.CurrentApp`, `testCaseInstructionPreViewObject.ClearTestCaseExecutionPreviewContainer`

### generatePreViewObject
- Signature: `func generatePreViewObject(testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) tempTestCasePreviewTestInstructionExecutionLogSplitContainer *container.Split`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Generates the Container structure for the PreView-container
- Internal calls: `NewHoverableRect`
- Selector calls: `container.NewCenter`, `widget.NewLabel`, `container.NewBorder`, `container.NewScroll`, `container.NewTabItem`, `container.NewAppTabs`, `testCaseTreePreViewOverlay.Hide`, `explorerTabOverlay.Hide`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
