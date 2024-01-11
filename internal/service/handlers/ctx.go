package handlers

import (
	"context"
	"net/http"

	"github.com/rarimo/rarime-auth-svc/internal/jwt"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	jwtKey
	claimKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxJWT(issuer *jwt.JWTIssuer) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, jwtKey, issuer)
	}
}

func CtxClaim(claim *jwt.AuthClaim) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, claimKey, claim)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func JWT(r *http.Request) *jwt.JWTIssuer {
	return r.Context().Value(jwtKey).(*jwt.JWTIssuer)
}

func Claim(r *http.Request) *jwt.AuthClaim {
	return r.Context().Value(claimKey).(*jwt.AuthClaim)
}
