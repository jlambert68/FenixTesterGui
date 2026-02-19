# testSuiteUI_GraphicalComponent_testSuiteDeletetionDate.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_GraphicalComponent_testSuiteDeletetionDate.go`
- Package: `testSuiteUI`
- Functions/Methods: `3`
- Imports: `12`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/soundEngine`
- `FenixTesterGui/testSuites/testSuitesModel`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `time`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `enableDeletionCheckbox`
- `newTestSuiteDeletionDateEntry`
- `tickerCountDownlabel`
- `tickerCountDownlabelDataBinding`
- `tickerDoneChannel`

## Functions and Methods
### countDownTicker
- Signature: `func countDownTicker(testSuiteUiModel *TestSuiteUiStruct)`
- Exported: `false`
- Control-flow features: `if, for/range, select, go, defer`
- Selector calls: `tickerCountDownlabelDataBinding.Set`, `fmt.Sprintf`, `fyne.Do`, `tickerCountDownlabel.Show`, `time.NewTicker`, `ticker.Stop`, `tickerCountDownlabel.Hide`, `enableDeletionCheckbox.SetChecked`

### flashScreen
- Signature: `func flashScreen(mainApp fyne.App, mainWindow fyne.Window)`
- Exported: `false`
- Control-flow features: `if, for/range, select, go, defer`
- Doc: Functions that hides the Fenix Screen and the flash the full screen
- Selector calls: `mainWindow.Hide`, `mainApp.NewWindow`, `redWindow.SetFullScreen`, `canvas.NewRectangle`, `fyne.NewContainerWithLayout`, `layout.NewMaxLayout`, `redWindow.SetContent`, `redWindow.Show`

### generateTestSuiteDeletionDateArea (method on `*TestSuiteUiStruct`)
- Signature: `func (*TestSuiteUiStruct) generateTestSuiteDeletionDateArea(testSuiteUuid string) (testSuiteDeletionDateAreaContainer *fyne.Container, err error)`
- Exported: `false`
- Control-flow features: `if, go, returns error`
- Doc: Generate the TestSuiteDeletionDate Area for the TestSuite
- Internal calls: `dateValidatorFunction`, `countDownTicker`
- Selector calls: `errors.New`, `fmt.Sprintf`, `fmt.Println`, `container.New`, `layout.NewVBoxLayout`, `layout.NewFormLayout`, `binding.NewString`, `widget.NewLabelWithData`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
