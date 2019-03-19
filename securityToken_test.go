package vonigo

import "testing"

func TestHasSecurityToken(t *testing.T) {
	if hasSecurityToken() {
		t.Error("This should return 'false' when securityToken is not set")
	}

	securityToken = "faketoken"

	if !hasSecurityToken() {
		t.Error("This should return true when the global securityToken is set.")
	}
}
