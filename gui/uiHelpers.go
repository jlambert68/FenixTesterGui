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
	uiServer.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage)

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
	uiTreeNodes := uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes
	// If not created then create the map
	if len(uiTreeNodes) == 0 {
		uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes = make(map[string]availableBuildingBlocksForUITreeNodesStruct)
	}

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
	for _, testInstruction := range testInstructionsAndTestContainersMessage.TestInstructionMessages {

		// *** Does Domain exist in map ***
		testInstructionTypeTestInstructionsRelationsMap, existInMap = uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid]

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerType-TestInstructionContainers-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstruction.DomainUuid,
				name:         testInstruction.DomainName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = uiServer.generateUITreeName(tempNode)

			_, existInMap = uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.DomainUuid]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.DomainUuid] = tempNode

			}
		}

		// *** Does TestInstructionType exist in map ***
		testInstructionMap, existInMap = testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid]

		// If TestInstructionType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap = make(map[string]availableTestInstructionStruct) // make(testInstructionMapType)
			testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstruction.TestInstructionTypeUuid,
				name:         testInstruction.TestInstructionTypeName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = uiServer.generateUITreeName(tempNode)

			_, existInMap = uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.TestInstructionTypeUuid]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.TestInstructionTypeUuid] = tempNode

			}
		}

		// *** Does TestInstruction exist in map ***
		_, existInMap = testInstructionMap[testInstruction.TestInstructionUuid]

		// Create the TestInstruction to be added to each leave node on UI-model for UI-Tree regarding Available Building Blocks
		tempTestInstruction := availableTestInstructionStruct{
			testInstructionNameInUITree: "",
			domainUuid:                  testInstruction.DomainUuid,
			domainName:                  testInstruction.DomainName,
			testInstructionTypeUuid:     testInstruction.TestInstructionTypeUuid,
			testInstructionTypeName:     testInstruction.TestInstructionTypeName,
			testInstructionUuid:         testInstruction.TestInstructionUuid,
			testInstructionName:         testInstruction.TestInstructionName,
		}

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempNode := availableBuildingBlocksForUITreeNodesStruct{
			nameInUITree: "",
			uuid:         testInstruction.TestInstructionUuid,
			name:         testInstruction.TestInstructionName,
		}
		// Set UI Node name in nodes
		tempNode.nameInUITree = uiServer.generateUITreeName(tempNode)
		tempTestInstruction.testInstructionNameInUITree = tempNode.nameInUITree

		// If TestInstruction doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap[testInstruction.TestInstructionUuid] = tempTestInstruction

			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.TestInstructionUuid]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.TestInstructionUuid] = tempNode

			}

			//testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap
			//uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

		}
	}
}

// Load all available TestInstructionContainers Building Blocks
func (uiServer *UIServerStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage) {

	uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap = make(map[string]map[string]map[string]availableTestInstructionContainerStruct) //make(fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMapType)
	var testInstructionContainerTypeTestInstructionContainersRelationsMap map[string]map[string]availableTestInstructionContainerStruct                                                               //testInstructionContainerTypeTestInstructionContainersRelationsMapType
	var testInstructionContainerMap map[string]availableTestInstructionContainerStruct                                                                                                                //testInstructionContainerMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	uiTreeNodes := uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes
	// If not created then create the map
	if len(uiTreeNodes) == 0 {
		uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes = make(map[string]availableBuildingBlocksForUITreeNodesStruct)
	}

	// Loop all TestInstructionContainers and extract all data to be used in Available Building Block UI-tree
	for _, testInstructionContainer := range testInstructionsAndTestContainersMessage.TestInstructionContainerMessages {

		// *** Does Domain exist in map ***
		testInstructionContainerTypeTestInstructionContainersRelationsMap, existInMap = uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid]

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerContainerType-TestInstructionContainerContainers-map
		if existInMap == false {
			testInstructionContainerTypeTestInstructionContainersRelationsMap = make(map[string]map[string]availableTestInstructionContainerStruct) //make(testInstructionContainerTypeTestInstructionContainersRelationsMapType)
			uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstructionContainer.DomainUuid,
				name:         testInstructionContainer.DomainName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = uiServer.generateUITreeName(tempNode)

			_, existInMap = uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.DomainUuid]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.DomainUuid] = tempNode

			}
		}

		// *** Does TestInstructionContainerType exist in map ***
		testInstructionContainerMap, existInMap = testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid]

		// If TestInstructionContainerType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap = make(map[string]availableTestInstructionContainerStruct) // make(testInstructionContainerMapType)
			testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid] = testInstructionContainerMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstructionContainer.TestInstructionContainerTypeUuid,
				name:         testInstructionContainer.TestInstructionContainerTypeName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = uiServer.generateUITreeName(tempNode)

			_, existInMap = uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.TestInstructionContainerTypeUuid]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.TestInstructionContainerTypeUuid] = tempNode

			}
		}

		// *** Does TestInstructionContainer exist in map ***
		_, existInMap = testInstructionContainerMap[testInstructionContainer.TestInstructionContainerUuid]

		// Create the TestInstructionContainer to be added to each leave node on UI-model for UI-Tree regarding Available Building Blocks
		tempTestInstructionContainer := availableTestInstructionContainerStruct{
			testInstructionContainerNameInUITree: "",
			domainUuid:                           testInstructionContainer.DomainUuid,
			domainName:                           testInstructionContainer.DomainName,
			testInstructionContainerTypeUuid:     testInstructionContainer.TestInstructionContainerTypeUuid,
			testInstructionContainerTypeName:     testInstructionContainer.TestInstructionContainerTypeName,
			testInstructionContainerUuid:         testInstructionContainer.TestInstructionContainerUuid,
			testInstructionContainerName:         testInstructionContainer.TestInstructionContainerName,
		}

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempNode := availableBuildingBlocksForUITreeNodesStruct{
			nameInUITree: "",
			uuid:         testInstructionContainer.TestInstructionContainerUuid,
			name:         testInstructionContainer.TestInstructionContainerName,
		}
		// Set UI Node name in nodes
		tempNode.nameInUITree = uiServer.generateUITreeName(tempNode)
		tempTestInstructionContainer.testInstructionContainerNameInUITree = tempNode.nameInUITree

		// If TestInstructionContainer doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap[testInstructionContainer.TestInstructionContainerUuid] = tempTestInstructionContainer

			// Add the TestInstructionContainer to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.TestInstructionContainerUuid]
			if existInMap == false {
				uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.TestInstructionContainerUuid] = tempNode

			}

			//testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid] = testInstructionContainerMap
			//uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

		}
	}
}
