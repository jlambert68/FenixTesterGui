# newOrEditTestDataPointGroupUI_BuildPopUpTable.go

## File Overview
- Path: `testDataSelector/newOrEditTestDataPointGroupUI/newOrEditTestDataPointGroupUI_BuildPopUpTable.go`
- Package: `newOrEditTestDataPointGroupUI`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `1`
- Imports: `3`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `fmt`
- `github.com/jlambert68/FenixScriptEngine/testDataEngine`
- `regexp`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### buildPopUpTableDataFromTestDataPointName
- Signature: `func buildPopUpTableDataFromTestDataPointName(tempTestDataPointRowName string, testDataModel *testDataEngine.TestDataModelStruct) tableData [][]string`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Build the Table Data, based on TestDataPointName, to be used when the popup table is shown to the user to pick from
- Internal calls: `string`
- External calls: `fmt.Println`, `re.FindStringSubmatch`, `regexp.MustCompile`, `testDataEngine.TestDataAreaUuidType`, `testDataEngine.TestDataDomainOrAreaNameType`, `testDataEngine.TestDataDomainUuidType`, `testDataEngine.TestDataValueNameType`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
