# testSuitesModel_allMandatoryFieldsHaveValues.go

## File Overview
- Path: `testSuites/testSuitesModel/testSuitesModel_allMandatoryFieldsHaveValues.go`
- Package: `testSuitesModel`
- Functions/Methods: `3`
- Imports: `2`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- None detected

## Imports
- `fmt`
- `log`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### checkIfAllMandatoryFieldsHaveValues (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) checkIfAllMandatoryFieldsHaveValues() (allMandatoryFieldsHaveValues bool, mandatoryFieldsHaveValuesNotificationText string)`
- Exported: `false`
- Control-flow features: `if`
- Doc: checkIfAllMandatoryFieldsHaveValues Checks if all mandatory fields in the TestSuite has gor any value
- Selector calls: `testSuiteModel.verifyMandatoryFieldsForMetaData`

### generateMandatoryMetaDataFieldsMap (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) generateMandatoryMetaDataFieldsMap() mandatoryMetaDataFieldsMap map[string]bool`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Generates a map containing all mandatory MetaDataFields for the selected DomainUuid ResponseMap map['GroupName-GroupItemName']bool
- Selector calls: `fmt.Sprintf`

### verifyMandatoryFieldsForMetaData (method on `*TestSuiteModelStruct`)
- Signature: `func (*TestSuiteModelStruct) verifyMandatoryFieldsForMetaData() (allMandatoryFieldsHasValues bool, mandatoryMetaDataFieldsString string)`
- Exported: `false`
- Control-flow features: `if, for/range, switch`
- Doc: Validates that all mandatory MetaData fields has values for specified DomainUuid
- Selector calls: `fmt.Sprintf`, `log.Fatalln`, `testSuiteModel.generateMandatoryMetaDataFieldsMap`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
