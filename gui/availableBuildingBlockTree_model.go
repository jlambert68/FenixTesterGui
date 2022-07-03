package gui

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// Gets the model used to drive the Available Building Blocks-Tree
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableBuildingBlocksModel() map[string][]string {

	return availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView
}

// Generate the model used to drive the Available Building Blocks-Tree
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) makeTreeUIModel() {

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
		// Add TestInstructionType to UI-tree model
		availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableBuildingBlocksModel.generateUITreeNameForTestInstructionsHeader(domain)] = testInstructionTypeNamesInUITree

		// For 'TestInstructionContainerHeaderName' add a list of all TestInstructionContainerTypes
		availableTestInstructionContainerTypesFromModel := availableBuildingBlocksModel.getAvailableTestInstructionContainerTypesFromModel(domain)
		var testInstructionContainerTypeNamesInUITree []string
		// Loop all TestInstructionContainerTypes and extract UI-tree name
		for _, testInstructionContainerTypeInUITree := range availableTestInstructionContainerTypesFromModel {
			testInstructionContainerTypeNamesInUITree = append(testInstructionContainerTypeNamesInUITree, testInstructionContainerTypeInUITree.nameInUITree)
		}
		// Add TestInstructionContainerType to UI-tree model
		availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableBuildingBlocksModel.generateUITreeNameForTestInstructionContainersHeader(domain)] = testInstructionContainerTypeNamesInUITree

		// For each 'TestInstructionType' add a list of all TestInstructions
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			var testInstructionNamesInUITree []string
			availableTestInstructions := availableBuildingBlocksModel.getAvailableTestInstructionsFromModel(domain, availableTestInstructionTypeFromModel)
			// Loop all TestInstructions and add the UI-tree name to array
			for _, availableTestInstruction := range availableTestInstructions {
				testInstructionNamesInUITree = append(testInstructionNamesInUITree, availableTestInstruction.nameInUITree)
			}
			// Add TestInstructions to UI-tree model
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
			// Add TestInstructionContainers to UI-tree model
			availableBuildingBlocksModel.availableBuildingBlockModelSuitedForFyneTreeView[availableTestInstructionContainerTypeFromModel.nameInUITree] = testInstructionContainerNamesInUITree
		}
	}
}

// Load all Available Building Blocks from Gui-server
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) loadAvailableBuildingBlocksFromServer() {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	testInstructionsAndTestContainersMessage = availableBuildingBlocksModel.grpcOut.SendListAllAvailableTestInstructionsAndTestInstructionContainers("s41797") //TODO change to use current logged in to computer user

	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage)

	fmt.Println(testInstructionsAndTestContainersMessage)

}

// Load all Pinned Building Blocks from Gui-server
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) loadPinnedBuildingBlocksFromServer() {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	testInstructionsAndTestContainersMessage = availableBuildingBlocksModel.grpcOut.SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers("s41797") //TODO change to use current logged in to computer user

	availableBuildingBlocksModel.loadModelWithPinnedBuildingBlocks(testInstructionsAndTestContainersMessage)

	//fmt.Println(testInstructionsAndTestContainersMessage)

}

// *********** Generate Names for UI-Tree (Start)***********

// Generate UI Tree name for 'Domain', TestInstructionType, TestInstruction, TestInstructionContainerType and TestInstructionContainer for the Available Building Blocks UI-Tree
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) generateUITreeName(node availableBuildingBlocksForUITreeNodesStruct, domainName string) (treeName string, pinnedTreeName string) {

	treeName = node.name + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	pinnedTreeName = node.name + " (" + domainName + ")" + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName, pinnedTreeName
}

// Generate UI Tree name for 'TestInstructionsHeader' for Available Building Blocks
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionsHeader(domain availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionsHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionContainersHeader' for Available Building Blocks
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) generateUITreeNameForTestInstructionContainersHeader(domain availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionContainersHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableDomainTreeNamesFromModel() (availableDomainTreeNamesList []string) {

	availableDomains := availableBuildingBlocksModel.getAvailableDomainsFromModel()

	for _, domain := range availableDomains {
		availableDomainTreeNamesList = append(availableDomainTreeNamesList, domain.nameInUITree)
	}

	return availableDomainTreeNamesList
}

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableDomainsFromModel() (availableDomains []availableBuildingBlocksForUITreeNodesStruct) {

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

// Extract all 'TestInstructionTypes', per Domain, with Names suited for Tree-model
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableTestInstructionTypesFromModel(domain availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionTypes []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionType' for specific domain
	testInstructionTypes := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	for testInstructionType := range testInstructionTypes {
		availableTestInstructionTypes = append(availableTestInstructionTypes, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionType])
	}

	return availableTestInstructionTypes
}

// Extract all 'TestInstructionContainerTypes', per Domain, with Names suited for Tree-model
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableTestInstructionContainerTypesFromModel(domain availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainerTypes []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainerType' for specific domain
	testInstructionContainerTypes := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	for testInstructionContainerType := range testInstructionContainerTypes {
		availableTestInstructionContainerTypes = append(availableTestInstructionContainerTypes, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainerType])
	}
	return availableTestInstructionContainerTypes
}

