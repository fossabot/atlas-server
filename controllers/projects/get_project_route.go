package projects

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atlas-io/atlas-server/dao"

	"github.com/atlas-io/atlas-server/models"

	"github.com/gorilla/mux"
)

// GetProjectRoute responds with a JSON `project` item
func GetProjectRoute(w http.ResponseWriter, r *http.Request) {
	paramsMap := mux.Vars(r)
	projectID := paramsMap["key"]

	dao.GetProject(projectID, func(project *models.Project) {
		var projectBytes []byte
		var marshalError error
		if projectBytes, marshalError = json.Marshal(&project); marshalError != nil {
			// JSON marshal failed
			log.Println(marshalError.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(marshalError.Error()))
			return
		}
		// JSON marshal successful
		w.Header().Set("Content-Type", "application/json")
		w.Write(projectBytes)
	}, func(err error) {
		// Database fetch failed
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	})
}
