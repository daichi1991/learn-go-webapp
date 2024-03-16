package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/daichi1991/learn-go-webapp/apperrors"
	"google.golang.org/api/idtoken"
)

const (
	googleClientID = "YOUR"
)

type userNameKey struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequireAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]

		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequireAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err := apperrors.CannotMakeValidator.Wrap(err, "cannot make validator")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err := apperrors.Unauthorized.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err := apperrors.Unauthorized.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
