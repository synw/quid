package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/synw/quid/quidlib/db"
	"github.com/synw/quid/quidlib/models"
	"github.com/synw/quid/quidlib/tokens"
)

// NamespaceInfo : info about a namespace
func NamespaceInfo(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	ID := int64(m["id"].(float64))

	nu, err := db.CountUsersForNamespace(ID)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "error counting users in namespace",
		})
	}

	g, err := db.SelectGroupsForNamespace(ID)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "error counting users in namespace",
		})
	}

	data := models.NamespaceInfo{
		NumUsers: nu,
		Groups:   g,
	}

	return c.JSON(http.StatusOK, &data)
}

// GetNamespaceKey : get the key for a namespace
func GetNamespaceKey(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	ID := int64(m["id"].(float64))

	found, data, err := db.SelectNamespaceKey(ID)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "error finding namespace key",
		})
	}
	if !found {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "namespace not found",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"key": data,
	})

}

// FindNamespace : namespace creation http handler
func FindNamespace(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	name := m["name"].(string)

	data, err := db.SelectNamespaceStartsWith(name)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "error finding namespace",
		})
	}
	return c.JSON(http.StatusOK, &data)

}

// DeleteNamespace : namespace creation http handler
func DeleteNamespace(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	ID := int64(m["id"].(float64))

	err := db.DeleteNamespace(ID)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{
			"error": "error deleting namespace",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "ok",
	})

}

// CreateNamespace : namespace creation http handler
func CreateNamespace(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	name := m["name"].(string)

	key := tokens.GenKey()
	ns, exists, err := createNamespace(name, key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "error creating namespace",
		})
	}
	if exists {
		return c.JSON(http.StatusConflict, echo.Map{
			"error": "namespace already exists",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"namespace_id": ns.ID,
	})
}

// createNamespace : create a namespace
func createNamespace(name string, key string) (models.Namespace, bool, error) {
	ns := models.Namespace{}

	exists, err := db.NamespaceExists(name)
	if err != nil {
		return ns, false, err
	}
	if exists {
		return ns, true, nil
	}

	uid, err := db.CreateNamespace(name, key)
	if err != nil {
		return ns, false, err
	}
	ns.ID = uid
	ns.Name = name
	return ns, false, nil
}