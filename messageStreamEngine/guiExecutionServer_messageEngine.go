package messageStreamEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/gcp"
	"FenixTesterGui/grpc_out_GuiExecutionServer"
	"context"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"io"
)

// TesterGui opens the gPRC-channel to have messages streamed back to TesterGui from GuiExecutionServer
func (messageStreamEngineObject *MessageStreamEngineStruct) initiateGuiExecutionServerRequestForMessages() {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "3a20b488-3b03-4be7-b2ea-0cae5c82ffc4",
	}).Debug("Incoming 'initiateGuiExecutionServerRequestForMessages'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "23978650-fd08-4b6d-b6a4-238ac030c207",
	}).Debug("Outgoing 'initiateGuiExecutionServerRequestForMessages'")

	var ctx context.Context
	var returnMessageAckNack bool

	ctx = context.Background()

	// Set up connection to Server
	ctx, err := messageStreamEngineObject.setConnectionToFenixGuiExecutionMessageServer(ctx)
	if err != nil {
		return
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background()) //, 30*time.Second)
	defer func() {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "19f720f0-bfe1-4b32-b605-02b094746169",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if sharedCode.ExecutionLocationForFenixGuiExecutionServer == sharedCode.GCP && sharedCode.GCPAuthentication == true {

		// Add Access token
		ctx, returnMessageAckNack, _ = gcp.GcpObject.GenerateGCPAccessToken(ctx, gcp.TargetServerGuiExecutionServer)
		if returnMessageAckNack == false {
			return
		}

	}

	// Set up call parameter
	userAndApplicationRunTimeIdentificationMessage := &fenixExecutionServerGuiGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
		ApplicationRunTimeUuid: sharedCode.ApplicationRunTimeUuid,
		UserId:                 sharedCode.CurrentUserId,
		ProtoFileVersionUsedByClient: fenixExecutionServerGuiGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
			grpc_out_GuiExecutionServer.GetHighestFenixGuiExecutionServerProtoFileVersion()),
	}

	// Start up streamClient from TestGui
	streamClient, err := fenixGuiExecutionServerSubscribeToMessagesClient.SubscribeToMessageStream(ctx, userAndApplicationRunTimeIdentificationMessage)

	// Couldn't connect to GuiExecutionServer
	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "436d9106-d1f7-4b62-ac97-fa8850ea32d1",
			"err": err,
		}).Error("Couldn't open streamClient from GuiExecutionServer. Will wait 1 second and try again")

		return
	}

	// Local channel to decide when Server stopped sending
	done := make(chan bool)

	// Connection to GuiExecutionServer is up so change Tray icon to "green", BUT ONLY when run as Tray App
	/*var myApplication fyne.App
	if common_config.FenixCAConnectorApplicationReference != nil {
		myApplication = *common_config.FenixCAConnectorApplicationReference
		myApplication.SetIcon(resources.ResourceFenix83green32x32Png)

	}
	*/

	// Run streamClient receiver as a go-routine
	go func() {
		for {
			subscribeToMessagesStreamResponse, err := streamClient.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":  "03a06dd6-36ef-4a3f-8e8c-69cd777a6739",
					"err": err,
				}).Error("Got some error when receiving 'subscribeToMessagesStreamResponse'-message from GuiExecutionServer, reconnect in a short time")

				done <- true //close(done)
				return

			}

			// Check if message counts as a "keep Alive message, message is 'nil
			if subscribeToMessagesStreamResponse.IsKeepAliveMessage == true {
				// Is a keep alive message
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":                                "b8c4e3fc-0fe6-439d-9a23-b4954e6d9f2c",
					"subscribeToMessagesStreamResponse": subscribeToMessagesStreamResponse,
				}).Debug("'Keep alive' message received from GuiExecutionServer")

			} else {
				// Is a standard subscribeToMessagesStreamResponse to process
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":                                "b845c8e3-8854-49d5-ad8c-f66431e40540",
					"subscribeToMessagesStreamResponse": subscribeToMessagesStreamResponse,
				}).Debug("Receive subscribeToMessagesStreamResponse from GuiExecutionServer")

				// Split incoming message from GuiExecutionServer into its individual messages and put on channel to MessagechannelEngine
				if subscribeToMessagesStreamResponse.ExecutionsStatus != nil {
					var channelCommandAndMessage ChannelCommandStruct
					channelCommandAndMessage = ChannelCommandStruct{
						ChannelCommand:          ChannelCommandExecutionsStatusesHaveBeUpdated,
						ExecutionsStatusMessage: subscribeToMessagesStreamResponse.ExecutionsStatus,
					}

					// Send message on channel
					executionStatusCommandChannel <- channelCommandAndMessage
				}

			}

		}
	}()

	// Server stopped sending so reconnect again after a short wait
	<-done

	/*
		if common_config.FenixCAConnectorApplicationReference != nil {
			myApplication = *common_config.FenixCAConnectorApplicationReference
			myApplication.SetIcon(resources.ResourceFenix83green32x32Png)

		}*/

	sharedCode.Logger.WithFields(logrus.Fields{
		"ID": "0b5fdb7c-91aa-4dfc-b587-7b6cef83d224",
	}).Debug("Server stopped sending so reconnect again in")

}
