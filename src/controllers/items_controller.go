package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gervi/bookstore_items-api/src/domain/items"
	"github.com/gervi/bookstore_items-api/src/domain/queries"
	"github.com/gervi/bookstore_items-api/src/services"
	"github.com/gervi/bookstore_items-api/src/utils/http_utils"
	oauth "github.com/gervi/bookstore_oauth-go"
	"github.com/gervi/bookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		respErr := rest_errors.NewNotAuthorizedError("not authorized")
		http_utils.RespondError(w, respErr)
		return
	}
	sellerID := oauth.GetCallerId(r)
	if sellerID == 0 {
		respErr := rest_errors.NewNotAuthorizedError("not authorized")
		http_utils.RespondError(w, respErr)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewDbError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()
	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerID
	result, respErr := services.ItemsService.Create(itemRequest)
	if respErr != nil {
		http_utils.RespondError(w, respErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)

		return
	}
	result, respErr := services.ItemsService.Get(vars["id"])
	if respErr != nil {
		http_utils.RespondError(w, respErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, result)

}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rest_errors.NewBadRequestError("invalid request body")
		return
	}
	defer r.Body.Close()
	var esQuery queries.EsQuery
	if err := json.Unmarshal(bytes, &esQuery); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}
	result, respErr := services.ItemsService.Search(esQuery)
	if respErr != nil {
		respErr := rest_errors.NewBadRequestError("invalid query parameters")
		http_utils.RespondError(w, respErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, result)
}

func (c *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)

		return
	}
	respErr := services.ItemsService.Delete(vars["id"])
	if respErr != nil {
		http_utils.RespondError(w, respErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, "Ok")
}

func (c *itemsController) Put(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		respErr := rest_errors.NewNotAuthorizedError("not authorized")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewDbError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()
	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}
	fmt.Println(&itemRequest)
	result, _ := services.ItemsService.Get(itemRequest.Id)
	fmt.Println(result)
	itemRequest.SoldQuantity = result.SoldQuantity
	itemRequest.Seller = result.Seller
	result, respErr := services.ItemsService.Update(itemRequest)
	if respErr != nil {
		http_utils.RespondError(w, respErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, result)
}
