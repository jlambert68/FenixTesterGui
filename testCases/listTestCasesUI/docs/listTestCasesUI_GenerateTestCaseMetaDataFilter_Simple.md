# listTestCasesUI_GenerateTestCaseMetaDataFilter_Simple.go

## File Overview
- Path: `testCases/listTestCasesUI/listTestCasesUI_GenerateTestCaseMetaDataFilter_Simple.go`
- Package: `listTestCasesUI`
- Functions/Methods: `8`
- Imports: `10`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ConvertMetaDataToNewMap`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `fmt`
- `fyne.io/fyne/v2`
- `fyne.io/fyne/v2/canvas`
- `fyne.io/fyne/v2/container`
- `fyne.io/fyne/v2/layout`
- `fyne.io/fyne/v2/widget`
- `github.com/jlambert68/Fast_BitFilter_MetaData/boolbits/boolbits`
- `image/color`
- `log`

## Declared Types
- `NewMetaDataInGroupStruct`

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### generateSimpleTestCaseMetaDataFilterContainer (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) generateSimpleTestCaseMetaDataFilterContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: generateSimpleTestCaseMetaDataFilterContainer Generates the GenerateTestCaseMetaDataFilterContainer containing the simple filter version
- Selector calls: `listTestCaseUIObject.loadTestCaseListTableTable`, `listTestCaseUIObject.calculateAndSetCorrectColumnWidths`, `listTestCaseUIObject.updateTestCasesListTable`, `boolbits.NewAllZerosEntry`, `fmt.Sprintf`, `err.Error`, `log.Fatalln`, `boolbits.NewAllOnesEntry`

### generateSimpleTestCaseMetaDataDomainFilterTopContainer (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) generateSimpleTestCaseMetaDataDomainFilterTopContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generates the top container having the Domain DropDown
- Internal calls: `newCustomMandatorySelectComboBoxWidget`
- Selector calls: `container.New`, `layout.NewVBoxLayout`, `layout.NewFormLayout`, `widget.NewLabel`, `testCaseOwnerDomainNameFormContainer.Add`, `canvas.NewRectangle`, `widget.NewSelect`, `listTestCaseUIObject.generateSimpleTestCaseMetaDataMainFilterContainer`

### generateSimpleTestCaseMetaDataDomainFilterBottomContainer (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) generateSimpleTestCaseMetaDataDomainFilterBottomContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) simpleTestCaseMetaDataDomainFilterBottomContainer *fyne.Container`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the bottom container having the Filter TestCases-list- and clear MetaData.selection
- Selector calls: `widget.NewButton`, `listTestCaseUIObject.calculateMetaDataFilterFunction`, `listTestCaseUIObject.generateSimpleTestCaseMetaDataMainFilterContainer`, `newSimpleTestCaseMetaDataMainFilterContainer.Refresh`, `widget.NewRadioGroup`, `autoFilterRadioGroup.SetSelected`, `container.NewHBox`

### generateSimpleTestCaseMetaDataMainFilterContainer (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) generateSimpleTestCaseMetaDataMainFilterContainer(domainUuidToGetMetaDataFor string, domainNameToGetMetaDataFor string, testCasesModel *testCaseModel.TestCasesModelsStruct) metaDataFilterArea *fyne.Container`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the main container having all the MetaData-filters
- Selector calls: `container.NewVBox`, `widget.NewLabel`, `fmt.Sprintf`, `testCaseModel.ConvertTestCaseMetaData`, `listTestCaseUIObject.buildGUIFromMetaDataGroupsMap`, `container.NewScroll`, `container.NewBorder`, `container.New`

### buildGUIFromMetaDataGroupsMap (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) buildGUIFromMetaDataGroupsMap(domainUUid string, metaDataGroupsOrder []string, metaDataGroupsSourceMapPtr *map[string]*testCaseModel.MetaDataGroupStruct, testCasesModels *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: buildGUIFromSlice builds a fyne.CanvasObject from your slice pointer
- Internal calls: `newCustomMandatorySelectComboBoxWidget`, `newCustomMandatoryCheckBoxGroupWidget`
- Selector calls: `log.Fatalln`, `canvas.NewRectangle`, `widget.NewSelect`, `fmt.Sprintf`, `boolbits.NewEntry`, `err.Error`, `listTestCaseUIObject.calculateMetaDataFilterFunction`, `sel.SetSelected`

### calcSelectWidth (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) calcSelectWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: calcSelectWidth returns the width needed to show the longest option
- Internal calls: `float32`
- Selector calls: `widget.NewSelect`, `tmp.Refresh`, `tmp.MinSize`

### calcCheckGroupWidth (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) calcCheckGroupWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: calcCheckGroupWidth returns the width needed to show the widest checkbox label
- Selector calls: `widget.NewCheckGroup`, `tmp.Refresh`, `tmp.MinSize`

### ConvertMetaDataToNewMap (method on `*ListTestCaseUIStruct`)
- Signature: `func (*ListTestCaseUIStruct) ConvertMetaDataToNewMap(tc *testCaseModel.TestCaseMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: ConvertMetaDataToNewMap transforms the TestCaseMetaDataStruct.MetaDataGroupsSlicePtr into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
