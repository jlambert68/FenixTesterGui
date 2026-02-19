# availableBuildingBlockTree_load_gRPC_results_into_model_test.go

## File Overview
- Path: `gui/availableBuildingBlockTree_load_gRPC_results_into_model_test.go`
- Package: `gui`
- Generated: `2026-02-19T14:23:17+01:00`
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
### TestLoadModelWithAvailableBuildingBlocks
- Signature: `func TestLoadModelWithAvailableBuildingBlocks(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Available TestInstructions & TestInstructionContainers are put in Available Building Blocks-testCaseModel in a correct way
- External calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `fmt.Println`, `fmt.Sprint`, `jsonpb.UnmarshalString`, `myLogger.WithFields`

### TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers
- Signature: `func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Available TestInstructionContainers are put in Available Building Blocks-testCaseModel in a correct way
- External calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers`, `fmt.Println`, `fmt.Sprint`, `jsonpb.UnmarshalString`, `myLogger.WithFields`

### TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions
- Signature: `func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Available TestInstructions are put in Available Building Blocks-testCaseModel in a correct way
- External calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `fmt.Println`, `fmt.Sprint`, `jsonpb.UnmarshalString`, `myLogger.WithFields`

### TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionAndTestInstructionContainers
- Signature: `func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionAndTestInstructionContainers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Pinned TestInstruction And TestInstructionContainers are put in Pinned Building Blocks-testCaseModel in a correct way
- External calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks`, `fmt.Sprint`, `jsonpb.UnmarshalString`, `myLogger.WithFields`

### TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers
- Signature: `func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Pinned TestInstructionContainers are put in Pinned Building Blocks-testCaseModel in a correct way
- External calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers`, `fmt.Sprint`, `jsonpb.UnmarshalString`, `myLogger.WithFields`

### TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructions
- Signature: `func TestLoadModelWithPinnedBuildingBlocksRegardingTestInstructions(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that Pinned TestInstructions are put in Pinned Building Blocks-testCaseModel in a correct way
- External calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions`, `fmt.Sprint`, `jsonpb.UnmarshalString`, `myLogger.WithFields`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
