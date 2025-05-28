package testCaseModel

import (
	"errors"
	"fmt"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// VerifyTestCaseHashTowardsDatabase - Verify if the Hash for the TestCase is the same as the one in the database
func (testCaseModel *TestCasesModelsStruct) VerifyTestCaseHashTowardsDatabase(testCaseUuid string) (
	hashIsTheSame bool, err error) {

	var existsInMap bool

	// Get current TestCase
	_, existsInMap = testCaseModel.TestCasesMap[testCaseUuid]
	if existsInMap == false {

		errorId := "13f7c602-b8b7-427d-92b0-335556c071f1"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMap [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// Create TestCase-hash
	var testcaseHash string
	_, _, _, _, _, _, _, _, testcaseHash, err = testCaseModel.generateTestCaseForGrpcAndHash(testCaseUuid, false)
	if err != nil {
		return false, err
	}

	// Get Hash from Database via gRPC
	var testCaseHashRespons *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashResponse
	var testCaseUuidList []string
	testCaseUuidList = []string{testCaseUuid}

	// Some error when retrieving from Database
	testCaseHashRespons = testCaseModel.GrpcOutReference.LoadHashesForTestCases(testCaseModel.CurrentUser, testCaseUuidList)
	if testCaseHashRespons.AckNack.AckNack == false {
		errorId := "eadc89a7-eb1d-4c96-b89d-5a2f98996a2a"
		err = errors.New(fmt.Sprintf("Couldn't get Hash stored in Database for testcase '%s'. Message returned: '%s' [ErrorID: %s]", testCaseUuid, testCaseHashRespons.AckNack.Comments, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// More than one Hash was returned
	if len(testCaseHashRespons.TestCasesHashes) > 1 {
		errorId := "63aca654-1a61-43d1-ab1f-4d375633dab5"
		err = errors.New(fmt.Sprintf("More then one Hash was returned from Database for testcase '%s' [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// Check if current Hash is the same as the one stored in Database
	if testcaseHash == testCaseHashRespons.TestCasesHashes[0].GetTestCaseHash() {
		hashIsTheSame = true
	} else {
		hashIsTheSame = false
	}

	return hashIsTheSame, err
}

// VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase - Verify if the latest Loaded or Saved Hash for the TestCase is the same as the one in the database
func (testCaseModel *TestCasesModelsStruct) VerifyLatestLoadedOrSavedTestCaseHashTowardsDatabase(testCaseUuid string) (
	hashIsTheSame bool, err error) {

	var existsInMap bool
	var tempTestCasePtr *TestCaseModelStruct

	// Get current TestCase
	tempTestCasePtr, existsInMap = testCaseModel.TestCasesMap[testCaseUuid]
	if existsInMap == false {

		errorId := "959d258d-2f83-46e5-8aba-8655b8fb27b2"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMap [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// Create TestCase-hash
	var testcaseHash string
	testcaseHash = tempTestCasePtr.TestCaseHashWhenTestCaseWasSavedOrLoaded

	// Get Hash from Database via gRPC
	var testCaseHashRespons *fenixGuiTestCaseBuilderServerGrpcApi.TestCasesHashResponse
	var testCaseUuidList []string
	testCaseUuidList = []string{testCaseUuid}

	// Some error when retrieving from Database
	testCaseHashRespons = testCaseModel.GrpcOutReference.LoadHashesForTestCases(testCaseModel.CurrentUser, testCaseUuidList)
	if testCaseHashRespons.AckNack.AckNack == false {
		errorId := "eadc89a7-eb1d-4c96-b89d-5a2f98996a2a"
		err = errors.New(fmt.Sprintf("Couldn't get Hash stored in Database for testcase '%s'. Message returned: '%s' [ErrorID: %s]", testCaseUuid, testCaseHashRespons.AckNack.Comments, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// More than one Hash was returned
	if len(testCaseHashRespons.TestCasesHashes) > 1 {
		errorId := "63aca654-1a61-43d1-ab1f-4d375633dab5"
		err = errors.New(fmt.Sprintf("More then one Hash was returned from Database for testcase '%s' [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// Check if current Hash is the same as the one stored in Database
	if testcaseHash == testCaseHashRespons.TestCasesHashes[0].GetTestCaseHash() {
		hashIsTheSame = true
	} else {
		hashIsTheSame = false
	}

	return hashIsTheSame, err
}

// TestCaseHashIsChangedSinceLoadedOrSaved - Verify if the Hash for the TestCase is the same as the one when TestCasesMap was last Loaded or Saved
func (testCaseModel *TestCasesModelsStruct) TestCaseHashIsChangedSinceLoadedOrSaved(testCaseUuid string) (
	hashIsChanged bool, err error) {

	var existsInMap bool
	var tempTestCasePtr *TestCaseModelStruct

	// Get current TestCase
	tempTestCasePtr, existsInMap = testCaseModel.TestCasesMap[testCaseUuid]
	if existsInMap == false {

		errorId := "d9b6aa9e-0cc4-4424-8d74-c794b44bbcd6"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCasesMap [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) // TODO Send on Error-channel

		return false, err
	}

	// Create TestCase-hash
	var testcaseHash string
	_, _, _, _, _, _, _, _, testcaseHash, err = testCaseModel.generateTestCaseForGrpcAndHash(testCaseUuid, false)
	if err != nil {
		return false, err
	}

	// Is Hash chaned or not
	if testcaseHash != tempTestCasePtr.TestCaseHashWhenTestCaseWasSavedOrLoaded {
		hashIsChanged = true
	} else {
		hashIsChanged = false
	}

	return hashIsChanged, err
}
