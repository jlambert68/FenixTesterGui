package testSuitesModel

import (
	"FenixTesterGui/testCase/testCaseModel"
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
)

// Pointer to the object holding all TestSuites (visible in a Tab)
var TestSuitesModelPtr *TestSuitesModelStruct

// Holding all information about TestSuites
type TestSuitesModelStruct struct {
	TestSuitesMapPtr            *map[string]*TestSuiteModelStruct // Map-key = 'TestSuiteUuid'
	TestSuiteMetaDataForDomains TestSuiteMetaDataForDomainsStruct
}

// TestSuiteModelStruct
// Holdning all information about a specific TestSuite
type TestSuiteModelStruct struct {

	// Object holding the data that can't be changed directly via the UI
	testSuiteModelDataThatCanNotBeChangedFromUI testSuiteModelDataThatCaNotBeChangedFromUIStruct

	// Object holding info about if OwnerDomain and TestEnvironment has got any value, used for Locking parts of the GUI in the TestSuite
	lockValuesForOwnerDomainAndTestEnvironment lockValuesForOwnerDomainAndTestEnvironmentStruct

	// Structure that keeps data used after saving was successfully performed
	savedTestSuiteUIModelBinding TestSuiteUIModelBindingStruct

	// Structure that keeps data used for saving before a successful save was performed
	NoneSavedTestSuiteUIModelBinding TestSuiteUIModelBindingStruct

	// TestSuite UI components bindings
	TestSuiteUIModelBinding TestSuiteUIModelBindingStruct

	// Reference to the TestCase-model
	testCasesModel *testCaseModel.TestCasesModelsStruct
}

// Storing values regarding if Owner domain and TestEnvironment has got any value selected by the user
type lockValuesForOwnerDomainAndTestEnvironmentStruct struct {
	OwnerDomainHasValue     bool
	TestEnvironmentHasValue bool
	LockButtonHaBeenClicked bool
}

// Object holding the data that can't be changed directly via the UI
type testSuiteModelDataThatCaNotBeChangedFromUIStruct struct {
	testSuiteUuid              string // The TestSuites Uuid
	testSuiteVersion           uint32 // The current TestSuite version
	createdByGcpLogin          string // The person that did log in towards GCP
	createdByComputerLogin     string // The person that is logged into the computer
	createdDate                string // The date when the TestSuite was first created
	lastChangedByGcpLogin      string // The person that did log in towards GCP when TestSuite was last changed and saved
	lastChangedByComputerLogin string // The person that is logged into the computer when TestSuite was last changed and saved
	lastChangedDate            string // The date when the TestSuite was last changed and saved
	testSuiteSavedMessageHash  string // Hash of save TestSuiteMessage
}

// TestSuiteUIModelBindingStruct
// Holding bindings to textboxes, dropDown, in UI
type TestSuiteUIModelBindingStruct struct {
	TestSuiteDeletionDate         string // Date that TestSuite will be deleted
	TestSuiteName                 string // The Name of the TestSuite
	TestSuiteDescription          string // A description for the TestSuite
	TestSuiteOwnerDomainUuid      string // The Uuid for the Domain that owns the TestSuite
	TestSuiteOwnerDomainName      string // The Name for the Domain that owns the TestSuite
	TestSuiteExecutionEnvironment string // The execution environment where the TestSuite will be executed

	TestSuiteIsNew        bool
	TestSuiteTestDataHash string                                       // Hash used to check if changes has been done to TestData
	TestDataPtr           *testDataEngine.TestDataForGroupObjectStruct // Holding all selected TestData

	TestSuiteMetaDataHash string                   // Hash used to check if changes has been done to MetaData
	TestSuiteMetaDataPtr  *TestSuiteMetaDataStruct // Holding the current TestSuiteMetaDataSet and what has been selected

	TestSuiteTypeHash string              // The hash of the TestSuiteType
	TestSuiteType     TestSuiteTypeStruct // Defines the type for the TestSuite, i.e. Standard, Pig...

}

type TestSuiteTypeStruct struct {
	TestSuiteType     TestSuiteTypeType
	TestSuiteTypeName TestSuiteTypeNameType
}

const (
	TestSuiteTypeIsNotDefined TestSuiteTypeType = iota
	TestSuiteTypeIsStandard
	TestSuiteTypeIsPig
)

const (
	TestSuiteTypeNameNotDefined TestSuiteTypeNameType = "Not Defined"
	TestSuiteTypeNameStandard   TestSuiteTypeNameType = "Standard"
	TestSuiteTypeNamePig        TestSuiteTypeNameType = "Pig"
)

// TestSuiteTypeNameMap
// The Map the gives the name for the TestSuite-type
var TestSuiteTypeNameMap = map[TestSuiteTypeType]TestSuiteTypeNameType{
	TestSuiteTypeIsNotDefined: TestSuiteTypeNameNotDefined,
	TestSuiteTypeIsStandard:   TestSuiteTypeNameStandard,
	TestSuiteTypeIsPig:        TestSuiteTypeNamePig,
}

// TestSuiteTypeNameType
// The type defined for the Type-name of the TestSuite
type TestSuiteTypeNameType string

// TestSuiteTypeType
// The type defining what 'type' the TestSuite is of
type TestSuiteTypeType uint8

