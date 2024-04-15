package sharedCode

import (
	"github.com/rs/zerolog"
	"os"
)

type LogMessage struct {
	Level       zerolog.Level
	MessageUuid string
	Message     string
	Context     map[string]interface{}
}

const logChannelSize = 1000
const logChannelWarningLeverl = 800

// Create the log channel where log messages arrives
var logChannel = make(chan LogMessage, 1000) // Adjust buffer size as needed

// InitiateLoggerEngine
// Initiate and start up the "logger engine"
func InitiateLoggerEngine(filePathName string) {
	go startLoggerEngine()
}

// Run the logger engine
func startLoggerEngine() {

	// Initiate the logger
	var log zerolog.Logger
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Loop forever
	for {

		// Wait for a message on channel
		msg := <-logChannel

		// Depending on the log level, process the message accordingly
		switch msg.Level {

		// Log DebugLevel
		case zerolog.DebugLevel:
			log.Debug().Fields(msg.Context).Msg(msg.Message)

		// Log InfoLevel
		case zerolog.InfoLevel:

			log.Info().Fields(msg.Context).Msg(msg.Message)

		// Log WarnLevel
		case zerolog.WarnLevel:
			log.Warn().Fields(msg.Context).Msg(msg.Message)

		// Log ErrorLevel
		case zerolog.ErrorLevel:
			log.Error().Fields(msg.Context).Msg(msg.Message)

		// Log FatalLevel
		case zerolog.FatalLevel:
			log.Fatal().Fields(msg.Context).Msg(msg.Message)

		default:
			id := "90c27674-6496-4bf0-b516-f49e00fa1795"
			logMessage := "Unhandled 'loglevel'"
			messageVariable := map[string]interface{}{"msg.Level": msg.Level}
			LogFatalLevel(id, logMessage, messageVariable)

		}
	}
}

// LogDebugLevel
// Puts one log message of type 'DebugLevel' on the log channel
func LogDebugLevel(messageUuid string, message string, context map[string]interface{}) {
	logChannel <- LogMessage{
		Level:       zerolog.DebugLevel,
		MessageUuid: messageUuid,
		Message:     message,
		Context:     context}
}

// LogInfoLevel
// Puts one log message of type 'InfoLevel' on the log channel
func LogInfoLevel(messageUuid string, message string, context map[string]interface{}) {
	logChannel <- LogMessage{
		Level:       zerolog.InfoLevel,
		MessageUuid: messageUuid,
		Message:     message,
		Context:     context}
}

// LogWarningLevel
// Puts one log message of type 'WarnLevel' on the log channel
func LogWarningLevel(messageUuid string, message string, context map[string]interface{}) {
	logChannel <- LogMessage{
		Level:       zerolog.WarnLevel,
		MessageUuid: messageUuid,
		Message:     message,
		Context:     context}
}

// LogErrorLevel
// Puts one log message of type 'ErrorLevel' on the log channel
func LogErrorLevel(messageUuid string, message string, context map[string]interface{}) {
	logChannel <- LogMessage{
		Level:       zerolog.ErrorLevel,
		MessageUuid: messageUuid,
		Message:     message,
		Context:     context}
}

// LogFatalLevel
// Puts one log message of type 'FatalLevel' on the log channel
func LogFatalLevel(messageUuid string, message string, context map[string]interface{}) {
	logChannel <- LogMessage{
		Level:       zerolog.FatalLevel,
		MessageUuid: messageUuid,
		Message:     message,
		Context:     context}
}
