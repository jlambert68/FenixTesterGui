# testSuiteUI_GenerateLeftSideBuildTestSuite.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_GenerateLeftSideBuildTestSuite.go`
- Package: `testSuiteUI`
- Functions/Methods: `1`
- Imports: `7`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testCases/listTestCasesUI`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateLeftSideBuildTestSuiteContainer (method on `TestSuiteUiStruct`)
- Signature: `func (TestSuiteUiStruct) generateLeftSideBuildTestSuiteContainer(testSuiteUuid string, testCasesModel *testCaseModel.TestCasesModelsStruct, preViewAndFilterTabsUsedForCreateTestSuite *container.AppTabs) (leftSideBuildTestSuiteContainer *fyne.Container, err error)`
- Exported: `false`
- Control-flow features: `if, returns error`
- Doc: Generate leftSideBuildTestSuite - Main information for TestSuite
- Selector calls: `container.NewVBox`, `widget.NewLabel`, `container.NewStack`, `container.NewBorder`, `widget.NewAccordionItem`, `widget.NewAccordion`, `testSuiteUiModel.generateTestSuiteDeletionDateArea`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
