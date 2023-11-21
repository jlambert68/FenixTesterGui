package main

import (
	sharedCode "FenixTesterGui/common_code"
	_ "embed"
	"strings"

	//"flag"
	"fmt"
	"os"

	//"github.com/getlantern/systray"
	//"github.com/getlantern/systray/example/icon"

	_ "net/http/pprof"
	"os/user"
)

// Embedded resources into the binary
// The icon used
//
//go:embed resources/fenix_icon_32x32_icon.ico
var embededfenixIcon []byte

func main() {
	//time.Sleep(15 * time.Second)
	//defer profile.Start(profile.ProfilePath(".")).Stop()

	// Used for profiling
	/*
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	*/

	// Set current user
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sharedCode.CurrentUserId = strings.ReplaceAll(currentUser.Username, "\\", "")

	fenixGuiBuilderServerMain()
}

/*
// SysTray Application - StartUp
func onReady() {

	systray.SetIcon(embededfenixIcon)
	systray.SetTitle("Fenix-GUI REST -> gRPC Proxy")
	systray.SetTooltip("Fenix-GUI REST -> gRPC Proxy")
	mQuit := systray.AddMenuItem("Quit", "Quit the Fenix-GUI REST -> gRPC Proxy")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)

	// Run menu handles as go-routine
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

}


*/
// SysTray Application - On exit
func onExit() {
	// clean up here, and exit the program
	os.Exit(0)

}

/*
// Convert variables that can be injected at build time into Map, to be able to be dynamically used when getting Environment variables
func convertBuildInjectedVariablesToMapStructure() {

	//GCP
	buildVariablesMap["BuildVariableGCPAuthentication"] = BuildVariableGCPAuthentication

	// GUI-ExecutionServer
	buildVariablesMap["BuildVariableExecutionLocationForFenixGuiExecutionServer"] = BuildVariableExecutionLocationForFenixGuiExecutionServer
	buildVariablesMap["BuildVariableFenixGuiExecutionServerAddress"] = BuildVariableFenixGuiExecutionServerAddress
	buildVariablesMap["BuildVariableFenixGuiExecutionServerPort"] = BuildVariableFenixGuiExecutionServerPort

	// GUI-TestCaseBuilderServer
	buildVariablesMap["BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer"] = BuildVariableExecutionLocationForFenixGuiTestCaseBuilderServer
	buildVariablesMap["BuildVariableFenixGuiTestCaseBuilderServerAddress"] = BuildVariableFenixGuiTestCaseBuilderServerAddress
	buildVariablesMap["BuildVariableFenixGuiTestCaseBuilderServerPort"] = BuildVariableFenixGuiTestCaseBuilderServerPort

	// This Application
	buildVariablesMap["BuildVariableExecutionLocationForThisApplication"] = BuildVariableExecutionLocationForThisApplication
	buildVariablesMap["BuildVariableFYNE_SCALE"] = BuildVariableFYNE_SCALE
	buildVariablesMap["BuildVariableRunAsTrayApplication"] = BuildVariableRunAsTrayApplication
	buildVariablesMap["BuildVariableApplicationGrpcPort"] = BuildVariableApplicationGrpcPort

}
*/
