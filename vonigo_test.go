package vonigo

import (
	"fmt"
	"testing"
)

func TestCanInitializeThePackage(t *testing.T) {
	//	logger := logrus.New()

	// Setting up a new config
	c := Config{}

	// Test to make sure config has valid properties
	err := Init(c)
	fmt.Println(err)

	if err == nil && err.Error() != "BaseURL is required." {
		t.Error("Expected error message: 'BaseURL is required.'")
	}

	c.BaseURL = "https://example.com"
	err = Init(c)

	if err == nil && err.Error() != "AppVersion is required." {
		t.Error("Expected error message: 'AppVersion is required'")
	}

	c.AppVersion = "1"
	err = Init(c)

	if err == nil && err.Error() != "Username is required." {
		t.Error("Expected Error: 'Username is required.'")
	}

	c.Username = "username"
	err = Init(c)

	if err == nil && err.Error() != "Password is required." {
		t.Error("Expected error message: 'Password is required.")
	}

	c.Password = "password"
	err = Init(c)

	// Check to ensure global variables have been set properly
	if baseURL != "https://example.com" {
		t.Error("Package variable 'baseURL' was not set properly")
	}
	if appVersion != "1" {
		t.Error("Package variable 'appVersion' was not set properly")
	}
	if username != "username" {
		t.Error("Package variable 'username' was not set properly")
	}

	if password != "password" {
		t.Error("Package variable 'password' was not set properly")
	}

}
