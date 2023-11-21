package auth

type Authenticator struct {
	Domain       string
	Audience     string
	ClientId     string
	ClientSecret string
}

func New() Authenticator {
	return Authenticator{
		//TODO: once auth0 tenant is set up, set these env vars between envs
		// Domain:       os.Getenv(""),
		// Audience:     os.Getenv(""),
		// ClientId:     os.Getenv(""),
		// ClientSecret: os.Getenv(""),
	}
}
