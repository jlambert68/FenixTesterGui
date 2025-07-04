package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/soundEngine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
)

// ExecuteOneTestSuiteWithOneTestDataSet
// Execute one TestSuit's all TestCases with one TestDataSets
func (testSuiteModel *TestSuiteModelStruct) ExecuteOneTestSuiteWithOneTestDataSet() {

	// Define variable for TestData
	var testDataForTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestDataForTestCaseExecutionMessage

	// Are there any TestData
	if testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr.ChosenTestDataPointsPerGroupMap != nil {
		// There might be TestData

		testDataPtr := testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr

		testDataDomainAndAreaNameToUuidMap := *testDataEngine.TestDataModel.TestDataDomainAndAreaNameToUuidMap

		// Check if there are TestData to pick, but no TestData was picked. Should the Execution execute without TestData
		var executeWithOutTestData bool

		// Check if there exists TestData but the user didn't pick any
		if len(testDataDomainAndAreaNameToUuidMap) > 0 &&
			(len(testDataPtr.SelectedTestDataGroup) == 0 ||
				len(testDataPtr.SelectedTestDataPoint) == 0 ||
				len(testDataPtr.SelectedTestDataPointRowSummary) == 0) {

			var syncChannel chan bool
			syncChannel = make(chan bool, 1)

			//mainWindow := *sharedCode.FenixMasterWindowPtr

			dialog.ShowConfirm("Confirmation", "There are TestData to chose from, "+
				"but no TestData was selected for the TestSuite."+
				"\n\nWould you like to execute the TestSuite WITHOUT TestData?",
				func(response bool) {

					executeWithOutTestData = response

					syncChannel <- true

				},
				*sharedCode.FenixMasterWindowPtr)

			/*
				fyne.Do(func() {
					dialog.NewConfirm("Confirmation", "There are TestData to chose from, "+
						"but no TestData was selected for the TestSuite."+
						"\n\nWould you like to execute the TestSuite WITHOUT TestData?", func(response bool) {

						executeWithOutTestData = response

						syncChannel <- true

					}, mainWindow).Show()
				})

			*/

			// Wait for user action
			<-syncChannel

			// User wants to add testdata
			if executeWithOutTestData == false {

				return
			}
		}

		if executeWithOutTestData == true {

			// TestData exist but not chosen
			testDataForTestCaseExecution = &fenixExecutionServerGuiGrpcApi.TestDataForTestCaseExecutionMessage{}

		} else {

			// Extract values from TestData-model
			testDataPointRowUuid := testDataPtr.TestDataColumnDataNameToValueMap["TestDataPointRowUuid"]
			testDataModelMap := *testDataEngine.TestDataModel.TestDataModelMap
			testDataAreasMap := *testDataModelMap[testDataEngine.TestDataDomainUuidType(
				testDataPtr.SelectedTestDataDomainUuid)].TestDataAreasMap
			testDataFileSha256Hash := testDataAreasMap[testDataEngine.TestDataAreaUuidType(
				testDataPtr.SelectedTestDataTestDataAreaUuid)].TestDataFileSha256Hash

			// Create the TestDataValueMap
			var testDataValueMap map[string]*fenixExecutionServerGuiGrpcApi.TestDataValueMapValueMessage
			testDataValueMap = make(map[string]*fenixExecutionServerGuiGrpcApi.TestDataValueMapValueMessage)
			for headerDataName, dataValue := range testDataPtr.TestDataColumnDataNameToValueMap {
				var testDataValueMapValueMessage fenixExecutionServerGuiGrpcApi.TestDataValueMapValueMessage
				testDataValueMapValueMessage = fenixExecutionServerGuiGrpcApi.TestDataValueMapValueMessage{
					HeaderDataName:                    headerDataName,
					TestDataValue:                     dataValue,
					TestDataValueIsReplaced:           false, // TODO implement this
					TestDataOriginalValueWhenReplaced: "",    // TODO implement this
				}

				// Add too TestDataValueMap
				testDataValueMap[headerDataName] = &testDataValueMapValueMessage

			}

			// Get selected TestData for execution

			testDataForTestCaseExecution = &fenixExecutionServerGuiGrpcApi.TestDataForTestCaseExecutionMessage{
				TestDataDomainUuid:         testDataPtr.SelectedTestDataDomainUuid,
				TestDataDomainName:         testDataPtr.SelectedTestDataDomainName,
				TestDataDomainTemplateName: testDataPtr.SelectedTestDataDomainTemplateName,
				TestDataAreaUuid:           testDataPtr.SelectedTestDataTestDataAreaUuid,
				TestDataAreaName:           testDataPtr.SelectedTestDataAreaName,
				TestDataValueMap:           testDataValueMap,
				TestDataRowIdentifier:      testDataPointRowUuid,
				TestDataFileSha256Hash:     string(testDataFileSha256Hash),
			}
		}
	} else {

		// No TestData exists
		testDataForTestCaseExecution = &fenixExecutionServerGuiGrpcApi.TestDataForTestCaseExecutionMessage{}
	}

	// Create message to be sent to GuiExecutionServer
	var initiateSingleTestSuiteExecutionRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionRequestMessage
	initiateSingleTestSuiteExecutionRequestMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionRequestMessage{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestSuiteUuid:                testSuiteModel.GetTestSuiteUuid(),
		ExecutionStatusReportLevel:   fenixExecutionServerGuiGrpcApi.ExecutionStatusReportLevelEnum_REPORT_ALL_STATUS_CHANGES_ON_EXECUTIONS, //fenixExecutionServerGuiGrpcApi.ExecutionStatusReportLevelEnum_REPORT_ALL_STATUS_CHANGES_ON_EXECUTIONS,
		TestDataForTestCaseExecution: testDataForTestCaseExecution,
	}

	// Initiate TestCaseExecution
	var initiateSingleTestSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestSuiteExecutionResponseMessage
	initiateSingleTestSuiteExecutionResponseMessage = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendInitiateTestSuiteExecution(initiateSingleTestSuiteExecutionRequestMessage)

	if initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.AckNack == false {

		errorId := "28d1c4d7-a8a7-42b5-843e-53fdd79b4d18"
		err := errors.New(fmt.Sprintf("couldn't execute TestSuite due to error: '%s', {error: %s} [ErrorID: %s]",
			initiateSingleTestSuiteExecutionResponseMessage.AckNackResponse.Comments, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return

	}

	// Notify User that the TestCase is execution
	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "TestSuiteExecution",
		Content: "The TestSuite is sent for Execution. See execution Tab for status.",
	})

}
