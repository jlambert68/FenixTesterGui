
# cat -e -t -v Makefile

#.DEFAULT_GOAL := buildAndRun
# Migration steps for new Golang gRPC installation
# https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable


RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6668

filename :=
filenamePartFirst := FenixGuiCrossBuild_
filenamePartLast := .exe
filenamePartFirstLinux=FenixBuild
filenamePartLinuxLast = .Linux
datetime := `date +'%y%m%d_%H%M%S'`

GenerateDateTime:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))

	echo $(fileName)

GenerateTrayIcons:
	./bundleIcons.sh

BuildExeForLinux:
	$(eval fileName := $(filenamePartFirstLinux)$(datetime)$(filenamePartLinuxLast))
	GOOD=linux GOARCH=amd64 go build -o $(fileName) -ldflags=" -X 'main.BuildVariableGCPAuthentication=false' -X 'main.BuildVariableUseServiceAccountForGuiExecutionServer=true' -X 'main.BuildVariableExecutionLocationForFenixGuiExecutionServer=GCP' -X 'main.BuildVariableFenixGuiExecutionServerAddress=fenixguiexecutionserver-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.BuildVariableFenixGuiExecutionServerPort=443' -X 'main.BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer=GCP' -X 'main.BuildVariableFenixGuiTestCaseBuilderServerAddress=fenixguitestcasebuilderserver-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.BuildVariableFenixGuiTestCaseBuilderServerPort=443' -X 'main.BuildVariableExecutionLocationForThisApplication=LOCALHOST_NODOCKER' -X 'main.BuildVariableFYNE_SCALE=0.6' -X 'main.BuildVariableRunAsTrayApplication=NO' -X 'main.BuildVariableApplicationGrpcPort=6668'" .

CrossBuildForWindows:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o $(fileName) -ldflags=" -X 'main.BuildVariableGCPAuthentication=false' -X 'main.BuildVariableUseServiceAccountForGuiExecutionServer=true' -X 'main.BuildVariableExecutionLocationForFenixGuiExecutionServer=GCP' -X 'main.BuildVariableFenixGuiExecutionServerAddress=fenixguiexecutionserver-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.BuildVariableFenixGuiExecutionServerPort=443' -X 'main.BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer=GCP' -X 'main.BuildVariableFenixGuiTestCaseBuilderServerAddress=fenixguitestcasebuilderserver-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.BuildVariableFenixGuiTestCaseBuilderServerPort=443' -X 'main.BuildVariableExecutionLocationForThisApplication=LOCALHOST_NODOCKER' -X 'main.BuildVariableFYNE_SCALE=0.6' -X 'main.BuildVariableRunAsTrayApplication=NO' -X 'main.BuildVariableApplicationGrpcPort=6668'" .
