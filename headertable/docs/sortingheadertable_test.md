# sortingheadertable_test.go

## File Overview
- Path: `headertable/sortingheadertable_test.go`
- Package: `headertable`
- Functions/Methods: `2`
- Imports: `6`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TestSortingLabel_OnTapped_CallsSorterAndUnsortsOthers`
- `Test_stringSort`

## Imports
- `fmt`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/test`
- `fyne.io/fyne/v2/widget`
- `github.com/stretchr/testify/assert`
- `testing`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### TestSortingLabel_OnTapped_CallsSorterAndUnsortsOthers
- Signature: `func TestSortingLabel_OnTapped_CallsSorterAndUnsortsOthers(t *testing.T)`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `NewSortingHeaderTable`
- Selector calls: `m.Refresh`, `test.Tap`, `assert.Equal`

### Test_stringSort
- Signature: `func Test_stringSort(t *testing.T)`
- Exported: `true`
- Control-flow features: `for/range`
- Internal calls: `stringSort`, `sortFn`
- Selector calls: `binding.BindStruct`, `assert.NoError`, `assert.Equal`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
