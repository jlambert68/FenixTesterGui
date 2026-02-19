# availableBuildingBlockTree_load_gRPC_results_into_model_test.go

## File Overview
- Path: `gui/availableBuildingBlockTree_load_gRPC_results_into_model_test.go`
- Package: `gui`
- Functions/Methods: `6`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TestLoadModelWithAvailableBuildingBlocks`
- `TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers`
- `TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions`
- `TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionAndTestInstructionContainers`
- `TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers`
- `TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructions`

## Imports
- `FenixTesterGui/grpc_out_GuiTestCaseBuilderServer`
- `FenixTesterGui/gui/UnitTestTestData`
- `fmt`
- `github.com/golang/protobuf/jsonpb`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/sirupsen/logrus`
- `github.com/stretchr/testify/assert`
- `testing`

## Declared Types
- None

## Declared Constants
- `printValues`

## Declared Variables
- None

## Functions and Methods
### TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions
- Signature: `func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Available TestInstructions are put in Available Building Blocks-testCaseModel in a correct way
- Selector calls: `jsonpb.UnmarshalString`, `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `myLogger.WithFields`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `assert.Equal`, `fmt.Sprint`, `fmt.Println`

### TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers
- Signature: `func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Available TestInstructionContainers are put in Available Building Blocks-testCaseModel in a correct way
- Selector calls: `jsonpb.UnmarshalString`, `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `myLogger.WithFields`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers`, `assert.Equal`, `fmt.Sprint`, `fmt.Println`

### TestLoadModelWithAvailableBuildingBlocks
- Signature: `func TestLoadModelWithAvailableBuildingBlocks(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Available TestInstructions & TestInstructionContainers are put in Available Building Blocks-testCaseModel in a correct way
- Selector calls: `jsonpb.UnmarshalString`, `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `myLogger.WithFields`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `assert.Equal`, `fmt.Sprint`, `fmt.Println`

### TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructions
- Signature: `func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructions(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Pinned TestInstructions are put in Pinned Building Blocks-testCaseModel in a correct way
- Selector calls: `jsonpb.UnmarshalString`, `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `myLogger.WithFields`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions`, `assert.Equal`, `fmt.Sprint`

### TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers
- Signature: `func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Pinned TestInstructionContainers are put in Pinned Building Blocks-testCaseModel in a correct way
- Selector calls: `jsonpb.UnmarshalString`, `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `myLogger.WithFields`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers`, `assert.Equal`, `fmt.Sprint`

### TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionAndTestInstructionContainers
- Signature: `func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionAndTestInstructionContainers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Pinned TestInstruction And TestInstructionContainers are put in Pinned Building Blocks-testCaseModel in a correct way
- Selector calls: `jsonpb.UnmarshalString`, `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `myLogger.WithFields`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks`, `assert.Equal`, `fmt.Sprint`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
