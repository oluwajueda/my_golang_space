package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	//get cookie

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	//if the user exists already, get user

	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	return u
}
