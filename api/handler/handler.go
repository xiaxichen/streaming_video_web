package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"streaming_video_web/api/dbops"
	"streaming_video_web/api/def"
	"streaming_video_web/api/middleware"
	"streaming_video_web/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &def.UserCreateStruct{}

	if err := json.Unmarshal(res, ubody); err != nil {
		middleware.SendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil {
		middleware.SendErrorResponse(w, def.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.UserName)
	su := &def.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		middleware.SendErrorResponse(w, def.ErrorInternalFaults)
		return
	} else {
		middleware.SendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
