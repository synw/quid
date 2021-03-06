package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	db "github.com/synw/quid/quidlib/server/db"
	"github.com/synw/quid/quidlib/tokens"
)

// RequestAdminAccessToken : request an access token from a refresh token
// for the quid namespace
func RequestAdminAccessToken(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	refreshToken, ok := m["refresh_token"].(string)
	if !ok {
		emo.ParamError("provide a refresh_token parameter")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "provide a refresh_token parameter",
		})
	}

	// get the namespace
	_, ns, err := db.SelectNamespaceFromName("quid")
	if err != nil {
		emo.Error(err)
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": true,
		})
	}

	// verify the refresh token
	var username string
	token, err := jwt.ParseWithClaims(refreshToken, &tokens.StandardRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ns.RefreshKey), nil
	})
	if claims, ok := token.Claims.(*tokens.StandardRefreshClaims); ok && token.Valid {
		username = claims.UserName
		fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
	} else {
		emo.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// get the user
	found, u, err := db.SelectNonDisabledUser(username, ns.ID)
	if !found {
		emo.Error("User not found: " + username)
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": true,
		})
	}

	// get the user groups names
	groupNames, err := db.SelectGroupsNamesForUser(u.ID)
	if err != nil {
		emo.Error("Groups error")
		log.Fatal(err)
		return err
	}
	// get the user orgs names
	orgsNames, err := db.SelectOrgsNamesForUser(u.ID)
	if err != nil {
		emo.Error("Orgs error")
		log.Fatal(err)
		return err
	}
	// check if the user is in the admin group
	isAdmin := false
	for _, gn := range groupNames {
		if gn == "quid_admin" {
			isAdmin = true
			break
		}
	}
	if !isAdmin {
		emo.Warning("Admin access token request from user", u.UserName, "that is not in the quid_admin group")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// generate the access token
	_, t, err := tokens.GenAccessToken(ns.Key, ns.MaxTokenTTL, u.UserName, groupNames, orgsNames, "5m")
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": true,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// RequestAccessToken : request an access token from a refresh token
func RequestAccessToken(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	refreshToken, ok := m["refresh_token"].(string)
	if !ok {
		emo.ParamError("provide a refresh_token parameter")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "provide a refresh_token parameter",
		})
	}
	namespace, ok := m["namespace"].(string)
	if !ok {
		emo.ParamError("provide a namespace parameter")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "provide a namespace parameter",
		})
	}
	timeout := c.Param("timeout")

	// get the namespace
	exists, ns, err := db.SelectNamespaceFromName(namespace)
	if !exists {
		emo.Error("The namepsace does not exist")
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": true,
		})
	}
	if err != nil {
		emo.Error(err)
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": true,
		})
	}

	// check if the endpoint is available
	if !ns.PublicEndpointEnabled {
		emo.Error("Public endpoint unanuthorized")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// verify the refresh token
	var username string
	token, err := jwt.ParseWithClaims(refreshToken, &tokens.StandardRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ns.RefreshKey), nil
	})
	if claims, ok := token.Claims.(*tokens.StandardRefreshClaims); ok && token.Valid {
		username = claims.UserName
		fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
	} else {
		emo.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// get the user
	found, u, err := db.SelectNonDisabledUser(username, ns.ID)
	if !found {
		emo.Error("User not found: " + username)
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": true,
		})
	}

	// get the user groups names
	groupNames, err := db.SelectGroupsNamesForUser(u.ID)
	if err != nil {
		emo.Error("Groups error")
		log.Fatal(err)
		return err
	}

	// get the user orgs names
	orgsNames, err := db.SelectOrgsNamesForUser(u.ID)
	if err != nil {
		emo.Error("Groups error")
		log.Fatal(err)
		return err
	}

	// generate the access token
	isAuth, t, err := tokens.GenAccessToken(ns.Key, ns.MaxTokenTTL, u.UserName, groupNames, orgsNames, timeout)
	if !isAuth {
		emo.Error("Timeout unauthorized")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": true,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// RequestRefreshToken : http login handler
func RequestRefreshToken(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	// username
	usernameParam, ok := m["username"]
	var username string
	if ok {
		username = usernameParam.(string)
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "provide a username",
		})
	}
	// password
	passwordParam, ok := m["password"]
	var password string
	if ok {
		password = passwordParam.(string)
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "provide a password",
		})
	}
	// namespace
	nsParam, ok := m["namespace"]
	var namespace string
	if ok {
		namespace = nsParam.(string)
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "provide a namespace",
		})
	}
	// timeout
	timeout := c.Param("timeout")

	// get the namespace
	exists, ns, err := db.SelectNamespaceFromName(namespace)
	if err != nil {
		return err
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "namespace does not exist",
		})
	}
	if !ns.PublicEndpointEnabled {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// check if the endpoint is available
	if !ns.PublicEndpointEnabled {
		emo.Error("Public endpoint unanuthorized")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// check if the user password matches
	isAuthorized, u, err := checkUserPassword(username, password, ns.ID)
	if err != nil {
		return err
	}
	// Respond with unauthorized status
	if isAuthorized == false {
		fmt.Println(username, "unauthorized")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	// generate the token
	isAuth, t, err := tokens.GenRefreshToken(ns.Name, ns.RefreshKey, ns.MaxRefreshTokenTTL, u.UserName, timeout)
	if err != nil {
		log.Fatal(err)
	}
	if !isAuth {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "max timeout exceeded",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