// Extract all 'TestInstructions', per TestInstructionType, with Names suited for Tree-model
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableTestInstructionsFromModel(domain availableBuildingBlocksForUITreeNodesStruct, testInstructionType availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructions []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructions' for specific TestInstructionType
	testInstructionTypes := availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	testInstructions := testInstructionTypes[testInstructionType.uuid]
	for testInstruction := range testInstructions {
		availableTestInstructions = append(availableTestInstructions, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction])
	}

	return availableTestInstructions
}

// Extract all 'TestInstructionContainers', per TestInstructionContainerType, with Names suited for Tree-model
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getAvailableTestInstructionContainersFromModel(domain availableBuildingBlocksForUITreeNodesStruct, testInstructionContainerType availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainers []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainers' for specific TestInstructionContainerType
	testInstructionContainerTypes := availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	testInstructionContainers := testInstructionContainerTypes[testInstructionContainerType.uuid]
	for testInstructionContainer := range testInstructionContainers {
		availableTestInstructionContainers = append(availableTestInstructionContainers, availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer])
	}

	return availableTestInstructionContainers
}

// Extract all 'Pinned TestInstructions' suited for Tree-model
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) getPinnedBuildingBlocksTreeNamesFromModel() (pinnedBuildingBlocks []string) {

	// Create the list of Pinned Building Blocks with names suited for UI-Trre
	for pinnedBuildingBlockTreeName, _ := range availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes {
		pinnedBuildingBlocks = append(pinnedBuildingBlocks, pinnedBuildingBlockTreeName)
	}

	return pinnedBuildingBlocks
}

// Pin one Available Building Block (TestInstruction or TestInstructionContainer, if it isn't already pinned
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) pinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree string) (err error) {

	// Verify that Name exists among available Building Blocks NodeNames
	nodeData, existsInMap := availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[nameInAvailableBuildingBlocksTree]

	if existsInMap == false {
		err = errors.New(nameInAvailableBuildingBlocksTree + " is missing among nodes i map")
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id":  "9d3510ec-8b9e-4490-bae9-0e6cf9c0a1cb",
			"err": err,
		}).Error(nameInAvailableBuildingBlocksTree + " is missing among nodes i map 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'")
		return err
	}

	// Verify that nod is not already pinned, equals exists in TreeNameToUuid for pinned Building Blocks
	tempPinnedNameInUITree := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[nodeData.uuid].pinnedNameInUITree

	_, existsInMap = availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[tempPinnedNameInUITree]
	if existsInMap == true {
		err = errors.New("building block is already pinned")
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id":  "e1d22ba0-072f-4be5-a2d9-4b73278c170c",
			"err": err,
		}).Error(nameInAvailableBuildingBlocksTree + " is already pinned, or exists in map 'availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes['")
		return err
	}

	// Do the Pin of the Building Block by adding it to 'pinnedBuildingBlocksForUITreeNodes'-map
	tempPinnedNodeData := uiTreeNodesNameToUuidStruct{
		uuid:              nodeData.uuid,
		buildingBlockType: nodeData.buildingBlockType,
	}
	availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[tempPinnedNameInUITree] = tempPinnedNodeData

	return err

	//TODO Rebuild UI-tree-model
}

// Unpin one pinned Available Building Block (TestInstruction or TestInstructionContainer
func (availableBuildingBlocksModel *availableBuildingBlocksModelStruct) unPinTestInstructionOrTestInstructionContainer(nameInAvailableBuildingBlocksTree string) (err error) {

	// Verify that Name exists among available Building Blocks NodeNames
	nodeData, existsInMap := availableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[nameInAvailableBuildingBlocksTree]

	if existsInMap == false {
		err = errors.New(nameInAvailableBuildingBlocksTree + " is missing among nodes i map")
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id":  "3e8af427-d2a7-4d01-95b0-45817e33fbc4",
			"err": err,
		}).Error(nameInAvailableBuildingBlocksTree + " is missing among nodes i map 'availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes'")
		return err
	}

	// Verify that nod is pinned, equals exists in TreeNameToUuid for pinned Building Blocks
	tempPinnedNameInUITree := availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[nodeData.uuid].pinnedNameInUITree

	_, existsInMap = availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes[tempPinnedNameInUITree]
	if existsInMap == false {
		err = errors.New("building block is not  pinned")
		availableBuildingBlocksModel.logger.WithFields(logrus.Fields{
			"id":  "be6e39f1-09dc-4532-9819-d516d8ca9661",
			"err": err,
		}).Error(nameInAvailableBuildingBlocksTree + " is not pinned, or exists in map 'availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes['")
		return err
	}

	// Do the UnPin of the Building Block by removing it to 'pinnedBuildingBlocksForUITreeNodes'-map
	delete(availableBuildingBlocksModel.pinnedBuildingBlocksForUITreeNodes, tempPinnedNameInUITree)

	return err

	//TODO Rebuild UI-tree-model
}
