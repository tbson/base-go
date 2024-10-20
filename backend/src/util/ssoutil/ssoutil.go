package ssoutil

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"

	"github.com/Nerzal/gocloak/v13"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type CallbackResult struct {
	AccessToken  string
	RefreshToken string
	Cliams       ctype.Dict
}

func getJwksUrl(realm string) string {
	jwksURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/certs",
		setting.KEYCLOAK_URL,
		realm,
	)
	return jwksURL
}

func getKeySet(jwksURL string) (jwk.Set, error) {
	localizer := localeutil.Get()
	ctx := context.Background()
	keySet, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.CannotExchangeAuthorizationCode,
		})
		return nil, errutil.New("", []string{msg})
	}
	return keySet, nil
}

func encodeState(stateData ctype.Dict) string {
	jsonData, _ := json.Marshal(stateData)
	return base64.URLEncoding.EncodeToString(jsonData)
}

func decodeState(stateStr string) (ctype.Dict, error) {
	jsonData, err := base64.URLEncoding.DecodeString(stateStr)
	if err != nil {
		return nil, err
	}
	var stateData ctype.Dict
	err = json.Unmarshal(jsonData, &stateData)
	if err != nil {
		return nil, err
	}
	return stateData, nil
}

func GetAuthUrl(
	realm string,
	clientId string,
	redirectUri string,
	state ctype.Dict,
) string {
	stateStr := encodeState(state)
	keycloakAuthURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/auth?client_id=%s&response_type=code&redirect_uri=%s&scope=openid profile email&state=%s",
		setting.KEYCLOAK_URL,
		realm,
		clientId,
		redirectUri,
		stateStr,
	)
	return keycloakAuthURL
}

func ValidateCallback(
	ctx context.Context,
	code string,
	stateStr string,
) (CallbackResult, error) {
	var result CallbackResult
	localizer := localeutil.Get()
	// Decode the state
	stateData, err := decodeState(stateStr)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.InvalidState,
		})
		return result, errutil.New("", []string{msg})
	}
	realm, ok := stateData["realm"].(string)
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.NoRealmFound,
		})
		return result, errutil.New("", []string{msg})
	}

	if code == "" {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.AuthorizationCodeNotFound,
		})
		return result, errutil.New("", []string{msg})
	}

	// Exchange the code for tokens
	client := gocloak.NewClient(setting.KEYCLOAK_API_URL)
	token, err := client.GetToken(ctx, realm, gocloak.TokenOptions{
		ClientID:     gocloak.StringP(setting.KEYCLOAK_DEFAULT_CLIENT_ID),
		ClientSecret: gocloak.StringP(setting.KEYCLOAK_DEFAULT_CLIENT_SECRET),
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

	claims, err := ValidateToken(idToken, realm)
	if err != nil {
		return result, err
	}

	result = CallbackResult{
		AccessToken:  accesToken,
		RefreshToken: refreshToken,
		Cliams:       claims,
	}
	return result, nil
}

func ValidateToken(tokenStr string, realm string) (ctype.Dict, error) {
	localizer := localeutil.Get()
	var result ctype.Dict
	jwksURL := getJwksUrl(realm)

	// Fetch the JWKS (public keys)
	keySet, err := getKeySet(jwksURL)
	if err != nil {
		return result, err
	}

	// Parse the JWT token to extract headers and claims
	token, err := jwt.Parse([]byte(tokenStr))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToParseToken,
		})
		return result, errutil.New("", []string{msg})
	}

	// Get the `kid` from the JWT header using jwt.Get()
	kid, ok := token.Get("kid")
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.NoKidFieldInJWTTokenHeader,
		})
		return result, errutil.New("", []string{msg})
	}

	// Look up the correct key from the JWKS using the `kid`
	key, found := keySet.LookupKeyID(kid.(string))
	if !found {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.UnableToFindKeyWithKid,
		})
		return result, errutil.New("", []string{msg})
	}

	// Extract the raw RSA public key
	var rawKey interface{}
	if err := key.Raw(&rawKey); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToCreateRawKey,
		})
		return result, errutil.New("", []string{msg})
	}

	// Ensure it's an RSA public key
	rsaKey, ok := rawKey.(*rsa.PublicKey)
	if !ok {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.ExpectedRSAKey,
		})
		return result, errutil.New("", []string{msg})
	}

	// Parse and verify the JWT token with the RSA public key, using jwa.RS256
	verifiedToken, err := jwt.Parse([]byte(tokenStr), jwt.WithKey(jwa.RS256, rsaKey))
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToVerifyToken,
		})
		return result, errutil.New("", []string{msg})
	}

	// Check if the token is expired by inspecting the "exp" claim
	if err := jwt.Validate(verifiedToken, jwt.WithAcceptableSkew(0)); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.TokenHasExpired,
		})
		return result, errutil.New("", []string{msg})
	}

	// If verification is successful, print the claims
	result = verifiedToken.PrivateClaims()

	return result, nil
}
