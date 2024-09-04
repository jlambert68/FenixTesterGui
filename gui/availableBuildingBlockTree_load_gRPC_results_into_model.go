package gui

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"github.com/sirupsen/logrus"
)

// Load Available Building Blocks, TestInstructions and TestInstructionContainers, from GUI-server into testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocks(
	testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	// Verify that AckNack Response is equal to AckNack = true
	if testInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id": "1c1d6645-4679-4140-8363-c3ed4c105540",
			"testInstructionsAndTestContainersMessage.AckNackResponse.Comments": testInstructionsAndTestContainersMessage.AckNackResponse.Comments,
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Reset AvailableBuildingBlocksForUITreeNodes
	availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes = make(map[string]AvailableBuildingBlocksForUITreeNodesStruct)

	// Reset 'pinnedBuildingBlocksForUITreeNodes'
	availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load TestInstructions
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage)

	// Load TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage)

	// Store the full available Building Blocks Structure
	availableBuildingBlocksModel.storeFullGrpcStructureForAvailableBuildingBlocks(testInstructionsAndTestContainersMessage)

	// Store list with Domains that can own a TestCase
	availableBuildingBlocksModel.storeDomainsThatCanOwnTestCases(testInstructionsAndTestContainersMessage, testCaseModeReference)

}

// Load Pinned Building Blocks, TestInstructions and TestInstructionContainers, from GUI-server into testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadModelWithPinnedBuildingBlocks(pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	// Verify that AckNack Response is equal to AckNack = true
	if pinnedTestInstructionsAndTestContainersMessage.AckNackResponse.AckNack == false {
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id": "707eb9c8-fd65-466e-aa82-1eca23b40345",
		}).Fatalln("Code should not come here if AckNack == false")
	}

	// Reset 'pinnedBuildingBlocksForUITreeNodes'
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes = make(map[string]uiTreeNodesNameToUuidStruct)

	// Load relations between tree-name and original UUID for TestInstructions
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructions(pinnedTestInstructionsAndTestContainersMessage)

	// Load relations between tree-name and original UUID for TestInstructionContainers
	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(pinnedTestInstructionsAndTestContainersMessage)
}

