# testSuiteUI_graphicalComponent_SuiteMetaData.go

## File Overview
- Path: `testSuites/testSuiteUI/testSuiteUI_graphicalComponent_SuiteMetaData.go`
- Package: `testSuiteUI`
- Functions/Methods: `3`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `GenerateMetaDataAreaForTestCase`

## Imports
- `FenixTesterGui/testSuites/testSuitesModel`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `image/color`
- `log`

## Declared Types
- `NewMetaDataInGroupStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### GenerateMetaDataAreaForTestCase (method on `TestSuiteUiStruct`)
- Signature: `func (TestSuiteUiStruct) GenerateMetaDataAreaForTestCase() (testSuiteMetaDataContainer *fyne.Container, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: Generate the MetaData Area for the TestCase
- Selector calls: `container.New`, `container.NewBorder`, `container.NewScroll`, `container.NewVBox`, `fyne.NewSize`, `layout.NewGridLayout`, `myContainer.MinSize`, `myContainerScroll.SetMinSize`

### buildGUIFromMetaDataGroupsMap (method on `TestSuiteUiStruct`)
- Signature: `func (TestSuiteUiStruct) buildGUIFromMetaDataGroupsMap(metaDataGroupsOrder []string, metaDataGroupsSourceMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct, metaDataGroupInTestSuitePtr *testSuitesModel.TestSuiteMetaDataStruct) fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: buildGUIFromSlice builds a Container from your slice pointer
- Internal calls: `calcCheckGroupWidth`, `calcSelectWidth`, `newCustomAttributeCheckBoxGroupWidget`, `newCustomSelectComboBoxWidget`
- Selector calls: `canvas.NewRectangle`, `chk.MinSize`, `chk.Refresh`, `container.New`, `container.NewVBox`, `fyne.NewSize`, `layout.NewGridWrapLayout`, `layout.NewHBoxLayout`

### convertMetaDataToNewMap (method on `TestSuiteUiStruct`)
- Signature: `func (TestSuiteUiStruct) convertMetaDataToNewMap(ts *testSuitesModel.TestSuiteMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: convertMetaDataToNewMap transforms the TestsUITEMetaDataStruct.MetaDataGroupsSlicePtr into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
