# types.go

## File Overview
- Path: `headertable/types.go`
- Package: `headertable`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `2`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `LoadFromFlashingTableCellsReferenceMap`
- `SaveToFlashingTableCellsReferenceMap`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/data/binding`
- `fyne.io/fyne/v2/widget`
- `sync`

## Declared Types
- `ColAttr`
- `Header`
- `TableOpts`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### LoadFromFlashingTableCellsReferenceMap
- Signature: `func LoadFromFlashingTableCellsReferenceMap(tableOptsReference *TableOpts, flashingTableCellsReferenceMapKey widget.TableCellID) (flashingTableCellReference *FlashingTableCellStruct, existInMap bool)`
- Exported: `true`
- Control-flow features: `defer`
- Doc: Load FlashingTableCellsReference from the FlashingTableCellsReferenceMap

### SaveToFlashingTableCellsReferenceMap
- Signature: `func SaveToFlashingTableCellsReferenceMap(tableOptsReference *TableOpts, flashingTableCellsReferenceMapKey widget.TableCellID, flashingTableCellReference *FlashingTableCellStruct)`
- Exported: `true`
- Control-flow features: `defer`
- Doc: Save FlashingTableCellsReference to the FlashingTableCellsReferenceMap

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
