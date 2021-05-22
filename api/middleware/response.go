package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"streaming_video_web/api/def"
)

func SendErrorResponse(w http.ResponseWriter, errResp def.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
