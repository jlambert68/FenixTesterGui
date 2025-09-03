package testSuitesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/soundEngine"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/jinzhu/copier"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"strings"
)

// SaveTestSuite
// Send TestSuite to TestCaseBuilderServer to saved in the Database
func (testSuiteModel *TestSuiteModelStruct) SaveTestSuite() (err error) {

	var isTestSuiteChanged bool
	var mandatoryFieldsHaveValues bool
	var mandatoryFieldsHaveValuesNotificationText string
	var valuesToBeHashed []string
	var messageHash string

	// Check if TestSuite is Changed
	isTestSuiteChanged = testSuiteModel.IsTestSuiteChanged()

	// Inform user that TestSuite is not changed
	if isTestSuiteChanged == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Save TestSuite",
			Content: fmt.Sprintf("TestSuite is not changed"),
		})

		return nil
	}

	//testSuiteModel.NoneSavedTestSuiteUIModelBinding = createEmptyAndInitiatedTestSuiteModel(testSuiteModel.testCasesModel)

	// Due to that DeepCopier can't copy 'TestDataPtr' it needs to be handled separate(I don't know why)
	var tempTestDataPtr *testDataEngine.TestDataForGroupObjectStruct
	tempTestDataPtr = testSuiteModel.TestSuiteUIModelBinding.TestDataPtr

	testSuiteModel.TestSuiteUIModelBinding.TestDataPtr = nil

	// Copy fields from 'TestSuiteUIModelBinding' to  'NoneSavedTestSuiteUIModelBinding' using deep copy
	err = copier.CopyWithOption(&testSuiteModel.NoneSavedTestSuiteUIModelBinding, &testSuiteModel.TestSuiteUIModelBinding, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "85acf490-2599-4a8a-bf9e-1451eef60f78"

		errorMsg := fmt.Sprintf("error copying 'TestSuiteUIModelBinding' using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}

	// Set back 'TestDataPtr' and put 'pointer-copy' in 'NoneSavedTestSuiteUIModelBinding'
	testSuiteModel.TestSuiteUIModelBinding.TestDataPtr = tempTestDataPtr
	testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestDataPtr = tempTestDataPtr

	// Check if all mandatory fields has values
	mandatoryFieldsHaveValues,
		mandatoryFieldsHaveValuesNotificationText = testSuiteModel.checkIfAllMandatoryFieldsHaveValues()

	// Inform user that mandatory field in TestSuite is missing
	if mandatoryFieldsHaveValues == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.UserNeedToRespondSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title: "Mandatory Fields",
			Content: fmt.Sprintf("Mandatory field '%s' is empty",
				mandatoryFieldsHaveValuesNotificationText),
		})

		return nil
	}

	// Copy all UI-fields to model
	//testSuiteModel.copyUiFieldsToModel()

	// Keeps a list of what is to be saved. Used to ensure that older versions of the TestSuite can later be loaded when new functionality has been added to the client
	var supportedTestSuiteDataToBeStored testSuiteImplementedFunctionsToBeStoredStruct
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap = make(map[testSuiteImplementedFunctionsType]bool)

	// Generate 'TestSuiteBasicInformation' to be added to full gRPC-message
	var testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage
	var testSuiteBasicInformationHash string
	testSuiteBasicInformation, testSuiteBasicInformationHash, err = testSuiteModel.
		generateTestSuiteBasicInformationMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteBasicInformationHash)

	// Generate 'DeleteDate' to be added to full gRPC-message
	var testSuiteDeleteDate string
	var testSuiteDeleteDateHash string
	testSuiteDeleteDate, testSuiteDeleteDateHash, err = testSuiteModel.generateTestSuiteDeleteDateMessageWhenSaving(
		&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteDeleteDateHash)

	// Convert 'TestSuiteType' into gRPC-message
	var testSuiteType *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage
	var testSuiteTypeHash string
	testSuiteType, testSuiteTypeHash, err = testSuiteModel.
		generateTestSuiteTypeMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteTypeHash)

	// Convert 'TestSuiteMetaData' into gRPC-message
	var testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage
	var testSuiteMetaDataHash string
	testSuiteMetaData, testSuiteMetaDataHash, err = testSuiteModel.
		generateTestSuiteMetaDataMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteMetaDataHash)

	// Convert 'TestSuiteTestData' into gRPC-message
	var testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage
	var testSuiteTestDataHash string
	testSuiteTestData, testSuiteTestDataHash, err = testSuiteModel.
		generateTestSuiteTestDataMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteTestDataHash)

	// Convert 'TestCasesInTestSuite' into gRPC-message
	var testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage
	var testCasesInTestSuiteHash string
	testCasesInTestSuite, testCasesInTestSuiteHash, err = testSuiteModel.
		generateTestCasesInTestSuiteMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testCasesInTestSuiteHash)

	// Convert 'TestSuitePreview' into gRPC-message
	var testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage
	var testSuitePreviewHash string
	testSuitePreview, testCasesInTestSuiteHash, err = testSuiteModel.
		generateTestSuitePreviewMessageWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuitePreviewHash)

	// Convert 'supportedTestSuiteDataToBeStored' to be added to full gRPC-message
	var testSuiteImplementedFunctionsMap map[int32]bool
	var testSuiteImplementedFunctionsMapHash string
	testSuiteImplementedFunctionsMap = make(map[int32]bool)
	testSuiteImplementedFunctionsMap, testSuiteImplementedFunctionsMapHash, err = testSuiteModel.
		generateTestSuiteImplementedFunctionsMapWhenSaving(&supportedTestSuiteDataToBeStored)
	valuesToBeHashed = append(valuesToBeHashed, testSuiteImplementedFunctionsMapHash)

	// Create MessageHash
	messageHash = sharedCode.HashValues(valuesToBeHashed, false)

	// Generate full gRPC-message to be sent to TestCaseBuilder-Server
	var fullTestSuiteMessage fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage
	fullTestSuiteMessage = fenixGuiTestCaseBuilderServerGrpcApi.FullTestSuiteMessage{
		TestSuiteBasicInformation:        testSuiteBasicInformation,
		TestSuiteTestData:                testSuiteTestData,
		TestSuitePreview:                 testSuitePreview,
		TestSuiteMetaData:                testSuiteMetaData,
		TestCasesInTestSuite:             testCasesInTestSuite,
		DeletedDate:                      testSuiteDeleteDate,
		UpdatedByAndWhen:                 nil, // Used when loading TestSuite
		TestSuiteType:                    testSuiteType,
		TestSuiteImplementedFunctionsMap: testSuiteImplementedFunctionsMap,
		MessageHash:                      messageHash,
	}

	// Send using gRPC
	returnMessage := testSuiteModel.testCasesModel.GrpcOutReference.SendSaveFullTestSuite(&fullTestSuiteMessage)

	if returnMessage == nil || returnMessage.AckNack == false {

		// Trigger System Notification sound
		soundEngine.PlaySoundChannel <- soundEngine.InvalidNotificationSound

		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Save TestSuite",
			Content: fmt.Sprintf("Got some error when trying to Save TestSuite"),
		})

		if returnMessage == nil {
			errorId := "6c8be10e-404e-49a3-b121-df05c952d5b8"
			err = errors.New(fmt.Sprintf("Got 'nil' back when trying to Save TestSuite [ErrorID: %s]", errorId))

		} else {
			errorId := "5f887e6f-db34-4ca0-8abf-4521657c2337"
			err = errors.New(fmt.Sprintf(returnMessage.Comments+"[ErrorID: %s]", errorId))

		}

		fmt.Println(err) // TODO Send on Error-channel

		return err
	}

	// Trigger System Notification sound
	soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Save TestSuite",
		Content: fmt.Sprintf("Success in saving TestSuite"),
	})

	// Update that the TestSuite is not New anymore and Update the TestSuiteVersion
	testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteIsNew = false
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteIsNew = false

	// Due to that DeepCopier can't copy 'TestDataPtr' it needs to be handled separate(I don't know why)
	var tempTestDataPtr2 *testDataEngine.TestDataForGroupObjectStruct
	tempTestDataPtr2 = testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestDataPtr
	testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr = nil
	testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestDataPtr = nil

	// Copy data from 'NoneSavedTestSuiteUIModelBinding' to 'savedTestSuiteUIModelBinding' using deep copy
	err = copier.CopyWithOption(&testSuiteModel.savedTestSuiteUIModelBinding, &testSuiteModel.NoneSavedTestSuiteUIModelBinding, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "2304dabe-7e12-4cf4-a021-597793ace220"

		errorMsg := fmt.Sprintf("error copying 'NoneSavedTestSuiteUIModelBinding' using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}

	// Set back 'TestDataPtr' and put 'pointer-copy' in 'NoneSavedTestSuiteUIModelBinding'
	testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestDataPtr = tempTestDataPtr2
	testSuiteModel.savedTestSuiteUIModelBinding.TestDataPtr = tempTestDataPtr2

	// Update The Hash for the saved TestSuite
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteSavedMessageHash = messageHash

	// Set the new TestSuiteVersion
	testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion = testSuiteBasicInformation.TestSuiteVersion

	return err

}

// Generates 'TestSuiteBasicInformation' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteBasicInformationMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteBasicInformation *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage,
	testSuiteBasicInformationHash string,
	err error) {

	// This TestSuite has stored 'testSuiteBasicInformationIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteBasicInformationIsSupported] = true

	// Generate TestSuiteVersion
	var testSuiteVersion uint32
	if testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteIsNew == true {
		testSuiteVersion = 1
	} else {
		testSuiteVersion = testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteVersion + 1
	}

	// Create 'testSuiteBasicInformation'
	testSuiteBasicInformation = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteBasicInformationMessage{
		DomainUuid:                    testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
		DomainName:                    testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
		TestSuiteUuid:                 testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid,
		TestSuiteVersion:              testSuiteVersion,
		TestSuiteName:                 testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteName,
		TestSuiteDescription:          testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteDescription,
		TestSuiteExecutionEnvironment: testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteExecutionEnvironment,
	}

	// Create the Hash of the Message
	var tempJson string
	tempJson = protojson.Format(testSuiteBasicInformation)
	testSuiteBasicInformationHash = sharedCode.HashSingleValue(tempJson)

	return testSuiteBasicInformation, testSuiteBasicInformationHash, err

}