// Load all available TestInstructions Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructions(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap = make(map[string]map[string]map[string]availableTestInstructionStruct) //make(fullDomainTestInstructionTypeTestInstructionRelationsMapType)
	var testInstructionTypeTestInstructionsRelationsMap map[string]map[string]availableTestInstructionStruct                                                      //testInstructionTypeTestInstructionsRelationsMapType
	var testInstructionMap map[string]availableTestInstructionStruct                                                                                              //testInstructionMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	uiTreeNodes := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes
	// If not created then create the map
	if len(uiTreeNodes) == 0 {
		availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes = make(map[string]AvailableBuildingBlocksForUITreeNodesStruct)
	}

	// Loop all TestInstructions and extract all data to be used in Available Building Block UI-tree
	for _, testInstruction := range testInstructionsAndTestContainersMessage.ImmatureTestInstructions {

		// only show TestInstruction that are enabled
		if testInstruction.BasicTestInstructionInformation.GetInvisibleBasicInformation().Enabled == false {
			continue
		}

		// *** Does Domain exist in map ***
		testInstructionTypeTestInstructionsRelationsMap, existInMap = availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid]

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerType-TestInstructionContainers-map
		if existInMap == false {
			testInstructionTypeTestInstructionsRelationsMap = make(map[string]map[string]availableTestInstructionStruct) //make(testInstructionTypeTestInstructionsRelationsMapType)
			availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := AvailableBuildingBlocksForUITreeNodesStruct{
				nameInUITree:      "",
				uuid:              testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid,
				name:              testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainName,
				BuildingBlockType: Undefined,
			}
			// Set UI Node name in node
			tempNode.nameInUITree, tempNode.pinnedNameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode, "Can not be pinned")

			_, existInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid]
			if existInMap == false {
				availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid] = tempNode

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
			tempNode := AvailableBuildingBlocksForUITreeNodesStruct{
				nameInUITree:      "",
				uuid:              testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid,
				name:              testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeName,
				BuildingBlockType: Undefined,
			}
			// Set UI Node name in node
			tempNode.nameInUITree, tempNode.pinnedNameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode, "Can not be pinned")

			_, existInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid]
			if existInMap == false {
				availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid] = tempNode

			}
		}

		// *** Does TestInstruction exist in map ***
		_, existInMap = testInstructionMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid]

		// Create the TestInstruction to be added to each leave node on UI-testCaseModel for UI-Tree regarding Available Building Blocks
		tempTestInstruction := availableTestInstructionStruct{
			testInstructionNameInUITree: "",
			domainUuid:                  testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainUuid,
			domainName:                  testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainName,
			testInstructionTypeUuid:     testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeUuid,
			testInstructionTypeName:     testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionTypeName,
			testInstructionUuid:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid,
			testInstructionName:         testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalName,
		}

		// Create simpler structure to be used vid UI-tree for Available Building Blocks
		tempNode := AvailableBuildingBlocksForUITreeNodesStruct{
			nameInUITree:      "",
			uuid:              testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid,
			name:              testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalName,
			BuildingBlockType: TestInstruction,
		}
		// Set UI Node name in nodes
		tempNode.nameInUITree, tempNode.pinnedNameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode, testInstruction.BasicTestInstructionInformation.NonEditableInformation.DomainName)
		tempTestInstruction.testInstructionNameInUITree = tempNode.nameInUITree

		// Add TestInstruction to TreeName to UUID -map
		tempTreeNameToUuidInstance := uiTreeNodesNameToUuidStruct{
			uuid:              tempNode.uuid,
			buildingBlockType: TestInstruction,
		}

		availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[tempNode.nameInUITree] = tempTreeNameToUuidInstance

		// If TestInstruction doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionMap[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid] = tempTestInstruction

			// Add the TestInstruction to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid]
			if existInMap == false {
				availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid] = tempNode

			}

			//testInstructionTypeTestInstructionsRelationsMap[testInstruction.TestInstructionTypeUuid] = testInstructionMap
			//AvailableBuildingBlocksModel.AvailableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[testInstruction.DomainUuid] = testInstructionTypeTestInstructionsRelationsMap

		}
	}
}

