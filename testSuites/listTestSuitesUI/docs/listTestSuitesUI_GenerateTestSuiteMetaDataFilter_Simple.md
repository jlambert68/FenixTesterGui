# listTestSuitesUI_GenerateTestSuiteMetaDataFilter_Simple.go

## File Overview
- Path: `testSuites/listTestSuitesUI/listTestSuitesUI_GenerateTestSuiteMetaDataFilter_Simple.go`
- Package: `listTestSuitesUI`
- Functions/Methods: `8`
- Imports: `11`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ConvertMetaDataToNewMap`

## Imports
- `FenixTesterGui/testCase/testCaseModel`
- `FenixTesterGui/testSuites/testSuitesModel`
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
### ConvertMetaDataToNewMap (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) ConvertMetaDataToNewMap(tc *testCaseModel.TestCaseMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: ConvertMetaDataToNewMap transforms the TestCaseMetaDataStruct.MetaDataGroupsSlicePtr into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.

### buildGUIFromMetaDataGroupsMap (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) buildGUIFromMetaDataGroupsMap(domainUUid string, metaDataGroupsOrder []string, metaDataGroupsSourceMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: buildGUIFromSlice builds a fyne.CanvasObject from your slice pointer
- Internal calls: `newCustomMandatoryCheckBoxGroupWidget`, `newCustomMandatorySelectComboBoxWidget`
- Selector calls: `boolbits.NewEntry`, `canvas.NewRectangle`, `chk.MinSize`, `chk.Refresh`, `container.New`, `container.NewVBox`, `err.Error`, `fmt.Sprintf`

### calcCheckGroupWidth (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) calcCheckGroupWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: calcCheckGroupWidth returns the width needed to show the widest checkbox label
- Selector calls: `tmp.MinSize`, `tmp.Refresh`, `widget.NewCheckGroup`

### calcSelectWidth (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) calcSelectWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: calcSelectWidth returns the width needed to show the longest option
- Internal calls: `float32`
- Selector calls: `tmp.MinSize`, `tmp.Refresh`, `widget.NewSelect`

### generateSimpleTestSuiteMetaDataDomainFilterBottomContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataDomainFilterBottomContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) simpleTestSuiteMetaDataDomainFilterBottomContainer *fyne.Container`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the bottom container having the Filter TestSuites-list- and clear MetaData.selection
- Selector calls: `autoFilterRadioGroup.SetSelected`, `container.NewHBox`, `listTestSuiteUIObject.calculateMetaDataFilterFunction`, `listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer`, `newSimpleTestSuiteMetaDataMainFilterContainer.Refresh`, `widget.NewButton`, `widget.NewRadioGroup`

### generateSimpleTestSuiteMetaDataDomainFilterTopContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataDomainFilterTopContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generates the top container having the Domain DropDown
- Internal calls: `newCustomMandatorySelectComboBoxWidget`
- Selector calls: `canvas.NewRectangle`, `container.New`, `container.NewVBox`, `layout.NewFormLayout`, `layout.NewVBoxLayout`, `listTestSuiteUIObject.calculateMetaDataFilterFunction`, `listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer`, `newSimpleTestSuiteMetaDataMainFilterContainer.Refresh`

### generateSimpleTestSuiteMetaDataFilterContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataFilterContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: generateSimpleTestSuiteMetaDataFilterContainer Generates the GenerateTestSuiteMetaDataFilterContainer containing the simple filter version
- Selector calls: `boolbits.NewAllOnesEntry`, `boolbits.NewAllZerosEntry`, `boolbits.NewBitSet`, `boolbits.NewEntry`, `booleanANDResultsEntry.And`, `booleanOrResultsEntry.Or`, `container.NewBorder`, `err.Error`

### generateSimpleTestSuiteMetaDataMainFilterContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataMainFilterContainer(domainUuidToGetMetaDataFor string, domainNameToGetMetaDataFor string, testCasesModel *testCaseModel.TestCasesModelsStruct) metaDataFilterArea *fyne.Container`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the main container having all the MetaData-filters
- Selector calls: `container.New`, `container.NewBorder`, `container.NewScroll`, `container.NewVBox`, `fmt.Sprintf`, `layout.NewGridLayout`, `listTestSuiteUIObject.buildGUIFromMetaDataGroupsMap`, `testSuitesModel.ConvertTestSuiteMetaData`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
