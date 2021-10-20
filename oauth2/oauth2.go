package oauth2

//基本配置
type OAuthConfig struct {
	// ClientID is the application's ID.
	// 应用ID
	ClientID string

	// ClientSecret is the application's secret.
	// 应用密钥
	ClientSecret string

	// Endpoint contains the resource server's token endpoint
	// URLs. These are constants specific to each server and are
	// often available via site-specific packages, such as
	// google.Endpoint or github.Endpoint.
	// 获取授权地址
	AuthURL string
	// 获取token地址
	TokenURL string

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	// 回调地址
	RedirectURL string

	// 需要申请的服务权限
	// Scope specifies optional requested permissions.
	Scopes []string
}
