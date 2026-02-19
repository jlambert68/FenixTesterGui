# uiHelpers.go

## File Overview
- Path: `gui/uiHelpers.go`
- Package: `gui`
- Functions/Methods: `6`
- Imports: `1`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `SetDialAddressString`
- `SetLogger`

## Imports
- `github.com/sirupsen/logrus`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- None

## Functions and Methods
### SetDialAddressString (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) SetDialAddressString(dialAddress string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetDialAddressString Set the Dial Address, which was received from environment variables

### SetDialAddressString (method on `*GlobalUIServerStruct`)
- Signature: `func (*GlobalUIServerStruct) SetDialAddressString(dialAddress string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetDialAddressString Set the Dial Address, which was received from environment variables

### SetDialAddressString (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) SetDialAddressString(dialAddress string)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetDialAddressString Set the Dial Address, which was received from environment variables

### SetLogger (method on `*AvailableBuildingBlocksModelStruct`)
- Signature: `func (*AvailableBuildingBlocksModelStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same logger reference as is used by central part of system

### SetLogger (method on `*GlobalUIServerStruct`)
- Signature: `func (*GlobalUIServerStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same logger reference as is used by central part of system

### SetLogger (method on `*UIServerStruct`)
- Signature: `func (*UIServerStruct) SetLogger(logger *logrus.Logger)`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: SetLogger Set to use the same logger reference as is used by central part of system

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
