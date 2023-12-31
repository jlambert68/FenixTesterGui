package gui

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sort"
)

// Gets the testCaseModel used to drive the Available Building Blocks-Tree
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableBuildingBlocksModel() map[string][]string {

	return availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView
}

// Generate the testCaseModel used to drive the Available Building Blocks-Tree
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) makeTreeUIModel() {

	availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView = map[string][]string{
		"":                            {TestCaseBuildingBlocksHeader},
		TestCaseBuildingBlocksHeader:  {PinnedBuildingBlocksHeader, AvailableBuildingBlocksHeader},
		PinnedBuildingBlocksHeader:    availableBuildingBlocksModel.getPinnedBuildingBlocksTreeNamesFromModel(),
		AvailableBuildingBlocksHeader: availableBuildingBlocksModel.getAvailableDomainTreeNamesFromModel(),
	}

	// Loop all Domains
	availableDomains := availableBuildingBlocksModel.getAvailableDomainsFromModel()

	// Sort 'Domains' in Name-order
	sort.SliceStable(availableDomains, func(i, j int) bool {
		if availableDomains[i].nameInUITree != availableDomains[j].nameInUITree {
			return availableDomains[i].nameInUITree < availableDomains[j].nameInUITree
		}

		return availableDomains[i].nameInUITree < availableDomains[j].nameInUITree
	})

	for _, domain := range availableDomains {
		// For each domain add TestInstructionHeaderName and TestInstructionContainerHeaderName
		availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[domain.nameInUITree] = []string{
			availableBuildingBlocksModel.generateUITreeNameForTestInstructionsHeader(domain),
			availableBuildingBlocksModel.generateUITreeNameForTestInstructionContainersHeader(domain)}

		// For 'TestInstructionHeaderName' add a list of all TestInstructionTypes
		availableTestInstructionTypesFromModel := availableBuildingBlocksModel.getAvailableTestInstructionTypesFromModel(domain)
		var testInstructionTypeNamesInUITree []string
		// Loop all TestInstructionTypes and extract UI-tree name
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			testInstructionTypeNamesInUITree = append(testInstructionTypeNamesInUITree, availableTestInstructionTypeFromModel.nameInUITree)
		}

		// Sort 'TestInstructionTypes' in Name-order
		sort.SliceStable(testInstructionTypeNamesInUITree, func(i, j int) bool {
			if testInstructionTypeNamesInUITree[i] != testInstructionTypeNamesInUITree[j] {
				return testInstructionTypeNamesInUITree[i] < testInstructionTypeNamesInUITree[j]
			}

			return testInstructionTypeNamesInUITree[i] < testInstructionTypeNamesInUITree[j]
		})

		// Add TestInstructionType to UI-tree testCaseModel
		availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableBuildingBlocksModel.generateUITreeNameForTestInstructionsHeader(domain)] = testInstructionTypeNamesInUITree

		// For 'TestInstructionContainerHeaderName' add a list of all TestInstructionContainerTypes
		availableTestInstructionContainerTypesFromModel := availableBuildingBlocksModel.getAvailableTestInstructionContainerTypesFromModel(domain)
		var testInstructionContainerTypeNamesInUITree []string
		// Loop all TestInstructionContainerTypes and extract UI-tree name
		for _, testInstructionContainerTypeInUITree := range availableTestInstructionContainerTypesFromModel {
			testInstructionContainerTypeNamesInUITree = append(testInstructionContainerTypeNamesInUITree, testInstructionContainerTypeInUITree.nameInUITree)
		}

		// Sort 'TestInstructionContainerTypes' in Name-order
		sort.SliceStable(testInstructionContainerTypeNamesInUITree, func(i, j int) bool {
			if testInstructionContainerTypeNamesInUITree[i] != testInstructionContainerTypeNamesInUITree[j] {
				return testInstructionContainerTypeNamesInUITree[i] < testInstructionContainerTypeNamesInUITree[j]
			}

			return testInstructionContainerTypeNamesInUITree[i] < testInstructionContainerTypeNamesInUITree[j]
		})

		// Add TestInstructionContainerType to UI-tree testCaseModel
		availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableBuildingBlocksModel.generateUITreeNameForTestInstructionContainersHeader(domain)] = testInstructionContainerTypeNamesInUITree

		// For each 'TestInstructionType' add a list of all TestInstructions
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			var testInstructionNamesInUITree []string
			availableTestInstructions := availableBuildingBlocksModel.getAvailableTestInstructionsFromModel(domain, availableTestInstructionTypeFromModel)
			// Loop all TestInstructions and add the UI-tree name to array
			for _, availableTestInstruction := range availableTestInstructions {
				testInstructionNamesInUITree = append(testInstructionNamesInUITree, availableTestInstruction.nameInUITree)
			}

			// Sort 'TestInstructions' in Name-order
			sort.SliceStable(testInstructionNamesInUITree, func(i, j int) bool {
				if testInstructionNamesInUITree[i] != testInstructionNamesInUITree[j] {
					return testInstructionNamesInUITree[i] < testInstructionNamesInUITree[j]
				}

				return testInstructionNamesInUITree[i] < testInstructionNamesInUITree[j]
			})

			// Add TestInstructions to UI-tree testCaseModel
			availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableTestInstructionTypeFromModel.nameInUITree] = testInstructionNamesInUITree
		}

		// For each 'TestInstructionContainerType' add a list of all TestInstructionContainers
		for _, availableTestInstructionContainerTypeFromModel := range availableTestInstructionContainerTypesFromModel {
			var testInstructionContainerNamesInUITree []string
			availableTestInstructionContainers := availableBuildingBlocksModel.getAvailableTestInstructionContainersFromModel(domain, availableTestInstructionContainerTypeFromModel)
			// Loop all TestInstructionContainers and add the UI-tree name to array
			for _, availableTestInstructionContainer := range availableTestInstructionContainers {
				testInstructionContainerNamesInUITree = append(testInstructionContainerNamesInUITree, availableTestInstructionContainer.nameInUITree)
			}

			// Sort 'TestInstructionContainers' in Name-order
			sort.SliceStable(testInstructionContainerNamesInUITree, func(i, j int) bool {
				if testInstructionContainerNamesInUITree[i] != testInstructionContainerNamesInUITree[j] {
					return testInstructionContainerNamesInUITree[i] < testInstructionContainerNamesInUITree[j]
				}

				return testInstructionContainerNamesInUITree[i] < testInstructionContainerNamesInUITree[j]
			})

			// Add TestInstructionContainers to UI-tree testCaseModel
			availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableTestInstructionContainerTypeFromModel.nameInUITree] = testInstructionContainerNamesInUITree
		}
	}
}