// TestSuiteMetaDataForDomainsStruct
// Holding all MetaData for all domains
type TestSuiteMetaDataForDomainsStruct struct {
	TestSuiteMetaDataForDomainsMap map[string]*TestSuiteMetaDataForDomainsForMapStruct // Key = DomainUuid
	UniqueMetaDataBitSets          UniqueMetaDataBitSetsStruct
}

// UniqueMetaDataBitSetsStruct
// Holding the unique Bitset for each of the  Domains, Groups, GroupItems, ItemValues
type UniqueMetaDataBitSetsStruct struct {
	DomainsBitSetMap                 map[string]*boolbits.BitSet // map key = DomainUuid
	MetaDataGroupsBitSetMap          map[string]*boolbits.BitSet // map key = GroupName
	MetaDataGroupItemsBitSetMap      map[string]*boolbits.BitSet // map key = GroupItemName
	MetaDataGroupItemValuesBitSetMap map[string]*boolbits.BitSet // map key = GroupItemValue
}

// TestSuiteMetaDataForDomainsMapStruct
// Hold one Domains all TestSuiteMetaData as original json and as a struct
type TestSuiteMetaDataForDomainsForMapStruct struct {
	TestSuiteMetaDataForDomainAsJsonPtr *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage // Hold one Domains all TestSuiteMetaData as original json
	TestSuiteMetaDataForDomainPtr       *TestSuiteMetaDataForDomainStruct                                                     // Hold one Domains all TestSuiteMetaData as a struct
}

// TestSuiteMetaDataForDomainStruct
// Struct holding the TestSuiteMetaData converted from the pure json-object
type TestSuiteMetaDataForDomainStruct struct {
	MetaDataGroups []struct {
		MetaDataGroupName string `json:"MetaDataGroupName"`
		MetaDataInGroup   []struct {
			MetaDataName   string   `json:"MetaDataName"`
			SelectType     string   `json:"SelectType"`
			Mandatory      string   `json:"Mandatory"`
			MetaDataValues []string `json:"MetaDataValues"`
		} `json:"MetaDataInGroup"`
	} `json:"MetaDataGroups"`
}

// TestSuiteMetaDataStruct
// Struct holding the current TestSuiteMetaDataSet and what has been selected
type TestSuiteMetaDataStruct struct {
	CurrentSelectedDomainUuid                               string                                                                                // Specifies the current selected Owner Domain for the TestSuite
	TestSuiteMetaDataMessageJsonForTestSuiteWhenLastSaved   *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseAndTestSuiteMetaDataForOneDomainMessage // The json used with the latest save version of the TestSuite
	TestSuiteMetaDataMessageStructForTestSuiteWhenLastSaved *TestSuiteMetaDataForDomainStruct
	MetaDataGroupsOrder                                     []string
	MetaDataGroupsMapPtr                                    *map[string]*MetaDataGroupStruct // holding MetaDataGroups and its MetaData. The key is the MetaDataGroupName
	SelectedTestSuiteMetaDataAsEntrySlice                   []*boolbits.Entry                // A slice holding all selected MetaData as boolbits-Entry types
}

// MetaDataGroupStruct
// Struct holding one MetaDataGroup and its MetaData
type MetaDataGroupStruct struct {
	MetaDataGroupName string
	//MetaDataInGroupPtr *[]*MetaDataInGroupStruct // Holding all MetaDataName and their values. It also holds what was selected
	MetaDataInGroupOrder  []string
	MetaDataInGroupMapPtr *map[string]*MetaDataInGroupStruct // Holding all MetaDataName and their values. It also holds what was selected
}

// MetaDataInGroupStruct
// Struct holding the available values, how they are selected and what was selected
type MetaDataInGroupStruct struct {
	MetaDataGroupName                          string             // The name of the MetaData-Group
	MetaDataName                               string             // The name of the MetaData-post
	SelectType                                 MetaDataSelectType // Is the MetaData-post single- or multi-select
	Mandatory                                  bool               // Is the MetaData-post mandatory or not
	AvailableMetaDataValues                    []string           // The available values for the MetaData-post
	SelectedMetaDataValueForSingleSelect       string             // The value selected for single select
	SelectedMetaDataValuesForMultiSelect       []string           // The values selected for multi select
	SelectedMetaDataValuesForMultiSelectMapPtr *map[string]string // The values selected for multi select
}

// MetaDataSelectType
// The type for the SelectType for the MetaData
type MetaDataSelectType uint8 // The type used for SelectType

const (
	MetaDataSelectType_NotSelected MetaDataSelectType = iota
	MetaDataSelectType_SingleSelect
	MetaDataSelectType_MultiSelect
)

// The type for the constants used to specify what has been stored in database
type testSuiteImplementedFucntionsType uint8

const (
	testSuiteBasicInformationIsSupported testSuiteImplementedFucntionsType = iota
	testSuiteTestDataIsSupported
	testSuitePreviewIsSupported
	testSuiteMetaDataIsSupported
	testCasesInTestSuiteIsSupported
	deletedDateIsSupported
	testSuiteTypeIsSupported
	testSuiteImplementedFunctionsMapIsSupported
)

type testSuiteImplementedFunctionsToBeStoredStruct struct {
	testSuiteImplementedFunctionsMap map[testSuiteImplementedFucntionsType]bool `json:"testSuiteImplementedFunctionsMap"`
}
