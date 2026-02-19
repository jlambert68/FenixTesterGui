# testCaseUI_newTextBoxAtributeEntry.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_newTextBoxAtributeEntry.go`
- Package: `testCaseUI`
- Functions/Methods: `3`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `FocusGained`
- `FocusLost`
- `NewAttributeEntry`

## Imports
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/widget`

## Declared Types
- `AttributeEntry`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### FocusGained (method on `*AttributeEntry`)
- Signature: `func (*AttributeEntry) FocusGained(x fyne.Focusable)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: FocusGained - Single Click on colorRectangle
- Selector calls: `fmt.Println`, `x.FocusGained`

### FocusLost (method on `*AttributeEntry`)
- Signature: `func (*AttributeEntry) FocusLost(x fyne.Focusable)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: FocusLost - Single Click on colorRectangle
- Selector calls: `fmt.Println`, `x.FocusLost`

### NewAttributeEntry (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) NewAttributeEntry(attributeUuid string) *AttributeEntry`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `myAttributeEntry.ExtendBaseWidget`, `widget.NewEntry`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
