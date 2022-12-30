package ehko

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/alertmanager/template"
)

// Alerts is an endpoint that receive alerts from e.g. Alertmanager (via a POST)
func Alerts(c echo.Context) error {
	var alert template.Data
	// Decode the alert into a slice of *v1.Alerts
	err := json.NewDecoder(c.Request().Body).Decode(&alert)
	if err != nil {
		return err
	}

	// Log the alert
	c.Logger().Infof("%+v\n", alert)

	return nil
}

// Raw logs whatever POSTed to this endpoint
func Raw(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Request().Body.Close()

	c.Logger().Info(string(b))
	return nil
}

// Responder takes the /responder/:code param and returns it as a HTTP response directly.
// E.g. /responder/501 would return a status code of HTTP 501 Not Implemented back to
// the client
func Responder(c echo.Context) error {
	code := c.Param("code")
	statusCode, err := strconv.Atoi(code)
	if err != nil {
		return err
	}
	c.Response().WriteHeader(statusCode)
	c.Response().Write([]byte(code))
	return nil
}
