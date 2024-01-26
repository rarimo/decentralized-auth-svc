package cookies

import (
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
)

type Cookies struct {
	Domain string `fig:"domain.required"`
}

func SetTokensCookies(w http.ResponseWriter, access, refresh, domain string) {
	refreshCookie := &http.Cookie{
		Name:     jwt.RefreshTokenType.String(),
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Domain:   domain,
	}

	http.SetCookie(w, refreshCookie)

	accessCookie := &http.Cookie{
		Name:     jwt.AccessTokenType.String(),
		Value:    access,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Domain:   domain,
	}

	http.SetCookie(w, accessCookie)
}

func ClearTokensCookies(w http.ResponseWriter, domain string) {
	refreshCookie := &http.Cookie{
		Name:     jwt.RefreshTokenType.String(),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Domain:   domain,
	}

	http.SetCookie(w, refreshCookie)

	accessCookie := &http.Cookie{
		Name:     jwt.AccessTokenType.String(),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Domain:   domain,
	}

	http.SetCookie(w, accessCookie)
}
