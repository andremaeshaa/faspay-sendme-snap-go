package snap

// Version is the current version of the SDK
const Version = "1.0.0"

// UserAgent is the user agent string sent with API requests
func UserAgent() string {
	return "FaspaySendMeSnapGo/" + Version
}
