# commandEngineChannelReader.go

## File Overview
- Path: `commandAndRuleEngine/commandEngineChannelReader.go`
- Package: `commandAndRuleEngine`
- Functions/Methods: `13`
- Imports: `22`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/executions/executionsModelForSubscriptions`
- `FenixTesterGui/grpc_out_GuiExecutionServer`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/dialog`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `github.com/sirupsen/logrus`
- `image/color`
- `log`
- `strconv`
- `sync`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### startCommandChannelReader (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) startCommandChannelReader()`
- Exported: `false`
- Control-flow features: `for/range, switch`
- Doc: Channel reader which is used for reading out commands to CommandEngine
- Selector calls: `commandAndRuleEngine.channelCommandCreateNewTestCase`, `commandAndRuleEngine.channelCommandSwap`, `commandAndRuleEngine.channelCommandRemove`, `commandAndRuleEngine.channelCommandSaveTestCase`, `commandAndRuleEngine.channelCommandExecuteTestCase`, `commandAndRuleEngine.channelCommandChangeActiveTestCase`, `commandAndRuleEngine.channelCommandOpenTestCase`, `commandAndRuleEngine.channelCommandRemoveTestCaseWithOutSaving`

### channelCommandCreateNewTestCase (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandCreateNewTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Execute command 'New TestCase' received from Channel reader
- Selector calls: `commandAndRuleEngine.NewTestCaseModel`, `fmt.Println`

### channelCommandSaveTestCase (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandSaveTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Execute command 'Save TestCase' received from Channel reader
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `fyne.CurrentApp`

### channelCommandExecuteTestCase (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandExecuteTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Execute command 'Save TestCase' received from Channel reader
- Internal calls: `string`, `int`
- Selector calls: `dialog.NewConfirm`, `testDataEngine.TestDataDomainUuidType`, `testDataEngine.TestDataAreaUuidType`, `fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum`, `grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion`, `errors.New`, `fmt.Sprintf`, `fmt.Println`

### channelCommandRemove (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandRemove(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Execute command 'Remove Element' received from Channel reader
- Selector calls: `commandAndRuleEngine.DeleteElementFromTestCaseModel`, `fmt.Println`

### channelCommandChangeActiveTestCase (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandChangeActiveTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Change the active TestCase and TestCase-tab
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`

### runPopUp
- Signature: `func runPopUp(w fyne.Window, uuidChannel chan<- string) modal *widget.PopUp`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: PopUp used function 'channelCommandOpenTestCase', below. Generates the ability for user enter Uuid
- Selector calls: `widget.NewEntry`, `widget.NewButton`, `modal.Hide`, `container.New`, `layout.NewHBoxLayout`, `layout.NewSpacer`, `layout.NewVBoxLayout`, `widget.NewLabel`

### channelCommandOpenTestCase (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandOpenTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Opens a saved TestCase from Database
- Internal calls: `runPopUp`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`

### channelCommandRemoveTestCaseWithOutSaving (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandRemoveTestCaseWithOutSaving(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Remove the TestCase without saving it to theDatabase

### channelCommandCloseOpenTestCaseWithOutSaving (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandCloseOpenTestCaseWithOutSaving(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Close open TestCase without saving it to theDatabase

### printDropZone
- Signature: `func printDropZone(index int)`
- Exported: `false`
- Control-flow features: `none detected`
- Selector calls: `log.Println`

### channelCommandSwap (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) channelCommandSwap(incomingChannelCommand sharedCode.ChannelCommandStruct)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, defer`
- Doc: Execute command 'Swap Element' received from Channel reader
- Selector calls: `commandAndRuleEngine.convertGrpcElementModelIntoTestCaseElementModel`, `container.NewVBox`, `dropZoneWaitGroup.Add`, `sharedCode.ConvertRGBAHexStringIntoRGBAColor`, `canvas.NewRectangle`, `fmt.Println`, `dropZoneWaitGroup.Done`, `container.NewMax`

### convertGrpcElementModelIntoTestCaseElementModel (method on `*CommandAndRuleEngineObjectStruct`)
- Signature: `func (*CommandAndRuleEngineObjectStruct) convertGrpcElementModelIntoTestCaseElementModel(immatureGrpcElementModelMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage) immatureTestCaseElementModel testCaseModel.ImmatureElementStruct`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: Convert gRPC-message for TI or TIC into model used within the TestCase-model

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
