package handlers

import (
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"net/http"
	"strings"
)

type AuthenticationMw struct {
	claimsDomain domain.JwtClaimsInterface
	apiKey       string
}
type AuthenticationMechanismInterface interface {
	JWTAuthenticationMW(next http.Handler) http.Handler
}

// TODO: take api key from env
func NewAuthenticationMw(claims domain.JwtClaimsInterface, apiKey string) *AuthenticationMw {
	return &AuthenticationMw{claimsDomain: claims, apiKey: apiKey}
}

func (a *AuthenticationMw) JWTAuthenticationMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := a.claimsDomain.GetClaimsFromAccessToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r = r.WithContext(a.claimsDomain.SetAccessJWTClaimsContext(r.Context(), claims))
		next.ServeHTTP(w, r)
	})
}

func (a *AuthenticationMw) APIKeyAuthenticationMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		clientApiKey := strings.TrimPrefix(authHeader, "Bearer ")

		//hashedApiKey := sha256.Sum256([]byte(a.apiKey))
		//hashedApiKeyAsString := string(hashedApiKey[:])

		if a.apiKey != clientApiKey {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
