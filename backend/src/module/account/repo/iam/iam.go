package iam

import (
	"context"
	"fmt"
	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/ssoutil"

	"github.com/Nerzal/gocloak/v13"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Repo struct {
	client *gocloak.GoCloak
}

func New(client *gocloak.GoCloak) Repo {
	return Repo{client: client}
}

func getJwksUrl(realm string) string {
	jwksURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/certs",
		setting.KEYCLOAK_URL,
		realm,
	)
	return jwksURL
}

func (r Repo) GetAuthUrl(
	realm string,
	clientId string,
	state ctype.Dict,
) string {
	stateStr := ssoutil.EncodeState(state)
	keycloakAuthURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/auth?client_id=%s&response_type=code&redirect_uri=%s&scope=openid profile email&state=%s",
		setting.KEYCLOAK_URL,
		realm,
		clientId,
		setting.KEYCLOAK_DEFAULT_REDIRECT_URI,
		stateStr,
	)
	return keycloakAuthURL
}

func (r Repo) GetLogoutUrl(
	realm string,
	clientId string,
) string {
	keycloakAuthURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/logout?client_id=%s&post_logout_redirect_uri=%s",
		setting.KEYCLOAK_URL,
		realm,
		clientId,
		setting.KEYCLOAK_DEFAULT_POST_LOGOUT_URI,
	)
	return keycloakAuthURL
}

func (r Repo) ValidateCallback(
	ctx context.Context,
	realm string,
	clientId string,
	clientSecret string,
	code string,
) (ssoutil.TokensAndClaims, error) {
	var result ssoutil.TokensAndClaims
	localizer := localeutil.Get()

	if code == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.AuthorizationCodeNotFound,
		})
		return result, errutil.New("", []string{msg})
	}

	// Exchange the code for tokens
	token, err := r.client.GetToken(ctx, realm, gocloak.TokenOptions{
		ClientID:     gocloak.StringP(clientId),
		ClientSecret: gocloak.StringP(clientSecret),
		RedirectURI:  gocloak.StringP(setting.KEYCLOAK_DEFAULT_REDIRECT_URI),
		Code:         gocloak.StringP(code),
		GrantType:    gocloak.StringP("authorization_code"),
	})

	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotExchangeAuthorizationCode,
		})
		return result, errutil.New("", []string{msg})
	}

	idToken := token.IDToken
	accesToken := token.AccessToken
	refreshToken := token.RefreshToken

	userInfo, err := r.ValidateToken(idToken, realm)
	if err != nil {
		return result, err
	}

	result = ssoutil.TokensAndClaims{
		AccessToken:  accesToken,
		RefreshToken: refreshToken,
		Realm:        realm,
		UserInfo:     userInfo,
	}
	return result, nil
}

// TODO: Implement checking kid
func (r Repo) ValidateToken(tokenStr string, realm string) (ssoutil.UserInfo, error) {
	localizer := localeutil.Get()
	result := ssoutil.UserInfo{}
	jwksURL := getJwksUrl(realm)

	// Fetch the JWKS (public keys)
	keySet, err := ssoutil.GetKeySet(jwksURL)
	if err != nil {
		return result, err
	}

	// Parse the JWT token to extract headers and claims
	token, err := jwt.Parse([]byte(tokenStr), jwt.WithKeySet(keySet))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToParseToken,
		})
		return result, errutil.New("", []string{msg})
	}

	// Check if the token is expired by inspecting the "exp" claim
	if err := jwt.Validate(token, jwt.WithAcceptableSkew(0)); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.TokenHasExpired,
		})
		return result, errutil.New("", []string{msg})
	}

	// If verification is successful, print the claims
	claims := token.PrivateClaims()
	result = ssoutil.UserInfo{
		Uid:       claims["preferred_username"].(string),
		Email:     claims["email"].(string),
		FirstName: claims["given_name"].(string),
		LastName:  claims["family_name"].(string),
	}

	return result, nil
}

func (r Repo) RefreshToken(
	ctx context.Context,
	realm string,
	refreshToken string,
) (ssoutil.TokensAndClaims, error) {
	var result ssoutil.TokensAndClaims
	localizer := localeutil.Get()
	if refreshToken == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.RefreshTokenNotFound,
		})
		return result, errutil.New("", []string{msg})
	}

	// Exchange the refresh token for new tokens
	token, err := r.client.RefreshToken(
		ctx,
		refreshToken,
		setting.KEYCLOAK_DEFAULT_CLIENT_ID,
		setting.KEYCLOAK_DEFAULT_CLIENT_SECRET,
		realm,
	)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotExchangeRefreshToken,
		})
		return result, errutil.New("", []string{msg})
	}

	idToken := token.IDToken
	accesToken := token.AccessToken
	refreshToken = token.RefreshToken

	userInfo, err := r.ValidateToken(idToken, realm)
	if err != nil {
		return result, err
	}

	result = ssoutil.TokensAndClaims{
		AccessToken:  accesToken,
		RefreshToken: refreshToken,
		Realm:        realm,
		UserInfo:     userInfo,
	}
	return result, nil
}