// Generates 'UsersChosenTestDataForTestSuiteMessage' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTestDataMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteTestData *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage,
	testSuiteTestDataHash string,
	err error) {

	// This TestSuite has stored 'testSuiteTestDataIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteTestDataIsSupported] = true

	// The gRPC-version of 'testDataPointNameMap'
	var chosenTestDataPointsPerGroupMapGrpc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestDataPointNameMapMessage
	chosenTestDataPointsPerGroupMapGrpc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestDataPointNameMapMessage)

	if testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestDataPtr != nil {

		// Loop TestDataGroups
		for testDataGroupName, testDataPointNameMap := range testSuiteModel.NoneSavedTestSuiteUIModelBinding.
			TestDataPtr.ChosenTestDataPointsPerGroupMap {

			var testDataPointNameMapMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestDataPointNameMapMessage
			testDataPointNameMapMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestDataPointNameMapMessage{}

			// The gRPC-version of 'testDataPointNameMap'
			var chosenTestDataRowsPerTestDataPointMapGprc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage
			chosenTestDataRowsPerTestDataPointMapGprc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage)

			// Extract TestDataPoints for the TestDataPointNameMap
			for testDataPointName, testDataPointsSlice := range *testDataPointNameMap {

				var testDataRowsGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage
				testDataRowsGrpc = &fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowsMessage{}

				// The gRPC-version of 'testDataPointsSlice'
				var testDataRowMessageSliceGrc []*fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowMessage

				// Loop the TestDataPoints
				for _, testDataPoint := range *testDataPointsSlice {

					// gRPC-version of 'SelectedTestDataPointUuidMap'
					var testDataPointRowValueSummaryMapGrpc map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
						TestDataPointRowValueSummaryMapMessage
					testDataPointRowValueSummaryMapGrpc = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
						TestDataPointRowValueSummaryMapMessage)

					// Loop the 'selected' TestDataRows in the TestDataPoint
					for _, testDataPointRowData := range testDataPoint.SelectedTestDataPointUuidMap {

						// Create gRPC-version of 'testDataPointRowData'
						var testDataPointRowDataGrpc *fenixGuiTestCaseBuilderServerGrpcApi.
							TestDataPointRowValueSummaryMapMessage
						testDataPointRowDataGrpc = &fenixGuiTestCaseBuilderServerGrpcApi.
							TestDataPointRowValueSummaryMapMessage{
							TestDataPointRowUuid:          string(testDataPointRowData.TestDataPointRowUuid),
							TestDataPointRowValuesSummary: string(testDataPointRowData.TestDataPointRowValuesSummary),
						}

						// Add 'testDataPointRowDataGrpc' to gRPC-version of 'SelectedTestDataPointUuidMap'
						testDataPointRowValueSummaryMapGrpc[string(testDataPointRowData.TestDataPointRowUuid)] =
							testDataPointRowDataGrpc
					}

					// Create the gRPC-version of 'testDataPoint'
					var testDataRowMessageGrpc *fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowMessage
					testDataRowMessageGrpc = &fenixGuiTestCaseBuilderServerGrpcApi.TestDataRowMessage{
						TestDataDomainUuid:              string(testDataPoint.TestDataDomainUuid),
						TestDataDomainName:              string(testDataPoint.TestDataDomainName),
						TestDataAreaUuid:                string(testDataPoint.TestDataAreaUuid),
						TestDataAreaName:                string(testDataPoint.TestDataAreaName),
						TestDataPointName:               string(testDataPoint.TestDataPointName),
						TestDataPointRowValueSummaryMap: testDataPointRowValueSummaryMapGrpc,
					}

					// Append to slice of TestDataPoint-message
					testDataRowMessageSliceGrc = append(testDataRowMessageSliceGrc, testDataRowMessageGrpc)
				}

				// Add 'testDataRowMessageSliceGrc' into "single message"
				testDataRowsGrpc.TestDataRows = testDataRowMessageSliceGrc

				// Store "single-message" with TestDataRows in map 'chosenTestDataRowsPerTestDataPointMapGprc'
				chosenTestDataRowsPerTestDataPointMapGprc[string(testDataPointName)] = testDataRowsGrpc

			}

			// And 'chosenTestDataRowsPerTestDataPointMapGprc' into "single message"
			testDataPointNameMapMessage.ChosenTestDataRowsPerTestDataPointMap = chosenTestDataRowsPerTestDataPointMapGprc

			// Store "singe-message with Group-data in map 'chosenTestDataPointsPerGroupMapGrpc'
			chosenTestDataPointsPerGroupMapGrpc[string(testDataGroupName)] = testDataPointNameMapMessage

		}
	}

	testSuiteTestData = &fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestSuiteMessage{
		ChosenTestDataPointsPerGroupMap: chosenTestDataPointsPerGroupMapGrpc,
		UsersSelectedTestDataPointRow: &fenixGuiTestCaseBuilderServerGrpcApi.UsersSelectedTestDataPointRowMessage{
			TestDataGroup:        "",
			TestDataPoint:        "",
			TestDataPointSummary: "",
		},
		HashOfThisMessageWithEmptyHashField: "",
	}

	// Generate Hash of gRPC-message and add it to the message
	tempJson := protojson.Format(testSuiteTestData)

	// Remove spaces before hashing, due to some bug that generates "double space" sometimes when running in non-debug-mode
	tempJson = strings.ReplaceAll(tempJson, " ", "")

	testSuiteTestDataHash = sharedCode.HashSingleValue(tempJson)
	testSuiteTestData.HashOfThisMessageWithEmptyHashField = testSuiteTestDataHash

	return testSuiteTestData, testSuiteTestDataHash, err

}

