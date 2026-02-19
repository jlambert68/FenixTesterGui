# testSuiteUI_graphicalComponent_owerDomainForTestSuite.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_graphicalComponent_owerDomainForTestSuite.go`
- Package: `testSuiteUI`
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
### generateOwnerDomainForTestSuiteArea (method on `*TestSuiteUiStruct`)
- Signature: `func (*TestSuiteUiStruct) generateOwnerDomainForTestSuiteArea(testCasesModel *testCaseModel.TestCasesModelsStruct) (testCaseOwnerDomainContainer *fyne.Container, testCaseOwnerDomainCustomSelectComboBox *customSelectComboBox, err error)`
- Exported: `false`
- Control-flow features: `if, for/range, returns error`
- Doc: Generate the OwnerDomain Area for the TestSuite
- Internal calls: `newCustomSelectComboBoxWidget`
- Selector calls: `container.New`, `layout.NewVBoxLayout`, `layout.NewFormLayout`, `canvas.NewRectangle`, `widget.NewSelect`, `testSuiteUiModel.lockUIUntilOwnerDomainAndTestEnvironmenIsSelected`, `testSuiteUiModel.generateTestEnvironmentForTestSuite`, `fmt.Sprintf`

### setSelectedOwnerDomainForTestSuiteArea (method on `*TestSuiteUiStruct`)
- Signature: `func (*TestSuiteUiStruct) setSelectedOwnerDomainForTestSuiteArea(tempCurrentOwnerDomainToBeChosenInDropDown string, newOwnerDomainSelect *widget.Select, valueIsValidWarningBox *canvas.Rectangle)`
- Exported: `false`
- Control-flow features: `if`
- Doc: Sets the Selected value for the DropDown specifying the Owner-Domain of the TestSuite
- Selector calls: `newOwnerDomainSelect.SetSelected`

### calcSelectWidth
- Signature: `func calcSelectWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: calcSelectWidth returns the width needed to show the longest option
- Internal calls: `float32`
- Selector calls: `widget.NewSelect`, `tmp.Refresh`, `tmp.MinSize`

### calcCheckGroupWidth
- Signature: `func calcCheckGroupWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: calcCheckGroupWidth returns the width needed to show the widest checkbox label
- Selector calls: `widget.NewCheckGroup`, `tmp.Refresh`, `tmp.MinSize`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
