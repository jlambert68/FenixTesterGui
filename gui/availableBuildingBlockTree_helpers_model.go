package gui

import (
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// Load Available Building Blocks, TestInstructions and TestInstructionContainers, from GUI-server
// And Store them in model
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	// Verify that AckNack Response is equal to AckNack = true
	if testInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id": "1c1d6645-4679-4140-8363-c3ed4c105540",
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Reset availableBuildingBlocksForUITreeNodes
	availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes = make(map[string]availableBuildingBlocksForUITreeNodesStruct)

	// Load TestInstructions
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage)

	// Load TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage)
}

// Load all available TestInstructions Building Blocks
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap = make(map[string]map[string]map[string]availableTestInstructionStruct) //make(fullDomainTestInstructionTypeTestInstructionRelationsMapType)
	var testInstructionTypeTestInstructionsRelationsMap map[string]map[string]availableTestInstructionStruct                                                      //testInstructionTypeTestInstructionsRelationsMapType
	var testInstructionMap map[string]availableTestInstructionStruct                                                                                              //testInstructionMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	uiTreeNodes := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes
	// If not created then create the map
	if len(uiTreeNodes) == 0 {
		availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes = make(map[string]availableBuildingBlocksForUITreeNodesStruct)
	}

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
	for _, testInstruction := range testInstructionsAndTestContainersMessage.ImmatureTestInstructions {

		// *** Does Domain exist in map ***
		testInstructionTypeTestInstructionsRelationsMap, existInMap = availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid]

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerType-TestInstructionContainers-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid,
				name:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode)

			_, existInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid]
			if existInMap == false {
				availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid] = tempNode

			}
		}

		// *** Does TestInstructionType exist in map ***
		testInstructionMap, existInMap = testInstructionTypeTestInstructionsRelationsMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid]

		// If TestInstructionType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap = make(map[string]availableTestInstructionStruct) // make(testInstructionMapType)
			testInstructionTypeTestInstructionsRelationsMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid] = testInstructionMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid,
				name:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode)

			_, existInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid]
			if existInMap == false {
				availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid] = tempNode

			}
		}

		// *** Does TestInstruction exist in map ***
		_, existInMap = testInstructionMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid]

		// Create the TestInstruction to be added to each leave node on UI-model for UI-Tree regarding Available Building Blocks
		tempTestInstruction := availableTestInstructionStruct{
			testInstructionNameInUITree: "",
			domainUuid:                  testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid,
			domainName:                  testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainName,
			testInstructionTypeUuid:     testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid,
			testInstructionTypeName:     testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeName,
			testInstructionUuid:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid,
			testInstructionName:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionName,
		}

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempNode := availableBuildingBlocksForUITreeNodesStruct{
			nameInUITree: "",
			uuid:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid,
			name:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionName,
		}
		// Set UI Node name in nodes
		tempNode.nameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode)
		tempTestInstruction.testInstructionNameInUITree = tempNode.nameInUITree

		// If TestInstruction doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid] = tempTestInstruction

			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid]
			if existInMap == false {
				availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid] = tempNode

			}

			//testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap
			//availableBuildingBlocksModel.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

		}
	}
}

// Load all available TestInstructionContainers Building Blocks
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap = make(map[string]map[string]map[string]availableTestInstructionContainerStruct) //make(fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMapType)
	var testInstructionContainerTypeTestInstructionContainersRelationsMap map[string]map[string]availableTestInstructionContainerStruct                                                      //testInstructionContainerTypeTestInstructionContainersRelationsMapType
	var testInstructionContainerMap map[string]availableTestInstructionContainerStruct                                                                                                       //testInstructionContainerMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	uiTreeNodes := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes
	// If not created then create the map
	if len(uiTreeNodes) == 0 {
		availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes = make(map[string]availableBuildingBlocksForUITreeNodesStruct)
	}

	// Loop all TestInstructionContainers and extract all data to be used in Available Building Block UI-tree
	for _, testInstructionContainer := range testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers {

		// *** Does Domain exist in map ***
		testInstructionContainerTypeTestInstructionContainersRelationsMap, existInMap = availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid]

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerContainerType-TestInstructionContainerContainers-map
		if existInMap == false {
			testInstructionContainerTypeTestInstructionContainersRelationsMap = make(map[string]map[string]availableTestInstructionContainerStruct) //make(testInstructionContainerTypeTestInstructionContainersRelationsMapType)
			availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid,
				name:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode)

			_, existInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid]
			if existInMap == false {
				availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid] = tempNode

			}
		}

		// *** Does TestInstructionContainerType exist in map ***
		testInstructionContainerMap, existInMap = testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid]

		// If TestInstructionContainerType doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap = make(map[string]availableTestInstructionContainerStruct) // make(testInstructionContainerMapType)
			testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid] = testInstructionContainerMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := availableBuildingBlocksForUITreeNodesStruct{
				nameInUITree: "",
				uuid:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid,
				name:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeName,
			}
			// Set UI Node name in node
			tempNode.nameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode)

			_, existInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid]
			if existInMap == false {
				availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid] = tempNode

			}
		}

		// *** Does TestInstructionContainer exist in map ***
		_, existInMap = testInstructionContainerMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid]

		// Create the TestInstructionContainer to be added to each leave node on UI-model for UI-Tree regarding Available Building Blocks
		tempTestInstructionContainer := availableTestInstructionContainerStruct{
			testInstructionContainerNameInUITree: "",
			domainUuid:                           testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid,
			domainName:                           testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainName,
			testInstructionContainerTypeUuid:     testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid,
			testInstructionContainerTypeName:     testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeName,
			testInstructionContainerUuid:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid,
			testInstructionContainerName:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerName,
		}

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempNode := availableBuildingBlocksForUITreeNodesStruct{
			nameInUITree: "",
			uuid:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid,
			name:         testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerName,
		}
		// Set UI Node name in nodes
		tempNode.nameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode)
		tempTestInstructionContainer.testInstructionContainerNameInUITree = tempNode.nameInUITree

		// If TestInstructionContainer doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = tempTestInstructionContainer

			// Add the TestInstructionContainer to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid]
			if existInMap == false {
				availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = tempNode

			}

			//testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid] = testInstructionContainerMap
			//availableBuildingBlocksModel.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

		}
	}
}
