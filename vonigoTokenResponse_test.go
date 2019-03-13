package vonigo

import (
	"testing"
)

func TestCanGetASecurityTokenFromVonigo(t *testing.T) {
	c := Config{
		AppVersion: "1",
		BaseURL:    "https://example.com",
		Password:   "secret",
		Username:   "secret",
		Company:    "company",
	}
	_ = Init(c)

	if securityToken == "" {
		t.Error("A security token should have been retrieved after package initialization")
	}
}

func TestTheTokenUrlCanBeGeneratedForUs(t *testing.T) {
	if !isInitialized {
		expected := "/api/v1/security/login/?appVersion=&company=&password=&userName="

		if expected != getTokenURL() {
			t.Error("We expected the default URL to be:'/api/v1/security/login/?appVersion=&company=&password=&userName='. Got back: '" + getTokenURL() + "'")
		}
	}

	c := Config{
		AppVersion: "1",
		BaseURL:    "https://example.com",
		Company:    "company",
		Password:   "secret",
		Username:   "username",
	}
	err := Init(c)
	if err != nil {
		t.Error("This test needs to be updated with the new required items for successful initialization. Error: " + err.Error())
	}

	expected := "https://example.com/api/v1/security/login/?appVersion=1&company=company&password=5ebe2294ecd0e0f08eab7690d2a6ee69&userName=username"
	if expected != getTokenURL() {
		t.Error("We expected the default URL to be:'https://example.com/api/v1/security/login/?appVersion=1&company=company&password=5ebe2294ecd0e0f08eab7690d2a6ee69&userName=username'. Got back: '" + getTokenURL() + "'")
	}
}
