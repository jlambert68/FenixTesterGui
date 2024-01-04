package gcp

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/notToBeSentToGithub"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/sirupsen/logrus"
	"github.com/toqueteos/webbrowser"
	"golang.org/x/net/context"
	"google.golang.org/api/idtoken"
	grpcMetadata "google.golang.org/grpc/metadata"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func (gcp *GcpObjectStruct) GenerateGCPAccessToken(ctx context.Context, targetServer TargetServerType) (
	appendedCtx context.Context, returnAckNack bool, returnMessage string) {

	// Chose correct Authentication method
	switch targetServer {
	/*
		case TargetServerGuiTestCaseBuilderServer:
			// Chose correct method for authentication
			if sharedCode.UseServiceAccountForGuiTestCaseBuilderServer == true {
				// Use Service account

				appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForServiceAccount(ctx, targetServer)
			} else {
				// User log into GCP via web
				appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
			}

				case TargetServerGuiExecutionServer:
					// Chose correct method for authentication
					if sharedCode.UseServiceAccountForGuiExecutionServer == true {
						// Use Service account

						appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForServiceAccount(ctx, targetServer)
					} else {
						// User log into GCP via web
						appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForAuthorizedUser(ctx)
					}


	*/

	case TargetServerGuiTestCaseBuilderServer:
		// Chose correct method for authentication
		if sharedCode.UseServiceAccountForGuiTestCaseBuilderServer == true {
			// Use Service account

			appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForServiceAccount(ctx, targetServer)
		} else {
			// User log into GCP via web
			appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForAuthorizedUserPubSub(ctx)
		}

	case TargetServerGuiExecutionServer:
		// Chose correct method for authentication
		if sharedCode.UseServiceAccountForGuiExecutionServer == true {
			// Use Service account

			appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForServiceAccount(ctx, targetServer)
		} else {
			// User log into GCP via web
			appendedCtx, returnAckNack, returnMessage = gcp.GenerateGCPAccessTokenForAuthorizedUserPubSub(ctx)
		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":           "6427a019-047d-4663-ae17-86446485ccc9",
			"targetServer": targetServer,
		}).Error("Unknown TargetServer")

		return nil, false, "Unknown TargetServer"

	}

	return appendedCtx, returnAckNack, returnMessage

}

// GenerateGCPAccessTokenForServiceAccount Generate Google access token for a service account. Used when running in GCP
func (gcp *GcpObjectStruct) GenerateGCPAccessTokenForServiceAccount(ctx context.Context, targetServer TargetServerType) (appendedCtx context.Context, returnAckNack bool, returnMessage string) {

	// Only create the token if there is none, or it has expired (or 5 minutes before expiration
	var safetyDuration time.Duration
	safetyDuration = 5 * time.Minute
	timeToCompareTo := gcp.gcpAccessTokenForServiceAccounts.Expiry.Add(safetyDuration)
	if gcp.gcpAccessTokenForServiceAccounts == nil || timeToCompareTo.After(time.Now()) {

		var gcpScope string
		var serviceAccountKeyJson []byte

		// Chose correct GCP-scope and Service Account-data
		switch targetServer {
		case TargetServerGuiTestCaseBuilderServer:
			gcpScope = notToBeSentToGithub.Gcp_scope_GuiTestCaseBuilderServer
			serviceAccountKeyJson = notToBeSentToGithub.ServiceAccountKeyJson_GuiTestCaseBuilderServer

		case TargetServerGuiExecutionServer:
			gcpScope = notToBeSentToGithub.Gcp_scope_GuiExecutionServer
			serviceAccountKeyJson = notToBeSentToGithub.ServiceAccountKeyJson_GuiExecutionServer

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":           "49781e38-0cf6-4b4f-9210-0257feea06a5",
				"targetServer": targetServer,
			}).Error("Unknown TargetServer")

			return nil, false, "Unknown TargetServer"

		}

		tokenSource, err := idtoken.NewTokenSource(ctx, gcpScope, idtoken.WithCredentialsJSON(serviceAccountKeyJson))

		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":  "c11aba6d-a177-4b88-a521-15bec26a3312",
				"err": err,
			}).Error("Couldn't generate access token")

			return nil, false, "Couldn't generate access token"
		}

		token, err := tokenSource.Token()
		//token, err := config.TokenSource(oauth2.NoContext).Token()

		if err != nil {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":  "0e0dd854-2088-419d-a0d1-035b6242c585",
				"err": err,
			}).Error("Problem getting the token")

			return nil, false, "Problem getting the token"
		} else {
			sharedCode.Logger.WithFields(logrus.Fields{
				"ID":    "1ee9fd6d-c83d-4dbb-bce7-fd4901ff3f87",
				"token": "Nothing to see", //token,
			}).Debug("Got Bearer Token")
		}

		gcp.gcpAccessTokenForServiceAccounts = token

	}

	sharedCode.Logger.WithFields(logrus.Fields{
		"ID": "fa86e0bd-9c47-4f1e-b9dd-2b1f08880b2b",
		"fenixGuiBuilderProxyServerObject.gcpAccessTokenForServiceAccounts": "Nothing to see", //fenixGuiBuilderProxyServerObject.gcpAccessTokenForServiceAccounts,
	}).Debug("Will use Bearer Token")

	// Add token to gRPC Request.
	appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.gcpAccessTokenForServiceAccounts.AccessToken)

	return appendedCtx, true, ""

}

