package testSuitesModel

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jlambert68/FenixScriptEngine/testDataEngine"
	"log"
)

// Copy Fields in Model, that is used by the UI, to models internal fields
func (testSuiteModel *TestSuiteModelStruct) copyUiFieldsToModel() {

	// Copy 'TestSuiteDeletionDate'
	testSuiteModel.testSuiteDeletionDate = testSuiteModel.TestSuiteUIModelBinding.TestSuiteDeletionDate

	// Copy 'TestSuiteName'
	testSuiteModel.testSuiteName = testSuiteModel.TestSuiteUIModelBinding.TestSuiteName

	// Copy 'TestSuiteDescription'
	testSuiteModel.testSuiteDescription = testSuiteModel.TestSuiteUIModelBinding.TestSuiteDescription

	// Copy 'TestSuiteOwnerDomainUuid'
	testSuiteModel.testSuiteOwnerDomainUuid = testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainUuid

	// Copy 'TestSuiteOwnerDomainName'
	testSuiteModel.testSuiteOwnerDomainName = testSuiteModel.TestSuiteUIModelBinding.TestSuiteOwnerDomainName

	// Copy 'TestSuiteExecutionEnvironment'
	testSuiteModel.testSuiteExecutionEnvironment = testSuiteModel.TestSuiteUIModelBinding.TestSuiteExecutionEnvironment

	// Copy changes for 'TestSuiteMetaDataHash'
	testSuiteModel.testSuiteMetaDataHash = testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataHash

	// Copy changes for 'TestSuiteMetaDataPtr'
	var copyTestSuiteMetaData TestSuiteMetaDataStruct
	var originalTestSuiteMetaData TestSuiteMetaDataStruct

	originalTestSuiteMetaData = *testSuiteModel.TestSuiteUIModelBinding.TestSuiteMetaDataPtr

	err := copier.CopyWithOption(&copyTestSuiteMetaData, &originalTestSuiteMetaData, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "b47a02a5-c035-4f2b-9760-91804d7eea57"

		errorMsg := fmt.Sprintf("error copying TestSuiteMetaDataStruct using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}
	testSuiteModel.testSuiteMetaDataPtr = &copyTestSuiteMetaData

	// Copy changes for 'TestSuiteTesDataHash'
	testSuiteModel.testSuiteTestDataHash = testSuiteModel.TestSuiteUIModelBinding.TestSuiteTesDataHash

	// Copy changes for 'TestDataPtr'
	var copyTestSuiteTestData testDataEngine.TestDataForGroupObjectStruct
	var originalTestSuiteTestData testDataEngine.TestDataForGroupObjectStruct

	originalTestSuiteTestData = *testSuiteModel.TestSuiteUIModelBinding.TestDataPtr

	err = copier.CopyWithOption(&copyTestSuiteTestData, &originalTestSuiteTestData, copier.Option{DeepCopy: true})
	if err != nil {

		errorID := "577cc5ef-7f99-48f6-8f22-34c8630ee491"

		errorMsg := fmt.Sprintf("error copying TestDataForGroupObjectStruct using 'copier'. error = '%s' [ErrorID: %s]",
			err.Error(),
			errorID)

		log.Fatalln(errorMsg)
	}
	testSuiteModel.testDataPtr = &copyTestSuiteTestData

}
