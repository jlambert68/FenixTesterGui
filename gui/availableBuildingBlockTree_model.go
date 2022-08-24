package gui

import (
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
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
		// Add TestInstructionType to UI-tree testCaseModel
		availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableBuildingBlocksModel.generateUITreeNameForTestInstructionsHeader(domain)] = testInstructionTypeNamesInUITree

		// For 'TestInstructionContainerHeaderName' add a list of all TestInstructionContainerTypes
		availableTestInstructionContainerTypesFromModel := availableBuildingBlocksModel.getAvailableTestInstructionContainerTypesFromModel(domain)
		var testInstructionContainerTypeNamesInUITree []string
		// Loop all TestInstructionContainerTypes and extract UI-tree name
		for _, testInstructionContainerTypeInUITree := range availableTestInstructionContainerTypesFromModel {
			testInstructionContainerTypeNamesInUITree = append(testInstructionContainerTypeNamesInUITree, testInstructionContainerTypeInUITree.nameInUITree)
		}
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
			// Add TestInstructionContainers to UI-tree testCaseModel
			availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableTestInstructionContainerTypeFromModel.nameInUITree] = testInstructionContainerNamesInUITree
		}
	}
}

// Load all Available Building Blocks from Gui-server
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadAvailableBuildingBlocksFromServer(testCaseModeReference *testCaseModel.TestCasesModelsStruct) {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	testInstructionsAndTestContainersMessage = availableBuildingBlocksModel.grpcOut.SendListAllAvailableTestInstructionsAndTestInstructionContainers("s41797") //TODO change to use current logged in to computer user

	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage)

	// Load TestCase-model with available Immature TestInstruction and TestInstructionContainers TODO Put Immature TI and TIC, and BONDS, in separate object
	testCaseModeReference.AvailableImmatureTestInstructionsMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage)
	testCaseModeReference.AvailableImmatureTestInstructionContainersMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage)

	// Loop TestInstructions and add to map
	for _, immatureTestInstruction := range testInstructionsAndTestContainersMessage.ImmatureTestInstructions {
		testCaseModeReference.AvailableImmatureTestInstructionsMap[immatureTestInstruction.BasicTestInstructionInformation.NonEditableInformation.TestInstructionUuid] = immatureTestInstruction
	}

	// Loop TestInstructionContainers and add to map
	for _, immatureTestInstructionContainer := range testInstructionsAndTestContainersMessage.ImmatureTestInstructionContainers {
		testCaseModeReference.AvailableImmatureTestInstructionContainersMap[immatureTestInstructionContainer.BasicTestInstructionContainerInformation.NonEditableInformation.TestInstructionContainerUuid] = immatureTestInstructionContainer
	}

	// fmt.Println(testInstructionsAndTestContainersMessage)

}

// Load all Pinned Building Blocks from Gui-server
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) loadPinnedBuildingBlocksFromServer() {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
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

		buildingBlock := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[pinnedNameObject.uuid]
		switch buildingBlock.buildingBlockType {

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
		UserId: "s41797", //TODO change to use dynamic or defined from GUI user-id
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			availableBuildingBlocksModel.grpcOut.GetHighestFenixGuiServerProtoFileVersion()),
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
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) generateUITreeName(node availableBuildingBlocksForUITreeNodesStruct, domainName string) (treeName string, pinnedTreeName string) {

	treeName = node.name + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	pinnedTreeName = node.name + " (" + domainName + ")" + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName, pinnedTreeName
}

// Generate UI Tree name for 'TestInstructionsHeader' for Available Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionsHeader(domain availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionsHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionContainersHeader' for Available Building Blocks
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionContainersHeader(domain availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

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
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableDomainsFromModel() (availableDomains []availableBuildingBlocksForUITreeNodesStruct) {

	// Extract Domain nodes from TestInstruction-map
	domainNodesInTestInstructionMap := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap
	for key := range domainNodesInTestInstructionMap {
		if key != TopNodeForAvailableDomainsMap {
			availableDomains = append(availableDomains, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[key])
		}
	}

	// Extract Domain nodes from TestInstructionContainer-map
	domainNodesInTestInstructionContainerMap := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap
	for domainUuid := range domainNodesInTestInstructionContainerMap {
		if domainUuid != TopNodeForAvailableDomainsMap {
			_, existsInMap := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domainUuid]
			if existsInMap == false {
				availableDomains = append(availableDomains, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[domainUuid])
			}
		}
	}

	return availableDomains
}

// Extract all 'TestInstructionTypes', per Domain, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionTypesFromModel(domain availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionTypes []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionType' for specific domain
	testInstructionTypes := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	for testInstructionType := range testInstructionTypes {
		availableTestInstructionTypes = append(availableTestInstructionTypes, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionType])
	}

	return availableTestInstructionTypes
}

// Extract all 'TestInstructionContainerTypes', per Domain, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionContainerTypesFromModel(domain availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainerTypes []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainerType' for specific domain
	testInstructionContainerTypes := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	for testInstructionContainerType := range testInstructionContainerTypes {
		availableTestInstructionContainerTypes = append(availableTestInstructionContainerTypes, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainerType])
	}
	return availableTestInstructionContainerTypes
}

// Extract all 'TestInstructions', per TestInstructionType, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionsFromModel(domain availableBuildingBlocksForUITreeNodesStruct, testInstructionType availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructions []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructions' for specific TestInstructionType
	testInstructionTypes := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	testInstructions := testInstructionTypes[testInstructionType.uuid]
	for testInstruction := range testInstructions {
		availableTestInstructions = append(availableTestInstructions, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction])
	}

	return availableTestInstructions
}

// Extract all 'TestInstructionContainers', per TestInstructionContainerType, with Names suited for Tree-testCaseModel
func (availableBuildingBlocksModel *AvailableBuildingBlocksModelStruct) getAvailableTestInstructionContainersFromModel(domain availableBuildingBlocksForUITreeNodesStruct, testInstructionContainerType availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainers []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainers' for specific TestInstructionContainerType
	testInstructionContainerTypes := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	testInstructionContainers := testInstructionContainerTypes[testInstructionContainerType.uuid]
	for testInstructionContainer := range testInstructionContainers {
		availableTestInstructionContainers = append(availableTestInstructionContainers, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer])
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
			}).Error(nameInAvailableBuildingBlocksTree + " is missing among nodes i map 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'")

		}
		return err
	}

	// Verify that nod is not already pinned, equals exists in TreeNameToUuid for pinned Building Blocks
	tempPinnedNameInUITree := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[nodeData.uuid].pinnedNameInUITree

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
		pinnedNodeName := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[nodeData.uuid].pinnedNameInUITree

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
	_, existsInMap = availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[pinnedBuildingBlock.uuid]

	if existsInMap == false {
		err = errors.New(pinnedNameInUITree + " is missing among all nodes map")
		if onlyForVerifying == false {
			availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
				"id":  "3e8af427-d2a7-4d01-95b0-45817e33fbc4",
				"err": err,
			}).Error(pinnedNameInUITree + " is missing among nodes i map 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'")
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
