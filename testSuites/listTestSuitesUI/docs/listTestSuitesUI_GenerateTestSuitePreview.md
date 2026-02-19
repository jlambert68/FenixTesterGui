# listTestSuitesUI_GenerateTestSuitePreview.go

## File Overview
- Path: `testSuites/listTestSuitesUI/listTestSuitesUI_GenerateTestSuitePreview.go`
- Package: `listTestSuitesUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateTestSuitePreviewContainer`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testCases/listTestCasesUI`
- `FenixTesterGui/testSuites/listTestSuitesModel`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### GenerateTestSuitePreviewContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) GenerateTestSuitePreviewContainer(testSuiteUuid string, testCasesModel *testCaseModel.TestCasesModelsStruct)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: GenerateTestSuitePreviewContainer Generates the PreViewContainer for the TestSuite
- Internal calls: `float32`
- External calls: `container.New`, `container.NewBorder`, `container.NewCenter`, `container.NewScroll`, `container.NewVBox`, `fmt.Sprintf`, `fyne.NewSize`, `layout.NewFormLayout`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
