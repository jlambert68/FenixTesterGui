package commandAndRuleEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"github.com/sirupsen/logrus"
	"image/color"
	"log"
	"strconv"
	"sync"
)

// Channel reader which is used for reading out commands to CommandEngine
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) startCommandChannelReader() {

	var incomingChannelCommand sharedCode.ChannelCommandStruct

	for {
		// Wait for incoming command over channel
		incomingChannelCommand = <-*commandAndRuleEngine.CommandChannelReference

		switch incomingChannelCommand.ChannelCommand {

		case sharedCode.ChannelCommandNewTestCase:
			commandAndRuleEngine.channelCommandCreateNewTestCase(incomingChannelCommand)

		case sharedCode.ChannelCommandSwapElement:
			commandAndRuleEngine.channelCommandSwap(incomingChannelCommand)

		case sharedCode.ChannelCommandRemoveElement:
			commandAndRuleEngine.channelCommandRemove(incomingChannelCommand)

		case sharedCode.ChannelCommandSaveTestCase:
			commandAndRuleEngine.channelCommandSaveTestCase(incomingChannelCommand)

		case sharedCode.ChannelCommandExecuteTestCase:
			commandAndRuleEngine.channelCommandExecuteTestCase(incomingChannelCommand)

		case sharedCode.ChannelCommandChangeActiveTestCase:
			commandAndRuleEngine.channelCommandChangeActiveTestCase(incomingChannelCommand)

		case sharedCode.ChannelCommandOpenTestCase:
			commandAndRuleEngine.channelCommandOpenTestCase(incomingChannelCommand)

		case sharedCode.ChannelCommandRemoveTestCaseWithOutSaving:
			commandAndRuleEngine.channelCommandRemoveTestCaseWithOutSaving(incomingChannelCommand)

		case sharedCode.ChannelCommandCloseOpenTestCaseWithOutSaving:
			commandAndRuleEngine.channelCommandCloseOpenTestCaseWithOutSaving(incomingChannelCommand)

		// No other command is supported
		default:
			//TODO Send Error over ERROR-channel
		}
	}

}

