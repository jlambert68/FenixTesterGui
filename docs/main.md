# main.go

## File Overview
- Path: `main.go`
- Package: `main`
- Functions/Methods: `2`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `main`

## Imports
- `FenixTesterGui/common_code`
- `embed`
- `fmt`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTemplateToThisDomain/version_1_0`
- `github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0`
- `os`
- `os/user`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `embededfenixIcon`

## Functions and Methods
### main
- Signature: `func main()`
- Exported: `false`
- Control-flow features: `if`
- Internal calls: `fenixGuiBuilderServerMain`
- Selector calls: `user.Current`, `fmt.Println`, `strings.ReplaceAll`, `testInstruction_SendTestDataToThisDomain_version_1_0.Initate_TestInstruction_FenixSentToUsersDomain_SendTestDataToThisDomain`, `testInstruction_SendTemplateToThisDomainversion_1_0.Initate_TestInstruction_FenixSentToUsersDomain_FenixOwnedSendTemplateToThisDomain`

### onExit
- Signature: `func onExit()`
- Exported: `false`
- Control-flow features: `none detected`
- Doc: SysTray Application - StartUp systray.SetIcon(embededfenixIcon)
- Selector calls: `os.Exit`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
