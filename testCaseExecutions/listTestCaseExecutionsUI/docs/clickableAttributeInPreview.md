# clickableAttributeInPreview.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/clickableAttributeInPreview.go`
- Package: `listTestCaseExecutionsUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `4`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Tapped`
- `TappedSecondary`

## Imports
- `FenixTesterGui/common_code`
- `FenixTesterGui/soundEngine`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/widget`
- `strings`
- `time`

## Declared Types
- `attributeTypeType`
- `clickableAttributeInPreviewStruct`

## Declared Constants
- `attributeIsOriginal`
- `attributeIsRunTimeChanged`
- `attributeTypeNotDefined`
- `numberOfVisibleCharacters`
- `numberOfVisibleRow`
- `visibleCharacterOfType`

## Declared Variables
- None

## Functions and Methods
### CreateRenderer (method on `*clickableAttributeInPreviewStruct`)
- Signature: `func (*clickableAttributeInPreviewStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: CreateRenderer Renderer (required by fyne.Widget)
- External calls: `container.NewBorder`, `container.NewScroll`, `labelBorderContainer.Refresh`, `labelBorderContainer.Resize`, `labelScrollContainer.Refresh`, `labelScrollContainer.Resize`, `labelScrollContainer.SetMinSize`, `lbl.MinSize`

### Tapped (method on `*clickableAttributeInPreviewStruct`)
- Signature: `func (*clickableAttributeInPreviewStruct) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Tapped Tapped interface clickableAttributeInPreviewStruct
- External calls: `attributeMessageBorderContainer.Refresh`, `attributeMessageStringBuilder.String`, `attributeMessageStringBuilder.WriteString`, `c.LeftClicked`, `clipboard.SetContent`, `container.NewBorder`, `container.NewScroll`, `container.NewVBox`

### TappedSecondary (method on `*clickableAttributeInPreviewStruct`)
- Signature: `func (*clickableAttributeInPreviewStruct) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TappedSecondary Optional: Handle secondary tap (right-click)
- External calls: `clipboard.SetContent`, `fenixMasterWindow.Clipboard`, `fmt.Sprintf`, `fyne.CurrentApp`

### newClickableAttributeInPreview
- Signature: `func newClickableAttributeInPreview(attributeValue string, attributeName string, testInstructionName string, leftClicked func(), rightClicked func(), attributeType attributeTypeType, testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) *clickableAttributeInPreviewStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new Attribute label
- External calls: `clickableAttributeInPreview.ExtendBaseWidget`, `time.Now`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
