package messageEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"context"
	"fyne.io/fyne/v2"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"io"
)

// InitiateConnectorRequestForProcessTestInstructionExecution
// This gPRC-methods is used when a Execution Connector needs to have its TestInstruction assignments using reverse streaming
// Execution Connector opens the gPRC-channel and assignments are then streamed back to Connector from Worker
func (messageEngineServerObject *messageEngineServerStruct) InitiateGuiExecutioNServerRequestForMessages() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "50706674-cd4c-494f-a0ae-796a4f7dd13c",
	}).Debug("Incoming 'InitiateGuiExecutioNServerRequestForMessages'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "a188b25c-f252-4ce0-96e0-7c82898c805a",
	}).Debug("Outgoing 'InitiateGuiExecutioNServerRequestForMessages'")

	var ctx context.Context
	var returnMessageAckNack bool

	ctx = context.Background()

	// Set up connection to Server
	ctx, err := messageEngineServerObject.setConnectionToFenixGuiExecutionMessageServer(ctx)
	if err != nil {
		return
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background()) //, 30*time.Second)
	defer func() {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "0966c46c-0d3c-44b9-a069-2609dc505619",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP && sharedCode.GCPAuthentication == true {

		// Add Access token
		ctx, returnMessageAckNack, _ = gcp.Gcp.GenerateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			return
		}

	}

	// Set up call parameter
	emptyParameter := &fenixExecutionWorkerGrpcApi.EmptyParameter{
		ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(common_config.GetHighestExecutionWorkerProtoFileVersion())}

	// Start up streamClient from Worker server
	streamClient, err := fenixExecutionWorkerGrpcClient.ConnectorRequestForProcessTestInstructionExecution(ctx, emptyParameter)

	// Couldn't connect to Worker
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "d9ab0434-1121-4e2e-95e7-3e1cc99656b0",
			"err": err,
		}).Error("Couldn't open streamClient from Worker Server. Will wait 1 second and try again")

		return
	}

	// Local channel to decide when Server stopped sending
	done := make(chan bool)

	// Connection to Worker is up so change Tray icon to "green", BUT ONLY when run as Tray App
	var myApplication fyne.App
	if common_config.FenixCAConnectorApplicationReference != nil {
		myApplication = *common_config.FenixCAConnectorApplicationReference
		myApplication.SetIcon(resources.ResourceFenix83green32x32Png)

	}

	// Run streamClient receiver as a go-routine
	go func() {
		for {
			processTestInstructionExecutionReveredRequest, err := streamClient.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":  "3439f49f-d7d5-477e-9a6b-cfa5ed355bfe",
					"err": err,
				}).Error("Got some error when receiving TestInstructionExecutionsRequests from Worker, reconnect in 1 second")

				done <- true //close(done)
				return

			}

			// Check if message counts as a "keep Alive message, message is 'nil
			if processTestInstructionExecutionReveredRequest.TestInstruction.TestInstructionName == "KeepAlive" {
				// Is a keep alive message
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "08b86c8d-81ba-4664-8cb5-8e53140dc870",
					"processTestInstructionExecutionReveredRequest": processTestInstructionExecutionReveredRequest,
				}).Debug("'Keep alive' message received from Worker")

			} else {
				// Is a standard TestInstruction to execute by Connector backend
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID": "d1ea4370-3e8e-4d2b-9626-a193213e091a",
					"processTestInstructionExecutionReveredRequest": processTestInstructionExecutionReveredRequest,
				}).Debug("Receive TestInstructionExecution from Worker")

				// Send response and start processing TestInstruction in parallell
				go func() {

					// Call 'CA' backend to convert 'TestInstruction' into useful structure later to be used by FangEngine
					var processTestInstructionExecutionReversedResponse *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionReversedResponse
					var fangEngineRestApiMessageValues *restCallsToCAEngine.FangEngineRestApiMessageStruct
					processTestInstructionExecutionReversedResponse, fangEngineRestApiMessageValues = ConvertTestInstructionIntoFangEngineStructure(processTestInstructionExecutionReveredRequest)

					// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
					couldSend, returnMessage := toExecutionWorkerObject.SendConnectorProcessTestInstructionExecutionReversedResponseToFenixWorkerServer(processTestInstructionExecutionReversedResponse)

					if couldSend == false {
						sharedCode.Logger.WithFields(logrus.Fields{
							"ID":            "95dddb21-0895-4016-9cb5-97ab4568f30b",
							"returnMessage": returnMessage,
						}).Error("Couldn't send response to Worker")

					} else {

						// Send TestInstruction to FangEngine using RestCall
						var finalTestInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage
						finalTestInstructionExecutionResultMessage = SendTestInstructionToFangEngineUsingRestCall(fangEngineRestApiMessageValues, processTestInstructionExecutionReveredRequest)

						// Send 'ProcessTestInstructionExecutionReversedResponse' back to worker over direct gRPC-call
						couldSend, returnMessage := toExecutionWorkerObject.SendReportCompleteTestInstructionExecutionResultToFenixWorkerServer(finalTestInstructionExecutionResultMessage)

						if couldSend == false {
							sharedCode.Logger.WithFields(logrus.Fields{
								"ID": "95dddb21-0895-4016-9cb5-97ab4568f30b",
								"finalTestInstructionExecutionResultMessage": finalTestInstructionExecutionResultMessage,
								"returnMessage": returnMessage,
							}).Error("Couldn't send response to Worker")
						}
					}
				}()

			}

		}
	}()

	// Server stopped sending so reconnect again in 5 seconds
	<-done

	sharedCode.Logger.WithFields(logrus.Fields{
		"ID": "0b5fdb7c-91aa-4dfc-b587-7b6cef83d224",
	}).Debug("Server stopped sending so reconnect again in 5 seconds")

}
