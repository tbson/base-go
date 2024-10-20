package infra

import (
	"fmt"
	"net/http"

	"src/common/ctype"
	"src/common/setting"
	"src/util/ssoutil"

	"github.com/labstack/echo/v4"
)

func GetLoginUrl(c echo.Context) error {
	state := ctype.Dict{
		"tenantId": "tenant1",
	}
	realm := setting.KEYCLOAK_DEFAULT_REALM
	clientId := setting.KEYCLOAK_DEFAULT_CLIENT_ID
	redirectUri := setting.KEYCLOAK_DEFAULT_REDIRECT_URI

	authUrl := ssoutil.GetAuthUrl(realm, clientId, redirectUri, state)
	fmt.Println(authUrl)
	return c.Redirect(http.StatusTemporaryRedirect, authUrl)
}

func Callback(c echo.Context) error {
	return c.JSON(http.StatusOK, "Callback")
}

/*
func Callback(c echo.Context) error {
	code := c.QueryParam("code")
	stateStr := c.QueryParam("state")

	// Decode the state
	stateData, err := ssoutil.DecodeState(stateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid state")
	}
	fmt.Println("stateData.............")
	fmt.Println(stateData)

	if code == "" {
		return c.JSON(http.StatusBadRequest, "Authorization code not found")
	}
	fmt.Println("case 1")
	client := gocloak.NewClient(setting.KEYCLOAK_API_URL)
	ctx := c.Request().Context()

	// Exchange the code for tokens
	token, err := client.GetToken(ctx, setting.KEYCLOAK_DEFAULT_REALM, gocloak.TokenOptions{
		ClientID:     gocloak.StringP(setting.KEYCLOAK_DEFAULT_CLIENT_ID),
		ClientSecret: gocloak.StringP(setting.KEYCLOAK_DEFAULT_CLIENT_SECRET),
		Code:         gocloak.StringP(code),
		RedirectURI:  gocloak.StringP(setting.KEYCLOAK_DEFAULT_REDIRECT_URI),
		GrantType:    gocloak.StringP("authorization_code"),
	})
	fmt.Println("case 2")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("case 3")

	// Use token.AccessToken, token.RefreshToken, etc.
	fmt.Println(token.IDToken)
	idToken, _ := jwt.Parse(token.IDToken, nil)
	fmt.Println("case 4")

	// Parse the ID token (without verifying the signature for simplicity)
	idTokenData, _, err := jwt.NewParser().ParseUnverified(idToken.Raw, jwt.MapClaims{})
	if err != nil {
		log.Fatalf("Error parsing ID token: %v", err)
	}
	fmt.Println("case 5")

	// Extract claims from the token
	if claims, ok := idTokenData.Claims.(jwt.MapClaims); ok {
		// Print some standard claims
		fmt.Println("Subject (User ID):", claims["sub"])
		fmt.Println("Name:", claims["name"])
		fmt.Println("Preferred Username:", claims["preferred_username"])
		fmt.Println("Email:", claims["email"])
		fmt.Println("Given Name:", claims["given_name"])
		fmt.Println("Family Name:", claims["family_name"])
	} else {
		log.Fatal("Invalid token claims")
	}

	return c.JSON(http.StatusOK, token)
}
*/
