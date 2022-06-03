package fenixTestGuiObject

import (
	"FenixTesterGui/grpc_in"
	"FenixTesterGui/grpc_out"
	"FenixTesterGui/restAPI"
)

type referencesStruct struct {
	GrpcIn  *grpc_in.GRPCInStruct
	GrpcOut *grpc_out.GRPCOutStruct
	RestAPI *restAPI.RestApiStruct
}

type FenixGuiBuilderProxyServerObjectStruct struct {
	LocalReferences referencesStruct
}
