package listTestCasesModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

// LoadTestCaseThatCanBeEditedByUser
// Load list with TestCasesMapPtr that the user can edit
func LoadTestCaseThatCanBeEditedByUser(
	testCasesModeReference *testCaseModel.TestCasesModelsStruct,
	testCaseUpdatedMinTimeStamp time.Time,
	testCaseExecutionUpdatedMinTimeStamp time.Time) {

	var listTestCasesThatCanBeEditedResponseMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage
	listTestCasesThatCanBeEditedResponseMessage = testCasesModeReference.GrpcOutReference.
		ListTestCasesThatCanBeEditedResponseMessage(testCaseUpdatedMinTimeStamp, testCaseExecutionUpdatedMinTimeStamp)

	if listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().AckNack == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "ebfe8adb-c224-4071-869b-f79cefde0dd3",
			"error": listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().Comments,
		}).Warning("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListTestCasesThatCanBeEditedResponseMessage'")

		return
	}

	// Store the slice with TestCasesMapPtr that a user can edit as a Map
	storeTestCaseThatCanBeEditedByUser(
		listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser(),
		testCasesModeReference)

	// Store the slice with TestCasesMapPtr
	//testCasesModeReference.TestCasesThatCanBeEditedByUserSlice = listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser()
	/*
		testCasesModeReference.TestCasesThatCanBeEditedByUserSlice = nil
		for _, tempTestCasesThatCanBeEditedByUser := range testCasesModeReference.TestCasesThatCanBeEditedByUserMap {
			testCasesModeReference.TestCasesThatCanBeEditedByUserSlice = append(
				testCasesModeReference.TestCasesThatCanBeEditedByUserSlice, tempTestCasesThatCanBeEditedByUser)
		}
	*/

}

