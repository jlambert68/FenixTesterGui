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
	for _, domain := range uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] {
		// For each domain add TestInstructionHeaderName and TestInstructionContainerHeaderName
		availableBuildingBlock[domain.domainNameInUITree] = []string{
			uiServer.generateUITreeNameForTestInstructionsHeader(domain),
			uiServer.generateUITreeNameForTestInstructionContainersHeader(domain)}

		// For 'TestInstructionHeaderName' add a list of all TestInstructionTypes
		availableTestInstructionTypesFromModel := uiServer.getAvailableTestInstructionTypesFromModel(domain)
		var testInstructionTypeNamesInUITree []string
		// Loop all TestInstructionTypes and extract UI-tree name
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			testInstructionTypeNamesInUITree = append(testInstructionTypeNamesInUITree, availableTestInstructionTypeFromModel.testInstructionTypeNameInUITree)
		}
		// Add TestInstructionType to UI-tree model
		availableBuildingBlock[uiServer.generateUITreeNameForTestInstructionsHeader(domain)] = testInstructionTypeNamesInUITree

		// For 'TestInstructionContainerHeaderName' add a list of all TestInstructionContainerTypes
		availableTestInstructionContainerTypesFromModel := uiServer.getAvailableTestInstructionContainerTypesFromModel(domain)
		var testInstructionContainerTypeNamesInUITree []string
		// Loop all TestInstructionContainerTypes and extract UI-tree name
		for _, testInstructionContainerTypeInUITree := range availableTestInstructionContainerTypesFromModel {
			testInstructionContainerTypeNamesInUITree = append(testInstructionContainerTypeNamesInUITree, testInstructionContainerTypeInUITree.testInstructionContainerTypeNameInUITree)
		}
		// Add TestInstructionContainerType to UI-tree model
		availableBuildingBlock[uiServer.generateUITreeNameForTestInstructionContainersHeader(domain)] = testInstructionContainerTypeNamesInUITree

		// For each 'TestInstructionType' add a list of all TestInstructions
		for _, availableTestInstructionTypeFromModel := range availableTestInstructionTypesFromModel {
			var testInstructionNamesInUITree []string
			availableTestInstructions := uiServer.getAvailableTestInstructionsFromModel(availableTestInstructionTypeFromModel)
			// Loop all TestInstructions and add the UI-tree name to array
			for _, availableTestInstruction := range availableTestInstructions {
				testInstructionNamesInUITree = append(testInstructionNamesInUITree, availableTestInstruction.testInstructionNameInUITree)
			}
			// Add TestInstructions to UI-tree model
			availableBuildingBlock[availableTestInstructionTypeFromModel.testInstructionTypeNameInUITree] = testInstructionNamesInUITree
		}

		// For each 'TestInstructionContainerType' add a list of all TestInstructionContainers
		for _, availableTestInstructionContainerTypeFromModel := range availableTestInstructionContainerTypesFromModel {
			var testInstructionContainerNamesInUITree []string
			availableTestInstructionContainers := uiServer.getAvailableTestInstructionContainersFromModel(availableTestInstructionContainerTypeFromModel)
			// Loop all TestInstructionContainers and add the UI-tree name to array
			for _, availableTestInstructionContainer := range availableTestInstructionContainers {
				testInstructionContainerNamesInUITree = append(testInstructionContainerNamesInUITree, availableTestInstructionContainer.testInstructionContainerNameInUITree)
			}
			// Add TestInstructionContainers to UI-tree model
			availableBuildingBlock[availableTestInstructionContainerTypeFromModel.testInstructionContainerTypeNameInUITree] = testInstructionContainerNamesInUITree
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

	var fenixGuiTestCaseBuilderServerGrpcApi *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	fenixGuiTestCaseBuilderServerGrpcApi = uiServer.grpcOut.SendGetTestInstructionsAndTestContainers("s41797")

	uiServer.loadModelWithAvailableBuildingBlocks(fenixGuiTestCaseBuilderServerGrpcApi)

	fmt.Println(fenixGuiTestCaseBuilderServerGrpcApi)

}

