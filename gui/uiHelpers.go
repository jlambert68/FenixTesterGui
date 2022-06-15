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
// And Store them in model
func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) {

	//type fullDomainTestInstructionTypeTestInstructionRelationsMapType map[string]map[string]map[string]availableTestInstructionStruct  // map[#DomainUUID#]map[#TestInstructionTypeUUID#]map[#TestInstructionUUID#][]testInstructionStruct
	//type testInstructionTypeTestInstructionsRelationsMapType map[string]map[string]availableTestInstructionStruct	// map[#TestInstructionTypeUUID#]map[#TestInstructionUUID#][]testInstructionStruct
	//type testInstructionMapType map[string]availableTestInstructionStruct // map[#TestInstructionUUID#][]testInstructionStruct
	//type testInstructionType availableTestInstructionStruct

	uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap = make(map[string]map[string]map[string]availableTestInstructionStruct) //make(fullDomainTestInstructionTypeTestInstructionRelationsMapType)
	var testInstructionTypeTestInstructionsRelationsMap map[string]map[string]availableTestInstructionStruct                                                               //testInstructionTypeTestInstructionsRelationsMapType
	var testInstructionMap map[string]availableTestInstructionStruct                                                                                                       //testInstructionMapType
	//var testInstructionInMap availableTestInstructionStruct

	// Simpler structure to store Available Building Blocks for UI-tree
	uiServer.availableBuildingBlocksModel.availableDomains = make(map[string][]availableDomainStruct)
	uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes = make(map[string][]availableTestInstructionTypeStruct)
	uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes = make(map[string][]availableTestInstructionContainerTypeStruct)
	uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions = make(map[string][]availableTestInstructionStruct)
	uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers = make(map[string][]availableTestInstructionContainerStruct)

	// Verify that AckNack Response is equal to AckNack = true
	if testInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		uiServer.logger.WithFields(logrus.Fields{
			"id": "1c1d6645-4679-4140-8363-c3ed4c105540",
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
	var existInMap bool
	for _, testInstruction := range testInstructionsAndTestContainersMessage.TestInstructionMessages {

		// *** Does Domain exist in map ***
		testInstructionTypeTestInstructionsRelationsMap, existInMap = uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
		tempDomain := availableDomainStruct{
			domainNameInUITree: "",
			domainUuid:         testInstruction.DomainUuid,
			domainName:         testInstruction.DomainName,
		}
		tempDomain.domainNameInUITree = uiServer.generateUITreeNameForDomain(tempDomain)

		// If Domain doesn't exist then add it to full Domain-TestInstructionType-TestInstructions-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// First occurrence, Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = []availableDomainStruct{tempDomain}
		} else {
			// Not first occurrence, Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempDomains := uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
			tempDomains = append(tempDomains, tempDomain)
			uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = tempDomains
		}

		// *** Does TestInstructionType exist in map ***
		testInstructionMap, existInMap = testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstructionType := availableTestInstructionTypeStruct{
			testInstructionTypeNameInUITree: "",
			domainUuid:                      testInstruction.DomainUuid,
			domainName:                      testInstruction.DomainName,
			testInstructionTypeUuid:         testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:         testInstruction.TestInstructionTypeName,
		}
		tempTestInstructionType.testInstructionTypeNameInUITree = uiServer.generateUITreeNameForTestInstructionType(tempTestInstructionType)

		// If TestInstructionType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap = make(map[string]availableTestInstructionStruct) // make(testInstructionMapType)
			testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap

			// Add the TestInstructionType to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid] = []availableTestInstructionTypeStruct{tempTestInstructionType}

		} else {
			// Add the TestInstructionType to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructionTypes := uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid]
			tempTestInstructionTypes = append(tempTestInstructionTypes, tempTestInstructionType)
			uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid] = tempTestInstructionTypes
		}

		// *** Does TestInstruction exist in map ***
		_, existInMap = testInstructionMap[testInstruction.TestInstructionUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstruction := availableTestInstructionStruct{
			testInstructionNameInUITree: "",
			domainUuid:                  testInstruction.DomainUuid,
			domainName:                  testInstruction.DomainName,
			testInstructionTypeUuid:     testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:     testInstruction.TestInstructionTypeName,
			testInstructionUuid:         testInstruction.TestInstructionUuid,
			testInstructionName:         testInstruction.TestInstructionName,
		}
		tempTestInstruction.testInstructionNameInUITree = uiServer.generateUITreeNameForTestInstruction(tempTestInstruction)

		// If TestInstruction doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap[testInstruction.TestInstructionUuid] = tempTestInstruction

			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid] = []availableTestInstructionStruct{tempTestInstruction}

		} else {
			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructions := uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid]
			tempTestInstructions = append(tempTestInstructions, tempTestInstruction)
			uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid] = tempTestInstructions
		}

		//testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap
		//uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

	}
}

