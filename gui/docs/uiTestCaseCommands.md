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
### copy (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) copy(element string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Copy(Element)
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`

### createTestCaseCommandParametersUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) createTestCaseCommandParametersUI() testCaseCommandsParametersUIObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `container.New`, `container.NewScroll`, `fmt.Printf`, `fmt.Println`, `layout.NewVBoxLayout`

### createTestCaseCommandsUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) createTestCaseCommandsUI() testCaseCommandsUIObject fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `availableBuildingBlocksInTestCaseSelectWidget.Refresh`, `container.New`, `container.NewScroll`, `fmt.Printf`, `fmt.Println`, `layout.NewFormLayout`, `layout.NewHBoxLayout`, `layout.NewVBoxLayout`

### cut (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) cut(element string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Cut(Element)
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`

### getUuidFromTreeName (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) getUuidFromTreeName(uiTreeName string) (buildingBlockUuid string, buildingBlockType BuildingBlock, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: GetUuidFromUiName Finds the UUID for from a UI-name like ' B0_BOND [3c8a3bc] [BOND] to live forever..'
- Selector calls: `errors.New`, `fmt.Sprintf`, `strings.Index`

### newTestCase (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) newTestCase()`
- Exported: `false`
- Control-flow features: `if`
- Doc: NewTestCase()
- Selector calls: `availableBuildingBlocksInTestCaseSelectWidget.Refresh`, `availableBuildingBlocksSelectWidget.Refresh`, `availableTestCasesSelectWidget.Refresh`, `bindedCommandListData.Prepend`, `fmt.Printf`, `fmt.Println`

### remove (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) remove(testCaseUuid string, elementUiNameoBeRemoved string)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Remove(ElementToBeRemoved)
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`, `fmt.Println`

### swapFromCopyBuffer (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) swapFromCopyBuffer(elementTobeSwappedOut string, copyBufferElementTobeSwappedIn string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: SwapFromCopyBuffer(ElementTobeSwappedOut, CopyBufferElementTobeSwappedIn)
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`

### swapFromCutBuffer (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) swapFromCutBuffer(elementTobeSwappedOut string, cutBufferElementTobeSwappedIn string)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: SwapFromCutBuffer(ElementTobeSwappedOut, CutBufferElementTobeSwappedIn)
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`

### swapFromNew (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) swapFromNew(testCaseUuid string, elementUiNameTobeSwappedOut string, newElementUiNameTobeSwappedIn string)`
- Exported: `false`
- Control-flow features: `if, switch`
- Doc: SwapFromNew(ElementTobeSwappedOut, NewElementTobeSwappedIn)
- Selector calls: `availableBuildingBlocksInTestCaseSelectWidget.Refresh`, `availableBuildingBlocksSelectWidget.Refresh`, `bindedCommandListData.Prepend`, `errors.New`, `fmt.Printf`, `fmt.Println`, `fmt.Sprintf`, `uiServer.getUuidFromTreeName`

### undoLastCommandOnStack (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) undoLastCommandOnStack()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: UndoLastCommandOnStack()
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`

### undoUndoLastCommandOnStack (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) undoUndoLastCommandOnStack()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: UndoUndoLastCommandOnStack()
- Selector calls: `bindedCommandListData.Prepend`, `fmt.Printf`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
