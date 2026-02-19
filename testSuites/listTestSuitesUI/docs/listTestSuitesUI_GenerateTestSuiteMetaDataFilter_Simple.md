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
### generateSimpleTestSuiteMetaDataFilterContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataFilterContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: generateSimpleTestSuiteMetaDataFilterContainer Generates the GenerateTestSuiteMetaDataFilterContainer containing the simple filter version
- Selector calls: `listTestSuiteUIObject.loadTestSuiteListTableTable`, `listTestSuiteUIObject.calculateAndSetCorrectColumnWidths`, `listTestSuiteUIObject.updateTestSuitesListTable`, `boolbits.NewAllZerosEntry`, `fmt.Sprintf`, `err.Error`, `log.Fatalln`, `boolbits.NewAllOnesEntry`

### generateSimpleTestSuiteMetaDataDomainFilterTopContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataDomainFilterTopContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generates the top container having the Domain DropDown
- Internal calls: `newCustomMandatorySelectComboBoxWidget`
- Selector calls: `container.New`, `layout.NewVBoxLayout`, `layout.NewFormLayout`, `widget.NewLabel`, `testSuiteOwnerDomainNameFormContainer.Add`, `canvas.NewRectangle`, `widget.NewSelect`, `listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer`

### generateSimpleTestSuiteMetaDataDomainFilterBottomContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataDomainFilterBottomContainer(testCasesModel *testCaseModel.TestCasesModelsStruct) simpleTestSuiteMetaDataDomainFilterBottomContainer *fyne.Container`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the bottom container having the Filter TestSuites-list- and clear MetaData.selection
- Selector calls: `widget.NewButton`, `listTestSuiteUIObject.calculateMetaDataFilterFunction`, `listTestSuiteUIObject.generateSimpleTestSuiteMetaDataMainFilterContainer`, `newSimpleTestSuiteMetaDataMainFilterContainer.Refresh`, `widget.NewRadioGroup`, `autoFilterRadioGroup.SetSelected`, `container.NewHBox`

### generateSimpleTestSuiteMetaDataMainFilterContainer (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) generateSimpleTestSuiteMetaDataMainFilterContainer(domainUuidToGetMetaDataFor string, domainNameToGetMetaDataFor string, testCasesModel *testCaseModel.TestCasesModelsStruct) metaDataFilterArea *fyne.Container`
- Exported: `false`
- Control-flow features: `if`
- Doc: Generates the main container having all the MetaData-filters
- Selector calls: `container.NewVBox`, `widget.NewLabel`, `fmt.Sprintf`, `testSuitesModel.ConvertTestSuiteMetaData`, `listTestSuiteUIObject.buildGUIFromMetaDataGroupsMap`, `container.NewScroll`, `container.NewBorder`, `container.New`

### buildGUIFromMetaDataGroupsMap (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) buildGUIFromMetaDataGroupsMap(domainUUid string, metaDataGroupsOrder []string, metaDataGroupsSourceMapPtr *map[string]*testSuitesModel.MetaDataGroupStruct) *fyne.Container`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: buildGUIFromSlice builds a fyne.CanvasObject from your slice pointer
- Internal calls: `newCustomMandatorySelectComboBoxWidget`, `newCustomMandatoryCheckBoxGroupWidget`
- Selector calls: `log.Fatalln`, `canvas.NewRectangle`, `widget.NewSelect`, `fmt.Sprintf`, `boolbits.NewEntry`, `err.Error`, `listTestSuiteUIObject.calculateMetaDataFilterFunction`, `sel.SetSelected`

### calcSelectWidth (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) calcSelectWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: calcSelectWidth returns the width needed to show the longest option
- Internal calls: `float32`
- Selector calls: `widget.NewSelect`, `tmp.Refresh`, `tmp.MinSize`

### calcCheckGroupWidth (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) calcCheckGroupWidth(values []string) float32`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: calcCheckGroupWidth returns the width needed to show the widest checkbox label
- Selector calls: `widget.NewCheckGroup`, `tmp.Refresh`, `tmp.MinSize`

### ConvertMetaDataToNewMap (method on `*ListTestSuiteUIStruct`)
- Signature: `func (*ListTestSuiteUIStruct) ConvertMetaDataToNewMap(tc *testCaseModel.TestCaseMetaDataStruct) map[string]map[string]*NewMetaDataInGroupStruct`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: ConvertMetaDataToNewMap transforms the TestCaseMetaDataStruct.MetaDataGroupsSlicePtr into a nested map[groupName][metaDataName] => *NewMetaDataInGroupStruct.

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
