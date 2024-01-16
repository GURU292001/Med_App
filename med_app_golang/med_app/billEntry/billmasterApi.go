package billEntry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

//TO recieve the request from thr front
type Bill_mst_Req_struct struct {
	User_id     string  `json:"User_id"`
	Login_id    int     `json:"login_id"`
	Bill_No     int     `json:"bill_No"`
	Bill_Amount float64 `json:"bill_Amount"`
	Bill_Gst    float64 `json:"bill_Gst"`
	Net_Price   float64 `json:"net_Price"`
}

func BillMasterApi(w http.ResponseWriter, r *http.Request) {
	log.Println("BillMasterApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var lBillmasterRec Bill_mst_Req_struct
		var lResponseRec Bill_Details_Resp_struct
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "BillMasterApi-01"+err.Error()
		} else {
			//unmarshal
			err := json.Unmarshal(body, &lBillmasterRec)
			//method calling
			lInsert, err := Bill_mst_mtd(lBillmasterRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "BillMasterApi-02"+err.Error()
			} else {
				if lInsert == 1 {
					lResponseRec.Status = "S"
					lResponseRec.ErrMsg = "No error"
				} else {
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "BillMasterApi-03"+err.Error()
				}
			}
			//marshall
			res, err := json.Marshal(lResponseRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "BillMasterApi-04"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("QuantityCheckApi (-)")
			}
		}
	}
}

//------------------methods--------------------------------------------------------

func Bill_mst_mtd(lbill_mstRec Bill_mst_Req_struct) (int64, error) {

	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		insert_query := `insert into medapp_bill_master 
		(bill_no,bill_date,bill_amount,bill_gst, net_price, login_id,
		created_by,created_date,updated_by,updated_date)
		values(?,curdate(),?,?,?,?,?,curdate(),?,curdate());`
		lrows, err := db.Exec(insert_query, lbill_mstRec.Bill_No, lbill_mstRec.Bill_Amount, lbill_mstRec.Bill_Gst, lbill_mstRec.Net_Price, lbill_mstRec.Login_id, lbill_mstRec.User_id, lbill_mstRec.User_id)
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
