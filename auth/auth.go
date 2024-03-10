package auth

// enforce interface at compile time
var _ AuthenticationManager = (*ApiKeyManager)(nil)

type ApiKeyManager struct {
	Key string
	//TODO:  portal is the default portal for ArcGIS Online - need to expose this when we add support for enterprise
	portal string
}

func NewApiKeyManager(key string) *ApiKeyManager {
	return &ApiKeyManager{
		Key:    key,
		portal: "https://www.arcgis.com/sharing/rest",
	}
}

func (a *ApiKeyManager) GetToken(url string) (string, error) {
	if a.Key == "" {
		return "", NewApiKeyError("API key is not set")
	}
	return a.Key, nil
}
