# testSuiteUI_GraphicalComponent_testSuiteDeletetionDate.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_GraphicalComponent_testSuiteDeletetionDate.go`
- Package: `testSuiteUI`
- Generated: `2026-02-19T14:23:17+01:00`
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
- External calls: `enableDeletionCheckbox.SetChecked`, `fmt.Sprintf`, `fyne.CurrentApp`, `fyne.Do`, `newTestSuiteDeletionDateEntry.SetText`, `ticker.Stop`, `tickerCountDownlabel.Hide`, `tickerCountDownlabel.Show`

### flashScreen
- Signature: `func flashScreen(mainApp fyne.App, mainWindow fyne.Window)`
- Exported: `false`
- Control-flow features: `if, for/range, select, go, defer`
- Doc: Functions that hides the Fenix Screen and the flash the full screen
- External calls: `canvas.NewRectangle`, `fmt.Println`, `fyne.CurrentApp`, `fyne.NewContainerWithLayout`, `layout.NewMaxLayout`, `mainApp.NewWindow`, `mainWindow.Hide`, `mainWindow.Show`

### generateTestSuiteDeletionDateArea (method on `*TestSuiteUiStruct`)
- Signature: `func (*TestSuiteUiStruct) generateTestSuiteDeletionDateArea(testSuiteUuid string) (testSuiteDeletionDateAreaContainer *fyne.Container, err error)`
- Exported: `false`
- Control-flow features: `if, go, returns error`
- Doc: Generate the TestSuiteDeletionDate Area for the TestSuite
- Internal calls: `countDownTicker`, `dateValidatorFunction`
- External calls: `binding.NewString`, `canvas.NewRectangle`, `container.New`, `container.NewBorder`, `container.NewHBox`, `container.NewVBox`, `deleteTestSuiteButton.Disable`, `deleteTestSuiteButton.Enable`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
