package token

import "os"

type TokenManager struct {
}

func (t *TokenManager) Token() string {
	return os.Getenv("BEARER_TOKEN")
}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}