func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) {

	//type fullDomainTestInstructionTypeTestInstructionRelationsMapType map[string]map[string]map[string]availableTestInstructionStruct  // map[#DomainUUID#]map[#TestInstructionTypeUUID#]map[#TestInstructionUUID#][]testInstructionStruct
	//type testInstructionTypeTestInstructionsRelationsMapType map[string]map[string]availableTestInstructionStruct	// map[#TestInstructionTypeUUID#]map[#TestInstructionUUID#][]testInstructionStruct
	//type testInstructionMapType map[string]availableTestInstructionStruct // map[#TestInstructionUUID#][]testInstructionStruct
	//type testInstructionType availableTestInstructionStruct

	uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap = make(map[string]map[string]map[string]availableTestInstructionStruct) //make(fullDomainTestInstructionTypeTestInstructionRelationsMapType)
	var testInstructionTypeTestInstructionsRelationsMap map[string]map[string]availableTestInstructionStruct                                                               //testInstructionTypeTestInstructionsRelationsMapType
	var testInstructionMap map[string]availableTestInstructionStruct                                                                                                       //testInstructionMapType
	//var testInstructionInMap availableTestInstructionStruct

	// Simpler structure to store Available Building Blocks for UI-tree
	uiServer.availableBuildingBlocksModel.availableDomains = make(map[string][]availableDomainStruct)
	uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes = make(map[string][]availableTestInstructionTypeStruct)
	uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes = make(map[string][]availableTestInstructionContainerTypeStruct)
	uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions = make(map[string][]availableTestInstructionStruct)
	uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers = make(map[string][]availableTestInstructionContainerStruct)

	// Verify that AckNack Response is equal to AckNack = true
	if testInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		uiServer.logger.WithFields(logrus.Fields{
			"id": "1c1d6645-4679-4140-8363-c3ed4c105540",
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
	var existInMap bool
	for _, testInstruction := range testInstructionsAndTestContainersMessage.TestInstructionMessages {

		// *** Does Domain exist in map ***
		testInstructionTypeTestInstructionsRelationsMap, existInMap = uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
		tempDomain := availableDomainStruct{
			domainNameInUITree: "",
			domainUuid:         testInstruction.DomainUuid,
			domainName:         testInstruction.DomainName,
		}
		tempDomain.domainNameInUITree = uiServer.generateUITreeNameForDomain(tempDomain)

		// If Domain doesn't exist then add it to full Domain-TestInstructionType-TestInstructions-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// First occurrence, Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = []availableDomainStruct{tempDomain}
		} else {
			// Not first occurrence, Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempDomains := uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
			tempDomains = append(tempDomains, tempDomain)
			uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = tempDomains
		}

		// *** Does TestInstructionType exist in map ***
		testInstructionMap, existInMap = testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstructionType := availableTestInstructionTypeStruct{
			testInstructionTypeNameInUITree: "",
			domainUuid:                      testInstruction.DomainUuid,
			domainName:                      testInstruction.DomainName,
			testInstructionTypeUuid:         testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:         testInstruction.TestInstructionTypeName,
		}
		tempTestInstructionType.testInstructionTypeNameInUITree = uiServer.generateUITreeNameForTestInstructionType(tempTestInstructionType)

		// If TestInstructionType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap = make(map[string]availableTestInstructionStruct) // make(testInstructionMapType)
			testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap

			// Add the TestInstructionType to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid] = []availableTestInstructionTypeStruct{tempTestInstructionType}

		} else {
			// Add the TestInstructionType to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructionTypes := uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid]
			tempTestInstructionTypes = append(tempTestInstructionTypes, tempTestInstructionType)
			uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid] = tempTestInstructionTypes
		}

		// *** Does TestInstruction exist in map ***
		_, existInMap = testInstructionMap[testInstruction.TestInstructionUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstruction := availableTestInstructionStruct{
			testInstructionNameInUITree: "",
			domainUuid:                  testInstruction.DomainUuid,
			domainName:                  testInstruction.DomainName,
			testInstructionTypeUuid:     testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:     testInstruction.TestInstructionTypeName,
			testInstructionUuid:         testInstruction.TestInstructionUuid,
			testInstructionName:         testInstruction.TestInstructionName,
		}
		tempTestInstruction.testInstructionNameInUITree = uiServer.generateUITreeNameForTestInstruction(tempTestInstruction)

		// If TestInstruction doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap[testInstruction.TestInstructionUuid] = tempTestInstruction

			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid] = []availableTestInstructionStruct{tempTestInstruction}

		} else {
			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructions := uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid]
			tempTestInstructions = append(tempTestInstructions, tempTestInstruction)
			uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid] = tempTestInstructions
		}

		//testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap
		//uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

	}
}