// Load all available TestInstructionContainers Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadModelWithAvailableBuildingBlocksRegardingTestInstructionContainers(testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap = make(map[string]map[string]map[string]availableTestInstructionContainerStruct) //make(fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMapType)
	var testInstructionContainerTypeTestInstructionContainersRelationsMap map[string]map[string]availableTestInstructionContainerStruct                                                      //testInstructionContainerTypeTestInstructionContainersRelationsMapType
	var testInstructionContainerMap map[string]availableTestInstructionContainerStruct                                                                                                       //testInstructionContainerMapType

	var existInMap bool

	// Simpler structure to store Available Building Blocks for UI-tree
	uiTreeNodes := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes
	// If not created then create the map
	if len(uiTreeNodes) == 0 {
		availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes = make(map[string]AvailableBuildingBlocksForUITreeNodesStruct)
	}

	// Loop all TestInstructionContainers and extract all data to be used in Available Building Block UI-tree
	for _, testInstructionContainer := range testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers {

		// only use TestInstructionContainer that are enabled
		if testInstructionContainer.BasicTestInstructionContainerInformation.GetInvisibleBasicInformation().Enabled == false {
			continue
		}

		// *** Does Domain exist in map ***
		testInstructionContainerTypeTestInstructionContainersRelationsMap, existInMap = availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid]

		// If Domain doesn't exist then add it to full Domain-TestInstructionContainerContainerType-TestInstructionContainerContainers-map
		if existInMap == false {
			testInstructionContainerTypeTestInstructionContainersRelationsMap = make(map[string]map[string]availableTestInstructionContainerStruct) //make(testInstructionContainerTypeTestInstructionContainersRelationsMapType)
			availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

			// If it is first occurrence in simpler structure then; Add the Domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			// Create simpler structure to be used vid UI-tree for Available Building Blocks
			// Add the domain to a simpler structure to be used vid UI-tree for Available Building Blocks
			tempNode := AvailableBuildingBlocksForUITreeNodesStruct{
				nameInUITree:      "",
				uuid:              testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid,
				name:              testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainName,
				BuildingBlockType: Undefined,
			}
			// Set UI Node name in node
			tempNode.nameInUITree, tempNode.pinnedNameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode, "Can not be pinned")

			_, existInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid]
			if existInMap == false {
				availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainUuid] = tempNode

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
			tempNode := AvailableBuildingBlocksForUITreeNodesStruct{
				nameInUITree:      "",
				uuid:              testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid,
				name:              testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeName,
				BuildingBlockType: Undefined,
			}
			// Set UI Node name in node
			tempNode.nameInUITree, tempNode.pinnedNameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode, "Can not be pinned")

			_, existInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid]
			if existInMap == false {
				availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerTypeUuid] = tempNode

			}
		}

		// *** Does TestInstructionContainer exist in map ***
		_, existInMap = testInstructionContainerMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid]

		// Create the TestInstructionContainer to be added to each leave node on UI-testCaseModel for UI-Tree regarding Available Building Blocks
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
		tempNode := AvailableBuildingBlocksForUITreeNodesStruct{
			nameInUITree:      "",
			uuid:              testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid,
			name:              testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerName,
			BuildingBlockType: TestInstructionContainer,
		}
		// Set UI Node name in nodes
		if len(tempNode.uuid) == 0 {
			fmt.Println("Debiug")
		}
		tempNode.nameInUITree, tempNode.pinnedNameInUITree = availableBuildingBlocksModel.generateUITreeName(tempNode, testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.DomainName)
		tempTestInstructionContainer.testInstructionContainerNameInUITree = tempNode.nameInUITree

		// Add TestInstructionContainer to TreeName to UUID -map
		tempTreeNameToUuidInstance := uiTreeNodesNameToUuidStruct{
			uuid:              tempNode.uuid,
			buildingBlockType: TestInstructionContainer,
		}

		availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[tempNode.nameInUITree] = tempTreeNameToUuidInstance

		// If TestInstructionContainer doesn't exist then add it with its full map-structure
		if existInMap == false {
			testInstructionContainerMap[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = tempTestInstructionContainer

			// Add the TestInstructionContainer to a simpler structure to be used vid UI-tree for Available Building Blocks
			_, existInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid]
			if existInMap == false {
				availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = tempNode

			}

			//testInstructionContainerTypeTestInstructionContainersRelationsMap[testInstructionContainer.TestInstructionContainerTypeUuid] = testInstructionContainerMap
			//AvailableBuildingBlocksModel.AvailableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[testInstructionContainer.DomainUuid] = testInstructionContainerTypeTestInstructionContainersRelationsMap

		}
	}
}

