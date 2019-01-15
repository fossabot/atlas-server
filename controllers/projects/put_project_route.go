package projects

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atlas-io/atlas-server/dao"
	"github.com/atlas-io/atlas-server/models"
)

// PutProjectRoute submits a new project
func PutProjectRoute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Supplied bad project data"))
		return
	}

	dao.PutProject(&project, func() {
		w.WriteHeader(http.StatusCreated)
	}, func(err error) {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	})
}
