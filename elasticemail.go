package elasticemail

const apiEndpoint = `https://api.elasticemail.com/v2/email/`

var elasticEmailAPIKeyEmailEnvVarName = "ELASTICEMAIL_APIKEY"

// SetAPIKeyEmailEnviVarName replaces the default "ELASTICEMAIL_APIKEY" environment variable name with a given custom
// It DOESN'T change the VALUE of environment key
func SetAPIKeyEmailEnviVarName(envVarName string) {
	elasticEmailAPIKeyEmailEnvVarName = envVarName
}
