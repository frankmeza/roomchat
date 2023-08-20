package users

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userMockData = UserProps{
		Bio:      "bio",
		Email:    "frank@frank.com",
		Name:     "frank",
		Username: "fraaank",
	}

	user = User{
		UserProps: userMockData,
	}

	usernameUrl = func(username string) string {
		return "/users/username/" + username
	}

	userPayload = `{"username":"fraaank"}`
)

func TestGetUser(t *testing.T) {
	dbConn, _ := db.GetDbConnection()
	dbConn.Debug().Create(&user)

	defer dbConn.Debug().Delete(&user)

	request := httptest.NewRequest(
		http.MethodGet,
		usernameUrl("fraaank"),
		strings.NewReader(userPayload),
	)

	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	echoServer := echo.New()
	recorder := httptest.NewRecorder()

	echoContext := echoServer.NewContext(request, recorder)
	echoContext.SetParamNames(constants.USERNAME)
	echoContext.SetParamValues("fraaank")

	handlerError := handleGetUser(echoContext)

	if assert.NoError(t, handlerError) {
		responseBody, err := io.ReadAll(recorder.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var result = make(map[string]interface{})
		if err := json.Unmarshal(responseBody, &result); err != nil {
			log.Fatalln(err)
		}

		username, isExists := result["username"]
		if isExists {
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, "fraaank", username)
		}
	}
}
