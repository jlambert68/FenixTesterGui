package gui

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (uiServer *UIServerStruct) SetLogger(logger *logrus.Logger) {

	//myUIServer = UIServerStruct{}
	uiServer.logger = logger

	return

}

// SetDialAddressString
// Set the Dial Address, which was received from environment variables
func (uiServer *UIServerStruct) SetDialAddressString(dialAddress string) {
	uiServer.fenixGuiBuilderServerAddressToDial = dialAddress

	return

}

// Load Available Building Blocks, TestInstructions and TestInstructionContainers, from GUI-server
// Store them in model
func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) (err error) {

	var availableDomains []availableDomainStruct
	var availableTestInstructionTypes []availableTestInstructionTypeStruct
	var availableTestInstructionContainerTypes []availableTestInstructionContainerTypeStruct
	var availableTestInstructions []availableTestInstructionStruct
	var availableTestInstructionContainers []availableTestInstructionContainerStruct

	availableDomainsMap := make(map[string]availableDomainStruct)
	domainsTestInstructionTypesMap := make(map[string][]availableTestInstructionTypeStruct)                                  // map[#DomainUUID#][]availableTestInstructionTypeStruct
	domainsTestInstructionContainerTypesMap := make(map[string][]availableTestInstructionContainerTypeStruct)                // map[#DomainUUID#][]availableTestInstructionContainerTypeStruct
	testInstructionTypesTestInstructionsMap := make(map[string][]availableTestInstructionStruct)                             // map[#TestInstructionTypeUUID#][]availableTestInstructionStruct
	testInstructionContainerTypesTestInstructionsContainersMap := make(map[string][]availableTestInstructionContainerStruct) // map[#TestInstructionContainerTypeUUID#][]availableTestInstructionContainerTypeStruct

	// Loop all TestInstructions and extract Domains
	var existInMap bool
	var tempDomains []availableDomainStruct
	for _, testInstruction := range testInstructionsAndTestContainersMessage.TestInstructionMessages {
		var tempDomain availableDomainStruct
		_, existInMap = availableDomainsMap[testInstruction.DomainUuid]
		// If Domain doesn't exist then add it to Domain-map
		if existInMap == false {
			tempDomain = availableDomainStruct{
				domainNameInUITree:"",
				domainUuid:         testInstruction.DomainUuid,
				domainName:         testInstruction.DomainName,
			}
			tempDomain.domainNameInUITree = uiServer.generateUITreeNameForDomain(tempDomain)
			tempDomains = append(tempDomains, tempDomain)
		}
	}

	// Loop all TestInstructionContainers and extract Domains
	for _, testInstructionContainer := range testInstructionsAndTestContainersMessage.TestInstructionContainerMessages {
		var tempDomain availableDomainStruct
		_, existInMap = availableDomainsMap[testInstructionContainer.DomainUuid]
		// If Domain doesn't exist then add it to Domain-map
		if existInMap == false {
			tempDomain = availableDomainStruct{
				domainNameInUITree: "",
				domainUuid:         testInstructionContainer.DomainUuid,
				domainName:         testInstructionContainer.DomainName,
			}
			tempDomain.domainNameInUITree = uiServer.generateUITreeNameForDomain(tempDomain)
			tempDomains = append(tempDomains, tempDomain)
		}
	}

	// Add all Domains to Building Block model
	uiServer.availableBuildingBlocksModel.availableDomains = tempDomains


	// Loop all TestInstructionContainers and extract Domains



	uiServer.availableBuildingBlocksModel.availableDomains =
		uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes =
	uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes =
		uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions =
	uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers =

	return err
}