// Store TestCasesMapPtr That Can Be Edited By User
func storeTestCaseThatCanBeEditedByUser(
	testCasesThatCanBeEditedByUserAsSlice []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage,
	testCasesModeReference *testCaseModel.TestCasesModelsStruct) {

	var err error
	var domainBitSetExistiInMap bool
	var groupBitSetExistInMap bool
	var itemBitSetExistInMap bool
	var valueBitSetExistInMap bool
	var latestTestCaseUpdatedMinTimeStamp time.Time
	var latestTestCaseExecutionUpdatedMinTimeStamp time.Time

	// Store the TestCaseThatCanBeEditedByUser-list in the TestCaseModel

	// Initiate Model if it hasn't been initialized
	if TestCasesThatCanBeEditedByUserMap == nil {
		TestCasesThatCanBeEditedByUserMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
			TestCaseThatCanBeEditedByUserMessage)
	}

	// Store the Available TemplateRepositoryApiUrls as a map structure in TestCase-struct
	for _, testCaseThatCanBeEditedByUser := range testCasesThatCanBeEditedByUserAsSlice {

		// Check for later 'latestTestCaseUpdatedMinTimeStamp'
		if latestTestCaseUpdatedMinTimeStamp.Before(testCaseThatCanBeEditedByUser.GetLastSavedTimeStamp().AsTime()) {
			latestTestCaseUpdatedMinTimeStamp = testCaseThatCanBeEditedByUser.GetLastSavedTimeStamp().AsTime()
		}

		// Check for later 'latestTestCaseExecutionUpdatedMinTimeStamp'
		if latestTestCaseExecutionUpdatedMinTimeStamp.Before(testCaseThatCanBeEditedByUser.GetLatestTestCaseExecutionStatusInsertTimeStamp().AsTime()) {
			latestTestCaseExecutionUpdatedMinTimeStamp = testCaseThatCanBeEditedByUser.GetLatestTestCaseExecutionStatusInsertTimeStamp().AsTime()
		}

		TestCasesThatCanBeEditedByUserMap[testCaseThatCanBeEditedByUser.GetTestCaseUuid()] =
			testCaseThatCanBeEditedByUser

		// Get PreViewMessage for TestCase
		var tempTestCasePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage
		tempTestCasePreview = testCaseThatCanBeEditedByUser.GetTestCasePreview().GetTestCasePreview()

		var selectedMetaDataValuesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
			TestCasePreviewStructureMessage_SelectedMetaDataValueMessage

		selectedMetaDataValuesMap = tempTestCasePreview.GetSelectedMetaDataValuesMap()

		// Generate the resultsEntry
		var resultsEntry *boolbits.Entry
		resultsEntry, err = boolbits.NewAllZerosEntry(64)
		if err != nil {

			errorId := "3739bf55-c0ec-426f-a302-758b2f72f3a1"
			errorMessage := fmt.Sprintf("couldn't produce new 'NewAllZerosEntry', {error: %s} [ErrorID: %s]",
				err.Error(), errorId)

			log.Fatalln(errorMessage)

		}

		// Loop all Filter settings and produce SimpleMetaDataFilterEntry
		for _, selectedMetaDataValue := range selectedMetaDataValuesMap {

			// Generate BitSet for Domain, Group, Item and Value
			var domainBitSet *boolbits.BitSet
			var metaDataGroupBitSet *boolbits.BitSet
			var metaDataItemBitSet *boolbits.BitSet
			var metaDataValueBitSet *boolbits.BitSet

			// Get BitSets
			domainBitSet, domainBitSetExistiInMap = testCasesModeReference.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				DomainsBitSetMap[selectedMetaDataValue.OwnerDomainUuid]
			metaDataGroupBitSet, groupBitSetExistInMap = testCasesModeReference.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				MetaDataGroupsBitSetMap[selectedMetaDataValue.MetaDataGroupName]
			metaDataItemBitSet, itemBitSetExistInMap = testCasesModeReference.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				MetaDataGroupItemsBitSetMap[selectedMetaDataValue.MetaDataName]
			metaDataValueBitSet, valueBitSetExistInMap = testCasesModeReference.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				MetaDataGroupItemValuesBitSetMap[selectedMetaDataValue.MetaDataNameValue]

			// Only produce the Entry if the all 4 of the BitSets still are valid
			if domainBitSetExistiInMap && groupBitSetExistInMap && itemBitSetExistInMap && valueBitSetExistInMap {

				var metaDataVaueEntry *boolbits.Entry
				metaDataVaueEntry, err = boolbits.NewEntry(domainBitSet, metaDataGroupBitSet, metaDataItemBitSet, metaDataValueBitSet)
				if err != nil {

					errorId := "42c1e26f-6345-459e-9ede-15ca118b47e7"
					errorMessage := fmt.Sprintf("couldn't produce new 'NewEntry', {error: %s} [ErrorID: %s]",
						err.Error(), errorId)

					log.Fatalln(errorMessage)

				}

				resultsEntry, err = resultsEntry.Or(metaDataVaueEntry)
				if err != nil {

					errorId := "061acd28-e7c0-4438-9c80-5382a9a99ac2"
					errorMessage := fmt.Sprintf("couldn't do 'Entry-OR', {error: %s} [ErrorID: %s]",
						err.Error(), errorId)

					log.Fatalln(errorMessage)

				}

			}

		}

		// Store filter results Entry for the TestCase
		if SimpleTestCaseMetaDataFilterEntryMap == nil {
			SimpleTestCaseMetaDataFilterEntryMap = make(map[string]*boolbits.Entry)
		}

		SimpleTestCaseMetaDataFilterEntryMap[testCaseThatCanBeEditedByUser.GetTestCaseUuid()] = resultsEntry

	}

	// Store 'latestTestCaseUpdatedMinTimeStamp' & 'latestTestCaseExecutionUpdatedMinTimeStamp' to be used with next Database-call
	LatestTestCaseUpdatedMinTimeStampForDatabaseCall = latestTestCaseUpdatedMinTimeStamp
	LatestTestCaseExecutionUpdatedMinTimeStampForDatabaseCall = latestTestCaseExecutionUpdatedMinTimeStamp

}
