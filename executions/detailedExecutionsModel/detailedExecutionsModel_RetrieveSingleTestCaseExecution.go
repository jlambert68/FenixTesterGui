package detailedExecutionsModel

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"errors"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// RetrieveSingleTestCaseExecution
// Retrieves a TestCaseExecution and all of its data belonging to the execution
// External call from (commandAndRuleEngine/commandEngineChannelReader.go) via command 'sharedCode.ChannelCommandExecuteTestCase:'
func RetrieveSingleTestCaseExecution(testCaseExecutionKey string) (err error) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id":                   "4be0dcc6-794c-4b35-8adb-9b04f88ac712",
		"testCaseExecutionKey": testCaseExecutionKey,
	}).Debug("Incoming 'RetrieveSingleTestCaseExecution'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "426abcf5-f18b-400f-818d-a4ae8fa73009",
	}).Debug("Outgoing 'RetrieveSingleTestCaseExecution'")

	// Extract individual parts of the 'TestCaseExecutionKeyMessage'
	var testCaseExecutionUuid string
	var testCaseExecutionVersion int
	var testCaseExecutionVersionError error

	testCaseExecutionUuid = testCaseExecutionKey[:len(testCaseExecutionKey)-1]

	testCaseExecutionVersion, testCaseExecutionVersionError = strconv.Atoi(testCaseExecutionKey[len(testCaseExecutionKey)-1:])
	if testCaseExecutionVersionError != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":                   "a7e1c59a-5e5f-47f2-ba7d-1a909eb90d68",
			"testCaseExecutionKey": testCaseExecutionKey,
			"testCaseExecutionKey[len(testCaseExecutionKey):]": testCaseExecutionKey[len(testCaseExecutionKey):],
		}).Error("Couldn't convert 'TestCaseExecutionVersion' from TestCaseExecutionKey into an integer")

		return testCaseExecutionVersionError
	}

	var testCaseExecutionKeyMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionKeyMessage
	testCaseExecutionKeyMessage = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionKeyMessage{
		TestCaseExecutionUuid:    testCaseExecutionUuid,
		TestCaseExecutionVersion: uint32(testCaseExecutionVersion),
	}

	var getSingleTestCaseExecutionRequest *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest
	getSingleTestCaseExecutionRequest = &fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionRequest{
		UserAndApplicationRunTimeIdentification: &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
			UserIdOnComputer:       sharedCode.CurrentUserIdLogedInOnComputer,
			GCPAuthenticatedUser:   sharedCode.CurrentUserAuthenticatedTowardsGCP,
			ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
		},
		TestCaseExecutionKey: testCaseExecutionKeyMessage,
	}

	// Do gRPC-call
	var getSingleTestCaseExecutionResponse *fenixExecutionServerGuiGrpcApi.GetSingleTestCaseExecutionResponse
	getSingleTestCaseExecutionResponse = grpc_out_GuiExecutionServer.GrpcOutGuiExecutionServerObject.
		SendGetSingleTestCaseExecution(getSingleTestCaseExecutionRequest)

	if getSingleTestCaseExecutionResponse.AckNackResponse.AckNack == false {
		return errors.New(getSingleTestCaseExecutionResponse.AckNackResponse.Comments)
	} else {

		// Add TestCaseExecution-details to repository via channelEngine
		var channelCommandDetailedExecutions ChannelCommandDetailedExecutionsStruct
		channelCommandDetailedExecutions = ChannelCommandDetailedExecutionsStruct{
			ChannelCommandDetailedExecutionsStatus:                            ChannelCommandFullDetailedExecutionsStatusUpdate,
			FullTestCaseExecutionResponseMessage:                              getSingleTestCaseExecutionResponse.TestCaseExecutionResponse,
			TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage: nil,
		}

		// Don't put on Channel if more than 9 items from max capacity
		var currentChannelSize int32
		currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
		if currentChannelSize > MessageChannelMaxSizeDetailedExecutionStatus-9 {
			for {
				time.Sleep(5 * time.Second)

				currentChannelSize = int32(len(DetailedExecutionStatusCommandChannel))
				if currentChannelSize < MessageChannelMaxSizeDetailedExecutionStatus-9 {
					break
				}
			}
		}

		// Send command ion channel
		DetailedExecutionStatusCommandChannel <- channelCommandDetailedExecutions

		sharedCode.Logger.WithFields(logrus.Fields{
			"Id":                                 "dd4b0e47-e8e2-4069-aaf9-371d7bc2b91d",
			"getSingleTestCaseExecutionResponse": getSingleTestCaseExecutionResponse,
		}).Debug("'getSingleTestCaseExecutionResponse' after calling GuiExecutionServer")

		return nil
	}
}
