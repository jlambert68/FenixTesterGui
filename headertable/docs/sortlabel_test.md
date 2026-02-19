# sortlabel_test.go

## File Overview
- Path: `headertable/sortlabel_test.go`
- Package: `headertable`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `4`
- Imports: `5`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TestNewSortingLabel`
- `TestSortingLabel_OnTapped_CyclesSortStates`
- `TestSortingLabel_SetState`
- `TestSortingLabel_nextState`

## Imports
- `fyne.io/fyne/v2/test`
- `github.com/PaulWaldo/fyne-headertable/headertable/data`
- `github.com/stretchr/testify/assert`
- `reflect`
- `testing`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### TestNewSortingLabel
- Signature: `func TestNewSortingLabel(t *testing.T)`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `NewSortingLabel`
- External calls: `reflect.DeepEqual`, `t.Errorf`, `t.Run`

### TestSortingLabel_OnTapped_CyclesSortStates
- Signature: `func TestSortingLabel_OnTapped_CyclesSortStates(t *testing.T)`
- Exported: `true`
- Control-flow features: `none detected`
- Internal calls: `NewSortingLabel`
- External calls: `assert.Equal`, `test.Tap`

### TestSortingLabel_SetState
- Signature: `func TestSortingLabel_SetState(t *testing.T)`
- Exported: `true`
- Control-flow features: `for/range`
- Internal calls: `NewSortingLabel`
- External calls: `assert.Equal`, `s.SetState`, `t.Run`

### TestSortingLabel_nextState
- Signature: `func TestSortingLabel_nextState(t *testing.T)`
- Exported: `true`
- Control-flow features: `if, for/range`
- External calls: `reflect.DeepEqual`, `t.Errorf`, `t.Run`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
