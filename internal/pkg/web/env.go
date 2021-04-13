package web

import "os"

// Server environment variables
const (
	// M2W_PORT variable determines the port to start the server on
	M2W_PORT = "M2W_PORT"

	// M2W_LOG_LEVEL variable determines the log level across the service
	M2W_LOG_LEVEL = "M2W_LOG_LEVEL"
)

// getPortFromEnv returns the M2W_PORT environment variable.
// For compatibility with some platforms (e.g. Heroku), it falls back to PORT if M2W_PORT is empty.
func getPortFromEnv() string {
	port, ok := os.LookupEnv(M2W_PORT)
	if !ok {
		port = os.Getenv("PORT")
	}
	return port
}

// getLogLevelFromEnv returns the M2W_LOG_LEVEL environment variable.
func getLogLevelFromEnv() string {
	return os.Getenv(M2W_LOG_LEVEL)
}

// hasEnv returns true if there exists an environment variable 'key'
func hasEnv(key string) bool {
	return os.Getenv(key) != ""
}