// Execute command 'New TestCase' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandCreateNewTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	// Create New TestCase
	testCaseUuid, err := commandAndRuleEngine.NewTestCaseModel()

	if err != nil {
		fmt.Println(err)

		return

	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedNewTestCase,
		CreateNewTestCaseUI:          true,
		ActiveTestCase:               testCaseUuid,
		TextualTestCaseSimple:        textualTestCaseSimple,
		TextualTestCaseComplex:       textualTestCaseComplex,
		TextualTestCaseExtended:      textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Execute command 'Save TestCase' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandSaveTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	currentTestCaseUuid := commandAndRuleEngine.Testcases.CurrentActiveTestCaseUuid

	// Save TestCase
	err := commandAndRuleEngine.Testcases.SaveFullTestCase(currentTestCaseUuid, commandAndRuleEngine.Testcases.CurrentUser)

	if err != nil {

		errorId := "b91f4270-babc-4432-9a9b-4769f1d662f9"
		err = errors.New(fmt.Sprintf("couldn't execute command 'SaveFullTestCase', {error: %s} [ErrorID: %s]", err, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return

	}

	// Extract the current TestCase UI model
	testCase_Model, existsInMap := commandAndRuleEngine.Testcases.TestCasesMap[currentTestCaseUuid]
	if existsInMap == false {
		errorId := "a08730c6-aa91-4b03-9144-eeeccc153d96"
		err := errors.New(fmt.Sprintf("testcase-model with TestCaseUuid '%s' is missing map for TestCasesMap [ErrorID: %s]", currentTestCaseUuid, errorId))

		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return

	}

	// Generate short version of UUID to put in TestCase Tab-Name
	var shortUUid string
	var tabName string
	var testCaseName string
	testCaseName = testCase_Model.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName

	shortUUid = commandAndRuleEngine.Testcases.GenerateShortUuidFromFullUuid(currentTestCaseUuid)

	// Shorten Tab-name if name is longer then 'testCaseTabNameVisibleLength'
	if len(testCaseName) > sharedCode.TestCaseTabNameVisibleLength {
		tabName = testCaseName[0:sharedCode.TestCaseTabNameVisibleLength]
	} else {
		tabName = testCaseName
	}

	tabName = tabName + " [" + shortUUid + "]"

	// Notify the user that the TestCase was Saved

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "TestCase Save",
		Content: tabName + " was successfully save in database",
	})

	// Send 'update TestCase name for Tab' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedUpdateTestCaseTabName,
		CreateNewTestCaseUI:          false,
		ActiveTestCase:               currentTestCaseUuid,
		TextualTestCaseSimple:        "",
		TextualTestCaseComplex:       "",
		TextualTestCaseExtended:      "",
		TestCaseTabName:              tabName,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Execute command 'Save TestCase' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandExecuteTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	// TestCaseUuid to execute
	var testCaseUuidToBeExecuted string
	testCaseUuidToBeExecuted = incomingChannelCommand.FirstParameter

	var existInMap bool
	var currentTestCasePtr *testCaseModel.TestCaseModelStruct

	currentTestCasePtr, existInMap = commandAndRuleEngine.Testcases.TestCasesMap[testCaseUuidToBeExecuted]
	if existInMap == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "8d25e3da-e1b5-4222-9446-d25dd7d930f1",
			"testCaseUuid": testCaseUuidToBeExecuted,
		}).Fatal("TestCase doesn't exist in TestCaseMap. This should not happen")
	}

	// Define variable for TestData
	var testDataForTestCaseExecution *fenixExecutionServerGuiGrpcApi.TestDataForTestCaseExecutionMessage

	// Are there any TestData
	if currentTestCasePtr.TestData != nil {
		// There might be TestData

		testDataDomainAndAreaNameToUuidMap := *testDataEngine.TestDataModel.TestDataDomainAndAreaNameToUuidMap

		// Check if there are TestData to pick, but no TestData was picked. Should the Execution execute without TestData
		var executeWithOutTestData bool

		// Check if there exists TestData but the user didn't pick any
		if len(testDataDomainAndAreaNameToUuidMap) > 0 &&
			(len(currentTestCasePtr.TestData.SelectedTestDataGroup) == 0 ||
				len(currentTestCasePtr.TestData.SelectedTestDataPoint) == 0 ||
				len(currentTestCasePtr.TestData.SelectedTestDataPointRowSummary) == 0) {

			var syncChannel chan bool
			syncChannel = make(chan bool, 1)

			dialog.NewConfirm("Confirmation", "There are TestData to chose from, "+
				"but no TestData was selected for the TestCase."+
				"\n\nWould you like to execute the TestCase WITHOUT TestData?", func(response bool) {

				executeWithOutTestData = response

				syncChannel <- true

			}, *sharedCode.FenixMasterWindowPtr).Show()

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
			testDataPointRowUuid := currentTestCasePtr.TestData.TestDataColumnDataNameToValueMap["TestDataPointRowUuid"]
			testDataModelMap := *testDataEngine.TestDataModel.TestDataModelMap
			testDataAreasMap := *testDataModelMap[testDataEngine.TestDataDomainUuidType(
				currentTestCasePtr.TestData.SelectedTestDataDomainUuid)].TestDataAreasMap
			testDataFileSha256Hash := testDataAreasMap[testDataEngine.TestDataAreaUuidType(
				currentTestCasePtr.TestData.SelectedTestDataTestDataAreaUuid)].TestDataFileSha256Hash

			// Create the TestDataValueMap
			var testDataValueMap map[string]*fenixExecutionServerGuiGrpcApi.TestDataValueMapValueMessage
			testDataValueMap = make(map[string]*fenixExecutionServerGuiGrpcApi.TestDataValueMapValueMessage)
			for headerDataName, dataValue := range currentTestCasePtr.TestData.TestDataColumnDataNameToValueMap {
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
				TestDataDomainUuid:         currentTestCasePtr.TestData.SelectedTestDataDomainUuid,
				TestDataDomainName:         currentTestCasePtr.TestData.SelectedTestDataDomainName,
				TestDataDomainTemplateName: currentTestCasePtr.TestData.SelectedTestDataDomainTemplateName,
				TestDataAreaUuid:           currentTestCasePtr.TestData.SelectedTestDataTestDataAreaUuid,
				TestDataAreaName:           currentTestCasePtr.TestData.SelectedTestDataAreaName,
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
	var initiateSingleTestCaseExecutionRequestMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionRequestMessage
	initiateSingleTestCaseExecutionRequestMessage = &fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionRequestMessage{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestCaseUuid:                 testCaseUuidToBeExecuted,
		TestDataSetUuid:              testCaseUuidToBeExecuted,                                                                              //TODO change into a correct 'TestDataSetUuid' when that is supported
		ExecutionStatusReportLevel:   fenixExecutionServerGuiGrpcApi.ExecutionStatusReportLevelEnum_REPORT_ALL_STATUS_CHANGES_ON_EXECUTIONS, //fenixExecutionServerGuiGrpcApi.ExecutionStatusReportLevelEnum_REPORT_ALL_STATUS_CHANGES_ON_EXECUTIONS,
		TestDataForTestCaseExecution: testDataForTestCaseExecution,
	}

	// Initiate TestCaseExecution
	var initiateSingleTestCaseExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.InitiateSingleTestCaseExecutionResponseMessage
	initiateSingleTestCaseExecutionResponseMessage = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendInitiateTestCaseExecution(initiateSingleTestCaseExecutionRequestMessage)

	if initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.AckNack == false {

		errorId := "524d549f-58b9-4b29-8668-0a322137b3cf"
		err := errors.New(fmt.Sprintf("couldn't execute command 'ExecuteTestCase' due to error: '%s', {error: %s} [ErrorID: %s]", initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.Comments, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return

	}

	// Notify User that the TestCase is execution
	// Notify the user

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "TestCase Execution",
		Content: "The TestCase is sent for Execution. See execution Tab for status.",
	})

	// Add TestCaseExecution to Executions-table for TestCaseExecutionOnQueue by send message to channel used to update OnQueue-table
	// Create Remove-message to be put on channel
	var onQueueTableAddRemoveChannelMessage executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct
	onQueueTableAddRemoveChannelMessage = executionsModelForSubscriptions.OnQueueTableAddRemoveChannelStruct{
		ChannelCommand: executionsModelForSubscriptions.OnQueueTableAddRemoveChannelAddCommand_AddAndFlash,
		AddCommandData: executionsModelForSubscriptions.OnQueueAddCommandDataStruct{
			TestCaseExecutionBasicInformation: initiateSingleTestCaseExecutionResponseMessage.TestCasesInExecutionQueue,
		},
	}

	// Put message on channel
	executionsModelForSubscriptions.OnQueueTableAddRemoveChannel <- onQueueTableAddRemoveChannelMessage

	// TestExecutionMapKey
	var tempTestExecutionMapKey string
	tempTestExecutionMapKey = initiateSingleTestCaseExecutionResponseMessage.TestCasesInExecutionQueue.TestCaseExecutionUuid +
		strconv.Itoa(int(initiateSingleTestCaseExecutionResponseMessage.TestCasesInExecutionQueue.TestCaseExecutionVersion))

	// Send message Executions Details handler to retrieve full TestCaseExecutions details
	_ = detailedExecutionsModel.RetrieveSingleTestCaseExecution(tempTestExecutionMapKey)

	fmt.Sprintf("Initiated TestCaseExecution for TestCase: '%s', testCaseUuidToBeExecuted")

}

// Execute command 'Remove Element' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandRemove(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	currentTestCaseUuid := incomingChannelCommand.ActiveTestCase
	elementToRemove := incomingChannelCommand.FirstParameter

	// Delete Element from TestCase
	err := commandAndRuleEngine.DeleteElementFromTestCaseModel(currentTestCaseUuid, elementToRemove)
	if err != nil {
		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return
	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(currentTestCaseUuid)
	if err != nil {
		fmt.Println(err)
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedUpdateTestCaseGraphics,
		CreateNewTestCaseUI:          false,
		ActiveTestCase:               currentTestCaseUuid,
		TextualTestCaseSimple:        textualTestCaseSimple,
		TextualTestCaseComplex:       textualTestCaseComplex,
		TextualTestCaseExtended:      textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Change the active TestCase and TestCase-tab
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandChangeActiveTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	var existInMap bool
	var tempTestCasePtr *testCaseModel.TestCaseModelStruct

	currentTestCaseUuid := incomingChannelCommand.ActiveTestCase

	// Verify that TestCase exists
	tempTestCasePtr, existInMap = commandAndRuleEngine.Testcases.TestCasesMap[currentTestCaseUuid]
	if existInMap == false {
		errorId := "a7645bee-7739-4ea3-a0f5-60d5339fb2e4"
		err := errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMap [ErrorID: %s]", currentTestCaseUuid, errorId))
		// TODO Send on Error-channel
		fmt.Println(err)

		return
	}

	// Set Active TestCase
	commandAndRuleEngine.Testcases.CurrentActiveTestCaseUuid = currentTestCaseUuid

	// If this is an already saved TestCase then check if there are changes in Database
	var testCaseHashIsTheSame bool
	var err2 error
	if tempTestCasePtr.ThisIsANewTestCase == false {
		testCaseHashIsTheSame, err2 = commandAndRuleEngine.Testcases.VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase(currentTestCaseUuid)
		fmt.Println("Is TestCase-Hash the same as Database-hash", testCaseHashIsTheSame, err2, currentTestCaseUuid)
	}

	// Check if current TestCase-hash has changed since TestCase was Saved or Loaded
	TestCaseHashHasChangedSincesSavedOrLoaded, err3 := commandAndRuleEngine.Testcases.TestCaseHashIsChangedSinceLoadedOrSaved(currentTestCaseUuid)
	fmt.Println("IS TestCase-Hash the changed since TestCase was Saved or Loaded", TestCaseHashHasChangedSincesSavedOrLoaded, err3, currentTestCaseUuid)

	// Generate short version of UUID to put in TestCase Tab-Name
	var shortUUid string
	var tabName string
	var testCaseName string
	testCaseName = tempTestCasePtr.LocalTestCaseMessage.BasicTestCaseInformationMessageEditableInformation.TestCaseName

	shortUUid = commandAndRuleEngine.Testcases.GenerateShortUuidFromFullUuid(currentTestCaseUuid)

	// Shorten Tab-name if name is longer then 'testCaseTabNameVisibleLength'
	if len(testCaseName) > sharedCode.TestCaseTabNameVisibleLength {
		tabName = testCaseName[0:sharedCode.TestCaseTabNameVisibleLength]
	} else {
		tabName = testCaseName
	}

	tabName = tabName + " [" + shortUUid + "]"

	if tempTestCasePtr.ThisIsANewTestCase == true {
		// New TestCase
		tabName = tabName + " (*)"
	} else {
		if testCaseHashIsTheSame == false {
			// TestCase changed in Database
			tabName = tabName + " (X)"
		} else {
			if TestCaseHashHasChangedSincesSavedOrLoaded == true {
				// TestCase updated in UI
				tabName = tabName + " (*)"
			}
		}
	}

	// Send 'update TestCase name for Tab' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedUpdateTestCaseTabName,
		CreateNewTestCaseUI:          false,
		ActiveTestCase:               currentTestCaseUuid,
		TextualTestCaseSimple:        "",
		TextualTestCaseComplex:       "",
		TextualTestCaseExtended:      "",
		TestCaseTabName:              tabName,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// PopUp used function 'channelCommandOpenTestCase', below. Generates the ability for user enter Uuid
func runPopUp(w fyne.Window, uuidChannel chan<- string) (modal *widget.PopUp) {

	var hBoxButtonContainer *fyne.Container
	var vBoxContainer *fyne.Container
	var uuidEntryBox *widget.Entry
	var okButton *widget.Button
	var closeButton *widget.Button

	uuidEntryBox = widget.NewEntry()

	okButton = widget.NewButton("Ok", func() {
		modal.Hide()
		uuidChannel <- uuidEntryBox.Text
	})
	closeButton = widget.NewButton("Close", func() {
		modal.Hide()
		uuidChannel <- ""
	})

	hBoxButtonContainer = container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		okButton,
		closeButton)

	vBoxContainer = container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Enter TestCaseUuid to open"),
		uuidEntryBox,
		hBoxButtonContainer)

	modal = widget.NewModalPopUp(
		vBoxContainer,
		w.Canvas(),
	)

	modal.Show()

	// Set the focus on the entry box
	w.Canvas().Focus(uuidEntryBox)

	return modal
}

// Opens a saved TestCase from Database
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandOpenTestCase(
	incomingChannelCommand sharedCode.ChannelCommandStruct) {

	var existInMap bool

	// Create return channel for UUID-value
	var returnChannelFromOpenTestCaseUuidPopUp chan string
	returnChannelFromOpenTestCaseUuidPopUp = make(chan string)

	//var modal *widget.PopUp
	_ = runPopUp(*commandAndRuleEngine.MasterFenixWindow, returnChannelFromOpenTestCaseUuidPopUp)

	// Way for UUID value from Open TestCase PopUp
	var uuidToOpen string
	uuidToOpen = <-returnChannelFromOpenTestCaseUuidPopUp

	// Verify if TestCase exists
	_, existInMap = commandAndRuleEngine.Testcases.TestCasesMap[uuidToOpen]
	if existInMap == true {

		/*
			// Change to Tab containing TestCaseUuid
			// Send command 'ChannelCommandChangeActiveTestCase' on command-channel
			commandEngineChannelMessage := sharedCode.ChannelCommandStruct{
				ChannelCommand:  sharedCode.ChannelCommandChangeActiveTestCase,
				FirstParameter:  "",
				SecondParameter: "",
				ActiveTestCase:  uuidToOpen,
				ElementType:     sharedCode.BuildingBlock(sharedCode.Undefined),
			}

			// Send command message over channel to Command and Rule Engine
			sharedCode.CommandChannel <- commandEngineChannelMessage
		*/
		// Set active TestCase
		commandAndRuleEngine.Testcases.CurrentActiveTestCaseUuid = uuidToOpen

		// Change to Tab containing TestCaseUuid
		// Send 'Switch to TestCase-tab' command over channel
		outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
			ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedSelectTestCaseTabBasedOnTestCaseUuid,
			CreateNewTestCaseUI:          false,
			ActiveTestCase:               uuidToOpen,
			TextualTestCaseSimple:        "",
			TextualTestCaseComplex:       "",
			TextualTestCaseExtended:      "",
			TestInstructionUuid:          "",
		}

		*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

		return
	}

	// Trigger Load the TestCase from Database
	err := commandAndRuleEngine.Testcases.LoadFullTestCaseFromDatabase(uuidToOpen, commandAndRuleEngine.Testcases.CurrentUser)

	if err != nil {

		errorId := "b9b31517-04f3-48e8-a05b-82ebcbba4307"
		err = errors.New(fmt.Sprintf("couldn't execute command 'LoadFullTestCaseFromDatabase', {error: %s} [ErrorID: %s]", err, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return

	}

	fmt.Println("TestCase was Load from Cloud-DB")

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(uuidToOpen)
	if err != nil {
		fmt.Println(err)
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedNewTestCase,
		CreateNewTestCaseUI:          true,
		ActiveTestCase:               uuidToOpen,
		TextualTestCaseSimple:        textualTestCaseSimple,
		TextualTestCaseComplex:       textualTestCaseComplex,
		TextualTestCaseExtended:      textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Remove the TestCase without saving it to theDatabase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandRemoveTestCaseWithOutSaving(
	incomingChannelCommand sharedCode.ChannelCommandStruct) {

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsRemoveTestCaseTabBasedOnTestCaseUuid,
		CreateNewTestCaseUI:          false,
		ActiveTestCase:               incomingChannelCommand.FirstParameter,
		TextualTestCaseSimple:        "",
		TextualTestCaseComplex:       "",
		TextualTestCaseExtended:      "",
		TestInstructionUuid:          "",
		TestCaseTabName:              "",
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Close open TestCase without saving it to theDatabase
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandCloseOpenTestCaseWithOutSaving(
	incomingChannelCommand sharedCode.ChannelCommandStruct) {

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsCloseTestCaseTabBasedOnTestCaseUuiWithOutSaving,
		CreateNewTestCaseUI:          false,
		ActiveTestCase:               incomingChannelCommand.ActiveTestCase,
		TextualTestCaseSimple:        "",
		TextualTestCaseComplex:       "",
		TextualTestCaseExtended:      "",
		TestInstructionUuid:          "",
		TestCaseTabName:              "",
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

func printDropZone(index int) {
	log.Println("Chosen DropZone: is ", index)
}

// Execute command 'Swap Element' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandSwap(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	currentTestCaseUuid := incomingChannelCommand.ActiveTestCase
	elementUuidTobeSwappedOut := incomingChannelCommand.FirstParameter
	elementUuidToBeSwappedIn := incomingChannelCommand.SecondParameter
	elementType := incomingChannelCommand.ElementType

	// Get the ImmatureElement To Swap In
	var immatureElementToSwapInTestCaseFormat testCaseModel.ImmatureElementStruct

	switch elementType {

	case sharedCode.TestInstruction:
		tempMap := commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap
		immatureElementToSwapInOriginal := tempMap[elementUuidToBeSwappedIn].ImmatureSubTestCaseModel
		immatureElementToSwapInTestCaseFormat = commandAndRuleEngine.convertGrpcElementModelIntoTestCaseElementModel(immatureElementToSwapInOriginal)

		// Handle DropZones
		availableDropZones := tempMap[elementUuidToBeSwappedIn].ImmatureTestInstructionInformation.AvailableDropZones
		numberOfDropZones := len(availableDropZones)

		// If there are more than one DropZone then pop up window so user can choose
		if numberOfDropZones > 1 {
			numberOfDropZones = 2
		}

		// Create DropZone window
		dropZoneContainer := container.NewVBox()
		var dropZoneWindow dialog.Dialog
		var dropZoneWaitGroup sync.WaitGroup
		dropZoneWaitGroup.Add(1)
		chosenDropZoneName := "NO_DROPZONE_WAS_CHOSEN"

		switch numberOfDropZones {
		case 0:

			// Set the TestInstructionColor to transparent
			//immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = "#00000000"
			// All TestInstructions have their initial values already set in UI-field for color
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = "No DropZone exists"
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = "No DropZone exists"

		case 1:
			// Move the uuid and color for the only DropZone onto all TestInstructions color-choice for UI
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = availableDropZones[0].DropZoneUuid
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = availableDropZones[0].DropZoneName
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = availableDropZones[0].DropZoneColor

		case 2:

			for _, dropZoneItem := range availableDropZones {

				dropZoneButton := &widget.Button{
					DisableableWidget: widget.DisableableWidget{},
					Text:              dropZoneItem.DropZoneName,
					Icon:              nil,
					Importance:        0,
					Alignment:         0,
					IconPlacement:     0,
					OnTapped:          nil,
				}

				// Create the Background colored rectangle
				dropZoneColor, err := sharedCode.ConvertRGBAHexStringIntoRGBAColor(dropZoneItem.DropZoneColor)
				if err != nil {
					return
				}
				backgroundColorRectangle := canvas.NewRectangle(dropZoneColor)

				// Create the DropZone Button
				dropZoneButton.OnTapped = func() {
					fmt.Println(dropZoneButton.Text)
					chosenDropZoneName = dropZoneButton.Text
					defer dropZoneWaitGroup.Done()
				}

				// Create text background rectangle for text to be more visible
				textbackgroundColor := color.RGBA{
					R: 0x33,
					G: 0x33,
					B: 0x33,
					A: 0xFF,
				}
				textBackgroundColorRectangle := canvas.NewRectangle(textbackgroundColor)

				// Create the DropZoneContainer
				dropZoneButtonContainer := container.NewMax(textBackgroundColorRectangle, dropZoneButton)
				dropZoneButtonContainer2 := container.NewHBox(layout.NewSpacer(), dropZoneButtonContainer, layout.NewSpacer())
				dropZoneButtonContainer2b := container.NewVBox(layout.NewSpacer(), dropZoneButtonContainer2, layout.NewSpacer())
				dropZoneButtonContainer3 := container.NewMax(backgroundColorRectangle, dropZoneButtonContainer2b)
				dropZoneButtonContainer3.Refresh()
				backgroundColorRectangle.SetMinSize(fyne.NewSize(0, dropZoneButton.Size().Height*1.3))
				textBackgroundColorRectangle.SetMinSize(fyne.NewSize(0, dropZoneButton.Size().Height*0.7))
				dropZoneButtonContainer3.Refresh()

				// Add the DropZoneButton-container to the object to be put up for user to chose DropZone from
				dropZoneContainer.Add(dropZoneButtonContainer3)

			}

			dropZoneContainer.Add(widget.NewSeparator())

			// Open up the DropZone-choser to the user
			dropZoneWindow = dialog.NewCustom("Choose DropZone", "Cancel", dropZoneContainer, *commandAndRuleEngine.MasterFenixWindow)
			dropZoneWindow.Show()

			// Wait for user to choose a DropZone
			dropZoneWaitGroup.Wait()
			dropZoneWindow.Hide()

			// If no DropZone was chosen then just exit
			if chosenDropZoneName == "NO_DROPZONE_WAS_CHOSEN" {
				return
			}

			// Find correct DropZone
			var dropZoneColorAsHexString string
			var dropZoneUuid string
			var dropZoneName string
			for _, dropZoneItem := range availableDropZones {

				if dropZoneItem.DropZoneName == chosenDropZoneName {
					dropZoneColorAsHexString = dropZoneItem.DropZoneColor
					dropZoneUuid = dropZoneItem.DropZoneUuid
					dropZoneName = dropZoneItem.DropZoneName
					break
				}

			}

			// Set the DropZoneUuid and TestInstructionColor from Chosen DropZone
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = dropZoneUuid
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = dropZoneName
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = dropZoneColorAsHexString

		}

		/*

			// If number of DropZones for Droppen object is greater than one the Open DropZone-picker Window
			dropZone1 := widget.NewButton("DropZone 1", func() {
				log.Println("Chosen DropZone: is One")
			})

			stateMachineDragAndDrop. stateMachineDragAndDrop.SourceUuid

			dropZoneColor, err :=
			dropZone1Container := container.NewMax(canvas.NewRectangle(color.Gray{}), dropZone1)

			dropZone2 := widget.NewButton("DropZone 1", func() {
				log.Println("Chosen DropZone: is Two")
			})
			dropZoneContainer := container.NewVBox(dropZone1, dropZone2)

			dropZoneWindow := dialog.NewCustom("Choose DropZone", "Cancel", dropZoneContainer, *masterFenixWindow)
			dropZoneWindow.Show()
		*/

	case sharedCode.TestInstructionContainer:
		tempMap := commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap
		immatureElementToSwapInOriginal := tempMap[elementUuidToBeSwappedIn].ImmatureSubTestCaseModel
		immatureElementToSwapInTestCaseFormat = commandAndRuleEngine.convertGrpcElementModelIntoTestCaseElementModel(immatureElementToSwapInOriginal)

		// Handle DropZones
		availableDropZones := tempMap[elementUuidToBeSwappedIn].ImmatureTestInstructionContainerInformation.AvailableDropZones
		numberOfDropZones := len(availableDropZones)

		// If there are more than one DropZone then pop up window so user can choose
		if numberOfDropZones > 1 {
			numberOfDropZones = 2
		}

		// Create DropZone window
		dropZoneContainer := container.NewVBox()
		var dropZoneWindow dialog.Dialog
		var dropZoneWaitGroup sync.WaitGroup
		dropZoneWaitGroup.Add(1)
		chosenDropZoneName := "NO_DROPZONE_WAS_CHOSEN"

		switch numberOfDropZones {
		case 0:

			// Set the TestInstructionColor to transparent
			//immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = "#00000000"
			// All TestInstructions have their initial values already set in UI-field for color
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = "No DropZone exists"
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = "No DropZone exists"

		case 1:
			// Move the uuid and color for the only DropZone onto all TestInstructions color-choice for UI
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = availableDropZones[0].DropZoneUuid
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = availableDropZones[0].DropZoneName
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = availableDropZones[0].DropZoneColor

		case 2:

			for _, dropZoneItem := range availableDropZones {

				dropZoneButton := &widget.Button{
					DisableableWidget: widget.DisableableWidget{},
					Text:              dropZoneItem.DropZoneName,
					Icon:              nil,
					Importance:        0,
					Alignment:         0,
					IconPlacement:     0,
					OnTapped:          nil,
				}

				// Create the Background colored rectangle
				dropZoneColor, err := sharedCode.ConvertRGBAHexStringIntoRGBAColor(dropZoneItem.DropZoneColor)
				if err != nil {
					return
				}
				backgroundColorRectangle := canvas.NewRectangle(dropZoneColor)

				// Create the DropZone Button
				dropZoneButton.OnTapped = func() {
					fmt.Println(dropZoneButton.Text)
					chosenDropZoneName = dropZoneButton.Text
					defer dropZoneWaitGroup.Done()
				}

				// Create text background rectangle for text to be more visible
				textbackgroundColor := color.RGBA{
					R: 0x33,
					G: 0x33,
					B: 0x33,
					A: 0xFF,
				}
				textBackgroundColorRectangle := canvas.NewRectangle(textbackgroundColor)

				// Create the DropZoneContainer
				dropZoneButtonContainer := container.NewMax(textBackgroundColorRectangle, dropZoneButton)
				dropZoneButtonContainer2 := container.NewHBox(layout.NewSpacer(), dropZoneButtonContainer, layout.NewSpacer())
				dropZoneButtonContainer2b := container.NewVBox(layout.NewSpacer(), dropZoneButtonContainer2, layout.NewSpacer())
				dropZoneButtonContainer3 := container.NewMax(backgroundColorRectangle, dropZoneButtonContainer2b)
				dropZoneButtonContainer3.Refresh()
				backgroundColorRectangle.SetMinSize(fyne.NewSize(0, dropZoneButton.Size().Height*1.3))
				textBackgroundColorRectangle.SetMinSize(fyne.NewSize(0, dropZoneButton.Size().Height*0.7))
				dropZoneButtonContainer3.Refresh()

				// Add the DropZoneButton-container to the object to be put up for user to chose DropZone from
				dropZoneContainer.Add(dropZoneButtonContainer3)

			}

			dropZoneContainer.Add(widget.NewSeparator())

			// Open up the DropZone-choser to the user
			dropZoneWindow = dialog.NewCustom("Choose DropZone", "Cancel", dropZoneContainer, *commandAndRuleEngine.MasterFenixWindow)
			dropZoneWindow.Show()

			// Wait for user to choose a DropZone
			dropZoneWaitGroup.Wait()
			dropZoneWindow.Hide()

			// If no DropZone was chosen then just exit
			if chosenDropZoneName == "NO_DROPZONE_WAS_CHOSEN" {
				return
			}

			// Find correct DropZone
			var dropZoneColorAsHexString string
			var dropZoneUuid string
			var dropZoneName string
			for _, dropZoneItem := range availableDropZones {

				if dropZoneItem.DropZoneName == chosenDropZoneName {
					dropZoneColorAsHexString = dropZoneItem.DropZoneColor
					dropZoneUuid = dropZoneItem.DropZoneUuid
					dropZoneName = dropZoneItem.DropZoneName
					break
				}

			}

			// Set the DropZoneUuid and TestInstructionColor from Chosen DropZone
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = dropZoneUuid
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = dropZoneName
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = dropZoneColorAsHexString

		}

	default:

		errorId := "6e8ed2ec-84df-42eb-a95d-41ba6920a9cb"
		err := errors.New(fmt.Sprintf("unknown Building BLock Type: '%s' [ErrorID: %s]", elementType, errorId))

		fmt.Println(err) //TODO Send error over error-channel

		// Exit function
		return

	}

	// Execute Swap of Elements
	err := commandAndRuleEngine.SwapElementsInTestCaseModel(currentTestCaseUuid, elementUuidTobeSwappedOut, &immatureElementToSwapInTestCaseFormat)
	if err != nil {
		fmt.Println(err)

		return //TODO Send error over error-channel
	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(currentTestCaseUuid)
	if err != nil {
		fmt.Println(err) //TODO Send error over error-channel
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedUpdateTestCaseGraphics,
		CreateNewTestCaseUI:          false,
		ActiveTestCase:               currentTestCaseUuid,
		TextualTestCaseSimple:        textualTestCaseSimple,
		TextualTestCaseComplex:       textualTestCaseComplex,
		TextualTestCaseExtended:      textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

	/*
		// Send 'Select TestInstruction in graphics' command over channel
		outgoingChannelCommandGraphicsUpdatedDataForTestInstruction := sharedCode.ChannelCommandGraphicsUpdatedStruct{
			ChannelCommandGraphicsUpdate: sharedCode.ChannelCommandGraphicsUpdatedSelectTestInstruction,
			CreateNewTestCaseUI:          false,
			ActiveTestCase:               currentTestCaseUuid,
			TextualTestCaseSimple:        textualTestCaseSimple,
			TextualTestCaseComplex:       textualTestCaseComplex,
			TextualTestCaseExtended:      textualTestCaseExtended,
			TestInstructionUuid:
		}

		*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedDataForTestInstruction
	*/

}

// Convert gRPC-message for TI or TIC into model used within the TestCase-model
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) convertGrpcElementModelIntoTestCaseElementModel(
	immatureGrpcElementModelMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage) (
	immatureTestCaseElementModel testCaseModel.ImmatureElementStruct) {

	// Initiate map used in TestCaseModel
	immatureTestCaseElementModel.ImmatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Loop over gRPC-element-model-structure
	for _, gRpcElementModel := range immatureGrpcElementModelMessage.TestCaseModelElements {
		immatureTestCaseElementModel.ImmatureElementMap[gRpcElementModel.ImmatureElementUuid] = *gRpcElementModel

	}

	// Set the first Element
	immatureTestCaseElementModel.FirstElementUuid = immatureGrpcElementModelMessage.FirstImmatureElementUuid

	return immatureTestCaseElementModel
}