// Generates 'TestSuitePreview' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuitePreviewMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage,
	testSuitePreviewHash string,
	err error) {

	// This TestSuite has stored 'testSuitePreviewSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuitePreviewSupported] = true

	// Create final SelectedTestSuiteMetaDataValuesMap
	var selectedTestSuiteMetaDataValuesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage
	selectedTestSuiteMetaDataValuesMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage)

	// Get the MetaData for the TestSuite
	var tempTestSuiteMetaData TestSuiteMetaDataStruct
	tempTestSuiteMetaData = *testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteMetaDataPtr

	// Get the MetaDataGroupsMap
	var tempMetaDataGroupsMap map[string]*MetaDataGroupStruct
	tempMetaDataGroupsMap = *tempTestSuiteMetaData.MetaDataGroupsMapPtr

	var selectedMetaDataValuesMapKey string

	// Loop all MetaDataGroups
	for _, tempMetaDataGroupPtr := range tempMetaDataGroupsMap {

		// Get the MetaDataGroup
		var tempMetaDataGroup MetaDataGroupStruct
		tempMetaDataGroup = *tempMetaDataGroupPtr

		// Loop all MetaDataGroupItems
		for _, tempMetaDataGroupItemPtr := range *tempMetaDataGroup.MetaDataInGroupMapPtr {

			// Get the MetaDataGroupItem
			var tempMetaDataGroupItem MetaDataInGroupStruct
			tempMetaDataGroupItem = *tempMetaDataGroupItemPtr

			// Is the Item a 'single select' or a 'multi select'
			switch tempMetaDataGroupItem.SelectType {

			case MetaDataSelectType_SingleSelect:

				if len(tempMetaDataGroupItem.SelectedMetaDataValueForSingleSelect) > 0 {
					// Add selected value to the 'SelectedMetaDataValuesMap'

					// Create the map-key
					// The Map key = 'OwnerDomainUuid.MetaDataGroupName.MetaDataName.MetaDataNameValue'
					selectedMetaDataValuesMapKey = fmt.Sprintf("%s.%s.%s.%s",
						testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
						tempMetaDataGroupItem.MetaDataGroupName,
						tempMetaDataGroupItem.MetaDataName,
						tempMetaDataGroupItem.SelectedMetaDataValueForSingleSelect)

					// Create the value to be inserted into the map
					var tempSelectedTestSuiteMetaDataValueMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage
					tempSelectedTestSuiteMetaDataValueMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage{
						OwnerDomainUuid:   testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
						OwnerDomainName:   testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
						MetaDataGroupName: tempMetaDataGroupItem.MetaDataGroupName,
						MetaDataName:      tempMetaDataGroupItem.MetaDataName,
						MetaDataNameValue: tempMetaDataGroupItem.SelectedMetaDataValueForSingleSelect,
						SelectType:        fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum(tempMetaDataGroupItem.SelectType),
						IsMandatory:       tempMetaDataGroupItem.Mandatory,
					}

					// Add selected value to the 'SelectedMetaDataValuesMap'
					selectedTestSuiteMetaDataValuesMap[selectedMetaDataValuesMapKey] = tempSelectedTestSuiteMetaDataValueMessage
				}

			case MetaDataSelectType_MultiSelect:

				if len(tempMetaDataGroupItem.SelectedMetaDataValuesForMultiSelect) > 0 {
					// Add selected value to the 'SelectedMetaDataValuesMap'

					// Loop SelectedMetaDataValuesForMultiSelect
					for _, tempSelectedMetaDataValueForMultiSelect := range tempMetaDataGroupItem.SelectedMetaDataValuesForMultiSelect {

						// Create the map-key
						selectedMetaDataValuesMapKey = fmt.Sprintf("%s.%s.%s.%s",
							testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
							tempMetaDataGroupItem.MetaDataGroupName,
							tempMetaDataGroupItem.MetaDataName,
							tempSelectedMetaDataValueForMultiSelect)

						// Create the value to be inserted into the map
						var tempSelectedTestSuiteMetaDataValueMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage
						tempSelectedTestSuiteMetaDataValueMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage{
							OwnerDomainUuid:   testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
							OwnerDomainName:   testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
							MetaDataGroupName: tempMetaDataGroupItem.MetaDataGroupName,
							MetaDataName:      tempMetaDataGroupItem.MetaDataName,
							MetaDataNameValue: tempSelectedMetaDataValueForMultiSelect,
							SelectType:        fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum(tempMetaDataGroupItem.SelectType),
							IsMandatory:       tempMetaDataGroupItem.Mandatory,
						}

						// Add selected value to the 'SelectedMetaDataValuesMap'
						selectedTestSuiteMetaDataValuesMap[selectedMetaDataValuesMapKey] = tempSelectedTestSuiteMetaDataValueMessage
					}
				}

			default:

				errorId := "34089b89-a29b-4a56-b6dc-026cceeacc77"

				errorMessage := fmt.Sprintf("Unknown SelectType for MetaDataGroupItem '%d' [ErrorID: %s]",
					tempMetaDataGroupItem.SelectType, errorId)

				log.Fatal(errorMessage)

			}

		}

	}

	// Create final TestSuitePreviewStructure-message
	var testSuitePreviewStructureMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage
	testSuitePreviewStructureMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage{
		TestSuiteUuid:                      testSuiteModel.testSuiteModelDataThatCanNotBeChangedFromUI.testSuiteUuid,
		TestSuiteName:                      testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteName,
		TestSuiteVersion:                   "0",
		DomainUuidThatOwnTheTestSuite:      testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
		DomainNameThatOwnTheTestSuite:      testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
		TestSuiteDescription:               testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteDescription,
		TestCasesInTestSuite:               nil,
		CreatedByGcpLoginUser:              testSuiteModel.GetLastChangedByGcpLogin(),
		CreatedByComputerLoginUser:         testSuiteModel.GetLastChangedByComputerLogin(),
		CreatedDate:                        testSuiteModel.GetCreatedDate(),
		LastSavedByUserOnComputer:          testSuiteModel.GetLastChangedByComputerLogin(),
		LastSavedByUserGCPAuthorization:    testSuiteModel.GetLastChangedByGcpLogin(),
		LastSavedTimeStamp:                 "",
		TestSuiteType:                      nil,
		SelectedTestSuiteMetaDataValuesMap: selectedTestSuiteMetaDataValuesMap,
	}

	// Create final TestSuitePreview-message
	testSuitePreview = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewMessage{
		TestSuitePreview:     testSuitePreviewStructureMessage,
		TestSuitePreviewHash: testSuitePreviewHash,
	}

	return testSuitePreview, testSuitePreviewHash, err

}

