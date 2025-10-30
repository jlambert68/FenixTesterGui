package testSuiteExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
)

var testSuiteExecutionsMapMutex = &sync.RWMutex{}

// readListOrDetailedMessagesType
// Type defining if messages should be read for TableList or for DetailedPreview
type readListOrDetailedMessagesTypeType uint8

const (
	readDetailedMessages readListOrDetailedMessagesTypeType = iota
	ReadListMessages
)

// InitiateTestSuiteExecutionsMap
// Add to the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) InitiateTestSuiteExecutionsMap() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "6a689fe7-0459-477e-a117-032bf48ae88c",
	}).Debug("Incoming - 'InitiateTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "fe601432-68fa-42ad-8390-391286d97893",
	}).Debug("Outgoing - 'InitiateTestSuiteExecutionsMap'")

	// Lock Map for Writing
	testSuiteExecutionsMapMutex.Lock()

	// Initiate map if it is not already done
	if testSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
		testSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap = make(map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.
			TestSuiteExecutionsListMessage)
	}

	//UnLock Map
	testSuiteExecutionsMapMutex.Unlock()
}

// ReadFromTestSuiteExecutionsMapForTableList
// Read from the TestSuiteExecutions-Map for the TableList for TestSuiteExecutions
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadFromTestSuiteExecutionsMapForTableList(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType) (
	testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "c31a6352-9d51-428b-8215-02792ce8cd71",
	}).Debug("Incoming - 'ReadFromTestSuiteExecutionsMapForTableList'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "98ac2743-7da3-4ee1-8c36-6daf6784119c",
	}).Debug("Outgoing - 'ReadFromTestSuiteExecutionsMapForTableList'")

	testSuiteExecutionsListMessage, _, existInMap = testSuiteExecutionsModel.ReadFromTestSuiteExecutionsMap(
		ReadListMessages,
		testSuiteExecutionsMapKey)

	return testSuiteExecutionsListMessage, existInMap

}

// ReadFromTestSuiteExecutionsMap
// Read from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadFromTestSuiteExecutionsMap(
	readListOrDetailedMessagesType readListOrDetailedMessagesTypeType,
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType) (
	testSuiteExecutionsListMessages *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	testSuiteExecutionResponseMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage,
	existInMap bool) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                             "01698693-b50e-4eeb-9057-326c70f52e5e",
		"readListOrDetailedMessagesType": readListOrDetailedMessagesType,
	}).Debug("Incoming - 'ReadFromTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d654d90b-c5bc-428d-b1a9-013e95e4851d",
	}).Debug("Outgoing - 'ReadFromTestSuiteExecutionsMap'")

	// Lock Map for Reading
	testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.RUnlock()

	// Different logic depending on if call is for TestSuiteExecutionsList or for Detailed TestSuiteExecutionPreview
	switch readListOrDetailedMessagesType {

	// Messages for TestSuiteExecutionsListTable
	case ReadListMessages:

		// Check if Map i nil
		if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
			return nil, nil, false
		}

		// Read Map
		testSuiteExecutionsListMessages, existInMap = TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap[testSuiteExecutionsMapKey]

		return testSuiteExecutionsListMessages, nil, existInMap

	case readDetailedMessages:

		// Check if Map i nil
		if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
			return nil, nil, false
		}

		// Read Map
		var tempDetailedTestSuiteExecutionsObjectsMap map[DetailedTestSuiteExecutionMapKeyType]*DetailedTestSuiteExecutionsMapObjectStruct
		tempDetailedTestSuiteExecutionsObjectsMap = *TestSuiteExecutionsModel.DetailedTestSuiteExecutionsObjectsMapPtr
		var tempDetailedTestSuiteExecutionsMapObjectPtr *DetailedTestSuiteExecutionsMapObjectStruct
		tempDetailedTestSuiteExecutionsMapObjectPtr, existInMap = tempDetailedTestSuiteExecutionsObjectsMap[DetailedTestSuiteExecutionMapKeyType(testSuiteExecutionsMapKey)] //LatestTestSuiteExecutionForEachTestSuiteUuid.latestTestSuiteExecutionForEachTestSuiteUuidMap[testSuiteExecutionsMapKey]

		if existInMap == true {
			testSuiteExecutionResponseMessage = tempDetailedTestSuiteExecutionsMapObjectPtr.DetailedTestSuiteExecution
		}

		return nil, testSuiteExecutionResponseMessage, existInMap

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                             "2e3a6094-118b-4e2b-8fc3-34e415b6c27e",
			"readListOrDetailedMessagesType": readListOrDetailedMessagesType,
		}).Fatalln("Unhandled incoming 'readListOrDetailedMessagesType'")

	}
	return nil, nil, false

}

// ReadAllFromTestSuiteExecutionsMapForTableList
// Read all from the TestSuiteExecutions-Map for the TableList for TestSuiteExecutions
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadAllFromTestSuiteExecutionsMapForTableList() (
	testSuiteExecutionsListMessages *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "61403cf9-30ef-4e27-ab68-b294fbbfbd45",
	}).Debug("Incoming - 'ReadAllFromTestSuiteExecutionsMapForTableList'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "1cefa358-c352-4acd-a437-49eb1bf95595",
	}).Debug("Outgoing - 'ReadAllFromTestSuiteExecutionsMapForTableList'")

	testSuiteExecutionsListMessages, _ = testSuiteExecutionsModel.ReadAllFromTestSuiteExecutionsMap(ReadListMessages)

	return testSuiteExecutionsListMessages

}