// Load all Pinned TestInstructions Building Blocks into testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadModelWithPinnedBuildingBlocksRegardingTestInstructions(pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	// If there are no Pinned TestInstructions then exit this function
	if len(pinnedTestInstructionsAndTestContainersMessage.AvailablePinnedTestInstructions) == 0 {
		return
	}

	// Loop through the Pinned TestInstructions and add them to the testCaseModel
	for _, pinnedTestInstruction := range pinnedTestInstructionsAndTestContainersMessage.AvailablePinnedTestInstructions {

		// Get Tree-name from testCaseModel and create reference between Pinned Tree-name and original UUID
		availableTestInstructionFromTreeNameModel, existsInMap := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[pinnedTestInstruction.TestInstructionUuid]

		// If the element is missing in map then there are something wrong
		if existsInMap == false {
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":                    "fa3890ab-d756-4deb-b66c-5357167904e5",
				"pinnedTestInstruction": pinnedTestInstruction,
			}).Fatalln("Some is wrong because couldn't find the 'pinnedTestInstruction' in tree-view-name-testCaseModel")
		}

		tempTreeNameToUuidForPinnedInstruction := uiTreeNodesNameToUuidStruct{
			uuid:              pinnedTestInstruction.TestInstructionUuid,
			buildingBlockType: TestInstruction,
		}

		// Add relation between pinned name and the original elements UUID
		availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[availableTestInstructionFromTreeNameModel.pinnedNameInUITree] = tempTreeNameToUuidForPinnedInstruction
	}

}

// Load all Pinned TestInstructions Building Blocks into testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadModelWithPinnedBuildingBlocksRegardingTestInstructionContainers(pinnedTestInstructionsAndTestInstructionsContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	// If there are no Pinned TestInstructions then exit this function
	if len(pinnedTestInstructionsAndTestInstructionsContainersMessage.AvailablePinnedPreCreatedTestInstructionContainers) == 0 {
		return
	}

	// Loop through the Pinned TestInstructionContainerss and add them to the testCaseModel
	for _, pinnedTestInstructionContainer := range pinnedTestInstructionsAndTestInstructionsContainersMessage.AvailablePinnedPreCreatedTestInstructionContainers {

		// Get Tree-name from testCaseModel and create reference between Pinned Tree-name and original UUID
		availableTestInstructionContainerFromTreeNameModel, existsInMap := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[pinnedTestInstructionContainer.TestInstructionContainerUuid]

		// If the element is missing in map then there are something wrong
		if existsInMap == false {
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":                             "6f2781ef-4bf8-40c1-8da4-0ac6996ae04b",
				"pinnedTestInstructionContainer": pinnedTestInstructionContainer,
			}).Fatalln("Some is wrong because couldn't find the 'pinnedTestInstructionContainer' in tree-view-name-testCaseModel")
		}

		tempTreeNameToUuidForPinnedInstruction := uiTreeNodesNameToUuidStruct{
			uuid:              pinnedTestInstructionContainer.TestInstructionContainerUuid,
			buildingBlockType: TestInstructionContainer,
		}

		// Add relation between pinned name and the original elements UUID
		availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[availableTestInstructionContainerFromTreeNameModel.pinnedNameInUITree] = tempTreeNameToUuidForPinnedInstruction
	}

}

// Store the full available Building Blocks Structure into the Available Building Blocks Model
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) storeFullGrpcStructureForAvailableBuildingBlocks(
	testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.
		AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	// Initiate TI-BuildingBlockType-map and TIC-BuildingBlockmap
	availableBuildingBlocksModel.allImmatureTestInstructionsBuildingBlocks = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage)
	availableBuildingBlocksModel.allImmatureTestInstructionContainerBuildingBlocks = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage)

	// If there are no available TI or TIC then exit
	if len(testInstructionsAndTestContainersMessage.ImmatureTestInstructions) == 0 &&
		len(testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers) == 0 {
		return
	}

	// Loop all TestInstruction-Building Blocks and add to model
	for _, testInstructionBuildingBlock := range testInstructionsAndTestContainersMessage.ImmatureTestInstructions {
		availableBuildingBlocksModel.allImmatureTestInstructionsBuildingBlocks[testInstructionBuildingBlock.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalUuid] = testInstructionBuildingBlock
	}

	// Loop all TestInstructionContainer-Building Blocks and add to model
	for _, testInstructionContainerBuildingBlock := range testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers {
		availableBuildingBlocksModel.allImmatureTestInstructionContainerBuildingBlocks[testInstructionContainerBuildingBlock.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = testInstructionContainerBuildingBlock
	}

}

