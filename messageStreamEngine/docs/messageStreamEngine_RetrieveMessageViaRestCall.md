# messageStreamEngine_RetrieveMessageViaRestCall.go

## File Overview
- Path: `messageStreamEngine/messageStreamEngine_RetrieveMessageViaRestCall.go`
- Package: `messageStreamEngine`
- Functions/Methods: `2`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/common_code`
- `bytes`
- `encoding/json`
- `errors`
- `fmt`
- `github.com/sirupsen/logrus`
- `io/ioutil`
- `net/http`

## Declared Types
- `ackRequest`
- `pullRequest`
- `pullResponse`

## Declared Constants
- `googlePubSubPullAckURL`
- `googlePubSubPullURL`
- `numberOfMessagesToBePulled`

## Declared Variables
- None

## Functions and Methods
### retrievePubSubMessagesViaRestApi
- Signature: `func retrievePubSubMessagesViaRestApi(subscriptionID string, oauth2Token string) (numberOfMessagesInPullResponse int, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, go, defer, returns error`
- Doc: Pull a maximum of 'numberOfMessagesToBePulled' from PubSub-subscription
- Internal calls: `sendAcknowledgeMessageViaRestApi`, `string`, `triggerProcessPubSubExecutionStatusMessage`
- Selector calls: `bytes.NewBuffer`, `client.Do`, `decoder.Decode`, `err.Error`, `errors.New`, `fmt.Println`, `fmt.Sprintf`, `http.NewRequest`

### sendAcknowledgeMessageViaRestApi
- Signature: `func sendAcknowledgeMessageViaRestApi(projectID string, subscriptionID string, ackID string, oauth2Token string) error`
- Exported: `false`
- Control-flow features: `if, defer, returns error`
- Doc: Send Acknowledge for one message, which was Pulled and execution was successful
- Internal calls: `string`
- Selector calls: `bytes.NewBuffer`, `client.Do`, `fmt.Errorf`, `fmt.Sprintf`, `http.NewRequest`, `ioutil.ReadAll`, `json.Marshal`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