// Load all Available Building Blocks from Gui-server
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadAvailableBuildingBlocksFromServer(testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//grpcOut := grpc_out_GuiTestCaseBuilderServer.GRPCOutStruct{}
	testInstructionsAndTestContainersMessage = availableBuildingBlocksModel.grpcOut.SendListAllAvailableTestInstructionsAndTestInstructionContainers("s41797") //TODO change to use current logged in to computer user

	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage)

	// Load TestCase-model with available Immature TestInstruction and TestInstructionContainers TODO Put Immature TI and TIC, and BONDS, in separate object
	testCaseModeReference.AvailableImmatureTestInstructionsMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage)
	testCaseModeReference.AvailableImmatureTestInstructionContainersMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage)

	testCaseModeReference.ImmatureDropZonesDataMap = make(map[string]testCaseModel.ImmatureDropZoneDataMapStruct)

	// Loop TestInstructions and add to map
	for immatureTestInstructionCounter, _ := range testInstructionsAndTestContainersMessage.ImmatureTestInstructions {

		// Need to do this because otherwise I don't get all subObjects(DropZone-data). Seems to be some bug
		var immatureTestInstruction *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage
		immatureTestInstruction = testInstructionsAndTestContainersMessage.ImmatureTestInstructions[immatureTestInstructionCounter]

		var tempImmatureTestInstructionMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage
		tempImmatureTestInstructionMessage = &fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage{
			BasicTestInstructionInformation:    immatureTestInstruction.BasicTestInstructionInformation,
			ImmatureTestInstructionInformation: immatureTestInstruction.ImmatureTestInstructionInformation,
			ImmatureSubTestCaseModel:           immatureTestInstruction.ImmatureSubTestCaseModel,
		}

		/*
			tempImmatureTestInstructionMessage = testCaseModeReference.AvailableImmatureTestInstructionsMap[immatureTestInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOrignalUuid]

			tempImmatureTestInstructionMessage.BasicTestInstructionInformation = immatureTestInstruction.BasicTestInstructionInformation
			tempImmatureTestInstructionMessage.ImmatureTestInstructionInformation = immatureTestInstruction.ImmatureTestInstructionInformation
			tempImmatureTestInstructionMessage.ImmatureSubTestCaseModel = immatureTestInstruction.ImmatureSubTestCaseModel
		*/
		testCaseModeReference.AvailableImmatureTestInstructionsMap[immatureTestInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOrignalUuid] = tempImmatureTestInstructionMessage

		//testCaseModeReference.AvailableImmatureTestInstructionsMap[immatureTestInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOrignalUuid] = immatureTestInstruction

		// Extract DropZone and add ti Map
		for _, dropZoneMessage := range immatureTestInstruction.ImmatureTestInstructionInformation.AvailableDropZones {

			// DropZoneUuid should not be empty
			if dropZoneMessage.DropZoneUuid == "" {
				errorId := "a8226610-8203-4305-a71d-1eb80923374c"
				err := errors.New(fmt.Sprintf("dropZoneUuid is emtpy in TestInstruction %s with name %s 'immatureTestInstruction.ImmatureTestInstructionInformation.AvailableDropZones ' [ErrorID: %s]", immatureTestInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOrignalUuid, immatureTestInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionOriginalName, errorId))

				fmt.Println(err) //TODO Send error over error-channel
			}

			// Verify that DropZoneUuid doesn't already exit in Map
			_, existInMap := testCaseModeReference.ImmatureDropZonesDataMap[dropZoneMessage.DropZoneUuid]
			if existInMap == true {

				errorId := "53d0d63d-777a-413e-a0e5-61102d73df0a"
				err := errors.New(fmt.Sprintf("dropZoneUuid %s already exist in ImmatureDropZonesDataMap [ErrorID: %s]", dropZoneMessage.DropZoneUuid, errorId))

				fmt.Println(err) //TODO Send error over error-channel

			} else {

				var tempImmatureDropZoneData testCaseModel.ImmatureDropZoneDataMapStruct

				tempImmatureDropZoneData = testCaseModel.ImmatureDropZoneDataMapStruct{
					DropZoneUuid:        dropZoneMessage.DropZoneUuid,
					DropZoneName:        dropZoneMessage.DropZoneName,
					DropZoneDescription: dropZoneMessage.DropZoneDescription,
					DropZoneMouseOver:   dropZoneMessage.DropZoneMouseOver,
					DropZoneColor:       dropZoneMessage.DropZoneColor,
					DropZonePreSetTestInstructionAttributesMap: make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage),
				}

				for _, dropZoneAttribute := range dropZoneMessage.DropZonePreSetTestInstructionAttributes {

					tempDropZonePreSetTestInstructionAttributeMessage := &fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage{
						TestInstructionAttributeType: dropZoneAttribute.TestInstructionAttributeType,
						TestInstructionAttributeUuid: dropZoneAttribute.TestInstructionAttributeUuid,
						TestInstructionAttributeName: dropZoneAttribute.TestInstructionAttributeName,
						AttributeValueAsString:       dropZoneAttribute.AttributeValueAsString,
						AttributeValueUuid:           dropZoneAttribute.AttributeValueUuid,
						AttributeActionCommand:       dropZoneAttribute.AttributeActionCommand,
					}

					// Add DropZone Attribute to DropZone Attributes Map
					tempImmatureDropZoneData.DropZonePreSetTestInstructionAttributesMap[dropZoneAttribute.TestInstructionAttributeUuid] = tempDropZonePreSetTestInstructionAttributeMessage

				}

				// Add all DropZone to DropZones-map
				testCaseModeReference.ImmatureDropZonesDataMap[dropZoneMessage.DropZoneUuid] = tempImmatureDropZoneData
			}
		}

	}
	//TODO make same changes to TestInstructionContainers as for TestInstructions
	// Loop TestInstructionContainers and add to map
	for immatureTestInstructionContainerCounter, _ := range testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers {

		// Need to do this because otherwise I don't get all subObjects(DropZone-data). Seems to be some bug
		var immatureTestInstructionContainer *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage
		immatureTestInstructionContainer = testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers[immatureTestInstructionContainerCounter]

		testCaseModeReference.AvailableImmatureTestInstructionContainersMap[immatureTestInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = immatureTestInstructionContainer
	}

	// fmt.Println(testInstructionsAndTestContainersMessage)

}

