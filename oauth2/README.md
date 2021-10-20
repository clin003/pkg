# oauth2

第三方 OAuth2.0 登录golang版本的简单实现


# 安装

`go get gitee.com/lyhuilin/pkg/oauth2`
`go get gitee.com/lyhuilin/pkg/oauth2/qq`

# emmm

```golang
	type OAuthConfg struct {
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
```

# 使用


## qq

```golang

	//QQ互联应用信息
	qqConf := &OAuthConfg{
		ClientID:     "xxx",
		ClientSecret: "xxx",
		RedirectURL:  "https://auth.nav.xin/oauth/qq",
		Scopes:       []string{"get_user_info", "getUnionId"},
		AuthURL:      "https://graph.qq.com/oauth2.0/authorize",
		TokenURL:     "https://graph.qq.com/oauth2.0/token",
	}
	qqAuth := NewAuth(qqConf)
	qqAuth.AuthCodeURL("state") //获取第三方登录地址
	code := "xxx"

	//	回调页收的code 获取token相关信息
	if qqToken, err := qqAuth.Token(code); err != nil {
		//t.Log(err)
	} else {
		//t.Log("qqToken", qqToken)
		//t.Log("AccessToken", qqToken.AccessToken)
		// 获取openInfo信息
		if qqOpenIDInfo, err := qqAuth.OpenInfo(qqToken.AccessToken, true); err != nil {
			//t.Log(err)
		} else {
			//t.Log("qqOpenIDInfo", qqOpenIDInfo)
			//t.Log("OpenID", qqOpenIDInfo.OpenID)
			//t.Log("UnionID", qqOpenIDInfo.UnionID)

			// 获取userinfo 信息
			if qqUserInfo, err := qqAuth.UserInfo(qqToken.AccessToken, qqOpenIDInfo.OpenID); err != nil {
				//t.Log(err)
			} else if qqUserInfo.Ret < 0 {
				//t.Log("qqUserInfo", qqUserInfo)
				//t.Log(qqUserInfo.Msg)
			} else {
				//t.Log("qqUserInfo", qqUserInfo)
				//t.Log(qqUserInfo.NickName)
				//t.Log(qqUserInfo.Figureurl_qq_1)
			}

		}
	}

```


