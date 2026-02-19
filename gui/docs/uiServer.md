# uiServer.go

## File Overview
- Path: `gui/uiServer.go`
- Package: `gui`
- Functions/Methods: `12`
- Imports: `39`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `StartUIServer`

## Imports
- `FenixTesterGui/commandAndRuleEngine`
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/executions/detailedExecutionsUI`
- `FenixTesterGui/executions/executionsModelForSubscriptions`
- `FenixTesterGui/executions/executionsModelForTestCaseExecutions`
- `FenixTesterGui/executions/executionsUIForExecutions`
- `FenixTesterGui/executions/executionsUIForSubscriptions`
- `FenixTesterGui/fenix_pig`
- `FenixTesterGui/grpc_out_GuiTestCaseBuilderServer`
- `FenixTesterGui/memoryUsage`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testCase/testCaseUI`
- `FenixTesterGui/testCaseExecutions/listTestCaseExecutionsUI`
- `FenixTesterGui/testCaseExecutions/testCaseExecutionsModel`
- `FenixTesterGui/testCaseSubscriptionHandler`
- `FenixTesterGui/testCases/listTestCasesModel`
- `FenixTesterGui/testCases/listTestCasesUI`
- `FenixTesterGui/testSuiteExecutions/listTestSuiteExecutionsUI`
- `FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel`
- `FenixTesterGui/testSuites/listTestSuitesUI`
- `FenixTesterGui/testSuites/testSuitesTabsUI`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/app`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/sirupsen/logrus`
- `image/color`
- `log`
- `os`
- `strconv`
- `time`

## Declared Types
- `CustomButton`
- `MouseHandler`
- `customRectangle`

## Declared Constants
- None

## Declared Variables
- `image`

## Functions and Methods
### StartUIServer (method on `*GlobalUIServerStruct`)
- Signature: `func (*GlobalUIServerStruct) StartUIServer()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `uiServer.SetLogger`, `uiServer.SetDialAddressString`, `detailedExecutionsModel.InitiateCommandChannelReaderForDetailedStatusUpdates`, `uiServer.startTestCaseUIServer`

### startTestCaseUIServer (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) startTestCaseUIServer()`
- Exported: `false`
- Control-flow features: `if, go`
- Doc: Main UI server module
- Internal calls: `createSplashWindow`, `fn`, `int`
- Selector calls: `app.NewWithID`, `theme.DarkTheme`, `listTestCasesModel.LoadTestCaseThatCanBeEditedByUser`, `time.Now`, `testCaseExecutionsModel.InitiateTestCaseExecutionModel`, `executionsModelForSubscriptions.InitiateAndStartChannelsUsedByListModel`, `executionsUIForSubscriptions.StartTableAddAndRemoveChannelReaders`, `executionsModelForSubscriptions.InitiateSubscriptionModelForTestCaseOnExecutionQueue`

### newCustomRect (method on `*customRectangle`)
- Signature: `func (*customRectangle) newCustomRect() *customRectangle`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `c.ExtendBaseWidget`

### MouseMoved (method on `*customRectangle`)
- Signature: `func (*customRectangle) MouseMoved(a *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: MouseMoved is called when a desktop pointer hovers over the widget
- Selector calls: `log.Println`, `fmt.Println`

### MouseIn (method on `*CustomButton`)
- Signature: `func (*CustomButton) MouseIn(e *desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`
- Internal calls: `float32`
- Selector calls: `fmt.Println`, `m.Position`, `fyne.NewPos`, `image.Move`, `image.Show`, `image.Refresh`

### MouseOut (method on `*CustomButton`)
- Signature: `func (*CustomButton) MouseOut()`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: fmt.Println("Mouse Moved")
- Selector calls: `fmt.Println`, `image.Hide`, `image.Refresh`

### loadUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) loadUI() fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `newAdaptiveSplit`
- Selector calls: `widget.NewLabel`, `widget.NewMultiLineEntry`, `uiServer.loadCompleteAvailableTestCaseBuildingBlocksUI`, `canvas.NewText`, `canvas.NewRectangle`, `draggingBackgroundRectangle.SetMinSize`, `draggingText.Size`, `fyne.NewSize`

### loadCompleteAvailableTestCaseBuildingBlocksUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) loadCompleteAvailableTestCaseBuildingBlocksUI() completeAvailableTestCaseBuildingBlocksUI fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if`
- Doc: Loads available TestInstructions and TestInstructionContainers and return the UI Bar and UI Tree-structure for them
- Selector calls: `widget.NewToolbar`, `widget.NewToolbarAction`, `theme.ContentRedoIcon`, `fmt.Println`, `uiServer.makeTreeUI`, `theme.DocumentSaveIcon`, `theme.ContentAddIcon`, `theme.ContentRemoveIcon`

### loadAvailableTestCaseBuildingBlocksTreeUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) loadAvailableTestCaseBuildingBlocksTreeUI() availableTestCaseBuildingBlocksTreeUI fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Loads current BuildingBlocksTree UI-structure

### loadCompleteCurrentTestCaseUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) loadCompleteCurrentTestCaseUI() completeCurrentTestCaseUIContainer fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Loads current TestCase return the UI-structure for it
- Internal calls: `newAdaptiveSplit`
- Selector calls: `widget.NewToolbar`, `widget.NewToolbarAction`, `theme.ContentRedoIcon`, `fmt.Println`, `theme.ContentCopyIcon`, `theme.ContentCutIcon`, `theme.ContentPasteIcon`, `container.NewAppTabs`

### loadCurrentTestCaseModelAreaUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) loadCurrentTestCaseModelAreaUI() currentTestCaseModelAreaUI fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Loads current TestCase testCaseModel and return the UI-structure for it
- Selector calls: `binding.NewString`, `widget.NewLabelWithData`, `container.NewVBox`

### loadCurrentTestCaseAttributesAreaUI (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) loadCurrentTestCaseAttributesAreaUI() currentTestCaseAttributesAreaUI fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Loads current TestCase attributes and return the UI-structure for it
- Selector calls: `widget.NewLabel`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