// Load all Pinned Building Blocks from Gui-server
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadPinnedBuildingBlocksFromServer() {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//grpcOut := grpc_out_GuiTestCaseBuilderServer.GRPCOutStruct{}
	testInstructionsAndTestContainersMessage = availableBuildingBlocksModel.grpcOut.SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers("s41797") //TODO change to use current logged in to computer user

	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(testInstructionsAndTestContainersMessage)

	//fmt.Println(testInstructionsAndTestContainersMessage)

}

// Save all Pinned Building Blocks to Gui-server
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) savePinnedBuildingBlocksFromServer() (err error) {

	var pinnedTestInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage
	var availablePinnedTestInstructions []*fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionMessage
	var availablePinnedPreCreatedTestInstructionContainers []*fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedPreCreatedTestInstructionContainerMessage

	// Loop testCaseModel for pinned Building Blocks
	for _, pinnedNameObject := range availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes {

		buildingBlock := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[pinnedNameObject.uuid]
		switch buildingBlock.BuildingBlockType {

		// Append TestInstruction
		case TestInstruction:
			availablePinnedTestInstruction := fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionMessage{
				TestInstructionUuid: buildingBlock.uuid,
				TestInstructionName: buildingBlock.name,
			}
			availablePinnedTestInstructions = append(availablePinnedTestInstructions, &availablePinnedTestInstruction)

			// Append TestInstructionContainer
		case TestInstructionContainer:
			availablePinnedTestInstructionContainer := fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedPreCreatedTestInstructionContainerMessage{
				TestInstructionContainerUuid: buildingBlock.uuid,
				TestInstructionContainerName: buildingBlock.name,
			}
			availablePinnedPreCreatedTestInstructionContainers = append(availablePinnedPreCreatedTestInstructionContainers, &availablePinnedTestInstructionContainer)

		}
	}

	pinnedTestInstructionsAndTestContainersMessage = &fenixGuiTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage{
		UserIdOnComputer:     sharedCode.CurrentUserIdLogedInOnComputer,
		GCPAuthenticatedUser: sharedCode.CurrentUserAuthenticatedTowardsGCP,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			availableBuildingBlocksModel.grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		AvailablePinnedTestInstructions:                    availablePinnedTestInstructions,
		AvailablePinnedPreCreatedTestInstructionContainers: availablePinnedPreCreatedTestInstructionContainers,
	}

	returnMessage := availableBuildingBlocksModel.grpcOut.SendSaveAllPinnedTestInstructionsAndTestInstructionContainers(pinnedTestInstructionsAndTestContainersMessage)

	if returnMessage.AckNack == false {
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"Id":  "c6fa24be-4473-457c-b899-a5d17c590096",
			"err": returnMessage.Comments,
		}).Error("Got some err when saving Pinned Building Blocks")

		return errors.New("got some err when saving Pinned Building Blocks")
	}

	return nil

}

