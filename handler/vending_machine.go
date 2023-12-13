package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"vending-machine/controller"
	"vending-machine/db/connection"
	"vending-machine/model"
	"vending-machine/utils"

	"github.com/gorilla/mux"
)

var (
	db = connection.ConnectMysql()
	h  = controller.NewBaseHandler(db)
)

// GetItemList : to get item data in list view
func GetItemList(w http.ResponseWriter, r *http.Request) {
	var (
		retDatas   []*model.Item
		err        error
		statusCode int
	)
	retDatas, statusCode, err = h.GetItemList()
	if err != nil {
		utils.ReturnResponse(w, statusCode, "", retDatas)
	}

	utils.ReturnResponse(w, statusCode, "", retDatas)
}

// BuyItem : to get item that can be bought by nominals that was inputted
func BuyItem(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		reqData    *model.ReqBuyData
		statusCode int
		totalNom   float64
		retData    string
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, utils.ErrorFailedReadRequest(), nil)
		return
	}

	nomAllowed := utils.GetEnvByKey("NOMINAL_ALLOWED")
	for _, v := range reqData.Input {
		if !strings.Contains(nomAllowed, strconv.Itoa(int(v))) {
			utils.ReturnResponse(w, http.StatusBadRequest, "Invalid denomination", nil)
			return
		}
		totalNom += v
	}

	retData, statusCode, err = h.BuyItems(totalNom)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, "Failed to buy items", nil)
		return
	}

	utils.ReturnResponse(w, statusCode, "", retData)
}

// AddItem : to add new item data that can be bought
func AddItem(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		reqData    *model.ReqData
		statusCode int
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, utils.ErrorFailedReadRequest(), nil)
		return
	}

	if reqData.Price == 0 {
		utils.ReturnResponse(w, http.StatusBadRequest, utils.ErrorRequired("price"), nil)
		return
	}

	statusCode, err = h.AddItem(reqData)
	if err != nil {
		utils.ReturnResponse(w, statusCode, utils.ErrorFailedExecData("insert", "item"), nil)
		return
	}

	utils.ReturnResponse(w, statusCode, utils.SuccessExecData("insert", "item"), nil)
}

// ModifyItem : to change existing item data
func ModifyItem(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		reqData        *model.ReqData
		statusCode, id int
		idStr          string
	)

	idStr = mux.Vars(r)["id"]
	id, _ = strconv.Atoi(idStr)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, utils.ErrorFailedReadRequest(), nil)
		return
	}

	statusCode, err = h.ModifyItem(int64(id), reqData)
	if err != nil {
		utils.ReturnResponse(w, statusCode, utils.ErrorFailedExecData("update", "item"), nil)
		return
	}

	utils.ReturnResponse(w, statusCode, utils.SuccessExecData("update", "item"), nil)
}

// RemoveItem : to delete data (change status into 2) so that it won't be shown when call get item list
func RemoveItem(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		statusCode, id int
	)

	idStr := mux.Vars(r)["id"]
	id, _ = strconv.Atoi(idStr)

	statusCode, err = h.RemoveItem(int64(id))
	if err != nil {
		utils.ReturnResponse(w, statusCode, utils.ErrorFailedExecData("delete", "item"), nil)
		return
	}

	utils.ReturnResponse(w, statusCode, utils.SuccessExecData("delete", "item"), nil)
}
