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

	if err == nil && err.Error() != "BaseURL is required" {
		t.Error("Expected error message: 'BaseURL is required'")
	}

	c.BaseURL = "https://example.com"
	err = Init(c)

	if err == nil && err.Error() != "AppVersion is required" {
		t.Error("Expected error message: 'AppVersion is required'")
	}

	c.AppVersion = "1"
	err = Init(c)

	if err == nil && err.Error() != "Username is required" {
		t.Error("Expected Error: 'Username is required'")
	}

	c.Username = "username"
	err = Init(c)

	if err == nil && err.Error() != "Password is required" {
		t.Error("Expected error message: 'Password is required'")
	}

	c.Password = "password"
	err = Init(c)

	if err == nil && err.Error() != "Company is required" {
		t.Error("Expected error message: 'Company is required'")
	}

	c.Company = "company"
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

	if password != "5f4dcc3b5aa765d61d8327deb882cf99" {
		t.Error("Package variable 'password' was not set properly, or it got the wrong MD5 hash")
	}

	if company != "company" {
		t.Error("Package variable 'company' was not set properly")
	}

}

// func TestCanGetClients(t *testing.T) {
// 	// Setting up a new config
// 	c := Config{}

// 	// Test to make sure config has valid properties
// 	err := Init(c)
// 	if err != nil {
// 		t.Error("Failed to initialize")
// 	}

// 	clients, err := GetClients()
// 	if err != nil {
// 		t.Error("Getting clients failed")
// 	}

// 	if clients.len <= 0 {
// 		t.Error("Failed to return any clients")
// 	}

// }
