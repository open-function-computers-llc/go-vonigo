package vonigo

import "testing"

func TestCanInitializeThePackage(t *testing.T) {
	if isInitialized {
		t.Error("the package should not be initialized to start with")
	}

	c := Config{}
	Init(c)

	if !isInitialized {
		t.Error("the package should be initialized after we call Init")
	}
}