// Generates 'TestSuiteMetaData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteMetaDataMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteMetaData *fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage,
	testSuiteMetaDataHash string,
	err error) {

	// Initiate TestSuiteMetaData
	testSuiteMetaData = &fenixGuiTestCaseBuilderServerGrpcApi.UserSpecifiedTestSuiteMetaDataMessage{
		CurrentSelectedDomainUuid: testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainUuid,
		CurrentSelectedDomainName: testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteOwnerDomainName,
		MetaDataGroupsMap:         make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage),
	}

	// This TestSuite has stored 'testSuiteMetaDataIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteMetaDataIsSupported] = true

	// SLice holding the values that will become the MetaDataSlice
	var valuesToBeHashedSlice []string
	var valueToBeHashed string

	// Check if there are any TestCaseMetaData, if not then just return
	if testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteMetaDataPtr == nil ||
		testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteMetaDataPtr.MetaDataGroupsMapPtr == nil {

		testSuiteMetaDataHash = sharedCode.HashSingleValue("")

		return testSuiteMetaData,
			testSuiteMetaDataHash,
			err
	}

	var tempMetaDataGroupsMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage
	tempMetaDataGroupsMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage)

	// Get the TestCaseMetaData from TestCaseMetaDataPtr
	//var TestCaseMetaData

	// Loop MetaDataGroups in the TestCase and extract each MetaDataGroup
	for tempMetaGroupNameInTestCase, tempMetaDataGroupInTestCasePtr := range *testSuiteModel.
		NoneSavedTestSuiteUIModelBinding.TestSuiteMetaDataPtr.MetaDataGroupsMapPtr {

		// Get the MetaDataGroupInTestCaseMap
		var tempMetaDataGroupInTestCaseMap map[string]*MetaDataInGroupStruct
		tempMetaDataGroupInTestCaseMap = *tempMetaDataGroupInTestCasePtr.MetaDataInGroupMapPtr

		// Create the MetaDataInGroupItem-map
		var metaDataInGroupMessage map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage
		metaDataInGroupMessage = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage)

		// Loop items in MetaDataGroup
		for tempMetaDataGroupItemNameInTestCase, tempMetaDataGroupItemInTestCase := range tempMetaDataGroupInTestCaseMap {

			// Create the MetaDataGroupItem for the gRPC-message
			var metaDataGroupItem fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage
			metaDataGroupItem = fenixGuiTestCaseBuilderServerGrpcApi.MetaDataInGroupMessage{
				MetaDataGroupName:                       tempMetaGroupNameInTestCase,
				MetaDataName:                            tempMetaDataGroupItemNameInTestCase,
				SelectType:                              fenixGuiTestCaseBuilderServerGrpcApi.MetaDataSelectTypeEnum(tempMetaDataGroupItemInTestCase.SelectType),
				IsMandatory:                             tempMetaDataGroupItemInTestCase.Mandatory,
				AvailableMetaDataValues:                 tempMetaDataGroupItemInTestCase.AvailableMetaDataValues,
				SelectedMetaDataValueForSingleSelect:    tempMetaDataGroupItemInTestCase.SelectedMetaDataValueForSingleSelect,
				SelectedMetaDataValuesForMultiSelect:    tempMetaDataGroupItemInTestCase.SelectedMetaDataValuesForMultiSelect,
				SelectedMetaDataValuesForMultiSelectMap: *tempMetaDataGroupItemInTestCase.SelectedMetaDataValuesForMultiSelectMapPtr,
			}

			// Add the MetaDataGroupItem for the gRPC-map-message
			metaDataInGroupMessage[tempMetaDataGroupItemNameInTestCase] = &metaDataGroupItem

			// Generate shared values, for SingleSelect and MultiSelect, to be hashed
			valueToBeHashed = fmt.Sprintf("%s-%s-%t-%d",
				tempMetaGroupNameInTestCase,
				tempMetaDataGroupItemNameInTestCase,
				tempMetaDataGroupItemInTestCase.Mandatory,
				tempMetaDataGroupItemInTestCase.SelectType)

			valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)

			// Add Available values to be hashed
			for _, tempMetaDataValue := range tempMetaDataGroupItemInTestCase.AvailableMetaDataValues {
				valueToBeHashed = fmt.Sprintf("%s",
					tempMetaDataValue)

				valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)
			}

			// Generate values to be hashed, depending on Single or Multiple Select
			switch tempMetaDataGroupItemInTestCase.SelectType {
			case MetaDataSelectType_SingleSelect:

				valueToBeHashed = fmt.Sprintf("%s",
					tempMetaDataGroupItemInTestCase.SelectedMetaDataValueForSingleSelect)

				valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)

			case MetaDataSelectType_MultiSelect:

				for _, tempMetaDataValue := range tempMetaDataGroupItemInTestCase.SelectedMetaDataValuesForMultiSelect {
					valueToBeHashed = fmt.Sprintf("%s",
						tempMetaDataValue)

					valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)
				}

			default:
				errorId := ""

				log.Fatalln(fmt.Sprintf("Unhandled tempMetaDataGroupItemInTestCase.SelectType",
					tempMetaDataGroupItemInTestCase.SelectType,
					errorId))
			}

		}

		// Create the MetaDataGroupMessage
		var metaDataGroupMessage fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage
		metaDataGroupMessage = fenixGuiTestCaseBuilderServerGrpcApi.MetaDataGroupMessage{
			MetaDataGroupName:  tempMetaGroupNameInTestCase,
			MetaDataInGroupMap: metaDataInGroupMessage,
		}

		// Add the MetaDataGroup-message for the gRPC-map-message
		tempMetaDataGroupsMap[tempMetaGroupNameInTestCase] = &metaDataGroupMessage

	}

	// Add the MetaDataGroupsMap to the gRPC-object
	testSuiteMetaData.MetaDataGroupsMap = tempMetaDataGroupsMap

	// Generate Hash of all sub-message-hashes
	testSuiteMetaDataHash = sharedCode.HashValues(valuesToBeHashedSlice, false)

	return testSuiteMetaData,
		testSuiteMetaDataHash,
		err

}

