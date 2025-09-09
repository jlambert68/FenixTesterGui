package listTestCasesModel

import (
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"time"
)

// Holden the complete filter for a TestCasePreView
var SimpleTestCaseMetaDataFilterEntryMap map[string]*boolbits.Entry // Key = TestCaseUuid

var TestCasesThatCanBeEditedByUserMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage

// Hold dates for last retrieved TestCase and TestCaseExecution.
var LatestTestCaseUpdatedMinTimeStampForDatabaseCall time.Time
var LatestTestCaseExecutionUpdatedMinTimeStampForDatabaseCall time.Time
