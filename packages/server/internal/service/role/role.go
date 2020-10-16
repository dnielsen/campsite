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
