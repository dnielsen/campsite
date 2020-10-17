package role

import "errors"

type Role string

const (
	ADMIN Role = "ADMIN"
	USER = "USER"
)

func (r Role) IsValid() error {
	switch r {
	case ADMIN, USER:
		return nil
	}
	return errors.New("invalid role")
}

func Contains(roles []Role, searchedRole Role) bool {
	for _, r := range roles {
		if r == searchedRole {
			return true
		}
	}
	return false
}
