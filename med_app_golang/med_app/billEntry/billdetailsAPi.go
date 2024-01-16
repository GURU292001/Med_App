package billEntry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

type Bill_Details_Req_struct struct {
	Bill_No         int                   `json:"bill_No"`
	User_id  string `json:"user_id"`
	Bill_detailsArr []Bill_Details_struct `json:"bill_detailsArr"`
}

type Bill_Details_struct struct {
	Medicine_Name string  `json:"medicine_Name"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Unit_Price    float64 `json:"unit_Price"`
	Amount        float64 `json:"amount"`
}

type Bill_Details_Resp_struct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func BillDetailsApi(w http.ResponseWriter, r *http.Request) {
	log.Println("BillDetailsApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var lflag = true
		var lBilldetailsRec Bill_Details_Req_struct
		var lResponseRec Bill_Details_Resp_struct
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "BillDetailsApi-01"+err.Error()
		} else {
			//unmarshal
			err := json.Unmarshal(body, &lBilldetailsRec)
			if err!=nil{
				log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "BillDetailsApi-02"+err.Error()
			}
			fmt.Println("lBilldetailsRec",lBilldetailsRec)
			//method calling
			for _, value := range lBilldetailsRec.Bill_detailsArr {
				lInsert, err := Bill_details_mtd(value,lBilldetailsRec.Bill_No,lBilldetailsRec.User_id)
				lUpdate,err:=QuantityUpdate_mtd(value,lBilldetailsRec.User_id)
				if err != nil {
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "BillDetailsApi-03"+err.Error()
				}
				if lInsert != 1 && lUpdate!=1 {
					lflag = false
				}
			}
			if lflag {
				lResponseRec.Status = "S"
				lResponseRec.ErrMsg = "No error"
			} else {
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "BillDetailsApi-04"+err.Error()
			}
			//marshall
			res, err := json.Marshal(lResponseRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "BillDetailsApi-05"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("BillDetailsApi (-)")
			}
		}
	}
}

//------------------methods--------------------------------------------------------
//billdetails update
func Bill_details_mtd(lbill_details_Rec Bill_Details_struct,lbill_no int, luser_id string ) (int64, error) {

	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		//query to insert the bill details in medapp_bill_details
		insert_query := `insert into medapp_bill_details 
		(bill_no, medicine_master_id, quantity, unit_price, amount
		,created_by,created_date,updated_by,updated_date )
		values(?,(select medicine_master_id 
			from medapp_medicine_master
			 where medicine_name=? and brand=?),?,?,?,?,curdate(),?,curdate())`
		lrows, err := db.Exec(insert_query,lbill_no,lbill_details_Rec.Medicine_Name,lbill_details_Rec.Brand, lbill_details_Rec.Quantity, lbill_details_Rec.Unit_Price, lbill_details_Rec.Amount,luser_id,luser_id)
		if err != nil {
			return 0, err
		} else {
			linsert, err := lrows.RowsAffected()
			if err != nil {
				fmt.Println(err)
			} else {
				return linsert, nil
			}

		}
	}
	return 0, nil
}
//stock quantity up
func QuantityUpdate_mtd(lqtyUpdateRec Bill_Details_struct,luser_id string ) (int64, error) {
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
				lupdate_query := `update medapp_stock 
						set quantity=quantity-?,
						updated_by=?,
						updated_date=curdate()
						where medicine_master_id =(select medicine_master_id 
						from medapp_medicine_master 
						where medicine_name=? and brand=?);`
				lrows, err := db.Exec(lupdate_query, lqtyUpdateRec.Quantity, luser_id, lqtyUpdateRec.Medicine_Name, lqtyUpdateRec.Brand)
				if err != nil {
					return 0, err
				} else {
					lUpdate ,err:= lrows.RowsAffected()
					if err!=nil{
						fmt.Println(err)
					}else{
						return lUpdate, nil
					}
				}
			// }
		}

		return 0,nil
	}