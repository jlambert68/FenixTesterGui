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
	testCaseModeReference *testCaseModel.TestCasesModelsStruct,
	testCaseUpdatedMinTimeStamp time.Time,
	testCaseExecutionUpdatedMinTimeStamp time.Time,
	testCaseModel *testCaseModel.TestCaseModelStruct) {

	var listTestCasesThatCanBeEditedResponseMessage *fenixGuiTestCaseBuilderServerGrpcApi.ListTestCasesThatCanBeEditedResponseMessage
	listTestCasesThatCanBeEditedResponseMessage = testCaseModeReference.GrpcOutReference.
		ListTestCasesThatCanBeEditedResponseMessage(testCaseUpdatedMinTimeStamp, testCaseExecutionUpdatedMinTimeStamp)

	if listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().AckNack == false {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":    "ebfe8adb-c224-4071-869b-f79cefde0dd3",
			"error": listTestCasesThatCanBeEditedResponseMessage.GetAckNackResponse().Comments,
		}).Warning("Problem to do gRPC-call to FenixTestGuiBuilderServer for 'loadTestCaseThatCanBeEditedByUser'")

		return
	}

	// Store the slice with TestCasesMapPtr that a user can edit as a Map
	storeTestCaseThatCanBeEditedByUser(
		listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser(),
		testCaseModeReference,
		testCaseModel )

	// Store the slice with TestCasesMapPtr
	//testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = listTestCasesThatCanBeEditedResponseMessage.GetTestCasesThatCanBeEditedByUser()
	/*
		testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = nil
		for _, tempTestCasesThatCanBeEditedByUser := range testCaseModeReference.TestCasesThatCanBeEditedByUserMap {
			testCaseModeReference.TestCasesThatCanBeEditedByUserSlice = append(
				testCaseModeReference.TestCasesThatCanBeEditedByUserSlice, tempTestCasesThatCanBeEditedByUser)
		}
	*/

}

// Store TestCasesMapPtr That Can Be Edited By User
func storeTestCaseThatCanBeEditedByUser(
	testCasesThatCanBeEditedByUserAsSlice []*fenixGuiTestCaseBuilderServerGrpcApi.TestCaseThatCanBeEditedByUserMessage,
	testCaseModeReference *testCaseModel.TestCasesModelsStruct,
	testCaseModel *testCaseModel.TestCaseModelStruct) {

	var err error
	var domainBitSetExistiInMap bool
	var groupBitSetExistiInMap bool
	var itemBitSetExistiInMap bool
	var valueBitSetExistiInMap bool

	// Store the TestCaseThatCanBeEditedByUser-list in the TestCaseModel
	testCaseModeReference.TestCasesThatCanBeEditedByUserMap = make(map[string]*fenixGuiTestCaseBuilderServerGrpcApi.
		TestCaseThatCanBeEditedByUserMessage)

	// Store the Available TemplateRepositoryApiUrls as a map structure in TestCase-struct
	for _, testCaseThatCanBeEditedByUser := range testCasesThatCanBeEditedByUserAsSlice {

		testCaseModeReference.TestCasesThatCanBeEditedByUserMap[testCaseThatCanBeEditedByUser.GetTestCaseUuid()] =
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
			domainBitSet, domainBitSetExistiInMap = testCasesModel.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				DomainsBitSetMap[selectedMetaDataValue.OwnerDomainUuid]
			metaDataGroupBitSet, groupBitSetExistiInMap  = testCasesModel.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				MetaDataGroupsBitSetMap[selectedMetaDataValue.MetaDataGroupName]
			metaDataItemBitSet, itemBitSetExistiInMap  = testCasesModel.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				MetaDataGroupItemsBitSetMap[selectedMetaDataValue.MetaDataName]
			metaDataValueBitSet, valueBitSetExistiInMap  = testCasesModel.TestCaseMetaDataForDomains.UniqueMetaDataBitSets.
				MetaDataGroupItemValuesBitSetMap[selectedMetaDataValue.MetaDataNameValue]


			// Only produce the Entry if the all 4 of the BitSets still are valid
			if 	domainBitSetExistiInMap && groupBitSetExistiInMap && itemBitSetExistiInMap &&  valueBitSetExistiInMap {

				var metaDataVaueEntry *boolbits.Entry
				metaDataVaueEntry, err = boolbits.NewEntry(domainBitSet, metaDataGroupBitSet, metaDataItemBitSet, metaDataValueBitSet)
				if err != nil {

					errorId := "42c1e26f-6345-459e-9ede-15ca118b47e7"
					errorMessage := fmt.Sprintf("couldn't produce new 'NewEntry', {error: %s} [ErrorID: %s]",
						err.Error(), errorId)

					log.Fatalln(errorMessage)

				}

				resultsEntry, err  = resultsEntry.Or(metaDataVaueEntry)
				if err != nil {

					errorId := "061acd28-e7c0-4438-9c80-5382a9a99ac2"
					errorMessage := fmt.Sprintf("couldn't do 'Entry-OR', {error: %s} [ErrorID: %s]",
						err.Error(), errorId)

					log.Fatalln(errorMessage)

				}

			}

			// Store results Entry in TestCase
			testCaseModeReference..SimpleTestCaseMetaDataFilterEntry



		}

	}

}
