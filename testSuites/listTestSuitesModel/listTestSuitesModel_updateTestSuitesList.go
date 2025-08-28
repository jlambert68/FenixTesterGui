package listTestSuitesModel

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

// LoadTestSuiteThatCanBeEditedByUser
// Load list with TestSuitesMapPtr that the user can edit
func LoadtestSuiteThatCanBeEditedByUser(
	testCasesModeReference *testCaseModel.TestCasesModelsStruct,
	testSuiteUpdatedMinTimeStamp time.Time,
	testSuiteExecutionUpdatedMinTimeStamp time.Time) {

	var listTestSuitesThatCanBeEditedResponseMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestSuitesResponseMessage
	listTestSuitesThatCanBeEditedResponseMessage = testCasesModeReference.GrpcOutReference.
		ListTestSuitesThatCanBeEditedResponseMessage(testSuiteUpdatedMinTimeStamp, testSuiteExecutionUpdatedMinTimeStamp)

	if listTestSuitesThatCanBeEditedResponseMessage.GetAckNackResponse().AckNack == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "e9868378-45a8-4cb7-ae61-98be8cb50736",
			"error": listTestSuitesThatCanBeEditedResponseMessage.GetAckNackResponse().Comments,
		}).Warning("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'ListTestSuitesThatCanBeEditedResponseMessage'")

		return
	}

	// Store the slice with TestSuitesMapPtr that a user can edit as a Map
	storeTestSuiteThatCanBeEditedByUser(
		listTestSuitesThatCanBeEditedResponseMessage.GetBasicTestSuiteInformation(),
		testCasesModeReference)

	// Store the slice with TestSuitesMapPtr
	//TestSuitesModeReference.TestSuitesThatCanBeEditedByUserSlice = listTestSuitesThatCanBeEditedResponseMessage.GetTestSuitesThatCanBeEditedByUser()
	/*
		TestSuitesModeReference.TestSuitesThatCanBeEditedByUserSlice = nil
		for _, tempTestSuitesThatCanBeEditedByUser := range TestSuitesModeReference.TestSuitesThatCanBeEditedByUserMap {
			TestSuitesModeReference.TestSuitesThatCanBeEditedByUserSlice = append(
				TestSuitesModeReference.TestSuitesThatCanBeEditedByUserSlice, tempTestSuitesThatCanBeEditedByUser)
		}
	*/

}

// Store TestSuitesMapPtr That Can Be Edited By User
func storeTestSuiteThatCanBeEditedByUser(
	TestSuitesThatCanBeEditedByUserAsSlice []*fenixGuiTestCaseBuilderServerGrpcApi.BasicTestSuiteInformationMessage,
	testCasesModeReference *testCaseModel.TestCasesModelsStruct) {

	var err error
	var domainBitSetExistiInMap bool
	var groupBitSetExistInMap bool
	var itemBitSetExistInMap bool
	var valueBitSetExistInMap bool

	// Store the TestSuiteThatCanBeEditedByUser-list in the TestCaseModel
	TestSuitesThatCanBeEditedByUserMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		BasicTestSuiteInformationMessage)

	// Store the Available TemplateRepositoryApiUrls as a map structure in TestCase-struct
	for _, testSuiteThatCanBeEditedByUser := range TestSuitesThatCanBeEditedByUserAsSlice {

		TestSuitesThatCanBeEditedByUserMap[testSuiteThatCanBeEditedByUser.NonEditableInformation.GetTestSuiteUuid()] =
			testSuiteThatCanBeEditedByUser

		// Get PreViewMessage for TestCase
		var tempTestSuitePreview *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage
		tempTestSuitePreview = testSuiteThatCanBeEditedByUser.GetTestSuitePreview().GetTestSuitePreview()

		var selectedMetaDataValuesMap map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
			TestSuitePreviewStructureMessage_SelectedTestSuiteMetaDataValueMessage

		selectedMetaDataValuesMap = tempTestSuitePreview.GetSelectedTestSuiteMetaDataValuesMap()

		// Generate the resultsEntry
		var resultsEntry *boolbits.Entry
		resultsEntry, err = boolbits.NewAllZerosEntry(64)
		if err != nil {

			errorId := "008089bc-6b0a-448a-a48f-62528bd4309c"
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
		if SimpleTestSuiteMetaDataFilterEntryMap == nil {
			SimpleTestSuiteMetaDataFilterEntryMap = make(map[string]*boolbits.Entry)
		}

		SimpleTestSuiteMetaDataFilterEntryMap[testSuiteThatCanBeEditedByUser.NonEditableInformation.GetTestSuiteUuid()] = resultsEntry

	}

}
