# testCaseUI_graphicalRepresentation_metaData.go

## File Overview
- Path: `testCase/testCaseUI/testCaseUI_graphicalRepresentation_metaData.go`
- Package: `testCaseUI`
- Functions/Methods: `5`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ConvertMetaDataToNewMap`
- `GenerateMetaDataAreaForTestCase`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `errors`
- `fmt`
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
### GenerateMetaDataAreaForTestCase (method on `*TestCasesUiModelStruct`)
- Signature: `func (*TestCasesUiModelStruct) GenerateMetaDataAreaForTestCase(tempTestCaseRef *testCaseModel.TestCaseModelStruct, testCaseUuid string, domainUuidToGetMetaDataFor string) (testCaseMetaDataArea fyne.CanvasObject, accordion *widget.Accordion, err error)`
- Exported: `true`
- Control-flow features: `if, returns error`
- Doc: Generate the MetaData Area for the TestCase
- Internal calls: `buildGUIFromMetaDataGroupsMap`
- Selector calls: `errors.New`, `fmt.Sprintf`, `testCaseModel.ConvertTestCaseMetaData`, `container.NewBorder`, `myContainer.MinSize`, `container.NewScroll`, `myContainerScroll.SetMinSize`, `fyne.NewSize`

### buildGUIFromMetaDataGroupsMap
- Signature: `func buildGUIFromMetaDataGroupsMap(testCaseUuid string, testCasesModelReference *testCaseModel.TestCasesModelsStruct, metaDataGroupsOrder []string, metaDataGroupsSourceMapPtr *map[string]*testCaseModel.MetaDataGroupStruct, metaDataGroupInTestCasePtr *testCaseModel.TestCaseMetaDataStruct) fyne.CanvasObject`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: buildGUIFromSlice builds a fyne.CanvasObject from your slice pointer
- Internal calls: `ConvertMetaDataToNewMap`, `NewCustomAttributeSelectComboBoxWidget`, `calcSelectWidth`, `newCustomAttributeCheckBoxGroupWidget`, `calcCheckGroupWidth`
- Selector calls: `log.Fatalln`, `canvas.NewRectangle`, `widget.NewSelect`, `sel.SetSelected`, `container.New`, `layout.NewGridWrapLayout`, `fyne.NewSize`, `sel.MinSize`

### calcSelectWidth
- Signature: `func calcSelectWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: calcSelectWidth returns the width needed to show the longest option
- Internal calls: `float32`
- Selector calls: `widget.NewSelect`, `tmp.Refresh`, `tmp.MinSize`

### calcCheckGroupWidth
- Signature: `func calcCheckGroupWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: calcCheckGroupWidth returns the width needed to show the widest checkbox label
- Selector calls: `widget.NewCheckGroup`, `tmp.Refresh`, `tmp.MinSize`

### ConvertMetaDataToNewMap
- Signature: `func ConvertMetaDataToNewMap(tc *testCaseModel.TestCaseMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: ConvertMetaDataToNewMap transforms the TestCaseMetaDataStruct.MetaDataGroupsSlicePtr into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
