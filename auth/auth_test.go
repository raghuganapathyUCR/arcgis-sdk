package auth

import (
	"testing"
)

func TestApiKeyManager_GetToken(t *testing.T) {
	apiKey := "test-api-key"
	manager := NewApiKeyManager(apiKey)

	token, err := manager.GetToken("some-url")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if token != apiKey {
		t.Errorf("Expected token to be %s, got %s", apiKey, token)
	}
}

func TestApiKeyManager_GetToken_NoApiKey(t *testing.T) {
	manager := NewApiKeyManager("")

	_, err := manager.GetToken("some-url")
	if err == nil {
		t.Errorf("Expected an error when API key is not set")
	} else {
		if _, ok := err.(*ApiKeyError); !ok {
			t.Errorf("Expected an ApiKeyError, got %T", err)
		}
	}
}
