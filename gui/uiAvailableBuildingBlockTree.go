package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
	"time"
)

//______________________________________________________________________________

var list map[string][]string
var tree *widget.Tree

func (uiServer *UIServerStruct) makeTreeUI() {
	list = map[string][]string{
		"":  {"A"},
		"A": {"B", "D"},
		"B": {"C"},
		"C": {"abc"},
		"D": {"E"},
		"E": {"F", "G"},
	}

	availableBuildingBlock := map[string][]string{
		"":                            {TestCaseBuildingBlocksHeader},
		TestCaseBuildingBlocksHeader:  {PinnedBuildingBlocksHeader, AvailableBuildingBlocksHeader},
		PinnedBuildingBlocksHeader:    {"Nothing here, yet"},
		AvailableBuildingBlocksHeader: uiServer.getAvailableDomainTreeNamesFromModel(),
	}

	// Generate TestInstructionTypes per Domain
	for _, domainTreeName := range uiServer.getAvailableDomainTreeNamesFromModel() {
		availableBuildingBlock[domainTreeName] = uiServer.getAvailableTestInstructionTypeTreeNamesFromModel(domainTreeName)

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
			return list[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := list[uid]

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
			obj.(*widget.Label).SetText(uid + time.Now().String())
			fmt.Println(tree.Size())
		},

		OnSelected: func(uid string) {
			fmt.Println(uid, list[uid])
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

	fmt.Println(fenixGuiTestCaseBuilderServerGrpcApi)

}

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

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (uiServer *UIServerStruct) getAvailableDomainTreeNamesFromModel() (availableDomainTreeNamesList []string) {

	availableDomainTreeNamesList = maps.Keys(uiServer.availableBuildingBlocksModel.availableDomains)

	return availableDomainTreeNamesList
}

// Extract all 'Domains', with Names suited for Tree-model, for the model tha underpins the UI Tree for Available Building Blocks
func (uiServer *UIServerStruct) getAvailableTestInstructionTypeTreeNamesFromModel(domainTreeName string) (availableTestInstructionTypeTreeNamesList []string) {

	// Extract Domain UUID
	domain := uiServer.availableBuildingBlocksModel.availableDomains[domainTreeName]
	if len(domain) != 1 {
		uiServer.logger.WithFields(logrus.Fields{
			"id": "338f2048-d4c1-4ee3-9efb-38680de44f32",
		}).Fatalln("Expected 'one' domain, but found '" + string(len(domain)) + "' domains")
	}
	domainUuid := domain[0].domainUuid

	availableTestInstructionTypeTreeNamesList = maps.Keys(uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes)

	for _, domainsTestInstructionType := range availableTestInstructionTypeTreeNamesList {
		uiServer.availableBuildingBlocksModel.domainsTestInstructionTypes[domainsTestInstructionType]
	}

	return availableTestInstructionTypeTreeNamesList
}
