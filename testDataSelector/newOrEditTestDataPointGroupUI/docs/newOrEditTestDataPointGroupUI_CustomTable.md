# newOrEditTestDataPointGroupUI_CustomTable.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_CustomTable.go`
- Package: `newOrEditTestDataPointGroupUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `25`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CreateRenderer`
- `Destroy`
- `Layout`
- `MinSize`
- `MouseIn`
- `MouseMoved`
- `MouseOut`
- `NewCustomTableWidget`
- `NewCustomWidget`
- `Objects`
- `Refresh`
- `SetCellID`
- `SetText`
- `Tapped`

## Imports
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/dialog`
- `fyne.io/fyne/v2/driver/desktop`
- `fyne.io/fyne/v2/theme`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `image/color`
- `sort`
- `sync`

## Declared Types
- `CustomTableWidget`
- `customWidget`
- `customWidgetRenderer`
- `sortDirection`

## Declared Constants
- `dataSortAscending`
- `dataSortDescending`
- `dataSortOrderNotSelected`

## Declared Variables
- `ascendingSortIndicatorIcon`
- `descendingSortIndicatorIcon`
- `notFocusFortSortingIcon`
- `selectedAndHoveredRowColor`
- `selectedRowColor`
- `tableMutex`

## Functions and Methods
### CreateRenderer (method on `*customWidget`)
- Signature: `func (*customWidget) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `if`
- Doc: CreateRenderer implements fyne.WidgetRenderer for customWidget
- External calls: `container.NewHBox`, `w.updateVisibility`

### Destroy (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### Layout (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Layout(size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `if`
- External calls: `fyne.NewSize`, `tempLabel.MinSize`, `tempLabel.Refresh`, `widget.NewLabel`

### MouseIn (method on `*customWidget`)
- Signature: `func (*customWidget) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- External calls: `w.Refresh`, `w.onHover`

### MouseMoved (method on `*customWidget`)
- Signature: `func (*customWidget) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### MouseOut (method on `*customWidget`)
- Signature: `func (*customWidget) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- External calls: `w.Refresh`, `w.onHover`

### NewCustomTableWidget
- Signature: `func NewCustomTableWidget(data [][]string, selectedTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct) *CustomTableWidget`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `NewCustomWidget`, `setColumnWidths`
- External calls: `table.ExtendBaseWidget`, `table.handleCellTapped`, `table.hoverRow`, `table.unhoverRow`, `tableMutex.Lock`, `tableMutex.Unlock`, `testDataEngine.TestDataPointRowUuidType`

### NewCustomWidget
- Signature: `func NewCustomWidget(isSelected bool, text string, tableRef *CustomTableWidget) *customWidget`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: NewCustomWidget creates a new customWidget
- External calls: `canvas.NewRectangle`, `theme.BackgroundColor`, `theme.CheckButtonCheckedIcon`, `w.ExtendBaseWidget`, `widget.NewIcon`, `widget.NewLabel`

### Objects (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### Refresh (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`
- External calls: `canvas.Refresh`

### SetCellID (method on `*customWidget`)
- Signature: `func (*customWidget) SetCellID(cellID widget.TableCellID)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetCellID sets the position of the cell in the Table

### SetText (method on `*customWidget`)
- Signature: `func (*customWidget) SetText(text string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetText sets the text of the label

### Tapped (method on `*customWidget`)
- Signature: `func (*customWidget) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Tapped handles tap events
- External calls: `w.onTapped`

### addDataPointToAvailableDataPointsAndRemoveFromSelectedDataPoints
- Signature: `func addDataPointToAvailableDataPointsAndRemoveFromSelectedDataPoints(testDataPointName testDataEngine.TestDataValueNameType, testDataPointRowUuidObject testDataEngine.TestDataPointRowUuidStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `setStateForSaveButtonAndGroupNameTextEntryExternalCall`
- External calls: `allAvailablePointsList.Refresh`, `selectedPointsList.Refresh`

### addDataPointToSelectedDataPointsAndRemoveFromAvailableDataPoints
- Signature: `func addDataPointToSelectedDataPointsAndRemoveFromAvailableDataPoints(testDataPointName testDataEngine.TestDataValueNameType, testDataPointRowUuidObject testDataEngine.TestDataPointRowUuidStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Add the TestDataPoint to the Selected-list and removes it from the Avaialables-lsit
- Internal calls: `setStateForSaveButtonAndGroupNameTextEntryExternalCall`, `sortDataPointsList`
- External calls: `allAvailablePointsList.Refresh`, `selectedPointsList.Refresh`

### generateTestDataPointRowUuidObject
- Signature: `func generateTestDataPointRowUuidObject(testDataPointRowUuid testDataEngine.TestDataPointRowUuidType, dataRow []string) TestDataPointRowUuidObject testDataEngine.TestDataPointRowUuidStruct`
- Exported: `false`
- Control-flow features: `if, for/range`
- External calls: `testDataEngine.TestDataPointRowValuesSummaryType`

### handleCellTapped (method on `*CustomTableWidget`)
- Signature: `func (*CustomTableWidget) handleCellTapped(cellID widget.TableCellID, table *CustomTableWidget)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `sortTable`, `updateRowsSelectedMap`
- External calls: `t.Refresh`, `table.Refresh`

### hoverRow (method on `*CustomTableWidget`)
- Signature: `func (*CustomTableWidget) hoverRow(row int)`
- Exported: `false`
- Control-flow features: `if, for/range`
- External calls: `customWidget.Refresh`, `t.Length`

### setColumnWidths
- Signature: `func setColumnWidths(table *widget.Table, data [][]string)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: setColumnWidths adapt all columns in the popup window to fit its headers and data
- External calls: `fyne.MeasureText`, `table.SetColumnWidth`, `theme.Padding`, `theme.TextSize`

### showTable
- Signature: `func showTable(w fyne.Window, data [][]string, selectedTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: showTable creates and shows a table for the selected node with data
- Internal calls: `NewCustomTableWidget`, `addDataPointToAvailableDataPointsAndRemoveFromSelectedDataPoints`, `addDataPointToSelectedDataPointsAndRemoveFromAvailableDataPoints`, `generateTestDataPointRowUuidObject`
- External calls: `container.NewScroll`, `dialog.NewCustomConfirm`, `fyne.NewSize`, `modal.Resize`, `modal.Show`, `table.Resize`, `testDataEngine.TestDataPointRowUuidType`, `testDataEngine.TestDataValueNameType`

### sortTable
- Signature: `func sortTable(data [][]string, column int, direction sortDirection)`
- Exported: `false`
- Control-flow features: `if`
- Doc: sortTable sorts a 2D string slice based on the specified column and direction, keeping the first row intact.
- External calls: `sort.Slice`

### unhoverRow (method on `*CustomTableWidget`)
- Signature: `func (*CustomTableWidget) unhoverRow(row int)`
- Exported: `false`
- Control-flow features: `if, for/range`
- External calls: `customWidget.Refresh`, `t.Length`

### updateRowsSelectedMap
- Signature: `func updateRowsSelectedMap(table *CustomTableWidget)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: updateRowsSelectedMap updates the map holding which row that is selected

### updateVisibility (method on `*customWidget`)
- Signature: `func (*customWidget) updateVisibility()`
- Exported: `false`
- Control-flow features: `if`
- Doc: updateVisibility updates the visibility of the label and icon based on the isIcon flag
- External calls: `theme.BackgroundColor`, `theme.PrimaryColor`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
