# executionStatusUpdatesHandler_channelReader.go

## File Overview
- Path: `messageStreamEngine/executionStatusUpdatesHandler_channelReader.go`
- Package: `messageStreamEngine`
- Functions/Methods: `4`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/detailedExecutionsModel`
- `FenixTesterGui/executions/executionsModelForSubscriptions`
- `FenixTesterGui/executions/executionsUIForSubscriptions`
- `errors`
- `fmt`
- `github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `strconv`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### startCommandChannelReader (method on `*MessageStreamEngineStruct`)
- Signature: `func (*MessageStreamEngineStruct) startCommandChannelReader()`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Channel reader which is used for reading out Status messages that is sent from GuiExecutionServer
- Internal calls: `int32`
- Selector calls: `time.Sleep`, `messageStreamEngineObject.processTestExecutionStatusChange`, `messageStreamEngineObject.initiateOpenMessageStreamToGuiExecutionServer`, `messageStreamEngineObject.initiateOpenMessageStreamToGuiExecutionServerInXSeconds`

### initiateOpenMessageStreamToGuiExecutionServer (method on `*MessageStreamEngineStruct`)
- Signature: `func (*MessageStreamEngineStruct) initiateOpenMessageStreamToGuiExecutionServer()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Call Worker to get TestInstructions to Execute, which is done as a message stream in the response from the Worker
- Selector calls: `messageStreamEngineObject.initiateOpenMessageStreamToGuiExecutionServerInXSeconds`

### initiateOpenMessageStreamToGuiExecutionServerInXSeconds (method on `*MessageStreamEngineStruct`)
- Signature: `func (*MessageStreamEngineStruct) initiateOpenMessageStreamToGuiExecutionServerInXSeconds(waitTimeInSeconds uint)`
- Exported: `false`
- Control-flow features: `if, go`
- Doc: Call Worker in X seconds, due to some connection error, to get TestInstructions to Execute, which is done as a message stream in the response from the Worker
- Selector calls: `time.Duration`, `time.Sleep`, `messageStreamEngineObject.initiateGuiExecutionServerRequestForMessages`

### processTestExecutionStatusChange (method on `*MessageStreamEngineStruct`)
- Signature: `func (*MessageStreamEngineStruct) processTestExecutionStatusChange(executionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage)`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Process TestExecutionStatus-change
- Internal calls: `int`
- Selector calls: `strconv.Itoa`, `executionsUIForSubscriptions.MoveTestCaseExecutionFromUnderExecutionToFinishedExecution`, `errors.New`, `fmt.Sprintf`, `fmt.Println`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
