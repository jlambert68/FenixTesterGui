# testSuiteUI_graphicalComponent_owerDomainForTestSuite.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_graphicalComponent_owerDomainForTestSuite.go`
- Package: `testSuiteUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `4`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/testSuitesCommandEngine`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
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
### calcCheckGroupWidth
- Signature: `func calcCheckGroupWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: calcCheckGroupWidth returns the width needed to show the widest checkbox label
- External calls: `tmp.MinSize`, `tmp.Refresh`, `widget.NewCheckGroup`

### calcSelectWidth
- Signature: `func calcSelectWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: calcSelectWidth returns the width needed to show the longest option
- Internal calls: `float32`
- External calls: `tmp.MinSize`, `tmp.Refresh`, `widget.NewSelect`

### generateOwnerDomainForTestSuiteArea (method on `*TestSuiteUiStruct`)
- Signature: `func (*TestSuiteUiStruct) generateOwnerDomainForTestSuiteArea(testCasesModel *testCaseModel.TestCasesModelsStruct) (testCaseOwnerDomainContainer *fyne.Container, testCaseOwnerDomainCustomSelectComboBox *customSelectComboBox, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generate the OwnerDomain Area for the TestSuite
- Internal calls: `newCustomSelectComboBoxWidget`
- External calls: `canvas.NewRectangle`, `container.New`, `container.NewStack`, `container.NewVBox`, `err.Error`, `fmt.Sprintf`, `fyne.Do`, `layout.NewFormLayout`

### setSelectedOwnerDomainForTestSuiteArea (method on `*TestSuiteUiStruct`)
- Signature: `func (*TestSuiteUiStruct) setSelectedOwnerDomainForTestSuiteArea(tempCurrentOwnerDomainToBeChosenInDropDown string, newOwnerDomainSelect *widget.Select, valueIsValidWarningBox *canvas.Rectangle)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Sets the Selected value for the DropDown specifying the Owner-Domain of the TestSuite
- External calls: `newOwnerDomainSelect.SetSelected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
