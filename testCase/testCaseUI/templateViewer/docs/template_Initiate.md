# template_Initiate.go

## File Overview
- Path: `testCase/testCaseUI/templateViewer/template_Initiate.go`
- Package: `templateViewer`
- Functions/Methods: `2`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateTemplateViewer`

## Imports
- `FenixTesterGui/importFilesFromGitHub`
- `FenixTesterGui/soundEngine`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/placeholderReplacementEngine`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### InitiateTemplateViewer
- Signature: `func InitiateTemplateViewer(mainWindow fyne.Window, myApp fyne.App, importedFilesPtr *[]importFilesFromGitHub.GitHubFile, testDataForGroupObject *testDataEngine.TestDataForGroupObjectStruct, randomUuidForScriptEngine string, choseTemplateName string, testDataPointGroupName string, testDataPointName string, testDataRowName string)`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `getTestGroupsFromTestDataEngineFunction`, `testDataPointsToStringSliceFunction`, `testDataRowSliceToStringSliceFunction`, `getTextFromRichText`
- Selector calls: `mainWindow.Hide`, `myApp.NewWindow`, `templateViewerWindow.Resize`, `fyne.NewSize`, `templateViewerWindow.CenterOnScreen`, `widget.NewSelect`, `widget.NewLabel`, `testDataForGroupObject.GetTestDataPointValuesMapBasedOnGroupPointNameAndSummaryValue`

### getTextFromRichText
- Signature: `func getTextFromRichText(richText *widget.RichText) string`
- Exported: `false`
- Control-flow features: `for/range, switch`
- Doc: Function to extract text from RichText
- Selector calls: `sb.WriteString`, `sb.String`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
