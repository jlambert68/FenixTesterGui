package testCaseModel

import (
	"FenixTesterGui/importFilesFromGitHub"
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
)

// LoadFullTestCaseFromDatabase - Load the TestCase from the Database into model
func (testCaseModel *TestCasesModelsStruct) LoadFullTestCaseFromDatabase(
	testCaseUuid string) (
	err error) {

	// Send LoadTesCase using gRPC
	var detailedTestCaseResponse *fenixGuiTestCaseBuilderServerGrpcApi.GetDetailedTestCaseResponse
	detailedTestCaseResponse = testCaseModel.GrpcOutReference.LoadDetailedTestCase(testCaseUuid)

	// Exit if something was wrong
	if detailedTestCaseResponse.AckNackResponse.AckNack == false {

		errorId := "ba195459-8902-4727-ab81-ae48cd616eea"
		err = errors.New(fmt.Sprintf(detailedTestCaseResponse.AckNackResponse.Comments+"[ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return err

	}

	// Create object that will hold complete TestCase in memory
	var tempTestCaseModel TestCaseModelStruct
	tempTestCaseModel = TestCaseModelStruct{
		LastLoadedTestCaseModelGRPCMessage:         *detailedTestCaseResponse.DetailedTestCase.TestCaseBasicInformation.TestCaseModel, // fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelMessage{},
		FirstElementUuid:                           detailedTestCaseResponse.DetailedTestCase.TestCaseBasicInformation.TestCaseModel.FirstMatureElementUuid,
		TestCaseModelMap:                           nil, // Created below
		TextualTestCaseRepresentationSimpleStack:   nil, // Created below
		TextualTestCaseRepresentationComplexStack:  nil, // Created below
		TextualTestCaseRepresentationExtendedStack: nil, // Created below
		CommandStack:                               nil,
		LastSavedCommandStack:                      lastSavedCommandStack{},
		CopyBuffer:                                 ImmatureElementStruct{},
		CutBuffer:                                  MatureElementStruct{},
		CutCommandInitiated:                        false,
		LocalTestCaseMessage: LocalTestCaseMessageStruct{
			BasicTestCaseInformationMessageNoneEditableInformation: *detailedTestCaseResponse.DetailedTestCase.
				TestCaseBasicInformation.BasicTestCaseInformation.GetNonEditableInformation(),
			BasicTestCaseInformationMessageEditableInformation: *detailedTestCaseResponse.DetailedTestCase.
				TestCaseBasicInformation.BasicTestCaseInformation.GetEditableInformation(),
			CreatedAndUpdatedInformation: *detailedTestCaseResponse.DetailedTestCase.
				TestCaseBasicInformation.GetCreatedAndUpdatedInformation(),
			DeleteTimeStamp: detailedTestCaseResponse.GetDetailedTestCase().GetDeletedDate(),
		},
		testCaseModelAdaptedForUiTree:            nil,
		CurrentSelectedTestCaseElement:           CurrentSelectedTestCaseElementStruct{},
		MatureTestInstructionMap:                 nil, // Created below
		MatureTestInstructionContainerMap:        nil, // Created below
		AttributesList:                           nil, // Initialized below
		ThisIsANewTestCase:                       false,
		TestCaseHash:                             detailedTestCaseResponse.DetailedTestCase.MessageHash,
		TestCaseHashWhenTestCaseWasSavedOrLoaded: "",
		TestDataHash:                             detailedTestCaseResponse.DetailedTestCase.GetTestCaseTestData().GetHashOfThisMessageWithEmptyHashField(),
		TestDataHashWhenTestCaseWasSavedOrLoaded: "",
		ImportedTemplateFilesFromGitHub:          nil,
		TestData:                                 nil,
		TestCasePreviewObject:                    nil,
		TestCaseMetaDataPtr:                      nil,
	}

	// Initialize AttributesList
	var tempAttributeStructSliceReference []*AttributeStruct
	tempAttributeStructSliceReference = make([]*AttributeStruct, 0)
	var tempAttributesList *AttributeStructSliceReferenceType
	tempAttributesList = (*AttributeStructSliceReferenceType)(&tempAttributeStructSliceReference)
	tempTestCaseModel.AttributesList = tempAttributesList

	// Generate 'TestCaseModelMap'
	tempTestCaseModel.TestCaseModelMap = make(map[string]MatureTestCaseModelElementStruct)
	for _, tempMatureTestCaseModelElementMessage := range detailedTestCaseResponse.DetailedTestCase.
		TestCaseBasicInformation.TestCaseModel.TestCaseModelElements {

		var tempMatureTestCaseModelElement MatureTestCaseModelElementStruct
		tempMatureTestCaseModelElement = MatureTestCaseModelElementStruct{
			MatureTestCaseModelElementMessage: fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
				OriginalElementUuid:      tempMatureTestCaseModelElementMessage.OriginalElementUuid,
				OriginalElementName:      tempMatureTestCaseModelElementMessage.OriginalElementName,
				MatureElementUuid:        tempMatureTestCaseModelElementMessage.MatureElementUuid,
				PreviousElementUuid:      tempMatureTestCaseModelElementMessage.PreviousElementUuid,
				NextElementUuid:          tempMatureTestCaseModelElementMessage.NextElementUuid,
				FirstChildElementUuid:    tempMatureTestCaseModelElementMessage.FirstChildElementUuid,
				ParentElementUuid:        tempMatureTestCaseModelElementMessage.ParentElementUuid,
				TestCaseModelElementType: tempMatureTestCaseModelElementMessage.TestCaseModelElementType,
			},
			MatureTestCaseModelElementMetaData: MatureTestCaseModelElementMetaDataStruct{
				ChosenDropZoneUuid:        "",
				ChosenDropZoneColorString: "",
			},
		}
		tempTestCaseModel.TestCaseModelMap[tempMatureTestCaseModelElementMessage.MatureElementUuid] = tempMatureTestCaseModelElement

	}

	// Generate 'MatureTestInstructionMap'
	tempTestCaseModel.MatureTestInstructionMap = make(map[string]MatureTestInstructionStruct)
	for _, tempMatureTestInstructionMessage := range detailedTestCaseResponse.DetailedTestCase.MatureTestInstructions.
		MatureTestInstructions {

		// Add TestInstruction
		var tempMatureTestInstruction MatureTestInstructionStruct
		tempMatureTestInstruction = MatureTestInstructionStruct{
			BasicTestInstructionInformation_NonEditableInformation: tempMatureTestInstructionMessage.
				BasicTestInstructionInformation.NonEditableInformation,
			BasicTestInstructionInformation_EditableInformation: tempMatureTestInstructionMessage.
				BasicTestInstructionInformation.EditableInformation,
			BasicTestInstructionInformation_InvisibleBasicInformation: tempMatureTestInstructionMessage.
				BasicTestInstructionInformation.InvisibleBasicInformation,
			MatureBasicTestInstructionInformation: tempMatureTestInstructionMessage.
				MatureTestInstructionInformation.MatureBasicTestInstructionInformation,
			CreatedAndUpdatedInformation: tempMatureTestInstructionMessage.
				MatureTestInstructionInformation.CreatedAndUpdatedInformation,
			TestInstructionAttributesList: make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
				MatureTestInstructionInformationMessage_TestInstructionAttributeMessage),
		}

		// Add Attributes for TestInstruction
		for _, tempTestInstructionAttributeMessage := range tempMatureTestInstructionMessage.
			MatureTestInstructionInformation.TestInstructionAttributesList {
			var tempTestInstructionAttributes fenixGuiTestCaseBuilderServerGrpcApi.
				MatureTestInstructionInformationMessage_TestInstructionAttributeMessage
			tempTestInstructionAttributes = fenixGuiTestCaseBuilderServerGrpcApi.
				MatureTestInstructionInformationMessage_TestInstructionAttributeMessage{
				BaseAttributeInformation: tempTestInstructionAttributeMessage.BaseAttributeInformation,
				AttributeInformation:     tempTestInstructionAttributeMessage.AttributeInformation,
			}

			tempMatureTestInstruction.TestInstructionAttributesList[tempTestInstructionAttributeMessage.
				BaseAttributeInformation.TestInstructionAttributeUuid] = &tempTestInstructionAttributes
		}

		tempTestCaseModel.MatureTestInstructionMap[tempMatureTestInstructionMessage.MatureTestInstructionInformation.
			MatureBasicTestInstructionInformation.TestInstructionMatureUuid] = tempMatureTestInstruction

	}

	// Generate 'MatureTestInstructionContainerMap'
	tempTestCaseModel.MatureTestInstructionContainerMap = make(map[string]MatureTestInstructionContainerStruct)
	for _, tempMatureTestInstructionContainerMessage := range detailedTestCaseResponse.DetailedTestCase.
		MatureTestInstructionContainers.MatureTestInstructionContainers {

		// Add TestInstructionContainer
		var tempMatureTestInstructionContainer MatureTestInstructionContainerStruct
		tempMatureTestInstructionContainer = MatureTestInstructionContainerStruct{
			NonEditableInformation: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.NonEditableInformation,
			EditableInformation: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.EditableInformation,
			InvisibleBasicInformation: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.InvisibleBasicInformation,
			EditableTestInstructionContainerAttributes: tempMatureTestInstructionContainerMessage.
				BasicTestInstructionContainerInformation.EditableTestInstructionContainerAttributes,
			MatureTestInstructionContainerInformation: tempMatureTestInstructionContainerMessage.
				MatureTestInstructionContainerInformation.MatureTestInstructionContainerInformation,
			CreatedAndUpdatedInformation: tempMatureTestInstructionContainerMessage.
				MatureTestInstructionContainerInformation.CreatedAndUpdatedInformation,
		}

		tempTestCaseModel.MatureTestInstructionContainerMap[tempMatureTestInstructionContainerMessage.
			MatureTestInstructionContainerInformation.MatureTestInstructionContainerInformation.
			TestInstructionContainerMatureUuid] = tempMatureTestInstructionContainer
	}

	// Generate TextualTestCaseRepresentationSimpleStack
	for _, tempTextualTestCaseRepresentationSimpleInstance := range detailedTestCaseResponse.DetailedTestCase.
		TestCaseExtraInformation.TestCaseTextualRepresentationHistory.TextualTestCaseRepresentationSimpleHistory {

		tempTestCaseModel.TextualTestCaseRepresentationSimpleStack =
			append(tempTestCaseModel.TextualTestCaseRepresentationSimpleStack, tempTextualTestCaseRepresentationSimpleInstance)
	}

	//Generate TextualTestCaseRepresentationComplexStack
	for _, tempTextualTestCaseRepresentationComplexInstance := range detailedTestCaseResponse.DetailedTestCase.
		TestCaseExtraInformation.TestCaseTextualRepresentationHistory.TextualTestCaseRepresentationComplexHistory {

		tempTestCaseModel.TextualTestCaseRepresentationComplexStack =
			append(tempTestCaseModel.TextualTestCaseRepresentationComplexStack, tempTextualTestCaseRepresentationComplexInstance)
	}

	// Generate TextualTestCaseRepresentationExtendedStack
	for _, tempTextualTestCaseRepresentationExtendedInstance := range detailedTestCaseResponse.DetailedTestCase.
		TestCaseExtraInformation.TestCaseTextualRepresentationHistory.TextualTestCaseRepresentationExtendedStackHistory {

		tempTestCaseModel.TextualTestCaseRepresentationExtendedStack =
			append(tempTestCaseModel.TextualTestCaseRepresentationExtendedStack, tempTextualTestCaseRepresentationExtendedInstance)
	}

	// Generate ImportedTemplateFilesFromGitHub
	for _, tempTestCaseTemplateFile := range detailedTestCaseResponse.GetDetailedTestCase().TestCaseTemplateFiles.GetTestCaseTemplateFile() {

		var githubFile importFilesFromGitHub.GitHubFile
		githubFile = importFilesFromGitHub.GitHubFile{
			Name:                tempTestCaseTemplateFile.Name,
			Type:                "",
			URL:                 tempTestCaseTemplateFile.URL,
			DownloadURL:         tempTestCaseTemplateFile.DownloadURL,
			Content:             nil,
			SHA:                 tempTestCaseTemplateFile.SHA,
			Size:                int(tempTestCaseTemplateFile.Size),
			FileContentAsString: tempTestCaseTemplateFile.FileContentAsString,
			FileHash:            tempTestCaseTemplateFile.FileHash,
		}

		tempTestCaseModel.ImportedTemplateFilesFromGitHub = append(
			tempTestCaseModel.ImportedTemplateFilesFromGitHub, githubFile)
	}

	// Generate users TestData for TestCase
	var usersChosenTestDataForTestCaseMessage *fenixGuiTestCaseBuilderServerGrpcApi.UsersChosenTestDataForTestCaseMessage
	usersChosenTestDataForTestCaseMessage = detailedTestCaseResponse.GetDetailedTestCase().GetTestCaseTestData()

	var chosenTestDataPointsPerGroupMap map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType
	chosenTestDataPointsPerGroupMap = make(map[testDataEngine.TestDataPointGroupNameType]*testDataEngine.TestDataPointNameMapType)

	var testDataPointGroups []testDataEngine.TestDataPointGroupNameType

	var testData *testDataEngine.TestDataForGroupObjectStruct
	testData = &testDataEngine.TestDataForGroupObjectStruct{
		TestDataPointGroups:             nil,
		TestDataPointsForAGroup:         nil,
		ChosenTestDataPointsPerGroupMap: chosenTestDataPointsPerGroupMap,
		ShouldUpdateMainWindow: testDataEngine.ResponseChannelStruct{
			ShouldBeUpdated:        false,
			TestDataPointGroupName: "",
		},
	}

	if usersChosenTestDataForTestCaseMessage != nil {
		// User has TestData stored for the TestCase

		// Loop all Groups with TestDataPoints in gRPC-message
		for testDataGroupNameGrpc, testDataGroupGrpc := range usersChosenTestDataForTestCaseMessage.ChosenTestDataPointsPerGroupMap {

			var testDataPointNameMap map[testDataEngine.TestDataValueNameType]*[]*testDataEngine.DataPointTypeForGroupsStruct
			testDataPointNameMap = make(map[testDataEngine.TestDataValueNameType]*[]*testDataEngine.DataPointTypeForGroupsStruct)

			var testDataPointNameMapAsObject testDataEngine.TestDataPointNameMapType

			for testDataPointNameGrpc, testDataPointGrpc := range testDataGroupGrpc.ChosenTestDataRowsPerTestDataPointMap {

				var dataPointTypeForGroups []*testDataEngine.DataPointTypeForGroupsStruct

				for _, testDataRowGrpc := range testDataPointGrpc.TestDataRows {

					// Initiate the maps in the struct, only 'selectedTestDataPointUuidMap' is filled with values
					var searchResultDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct
					searchResultDataPointUuidMap = make(map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)
					var availableTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct
					availableTestDataPointUuidMap = make(map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)
					var selectedTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct
					selectedTestDataPointUuidMap = make(map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)

					// Create the RowSummary for each row for 'dataPointTypeForGroup'
					for _, testDataPointRowValueSummaryGrpc := range testDataRowGrpc.TestDataPointRowValueSummaryMap {

						var testDataPointRowValuesSummary testDataEngine.TestDataPointRowUuidStruct
						testDataPointRowValuesSummary = testDataEngine.TestDataPointRowUuidStruct{
							TestDataPointRowUuid:          testDataEngine.TestDataPointRowUuidType(testDataPointRowValueSummaryGrpc.GetTestDataPointRowUuid()),
							TestDataPointRowValuesSummary: testDataEngine.TestDataPointRowValuesSummaryType(testDataPointRowValueSummaryGrpc.GetTestDataPointRowValuesSummary()),
						}

						// Add the RowUUid and the values summary to the map
						selectedTestDataPointUuidMap[testDataEngine.TestDataPointRowUuidType(testDataPointRowValueSummaryGrpc.GetTestDataPointRowUuid())] = testDataPointRowValuesSummary

					}

					// Create the 'dataPointTypeForGroup'
					var dataPointTypeForGroup *testDataEngine.DataPointTypeForGroupsStruct
					dataPointTypeForGroup = &testDataEngine.DataPointTypeForGroupsStruct{
						TestDataDomainUuid:            testDataEngine.TestDataDomainUuidType(testDataRowGrpc.GetTestDataDomainUuid()),
						TestDataDomainName:            testDataEngine.TestDataDomainNameType(testDataRowGrpc.GetTestDataDomainName()),
						TestDataAreaUuid:              testDataEngine.TestDataAreaUuidType(testDataRowGrpc.GetTestDataAreaUuid()),
						TestDataAreaName:              testDataEngine.TestDataAreaNameType(testDataRowGrpc.GetTestDataAreaName()),
						TestDataPointName:             testDataEngine.TestDataValueNameType(testDataRowGrpc.GetTestDataPointName()),
						SearchResultDataPointUuidMap:  searchResultDataPointUuidMap,
						AvailableTestDataPointUuidMap: availableTestDataPointUuidMap,
						SelectedTestDataPointUuidMap:  selectedTestDataPointUuidMap,
					}

					// Add 'dataPointTypeForGroup' to slice of 'dataPointTypeForGroups'
					dataPointTypeForGroups = append(dataPointTypeForGroups, dataPointTypeForGroup)

				}

				// Add the slice of TestDataRows to the map for TestDataPoints
				testDataPointNameMap[testDataEngine.TestDataValueNameType(testDataPointNameGrpc)] = &dataPointTypeForGroups

			}

			// Add to slice of GroupNames
			testDataPointGroups = append(testDataPointGroups, testDataEngine.TestDataPointGroupNameType(testDataGroupNameGrpc))

			// Move Map into object
			testDataPointNameMapAsObject = testDataPointNameMap

			chosenTestDataPointsPerGroupMap[testDataEngine.TestDataPointGroupNameType(testDataGroupNameGrpc)] = &testDataPointNameMapAsObject

		}

		testData.ChosenTestDataPointsPerGroupMap = chosenTestDataPointsPerGroupMap

		testData.TestDataPointGroups = testDataPointGroups

		tempTestCaseModel.TestData = testData

	} else {
		// User has no TestData stored for the TestCase
		testData.ChosenTestDataPointsPerGroupMap = chosenTestDataPointsPerGroupMap

		tempTestCaseModel.TestData = testData
	}

	// Update The Hash for the TestCase
	tempTestCaseModel.TestCaseHashWhenTestCaseWasSavedOrLoaded = detailedTestCaseResponse.DetailedTestCase.MessageHash
	tempTestCaseModel.TestDataHashWhenTestCaseWasSavedOrLoaded = detailedTestCaseResponse.GetDetailedTestCase().GetTestCaseTestData().GetHashOfThisMessageWithEmptyHashField()

	// Generate the TestCaseMetaData for the TestCase-model from the gRPC-data
	var tempMetaDataGroupsMap map[string]*MetaDataGroupStruct
	var tempMetaDataGroupsOrder []string
	tempMetaDataGroupsMap = make(map[string]*MetaDataGroupStruct)

	// Loop MetaDataGroups in gPRC-data
	for tempMetaDataGroupNameFromGrpc, tempMetaDataGroupFromGrpc := range detailedTestCaseResponse.GetDetailedTestCase().GetTestCaseMetaData().GetMetaDataGroupsMap() {

		tempMetaDataGroupsOrder = append(tempMetaDataGroupsOrder, tempMetaDataGroupNameFromGrpc)

		var tempMetaDataInGroupOrder []string

		var tempMetaDataInGroupMap map[string]*MetaDataInGroupStruct
		tempMetaDataInGroupMap = make(map[string]*MetaDataInGroupStruct)

		// Loop MetaDataGroupItems in MetaDataGroup
		for tempMetaDataGroupItemNameFromGrpc, tempMetaDataGroupItemFromGrpc := range tempMetaDataGroupFromGrpc.GetMetaDataInGroupMap() {

			tempMetaDataInGroupOrder = append(tempMetaDataInGroupOrder, tempMetaDataGroupItemNameFromGrpc)

			var tempMetaDataInGroup MetaDataInGroupStruct
			tempMetaDataInGroup = MetaDataInGroupStruct{
				MetaDataGroupName:                          tempMetaDataGroupNameFromGrpc,
				MetaDataName:                               tempMetaDataGroupItemNameFromGrpc,
				SelectType:                                 MetaDataSelectType(tempMetaDataGroupItemFromGrpc.SelectType),
				Mandatory:                                  tempMetaDataGroupItemFromGrpc.IsMandatory,
				AvailableMetaDataValues:                    tempMetaDataGroupItemFromGrpc.AvailableMetaDataValues,
				SelectedMetaDataValueForSingleSelect:       tempMetaDataGroupItemFromGrpc.SelectedMetaDataValueForSingleSelect,
				SelectedMetaDataValuesForMultiSelect:       tempMetaDataGroupItemFromGrpc.SelectedMetaDataValuesForMultiSelect,
				SelectedMetaDataValuesForMultiSelectMapPtr: nil,
			}

			// Generate the map holding the Selected values for the multi select
			var tempSelectedMetaDataValuesMap map[string]string
			tempSelectedMetaDataValuesMap = make(map[string]string)

			for tempMetaDataItemKeyFromGrpc, tempMetaDataItemValueFromGrpc := range tempMetaDataGroupItemFromGrpc.
				SelectedMetaDataValuesForMultiSelectMap {
				tempSelectedMetaDataValuesMap[tempMetaDataItemKeyFromGrpc] = tempMetaDataItemValueFromGrpc
			}

			// Store selected values map in main message
			tempMetaDataInGroup.SelectedMetaDataValuesForMultiSelectMapPtr = &tempSelectedMetaDataValuesMap

			// Add MetaDataGroupItem to 'MetaDataInGroupMap'
			tempMetaDataInGroupMap[tempMetaDataGroupItemNameFromGrpc] = &tempMetaDataInGroup

		}

		// Create the var MetaDataGroup-object
		var tempMetaDataGroup MetaDataGroupStruct
		tempMetaDataGroup = MetaDataGroupStruct{
			MetaDataGroupName:     tempMetaDataGroupNameFromGrpc,
			MetaDataInGroupOrder:  tempMetaDataInGroupOrder,
			MetaDataInGroupMapPtr: &tempMetaDataInGroupMap,
		}

		// Add  MetaDataGroup-object to 'tempMetaDataGroupsMap'
		tempMetaDataGroupsMap[tempMetaDataGroupNameFromGrpc] = &tempMetaDataGroup

	}

	// Create the full 'TestCaseMetaData-object'
	var tempTestCaseMetaData TestCaseMetaDataStruct
	tempTestCaseMetaData = TestCaseMetaDataStruct{
		CurrentSelectedDomainUuid: detailedTestCaseResponse.GetDetailedTestCase().
			GetTestCaseMetaData().GetCurrentSelectedDomainUuid(),
		TestCaseMetaDataMessageStructForTestCaseWhenLastSaved: nil,
		MetaDataGroupsOrder:  tempMetaDataGroupsOrder,
		MetaDataGroupsMapPtr: &tempMetaDataGroupsMap,
	}

	// Add converted TestCaseMetaData to 'TestCaseModel'
	tempTestCaseModel.TestCaseMetaDataPtr = &tempTestCaseMetaData

	// Add TestCase to map with all TestCasesMapPtr
	if testCaseModel.TestCasesMapPtr == nil {
		var tempTestCasesMapPtr map[string]*TestCaseModelStruct
		tempTestCasesMapPtr = make(map[string]*TestCaseModelStruct)
		testCaseModel.TestCasesMapPtr = &tempTestCasesMapPtr
	}

	// Create temporary instance to be used for verifying of Hash
	var tempTestCaseUuid string
	tempTestCaseUuid = "temp_" + testCaseUuid

	// Get TestCasesMap
	var testCasesMap map[string]*TestCaseModelStruct
	testCasesMap = *testCaseModel.TestCasesMapPtr

	testCasesMap[tempTestCaseUuid] = &tempTestCaseModel

	// Verify that calculated Hash is the same as the Stored Hash from the Database
	var generatedHash string
	_, _, _, _, _, _, _, _, generatedHash, err = testCaseModel.generateTestCaseForGrpcAndHash(tempTestCaseUuid, false)
	if err != nil {

		// Remove temporary stored TestCase
		delete(testCasesMap, tempTestCaseUuid)

		return err
	}

	// Check hash towards Hash from the Database
	if generatedHash != detailedTestCaseResponse.DetailedTestCase.MessageHash {

		errorId := "ab73a9b8-386c-4ee4-8c6b-594850acdebf"
		err = errors.New(fmt.Sprintf("after loading testcase '%s', from database, the Hash that is recreated ('%s') is not the"+
			" same as the one stored in database ('%s') [ErrorID: %s]",
			testCaseUuid, generatedHash, detailedTestCaseResponse.DetailedTestCase.MessageHash, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		// Remove temporary stored TestCase
		delete(testCasesMap, tempTestCaseUuid)

		return err

	}

	// Add TestCase to map with TestCasesMapPtr
	testCasesMap[testCaseUuid] = &tempTestCaseModel

	// Remove temporary stored TestCase
	delete(testCasesMap, tempTestCaseUuid)

	return err

}