// Store list with Domains that can own a TestCase
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) storeDomainsThatCanOwnTestCases(
	testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.
		AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	// Store the list with Domains that can own a TestCase
	availableBuildingBlocksModel.DomainsThatCanOwnTheTestCase = testInstructionsAndTestContainersMessage.DomainsThatCanOwnTheTestCase

	// Store the Domains in the TestCaseModel
	testCaseModeReference.DomainsThatCanOwnTheTestCaseMap = make(map[string]*testCaseModel.DomainThatCanOwnTheTestCaseStruct)

	var tempDomainNameShownInGui string

	// Store the Available Owner Domains as a map structure in TestCase-struct
	for _, tempDomainThatCanOwnTheTestCase := range testInstructionsAndTestContainersMessage.DomainsThatCanOwnTheTestCase {

		// Create Key and Domain Name to show in DropDown in Gui
		tempDomainNameShownInGui = testCaseModeReference.GenerateShortUuidFromFullUuid(tempDomainThatCanOwnTheTestCase.GetDomainUuid())
		tempDomainNameShownInGui = tempDomainThatCanOwnTheTestCase.GetDomainName() + " [" + tempDomainNameShownInGui + "]"

		var tempDomainsThatCanOwnTheTestCase *testCaseModel.DomainThatCanOwnTheTestCaseStruct
		tempDomainsThatCanOwnTheTestCase = &testCaseModel.DomainThatCanOwnTheTestCaseStruct{
			DomainUuid:           tempDomainThatCanOwnTheTestCase.GetDomainUuid(),
			DomainName:           tempDomainThatCanOwnTheTestCase.GetDomainName(),
			DomainNameShownInGui: tempDomainNameShownInGui,
		}
		testCaseModeReference.DomainsThatCanOwnTheTestCaseMap[tempDomainNameShownInGui] = tempDomainsThatCanOwnTheTestCase
	}

}

// Store list with TemplateRepositoryApiUrls
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) storeTemplateRepositoryApiUrls(
	templateRepositoryApiUrlsToBeStored []*fenixGuiTestCaseBuilderServerGrpcApi.RepositoryApiUrlResponseMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	// Store the list with TemplateRepositoryApiUrls
	availableBuildingBlocksModel.TemplateRepositoryApiUrls = templateRepositoryApiUrlsToBeStored

	// Store a pointer to 'TemplateRepositoryApiUrls'
	sharedCode.TemplateRepositoryApiUrlsPtr = &availableBuildingBlocksModel.TemplateRepositoryApiUrls

	// Store the TemplateRepositoryApiUrls-list in the TestCaseModel
	testCaseModeReference.TemplateRepositoryApiUrlMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		RepositoryApiUrlResponseMessage)

	// Store the Available TemplateRepositoryApiUrls as a map structure in TestCase-struct
	for _, templateRepositoryApiUrlToBeStored := range templateRepositoryApiUrlsToBeStored {

		testCaseModeReference.
			TemplateRepositoryApiUrlMap[templateRepositoryApiUrlToBeStored.GetRepositoryApiUrlName()] =
			templateRepositoryApiUrlToBeStored

	}

}