// *********** Generate Names for UI-Tree (Start)***********

// Generate UI Tree name for 'Domain', TestInstructionType, TestInstruction, TestInstructionContainerType and TestInstructionContainer for the Available Building Blocks UI-Tree
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) generateUITreeName(node AvailableBuildingBlocksForUITreeNodesStruct, domainName string) (treeName string, pinnedTreeName string) {

	treeName = node.name + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	pinnedTreeName = node.name + " (" + domainName + ")" + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName, pinnedTreeName
}

// Generate UI Tree name for 'TestInstructionsHeader' for Available Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionsHeader(domain AvailableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionsHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionContainersHeader' for Available Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionContainersHeader(domain AvailableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionContainersHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Extract all 'Domains', with Names suited for Tree-testCaseModel, for the testCaseModel tha underpins the UI Tree for Available Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableDomainTreeNamesFromModel() (availableDomainTreeNamesList []string) {

	availableDomains := availableBuildingBlocksModel.getAvailableDomainsFromModel()

	for _, domain := range availableDomains {
		availableDomainTreeNamesList = append(availableDomainTreeNamesList, domain.nameInUITree)
	}

	return availableDomainTreeNamesList
}

// Extract all 'Domains', with Names suited for Tree-testCaseModel, for the testCaseModel tha underpins the UI Tree for Available Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableDomainsFromModel() (availableDomains []AvailableBuildingBlocksForUITreeNodesStruct) {

	// Extract Domain nodes from TestInstruction-map
	domainNodesInTestInstructionMap := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap
	for key := range domainNodesInTestInstructionMap {
		if key != TopNodeForAvailableDomainsMap {
			availableDomains = append(availableDomains, availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[key])
		}
	}

	// Extract Domain nodes from TestInstructionContainer-map
	domainNodesInTestInstructionContainerMap := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap
	for domainUuid := range domainNodesInTestInstructionContainerMap {
		if domainUuid != TopNodeForAvailableDomainsMap {
			_, existsInMap := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domainUuid]
			if existsInMap == false {
				availableDomains = append(availableDomains, availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[domainUuid])
			}
		}
	}

	return availableDomains
}

