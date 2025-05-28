package messageStreamEngine

import (
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"google.golang.org/grpc"
)

type MessageStreamEngineStruct struct {
	ongoingTimerOrConnectionForCallingWorkerForTestInstructionsToExecute bool
}

var messageStreamEngineObject MessageStreamEngineStruct

var (
	// Standard gRPC Client
	remoteFenixGuiExecutionServerConnection *grpc.ClientConn

	fenixGuiExecutionServerSubscribeToMessagesClient fenixExecutionServerGuiGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClientClient
)

var highestFenixGuiExecutionServerProtoFileVersion int32 = -1

// Parameters used for channel to update status on Executions for TestCasesMap and TestInstructions
var executionStatusCommandChannel ExecutionStatusChannelType

type ExecutionStatusChannelType chan ChannelCommandStruct

const messageChannelMaxSize int32 = 100

type ChannelCommandType uint8

const (
	ChannelCommandExecutionsStatusesHaveBeUpdated ChannelCommandType = iota
	ChannelCommandTriggerRequestForTestInstructionExecutionToProcess
	ChannelCommandTriggerRequestForTestInstructionExecutionToProcessIn1Second
)

type ChannelCommandStruct struct {
	ChannelCommand          ChannelCommandType
	ExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage
}
