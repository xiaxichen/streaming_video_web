package main

import (
	"net/http"
	"streaming_video_web/api/def"
	"streaming_video_web/api/middleware"
	"streaming_video_web/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_userName = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	userName, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_userName, userName)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	userName := r.Header.Get(HEADER_FIELD_userName)
	if len(userName) == 0 {
		middleware.SendErrorResponse(w, def.ErrorNotAuthUser)
		return false
	}

	return true
}