// Extract all 'TestInstructionTypes', per Domain, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionTypesFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionTypes []AvailableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionType' for specific domain
	testInstructionTypes := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	for testInstructionType := range testInstructionTypes {
		availableTestInstructionTypes = append(availableTestInstructionTypes, availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionType])
	}

	return availableTestInstructionTypes
}

// Extract all 'TestInstructionContainerTypes', per Domain, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionContainerTypesFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainerTypes []AvailableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainerType' for specific domain
	testInstructionContainerTypes := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	for testInstructionContainerType := range testInstructionContainerTypes {
		availableTestInstructionContainerTypes = append(availableTestInstructionContainerTypes, availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainerType])
	}
	return availableTestInstructionContainerTypes
}

// Extract all 'TestInstructions', per TestInstructionType, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionsFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct, testInstructionType AvailableBuildingBlocksForUITreeNodesStruct) (availableTestInstructions []AvailableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructions' for specific TestInstructionType
	testInstructionTypes := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	testInstructions := testInstructionTypes[testInstructionType.uuid]
	for testInstruction := range testInstructions {
		availableTestInstructions = append(availableTestInstructions, availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstruction])
	}

	return availableTestInstructions
}

// Extract all 'TestInstructionContainers', per TestInstructionContainerType, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionContainersFromModel(domain AvailableBuildingBlocksForUITreeNodesStruct, testInstructionContainerType AvailableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainers []AvailableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainers' for specific TestInstructionContainerType
	testInstructionContainerTypes := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	testInstructionContainers := testInstructionContainerTypes[testInstructionContainerType.uuid]
	for testInstructionContainer := range testInstructionContainers {
		availableTestInstructionContainers = append(availableTestInstructionContainers, availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[testInstructionContainer])
	}

	return availableTestInstructionContainers
}

// Extract all 'Pinned TestInstructions' suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getPinnedBuildingBlocksTreeNamesFromModel() (pinnedBuildingBlocks []string) {

	// Create the list of Pinned Building Blocks with names suited for UI-Trre
	for pinnedBuildingBlockTreeName := range availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes {
		pinnedBuildingBlocks = append(pinnedBuildingBlocks, pinnedBuildingBlockTreeName)
	}

	return pinnedBuildingBlocks
}

// Verify that it is possible to Pin one Available Building Block (TestInstruction or TestInstructionContainer, if it isn't already pinned
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) verifyBeforePinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree string, onlyForVerifying bool) (err error) {

	// Verify that Name exists among available Building Blocks NodeNames
	nodeData, existsInMap := availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[nameInAvailableBuildingBlocksTree]

	if existsInMap == false {
		err = errors.New(nameInAvailableBuildingBlocksTree + " is missing among nodes i map")
		err = errors.New(nameInAvailableBuildingBlocksTree + " is missing among nodes i map")
		{
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":  "9d3510ec-8b9e-4490-bae9-0e6cf9c0a1cb",
				"err": err,
			}).Error(nameInAvailableBuildingBlocksTree + " is missing among nodes i map 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'")

		}
		return err
	}

	// Verify that nod is not already pinned, equals exists in TreeNameToUuid for pinned Building Blocks
	tempPinnedNameInUITree := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[nodeData.uuid].pinnedNameInUITree

	_, existsInMap = availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[tempPinnedNameInUITree]
	if existsInMap == true {
		err = errors.New("building block is already pinned")
		if onlyForVerifying == false {
			// Only create log message if we really tries to pin
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":  "e1d22ba0-072f-4be5-a2d9-4b73278c170c",
				"err": err,
			}).Error(nameInAvailableBuildingBlocksTree + " is already pinned, or exists in map 'availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes['")
		}
		return err
	}

	return err
}

