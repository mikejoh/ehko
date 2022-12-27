package ehko

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/alertmanager/types"
)

// Alerts is an endpoint that receive alerts from e.g. Alertmanager (via a POST)
func Alerts(c echo.Context) error {
	var alert types.Alert
	// Decode the alert into a slice of *v1.Alerts
	err := json.NewDecoder(c.Request().Body).Decode(&alert)
	if err != nil {
		return err
	}
	// Log the slice of *v1.Alerts
	c.Logger().Infof("%+v\n", alert)

	// Marshal it back to JSON to return it to the client
	b, err := json.Marshal(alert)
	if err != nil {
		return err
	}

	// Return it with pretty printed JSON
	// Convert byte slice into a json.RawMessage
	return c.JSONPretty(http.StatusOK, json.RawMessage(b), "  ")
}

// Responder takes the /responder/:code param and returns it as a HTTP status.
// E.g. /responder/501 would return a status code of HTTP 501 Not Implemented
// back to the client
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
