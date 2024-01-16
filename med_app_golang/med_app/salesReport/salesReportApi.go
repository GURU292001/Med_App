package salesreport

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

// salesreport structure to recieve from the bill_master table
type SalesReport_struct struct {
	Bill_No       int     `json:"bill_No"`
	Bill_Date     string  `json:"bill_Date"`
	Medicine_Name string  `json:"medicine_Name"`
	Quantity      int     `json:"quantity"`
	Amount        float64 `json:"amount"`
}

// salesreport request structure
type SalesReport_Req_struct struct {
	From string `json:"from"`
	TO   string `json:"to"`
}

// salesreport response structure
type SalesReport_Resp_struct struct {
	Status            string                 `json:"status"`
	ErrMsg            string                 `json:"errMsg"`
	SalesReport_Array [][]SalesReport_struct `json:"salesReport_Array"`
}

func SalesReportApi(w http.ResponseWriter, r *http.Request) {
	log.Println("SalesReportApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lSalesreportArr [][]SalesReport_struct
		var lSalesreportRec SalesReport_Req_struct
		var lResponseRec SalesReport_Resp_struct
		//read the body and get the form and to date
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "SalesReportApi-01"+err.Error()
		} else {
			//UNmarshall
			err := json.Unmarshal(body, &lSalesreportRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg ="SalesReportApi-02"+err.Error()
			} else {
				lSalesreportArr, err = SalesReport_mtd(lSalesreportRec)
				if err != nil {
					log.Println("AP01 ", err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "SalesReportApi-03"+err.Error()
				} else {
					lResponseRec.SalesReport_Array = lSalesreportArr
					lResponseRec.Status = "S"
					lResponseRec.ErrMsg = "no Error"
					//marshall
					res, err := json.Marshal(lResponseRec)
					if err != nil {
						log.Println(err)
						lResponseRec.Status = "E"
						lResponseRec.ErrMsg = "SalesReportApi-04"+err.Error()
					} else {
						fmt.Fprintln(w, string(res))
						log.Println("SalesReportApi Api (-)")
					}
				}
			}
		}

	}
}

//------------------methods--------------------------------------------------------

func SalesReport_mtd(ldatesRec SalesReport_Req_struct) ([][]SalesReport_struct, error) {
	//local variable to store the result rows
	var lSalesreportArr [][]SalesReport_struct
	var lSalesreportRec SalesReport_struct
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return nil, errors.New("1st query not executed")
	} else {
		defer db.Close()
		var lquery = `select mbd.bill_no ,mbd1.bill_date, (select mmm.medicine_name 
			from medapp_medicine_master mmm 
			where mmm.medicine_master_id=mbd.medicine_master_id) Medicine_Name,
			mbd.quantity ,mbd.amount 
			from medapp_bill_details mbd , medapp_bill_master mbd1
			where mbd.bill_no =mbd1.bill_no and mbd.bill_no in (select mbd1.bill_no
			from medapp_bill_master mbd1
			where mbd1.bill_date>=? and mbd1.bill_date<=?
			);`
		lrows, err := db.Query(lquery, ldatesRec.From, ldatesRec.TO)
		if err != nil {
			fmt.Println("error", err)
			return nil, errors.New("2nd query not executed")
		} else {
			for lrows.Next() {
				err := lrows.Scan(&lSalesreportRec.Bill_No, &lSalesreportRec.Bill_Date, &lSalesreportRec.Medicine_Name, &lSalesreportRec.Quantity, &lSalesreportRec.Amount)
				if err != nil {
					log.Println(err)
				} else {
					var lArr []SalesReport_struct
					lArr = append(lArr, lSalesreportRec)
					fmt.Println("larr:", lArr)
					lSalesreportArr = append(lSalesreportArr, lArr)
				}
			}
			return lSalesreportArr, nil
		}
	}
}
