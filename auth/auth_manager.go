package auth

// ITokenRequestOptions represents the options for a token request.
// TODO: add enterprise support
// type ITokenRequestOptions struct {

// }

// IAuthenticationManager defines the interface for authentication managers.
type AuthenticationManager interface {
	// GetToken(url string, requestOptions *ITokenRequestOptions) (string, error)
	GetToken(url string) (string, error)
	// GetDomainCredentials(url string) (string, error)
	// CanRefresh() bool
	// RefreshCredentials(requestOptions *ITokenRequestOptions) error
}
