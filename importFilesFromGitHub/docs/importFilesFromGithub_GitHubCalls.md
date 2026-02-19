# importFilesFromGithub_GitHubCalls.go

## File Overview
- Path: `importFilesFromGitHub/importFilesFromGithub_GitHubCalls.go`
- Package: `importFilesFromGitHub`
- Functions/Methods: `2`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `encoding/json`
- `fmt`
- `io/ioutil`
- `log`
- `net/http`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### getFileListFromGitHub (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) getFileListFromGitHub(apiUrl string)`
- Exported: `false`
- Control-flow features: `if, defer`
- Doc: List files and folders for a certain GitHub url
- Selector calls: `client.Do`, `err.Error`, `http.NewRequest`, `ioutil.ReadAll`, `json.Unmarshal`, `log.Fatalf`

### loadFileContent (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) loadFileContent(file GitHubFile) ([]byte, error)`
- Exported: `false`
- Control-flow features: `if, defer, returns error`
- Doc: Load the files content from GitHub
- Selector calls: `fmt.Errorf`, `http.Get`, `ioutil.ReadAll`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