// DoneChannel - channel used for to close down local web server
var DoneChannel chan bool

func (gcp *GcpObjectStruct) GenerateGCPAccessTokenForAuthorizedUser(ctx context.Context) (appendedCtx context.Context, returnAckNack bool, returnMessage string) {

	sharedCode.Logger.WithFields(logrus.Fields{
		"id": "df375b80-449b-4b6f-96ba-1600146f8860",
	}).Debug("Incoming 'GenerateGCPAccessTokenForAuthorizedUser'")

	defer sharedCode.Logger.WithFields(logrus.Fields{
		"id": "623fffe0-ff28-42ef-b8ea-8b2875edc32f",
	}).Debug("Outgoing 'GenerateGCPAccessTokenForAuthorizedUser'")

	// Secure that User is initiated
	gcp.initiateUserObject()

	// Only aloud one parallel instance of the code below to be run
	gcp.mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts.Lock()
	defer gcp.mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts.Unlock()

	// Only create the token if there is none, or it has expired (or 5 minutes before expiration
	var safetyDuration time.Duration
	safetyDuration = 5 * time.Minute
	timeToCompareTo := gcp.gcpAccessTokenForAuthorizedAccounts.ExpiresAt.Add(safetyDuration)
	if gcp.gcpAccessTokenForAuthorizedAccounts.IDToken != "" && timeToCompareTo.After(time.Now()) {
		// We already have a ID-token that can be used, so return that
		appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.gcpAccessTokenForAuthorizedAccounts.IDToken)

		return appendedCtx, true, ""
	}

	// Need to create a new ID-token

	key := sharedCode.ApplicationRunTimeUuid //"Secret-session-keyxyz" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                     // 30 days
	isProd := false                          // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		// Use 'Fenix End User Authentication'
		google.New(
			sharedCode.AuthClientId,
			sharedCode.AuthClientSecret,
			"http://localhost:3000/auth/google/callback",
			"email", "profile", "https://www.googleapis.com/auth/pubsub"),
	)

	router := pat.New()

	router.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {

			fmt.Fprintln(res, err)

			return
		}
		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(res, user)

		// Save ID-token
		gcp.gcpAccessTokenForAuthorizedAccounts = user

		// Trigger Close of Web Server, and 'true' means that a ID-to
		DoneChannel <- true

	})

	router.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	router.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)

	})

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(res, false)

	})

	// Initiate channel used to stop server
	DoneChannel = make(chan bool, 1)

	// Initiate http server
	localWebServer := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// Start Local Web Server as go routine
	go gcp.startLocalWebServer(localWebServer)

	gcp.logger.WithFields(logrus.Fields{
		"ID": "b0f86842-c0f7-4473-be2e-8e5e21ae82f3",
	}).Debug("Local webServer Started")

	// Wait for message in channel to stop local web server
	gotIdTokenResult := <-DoneChannel

	// Shutdown local web server before leaving
	gcp.stopLocalWebServer(context.Background(), localWebServer)

	// Depending on the outcome of getting a token return different results
	if gotIdTokenResult == true {
		// Success in getting an ID-token
		appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.gcpAccessTokenForAuthorizedAccounts.IDToken)

		return appendedCtx, true, ""
	} else {
		// Didn't get any ID-token
		return nil, false, "Couldn't generate access token"
	}

}