// Generates 'TestCasesInTestSuite' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestCasesInTestSuiteMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testCasesInTestSuite *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage,
	testCasesInTestSuiteHash string,
	err error) {

	var valuesToBeHashedSlice []string
	var valueToBeHashed string

	// Initiate 'testCasesInTestSuite'
	testCasesInTestSuite = &fenixGuiTestCaseBuilderServerGrpcApi.TestCasesInTestSuiteMessage{
		TestCasesInTestSuite: nil,
	}

	// This TestSuite has stored 'testCasesInTestSuiteIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testCasesInTestSuiteIsSupported] = true

	// Get the TestCasesMap
	var testCasesInTestSuiteMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
	testCasesInTestSuiteMap = *testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestCasesInTestSuitePtr.TestCasesInTestSuiteMapPtr

	// Loop added TestCases and add to gRPC-message
	if testCasesInTestSuiteMap != nil {
		for _, testCasesInTestSuiteMessage := range testCasesInTestSuiteMap {

			var tempTestCaseInTestSuiteMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage
			tempTestCaseInTestSuiteMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestCaseInTestSuiteMessage{
				DomainUuid:   testCasesInTestSuiteMessage.GetDomainUuid(),
				DomainName:   testCasesInTestSuiteMessage.GetDomainName(),
				TestCaseUuid: testCasesInTestSuiteMessage.GetTestCaseUuid(),
				TestCaseName: testCasesInTestSuiteMessage.GetTestCaseName(),
			}

			testCasesInTestSuite.TestCasesInTestSuite = append(testCasesInTestSuite.TestCasesInTestSuite, tempTestCaseInTestSuiteMessage)

			// Create value to be hashed
			valueToBeHashed = fmt.Sprintf("%s-%s-%s-%s",
				testCasesInTestSuiteMessage.GetDomainUuid(),
				testCasesInTestSuiteMessage.GetDomainName(),
				testCasesInTestSuiteMessage.GetTestCaseUuid(),
				testCasesInTestSuiteMessage.GetTestCaseName())

			valuesToBeHashedSlice = append(valuesToBeHashedSlice, valueToBeHashed)

		}
	}

	// Create TestCasesInTestSuiteHash
	if len(valuesToBeHashedSlice) > 0 {
		testCasesInTestSuiteHash = sharedCode.HashValues(valuesToBeHashedSlice, true)
	} else {
		testCasesInTestSuiteHash = sharedCode.HashSingleValue("")
	}

	return testCasesInTestSuite, testCasesInTestSuiteHash, err

}

