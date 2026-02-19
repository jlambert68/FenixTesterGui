# testCaseModel_generateTestCaseMetaDataStructureForSelectedDomain.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_generateTestCaseMetaDataStructureForSelectedDomain.go`
- Package: `testCaseModel`
- Functions/Methods: `3`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `ConvertTestCaseMetaData`
- `GenerateTestCaseMetaDataStructureForSelectedDomain`

## Imports
- `FenixTesterGui/common_code`
- `errors`
- `fmt`
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### ConvertTestCaseMetaData
- Signature: `func ConvertTestCaseMetaData(testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct) (*map[string]*MetaDataGroupStruct, []string)`
- Exported: `true`
- Control-flow features: `if, for/range, switch`
- Doc: ConvertTestCaseMetaData converts the JSON‐parsed TestCaseMetaDataForDomainStruct into the GUI‐friendly slice *[]*MetaDataGroupStruct.

### GenerateTestCaseMetaDataStructureForSelectedDomain (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) GenerateTestCaseMetaDataStructureForSelectedDomain(testCaseUuid string, selectedDomainUuid string) err error`
- Exported: `true`
- Control-flow features: `if, for/range, returns error`
- Doc: GenerateTestCaseMetaDataStructureForSelectedDomain - Verify if the Hash for the TestCase is the same as the one in the database
- Internal calls: `buildMetaDataGroups`
- Selector calls: `errors.New`, `fmt.Println`, `fmt.Sprintf`

### buildMetaDataGroups
- Signature: `func buildMetaDataGroups(testCaseMetaDataForDomain *TestCaseMetaDataForDomainStruct) *map[string]*MetaDataGroupStruct`
- Exported: `false`
- Control-flow features: `if, for/range, switch`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
