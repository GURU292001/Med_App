package stockview

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	database "medapp/database"
	"net/http"
)

// stockview structure to recieve from the stock table
type Stockview_struct struct {
	Medicine_Name string  `json:"medicine_Name"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Unit_Price    float64 `json:"unit_Price"`
}

// stock view response structure
type Stockview_Resp_struct struct {
	Status          string               `json:"status"`
	ErrMsg          string               `json:"errMsg"`
	Stockview_Array [][]Stockview_struct `json:"stockview_array"`
}

func StockviewApi(w http.ResponseWriter, r *http.Request) {
	log.Println("stockview Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTION")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lStockviewArr [][]Stockview_struct
		var lResponseRec Stockview_Resp_struct
		lStockviewArr, err := Stockview_mtd()
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "StockviewApi-01"+err.Error()
		} else {
			fmt.Println("lStockviewArr", lStockviewArr)
			lResponseRec.Stockview_Array = lStockviewArr
			lResponseRec.Status = "S"
			lResponseRec.ErrMsg = "no Error"
			//marshall
			res, err := json.Marshal(lResponseRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg =  "StockviewApi-02"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("stockview Api (-)")
			}
		}
	}
}

//------------------methods--------------------------------------------------------

func Stockview_mtd() ([][]Stockview_struct, error) {
	//local variable to store the result rows
	var lStockviewArr [][]Stockview_struct
	var lStockviewRec Stockview_struct
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return nil, errors.New("query not executed")
	} else {
		defer db.Close() //nvl(ms.unit_price ,'-')
		lrows, err := db.Query(`select nvl(mmm.medicine_name,'-') ,nvl(mmm.brand,'-') ,nvl(ms.quantity,'-') ,nvl(ms.unit_price ,'-')
		from medapp_medicine_master mmm ,medapp_stock ms 
		where mmm.medicine_master_id =ms.medicine_master_id; `)
		if err != nil {
			return nil, errors.New("query not executed")
		} else {
			for lrows.Next() {
				err := lrows.Scan(&lStockviewRec.Medicine_Name, &lStockviewRec.Brand, &lStockviewRec.Quantity, &lStockviewRec.Unit_Price)
				if err != nil {
					fmt.Println("3- st err ")
					log.Println(err)
				} else {
					var lArr []Stockview_struct
					lArr = append(lArr, lStockviewRec)
					lStockviewArr = append(lStockviewArr, lArr)
				}
			}
			return lStockviewArr, nil
		}
	}
}
