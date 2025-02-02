package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/merico-dev/lake/api"
	"github.com/merico-dev/lake/utils"
	"github.com/stretchr/testify/mock"
)

func TestNewTask(t *testing.T) {
	r := gin.Default()
	api.RegisterRouter(r)

	type services struct {
		mock.Mock
	}

	// fakeTask := models.Task{}
	testObj := new(services)
	testObj.On("CreateTask").Return(true, nil)

	w := httptest.NewRecorder()
	params := strings.NewReader(`[[{ "plugin": "jira", "options": { "host": "www.jira.com" } }]]`)
	req, _ := http.NewRequest("POST", "/task", params)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)
	resp := w.Body.String()
	tasks, err := utils.JsonToMap(resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tasks[0][0]["plugin"], "jira")
}
