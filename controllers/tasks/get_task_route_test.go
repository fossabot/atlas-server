package tasks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atlas-io/atlas-server/models"
)

func TestGetTaskRoute(t *testing.T) {
	old := daoGetTask
	defer func() {
		daoGetTask = old
	}()

	task := new(models.Task)
	task.ID = "RDC-123"

	daoGetTask = func(id string, onSuccess func(*models.Task), onError func(error)) {
		onSuccess(task)
	}

	req, err := http.NewRequest("GET", "/tasks/RDC-123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTaskRoute)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	obtainedTask := new(models.Task)
	if err := json.Unmarshal([]byte(rr.Body.String()), obtainedTask); err != nil {
		t.Errorf("Got error %v during unmarshall\n", err.Error())
	}

	if obtainedTask.ID != task.ID {
		t.Error("Mismatched results")
	}
}
