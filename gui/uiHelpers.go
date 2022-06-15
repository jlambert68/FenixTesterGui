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

	// Verify that AckNack Response is equal to AckNack = true
	if testInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		uiServer.logger.WithFields(logrus.Fields{
			"id": "1c1d6645-4679-4140-8363-c3ed4c105540",
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Load TestInstructions
	//uiServer.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage)

	// Load TestInstructionContainers
	uiServer.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage)
}

// Load all available TestInstructions Building Blocks
func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) {

	uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap = make(map[string]map[string]map[string]availableTestInstructionStruct) //make(fullDomainTestInstructionTypeTestInstructionRelationsMapType)
	var testInstructionTypeTestInstructionsRelationsMap map[string]map[string]availableTestInstructionStruct                                                               //testInstructionTypeTestInstructionsRelationsMapType
	var testInstructionMap map[string]availableTestInstructionStruct                                                                                                       //testInstructionMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
	if existInMap == false {
		uiServer.availableBuildingBlocksModel.availableDomains = make(map[string][]availableDomainStruct)
	}
	uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes = make(map[string][]availableTestInstructionTypeStruct)
	uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions = make(map[string][]availableTestInstructionStruct)

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
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

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerType-TestInstructionContainers-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = []availableDomainStruct{tempDomain}

				// Also add reference to on key-level
				uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid] = []availableDomainStruct{tempDomain}
			} else {
				_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid]
				if existInMap == false {
					// Also add reference to on key-level
					uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid] = []availableDomainStruct{tempDomain}
				}
			}
		} else {
			// If it is not first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Only add if it is not already in there
			_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid]
			if existInMap == false {
				tempDomains := uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
				tempDomains = append(tempDomains, tempDomain)
				uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = tempDomains

				// Also add reference to on key-level
				uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid] = []availableDomainStruct{tempDomain}
			}

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

// Load all available TestInstructionContainers Building Blocks
func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) {

	uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap = make(map[string]map[string]map[string]availableTestInstructionContainerStruct) //make(fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMapType)
	var testInstructionContainerTypeTestInstructionContainersRelationsMap map[string]map[string]availableTestInstructionContainerStruct                                                               //testInstructionContainerTypeTestInstructionContainersRelationsMapType
	var testInstructionContainerMap map[string]availableTestInstructionContainerStruct                                                                                                                //testInstructionContainerMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
	if existInMap == false {
		uiServer.availableBuildingBlocksModel.availableDomains = make(map[string][]availableDomainStruct)
	}
	uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes = make(map[string][]availableTestInstructionContainerTypeStruct)
	uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers = make(map[string][]availableTestInstructionContainerStruct)

	// Loop all TestInstructionContainers and extract all data to be used in Available Building Block UI-tree
	for _, testInstructionContainer := range testInstructionsAndTestContainersMessage.TestInstructionContainerMessages {

		// *** Does Domain exist in map ***
		testInstructionContainerTypeTestInstructionContainersRelationsMap, existInMap = uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
		tempDomain := availableDomainStruct{
			domainNameInUITree: "",
			domainUuid:         testInstructionContainer.DomainUuid,
			domainName:         testInstructionContainer.DomainName,
		}
		tempDomain.domainNameInUITree = uiServer.generateUITreeNameForDomain(tempDomain)

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerType-TestInstructionContainers-map
		if existInMap == false {
			testInstructionContainerTypeTestInstructionContainersRelationsMap = make(map[string]map[string]availableTestInstructionContainerStruct) //make(testInstructionContainerTypeTestInstructionContainersRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = []availableDomainStruct{tempDomain}

				// Also add reference to on key-level
				uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid] = []availableDomainStruct{tempDomain}
			} else {
				_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid]
				if existInMap == false {
					// Also add reference to on key-level
					uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid] = []availableDomainStruct{tempDomain}
				}
			}
		} else {
			// If it is not first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Only add if it is not already in there
			_, existInMap = uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid]
			if existInMap == false {
				tempDomains := uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap]
				tempDomains = append(tempDomains, tempDomain)
				uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] = tempDomains

				// Also add reference to on key-level
				uiServer.availableBuildingBlocksModel.availableDomains[tempDomain.domainUuid] = []availableDomainStruct{tempDomain}
			}

		}

		// *** Does TestInstructionContainerType exist in map ***
		testInstructionContainerMap, existInMap = testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstructionContainerType := availableTestInstructionContainerTypeStruct{
			testInstructionContainerTypeNameInUITree: "",
			domainUuid:                               testInstructionContainer.DomainUuid,
			domainName:                               testInstructionContainer.DomainName,
			testInstructionContainerTypeUuid:         testInstructionContainer.TestInstructionContainerTypeUuid,
			testInstructionContainerTypeName:         testInstructionContainer.TestInstructionContainerTypeName,
		}
		tempTestInstructionContainerType.testInstructionContainerTypeNameInUITree = uiServer.generateUITreeNameForTestInstructionContainerType(tempTestInstructionContainerType)

		// If TestInstructionContainerType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap = make(map[string]availableTestInstructionContainerStruct) // make(testInstructionContainerMapType)
			testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid] = testInstructionContainerMap

			// Add the TestInstructionContainerType to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes[testInstructionContainer.DomainUuid] = []availableTestInstructionContainerTypeStruct{tempTestInstructionContainerType}

		} else {
			// Add the TestInstructionContainerType to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructionContainerTypes := uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes[testInstructionContainer.DomainUuid]
			tempTestInstructionContainerTypes = append(tempTestInstructionContainerTypes, tempTestInstructionContainerType)
			uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes[testInstructionContainer.DomainUuid] = tempTestInstructionContainerTypes
		}

		// *** Does TestInstructionContainer exist in map ***
		_, existInMap = testInstructionContainerMap[testInstructionContainer.TestInstructionContainerUuid]

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempTestInstructionContainer := availableTestInstructionContainerStruct{
			testInstructionContainerNameInUITree: "",
			domainUuid:                           testInstructionContainer.DomainUuid,
			domainName:                           testInstructionContainer.DomainName,
			testInstructionContainerTypeUuid:     testInstructionContainer.TestInstructionContainerTypeUuid,
			testInstructionContainerTypeName:     testInstructionContainer.TestInstructionContainerTypeName,
			testInstructionContainerUuid:         testInstructionContainer.TestInstructionContainerUuid,
			testInstructionContainerName:         testInstructionContainer.TestInstructionContainerName,
		}
		tempTestInstructionContainer.testInstructionContainerNameInUITree = uiServer.generateUITreeNameForTestInstructionContainer(tempTestInstructionContainer)

		// If TestInstructionContainer doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap[testInstructionContainer.TestInstructionContainerUuid] = tempTestInstructionContainer

			// Add the TestInstructionContainer to a simpler structure to be used vid UI-tree for Available Building Blocks
			uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers[testInstructionContainer.TestInstructionContainerTypeUuid] = []availableTestInstructionContainerStruct{tempTestInstructionContainer}

		} else {
			// Add the TestInstructionContainer to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempTestInstructionContainers := uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers[testInstructionContainer.TestInstructionContainerTypeUuid]
			tempTestInstructionContainers = append(tempTestInstructionContainers, tempTestInstructionContainer)
			uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers[testInstructionContainer.TestInstructionContainerTypeUuid] = tempTestInstructionContainers
		}

		//testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid] = testInstructionContainerMap
		//uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

	}
}
