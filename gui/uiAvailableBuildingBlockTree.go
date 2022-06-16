package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	//"golang.org/x/exp/maps"
)

//______________________________________________________________________________

var list map[string][]string
var tree *widget.Tree
var availableBuildingBlock map[string][]string

func (uiServer *UIServerStruct) makeTreeUI() {
	list = map[string][]string{
		"":  {"A"},
		"A": {"B", "D"},
		"B": {"C"},
		"C": {"abc"},
		"D": {"E"},
		"E": {"F", "G"},
	}

	availableBuildingBlock = map[string][]string{
		"":                            {TestCaseBuildingBlocksHeader},
		TestCaseBuildingBlocksHeader:  {PinnedBuildingBlocksHeader, AvailableBuildingBlocksHeader},
		PinnedBuildingBlocksHeader:    {"Nothing here, yet"},
		AvailableBuildingBlocksHeader: uiServer.getAvailableDomainTreeNamesFromModel(),
	}

	// Loop all Domains
	availableDomains := uiServer.getAvailableDomainsFromModel()
	for _, domain := range availableDomains {
		// For each domain add TestInstructionHeaderName and TestInstructionContainerHeaderName
		availableBuildingBlock[domain.nameInUITree] = []string{
			uiServer.generateUITreeNameForTestInstructionsHeader(domain),
			uiServer.generateUITreeNameForTestInstructionContainersHeader(domain)}

		// For 'TestInstructionHeaderName' add a list of all TestInstructionTypes
		availableTestInstructionTypesFromModel := uiServer.getAvailableTestInstructionTypesFromModel(domain)
		var testInstructionTypeNamesInUITree []string
		// Loop all TestInstructionTypes and extract UI-tree name
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			testInstructionTypeNamesInUITree = append(testInstructionTypeNamesInUITree, availableTestInstructionTypeFromModel.nameInUITree)
		}
		// Add TestInstructionType to UI-tree model
		availableBuildingBlock[uiServer.generateUITreeNameForTestInstructionsHeader(domain)] = testInstructionTypeNamesInUITree

		// For 'TestInstructionContainerHeaderName' add a list of all TestInstructionContainerTypes
		availableTestInstructionContainerTypesFromModel := uiServer.getAvailableTestInstructionContainerTypesFromModel(domain)
		var testInstructionContainerTypeNamesInUITree []string
		// Loop all TestInstructionContainerTypes and extract UI-tree name
		for _, testInstructionContainerTypeInUITree := range availableTestInstructionContainerTypesFromModel {
			testInstructionContainerTypeNamesInUITree = append(testInstructionContainerTypeNamesInUITree, testInstructionContainerTypeInUITree.nameInUITree)
		}
		// Add TestInstructionContainerType to UI-tree model
		availableBuildingBlock[uiServer.generateUITreeNameForTestInstructionContainersHeader(domain)] = testInstructionContainerTypeNamesInUITree

		// For each 'TestInstructionType' add a list of all TestInstructions
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			var testInstructionNamesInUITree []string
			availableTestInstructions := uiServer.getAvailableTestInstructionsFromModel(domain, availableTestInstructionTypeFromModel)
			// Loop all TestInstructions and add the UI-tree name to array
			for _, availableTestInstruction := range availableTestInstructions {
				testInstructionNamesInUITree = append(testInstructionNamesInUITree, availableTestInstruction.nameInUITree)
			}
			// Add TestInstructions to UI-tree model
			availableBuildingBlock[availableTestInstructionTypeFromModel.nameInUITree] = testInstructionNamesInUITree
		}

		// For each 'TestInstructionContainerType' add a list of all TestInstructionContainers
		for _, availableTestInstructionContainerTypeFromModel := range availableTestInstructionContainerTypesFromModel {
			var testInstructionContainerNamesInUITree []string
			availableTestInstructionContainers := uiServer.getAvailableTestInstructionContainersFromModel(domain, availableTestInstructionContainerTypeFromModel)
			// Loop all TestInstructionContainers and add the UI-tree name to array
			for _, availableTestInstructionContainer := range availableTestInstructionContainers {
				testInstructionContainerNamesInUITree = append(testInstructionContainerNamesInUITree, availableTestInstructionContainer.nameInUITree)
			}
			// Add TestInstructionContainers to UI-tree model
			availableBuildingBlock[availableTestInstructionContainerTypeFromModel.nameInUITree] = testInstructionContainerNamesInUITree
		}
	}

	/*

		tree = widget.NewTreeWithStrings(list)
		tree.OnSelected = func(id string) {
			dbg.Green("Tree node selected: %s", id)

		}
		tree.OnUnselected = func(id string) {
			dbg.Red("Tree node unselected: %s", id)
		}

		tree.OpenAllBranches()

	*/

	tree = &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return availableBuildingBlock[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := availableBuildingBlock[uid]

			return ok && len(children) > 0
		},

		CreateNode: func(branch bool) fyne.CanvasObject {
			fmt.Println("CreateNode: ")
			return widget.NewLabel("Collection Widgets: ")
		},

		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			fmt.Println("UpdateNode: ", uid)
			/*
				_, ok := list[uid]
				if !ok {
					fyne.LogError("Missing tutorial panel: "+uid, nil)
					return
				}
			*/
			obj.(*widget.Label).SetText(uid) // + time.Now().String())
			fmt.Println(tree.Size())
		},

		OnSelected: func(uid string) {
			fmt.Println(uid, availableBuildingBlock[uid])
			//if t, ok := list[uid]; ok {
			//	fmt.Println(tree.Root)
			//	fmt.Println(t)

			//}
		},
	}

}

