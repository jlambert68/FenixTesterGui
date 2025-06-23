package testSuitesModel

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"log"
)

// Copy Fields in Model to fields that is used by the UI
func (testSuiteModel *TestSuiteModelStruct) copyModelToUiFields() {

	// Copy 'TestSuiteDeletionDate'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate = testSuiteModel.testSuiteDeletionDate

	// Copy 'TestSuiteName'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteName = testSuiteModel.testSuiteName

	// Copy 'TestSuiteDescription'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteDescription = testSuiteModel.testSuiteDescription

	// Copy 'TestSuiteOwnerDomainUuid'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid = testSuiteModel.testSuiteOwnerDomainUuid

	// Copy 'TestSuiteExecutionEnvironment'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment = testSuiteModel.testSuiteExecutionEnvironment

	// Copy changes for 'TestSuiteMetaDataHash'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataHash = testSuiteModel.testSuiteMetaDataHash

	// Copy changes for 'TestSuiteMetaDataPtr'
	var copyTestSuiteMetaData TestSuiteMetaDataStruct
	var originalTestSuiteMetaData TestSuiteMetaDataStruct

	originalTestSuiteMetaData = *testSuiteModel.testSuiteMetaDataPtr

	err := copier.CopyWithOption(&copyTestSuiteMetaData, &originalTestSuiteMetaData, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "3e4f465a-8a41-4e44-81f8-47f9a085e778"

		errorMsg := fmt.Sprintf("error copying TestSuiteMetaDataStruct using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataPtr = &copyTestSuiteMetaData

	// Copy changes for 'TestSuiteTesDataHash'
	testSuiteModel.TestSuiteUIModelBinding.TestSuiteTesDataHash = testSuiteModel.testSuiteTestDataHash

	// Copy changes for 'TestDataPtr'
	var copyTestSuiteTestData testDataEngine.TestDataForGroupObjectStruct
	var originalTestSuiteTestData testDataEngine.TestDataForGroupObjectStruct

	originalTestSuiteTestData = *testSuiteModel.testDataPtr

	err = copier.CopyWithOption(&copyTestSuiteTestData, &originalTestSuiteTestData, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "82a55f35-ccf3-4cd3-9b1c-6159d0a5e73d"

		errorMsg := fmt.Sprintf("error copying TestDataForGroupObjectStruct using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}
	testSuiteModel.TestSuiteUIModelBinding.TestDataPtr = &copyTestSuiteTestData

}
