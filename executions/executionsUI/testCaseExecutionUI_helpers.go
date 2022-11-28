package executionsUI

import "fyne.io/fyne/v2/data/binding"

// Remove item from the DataItem-slice and keep order
func remove(slice []binding.DataMap, s int) []binding.DataMap {
	return append(slice[:s], slice[s+1:]...)
}