// Pin one Available Building Block (TestInstruction or TestInstructionContainer, if it isn't already pinned
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) pinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree string) (err error) {

	// Verify that node can be pinned
	err = availableBuildingBlocksModel.verifyBeforePinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree, false)

	if err == nil {
		// Do the Pin of the Building Block by adding it to 'pinnedBuildingBlocksForUITreeNodes'-map

		//Extract node-uuid and node type
		nodeData, _ := availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[nameInAvailableBuildingBlocksTree] //
		pinnedNodeName := availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[nodeData.uuid].pinnedNameInUITree

		tempPinnedNodeData := uiTreeNodesNameToUuidStruct{
			uuid:              nodeData.uuid,
			buildingBlockType: nodeData.buildingBlockType,
		}
		availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[pinnedNodeName] = tempPinnedNodeData

	}

	return err

}

// Verify that it is possible to Unpin one pinned Available Building Block (TestInstruction or TestInstructionContainer
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) verifyBeforeUnPinTestInstructionOrTestInstructionContainer(pinnedNameInUITree string, onlyForVerifying bool) (err error) {

	// Verify that nod is pinned, equals exists in TreeNameToUuid for pinned Building Blocks
	pinnedBuildingBlock, existsInMap := availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[pinnedNameInUITree]
	if existsInMap == false {
		err = errors.New("building block is not  pinned")
		if onlyForVerifying == false {
			// Only create log message if we really tries to unpin
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":  "be6e39f1-09dc-4532-9819-d516d8ca9661",
				"err": err,
			}).Error(pinnedNameInUITree + " is not pinned, or exists in map 'availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes['")
		}
		return err
	}

	// Verify that Name exists among available Building Blocks
	_, existsInMap = availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes[pinnedBuildingBlock.uuid]

	if existsInMap == false {
		err = errors.New(pinnedNameInUITree + " is missing among all nodes map")
		if onlyForVerifying == false {
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":  "3e8af427-d2a7-4d01-95b0-45817e33fbc4",
				"err": err,
			}).Error(pinnedNameInUITree + " is missing among nodes i map 'availableBuildingBlocksModel.AvailableBuildingBlocksForUITreeNodes'")
		}
		return err
	}

	return err

}

// Unpin one pinned Available Building Block (TestInstruction or TestInstructionContainer
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) unPinTestInstructionOrTestInstructionContainer(pinnedNameInUITree string) (err error) {

	// Verify that node can be unpinned
	err = availableBuildingBlocksModel.verifyBeforeUnPinTestInstructionOrTestInstructionContainer(pinnedNameInUITree, false)

	if err == nil {
		// Do the UnPin of the Building Block by removing it to 'pinnedBuildingBlocksForUITreeNodes'-map
		delete(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes, pinnedNameInUITree)

	}

	return err
}

// List all Available Building Block (TestInstruction or TestInstructionContainer
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) listAllAvailableBuidlingBlocks() (availableBuidlingBlocksList []string) {

	// Loop all available building blocks and create list to be used in DropDown
	for uuidKey, buidingBlock := range availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid {

		switch buidingBlock.buildingBlockType {

		case TestInstruction:
			availableBuidlingBlocksList = append(availableBuidlingBlocksList, uuidKey+" [TI]")

		case TestInstructionContainer:
			availableBuidlingBlocksList = append(availableBuidlingBlocksList, uuidKey+" [TIC]")

		default:
			availableBuidlingBlocksList = append(availableBuidlingBlocksList, uuidKey+" [UNKNOWN]")
		}
	}

	return availableBuidlingBlocksList
}
