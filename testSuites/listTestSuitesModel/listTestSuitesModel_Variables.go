package listTestSuitesModel

import (
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// Holden the complete filter for a TestSuitePreView
var SimpleTestSuiteMetaDataFilterEntryMap map[string]*boolbits.Entry // Key = TestSuiteUuid

var TestSuitesThatCanBeEditedByUserMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.BasicTestSuiteInformationMessage
