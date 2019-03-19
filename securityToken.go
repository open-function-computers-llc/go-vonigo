package vonigo

var securityToken string

func hasSecurityToken() bool {
	if securityToken == "" {
		return false
	}

	return true
}
