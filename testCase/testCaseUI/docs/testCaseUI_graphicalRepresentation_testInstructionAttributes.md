# testCaseUI_graphicalRepresentation_testInstructionAttributes.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_graphicalRepresentation_testInstructionAttributes.go`
- Package: `testCaseUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `6`
- Imports: `16`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0`
- `image/color`
- `log`
- `regexp`
- `sort`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### extractResponseVariablesFromTestInstruction (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) extractResponseVariablesFromTestInstruction(allowedResponseVariablesTypeUuidPtr *[]string, currentElement *testCaseModel.MatureTestCaseModelElementStruct, testInstructionWithCorrectResponseVariablesTypePtr *[]*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct) err error`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Extract ResponseVariables from TestInstruction, used within the two traverse functions
- External calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`, `tempImmatureTestInstructionMessage.GetResponseVariablesMapStructure`, `tempResponseVariable.GetResponseVariableTypeUuid`

### generateAttributeRow (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) generateAttributeRow(currentTestCaseUuid string, attributeItem *testCaseModel.AttributeStruct, attributesList *testCaseModel.AttributeStructSliceReferenceType, attributesFormContainer *fyne.Container, currentTestCase *testCaseModel.TestCaseModelStruct, testInstructionElementMatureUuid string, immatureTestInstructionUuid string)`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Generate and add an 'attribute row' to be used in attributes
- Internal calls: `NewCustomAttributeSelectComboBoxWidget`, `int32`, `newCustomAttributeEntryWidget`, `string`
- External calls: `attributesFormContainer.Add`, `attributesFormContainer.Refresh`, `canvas.NewRectangle`, `err.Error`, `errors.New`, `fmt.Println`, `fmt.Sprintf`, `immatureTestInstruction.GetResponseVariablesMapStructure`

### generateAttributeStringListData (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) generateAttributeStringListData(testCaseUuid string, testInstructionElementMatureUuid string) (attributesListRef testCaseModel.AttributeStructSliceReferenceType, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Generate structure for 'binding.StringList' regarding Attribute values
- Internal calls: `int32`
- External calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`, `regexp.MustCompile`, `sort.SliceStable`

### generateTestCaseAttributesAreaForTestCase (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) generateTestCaseAttributesAreaForTestCase(testCaseUuid string, testInstructionElementMatureUuid string) (testCaseAttributesArea fyne.CanvasObject, testInstructionAttributesAccordion *widget.Accordion, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generate the TestCaseAttributes Area for the TestCase
- External calls: `attributesContainer.Add`, `attributesFormContainer.Add`, `container.New`, `container.NewScroll`, `container.NewVBox`, `errors.New`, `fmt.Println`, `fmt.Sprintf`

### recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch(currentTestCase *testCaseModel.TestCaseModelStruct, elementUuidToCheck string, allowedResponseVariablesTypeUuidPtr *[]string, testInstructionWithCorrectResponseVariablesTypePtr *[]*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct) err error`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: When the traverse logic comes to a TestInstructionContainer it will start to traverse down in this path, going first down and then right. This is use in UI for user to chose from. Then this information is used in runtime to get the value th...
- External calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`, `log.Fatalln`, `testCasesUiCanvasObject.extractResponseVariablesFromTestInstruction`, `testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch`

### recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch(currentTestCase *testCaseModel.TestCaseModelStruct, elementUuidToCheck string, allowedResponseVariablesTypeUuidPtr *[]string, testInstructionWithCorrectResponseVariablesTypePtr *[]*testCaseModel.MatureTestInstructionWithCorrectResponseVariablesTypeStruct, thisIsTheStartElement bool, previousProcessedElementUuid string, previousProcessedElementsParentUuid string) err error`
- Exported: `false`
- Control-flow features: `if, switch, returns error`
- Doc: Traverse the element model to left and/or upward to the top element. This to be able to find all matching response variables. When the traverse logic comes to a TestInstructionContainer it will start to traverse down in this path, going fir...
- External calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`, `log.Fatalln`, `testCasesUiCanvasObject.extractResponseVariablesFromTestInstruction`, `testCasesUiCanvasObject.recursiveTraverseDownwardsTestInstructionContainerElementsForResponseVariablesThatMatch`, `testCasesUiCanvasObject.recursiveTraverseUpwardsTestInstructionContainerElementsForResponseVariablesThatMatch`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
