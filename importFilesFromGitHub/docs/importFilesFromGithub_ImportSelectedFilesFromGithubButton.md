# importFilesFromGithub_ImportSelectedFilesFromGithubButton.go

## File Overview
- Path: `importFilesFromGitHub/importFilesFromGithub_ImportSelectedFilesFromGithubButton.go`
- Package: `importFilesFromGitHub`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `3`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `encoding/base64`
- `encoding/json`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/dialog`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixSyncShared`
- `log`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### decodeBase64Content (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) decodeBase64Content(encodedContent string) (string, error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Decode the file content from base64 to string
- Internal calls: `string`

### extractContentFromJson (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) extractContentFromJson(jsonData string) (string, error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Extra the file content from the json
- External calls: `json.Unmarshal`

### generateImportSelectedFilesFromGithubButton (method on `*ImportFilesFromGitHubStruct`)
- Signature: `func (*ImportFilesFromGitHubStruct) generateImportSelectedFilesFromGithubButton(parentWindow fyne.Window)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generate the button that imports the selected files from Github
- Internal calls: `string`
- External calls: `dialog.ShowError`, `fenixMainWindow.Show`, `fenixSyncShared.HashSingleValue`, `importFilesFromGitHubObject.decodeBase64Content`, `importFilesFromGitHubObject.extractContentFromJson`, `importFilesFromGitHubObject.loadFileContent`, `log.Fatalf`, `parentWindow.Close`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