func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) {

	uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap = make(map[string]map[string]map[string]availableTestInstructionStruct) //make(fullDomainTestInstructionTypeTestInstructionRelationsMapType)
	var testInstructionTypeTestInstructionsRelationsMap map[string]map[string]availableTestInstructionStruct                                                               //testInstructionTypeTestInstructionsRelationsMapType
	var testInstructionMap map[string]availableTestInstructionStruct                                                                                                       //testInstructionMapType

	// Simpler structure to store Available Building Blocks for UI-tree
	uiServer.availableBuildingBlocksModel.availableDomains = make(map[string][]availableDomainStruct)
	uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes = make(map[string][]availableTestInstructionTypeStruct)
	uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes = make(map[string][]availableTestInstructionContainerTypeStruct)
	uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions = make(map[string][]availableTestInstructionStruct)
	uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers = make(map[string][]availableTestInstructionContainerStruct)

	// Verify that AckNack Response is equal to AckNack = true
	if testInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		uiServer.logger.WithFields(logrus.Fields{
			"id": "1c1d6645-4679-4140-8363-c3ed4c105540",
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
	var existInMap bool
	for _, testInstruction := range testInstructionsAndTestContainersMessage.TestInstructionMessages {

		// *** Does Domain exist in map ***
		testInstructionTypeTestInstructionsRelationsMap, existInMap = uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
		tempDomain := availableDomainStruct{
			domainNameInUITree: "",
			domainUuid:         testInstruction.DomainUuid,
			domainName:         testInstruction.DomainName,
		}
		tempDomain.domainNameInUITree = uiServer.generateUITreeNameForDomain(tempDomain)

		// If Domain doesn't exist then add it to full Domain-TestInstructionType-TestInstructions-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// First occurrence, Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = []availableDomainStruct{tempDomain}
		} else {
			// Not first occurrence, Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempDomains := uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
			tempDomains = append(tempDomains, tempDomain)
			uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = tempDomains
		}

		// *** Does TestInstructionType exist in map ***
		testInstructionMap, existInMap = testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstructionType := availableTestInstructionTypeStruct{
			testInstructionTypeNameInUITree: "",
			domainUuid:                      testInstruction.DomainUuid,
			domainName:                      testInstruction.DomainName,
			testInstructionTypeUuid:         testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:         testInstruction.TestInstructionTypeName,
		}
		tempTestInstructionType.testInstructionTypeNameInUITree = uiServer.generateUITreeNameForTestInstructionType(tempTestInstructionType)

		// If TestInstructionType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap = make(map[string]availableTestInstructionStruct) // make(testInstructionMapType)
			testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap

			// Add the TestInstructionType to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid] = []availableTestInstructionTypeStruct{tempTestInstructionType}

		} else {
			// Add the TestInstructionType to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructionTypes := uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid]
			tempTestInstructionTypes = append(tempTestInstructionTypes, tempTestInstructionType)
			uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[testInstruction.DomainUuid] = tempTestInstructionTypes
		}

		// *** Does TestInstruction exist in map ***
		_, existInMap = testInstructionMap[testInstruction.TestInstructionUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstruction := availableTestInstructionStruct{
			testInstructionNameInUITree: "",
			domainUuid:                  testInstruction.DomainUuid,
			domainName:                  testInstruction.DomainName,
			testInstructionTypeUuid:     testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:     testInstruction.TestInstructionTypeName,
			testInstructionUuid:         testInstruction.TestInstructionUuid,
			testInstructionName:         testInstruction.TestInstructionName,
		}
		tempTestInstruction.testInstructionNameInUITree = uiServer.generateUITreeNameForTestInstruction(tempTestInstruction)

		// If TestInstruction doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap[testInstruction.TestInstructionUuid] = tempTestInstruction

			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid] = []availableTestInstructionStruct{tempTestInstruction}

		} else {
			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructions := uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid]
			tempTestInstructions = append(tempTestInstructions, tempTestInstruction)
			uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstruction.TestInstructionTypeUuid] = tempTestInstructions
		}

		//testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap
		//uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

	}
