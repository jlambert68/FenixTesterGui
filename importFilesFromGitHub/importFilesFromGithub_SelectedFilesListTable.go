package importFilesFromGitHub

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

type clickableLabel struct {
	widget.Label
	onDoubleTap func()
	lastTapTime time.Time
	isClickable bool
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) newClickableLabel(text string, onDoubleTap func(), tempIsClickable bool) *clickableLabel {
	l := &clickableLabel{
		widget.Label{Text: text},
		onDoubleTap,
		time.Now(),
		tempIsClickable}
	l.ExtendBaseWidget(l)
	return l
}

func (l *clickableLabel) Tapped(e *fyne.PointEvent) {
	if l.isClickable == false {
		return
	}

	if time.Since(l.lastTapTime) < 500*time.Millisecond {
		if l.onDoubleTap != nil {
			l.onDoubleTap()
		}
	}
	l.lastTapTime = time.Now()
}

func (l *clickableLabel) TappedSecondary(*fyne.PointEvent) {
	// Implement if you need right-click (secondary tap) actions.
}

func (l *clickableLabel) MouseIn(*desktop.MouseEvent)    {}
func (l *clickableLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *clickableLabel) MouseOut()                      {}

// Create the UI-list that holds the selected files
func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) generateSelectedFilesListTable(parentWindow fyne.Window) {
	// Correctly initialize the selectedFilesTable as a new table
	importFilesFromGitHubObject.selectedFilesTable = widget.NewTable(
		func() (int, int) { return 0, 2 }, // Start with zero rows, 2 columns
		func() fyne.CanvasObject {
			return widget.NewLabel("") // Create cells as labels
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			// This should be filled when updating the table
		},
	)

	/*
			selectedFilesTable = widget.NewList(
				func() int { return len(selectedFiles) },
				func() fyne.CanvasObject { return widget.NewLabel("") },
				func(i widget.ListItemID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(selectedFiles[i].Name)
				},
			)



		selectedFilesTable.OnSelected = func(id widget.ListItemID) {

		}
	*/

}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) UpdateSelectedFilesTable() {

	importFilesFromGitHubObject.selectedFilesTable.Length = func() (int, int) {
		return len(importFilesFromGitHubObject.selectedFiles), 2
	}
	importFilesFromGitHubObject.selectedFilesTable.CreateCell = func() fyne.CanvasObject {
		return importFilesFromGitHubObject.newClickableLabel("", func() {}, false)

	}
	importFilesFromGitHubObject.selectedFilesTable.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		switch id.Col {
		case 0:
			// For the "Name" column, use the clickable label
			clickable := cell.(*clickableLabel)
			clickable.SetText(importFilesFromGitHubObject.selectedFiles[id.Row].Name)
			clickable.isClickable = true

			clickable.onDoubleTap = func() {

				// Remove the file from selectedFiles and refresh the list
				for fileIndex, file := range importFilesFromGitHubObject.selectedFiles {
					if file.URL == importFilesFromGitHubObject.selectedFiles[id.Row].URL {
						importFilesFromGitHubObject.selectedFiles = append(importFilesFromGitHubObject.selectedFiles[:fileIndex], importFilesFromGitHubObject.selectedFiles[fileIndex+1:]...)
						importFilesFromGitHubObject.selectedFilesTable.Unselect(id)
						importFilesFromGitHubObject.selectedFilesTable.Refresh()
						importFilesFromGitHubObject.UpdateSelectedFilesTable()
						break
					}
				}

			}

		case 1:
			// For the "URL" column, use a regular label

			nonClickable := cell.(*clickableLabel)
			nonClickable.SetText(importFilesFromGitHubObject.selectedFiles[id.Row].URL)
		}
	}

	maxNameWidth := float32(150) // Start with a minimum width
	maxUrlWidth := float32(250)  // Start with a minimum width
	for _, file := range importFilesFromGitHubObject.selectedFiles {
		textNameWidth := fyne.MeasureText(file.Name, theme.TextSize(), fyne.TextStyle{}).Width
		textUrlWidth := fyne.MeasureText(file.URL, theme.TextSize(), fyne.TextStyle{}).Width
		if textNameWidth > maxNameWidth {
			maxNameWidth = textNameWidth
		}
		if textUrlWidth > maxUrlWidth {
			maxUrlWidth = textUrlWidth
		}
	}

	importFilesFromGitHubObject.selectedFilesTable.SetColumnWidth(0, maxNameWidth+theme.Padding()*4) // Add padding
	importFilesFromGitHubObject.selectedFilesTable.SetColumnWidth(1, maxUrlWidth+theme.Padding()*4)  // Path column width can be static or calculated similarly

	importFilesFromGitHubObject.selectedFilesTable.Refresh()

}

type customLabel struct {
	widget.Label
	onDoubleTap func()
	lastTap     time.Time
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) newCustomLabel(text string, onDoubleTap func()) *customLabel {
	l := &customLabel{Label: widget.Label{Text: text}, onDoubleTap: onDoubleTap, lastTap: time.Now()}
	l.ExtendBaseWidget(l)
	return l
}

func (l *customLabel) Tapped(e *fyne.PointEvent) {
	now := time.Now()
	if now.Sub(l.lastTap) < 500*time.Millisecond { // 500 ms as double-click interval
		if l.onDoubleTap != nil {
			l.onDoubleTap()
		}
	}
	l.lastTap = now
}

func (l *customLabel) TappedSecondary(*fyne.PointEvent) {
	// Implement if you need right-click (secondary tap) actions.
}

func (l *customLabel) MouseIn(*desktop.MouseEvent)    {}
func (l *customLabel) MouseMoved(*desktop.MouseEvent) {}
func (l *customLabel) MouseOut()                      {}

/*
type coloredLabelItem struct {
	text  string
	color color.Color
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) newColoredLabelItem(text string, color color.Color) *coloredLabelItem {
	return &coloredLabelItem{text: text, color: color}
}

func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) (item *coloredLabelItem) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(item.text)
	label.color = item.color
	label.Refresh()

	return widget.NewSimpleRenderer(label)
}

*/
