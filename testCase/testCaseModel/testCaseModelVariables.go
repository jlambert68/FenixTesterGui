package testCaseModel

import (
	"FenixTesterGui/grpc_out"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"time"
)

const numberOfCharactersfromUuid = 8

// Node colors in RGBA-format
const (
	nodeColor_Bond_B0         = "#404040FF"
	nodeColor_Bond_B1         = "#606060FF"
	nodeColor_Swappeble_Bonds = "#989898FF"
	nodeColor_TI_TIC          = "#888888FF"
	nodeColor_X_Bonds         = "#505050FF"
	nodeColor_B10X_Bonds      = "#505050FF"
	nodeColor_TIx_TICx        = "#F8F8F8FF"
)

type TestCasesModelsStruct struct {
	TestCases   map[string]TestCaseModelStruct // Holds the Model for all the TestCase-models
	CurrentUser string                         // Current logged-in user TODO Put this in a more global structure
	//subSystemsCrossReferences *gui.SubSystemsCrossReferencesStruct
	GrpcOutReference                              *grpc_out.GRPCOutStruct
	CurrentActiveTestCaseUuid                     string                                                                                                                                               // The TestCase that should be worked on both by the model and UI
	AvailableBondsMap                             map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage // A copy of available Bonds //TODO should be placed in one common object
	AvailableImmatureTestInstructionsMap          map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage
	AvailableImmatureTestInstructionContainersMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage
	ImmatureTestInstructionAttributesMap          map[string]map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage // map[TestInstructionUuid]map[TestInstructionAttributeUuid]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage
	ImmatureDropZonesDataMap                      map[string]ImmatureDropZoneDataMapStruct                                                                                             // map[DropZoneUuid]ImmatureDropZoneDataMapStruct

	//AvailableBuildingBlocksModel                  *gui.AvailableBuildingBlocksModelStruct

}

type ImmatureDropZoneDataMapStruct struct {
	DropZoneUuid                               string                                                                                                                                                            `protobuf:"bytes,1,opt,name=DropZoneUuid,proto3" json:"DropZoneUuid,omitempty"`                                                       // A DropZone, UUID, for the TestInstruction
	DropZoneName                               string                                                                                                                                                            `protobuf:"bytes,2,opt,name=DropZoneName,proto3" json:"DropZoneName,omitempty"`                                                       // A DropZone, Name, for the TestInstruction
	DropZoneDescription                        string                                                                                                                                                            `protobuf:"bytes,3,opt,name=DropZoneDescription,proto3" json:"DropZoneDescription,omitempty"`                                         // Description of the DropZone
	DropZoneMouseOver                          string                                                                                                                                                            `protobuf:"bytes,4,opt,name=DropZoneMouseOver,proto3" json:"DropZoneMouseOver,omitempty"`                                             // The mouse over text for the DropZone
	DropZoneColor                              string                                                                                                                                                            `protobuf:"bytes,5,opt,name=DropZoneColor,proto3" json:"DropZoneColor,omitempty"`                                                     // The color used for presenting the DropsZone, e.g. #FAF437
	DropZonePreSetTestInstructionAttributesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionInformationMessage_AvailableDropZoneMessage_DropZonePreSetTestInstructionAttributeMessage `protobuf:"bytes,6,rep,name=DropZonePreSetTestInstructionAttributes,proto3" json:"DropZonePreSetTestInstructionAttributes,omitempty"` // A list of the attributes and their pre-set values

}

type TestCaseModelStruct struct {
	LastLoadedTestCaseModelGRPCMessage         fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage
	FirstElementUuid                           string
	TestCaseModelMap                           map[string]MatureTestCaseModelElementStruct
	TextualTestCaseRepresentationSimpleStack   []string
	TextualTestCaseRepresentationComplexStack  []string
	TextualTestCaseRepresentationExtendedStack []string
	CommandStack                               []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
	LastSavedCommandStack                      lastSavedCommandStack
	CopyBuffer                                 ImmatureElementStruct
	CutBuffer                                  MatureElementStruct
	CutCommandInitiated                        bool
	LocalTestCaseMessage                       LocalTestCaseMessageStruct
	testCaseModelAdaptedForUiTree              map[string][]TestCaseModelAdaptedForUiTreeDataStruct // Model used for Creating the Tree-view for the TestCase-model
	CurrentSelectedTestCaseElement             currentSelectedTestCaseElementStruct
	MatureTestInstructionMap                   map[string]MatureTestInstructionStruct
}

type MatureTestInstructionStruct struct {
	MatureBasicTestInstructionInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_MatureBasicTestInstructionInformationMessage `protobuf:"bytes,1,opt,name=MatureBasicTestInstructionInformation,proto3" json:"MatureBasicTestInstructionInformation,omitempty"` // The Basic information for the Matures TestInstruction
	CreatedAndUpdatedInformation          *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_CreatedAndUpdatedInformationMessage          `protobuf:"bytes,2,opt,name=CreatedAndUpdatedInformation,proto3" json:"CreatedAndUpdatedInformation,omitempty"`                   // Information regarding who did what and when
	TestInstructionAttributesList         map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage   `protobuf:"bytes,3,rep,name=TestInstructionAttributesList,proto3" json:"TestInstructionAttributesList,omitempty"`                 // All attributes that belongs to the TestInstruction

}

type currentSelectedTestCaseElementStruct struct {
	CurrentSelectedTestCaseElementUuid string
	CurrentSelectedTestCaseElementName string
}

type MatureTestCaseModelElementStruct struct {
	MatureTestCaseModelElementMessage  fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
	MatureTestCaseModelElementMetaData MatureTestCaseModelElementMetaDataStruct
	//MatureTestCaseModelElementAttributes MatureTestCaseModelElementAttributesStruct
}

type MatureTestCaseModelElementMetaDataStruct struct {
	ChosenDropZoneUuid        string
	ChosenDropZoneColorString string
}

// MatureTestCaseModelElementAttributesStruct - AttributeUuid as map-key
type MatureTestCaseModelElementAttributesStruct struct {
	AttributesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
}

type TestCaseModelAdaptedForUiTreeDataStruct struct {
	Uuid                     string
	OriginalUuid             string
	NodeName                 string
	NodeColor                string
	TestInstructionTypeColor string
	NodeTypeEnum             fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum // TestCaseModelElementTypeEnum fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_name["int32"]
	CanBeDeleted             bool
	CanBeSwappedOut          bool
}

type lastSavedCommandStack struct {
	savedTimeStamp time.Time
	userId         string
	commandStack   []fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage_TestCaseModelCommandMessage
}

// ImmatureElementStruct
// Used for Swapping in a New Element-structure or the Copy-Buffer
type ImmatureElementStruct struct {
	FirstElementUuid    string
	ChosenDropZoneUuid  string
	ChosenDropZoneName  string
	ChosenDropZoneColor string
	ImmatureElementMap  map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage
}

// MatureElementStruct
// Used when converting an Immature Element-structure into a Mature-structure to be used in the TestCase. Cut-buffer also use this structure
type MatureElementStruct struct {
	FirstElementUuid    string
	ChosenDropZoneUuid  string
	ChosenDropZoneName  string
	ChosenDropZoneColor string
	MatureElementMap    map[string]fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage
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
