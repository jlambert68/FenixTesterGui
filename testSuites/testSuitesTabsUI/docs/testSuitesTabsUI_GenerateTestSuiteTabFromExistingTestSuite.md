# testSuitesTabsUI_GenerateTestSuiteTabFromExistingTestSuite.go

## File Overview
- Path: `testSuites/testSuitesTabsUI/testSuitesTabsUI_GenerateTestSuiteTabFromExistingTestSuite.go`
- Package: `testSuitesTabsUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateTestSuiteTabFromExistingTestSuite`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/testSuiteUI`
- `FenixTesterGui/testSuites/testSuitesCommandEngine`
- `FenixTesterGui/testSuites/testSuitesModel`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### GenerateTestSuiteTabFromExistingTestSuite
- Signature: `func GenerateTestSuiteTabFromExistingTestSuite(testSuiteUuidToOpen string, testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `true`
- Control-flow features: `if, go`
- External calls: `container.NewBorder`, `container.NewTabItem`, `container.NewVBox`, `err.Error`, `fmt.Sprintf`, `fyne.CurrentApp`, `fyne.Do`, `newTestSuiteModel.ExecuteOneTestSuiteWithAllItsTestDataSets`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