// Generates 'TestSuiteDeleteData' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteDeleteDateMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteDeleteDate string,
	testSuiteDeleteDateHash string,
	err error) {

	// This TestSuite has stored 'deletedDateIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[deletedDateIsSupported] = true

	// Create 'testSuiteDeleteDate'
	testSuiteDeleteDate = testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteDeletionDate

	// Create the Hash of the Message
	testSuiteDeleteDateHash = sharedCode.HashSingleValue(testSuiteDeleteDate)

	return testSuiteDeleteDate, testSuiteDeleteDateHash, err

}

// Generates 'TestCasesInTestSuite' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteTypeMessageWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteTypeMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage,
	testSuiteTypeHash string,
	err error) {

	// This TestSuite has stored 'testSuiteTypeIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteTypeIsSupported] = true

	// Create 'testSuiteTypeMessage'
	testSuiteTypeMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeMessage{
		TestSuiteType: fenixGuiTestCaseBuilderServerGrpcApi.TestSuiteTypeEnum(
			testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteType.TestSuiteType),
		TestSuiteTypeName: string(testSuiteModel.NoneSavedTestSuiteUIModelBinding.TestSuiteType.TestSuiteTypeName),
	}

	// Create Hash for 'testSuiteTypeIsSupported'
	var valuesToBeHashed []string
	valuesToBeHashed = append(valuesToBeHashed, fmt.Sprintf("%d", testSuiteTypeMessage.TestSuiteType))
	valuesToBeHashed = append(valuesToBeHashed, fmt.Sprintf("%s", testSuiteTypeMessage.TestSuiteTypeName))

	// Create the Hash of the Message
	testSuiteTypeHash = sharedCode.HashValues(valuesToBeHashed, true)

	return testSuiteTypeMessage, testSuiteTypeHash, err

}

