package validation

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
	"github.com/korniykom/Chatters-Backend-Go/internal/service"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,32}$`)

func ValidateUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return service.ErrInvalidUsername
	}

	return nil
}

func ValidateEmail(email  string) error {

	address, err := mail.ParseAddress(email)
	if err != nil {
		return service.ErrInvalidEmail
	}

	if address.Address != email {
		return service.ErrInvalidEmail
	}

	return nil
}


func NormalizeRegistrationRequest(
	req domain.RegisterRequest,
) domain.RegisterRequest {
	return domain.RegisterRequest{
		Username: strings.TrimSpace(req.Username),
		Email:    strings.ToLower(strings.TrimSpace(req.Email)),
	}
}