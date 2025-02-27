package headertable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"sync"
)

type ColAttr struct {
	Alignment          fyne.TextAlign
	Header             string
	Name               string
	TextStyle          fyne.TextStyle
	WidthPercent       int
	Wrapping           fyne.TextWrap
	HeaderWidth        float32
	MaxColumnDataWidth float32
}

type TableOpts struct {
	Bindings                            []binding.DataMap
	ColAttrs                            []ColAttr
	OnDataCellSelect                    func(cellID widget.TableCellID)
	RefWidth                            string
	HeaderLabel                         string
	FlashingTableCellsReferenceMap      map[widget.TableCellID]*FlashingTableCellStruct
	FlashingTableCellsReferenceMapMutex *sync.RWMutex
}

type Header struct {
	widget.Table
}

// Load FlashingTableCellsReference from the FlashingTableCellsReferenceMap
func LoadFromFlashingTableCellsReferenceMap(
	tableOptsReference *TableOpts,
	flashingTableCellsReferenceMapKey widget.TableCellID) (
	flashingTableCellReference *FlashingTableCellStruct,
	existInMap bool) {

	// Lock Map for Reading
	tableOptsReference.FlashingTableCellsReferenceMapMutex.RLock()

	//UnLock Map
	defer tableOptsReference.FlashingTableCellsReferenceMapMutex.RUnlock()

	// Read Map
	flashingTableCellReference, existInMap = tableOptsReference.FlashingTableCellsReferenceMap[flashingTableCellsReferenceMapKey]

	return flashingTableCellReference, existInMap
}

// Save FlashingTableCellsReference to the FlashingTableCellsReferenceMap
func SaveToFlashingTableCellsReferenceMap(
	tableOptsReference *TableOpts,
	flashingTableCellsReferenceMapKey widget.TableCellID,
	flashingTableCellReference *FlashingTableCellStruct) {

	// Lock Map for Writing
	tableOptsReference.FlashingTableCellsReferenceMapMutex.Lock()

	//UnLock Map
	defer tableOptsReference.FlashingTableCellsReferenceMapMutex.Unlock()

	// Save to Map
	tableOptsReference.FlashingTableCellsReferenceMap[flashingTableCellsReferenceMapKey] = flashingTableCellReference

}
