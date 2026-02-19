# testCaseModel_verifyMandatoryFieldsForMetaData.go

## File Overview
- Path: `testCase/testCaseModel/testCaseModel_verifyMandatoryFieldsForMetaData.go`
- Package: `testCaseModel`
- Generated: `2026-02-19T14:23:17+01:00`
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
### generateMandatoryMetaDataFieldsMap (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) generateMandatoryMetaDataFieldsMap(ownerDomainUuid string) mandatoryMetaDataFieldsMap map[string]bool`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generates a map containing all mandatory MetaDataFields for the selected DomainUuid ResponseMap map['GroupName-GroupItemName']bool
- External calls: `fmt.Sprintf`

### verifyMandatoryFieldsForMetaData (method on `*TestCasesModelsStruct`)
- Signature: `func (*TestCasesModelsStruct) verifyMandatoryFieldsForMetaData(domainUuid string, currentTestCasePtr *TestCaseModelStruct, shouldBeSaved bool) err error`
- Exported: `false`
- Control-flow features: `if, for/range, switch, returns error`
- Doc: Validates that all mandatory MetaData fields has values for specified DomainUuid
- External calls: `fmt.Errorf`, `fmt.Sprintf`, `fyne.CurrentApp`, `log.Fatalln`, `testCaseModel.generateMandatoryMetaDataFieldsMap`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
