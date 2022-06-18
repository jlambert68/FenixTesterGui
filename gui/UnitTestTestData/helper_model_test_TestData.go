package UnitTestTestData

import (
	common_config "FenixTesterGui/common_code"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// Availalbe Buoilding Blocks; ABB

// ******* START ABB001 *******
// Result when asking for Available Building Blocks
var TestInstructionsAndTestInstructionsContainersRespons_ABB001 string = `{
		"TestInstructionMessages": [
	{
	"DomainUuid": "78a97c41-a098-4122-88d2-01ed4b6c4844",
	"DomainName": "Custody Arrangement",
	"TestInstructionUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
	"TestInstructionName": "Just the name",
	"TestInstructionTypeUuid": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
	"TestInstructionTypeName": "The type of the TestInstruction",
	"TestInstructionDescription": "En vanlig typ",
	"TestInstructionMouseOverText": "This will be shown when hovering above this TestInstruction",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-04-29T15:42:15Z"
	}
	],
	"TestInstructionContainerMessages": [
	{
	"DomainUuid": "e81b9734-5dce-43c9-8d77-3368940cf126",
	"DomainName": "Fenix",
	"TestInstructionContainerUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
	"TestInstructionContainerName": "Emtpy serial processed TestInstructionsContainer",
	"TestInstructionContainerTypeUuid": "b107bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerTypeName": "Base containers",
	"TestInstructionContainerDescription": "Children of this container is processed in serial",
	"TestInstructionContainerMouseOverText": "Children of this container is processed in serial",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-05-02T10:10:07Z",
	"TestInstructionContainerExecutionType": "SERIAL_PROCESSED",
	"TestInstructionContainerChildren": []
	},
	{
	"DomainUuid": "e81b9734-5dce-43c9-8d77-3368940cf126",
	"DomainName": "Fenix",
	"TestInstructionContainerUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerName": "Emtpy parallelled processed TestInstructionsContainer",
	"TestInstructionContainerTypeUuid": "b107bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerTypeName": "Base containers",
	"TestInstructionContainerDescription": "Children of this container is processed in parallel",
	"TestInstructionContainerMouseOverText": "Children of this container is processed in parallel",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-05-02T10:08:28Z",
	"TestInstructionContainerExecutionType": "PARALLELLED_PROCESSED",
	"TestInstructionContainerChildren": []
	},
	{
	"DomainUuid": "78a97c41-a098-4122-88d2-01ed4b6c4844",
	"DomainName": "Custody Arrangement",
	"TestInstructionContainerUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
	"TestInstructionContainerName": "Emtpy parallelled processed Turbo TestInstructionsContainer",
	"TestInstructionContainerTypeUuid": "ca07bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerTypeName": "CA Base containers",
	"TestInstructionContainerDescription": "Children of this CA-container is processed in parallel",
	"TestInstructionContainerMouseOverText": "Children of this CA-container is processed in parallel",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-06-16T16:09:43Z",
	"TestInstructionContainerExecutionType": "PARALLELLED_PROCESSED",
	"TestInstructionContainerChildren": []
	}
	],
	"ackNackResponse": {
	"AckNack": true,
	"Comments": "",
	"ErrorCodes": []
	}
	}`

// *** Expected result, regarding TestInstructions, when asking for Available Building Blocks

// All TestDomains, TestInstructionTypes and TestInstructions
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001 string = "map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name} 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:{The type of the TestInstruction [513dd8f] 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction} 78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement}]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name}]]]"

// *** Expected result, regarding TestInstructionContainerss, when asking for Available Building Blocks

// All TestDomains, TestInstructionContainerTypes and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement} aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer} b107bdd9-4152-4020-b3f0-fc750b45885e:{Base containers [b107bdd] b107bdd9-4152-4020-b3f0-fc750b45885e Base containers} ca07bdd9-4152-4020-b3f0-fc750b45885e:{CA Base containers [ca07bdd] ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers} e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} e81b9734-5dce-43c9-8d77-3368940cf126:{Fenix [e81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[ca07bdd9-4152-4020-b3f0-fc750b45885e:map[aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer}]] e81b9734-5dce-43c9-8d77-3368940cf126:map[b107bdd9-4152-4020-b3f0-fc750b45885e:map[e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]]]"

// *** Expected result, regarding TestInstructions & TestInstructionContainerss, when asking for Available Building Blocks

// All TestDomains, TestInstructionTypes, TestInstructions, TestInstructionContainerTypes and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005 string = "map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name} 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:{The type of the TestInstruction [513dd8f] 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction} 78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement} aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer} b107bdd9-4152-4020-b3f0-fc750b45885e:{Base containers [b107bdd] b107bdd9-4152-4020-b3f0-fc750b45885e Base containers} ca07bdd9-4152-4020-b3f0-fc750b45885e:{CA Base containers [ca07bdd] ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers} e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} e81b9734-5dce-43c9-8d77-3368940cf126:{Fenix [e81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name}]]]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[ca07bdd9-4152-4020-b3f0-fc750b45885e:map[aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer}]] e81b9734-5dce-43c9-8d77-3368940cf126:map[b107bdd9-4152-4020-b3f0-fc750b45885e:map[e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]]]"

// ******* END ABB001 *******

// Init the logger for UnitTests
func InitLoggerForTest(filename string) (myTestLogger *logrus.Logger) {
	myTestLogger = logrus.StandardLogger()

	switch common_config.LoggingLevel {

	case logrus.DebugLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	case logrus.InfoLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	case logrus.WarnLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	default:
		log.Println("Not correct value for debugging-level, this was used: ", common_config.LoggingLevel)
		os.Exit(0)

	}

	logrus.SetLevel(common_config.LoggingLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out

	if filename == "" {
		myTestLogger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			myTestLogger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}

	// Should only be done from init functions
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))

	return myTestLogger
}

// ********************************************************************************************************************
// Check if testdata is using correct proto-file version
func IsTestDataUsingCorrectTestDataProtoFileVersion(usedProtoFileVersion fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum) (returnMessage *fenixTestCaseBuilderServerGrpcApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum
	var protoFileUsed fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(getHighestFenixTestDataProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}
		/*
			fenixGuiTestCaseBuilderServerObject.logger.WithFields(logrus.Fields{
				"id": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
			}).Debug("Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "' for Client: " + callingClientUuid)
		*/
		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
var highestFenixProtoFileVersion int32

// Get the highest FenixProtoFileVersionEnumeration
func getHighestFenixTestDataProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixProtoFileVersion != -1 {
		return highestFenixProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixProtoFileVersion = maxValue

	return highestFenixProtoFileVersion
}
