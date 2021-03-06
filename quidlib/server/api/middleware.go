package api

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo-contrib/session"
	"github.com/synw/quid/quidlib/conf"
	"github.com/synw/quid/quidlib/tokens"

	"github.com/labstack/echo/v4"
)

// AdminMiddleware : check the token claim to see if the user is admin
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check the access token to control that the user is admin
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*tokens.StandardAccessClaims)
		isAdmin := false
		for _, g := range claims.Groups {
			if g == "quid_admin" {
				isAdmin = true
				break
			}
		}
		if !isAdmin {
			emo.ParamError("The user " + claims.UserName + " is not in the quid_admin group")
			return c.NoContent(http.StatusUnauthorized)
		}
		// check session data in prouction
		if conf.IsDevMode {
			return next(c)
		}
		sess, _ := session.Get("session", c)
		if sess.Values["is_admin"] == "true" {
			return next(c)
		}

		emo.Error("Unauthorized session from admin middleware")
		return c.NoContent(http.StatusUnauthorized)
	}
}
