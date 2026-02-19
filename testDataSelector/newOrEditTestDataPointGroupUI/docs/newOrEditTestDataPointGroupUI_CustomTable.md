# newOrEditTestDataPointGroupUI_CustomTable.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_CustomTable.go`
- Package: `newOrEditTestDataPointGroupUI`
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
### NewCustomWidget
- Signature: `func NewCustomWidget(isSelected bool, text string, tableRef *CustomTableWidget) *customWidget`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: NewCustomWidget creates a new customWidget
- Selector calls: `widget.NewLabel`, `widget.NewIcon`, `theme.CheckButtonCheckedIcon`, `canvas.NewRectangle`, `theme.BackgroundColor`, `w.ExtendBaseWidget`

### CreateRenderer (method on `*customWidget`)
- Signature: `func (*customWidget) CreateRenderer() fyne.WidgetRenderer`
- Exported: `true`
- Control-flow features: `if`
- Doc: CreateRenderer implements fyne.WidgetRenderer for customWidget
- Selector calls: `w.updateVisibility`, `container.NewHBox`

### Layout (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Layout(size fyne.Size)`
- Exported: `true`
- Control-flow features: `none detected`

### MinSize (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) MinSize() fyne.Size`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `widget.NewLabel`, `tempLabel.Refresh`, `fyne.NewSize`, `tempLabel.MinSize`

### Refresh (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Refresh()`
- Exported: `true`
- Control-flow features: `none detected`
- Selector calls: `canvas.Refresh`

### Objects (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Objects() []fyne.CanvasObject`
- Exported: `true`
- Control-flow features: `none detected`

### Destroy (method on `*customWidgetRenderer`)
- Signature: `func (*customWidgetRenderer) Destroy()`
- Exported: `true`
- Control-flow features: `none detected`

### updateVisibility (method on `*customWidget`)
- Signature: `func (*customWidget) updateVisibility()`
- Exported: `false`
- Control-flow features: `if`
- Doc: updateVisibility updates the visibility of the label and icon based on the isIcon flag
- Selector calls: `theme.PrimaryColor`, `theme.BackgroundColor`

### SetText (method on `*customWidget`)
- Signature: `func (*customWidget) SetText(text string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetText sets the text of the label

### SetCellID (method on `*customWidget`)
- Signature: `func (*customWidget) SetCellID(cellID widget.TableCellID)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetCellID sets the position of the cell in the Table

### MouseIn (method on `*customWidget`)
- Signature: `func (*customWidget) MouseIn(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `w.onHover`, `w.Refresh`

### MouseOut (method on `*customWidget`)
- Signature: `func (*customWidget) MouseOut()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `w.onHover`, `w.Refresh`

### MouseMoved (method on `*customWidget`)
- Signature: `func (*customWidget) MouseMoved(*desktop.MouseEvent)`
- Exported: `true`
- Control-flow features: `none detected`

### Tapped (method on `*customWidget`)
- Signature: `func (*customWidget) Tapped(*fyne.PointEvent)`
- Exported: `true`
- Control-flow features: `if`
- Doc: Tapped handles tap events
- Selector calls: `w.onTapped`

### NewCustomTableWidget
- Signature: `func NewCustomTableWidget(data [][]string, selectedTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct) *CustomTableWidget`
- Exported: `true`
- Control-flow features: `if, for/range`
- Internal calls: `NewCustomWidget`, `setColumnWidths`
- Selector calls: `table.hoverRow`, `table.unhoverRow`, `table.handleCellTapped`, `tableMutex.Lock`, `tableMutex.Unlock`, `table.ExtendBaseWidget`, `testDataEngine.TestDataPointRowUuidType`

### handleCellTapped (method on `*CustomTableWidget`)
- Signature: `func (*CustomTableWidget) handleCellTapped(cellID widget.TableCellID, table *CustomTableWidget)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `sortTable`, `updateRowsSelectedMap`
- Selector calls: `table.Refresh`, `t.Refresh`

### hoverRow (method on `*CustomTableWidget`)
- Signature: `func (*CustomTableWidget) hoverRow(row int)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Selector calls: `t.Length`, `customWidget.Refresh`

### unhoverRow (method on `*CustomTableWidget`)
- Signature: `func (*CustomTableWidget) unhoverRow(row int)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Selector calls: `t.Length`, `customWidget.Refresh`

### sortTable
- Signature: `func sortTable(data [][]string, column int, direction sortDirection)`
- Exported: `false`
- Control-flow features: `if`
- Doc: sortTable sorts a 2D string slice based on the specified column and direction, keeping the first row intact.
- Selector calls: `sort.Slice`

### updateRowsSelectedMap
- Signature: `func updateRowsSelectedMap(table *CustomTableWidget)`
- Exported: `false`
- Control-flow features: `for/range`
- Doc: updateRowsSelectedMap updates the map holding which row that is selected

### setColumnWidths
- Signature: `func setColumnWidths(table *widget.Table, data [][]string)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: setColumnWidths adapt all columns in the popup window to fit its headers and data
- Selector calls: `fyne.MeasureText`, `theme.TextSize`, `theme.Padding`, `table.SetColumnWidth`

### showTable
- Signature: `func showTable(w fyne.Window, data [][]string, selectedTestDataPointUuidMap map[testDataEngine.TestDataPointRowUuidType]testDataEngine.TestDataPointRowUuidStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: showTable creates and shows a table for the selected node with data
- Internal calls: `NewCustomTableWidget`, `generateTestDataPointRowUuidObject`, `addDataPointToSelectedDataPointsAndRemoveFromAvailableDataPoints`, `addDataPointToAvailableDataPointsAndRemoveFromSelectedDataPoints`
- Selector calls: `table.Resize`, `fyne.NewSize`, `container.NewScroll`, `dialog.NewCustomConfirm`, `testDataEngine.TestDataValueNameType`, `testDataEngine.TestDataPointRowUuidType`, `modal.Resize`, `modal.Show`

### generateTestDataPointRowUuidObject
- Signature: `func generateTestDataPointRowUuidObject(testDataPointRowUuid testDataEngine.TestDataPointRowUuidType, dataRow []string) TestDataPointRowUuidObject testDataEngine.TestDataPointRowUuidStruct`
- Exported: `false`
- Control-flow features: `if, for/range`
- Selector calls: `testDataEngine.TestDataPointRowValuesSummaryType`

### addDataPointToSelectedDataPointsAndRemoveFromAvailableDataPoints
- Signature: `func addDataPointToSelectedDataPointsAndRemoveFromAvailableDataPoints(testDataPointName testDataEngine.TestDataValueNameType, testDataPointRowUuidObject testDataEngine.TestDataPointRowUuidStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Add the TestDataPoint to the Selected-list and removes it from the Avaialables-lsit
- Internal calls: `sortDataPointsList`, `setStateForSaveButtonAndGroupNameTextEntryExternalCall`
- Selector calls: `allAvailablePointsList.Refresh`, `selectedPointsList.Refresh`

### addDataPointToAvailableDataPointsAndRemoveFromSelectedDataPoints
- Signature: `func addDataPointToAvailableDataPointsAndRemoveFromSelectedDataPoints(testDataPointName testDataEngine.TestDataValueNameType, testDataPointRowUuidObject testDataEngine.TestDataPointRowUuidStruct)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `setStateForSaveButtonAndGroupNameTextEntryExternalCall`
- Selector calls: `allAvailablePointsList.Refresh`, `selectedPointsList.Refresh`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
