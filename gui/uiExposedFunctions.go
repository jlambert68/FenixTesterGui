package gui

import (
	"errors"
	"fmt"
)

func (uiServer *UIServerStruct) GetDraggedBuildingBlockType(draggedObjectsName string) (buildingBlockType BuildingBlock, err error) {

	draggedElement, existsInMap := uiServer.AvailableBuildingBlocksModel.allBuildingBlocksTreeNameToUuid[draggedObjectsName]
	if existsInMap == false {
		errorId := "545c02bd-2287-40f1-8a09-0f0969398278"
		err = errors.New(fmt.Sprintf("dragged element with name '%s' doesn't exist in map with all availeable building blocks that is name prepared for UI-tree to Object mapping [ErrorID: %s]", draggedObjectsName, errorId))

		return 0, err
	}

	buildingBlockType = draggedElement.buildingBlockType

	return buildingBlockType, err
}
