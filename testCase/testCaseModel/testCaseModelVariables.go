package testCaseModel

import (
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"FenixTesterGui/importFilesFromGitHub"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"regexp"
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
	GrpcOutReference                              *grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct
	CurrentActiveTestCaseUuid                     string                                                                                                                                               // The TestCase that should be worked on both by the model and UI
	AvailableBondsMap                             map[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureBondsMessage_ImmatureBondMessage // A copy of available Bonds //TODO should be placed in one common object
	AvailableImmatureTestInstructionsMap          map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage
	AvailableImmatureTestInstructionContainersMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage
	ImmatureTestInstructionAttributesMap          map[string]map[string]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage // map[TestInstructionUuid]map[TestInstructionAttributeUuid]*fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestInstructionAttributesMessage_TestInstructionAttributeMessage
	ImmatureDropZonesDataMap                      map[string]ImmatureDropZoneDataMapStruct                                                                                             // map[DropZoneUuid]ImmatureDropZoneDataMapStruct
	DomainsThatCanOwnTheTestCaseMap               map[string]*DomainThatCanOwnTheTestCaseStruct
	TemplateRepositoryApiUrlMap                   map[string]*fenixGuiTestCaseBuilderServerGrpcApi.RepositoryApiUrlResponseMessage
	TestCasesThatCanBeEditedByUserMap             map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage
	//TestCasesThatCanBeEditedByUserSlice           []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage

	//AvailableBuildingBlocksModel                  *gui.AvailableBuildingBlocksModelStruct

}

type DomainThatCanOwnTheTestCaseStruct struct {
	DomainUuid           string
	DomainName           string
	DomainNameShownInGui string
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
	CurrentSelectedTestCaseElement             CurrentSelectedTestCaseElementStruct
	MatureTestInstructionMap                   map[string]MatureTestInstructionStruct
	MatureTestInstructionContainerMap          map[string]MatureTestInstructionContainerStruct

	AttributesList                           *AttributeStructSliceReferenceType
	ThisIsANewTestCase                       bool
	TestCaseHash                             string
	TestCaseHashWhenTestCaseWasSavedOrLoaded string
	TestDataHash                             string
	TestDataHashWhenTestCaseWasSavedOrLoaded string

	ImportedTemplateFilesFromGitHub []importFilesFromGitHub.GitHubFile
	TestData                        *testDataEngine.TestDataForGroupObjectStruct

	TestCasePreviewObject *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage
}

type AttributeStructSliceReferenceType []*AttributeStruct

type AttributeStruct struct {
	AttributeUuid                             string
	AttributeName                             string
	AttributeValue                            string
	AttributeChangedValue                     string
	AttributeTypeName                         string
	AttributeType                             fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionAttributeTypeEnum
	AttributeTextBoxProperty                  *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty
	AttributeComboBoxProperty                 *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty
	AttributeResponseVariableComboBoxProperty *AttributeResponseVariableComboBoxPropertyStruct
	EntryRef                                  *widget.Entry
	SelectRef                                 *widget.Select
	AttributeIsChanged                        bool
	TestInstructionElementMatureUuidUuid      string
	AttributeValueIsValidRegExAsString        string
	CompileRegEx                              *regexp.Regexp
	AttributeValueIsValid                     bool
	AttributeValueIsValidWarningBox           *canvas.Rectangle
}

type AttributeResponseVariableComboBoxPropertyStruct struct {
	AttributeResponseVariableComboBoxProperty              *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeResponseVariableComboBoxProperty
	MatureTestInstructionsWithCorrectResponseVariablesType *[]*MatureTestInstructionWithCorrectResponseVariablesTypeStruct
}

// Holding one Mature TestInstruction, Uuid and Name, that match a response variable type
type MatureTestInstructionWithCorrectResponseVariablesTypeStruct struct {
	MatureTestInstructionUuidWithCorrectResponseVariablesType string
	MatureTestInstructionNameWithCorrectResponseVariablesType string
	MatureTestInstructionComboBoxOptionsName                  string
}

type MatureTestInstructionStruct struct {
	/*
		MatureTestInstructions          *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage
			[]*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage_MatureTestInstructionMessage
				BasicTestInstructionInformation  *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage
					xy NonEditableInformation    *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_NonEditableBasicInformationMessage
					xy EditableInformation       *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_EditableBasicInformationMessage
					xy InvisibleBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_InvisibleBasicInformationMessage
				MatureTestInstructionInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage
					xy MatureBasicTestInstructionInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_MatureBasicTestInstructionInformationMessage
					xy CreatedAndUpdatedInformation          *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_CreatedAndUpdatedInformationMessage
					TestInstructionAttributesList         []*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
						xy BaseAttributeInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_BaseAttributeInformationMessage
						AttributeInformation     *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage
							xy InputTextBoxProperty      *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputTextBoxProperty
							xy InputComboBoxProperty     *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputComboBoxProperty
							xy InputFileSelectorProperty *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage_AttributeInformationMessage_TestInstructionAttributeInputFileSelectorProperty
	*/

	// BasicTestInstructionInformation
	BasicTestInstructionInformation_NonEditableInformation    *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_NonEditableBasicInformationMessage
	BasicTestInstructionInformation_EditableInformation       *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_EditableBasicInformationMessage
	BasicTestInstructionInformation_InvisibleBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionInformationMessage_InvisibleBasicInformationMessage

	// Specific Mature information
	MatureBasicTestInstructionInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_MatureBasicTestInstructionInformationMessage `protobuf:"bytes,1,opt,name=MatureBasicTestInstructionInformation,proto3" json:"MatureBasicTestInstructionInformation,omitempty"` // The Basic information for the Matures TestInstruction
	CreatedAndUpdatedInformation          *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_CreatedAndUpdatedInformationMessage          `protobuf:"bytes,2,opt,name=CreatedAndUpdatedInformation,proto3" json:"CreatedAndUpdatedInformation,omitempty"`                   // Information regarding who did what and when
	TestInstructionAttributesList         map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionInformationMessage_TestInstructionAttributeMessage   `protobuf:"bytes,3,rep,name=TestInstructionAttributesList,proto3" json:"TestInstructionAttributesList,omitempty"`                 // All attributes that belongs to the TestInstruction

	/*
		FullTestCaseMessage *fenixGuiTestCaseBuilderServerGrpcApi.FullTestCaseMessage
			TestCaseBasicInformation        *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage
			MatureTestInstructions          *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionsMessage
			MatureTestInstructionContainers *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainersMessage
	*/
}

type MatureTestInstructionContainerStruct struct {
	NonEditableInformation                     *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_NonEditableBasicInformationMessage
	EditableInformation                        *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_EditableBasicInformationMessage
	InvisibleBasicInformation                  *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_InvisibleBasicInformationMessage
	EditableTestInstructionContainerAttributes *fenixGuiTestCaseBuilderServerGrpcApi.BasicTestInstructionContainerInformationMessage_EditableTestInstructionContainerAttributesMessage

	MatureTestInstructionContainerInformation *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerInformationMessage_MatureTestInstructionContainerInformationMessage
	CreatedAndUpdatedInformation              *fenixGuiTestCaseBuilderServerGrpcApi.MatureTestInstructionContainerInformationMessage_CreatedAndUpdatedInformationMessage
}

type CurrentSelectedTestCaseElementStruct struct {
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
	BasicTestCaseInformationMessageNoneEditableInformation fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage_NonEditableBasicInformationMessage  // All Basic information that can be shown in GUI but can't be changed by the user
	BasicTestCaseInformationMessageEditableInformation     fenixGuiTestCaseBuilderServerGrpcApi.BasicTestCaseInformationMessage_EditableBasicInformationMessage     // All Basic information that can be shown in GUI and can be changed by the user
	CreatedAndUpdatedInformation                           fenixGuiTestCaseBuilderServerGrpcApi.TestCaseBasicInformationMessage_CreatedAndUpdatedInformationMessage // Information regarding who did what and when
	DeleteTimeStamp                                        string                                                                                                   // YYYY-MM-DD

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
