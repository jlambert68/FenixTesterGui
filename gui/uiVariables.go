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
	TestCaseBuildingBlocksHeader    = "TestCase Building Blocks"
	PinnedBuildingBlocksHeader      = "Pinned Building Blocks"
	AvailableBuildingBlocksHeader   = "Available Building Blocks"
	TestInstructionsHeader          = "TestInstructions"
	TestInstructionContainersHeader = "TestInstructionContainers"
)

const numberOfCharactersfromUuid = 8

const TopNodeForAvailableDomainsMap = "TOP_NODE"

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
	/*
		   	list = map[string][]string{
		   	"":  {TestCaseBuildingBlocksHeader},
		   	TestCaseBuildingBlocksHeader: {PinnedBuildingBlocksHeader, AvailableBuildingBlocksHeader},
		   	PinnedBuildingBlocksHeader: {[]TestInstructions, []TestInstructionContainer}


		    AvailableBuildingBlocksHeader: {[]Domains},
			[0]Domains: {TestInstructionsHeader, TestInstructionContainersHeader},
			TestInstructionsHeader: {[]TestInstructionTypes},
			[0]TestInstructionTypes: {[]TestInstructions},
			[1]TestInstructionTypes: {[]TestInstructions},
			TestInstructionContainersHeader: {[]TestInstructionContainerTypes},
			[0]TestInstructionContainerTypes: {[]TestInstructionContainer},
			[1]TestInstructionContainerTypes: {[]TestInstructionContainer},

			[1]Domains: {TestInstructionsHeader, TestInstructionContainersHeader},
			TestInstructionsHeader: {[]TestInstructionTypes},
			[0]TestInstructionTypes: {[]TestInstructions},
			[1]TestInstructionTypes: {[]TestInstructions},
			TestInstructionContainersHeader: {[]TestInstructionContainerTypes},
			[0]TestInstructionContainerTypes: {[]TestInstructionContainer},
			[1]TestInstructionContainerTypes: {[]TestInstructionContainer},

			Key Name construction
			Domains = Domain Name [f107bdd9]
			TestInstructionsHeader = TestInstructionsHeader [Domain UUID]
			TestInstructionContainersHeader = TestInstructionContainersHeader [Domain UUID]
			TestInstructionTypes = TestInstructionType Name [TestInstructionType UUID]
			TestInstructionTypes = TestInstructionContainerType Name [TestInstructionContainerType UUID]
			TestInstructions = TestInstruction Name [TestInstruction UUID]
			TestInstructionContainers = TestInstructionContainer Name [TestInstructionContainer UUID]



		   }
	*/

	fullDomainTestInstructionTypeTestInstructionRelationsMap                   map[string]map[string]map[string]availableTestInstructionStruct
	fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap map[string]map[string]map[string]availableTestInstructionContainerStruct

	availableDomains                                        map[string][]availableDomainStruct                       // map[#DomainUUID#]availableDomainStruct
	domainsTestInstructionTypes                             map[string][]availableTestInstructionTypeStruct          // map[#DomainUUID#][]availableTestInstructionTypeStruct
	domainsTestInstructionContainerTypes                    map[string][]availableTestInstructionContainerTypeStruct // map[#DomainUUID#][]availableTestInstructionContainerTypeStruct
	testInstructionTypesTestInstructions                    map[string][]availableTestInstructionStruct              // map[#TestInstructionTypeUUID#][]availableTestInstructionStruct
	testInstructionContainerTypesTestInstructionsContainers map[string][]availableTestInstructionContainerStruct     // map[#TestInstructionContainerTypeUUID#"][]availableTestInstructionContainerTypeStruct

}

type UIServerStruct struct {
	logger                             *logrus.Logger
	fyneApp                            fyne.App
	tree                               *widget.Label
	content                            *widget.Entry
	grpcOut                            grpc_out.GRPCOutStruct
	fenixGuiBuilderServerAddressToDial string
	availableBuildingBlocks            availableBuildingBlocksStruct
	availableBuildingBlocksModel       availableBuildingBlocksModelStruct
}

//var myUIServer UIServerStruct
