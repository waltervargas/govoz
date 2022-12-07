package govoz

const (
	sampleRate = 44100 / 2
	seconds    = 2
)

func RunAs(mode string, url string) error {
	// run as a server or as a client?
	switch mode {
	case "client":
		return runHTTPClient(url)
	default:
		return runHTTPServer()
	}
}
