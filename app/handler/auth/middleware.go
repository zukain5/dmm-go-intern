package auth

import (
	"context"
	"net/http"
	"strings"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
)

var contextKey = new(struct{})

// Auth by header
func Middleware(ar repository.Account) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// ヘッダーから Username を取り出すだけの超安易な認証
			a := r.Header.Get("Authentication")
			pair := strings.SplitN(a, " ", 2)
			if len(pair) < 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			authType := pair[0]
			if !strings.EqualFold(authType, "username") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			username := pair[1]
			if account, err := ar.FindByUsername(ctx, username); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else if account == nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			} else {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKey, account)))
			}
		})
	}
}

// Read Account data from authorized request
func AccountOf(r *http.Request) *object.Account {
	if cv := r.Context().Value(contextKey); cv == nil {
		return nil

	} else if account, ok := cv.(*object.Account); !ok {
		return nil

	} else {
		return account

	}
}
