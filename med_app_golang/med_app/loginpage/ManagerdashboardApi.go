package loginpage

import (
	"encoding/json"
	"fmt"
	"log"
	database "medapp/database"
	"net/http"
)

// Manager response response structure
type Manager_Resp_struct struct {
	Status               string  `json:"status"`
	ErrMsg               string  `json:"errMsg"`
	Todaysales           float64 `json:"todaySales"`
	CurrentInventryValue float64 `json:"currentInventryValue"`
}

func MgrDashboardApi(w http.ResponseWriter, r *http.Request) {
	log.Println("MgrDashboardApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lResponseRec Manager_Resp_struct
		lResponseRec, err := Mgrdashboard_mtd()
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg =  "MgrDashboardApi-01"+err.Error()
		} else {
			lResponseRec.Status = "S"
			lResponseRec.ErrMsg ="No Error"
			//marshall
			res, err := json.Marshal(lResponseRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg =  "MgrDashboardApi-02"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("MgrDashboardApi Api (-)")
			}
		}
	}
}

//------------------methods--------------------------------------------------------

func Mgrdashboard_mtd() (Manager_Resp_struct, error) {
	//local variable to store the result rows
	var lManagerResRec Manager_Resp_struct
	var lTodaysalesquery string
	var lCurrentInventryValue_query string
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return lManagerResRec, err
	} else {
		defer db.Close()
		// today total sales
		lTodaysalesquery = `select nvl(sum(bill_amount),0)
		from medapp_bill_master mbm 
		where bill_date =curdate() ;`
		value, err := db.Query(lTodaysalesquery)
		if err != nil {
			return lManagerResRec, err
		} else {
			for value.Next() {
				err := value.Scan(&lManagerResRec.Todaysales)
				if err != nil {
					log.Println(err)
				}
			}
		}
		lCurrentInventryValue_query = `select nvl(sum(ms.quantity *unit_price) ,0)
		from medapp_stock ms ;`
		lrows, err := db.Query(lCurrentInventryValue_query)
		if err != nil {
			return lManagerResRec, err
		} else {
			for lrows.Next() {
				err := lrows.Scan(&lManagerResRec.CurrentInventryValue)
				if err != nil {
					log.Println(err)
				}
			}
			return lManagerResRec, nil
		}
	}
}
