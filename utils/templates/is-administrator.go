package templatesutils

import (
	"fmt"
	"forum/models"
	"slices"
)

func isAdministrator(v interface{}) bool {
	user, ok := v.(models.User)
	fmt.Println(user.Roles)
	if !ok {
		return false
	}
	return slices.Contains(user.Roles, "admin")
}
