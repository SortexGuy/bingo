package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type User struct {
	ID            string
	Name          string
	Email         string
	EmailVerified bool
	Image         string
	CreatedAt     string
	UpdatedAt     string
}

func UserFromHeader(ctx context.Context, header http.Header) (User, error) {
	keyset, err := jwk.Fetch(ctx, "http://localhost:3000/api/auth/jwks")

	if err != nil {
		return User{}, err
	}

	token, err := jwt.ParseHeader(header, "Authorization", jwt.WithKeySet(keyset))

	if err != nil {
		return User{}, err
	}

	UserID, exist := token.Subject()

	if !exist {
		return User{}, errors.New("missing user id")
	}

	var name string
	token.Get("name", &name)

	var email string
	token.Get("email", &email)

	var emailVerified bool
	token.Get("emailVerified", &emailVerified)

	var image string
	token.Get("image", &image)

	var createdAt string
	token.Get("createdAt", &createdAt)

	var updatedAt string
	token.Get("updatedAt", &updatedAt)

	return User{
		ID:            UserID,
		Name:          name,
		Email:         email,
		EmailVerified: emailVerified,
		Image:         image,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}
