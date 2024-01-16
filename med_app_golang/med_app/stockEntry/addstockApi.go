package stockEntry

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

// Add stock  structure to recieve the data from the front
type Addstock_Req_struct struct {
	User_id       string `json:"user_id"`
	Medicine_Name string `json:"medicine_Name"`
	Brand         string `json:"brand"`
}

// Adduser response structure
type Addstock_Resp_struct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func AddStockIApi(w http.ResponseWriter, r *http.Request) {
	log.Println("AddStockApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var lAddStockReqRec Addstock_Req_struct
		var lResponseRec Addstock_Resp_struct
		//read the body and get the form and to date
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "AddStockIApi-01"+err.Error()
		} else {
			//UNmarshall
			err := json.Unmarshal(body, &lAddStockReqRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg ="AddStockIApi-02"+err.Error()
			} else {
				fmt.Println("lAddStockReqRec: ", lAddStockReqRec)
				result, err := Addstock_mtd(lAddStockReqRec)
				if err != nil {
					log.Println(err)
					lResponseRec.Status = "Error"
					lResponseRec.ErrMsg ="AddStockIApi-03"+err.Error()
				} else {
					//checking whether the insert operation is done or not
					if result == 1 {
						lResponseRec.Status = "S"
						lResponseRec.ErrMsg = "no Error"
					} else {
						lResponseRec.Status = "E"
						lResponseRec.ErrMsg = "AddStockIApi-04"+err.Error()
					}
					//marshall
					res, err := json.Marshal(lResponseRec)
					if err != nil {
						log.Println(err)
						lResponseRec.Status = "E"
						lResponseRec.ErrMsg = "AddStockIApi-05"+err.Error()
					} else {
						fmt.Fprintln(w, string(res))

					}
				}
			}
		}

	}
	log.Println("AddStockApi(-)")
}

//------------------methods--------------------------------------------------------

func Addstock_mtd(laddstockRec Addstock_Req_struct) (int64, error) {

	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, errors.New("query not executed")
	} else {
		defer db.Close()
		var lquery = `insert into medapp_medicine_master 
		(medicine_name,brand, created_by, created_date, updated_by, updated_date)
		values(?,?,?,curdate(),?,curdate());`
		lrows, err := db.Exec(lquery, laddstockRec.Medicine_Name, laddstockRec.Brand, laddstockRec.User_id, laddstockRec.User_id)
		fmt.Println()
		if err != nil {
			return 0, errors.New("query not executed")
		} else {
			insert, err := lrows.RowsAffected()
			fmt.Println("insert", insert)
			if err != nil {
				return 0, errors.New("query not executed")
			}
			return insert, nil
		}
	}
}
