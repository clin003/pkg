package qq

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"

	. "gitee.com/lyhuilin/pkg/oauth2"
)

//@ qq 结构 ------------------------------------------------- start
type OAuth struct {
	Conf *OAuthConfg
}

type Token struct {
	// 授权令牌，Access_Token。
	AccessToken string `json:"access_token"` //授权令牌，Access_Token。

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string `json:"token_type,omitempty"`

	// 在授权自动续期步骤中，获取新的Access_Token时需要提供的参数。
	// 注：refresh_token仅一次有效
	RefreshToken string `json:"refresh_token,omitempty"`

	// 该access token的有效期，单位为秒。
	Expiry string `json:"expires_in,omitempty"`

	// raw optionally contains extra metadata from the server
	// when updating a token.
	raw interface{}
}
type OpenIDInfo struct {
	Client_ID string `json:"client_id"`
	OpenID    string `json:"openid"`
	UnionID   string `json:"unionid",omitempty`
}
type QQUserInfo struct {
	// 返回码
	Ret int `json:"ret"`
	// 如果ret<0，会有相应的错误信息提示，返回数据全部用UTF-8编码。
	Msg string `json:"msg",omitempty`
	// 用户在QQ空间的昵称。
	NickName string `json:"nickname",omitempty`
	// 大小为30×30像素的QQ空间头像URL。
	Figureurl string `json:"figureurl",omitempty`
	// 大小为50×50像素的QQ空间头像URL。
	Figureurl_1 string `json:"figureurl_1",omitempty`
	// 大小为100×100像素的QQ空间头像URL。
	Figureurl_2 string `json:"figureurl_2",omitempty`
	// 大小为40×40像素的QQ头像URL。
	Figureurl_qq_1 string `json:"figureurl_qq_1"`
	// 大小为100×100像素的QQ头像URL。需要注意，不是所有的用户都拥有QQ的100x100的头像，但40x40像素则是一定会有。
	Figureurl_qq_2 string `json:"figureurl_qq_2",omitempty`
	// 性别。 如果获取不到则默认返回"男"
	Gender string `json:"gender",omitempty`
}

//@ qq 结构 ------------------------------------------------- end

//获取登录地址
func (c *OAuth) AuthCodeURL(state string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Conf.AuthURL)

	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.Conf.ClientID},
	}
	if c.Conf.RedirectURL != "" {
		v.Set("redirect_uri", c.Conf.RedirectURL)
	}
	if len(c.Conf.Scopes) > 0 {
		v.Set("scope", strings.Join(c.Conf.Scopes, ","))
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		v.Set("state", state)
	}
	if strings.Contains(c.Conf.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
	// return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + e.Conf.Appid + "&redirect_uri=" + e.Conf.Rurl + "&state=" + state
}

//获取Token地址
func (c *OAuth) getTokenURL(code string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Conf.TokenURL)

	v := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {c.Conf.ClientID},
		"client_secret": {c.Conf.ClientSecret},
		"fmt":           {"json"},
	}
	if c.Conf.RedirectURL != "" {
		v.Set("redirect_uri", c.Conf.RedirectURL)
	}
	if code != "" {
		v.Set("code", code)
	}

	if strings.Contains(c.Conf.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

//获取token
func (c *OAuth) Token(code string) (*Token, error) {
	apiUrl := c.getTokenURL(code)
	// fmt.Println(apiUrl)
	body, err := HttpGetByte(apiUrl)
	if err != nil {
		return nil, err
	}

	// body, err := util.GetUrlToByte(tokenURL)
	// if err != nil {
	// 	return nil, err
	// }

	var q Token
	if err := json.Unmarshal(body, &q); err != nil {
		return nil, err
	}
	return &q, nil
}

// https://wiki.connect.qq.com/unionid%E4%BB%8B%E7%BB%8D
func (c *OAuth) getOpenInfoURL(access_token string, isUnion bool) string {
	var buf bytes.Buffer
	buf.WriteString("https://graph.qq.com/oauth2.0/me")

	v := url.Values{
		"access_token": {access_token},
		"fmt":          {"json"},
	}
	if isUnion {
		v.Set("unionid", "1")
	}

	if strings.Contains(c.Conf.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

//获取第三方id
func (c *OAuth) OpenInfo(access_token string, isUnion bool) (*OpenIDInfo, error) {
	apiUrl := c.getOpenInfoURL(access_token, isUnion)

	// fmt.Println(apiUrl)
	body, err := HttpGetByte(apiUrl)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	var q OpenIDInfo
	if err := json.Unmarshal(body, &q); err != nil {
		return nil, err
	} else {
		return &q, nil
	}
}

// https://wiki.connect.qq.com/get_user_info
func (c *OAuth) getUserInfoURL(access_token, openid string) string {
	var buf bytes.Buffer
	buf.WriteString("https://graph.qq.com/user/get_user_info")

	v := url.Values{
		"access_token":       {access_token},
		"oauth_consumer_key": {c.Conf.ClientID},
		"openid":             {openid},
		// "fmt":                {"json"},
	}

	if strings.Contains(c.Conf.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

//获取第三方用户信息
func (c *OAuth) UserInfo(access_token string, openid string) (*QQUserInfo, error) {
	apiUrl := c.getUserInfoURL(access_token, openid)

	// fmt.Println(apiUrl)
	body, err := HttpGetByte(apiUrl)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	var q QQUserInfo
	if err := json.Unmarshal(body, &q); err != nil {
		return nil, err
	} else {
		return &q, nil
	}
}

//构造方法
func NewAuth(config *OAuthConfg) *OAuth {
	return &OAuth{
		Conf: config,
	}
}
