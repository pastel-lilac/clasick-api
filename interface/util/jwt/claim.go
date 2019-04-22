package jwt

type Claim struct {
	Issuer         string
	Subject        string
	Audience       string
	ExpirationTime uint32
	NotBefore      uint32
	IssuedAt       uint32
	JWTID          string
	Type           string
}
