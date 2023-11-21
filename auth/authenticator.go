package auth

type Authenticator struct {
	Domain       string
	Audience     string
	ClientId     string
	ClientSecret string
}

func New() Authenticator {
	return Authenticator{}
}
