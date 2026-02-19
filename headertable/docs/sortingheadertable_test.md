# sortingheadertable_test.go

## File Overview
- Path: `headertable/sortingheadertable_test.go`
- Package: `headertable`
- Generated: `2026-02-19T14:23:17+01:00`
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
- External calls: `assert.Equal`, `m.Refresh`, `test.Tap`

### Test_stringSort
- Signature: `func Test_stringSort(t *testing.T)`
- Exported: `true`
- Control-flow features: `for/range`
- Internal calls: `sortFn`, `stringSort`
- External calls: `assert.Equal`, `assert.NoError`, `binding.BindStruct`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