func (gcp *GcpObjectStruct) GenerateGCPAccessTokenForAuthorizedUserPubSub(ctx context.Context) (appendedCtx context.Context, returnAckNack bool, returnMessage string) {

	// Secure that User is initiated
	gcp.initiatAccessTokenForAuthorizedAccountsPubSubObject()

	router := pat.New()
	var url string

	// Initiate http server
	localWebServer := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// Only create the token if there is none, or it has expired (or 5 minutes before expiration
	var safetyDuration time.Duration
	var timeToCompareTo time.Time
	safetyDuration = -5 * time.Minute
	if gcp.gcpAccessTokenForAuthorizedAccountsPubSub.IDToken != "" {

		timeToCompareTo = gcp.refreshTokenResponse.ExpiresAt.Add(safetyDuration)
	}
	if gcp.gcpAccessTokenForAuthorizedAccountsPubSub.IDToken != "" && timeToCompareTo.After(time.Now()) {
		// We already have a ID-token that can be used, so return that
		appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.refreshTokenResponse.IDToken)

		return appendedCtx, true, ""

	} else if gcp.gcpAccessTokenForAuthorizedAccountsPubSub.IDToken != "" && timeToCompareTo.Before(time.Now()) {
		client := &http.Client{
			// Configure the client if necessary. For example, set a timeout:
			Timeout: time.Second * 30,
		}

		refreshTokenResponseMessage, err := refreshToken(client, gcp.gcpAccessTokenForAuthorizedAccountsPubSub.RefreshToken)
		if err != nil {
			fmt.Println("err: ", err)

			return nil, false, err.Error()

		} else {

			// When no refresh token was received then ask user to close the web browser containing previous log in credentials
			if gcp.gcpAccessTokenForAuthorizedAccountsPubSub.RefreshToken == "" {
				url = "http://localhost:3000/close_this_browser"
				go gcp.startLocalWebServerExpanded(localWebServer, url)

				return nil, false, "Missing Refresh token"
			}

			// Store Refresh response
			gcp.refreshTokenResponse = refreshTokenResponseMessage

			//
			appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.refreshTokenResponse.IDToken)
			return appendedCtx, true, ""
		}

	}

	// Need to create a new ID-token

	key := sharedCode.ApplicationRunTimeUuid // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                     // 30 days
	isProd := false                          // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		// Use 'Fenix End User Authentication'

		google.New(
			sharedCode.AuthClientId,
			sharedCode.AuthClientSecret,
			"http://localhost:3000/auth/google/callback",
			"email", "profile", "https://www.googleapis.com/auth/pubsub"),
	)

	router.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {

			fmt.Fprintln(res, err)

			return
		}

		// Save ID-token
		gcp.gcpAccessTokenForAuthorizedAccountsPubSub = user

		// When we got an Refresh Token then inform of Success
		// When there was no Refresh Token then inform user to close Browser and restart
		if len(user.RefreshToken) > 0 {
			// Got Refresh Token

			// Store who is authenticated towards GCP
			sharedCode.CurrentUserAuthenticatedTowardsGCP = user.Email

			t, _ := template.ParseFiles("templates/success.html")
			t.Execute(res, user)

			// Trigger Close of Web Server, and 'true' means that a ID-to
			DoneChannel <- true

		} else {
			// Didn't get Refresh Token
			t, _ := template.ParseFiles("templates/close_this_browser.html")
			t.Execute(res, false)

			// Trigger Close of Web Server, and 'false' means no Refresh Token
			DoneChannel <- false

		}

	})

	router.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)

	})

	//
	router.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	// Start page for web server for user to be able to login into GCP
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		//res.Header().Set("state-token", "offline")
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(res, false)
	})

	// Show Text telling user to close down web browser due to that no Refresh Token can be retrieved
	// as long as browser window is open
	/*
		router.Get("/closethisbrowser", func(res http.ResponseWriter, req *http.Request) {
			//res.Header().Set("state-token", "offline")
			t, _ := template.ParseFiles("templates/close_this_browser.html")
			t.Execute(res, false)

			// Shutdown local web server
			gcp.stopLocalWebServer(context.Background(), localWebServer)

		})

	*/

	// Initiate channel used to stop server
	DoneChannel = make(chan bool, 1)

	// Start Local Web Server as go routine
	url = "http://localhost:3000"
	go gcp.startLocalWebServerExpanded(localWebServer, url)

	sharedCode.Logger.WithFields(logrus.Fields{
		"ID": "689d42de-3cc0-4237-b1e9-3a6c769f65ea",
	}).Debug("Local webServer Started")

	// Wait for message in channel to stop local web server
	gotIdTokenResult := <-DoneChannel

	// Shutdown local web server
	gcp.stopLocalWebServer(context.Background(), localWebServer)

	// Depending on the outcome of getting a token return different results
	if gotIdTokenResult == true {
		// Success in getting an ID-token first time so use RefreshToken to fill RefreshTokenMessage
		client := &http.Client{
			// Configure the client if necessary. For example, set a timeout:
			Timeout: time.Second * 30,
		}
		refreshTokenResponseMessage, err := refreshToken(client, gcp.gcpAccessTokenForAuthorizedAccountsPubSub.RefreshToken)
		if err != nil {
			fmt.Println("err: ", err)

			return nil, false, err.Error()

		} else {

			// When no refresh token was received then ask user to close the web browser containing previous log in credentials
			if gcp.gcpAccessTokenForAuthorizedAccountsPubSub.RefreshToken == "" {
				url = "http://localhost:3000/closethisbrowser"
				gcp.startLocalWebServerExpanded(localWebServer, url)

				time.Sleep(10 * time.Second)

				return nil, false, "Missing Refresh token"
			}

			// Store Refresh response
			gcp.refreshTokenResponse = refreshTokenResponseMessage

			//
			appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.refreshTokenResponse.IDToken)
			return appendedCtx, true, ""
		}

		appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+gcp.gcpAccessTokenForAuthorizedAccountsPubSub.IDToken)

		return appendedCtx, true, ""
	} else {
		// Didn't get any ID-token
		return nil, false, "Couldn't generate access token"
	}

}

