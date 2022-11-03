package restAPI

import (
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

// Structs used when converting json messages in restAPI

type RestUserMessageStruct struct {
	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

type RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct struct {
	UserId                                 string                                                                                           `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	PinnedTestInstructionMessages          []*fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionMessage                    `protobuf:"bytes,3,rep,name=AvailablePinnedTestInstructionMessage,proto3" json:"PinnedTestInstructionMessages,omitempty"`
	PinnedTestInstructionContainerMessages []*fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedPreCreatedTestInstructionContainerMessage `protobuf:"bytes,4,rep,name=AvailablePinnedPreCreatedTestInstructionContainerMessage,proto3" json:"PinnedTestInstructionContainerMessages,omitempty"`
}

func (restAPI *RestApiStruct) RestAPIServer() {
	log.Println("starting API server")

	// Setting logger in gRPC Out
	//restAPI.GrpcOut.SetLogger(restAPI.logger)

	// Setting Dial address to call GuiServer on
	restAPI.GrpcOut.SetDialAddressString(restAPI.fenixGuiBuilderServerAddressToDial)

	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", restAPI.HealthCheck).Methods("GET")
	router.HandleFunc("/are-guibuilderserver-alive", restAPI.RestSendAreYouAliveToFenixGuiBuilderServer).Methods("GET")
	router.HandleFunc("/testinstructions-and-testinstructioncontainers", restAPI.RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer).Methods("GET")
	router.HandleFunc("/pinned-testinstructions-and-testinstructioncontainers", restAPI.RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer).Methods("GET")
	router.HandleFunc("/pinned-testinstructions-and-testinstructioncontainers", restAPI.RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer).Methods("POST")

	http.Handle("/", router)

	//start and listen to requests
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(" Couldn't start RestServer")
	}

}

func (restAPI *RestApiStruct) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	// curl --request GET localhost:8080/health-check

	restAPI.logger.WithFields(logrus.Fields{
		"id": "fb3c1ecb-3da8-4d27-b1c4-16d5120e7125",
	}).Debug("Incoming 'RestApi - /health-check'")

	defer restAPI.logger.WithFields(logrus.Fields{
		"id": "fab7676d-c303-4b20-8980-397d7a59282e",
	}).Debug("Outgoing 'RestApi - /health-check'")

	// Set OK in Header
	w.WriteHeader(http.StatusOK)

	// Create Response message
	_, err := fmt.Fprintf(w, "API is up and running")
	if err != nil {
		log.Fatalln(" Couldn't create Response message")
	}

}

func (restAPI *RestApiStruct) RestSendAreYouAliveToFenixGuiBuilderServer(w http.ResponseWriter, _ *http.Request) {
	// curl --request GET localhost:8080/are-guibuilderserver-alive

	restAPI.logger.WithFields(logrus.Fields{
		"id": "0645d30c-4479-49ab-bb72-9bc3fac329a5",
	}).Debug("Incoming 'RestApi - /are-guibuilderserver-alive'")

	defer restAPI.logger.WithFields(logrus.Fields{
		"id": "cc168cfe-3544-4946-93d4-d2325893f8cd",
	}).Debug("Outgoing 'RestApi - /are-guibuilderserver-alive'")

	// gRPC -response
	var response *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse

	// Do gRPC-call
	response = restAPI.GrpcOut.SendAreYouAliveToFenixGuiBuilderServer()

	// Create Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convert gRPC-response into json
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// if error then just exit TODO Create correct response message
		return
	}

	// Create Response message
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalln(" Couldn't create Response message")
	}

}

func (restAPI *RestApiStruct) RestSendGetInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request) {
	/*
		curl -X GET \
		localhost:8080/testinstructions-and-testinstructioncontainers \
		-H 'Content-Type: application/json' \
		-d '{"UserId":"s41797"}'
	*/
	restAPI.logger.WithFields(logrus.Fields{
		"id": "0645d30c-4479-49ab-bb72-9bc3fac329a5",
	}).Debug("Incoming 'RestApi - (GET) /testinstructions-and-testinstructioncontainers'")

	defer restAPI.logger.WithFields(logrus.Fields{
		"id": "cc168cfe-3544-4946-93d4-d2325893f8cd",
	}).Debug("Outgoing 'RestApi - (GET) /testinstructions-and-testinstructioncontainers'")

	// gRPC -response
	var response *fenixGuiTestCaseBuilderServerGrpcApi.AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Variable where Rest-json-payload will end up in
	jsonData := &RestUserMessageStruct{}

	// Extract and Validate json body
	err := extractAndValidateJsonBody(&w, r, jsonData)
	if err != nil {
		// If something went wrong then just exit
		return
	}

	// Do gRPC-call
	response = restAPI.GrpcOut.SendListAllAvailableTestInstructionsAndTestInstructionContainers(jsonData.UserId)

	// Create Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convert gRPC-response into json
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// if error then just exit TODO Create correct response message
		_, err := fmt.Fprintf(w, err.Error())
		if err != nil {
			log.Fatalln(" Couldn't create Response message")
		}

		return
	}

	// Create Response message
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalln(" Couldn't create Response message")
	}
}

func (restAPI *RestApiStruct) RestSendGetPinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request) {
	// curl --request GET localhost:8080/pinned-testinstructions-and-testinstructioncontainers/s41797
	/*
		curl -X GET \
		localhost:8080/pinned-testinstructions-and-testinstructioncontainers \
		-H 'Content-Type: application/json' \
		-d '{"UserId":"s41797"}'
	*/

	restAPI.logger.WithFields(logrus.Fields{
		"id": "2472dda1-701d-4b23-8326-757e43df4af4",
	}).Debug("Incoming 'RestApi - /pinned-testinstructions-and-testinstructioncontainers'")

	defer restAPI.logger.WithFields(logrus.Fields{
		"id": "db318ff4-ad36-43d4-a8d4-3e0ac4ff08c6",
	}).Debug("Outgoing 'RestApi - /pinned-testinstructions-and-testinstructioncontainers'")

	// gRPC -response
	var response *fenixGuiTestCaseBuilderServerGrpcApi.AvailablePinnedTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage

	// Variable where Rest-json-payload will end up in
	jsonData := &RestUserMessageStruct{}

	// Extract and Validate json body
	err := extractAndValidateJsonBody(&w, r, jsonData)
	if err != nil {
		// If something went wrong then just exit
		return
	}

	// Do gRPC-call
	response = restAPI.GrpcOut.SendListAllAvailablePinnedTestInstructionsAndTestInstructionContainers(jsonData.UserId)

	// Create Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convert gRPC-response into json
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// if error then just exit TODO Create correct response message
		_, err := fmt.Fprintf(w, err.Error())
		if err != nil {
			log.Fatalln(" Couldn't create Response message")
		}

		return
	}

	// Create Response message
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalln(" Couldn't create Response message")
	}
}

func (restAPI *RestApiStruct) RestSendSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServer(w http.ResponseWriter, r *http.Request) {
	// curl --request POST localhost:8080/pinned-testinstructions-and-testinstructioncontainers/s41797
	/*
		curl -X POST localhost:8080/pinned-testinstructions-and-testinstructioncontainers \
		-H 'Content-Type: application/json' \
		-d '{"UserId":"s41797","PinnedTestInstructionMessages":[{"TestInstructionUuid":"2f130d7e-f8aa-466f-b29d-0fb63608c1a6","TestInstructionName":"TestInstructionName 1"}],"PinnedTestInstructionContainerMessages":[{"TestInstructionContainerUuid":"b107bdd9-4152-4020-b3f0-fc750b45885e","TestInstructionContainerName":"TestInstructionContainerName 1"},{"TestInstructionContainerUuid":"e81b9734-5dce-43c9-8d77-3368940cf126","TestInstructionContainerName":"TestInstructionContainerName"}]}'
	*/
	// curl -X POST localhost:8080/pinned-testinstructions-and-testinstructioncontainers -H 'Content-Type: application/json' -d '{"UserId":"s41797","PinnedTestInstructionMessages":[{"TestInstructionUuid":"myUuid", "TestInstructionName":"myName"}],"PinnedTestInstructionContainerMessages":[{"TestInstructionContainerUuid":"myUuid2", "TestInstructionContainerName":"myName2"}]}'
	restAPI.logger.WithFields(logrus.Fields{
		"id": "2472dda1-701d-4b23-8326-757e43df4af4",
	}).Debug("Incoming 'RestApi - (POST) /pinned-testinstructions-and-testinstructioncontainers'")

	defer restAPI.logger.WithFields(logrus.Fields{
		"id": "db318ff4-ad36-43d4-a8d4-3e0ac4ff08c6",
	}).Debug("Outgoing 'RestApi - (POST) /pinned-testinstructions-and-testinstructioncontainers'")

	// Variable where Rest-json-payload will end up in
	jsonData := &RestSavePinnedInstructionsAndTestInstructionContainersToFenixGuiBuilderServerStruct{}

	// Extract and Validate json body
	err := extractAndValidateJsonBody(&w, r, jsonData)
	if err != nil {
		// If something went wrong then just exit
		return
	}

	// Create gRPC -response variable
	var response *fenixGuiTestCaseBuilderServerGrpcApi.AckNackResponse

	grpcOut := grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{}

	// Create input message for gRPC-call
	pinnedTestInstructionsAndTestContainersMessage := &fenixGuiTestCaseBuilderServerGrpcApi.SavePinnedTestInstructionsAndPreCreatedTestInstructionContainersMessage{
		UserId: jsonData.UserId,
		ProtoFileVersionUsedByClient: fenixGuiTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(
			grpcOut.GetHighestFenixGuiTestCaseBuilderServerProtoFileVersion()),
		AvailablePinnedTestInstructions:                    jsonData.PinnedTestInstructionMessages,
		AvailablePinnedPreCreatedTestInstructionContainers: jsonData.PinnedTestInstructionContainerMessages,
	}

	// Do gRPC-call
	response = restAPI.GrpcOut.SendSaveAllPinnedTestInstructionsAndTestInstructionContainers(pinnedTestInstructionsAndTestContainersMessage)

	// Create Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convert gRPC-response into json
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		// if error then just exit TODO Create correct response message
		_, err := fmt.Fprintf(w, err.Error())
		if err != nil {
			log.Fatalln(" Couldn't create Response message")
		}

		return
	}

	// Create Response message
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalln(" Couldn't create Response message")
	}
}
