package tasks

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atlas-io/atlas-server/dao"
	"github.com/atlas-io/atlas-server/models"

	"github.com/gorilla/mux"
)

var daoGetTask = dao.GetTask

// GetTaskRoute gets a task by its `taskid`
func GetTaskRoute(w http.ResponseWriter, r *http.Request) {
	paramsMap := mux.Vars(r)
	taskID := paramsMap["id"]

	daoGetTask(taskID, func(task *models.Task) {
		// Database fetch successful
		var taskBytes []byte
		var marshalError error
		if taskBytes, marshalError = json.Marshal(task); marshalError != nil {
			// JSON marshal failed
			log.Println(marshalError.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(marshalError.Error()))
			return
		}

		// JSON marshal successful
		w.Header().Set("Content-Type", "application/json")
		w.Write(taskBytes)
	}, func(err error) {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	})
}
