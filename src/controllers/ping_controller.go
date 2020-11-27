package controllers

import (
	"net/http"

	"github.com/gervi/bookstore_items-api/src/utils/http_utils"
)

var (
	PingController pingControllerInterface = &pingController{}
)

const (
	pong = "pong"
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	http_utils.RespondJson(w, http.StatusOK, "pong")
}
