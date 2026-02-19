# newOrEditTestDataPointGroupUI_SelectedTestDataComponent.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_SelectedTestDataComponent.go`
- Package: `newOrEditTestDataPointGroupUI`
- Functions/Methods: `2`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateSelectedPointsListUIComponent
- Signature: `func generateSelectedPointsListUIComponent(newOrEditTestDataPointGroupWindowPtr *fyne.Window, incomingGroupName testDataEngine.TestDataPointGroupNameType, isNew bool, newOrEditedChosenTestDataPointsThisGroupMapPtr *map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType, testDataForGroupObject *testDataEngine.TestDataForGroupObjectStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Create and configure the list-component of selected TestDataPoints
- Internal calls: `setStateForSaveButtonAndGroupNameTextEntry`, `string`
- Selector calls: `container.NewBorder`, `container.NewHBox`, `container.NewVBox`, `fmt.Sprintf`, `nameEntry.SetPlaceHolder`, `nameEntry.SetText`, `newOrEditTestDataPointGroupWindow.Close`, `testDataEngine.TestDataPointGroupNameType`

### setStateForSaveButtonAndGroupNameTextEntry
- Signature: `func setStateForSaveButtonAndGroupNameTextEntry(entryValue string, nameStatusLabel *widget.Label, saveButton *widget.Button, isNew bool, incomingGroupName testDataEngine.TestDataPointGroupNameType, testDataForGroupObject *testDataEngine.TestDataForGroupObjectStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Set the State for Save Button and the GroupName Entry
- Internal calls: `string`
- Selector calls: `nameStatusLabel.SetText`, `saveButton.Disable`, `saveButton.Enable`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
