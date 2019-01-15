package tasks

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atlas-io/atlas-server/dao"

	"github.com/atlas-io/atlas-server/models"
)

// PutTaskRoute submits a new task
func PutTaskRoute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Supplied bad task data"))
		return
	}

	dao.PutTask(&task, func() {
		w.WriteHeader(http.StatusCreated)
	}, func(err error) {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	})
}
