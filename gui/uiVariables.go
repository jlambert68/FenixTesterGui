package gui

import (
	"FenixTesterGui/commandAndRuleEngine"
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCase/testCaseUI"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

/*
type availableBuildingBlocksStruct struct {
	availableTestInstructionsAndTestInstructionContainers       *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage
	availablePinnedTestInstructionsAndTestInstructionContainers *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage
}

*/

// Constants used for Available BuildingBlocks-tree as headers
const (
	TestCaseBuildingBlocksHeader    = "TestCase Building Blocks"
	PinnedBuildingBlocksHeader      = "Pinned Building Blocks"
	AvailableBuildingBlocksHeader   = "Available Building Blocks"
	TestInstructionsHeader          = "TestInstructions"
	TestInstructionContainersHeader = "TestInstructionContainers"
)

const numberOfCharactersfromUuid = 8

const TopNodeForAvailableDomainsMap = "TOP_NODE"

type AvailableBuildingBlocksForUITreeNodesStruct struct {
	nameInUITree       string
	pinnedNameInUITree string
	uuid               string
	name               string
	BuildingBlockType  BuildingBlock
}

type uiTreeNodesNameToUuidStruct struct {
	uuid              string
	buildingBlockType BuildingBlock
}

type BuildingBlock int

const (
	Undefined BuildingBlock = iota
	TestInstruction
	TestInstructionContainer
)

/*
	type availableDomainStruct struct {
		domainNameInUITree string
		domainUuid         string
		domainName         string
	}

	type availableTestInstructionTypeStruct struct {
		testInstructionTypeNameInUITree string
		domainUuid                      string
		domainName                      string
		testInstructionTypeUuid         string
		testInstructionTypeName         string
	}

	type availableTestInstructionContainerTypeStruct struct {
		testInstructionContainerTypeNameInUITree string
		domainUuid                               string
		domainName                               string
		testInstructionContainerTypeUuid         string
		testInstructionContainerTypeName         string
	}
*/
type availableTestInstructionStruct struct {
	testInstructionNameInUITree string
	domainUuid                  string
	domainName                  string
	testInstructionTypeUuid     string
	testInstructionTypeName     string
	testInstructionUuid         string
	testInstructionName         string
}

type availableTestInstructionContainerStruct struct {
	testInstructionContainerNameInUITree string
	domainUuid                           string
	domainName                           string
	testInstructionContainerTypeUuid     string
	testInstructionContainerTypeName     string
	testInstructionContainerUuid         string
	testInstructionContainerName         string
}

type AvailableBuildingBlocksModelStruct struct {
	// + TestCase Building Blocks
	//    + Pinned Building Blocks
	//       TestInstruction 1 [c107bdd9] - Pinned
	//       TestInstructionContainer 1 [f107bdd9] - Pinned

	//    +Available Building Blocks
	//       +Domain Name 1 [a107bdd9]
	//          +TestInstructions [a107bdd9]
	//              +TestInstructionType 1 [b107bdd9]
	//                 TestInstruction 1 [c107bdd9]
	//                 TestInstruction 2 [d107bdd9]
	//              +TestInstructionType 2 [b107bdd9]
	//                 TestInstruction 21 [c107bdd9]
	//                 TestInstruction 22 [d107bdd9]
	//         +TestInstructionContainers [a107bdd9]
	//            +TestInstructionContainerType 1 [e107bdd9]
	//                TestInstructionContainer 1 [f107bdd9]
	//                TestInstructionContainer 2 [g107bdd9]
	//             +TestInstructionContainerType 2 [e107bdd9]
	//                TestInstructionContainer 21 [f107bdd9]
	//                TestInstructionContainer 22 [g107bdd9]

	logger                                                                     *logrus.Logger
	fenixGuiBuilderServerAddressToDial                                         string
	fullDomainTestInstructionTypeTestInstructionRelationsMap                   map[string]map[string]map[string]availableTestInstructionStruct
	fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap map[string]map[string]map[string]availableTestInstructionContainerStruct
	AvailableBuildingBlocksForUITreeNodes                                      map[string]AvailableBuildingBlocksForUITreeNodesStruct // map[uuid]AvailableBuildingBlocksForUITreeNodesStruct
	pinnedBuildingBlocksForUITreeNodes                                         map[string]uiTreeNodesNameToUuidStruct                 //map[nameInUITree]uiTreeNodesNameToUuidStruct
	grpcOut                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct
	availableBuildingBlockModelSuitedForFyneTreeView                           map[string][]string
	allBuildingBlocksTreeNameToUuid                                            map[string]uiTreeNodesNameToUuidStruct
	clickedNodeName                                                            string

	allImmatureTestInstructionsBuildingBlocks         map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage          // Map of all Available Building Blocks regarding TestInstructions, Immature UUID as Map-key
	allImmatureTestInstructionContainerBuildingBlocks map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage // Map of all Available Building Blocks regarding TestInstructionContainers, Immature UUID as Map-key

	// temp Variable for Current TestCase Textual Structure
	currentTestCaseTextualStructureSimple   binding.String
	currentTestCaseTextualStructureComplex  binding.String
	currentTestCaseTextualStructureExtended binding.String

	// Domains that can own a TestCase for this user
	DomainsThatCanOwnTheTestCase []*fenixGuiTestCaseBuilderServerGrpcApi.DomainsThatCanOwnTheTestCaseMessage

	// URLs to repositories with Template for this user
	TemplateRepositoryApiUrls []*fenixGuiTestCaseBuilderServerGrpcApi.RepositoryApiUrlResponseMessage

	// The TestCaseMetaData for Domain that the user has acces to when creating TestCasesMapPtr
	TestCaseMetaDataForDomains []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage

	// TestData that this User can use
	TestData []*fenixGuiTestCaseBuilderServerGrpcApi.TestDataFromOneSimpleTestDataAreaFileMessage

	// Users available ExecutionDomains to be used with Fenix-created TestInstructions that should be sent to other Domain then Fenix
	executionDomainsThatCanReceiveDirectTargetedTestInstructionsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
									ExecutionDomainsThatCanReceiveDirectTargetedTestInstructionsMessage
}
type GlobalUIServerStruct struct {
	logger                             *logrus.Logger
	fenixGuiBuilderServerAddressToDial string
	//fenixApp                            fyne.App
	//tree                               *widget.Label
	//content                            *widget.Entry
	//grpcOut                            grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct

	//availableBuildingBlocks            availableBuildingBlocksStruct
	//AvailableBuildingBlocksModel       AvailableBuildingBlocksModelStruct
	//AvailableBuildingBlocksModel AvailableBuildingBlocksModelStruct
}

var localUIServer UIServerStruct

type UIServerStruct struct {
	logger                             *logrus.Logger
	fenixApp                           fyne.App
	fenixMasterWindow                  fyne.Window
	tree                               *widget.Label
	content                            *widget.Entry
	grpcOut                            grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct
	fenixGuiBuilderServerAddressToDial string
	//availableBuildingBlocks            availableBuildingBlocksStruct
	//AvailableBuildingBlocksModel       AvailableBuildingBlocksModelStruct
	AvailableBuildingBlocksModel AvailableBuildingBlocksModelStruct
	testCasesModel               testCaseModel.TestCasesModelsStruct
	commandAndRuleEngine         commandAndRuleEngine.CommandAndRuleEngineObjectStruct
	//subSystemsCrossReferences    SubSystemsCrossReferencesStruct
	testCasesUiModel testCaseUI.TestCasesUiModelStruct
}

/*
type SubSystemsCrossReferencesStruct struct {
	//AvailableBuildingBlocksModelReference *AvailableBuildingBlocksModelStruct
	//TestCasesModelReference               *testCaseModel.TestCasesModelsStruct
	//CommandAndRuleEnginReference          *commandAndRuleEngine.CommandAndRuleEngineObjectStruct
	GrpcOutReference *grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct
}


*/
/*
// The testCaseModel for available Building Blocks used within the Tree-view in GUI
type availableBuildingBlocksStruct struct {
	logger                       *logrus.Logger

	fenixGuiBuilderServerAddressToDial string
	AvailableBuildingBlocksModel AvailableBuildingBlocksModelStruct
}

*/

//var myUIServer UIServerStruct
