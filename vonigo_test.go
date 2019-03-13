package vonigo

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestCanInitializeThePackage(t *testing.T) {
	if isInitialized {
		t.Error("the package should not be initialized to start with")
	}

	logger := logrus.New()
	c := Config{
		Logger: logger,
	}
	Init(c)

	if !isInitialized {
		t.Error("the package should be initialized after we call Init")
	}
}