// Store TestData that user can use within TestCases
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) storeTestData(
	testDataFromSimpleTestDataAreaFiles []*fenixGuiTestCaseBuilderServerGrpcApi.TestDataFromOneSimpleTestDataAreaFileMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	// Store the TestData
	availableBuildingBlocksModel.TestData = testDataFromSimpleTestDataAreaFiles

	// Loop all TestDataFiles for TestData-Areas and add to the TestData-model
	var testDataFromTestDataArea testDataEngine.TestDataFromSimpleTestDataAreaStruct
	for _, testDataFromOneSimpleTestDataAreaFile := range testDataFromSimpleTestDataAreaFiles {

		// Convert Headers
		var header struct {
			ShouldHeaderActAsFilter bool
			HeaderName              string
			HeaderUiName            string
		}
		var headers []struct {
			ShouldHeaderActAsFilter bool
			HeaderName              string
			HeaderUiName            string
		}
		for _, rawHeader := range testDataFromOneSimpleTestDataAreaFile.HeadersForTestDataFromOneSimpleTestDataAreaFile {

			// Set values to 'header'
			header.ShouldHeaderActAsFilter = rawHeader.GetShouldHeaderActAsFilter()
			header.HeaderName = rawHeader.GetHeaderName()
			header.HeaderUiName = rawHeader.GetHeaderUiName()

			// Add to the slice of headers
			headers = append(headers, header)
		}

		// Convert TestDataRows
		var row []string
		var rows [][]string

		for _, simpleTestDataRow := range testDataFromOneSimpleTestDataAreaFile.SimpleTestDataRows {

			// Set values to 'row'
			row = simpleTestDataRow.GetTestDataValue()

			// Add to the slice of headers
			rows = append(rows, row)
		}

		// Populate the TestDataFromTestDataArea-structure
		testDataFromTestDataArea = testDataEngine.TestDataFromSimpleTestDataAreaStruct{
			TestDataDomainUuid:         testDataFromOneSimpleTestDataAreaFile.GetTestDataDomainUuid(),
			TestDataDomainName:         testDataFromOneSimpleTestDataAreaFile.GetTestDataDomainName(),
			TestDataDomainTemplateName: testDataFromOneSimpleTestDataAreaFile.GetTestDataDomainTemplateName(),
			TestDataAreaUuid:           testDataFromOneSimpleTestDataAreaFile.GetTestDataAreaUuid(),
			TestDataAreaName:           testDataFromOneSimpleTestDataAreaFile.GetTestDataAreaName(),
			Headers:                    headers,
			TestDataRows:               rows,
			TestDataFileSha256Hash:     testDataFromOneSimpleTestDataAreaFile.GetTestDataFileSha256Hash(),
		}

		// Add TestData to TestDataModel
		testDataEngine.AddTestDataToTestDataModel(testDataFromTestDataArea)
	}

}

// Store Users available ExecutionDomains to be used with Fenix-created TestInstructions that should be sent to other Domain then Fenix
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) storeUsersAvailableExecutionDomains(
	executionDomainsThatCanReceiveDirectTargetedTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.
		ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	// Initiate the Map
	availableBuildingBlocksModel.
		executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage)

	// Loop slice with ExecutionDomains and add to the Map
	for _, executionDomainsThatCanReceiveDirectTargetedTestInstruction := range executionDomainsThatCanReceiveDirectTargetedTestInstructions {
		availableBuildingBlocksModel.
			executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap[executionDomainsThatCanReceiveDirectTargetedTestInstruction.GetNameUsedInGui()] = executionDomainsThatCanReceiveDirectTargetedTestInstruction
	}

	// Store a pointer to 'ExecutionDomains'
	sharedCode.ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMapPtr =
		&availableBuildingBlocksModel.executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap

}

// Convert gRPC-message for TI or TIC into model used within the TestCase-model
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) convertGrpcElementModelIntoTestCaseElementModel(immatureGrpcElementModelMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage) (immatureTestCaseElementModel testCaseModel.ImmatureElementStruct) {

	// Initiate map used in TestCaseModel
	immatureTestCaseElementModel.ImmatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Loop over gRPC-element-model-structure
	for _, gRpcElementModel := range immatureGrpcElementModelMessage.TestCaseModelElements {
		immatureTestCaseElementModel.ImmatureElementMap[gRpcElementModel.ImmatureElementUuid] = *gRpcElementModel
	}

	// Set the first Element
	immatureTestCaseElementModel.FirstElementUuid = immatureGrpcElementModelMessage.FirstImmatureElementUuid

	return immatureTestCaseElementModel
}