// ReadAllFromTestSuiteExecutionsMap
// Read all from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) ReadAllFromTestSuiteExecutionsMap(
	readListOrDetailedMessagesType readListOrDetailedMessagesTypeType) (
	testSuiteExecutionsListMessages *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage,
	testSuiteExecutionResponseMessages *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionResponseMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                             "f11158ad-3495-42ab-85aa-68c7c2272c00",
		"readListOrDetailedMessagesType": readListOrDetailedMessagesType,
	}).Debug("Incoming - 'ReadAllFromTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9589f7cc-5329-4bb8-b8ba-8582c214773e",
	}).Debug("Outgoing - 'ReadAllFromTestSuiteExecutionsMap'")

	// Lock Map for Reading
	testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.RUnlock()

	// Different logic depending on if call is for TestSuiteExecutionsList or for Detailed TestSuiteExecutionPreview
	switch readListOrDetailedMessagesType {

	// Messages for TestSuiteExecutionsListTable
	case ReadListMessages:

		// Check if Map i nil
		if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {
			return nil, nil
		}

		var tempTestSuiteExecutionsListMessage []*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

		// Loop all items in map and add to response slice
		for _, testSuiteExecutionListMessage := range TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap {
			tempTestSuiteExecutionsListMessage = append(tempTestSuiteExecutionsListMessage, testSuiteExecutionListMessage)
		}

		return &tempTestSuiteExecutionsListMessage, nil

	case readDetailedMessages:

		// Not implemented should never be called
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                             "e158a5b3-573a-4e65-92ba-09e29bf5f075",
			"readListOrDetailedMessagesType": readListOrDetailedMessagesType,
		}).Fatalln("Not implemented should never be called")

		return nil, nil

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                             "728024b3-65ce-4b9f-8fd6-769e7787e524",
			"readListOrDetailedMessagesType": readListOrDetailedMessagesType,
		}).Fatalln("Unhandled incoming 'readListOrDetailedMessagesType'")

	}

	return nil, nil

}

// GetNumberOfTestSuiteExecutionsRetrievedFromDatabase
// Read all from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) GetNumberOfTestSuiteExecutionsRetrievedFromDatabase() (
	numberOfTestSuiteExecutionsRetrievedFromDatabase int) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "29fa838d-d52e-48bf-b399-7abf73cf4457",
	}).Debug("Incoming - 'GetNumberOfTestSuiteExecutionsRetrievedFromDatabase'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "27f1a1a9-bdbd-4fed-9ee4-764801be229a",
	}).Debug("Outgoing - 'GetNumberOfTestSuiteExecutionsRetrievedFromDatabase'")

	// Lock Map for Reading
	testSuiteExecutionsMapMutex.RLock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.RUnlock()

	var testSuiteExecutionsListMessage *[]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

	// Get all data retrieved and get hwo many they are
	testSuiteExecutionsListMessage = testSuiteExecutionsModel.ReadAllFromTestSuiteExecutionsMapForTableList()

	numberOfTestSuiteExecutionsRetrievedFromDatabase = len(*testSuiteExecutionsListMessage)

	return numberOfTestSuiteExecutionsRetrievedFromDatabase

}

// AddToTestSuiteExecutionsMap
// Add to the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) AddToTestSuiteExecutionsMap(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType,
	testSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "9984f6c6-996f-4039-b431-9a221ea36fa2",
	}).Debug("Incoming - 'AddToTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "11d9fa55-7bbe-4d8b-9115-848dc37774fb",
	}).Debug("Outgoing - 'AddToTestSuiteExecutionsMap'")

	// Lock Map for Writing
	testSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {

		TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap = make(
			map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage)
	}

	// Save to TestSuiteExecutions-Map
	TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap[testSuiteExecutionsMapKey] = testSuiteExecutionsListMessage

}

// DeleteFromTestSuiteExecutionsMap
// Delete from the TestSuiteExecutions-Map
func (testSuiteExecutionsModel TestSuiteExecutionsModelStruct) DeleteFromTestSuiteExecutionsMap(
	testSuiteExecutionsMapKey TestSuiteExecutionUuidType) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "85632715-c3a4-435c-bd5e-3a9870f3b690",
	}).Debug("Incoming - 'DeleteFromTestSuiteExecutionsMap'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "d1f707ac-8965-4119-98d2-930aad1cea78",
	}).Debug("Outgoing - 'DeleteFromTestSuiteExecutionsMap'")

	// Lock Map for Writing
	testSuiteExecutionsMapMutex.Lock()

	//UnLock Map
	defer testSuiteExecutionsMapMutex.Unlock()

	// Check if Map i nil
	if TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.latestTestSuiteExecutionForEachTestSuiteUuidMap == nil {

		TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
			latestTestSuiteExecutionForEachTestSuiteUuidMap = make(
			map[TestSuiteExecutionUuidType]*fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage)

		return
	}

	// Save to TestSuiteExecutions-Map
	delete(TestSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.
		latestTestSuiteExecutionForEachTestSuiteUuidMap,
		testSuiteExecutionsMapKey)

}
