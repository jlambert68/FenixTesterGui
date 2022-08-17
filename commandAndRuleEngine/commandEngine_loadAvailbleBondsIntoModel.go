package commandAndRuleEngine

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// LoadAvailableBondsFromServer Load all Available Bonds from Gui-server
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) LoadAvailableBondsFromServer() {

	var availableImmatureBondsMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	availableImmatureBondsMessage = commandAndRuleEngine.GrpcOutReference.ListAllAvailableBonds("s41797") //TODO change to use current logged in to computer user

	commandAndRuleEngine.loadModelWithAvailableBonds(availableImmatureBondsMessage)

}

// Load Model with Available Bonds
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) loadModelWithAvailableBonds(availableImmatureBondsMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage) {

	var availableImmatureBondsMap map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage
	availableImmatureBondsMap = make(map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage)

	// Loop all bonds and put inte map
	for _, immantureBond := range availableImmatureBondsMessage.ImmatureBonds {
		availableImmatureBondsMap[immantureBond.BasicBondInformation.VisibleBondAttributes.TestCaseModelElementType] = immantureBond
	}

	// Save bonds-map into model for Available Bonds
	commandAndRuleEngine.availableBondsMap = availableImmatureBondsMap

	// Save Copy of Bonds in TestCase //TODO Place all Bonds and Immature TI and Immature TIC in separate object
	commandAndRuleEngine.Testcases.AvailableBondsMap = availableImmatureBondsMap

}
