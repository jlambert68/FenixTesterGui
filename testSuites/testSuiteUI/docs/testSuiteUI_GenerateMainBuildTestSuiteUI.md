# testSuiteUI_GenerateMainBuildTestSuiteUI.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_GenerateMainBuildTestSuiteUI.go`
- Package: `testSuiteUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateBuildNewTestSuiteUI`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/testSuitesModel`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `image/color`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### GenerateBuildNewTestSuiteUI (method on `TestSuiteUiStruct`)
- Signature: `func (TestSuiteUiStruct) GenerateBuildNewTestSuiteUI(testCasesModel *testCaseModel.TestCasesModelsStruct, newTestSuiteModel *testSuitesModel.TestSuiteModelStruct) (newTestSuiteUIContainer *fyne.Container, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: GenerateBuildNewTestSuiteUI Create the UI used for creating new TestSuites
- Internal calls: `NewHoverableRect`, `generateRightSideBuildTestSuiteContainer`
- External calls: `container.New`, `container.NewBorder`, `container.NewHSplit`, `container.NewVBox`, `err.Error`, `layout.NewStackLayout`, `leftCreateTestSuiteOverlay.Hide`, `newTestSuiteModel.GetTestSuiteUuid`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
