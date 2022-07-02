package gui

import (
	"FenixTesterGui/grpc_out"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
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

type availableBuildingBlocksForUITreeNodesStruct struct {
	nameInUITree       string
	pinnedNameInUITree string
	uuid               string
	name               string
	buildingBlockType  buildingBlock
}

type uiTreeNodesNameToUuidStruct struct {
	uuid              string
	buildingBlockType buildingBlock
}

type buildingBlock int

const (
	Undefined buildingBlock = iota
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

type availableBuildingBlocksModelStruct struct {
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
	availableBuildingBlocksForUITreeNodes                                      map[string]availableBuildingBlocksForUITreeNodesStruct
	pinnedBuildingBlocksForUITreeNodes                                         map[string]uiTreeNodesNameToUuidStruct //map[nameInUITree]uiTreeNodesNameToUuidStruct
	grpcOut                                                                    grpc_out.GRPCOutStruct
	availableBuildingBlockModelSuitedForFyneTreeView                           map[string][]string
	allBuildingBlocksTreeNameToUuid                                            map[string]uiTreeNodesNameToUuidStruct
}

type GlobalUIServerStruct struct {
	logger                             *logrus.Logger
	fenixGuiBuilderServerAddressToDial string
	//fyneApp                            fyne.App
	//tree                               *widget.Label
	//content                            *widget.Entry
	//grpcOut                            grpc_out.GRPCOutStruct

	//availableBuildingBlocks            availableBuildingBlocksStruct
	//availableBuildingBlocksModel       availableBuildingBlocksModelStruct
	//availableBuildingBlocksModel availableBuildingBlocksModelStruct
}

var localUIServer UIServerStruct

type UIServerStruct struct {
	logger                             *logrus.Logger
	fyneApp                            fyne.App
	tree                               *widget.Label
	content                            *widget.Entry
	grpcOut                            grpc_out.GRPCOutStruct
	fenixGuiBuilderServerAddressToDial string
	//availableBuildingBlocks            availableBuildingBlocksStruct
	//availableBuildingBlocksModel       availableBuildingBlocksModelStruct
	availableBuildingBlocksModel availableBuildingBlocksModelStruct
}

/*
// The model for available Building Blocks used within the Tree-view in GUI
type availableBuildingBlocksStruct struct {
	logger                       *logrus.Logger

	fenixGuiBuilderServerAddressToDial string
	availableBuildingBlocksModel availableBuildingBlocksModelStruct
}

*/

//var myUIServer UIServerStruct