// GetGcpAccessTokenForAuthorizedAccountsPubSub
// Get Access token to be used for contacting PubSub
func (gcp *GcpObjectStruct) GetGcpAccessTokenForAuthorizedAccountsPubSub() string {
	return gcp.refreshTokenResponse.AccessToken
	//return gcp.gcpAccessTokenForAuthorizedAccountsPubSub.AccessToken
}

// RefreshTokenResponse represents the JSON response from the OAuth2 provider.
type RefreshTokenResponse struct {
	AccessToken  string    `json:"access_token"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
	//Scope        string `json:"scope"`
	TokenType string `json:"token_type"`
	IDToken   string `json:"id_token"`
	// Include other fields as necessary
}

func refreshToken(client *http.Client, refreshToken string) (*RefreshTokenResponse, error) {
	// The URL for the token endpoint will vary based on the OAuth2 provider.
	tokenEndpoint := "https://oauth2.googleapis.com/token"

	requestData := map[string]string{
		"client_id":     sharedCode.AuthClientId,
		"client_secret": sharedCode.AuthClientSecret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}
	jsonValue, _ := json.Marshal(requestData)

	response, err := http.Post(tokenEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	//response, err := client.Post(tokenEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// Handle non-200 responses
		fmt.Println(response.StatusCode)
		return nil, err
	}

	var tokenResponse RefreshTokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, err
	}

	// Build time when Token expires
	var expireDuration time.Duration
	expireDuration = time.Duration(tokenResponse.ExpiresIn) * time.Second
	tokenResponse.ExpiresAt = time.Now().Add(expireDuration)

	return &tokenResponse, nil
}

// Start and run Local Web Server
func (gcp *GcpObjectStruct) startLocalWebServer(webServer *http.Server) {

	go func() {
		time.Sleep(1 * time.Second)
		err := webbrowser.Open("http://localhost:3000")

		if err != nil {
			gcp.logger.WithFields(logrus.Fields{
				"ID":  "17bc0305-4594-48e1-bb8d-c642579e5e56",
				"err": err,
			}).Fatalln("Couldn't open the web browser")
		}

	}()
	err := webServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		GcpObject.logger.WithFields(logrus.Fields{
			"ID": "8226cf74-0cdc-4e29-a441-116504b4b333",
		}).Fatalf("Local Web Server failed to listen: %s\n", err)

	}
}

// Start and run Local Web Server
func (gcp *GcpObjectStruct) startLocalWebServerExpanded(webServer *http.Server, url string) {

	var cmd *exec.Cmd

	/*
		switch runtime.GOOS {
		case "windows":
			// Command for Windows
			cmd = exec.Command("cmd", "/C", "start", "chrome", "--new-window", "--guest", url)
		case "darwin":
			// Command for macOS
			cmd = exec.Command("open", "-a", "Google Chrome", "--args", "--new-window", "--guest", url)
		case "linux":
			// Command for Linux
			cmd = exec.Command("google-chrome", "--new-window", "--guest", url)
		default:
			panic("Unsupported operating system")
		}

	*/
	// Determine the operating system
	switch runtime.GOOS {
	case "windows":
		// Command for Windows
		cmd = exec.Command("cmd", "/C", "start", "chrome", "--incognito", url)
	case "darwin":
		// Command for macOS
		cmd = exec.Command("open", "-a", "Google Chrome", "--args", "--incognito", url)
	case "linux":
		// Command for Linux
		cmd = exec.Command("google-chrome", "--incognito", url)
	default:
		panic("Unsupported operating system")
	}

	// Execute the command
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	// Print the PID of the process
	fmt.Printf("Chrome started with PID: %d\n", cmd.Process.Pid)

	// Kill the process
	//if err := cmd.Process.Kill(); err != nil {
	//	panic(err)
	//}
	//err := webbrowser.Open("http://localhost:3000")

	if err != nil {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "17bc0305-4594-48e1-bb8d-c642579e5e56",
			"err": err,
		}).Fatalf("Couldn't open the web browser")
	}

	// Kill the process before leave
	defer func() {
		if err := cmd.Process.Kill(); err != nil {
			panic(err)
		}

		fmt.Println("Chrome process killed")
	}()

	err = webServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID": "8226cf74-0cdc-4e29-a441-116504b4b333",
		}).Fatalf("Local Web Server failed to listen: %s\n", err)

	}

	sharedCode.Logger.WithFields(logrus.Fields{
		"ID":  "844f2c3e-c271-4f95-ba9c-4eec9a206811",
		"err": err.Error(),
	}).Debug("Web Server was stopped")
}

// Close down Local Web Server
func (gcp *GcpObjectStruct) stopLocalWebServer(
	ctx context.Context, webServer *http.Server) {

	gcp.logger.WithFields(logrus.Fields{
		"ID": "1f4e0354-2a09-4a1d-be61-67ecda781142",
	}).Debug("Trying to stop local web server")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := webServer.Shutdown(ctx)
	if err != nil {
		if err != nil {
			gcp.logger.WithFields(logrus.Fields{
				"ID": "ea06dfab-39b9-4df6-b3ca-7f5f56b3cb91",
			}).Fatalf("Local Web Server Shutdown Failed:%+v", err)

		} else {
			gcp.logger.WithFields(logrus.Fields{
				"ID": "ea06dfab-39b9-4df6-b3ca-7f5f56b3cb91",
			}).Debug("Local Web Server Exited Properly")
		}

	}

}

// SetLogger
// Set to use the same Logger reference as is used by central part of system
func (gcp *GcpObjectStruct) SetLogger(logger *logrus.Logger) {

	//grpcOutVaraible = GRPCOutStruct{}

	gcp.logger = logger

	return

}

// initiateUserObject
func (gcp *GcpObjectStruct) initiateUserObject() {

	// Only do initiation if it's not done before

	if gcp.mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts == nil {
		gcp.gcpAccessTokenForAuthorizedAccounts = goth.User{}
		gcp.mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts = &sync.Mutex{}
	}

	return

}

// initiatAccessTokenForAuthorizedAccountsPubSubObject
func (gcp *GcpObjectStruct) initiatAccessTokenForAuthorizedAccountsPubSubObject() {

	// Only do initiation if it's not done before

	if gcp.mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts == nil {
		gcp.gcpAccessTokenForAuthorizedAccountsPubSub = goth.User{}
		gcp.mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts = &sync.Mutex{}
	}

	return

}
