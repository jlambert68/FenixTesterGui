
# cat -e -t -v Makefile

#.DEFAULT_GOAL := buildAndRun
# Migration steps for new Golang gRPC installation
# https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable


RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6672

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

BuildExeForWindows:
#	fyne-cross windows -arch=amd64 --ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X 'main.gcpAuthentication=false'"
#	GOOD=windows GOARCH=amd64 go build -o FenixCAConnectorWindow.exe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X  'main.gcpAuthentication=true' -X 'main.caEngineAddress=127.0.0.1' -X 'main.caEngineAddressPath=/"
	env GOOD=windows GOARCH=amd64 go build  -o FenixGui.WindowsExe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.executionConnectorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.executionWorkerAddress=fenixexecutionworker-ca-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.executionWorkerPort=443' -X 'main.gcpAuthentication=true' -X 'main.caEngineAddress=127.0.0.1' -X 'main.caEngineAddressPath=x' -X 'main.useInternalWebServerForTest=true' -X 'main.useServiceAccount=true'" /home/jlambert/egen_kod/go/go_workspace/src/jlambert/FenixCAConnector
BuildExeForLinux:
	$(eval fileName := $(filenamePartFirstLinux)$(datetime)$(filenamePartLinuxLast))
	GOOD=linux GOARCH=amd64 go build -o $(fileName)  -ldflags=" -X 'main.BuildVariableDB_HOST=xxxx' -X 'main.BuildVariableDB_NAME=xxxx' -X 'main.BuildVariableDB_PASS=xxxx' -X 'main.BuildVariableDB_PORT=xxxx' -X 'main.BuildVariableDB_SCHEMA=xxxx' -X 'main.BuildVariableDB_USER=xxxx' -X 'main.BuildVariableExecutionLocation=LOCALHOST_NODOCKER' -X 'main.BuildVariableExecutionLocationFenixGuiServer=GCP' -X 'main.BuildVariableFenixGuiBuilderProxyServerAddress=127.0.0.1' -X 'main.BuildVariableFenixGuiBuilderProxyServerAdminPort=6672' -X 'main.BuildVariableFenixGuiBuilderProxyServerPort=6671'-X 'main.BuildVariableFenixGuiBuilderServerAddress=fenixguitestcasebuilderserver-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.BuildVariableFenixGuiBuilderServerPort=443' -X 'main.BuildVariableRunAsTrayApplication=NO'" .

CrossBuildForWindows:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o $(fileName) -ldflags=" -X 'main.BuildVariableDB_HOST=xxxx' -X 'main.BuildVariableDB_NAME=xxxx' -X 'main.BuildVariableDB_PASS=xxxx' -X 'main.BuildVariableDB_PORT=xxxx' -X 'main.BuildVariableDB_SCHEMA=xxxx' -X 'main.BuildVariableDB_USER=xxxx' -X 'main.BuildVariableExecutionLocation=LOCALHOST_NODOCKER' -X 'main.BuildVariableExecutionLocationFenixGuiServer=GCP' -X 'main.BuildVariableFenixGuiBuilderProxyServerAddress=127.0.0.1' -X 'main.BuildVariableFenixGuiBuilderProxyServerAdminPort=6672' -X 'main.BuildVariableFenixGuiBuilderProxyServerPort=6671'-X 'main.BuildVariableFenixGuiBuilderServerAddress=fenixguitestcasebuilderserver-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.BuildVariableFenixGuiBuilderServerPort=443' -X 'main.BuildVariableRunAsTrayApplication=NO'" .