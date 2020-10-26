
package handler

import "campsite/packages/server/internal/service/role"

// We're defining such constants so that we don't make typos.
const (
	ID = "id"
	FILENAME = "filename"
)

var ADMIN_ONLY_ROLE_WHITELIST = []role.Role{role.ADMIN}