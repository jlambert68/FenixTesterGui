# uiTestCaseCommands.go

## File Overview
- Path: `gui/uiTestCaseCommands.go`
- Package: `gui`
- Functions/Methods: `12`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `strings`

## Declared Types
- None

## Declared Constants
- `CommandCopy`
- `CommandCut`
- `CommandNewTestcase`
- `CommandRemoveElementFromTestcase`
- `CommandSwapFromCopyBuffer`
- `CommandSwapFromCutBuffer`
- `CommandSwapFromNewComponent`
- `CommandUndoLastCommandOnStack`
- `CommandUndoUndoLastCommandOnStack`

## Declared Variables
- `availableBuildingBlocksInTestCaseSelectWidget`
- `availableBuildingBlocksSelectWidget`
- `availableTestCasesSelectWidget`

## Functions and Methods
### createTestCaseCommandParametersUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) createTestCaseCommandParametersUI() testCaseCommandsParametersUIObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `fmt.Println`, `fmt.Printf`, `container.New`, `layout.NewVBoxLayout`, `container.NewScroll`

### createTestCaseCommandsUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) createTestCaseCommandsUI() testCaseCommandsUIObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `widget.NewLabel`, `fmt.Println`, `availableBuildingBlocksInTestCaseSelectWidget.Refresh`, `fmt.Printf`, `widget.NewButton`, `uiServer.newTestCase`, `uiServer.remove`, `uiServer.swapFromNew`

### newTestCase (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) newTestCase()`
- Exported: `false`
- Control-flow features: `if`
- Doc: NewTestCase()
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`, `fmt.Println`, `availableTestCasesSelectWidget.Refresh`, `availableBuildingBlocksSelectWidget.Refresh`, `availableBuildingBlocksInTestCaseSelectWidget.Refresh`

### remove (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) remove(testCaseUuid string, elementUiNameoBeRemoved string)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Remove(ElementToBeRemoved)
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`, `fmt.Println`

### swapFromNew (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) swapFromNew(testCaseUuid string, elementUiNameTobeSwappedOut string, newElementUiNameTobeSwappedIn string)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`, `fmt.Println`, `uiServer.getUuidFromTreeName`, `errors.New`, `fmt.Sprintf`, `availableBuildingBlocksSelectWidget.Refresh`, `availableBuildingBlocksInTestCaseSelectWidget.Refresh`

### copy (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) copy(element string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Copy(Element)
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`

### swapFromCopyBuffer (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) swapFromCopyBuffer(elementTobeSwappedOut string, copyBufferElementTobeSwappedIn string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: SwapFromCopyBuffer(ElementTobeSwappedOut, CopyBufferElementTobeSwappedIn)
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`

### cut (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) cut(element string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Cut(Element)
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`

### swapFromCutBuffer (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) swapFromCutBuffer(elementTobeSwappedOut string, cutBufferElementTobeSwappedIn string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: SwapFromCutBuffer(ElementTobeSwappedOut, CutBufferElementTobeSwappedIn)
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`

### undoLastCommandOnStack (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) undoLastCommandOnStack()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: UndoLastCommandOnStack()
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`

### undoUndoLastCommandOnStack (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) undoUndoLastCommandOnStack()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: UndoUndoLastCommandOnStack()
- Selector calls: `fmt.Printf`, `bindedCommandListData.Prepend`, `fmt.Println`

### getUuidFromTreeName (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) getUuidFromTreeName(uiTreeName string) (buildingBlockUuid string, buildingBlockType BuildingBlock, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: GetUuidFromUiName Finds the UUID for from a UI-name like ' B0_BOND [3c8a3bc] [BOND] to live forever..'
- Selector calls: `strings.Index`, `errors.New`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