// Load all Available Building Blocks from Gui-server
func (uiServer *UIServerStruct) loadAvailableBuildingBlocksFromServer() {

	var testInstructionsAndTestContainersMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	testInstructionsAndTestContainersMessage = uiServer.grpcOut.SendGetTestInstructionsAndTestContainers("s41797")

	uiServer.loadModelWithAvailableBuildingBlocks(testInstructionsAndTestContainersMessage)

	fmt.Println(testInstructionsAndTestContainersMessage)

}

// *********** Generate Names for UI-Tree (Start)***********

// Generate UI Tree name for 'Domain', TestInstructionType, TestInstruction, TestInstructionContainerType and TestInstructionContainer for the Available Building Blocks UI-Tree
func (uiServer *UIServerStruct) generateUITreeName(node availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = node.name + " [" + node.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

/*
// Generate UI Tree name for 'Domain' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForDomain(domain availableDomainStruct) (treeName string) {

	treeName = domain.domainName + " [" + domain.domainUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

*/

// Generate UI Tree name for 'TestInstructionsHeader' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionsHeader(domain availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionsHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionContainersHeader' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionContainersHeader(domain availableBuildingBlocksForUITreeNodesStruct) (treeName string) {

	treeName = TestInstructionContainersHeader + " [" + domain.uuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

/*
// Generate UI Tree name for 'TestInstructionTypes' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionType(testInstructionType availableTestInstructionTypeStruct) (treeName string) {

	treeName = testInstructionType.testInstructionTypeName + " [" + testInstructionType.testInstructionTypeUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}



// Generate UI Tree name for 'TestInstructionContainerTypes' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionContainerType(testInstructionContainerType availableTestInstructionContainerTypeStruct) (treeName string) {

	treeName = testInstructionContainerType.testInstructionContainerTypeName + " [" + testInstructionContainerType.testInstructionContainerTypeUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructions' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstruction(testInstruction availableTestInstructionStruct) (treeName string) {

	treeName = testInstruction.testInstructionName + " [" + testInstruction.testInstructionUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionContainers' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionContainer(testInstructionContainer availableTestInstructionContainerStruct) (treeName string) {

	treeName = testInstructionContainer.testInstructionContainerName + " [" + testInstructionContainer.testInstructionContainerUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

*/

// *********** Generate Names for UI-Tree (End)***********

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (uiServer *UIServerStruct) getAvailableDomainTreeNamesFromModel() (availableDomainTreeNamesList []string) {

	availableDomains := uiServer.getAvailableDomainsFromModel()

	for _, domain := range availableDomains {
		availableDomainTreeNamesList = append(availableDomainTreeNamesList, domain.nameInUITree)
	}

	return availableDomainTreeNamesList
}

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (uiServer *UIServerStruct) getAvailableDomainsFromModel() (availableDomains []availableBuildingBlocksForUITreeNodesStruct) {

	// Extract Domain nodes from TestInstruction-map
	domainNodesInTestInstructionMap := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap
	for key := range domainNodesInTestInstructionMap {
		if key != TopNodeForAvailableDomainsMap {
			availableDomains = append(availableDomains, uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[key])
		}
	}

	// Extract Domain nodes from TestInstructionContainer-map
	domainNodesInTestInstructionContainerMap := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap
	for domainUuid := range domainNodesInTestInstructionContainerMap {
		if domainUuid != TopNodeForAvailableDomainsMap {
			_, existsInMap := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domainUuid]
			if existsInMap == false {
				availableDomains = append(availableDomains, uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[domainUuid])
			}
		}
	}

	return availableDomains
}

// Extract all 'TestInstructionTypes', per Domain, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionTypesFromModel(domain availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionTypes []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionType' for specific domain
	testInstructionTypes := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	for testInstructionType := range testInstructionTypes {
		availableTestInstructionTypes = append(availableTestInstructionTypes, uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionType])
	}

	return availableTestInstructionTypes
}

// Extract all 'TestInstructionContainerTypes', per Domain, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionContainerTypesFromModel(domain availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainerTypes []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainerType' for specific domain
	testInstructionContainerTypes := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	for testInstructionContainerType := range testInstructionContainerTypes {
		availableTestInstructionContainerTypes = append(availableTestInstructionContainerTypes, uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainerType])
	}
	return availableTestInstructionContainerTypes
}

// Extract all 'TestInstructions', per TestInstructionType, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionsFromModel(domain availableBuildingBlocksForUITreeNodesStruct, testInstructionType availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructions []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructions' for specific TestInstructionType
	testInstructionTypes := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionTypeTestInstructionRelationsMap[domain.uuid]
	testInstructions := testInstructionTypes[testInstructionType.uuid]
	for testInstruction := range testInstructions {
		availableTestInstructions = append(availableTestInstructions, uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstruction])
	}

	return availableTestInstructions
}

// Extract all 'TestInstructionContainers', per TestInstructionContainerType, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionContainersFromModel(domain availableBuildingBlocksForUITreeNodesStruct, testInstructionContainerType availableBuildingBlocksForUITreeNodesStruct) (availableTestInstructionContainers []availableBuildingBlocksForUITreeNodesStruct) {

	// Create the list of 'TestInstructionContainers' for specific TestInstructionContainerType
	testInstructionContainerTypes := uiServer.availableBuildingBlocksModel.fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap[domain.uuid]
	testInstructionContainers := testInstructionContainerTypes[testInstructionContainerType.uuid]
	for testInstructionContainer := range testInstructionContainers {
		availableTestInstructionContainers = append(availableTestInstructionContainers, uiServer.availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes[testInstructionContainer])
	}

	return availableTestInstructionContainers
}
