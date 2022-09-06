package testCaseModel

import (
	"fmt"
	"github.com/getlantern/errors"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// Load Model with Available Immature TestInstruction Attributes
func (testCaseModel *TestCasesModelsStruct) LoadModelWithImmatureTestInstructionAttributes() {

	var immatureTestInstructionAttributesMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage

	immatureTestInstructionAttributesMessage = testCaseModel.GrpcOutReference.ListAllTestInstructionAttributes("s41797") //TODO change to use current logged in to computer user

	if immatureTestInstructionAttributesMessage.AckNackResponse.AckNack == false {

		errorId := "7842e5ac-aaea-415f-92ea-fe5d998a4fee"
		err := errors.New(fmt.Sprintf("Response: '%s' [ErrorID: %s]", immatureTestInstructionAttributesMessage.AckNackResponse.Comments, errorId))

		fmt.Println(err.Error()) //TODO send to error channel
	}

	var immatureTestInstructionAttributesMap map[string]map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage
	immatureTestInstructionAttributesMap = make(map[string]map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage)

	// Loop all Attributes and put inte map
	for _, attribute := range immatureTestInstructionAttributesMessage.TestInstructionAttributesList {

		// Create a map for the Attribute useing the following map-format:  map[TestInstructionUuid]map[TestInstructionAttributeUuid]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage
		testInstructionUuidMap, existInMap := immatureTestInstructionAttributesMap[attribute.TestInstructionUuid]
		if existInMap == true {
			testInstructionUuidMap[attribute.TestInstructionAttributeUuid] = attribute
			immatureTestInstructionAttributesMap[attribute.TestInstructionUuid] = testInstructionUuidMap
		} else {
			testInstructionUuidMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage)
			testInstructionUuidMap[attribute.TestInstructionAttributeUuid] = attribute
			immatureTestInstructionAttributesMap[attribute.TestInstructionUuid] = testInstructionUuidMap
		}

	}

	// Save Attributes in TestCase-model //TODO Place Attributes, Bonds and Immature TI and Immature TIC in separate object
	testCaseModel.ImmatureTestInstructionAttributesMap = immatureTestInstructionAttributesMap

}
