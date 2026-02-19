# testCaseModel_verifyMandatoryFieldsForMetaData.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_verifyMandatoryFieldsForMetaData.go`
- Package: `testCaseModel`
- Functions/Methods: `2`
- Imports: `4`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `FenixTesterGui/soundEngine`
- `fmt`
- `fyne.io/fyne/v2`
- `log`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### verifyMandatoryFieldsForMetaData (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) verifyMandatoryFieldsForMetaData(domainUuid string, currentTestCasePtr *TestCaseModelStruct, shouldBeSaved bool) err error`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Validates that all mandatory MetaData fields has values for specified DomainUuid
- Selector calls: `testCaseModel.generateMandatoryMetaDataFieldsMap`, `fmt.Sprintf`, `log.Fatalln`, `fmt.Errorf`, `fyne.CurrentApp`

### generateMandatoryMetaDataFieldsMap (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateMandatoryMetaDataFieldsMap(ownerDomainUuid string) mandatoryMetaDataFieldsMap map[string]bool`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generates a map containing all mandatory MetaDataFields for the selected DomainUuid ResponseMap map['GroupName-GroupItemName']bool
- Selector calls: `fmt.Sprintf`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
