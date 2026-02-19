# clickableAttributeInPreview.go

## File Overview
- Path: `testCaseExecutions/listTestCaseExecutionsUI/clickableAttributeInPreview.go`
- Package: `listTestCaseExecutionsUI`
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
### newClickableAttributeInPreview
- Signature: `func newClickableAttributeInPreview(attributeValue string, attributeName string, testInstructionName string, leftClicked func(), rightClicked func(), attributeType attributeTypeType, testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) *clickableAttributeInPreviewStruct`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: Used for creating a new Attribute label
- Selector calls: `time.Now`, `clickableAttributeInPreview.ExtendBaseWidget`

### CreateRenderer (method on `*clickableAttributeInPreviewStruct`)
- Signature: `func (*clickableAttributeInPreviewStruct) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: CreateRenderer Renderer (required by fyne.Widget)
- Selector calls: `widget.NewLabel`, `lbl.MinSize`, `maxlabelStringBuilder.WriteString`, `maxlabelStringBuilder.String`, `tempMaxWidget.Refresh`, `tempMaxWidget.MinSize`, `container.NewBorder`, `labelBorderContainer.Resize`

### Tapped (method on `*clickableAttributeInPreviewStruct`)
- Signature: `func (*clickableAttributeInPreviewStruct) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Tapped Tapped interface clickableAttributeInPreviewStruct
- Selector calls: `container.NewVBox`, `widget.NewButton`, `fenixMasterWindow.Clipboard`, `textToCopy.WriteString`, `fmt.Sprintf`, `clipboard.SetContent`, `textToCopy.String`, `fyne.CurrentApp`

### TappedSecondary (method on `*clickableAttributeInPreviewStruct`)
- Signature: `func (*clickableAttributeInPreviewStruct) TappedSecondary(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: TappedSecondary Optional: Handle secondary tap (right-click)
- Selector calls: `fenixMasterWindow.Clipboard`, `clipboard.SetContent`, `fyne.CurrentApp`, `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
