# gcp_helpers.go

## File Overview
- Path: `gcp/gcp_helpers.go`
- Package: `gcp`
- Functions/Methods: `12`
- Imports: `21`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateGCPAccessToken`
- `GenerateGCPAccessTokenForAuthorizedUser`
- `GenerateGCPAccessTokenForAuthorizedUserPubSub`
- `GenerateGCPAccessTokenForServiceAccount`
- `GetGcpAccessTokenForAuthorizedAccountsPubSub`
- `SetLogger`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/notToBeSentToGithub`
- `bytes`
- `encoding/json`
- `fmt`
- `github.com/gorilla/pat`
- `github.com/gorilla/sessions`
- `github.com/markbates/goth`
- `github.com/markbates/goth/gothic`
- `github.com/markbates/goth/providers/google`
- `github.com/sirupsen/logrus`
- `github.com/toqueteos/webbrowser`
- `golang.org/x/net/context`
- `google.golang.org/api/idtoken`
- `google.golang.org/grpc/metadata`
- `html/template`
- `net/http`
- `os/exec`
- `runtime`
- `sync`
- `time`

## Declared Types
- `RefreshTokenResponse`

## Declared Constants
- None

## Declared Variables
- `DoneChannel`

## Functions and Methods
### GenerateGCPAccessToken (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) GenerateGCPAccessToken(ctx context.Context, targetServer TargetServerType) (appendedCtx context.Context, returnAckNack bool, returnMessage string)`
- Exported: `true`
- Control-flow features: `if, switch`
- Selector calls: `gcp.GenerateGCPAccessTokenForServiceAccount`, `gcp.GenerateGCPAccessTokenForAuthorizedUserPubSub`

### GenerateGCPAccessTokenForServiceAccount (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) GenerateGCPAccessTokenForServiceAccount(ctx context.Context, targetServer TargetServerType) (appendedCtx context.Context, returnAckNack bool, returnMessage string)`
- Exported: `true`
- Control-flow features: `if, switch`
- Doc: GenerateGCPAccessTokenForServiceAccount Generate Google access token for a service account. Used when running in GCP
- Selector calls: `timeToCompareTo.After`, `time.Now`, `idtoken.NewTokenSource`, `idtoken.WithCredentialsJSON`, `tokenSource.Token`, `grpcMetadata.AppendToOutgoingContext`

### GenerateGCPAccessTokenForAuthorizedUser (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) GenerateGCPAccessTokenForAuthorizedUser(ctx context.Context) (appendedCtx context.Context, returnAckNack bool, returnMessage string)`
- Exported: `true`
- Control-flow features: `if, go, defer`
- Selector calls: `gcp.initiateUserObject`, `timeToCompareTo.After`, `time.Now`, `grpcMetadata.AppendToOutgoingContext`, `sessions.NewCookieStore`, `store.MaxAge`, `goth.UseProviders`, `google.New`

### GenerateGCPAccessTokenForAuthorizedUserPubSub (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) GenerateGCPAccessTokenForAuthorizedUserPubSub(ctx context.Context) (appendedCtx context.Context, returnAckNack bool, returnMessage string)`
- Exported: `true`
- Control-flow features: `if, go`
- Internal calls: `refreshToken`
- Selector calls: `gcp.initiatAccessTokenForAuthorizedAccountsPubSubObject`, `pat.New`, `timeToCompareTo.After`, `time.Now`, `grpcMetadata.AppendToOutgoingContext`, `timeToCompareTo.Before`, `fmt.Println`, `err.Error`

### GetGcpAccessTokenForAuthorizedAccountsPubSub (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) GetGcpAccessTokenForAuthorizedAccountsPubSub() string`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GetGcpAccessTokenForAuthorizedAccountsPubSub Get Access token to be used for contacting PubSub

### refreshToken
- Signature: `func refreshToken(client *http.Client, refreshToken string) (*RefreshTokenResponse, error)`
- Exported: `false`
- Control-flow features: `if, defer, returns error`
- Selector calls: `json.Marshal`, `http.Post`, `bytes.NewBuffer`, `fmt.Println`, `json.NewDecoder`, `time.Duration`, `time.Now`

### startLocalWebServer (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) startLocalWebServer(webServer *http.Server)`
- Exported: `false`
- Control-flow features: `if, go`
- Doc: Start and run Local Web Server
- Selector calls: `time.Sleep`, `webbrowser.Open`, `webServer.ListenAndServe`

### startLocalWebServerExpanded (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) startLocalWebServerExpanded(webServer *http.Server, url string)`
- Exported: `false`
- Control-flow features: `if, switch, defer`
- Doc: Start and run Local Web Server
- Selector calls: `exec.Command`, `cmd.Start`, `fmt.Printf`, `fmt.Println`, `webServer.ListenAndServe`, `err.Error`

### stopLocalWebServer (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) stopLocalWebServer(ctx context.Context, webServer *http.Server)`
- Exported: `false`
- Control-flow features: `if, defer`
- Doc: Close down Local Web Server
- Internal calls: `cancel`
- Selector calls: `context.WithTimeout`, `webServer.Shutdown`

### SetLogger (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same Logger reference as is used by central part of system

### initiateUserObject (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) initiateUserObject()`
- Exported: `false`
- Control-flow features: `if`
- Doc: initiateUserObject

### initiatAccessTokenForAuthorizedAccountsPubSubObject (method on `*GcpObjectStruct`)
- Signature: `func (*GcpObjectStruct) initiatAccessTokenForAuthorizedAccountsPubSubObject()`
- Exported: `false`
- Control-flow features: `if`
- Doc: initiatAccessTokenForAuthorizedAccountsPubSubObject

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
