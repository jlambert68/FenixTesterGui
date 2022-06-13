package gui

import (
	"FenixTesterGui/grpc_out"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

type availableBuildingBlocksStruct struct {
	availableTestInstructionsAndTestInstructionContainers       *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage
	availablePinnedTestInstructionsAndTestInstructionContainers *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage
}

// Constants used for Available BuildingBlocks-tree as headers
const (
	TestCaseBuildingBlocksHeader  = "TestCase Building Blocks"
	PinnedBuildingBlocksHeader    = "Pinned Building Blocks"
	AvailableBuildingBlocksHeader = "Available Building Blocks"
)

type availableBuildingBlocksModelStruct struct {
	// + TestCase Building Blocks
	//    + Pinned Building Blocks
	//       TestInstruction 1 [c107bdd9] - Pinned
	//       TestInstructionContainer 1 [f107bdd9] - Pinned

	//    +Available Building Blocks
	//       +Domain Name 1 [a107bdd9]
	//          +TestInstructionType 1 [b107bdd9]
	//             TestInstruction 1 [c107bdd9]
	//             TestInstruction 2 [d107bdd9]
	//          +TestInstructionType 2 [b107bdd9]
	//             TestInstruction 21 [c107bdd9]
	//             TestInstruction 22 [d107bdd9]
	//          +TestInstructionContainerType 1 [e107bdd9]
	//             TestInstructionContainer 1 [f107bdd9]
	//             TestInstructionContainer 2 [g107bdd9]
	//          +TestInstructionContainerType 2 [e107bdd9]
	//             TestInstructionContainer 21 [f107bdd9]
	//             TestInstructionContainer 22 [g107bdd9]
	/*
	   	list = map[string][]string{
	   	"":  {TestCaseBuildingBlocksHeader},
	   	TestCaseBuildingBlocksHeader: {PinnedBuildingBlocksHeader, AvailableBuildingBlocksHeader},
	   	PinnedBuildingBlocksHeader: [xxx,xxx]
	    AvailableBuildingBlocksHeader: {[]TestInstructionTypes},
	    AvailableBuildingBlocksHeader: {[]TestInstructionTypes},
	   	"C": {"abc"},
	   	"D": {"E"},
	   	"E": {"F", "G"},
	   }
	*/

}

type UIServerStruct struct {
	logger                             *logrus.Logger
	fyneApp                            fyne.App
	tree                               *widget.Label
	content                            *widget.Entry
	grpcOut                            grpc_out.GRPCOutStruct
	fenixGuiBuilderServerAddressToDial string
	availableBuildingBlocks            availableBuildingBlocksStruct
}

//var myUIServer UIServerStruct