// Generates 'TestSuiteImplementedFunctionsMap' to be added to full gRPC-message
func (testSuiteModel *TestSuiteModelStruct) generateTestSuiteImplementedFunctionsMapWhenSaving(
	supportedTestSuiteDataToBeStored *testSuiteImplementedFunctionsToBeStoredStruct) (
	testSuiteImplementedFunctionsMap map[int32]bool,
	testSuiteImplementedFunctionsMapHash string,
	err error) {

	// This TestSuite has stored 'testSuiteImplementedFunctionsMapIsSupported'
	supportedTestSuiteDataToBeStored.testSuiteImplementedFunctionsMap[testSuiteImplementedFunctionsMapIsSupported] = true

	// Create 'testSuiteImplementedFunctionsMap'
	testSuiteImplementedFunctionsMap = make(map[int32]bool)

	// Loop 'supportedTestSuiteDataToBeStored' and convert to gRPC-message
	var valuesToBeHashed []string
	for testSuiteImplementedFunction, testSuiteImplementedFunctionValue := range supportedTestSuiteDataToBeStored.
		testSuiteImplementedFunctionsMap {

		testSuiteImplementedFunctionsMap[int32(testSuiteImplementedFunction)] = testSuiteImplementedFunctionValue
		valuesToBeHashed = append(valuesToBeHashed, fmt.Sprintf("%d", testSuiteImplementedFunction))
	}

	// Create the Hash of the Message
	testSuiteImplementedFunctionsMapHash = sharedCode.HashValues(valuesToBeHashed, true)

	return testSuiteImplementedFunctionsMap, testSuiteImplementedFunctionsMapHash, err

}
