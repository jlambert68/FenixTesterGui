# logEngine.go

## File Overview
- Path: `common_code/logEngine.go`
- Package: `sharedCode`
- Functions/Methods: `7`
- Imports: `2`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `InitiateLoggerEngine`
- `LogDebugLevel`
- `LogErrorLevel`
- `LogFatalLevel`
- `LogInfoLevel`
- `LogWarningLevel`

## Imports
- `github.com/rs/zerolog`
- `os`

## Declared Types
- `LogMessage`

## Declared Constants
- `logChannelSize`
- `logChannelWarningLeverl`

## Declared Variables
- `logChannel`

## Functions and Methods
### InitiateLoggerEngine
- Signature: `func InitiateLoggerEngine(filePathName string)`
- Exported: `true`
- Control-flow features: `go`
- Doc: InitiateLoggerEngine Initiate and start up the "logger engine"
- Internal calls: `startLoggerEngine`

### LogDebugLevel
- Signature: `func LogDebugLevel(messageUuid string, message string, context map[string]interface{})`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LogDebugLevel Puts one log message of type 'DebugLevel' on the log channel

### LogErrorLevel
- Signature: `func LogErrorLevel(messageUuid string, message string, context map[string]interface{})`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LogErrorLevel Puts one log message of type 'ErrorLevel' on the log channel

### LogFatalLevel
- Signature: `func LogFatalLevel(messageUuid string, message string, context map[string]interface{})`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LogFatalLevel Puts one log message of type 'FatalLevel' on the log channel

### LogInfoLevel
- Signature: `func LogInfoLevel(messageUuid string, message string, context map[string]interface{})`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LogInfoLevel Puts one log message of type 'InfoLevel' on the log channel

### LogWarningLevel
- Signature: `func LogWarningLevel(messageUuid string, message string, context map[string]interface{})`
- Exported: `true`
- Control-flow features: `none detected`
- Doc: LogWarningLevel Puts one log message of type 'WarnLevel' on the log channel

### startLoggerEngine
- Signature: `func startLoggerEngine()`
- Exported: `false`
- Control-flow features: `for/range, switch`
- Doc: Run the logger engine
- Internal calls: `LogFatalLevel`
- Selector calls: `log.Debug`, `log.Error`, `log.Fatal`, `log.Info`, `log.Warn`, `zerolog.New`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
