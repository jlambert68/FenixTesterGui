# testCaseExecutionUI_helpers.go

## File Overview
- Path: `executions/executionsUIForSubscriptions/testCaseExecutionUI_helpers.go`
- Package: `executionsUIForSubscriptions`
- Functions/Methods: `3`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ResizeTableColumns`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/executions/executionsModelForSubscriptions`
- `FenixTesterGui/headertable`
- `errors`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/widget`
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- `headerColumnExtraWidth`

## Declared Variables
- None

## Functions and Methods
### ResizeTableColumns
- Signature: `func ResizeTableColumns(t *headertable.SortingHeaderTable)`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `float32`
- Selector calls: `b1.GetItem`, `fyne.NewSize`, `widget.NewLabel`

### remove
- Signature: `func remove(slice []binding.DataMap, s int) []binding.DataMap`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Remove item from the DataItem-slice and keep order

### verifyThatTestCaseExecutionIsNotInUse
- Signature: `func verifyThatTestCaseExecutionIsNotInUse(subscriptionsForTestCaseExecutionMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapOverallType) err error`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Selector calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
