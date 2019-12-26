package auth

import (
	"github.com/nengkhoiba/go-framework-dda/utils"
	"net/http"
)

func (a *App) somefunction(w http.ResponseWriter, r *http.Request) {

	//some type struct
	type sometype struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	}

	//add some value
	se := sometype{}
	se.ID = 1
	se.Name = "Tom"

	//response as json
	utils.RespondWithJSON(w, http.StatusOK, se)
}
