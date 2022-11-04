package grpc_in

import (
	sharedCode "FenixTesterGui/common_code"
	fenixUserGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixUserGui/fenixUserGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// Used to keep track of latest Proto-file version
var highestFenixProtoFileVersion int32 = -1

// IsClientUsingCorrectTestDataProtoFileVersion ********************************************************************************************************************
// Check if Calling Client is using correct proto-file version
func IsClientUsingCorrectTestDataProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum) (returnMessage *fenixUserGuiGrpcApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum
	var protoFileUsed fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum(getHighestFenixUserGuiServerProtoFileVersion())

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
		var errorCodes []fenixUserGuiGrpcApi.ErrorCodesEnum
		var errorCode fenixUserGuiGrpcApi.ErrorCodesEnum

		errorCode = fenixUserGuiGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixUserGuiGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
		}).Info("Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "' for Client: " + callingClientUuid)

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
// Get the highest FenixProtoFileVersionEnumeration
func getHighestFenixUserGuiServerProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixProtoFileVersion != -1 {
		return highestFenixProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixUserGuiGrpcApi.CurrentFenixUserGuiProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixProtoFileVersion = maxValue

	return highestFenixProtoFileVersion
}

// SetLogger
// Set to use the same logger reference as is used by central part of system
func (grpcIn *GRPCInStruct) SetLogger(logger *logrus.Logger) {
	grpcIn.logger = logger

	return

}