// *********** Generate Names for UI-Tree (Start)***********

// Generate UI Tree name for 'Domain' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForDomain(domain availableDomainStruct) (treeName string) {

	treeName = domain.domainName + " [" + domain.domainUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionsHeader' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionsHeader(domain availableDomainStruct) (treeName string) {

	treeName = TestInstructionsHeader + " [" + domain.domainUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

// Generate UI Tree name for 'TestInstructionContainersHeader' for Available Building Blocks
func (uiServer *UIServerStruct) generateUITreeNameForTestInstructionContainersHeader(domain availableDomainStruct) (treeName string) {

	treeName = TestInstructionContainersHeader + " [" + domain.domainUuid[0:numberOfCharactersfromUuid-1] + "]"

	return treeName
}

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

// *********** Generate Names for UI-Tree (End)***********

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (uiServer *UIServerStruct) getAvailableDomainTreeNamesFromModel() (availableDomainTreeNamesList []string) {

	for _, domain := range uiServer.availableBuildingBlocksModel.availableDomains[TopNodeForAvailableDomainsMap] {
		availableDomainTreeNamesList = append(availableDomainTreeNamesList, domain.domainNameInUITree)
	}

	return availableDomainTreeNamesList
}

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
/*
func (uiServer *UIServerStruct) getAvailableDomainsFromModel() (availableDomains []availableDomainStruct) {

	//availableDomains = maps.Keys(uiServer.availableBuildingBlocksModel.availableDomains)

	return availableDomains
}
*/

// Extract all 'TestInstructionTypes', per Domain, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionTypesFromModel(domain availableDomainStruct) (availableTestInstructionTypesList []availableTestInstructionTypeStruct) {

	// Create the list of 'TestInstructionTypeTreeNames' for specific domain
	for _, domainsTestInstructionType := range uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[domain.domainUuid] {
		// Add each TestInstructionType to list
		availableTestInstructionTypesList = append(availableTestInstructionTypesList, domainsTestInstructionType)
	}

	return availableTestInstructionTypesList
}

// Extract all 'TestInstructionContainerTypes', per Domain, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionContainerTypesFromModel(domain availableDomainStruct) (availableTestInstructionContainerTypesList []availableTestInstructionContainerTypeStruct) {

	// Create the list of 'TestInstructionContainerTypeTreeNames' for specific domain
	for _, domainsTestInstructionContainerType := range uiServer.availableBuildingBlocksModel.domainsTestInstructionContainerTypes[domain.domainUuid] {
		// Add each TestInstructionType to list
		availableTestInstructionContainerTypesList = append(availableTestInstructionContainerTypesList, domainsTestInstructionContainerType)
	}

	return availableTestInstructionContainerTypesList
}

// Extract all 'TestInstructions', per TestInstructionType, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionsFromModel(testInstructionType availableTestInstructionTypeStruct) (availableTestInstructions []availableTestInstructionStruct) {

	// Create the list of 'TestInstructions' for specific TestInstructionType
	for _, testInstruction := range uiServer.availableBuildingBlocksModel.testInstructionTypesTestInstructions[testInstructionType.testInstructionTypeUuid] {
		// Add each TestInstruction to list
		availableTestInstructions = append(availableTestInstructions, testInstruction)
	}

	return availableTestInstructions
}

// Extract all 'TestInstructionContainers', per TestInstructionContainerType, with Names suited for Tree-model
func (uiServer *UIServerStruct) getAvailableTestInstructionContainersFromModel(testInstructionContainerType availableTestInstructionContainerTypeStruct) (availableTestInstructionContainers []availableTestInstructionContainerStruct) {

	// Create the list of 'TestInstructionContainers' for specific TestInstructionContainerType
	for _, testInstructionContainer := range uiServer.availableBuildingBlocksModel.testInstructionContainerTypesTestInstructionsContainers[testInstructionContainerType.testInstructionContainerTypeUuid] {
		// Add each TestInstructionContainer to list
		availableTestInstructionContainers = append(availableTestInstructionContainers, testInstructionContainer)
	}

	return availableTestInstructionContainers
}
