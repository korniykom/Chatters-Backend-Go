package validation

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/korniykom/Chatters-Backend-Go/internal/domain"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,32}$`)

func ValidateUsername(username string) bool {
	return usernameRegex.MatchString(username)
}

func ValidateEmail(email string) bool {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return address.Address == email
}

func NormalizeRegistrationRequest(
	req domain.RegisterRequest,
) domain.RegisterRequest {
	return domain.RegisterRequest{
		Username: strings.TrimSpace(req.Username),
		Email:    strings.ToLower(strings.TrimSpace(req.Email)),
	}
}