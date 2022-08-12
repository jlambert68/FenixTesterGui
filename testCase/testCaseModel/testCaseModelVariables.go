package testCaseModel

import (
	"FenixTesterGui/grpc_out"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"time"
)

const numberOfCharactersfromUuid = 8

type TestCasesModelsStruct struct {
	TestCases   map[string]TestCaseModelStruct // Holds the Model for all the TestCase-models
	CurrentUser string                         // Current logged-in user TODO Put this in a more global structure
	//subSystemsCrossReferences *gui.SubSystemsCrossReferencesStruct
	GrpcOutReference          *grpc_out.GRPCOutStruct
	CurrentActiveTestCaseUuid string // The TestCase that should be worked on both by the model and UI

}

type TestCaseModelStruct struct {
	LastLoadedTestCaseModelGRPCMessage         fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	FirstElementUuid                           string
	TestCaseModelMap                           map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
	TextualTestCaseRepresentationSimpleStack   []string
	TextualTestCaseRepresentationComplexStack  []string
	TextualTestCaseRepresentationExtendedStack []string
	CommandStack                               []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
	LastSavedCommandStack                      lastSavedCommandStack
	CopyBuffer                                 ImmatureElementStruct
	CutBuffer                                  MatureElementStruct
	CutCommandInitiated                        bool
	LocalTestCaseMessage                       LocalTestCaseMessageStruct
}

type lastSavedCommandStack struct {
	savedTimeStamp time.Time
	userId         string
	commandStack   []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
}

// ImmatureElementStruct
// Used for Swapping in a New Element-structure or the Copy-Buffer
type ImmatureElementStruct struct {
	FirstElementUuid   string
	ImmatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage
}

// MatureElementStruct
// Used when converting an Immature Element-structure into a Mature-structure to be used in the TestCase. Cut-buffer also use this structure
type MatureElementStruct struct {
	FirstElementUuid string
	MatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}

// LocalTestCaseMessageStruct
// A message holding one TestCase in
type LocalTestCaseMessageStruct struct {
	BasicTestCaseInformationMessageNoneEditableInformation fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage_NonEditableBasicInformationMessage // All Basic information that can be shown in GUI but can't be changed by the user
	BasicTestCaseInformationMessageEditableInformation     fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage_EditableBasicInformationMessage    // All Basic information that can be shown in GUI and can be changed by the user
	CreatedAndUpdatedInformation                           fenixGuiTestCaseBuilderServerGrpcApi.TestCaseMessage_CreatedAndUpdatedInformationMessage                // Information regarding who did what and when

	//TestCaseModel                fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage                                // Hold the model of how the TestCase is constructed
	//TestCaseMetaData             fenixGuiTestCaseBuilderServerGrpcApi.TestCaseMetaDataMessage                             // Holds the metadata for the TestCase, set by the user to classify the TestCase
	//TestCaseFiles                fenixGuiTestCaseBuilderServerGrpcApi.TestCaseFilesMessage                               // All files connected to the TestCase
}

/*
type MatureElementStruct struct {
	FirstElementUuid string
	MatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
}
*/

/*
type ImmatureElementStruct struct {
	FirstElementUuid   string
	ImmatureElementMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage
}

*/

const NotApplicable = "N/A"
