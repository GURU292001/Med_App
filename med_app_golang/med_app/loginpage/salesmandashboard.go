package loginpage

import (
	"encoding/json"
	"fmt"
	"log"
	database "medapp/database"
	"net/http"
	"strconv"
)


// salseman response structure
type Salesman_Resp_struct struct {
	Status     string  `json:"status"`
	ErrMsg     string  `json:"errMsg"`
	TodaySales float64 `json:"todaySales"`
	Yesterdaysales float64 `json:"yesterdaysales"`
}

func SalesmanDashboardApi(w http.ResponseWriter, r *http.Request) {
	log.Println("SalesmanDashboardApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "HEADER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
	
		var lSalesmanRespRec Salesman_Resp_struct

		lheader := r.Header.Get("HEADER")
		lLogin_id, err := strconv.Atoi(lheader)
		fmt.Println("login_id", lLogin_id)
		if err!=nil{
			fmt.Println(err)
		}
		lTodaysales, lyesterdaysales, err := Salesman_dashboard_mtd(lLogin_id)
		fmt.Println("lTodaysales",lTodaysales)
		fmt.Println("lyesterdaysales",lyesterdaysales)
		if err != nil {
			log.Println(err)
			lSalesmanRespRec.Status = "E"
			lSalesmanRespRec.ErrMsg =  "SalesmanDashboardApi-01"+err.Error()
		} else {
			lSalesmanRespRec.TodaySales = lTodaysales
			lSalesmanRespRec.Yesterdaysales = lyesterdaysales
			lSalesmanRespRec.Status = "S"
			lSalesmanRespRec.ErrMsg = "No Error"
			//marshall
			res, err := json.Marshal(lSalesmanRespRec)
			if err != nil {
				log.Println(err)
				lSalesmanRespRec.Status = "E"
				lSalesmanRespRec.ErrMsg = "SalesmanDashboardApi-02"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("salesmanDashboard Api (-)")
			}
		}

	}
}

//---------------method---------------------------------------------

func Salesman_dashboard_mtd(login_id int) (float64, float64, error) {
	//local variable to store the result rows
	var lTodaysales float64
	var ldifferenc float64
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, 0, err
	} else {
		defer db.Close()
		// today total sales
		value, err := db.Query(`select nvl(sum(bill_amount),0)
		from medapp_bill_master 
		where bill_date =curdate()
		and login_id =?;`, login_id)
		if err != nil {
			return 0, 0, err
		} else {
			value.Next() 
				err := value.Scan(&lTodaysales)
				if err != nil {
					log.Println(err)
				}
			
		}
		// to get sales difference (today sales - yesterday sales)
		lrows, err := db.Query(`select sum(bill_amount) 
		from medapp_bill_master  
		where login_id =? and bill_date =subdate(curdate(),1)`,login_id)
		if err != nil {
			return 0, 0, err
		} else {
			   lrows.Next() 
				err := lrows.Scan(&ldifferenc)
				if err != nil {
					log.Println(err)
				}
			
		}
		return lTodaysales, ldifferenc, nil
	}
}
