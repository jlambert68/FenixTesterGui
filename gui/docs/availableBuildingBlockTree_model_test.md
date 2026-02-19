# availableBuildingBlockTree_model_test.go

## File Overview
- Path: `gui/availableBuildingBlockTree_model_test.go`
- Package: `gui`
- Functions/Methods: `5`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TestThatAlreadyPinnedBuildingBlockNotCanBePinned`
- `TestThatNonExistingBuildBlockCanBePinned`
- `TestThatNonExistingBuildingBlockCanNotBeUnPinned`
- `TestToPinBuildingBlockCanBePinned`
- `TestUnPinBuildingBlockThatIsPinned`

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
- None

## Declared Variables
- None

## Functions and Methods
### TestThatAlreadyPinnedBuildingBlockNotCanBePinned
- Signature: `func TestThatAlreadyPinnedBuildingBlockNotCanBePinned(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that an already pinned Building Block can be pinned
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.getPinnedBuildingBlocksTreeNamesFromModel`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks`, `availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer`, `fmt.Sprint`

### TestThatNonExistingBuildBlockCanBePinned
- Signature: `func TestThatNonExistingBuildBlockCanBePinned(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that a non-existing Building Block can't be pinned
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer`, `fmt.Sprint`, `jsonpb.UnmarshalString`

### TestThatNonExistingBuildingBlockCanNotBeUnPinned
- Signature: `func TestThatNonExistingBuildingBlockCanNotBeUnPinned(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that a non-existing Building Block can't be unpinned
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions`, `availableBuildingBlocksModel.unPinTestInstructionOrTestInstructionContainer`, `fmt.Sprint`, `jsonpb.UnmarshalString`

### TestToPinBuildingBlockCanBePinned
- Signature: `func TestToPinBuildingBlockCanBePinned(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that a non-pinned Building Block can be pinned
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks`, `availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer`, `fmt.Sprint`, `jsonpb.UnmarshalString`

### TestUnPinBuildingBlockThatIsPinned
- Signature: `func TestUnPinBuildingBlockThatIsPinned(t *testing.T)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Checks that an pinned Building Block exist among pinned building blocks func TestThatBuildingBlockToBeUnPinnedExistsAmongPinned(t *testing.T) {
- Selector calls: `UnitTestTestData.InitLoggerForTest`, `UnitTestTestData.IsTestDataUsingCorrectTestDataProtoFileVersion`, `assert.Equal`, `availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks`, `availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks`, `availableBuildingBlocksModel.unPinTestInstructionOrTestInstructionContainer`, `fmt.Sprint`, `jsonpb.UnmarshalString`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
