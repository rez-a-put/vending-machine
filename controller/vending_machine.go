package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"vending-machine/model"
	"vending-machine/repository"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

// GetItemList : function to get list of items
func (h *BaseHandler) GetItemList() (retDatas []*model.Item, statusCode int, err error) {
	retDatas, err = repository.GetItemList(h.db, nil, "")
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retDatas, http.StatusOK, nil
}

// BuyItems : function to buy items and to get returned item that bought based on nominal inputted
func (h *BaseHandler) BuyItems(totalNom float64) (retString string, statusCode int, err error) {
	var (
		itemDatas []*model.Item
		mapBought = make(map[int64]int64)
		mapItems  = make(map[int64]*model.Item)
	)

	itemDatas, err = repository.GetItemList(h.db, map[string]interface{}{"status": 1, "total_payment": totalNom}, "-price")
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	for i := 0; i < len(itemDatas); {
		if totalNom >= itemDatas[i].Price {
			_, isExist := mapBought[itemDatas[i].ID]
			if isExist {
				mapBought[itemDatas[i].ID]++
			} else {
				mapBought[itemDatas[i].ID] = 1
			}

			mapItems[itemDatas[i].ID] = itemDatas[i]
			totalNom -= itemDatas[i].Price

			if totalNom <= 0 {
				break
			}
		} else {
			i++
		}
	}

	for i, v := range mapBought {
		retString += strconv.Itoa(int(v)) + " " + mapItems[i].Name + ", "
	}
	retString = strings.TrimSuffix(retString, ", ")

	return retString, http.StatusOK, nil
}

// AddItem : function to add new items
func (h *BaseHandler) AddItem(reqData *model.ReqData) (statusCode int, err error) {
	var values []interface{}

	values = append(values, reqData.Name)
	values = append(values, reqData.Price)
	values = append(values, 1)

	err = repository.AddItem(h.db, values)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, nil
}

// ModifyItem : function to modify existing item
func (h *BaseHandler) ModifyItem(id int64, reqData *model.ReqData) (statusCode int, err error) {
	var (
		mapValues    = make(map[string]interface{})
		rowsAffected int64
	)

	if reqData.Name != "" {
		mapValues["name"] = reqData.Name
	}

	if reqData.Price > 0 {
		mapValues["price"] = reqData.Price
	}

	rowsAffected, err = repository.ModifyItem(h.db, id, mapValues)
	if err != nil || rowsAffected == 0 {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

// RemoveItem : function to delete existing item and won't be shown when get item list
func (h *BaseHandler) RemoveItem(id int64) (statusCode int, err error) {
	var rowsAffected int64

	rowsAffected, err = repository.RemoveItem(h.db, id)
	if err != nil || rowsAffected == 0 {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
