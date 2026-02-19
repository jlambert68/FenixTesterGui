# testCaseUI_rightClickableAccordion.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_rightClickableAccordion.go`
- Package: `testCaseUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `4`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `TappedSecondary`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/widget`
- `log`

## Declared Types
- `clickableAccordion`
- `clickableAccordionItem`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### TappedSecondary (method on `*clickableAccordion`)
- Signature: `func (*clickableAccordion) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- External calls: `log.Println`

### TappedSecondary (method on `*clickableAccordionItem`)
- Signature: `func (*clickableAccordionItem) TappedSecondary(_ *fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- External calls: `log.Println`

### newClickableAccordion (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) newClickableAccordion(accordionItem *widget.AccordionItem, isClickable bool, testCaseUuid, testInstructionUuid string) *clickableAccordion`
- Exported: `false`
- Control-flow features: `none detected`
- External calls: `accordion.Append`, `accordion.ExtendBaseWidget`

### newClickableAccordionItem (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) newClickableAccordionItem(title string, detail fyne.CanvasObject) *clickableAccordionItem`
- Exported: `false`
- Control-flow features: `none detected`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
