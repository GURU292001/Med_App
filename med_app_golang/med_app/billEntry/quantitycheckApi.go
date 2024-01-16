package billEntry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

type AddItem_Req_struct struct {
	User_id       string `json:"user_id"`
	Medicine_Name string `json:"medicine_Name"`
	Quantity      int    `json:"quantity"`
}
type AddItem_Resp_struct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
	Availqty int `json:"availqty"`
}

type Bill_Resp_struct struct{
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
	Bill_no int `json:"bill_no"`
}

func QuantityCheckApi(w http.ResponseWriter, r *http.Request) {
	log.Println("QuantityCheckApi  (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var lAvail_quantity int 
		var lAddItemRec AddItem_Req_struct
		var lResponseRec AddItem_Resp_struct
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "Error"
		} else {
			//unmarshal
			err := json.Unmarshal(body, &lAddItemRec)
			//method calling
			lAvail_quantity, err = Quantitycheck_mtd(lAddItemRec.Medicine_Name)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "Error"
			} else {
				if lAvail_quantity<lAddItemRec.Quantity{
					lResponseRec.Status = "Notexist"
					lResponseRec.ErrMsg = "Quantity is higher than the available"
					lResponseRec.Availqty=lAvail_quantity
					}else{
						lResponseRec.Status = "exist"
						lResponseRec.ErrMsg = "No error"
						lResponseRec.Availqty=lAvail_quantity
				}
				}
				//marshall
				res, err := json.Marshal(lResponseRec)
				if err != nil {
					log.Println(err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "Error"
				} else {
					fmt.Fprintln(w, string(res))
					log.Println("QuantityCheckApi (-)")
				}
			}
		}
	}


func BillnoApi(w http.ResponseWriter, r *http.Request){
	log.Println("bill_no  (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method  == "GET"{
		//LOCAL variable
		var lbill_no int
		var lResponseRec Bill_Resp_struct
		lbill_no,err:=Billno_mtd()
		if err!=nil{
			log.Println(err)
			lResponseRec.Status="E"
			lResponseRec.ErrMsg="BillnoApi-01"+err.Error()
		}else{
			lResponseRec.Bill_no=lbill_no
			lResponseRec.Status="S"
			lResponseRec.ErrMsg="No Error"
		}
		//marshall
		res,err:=json.Marshal(lResponseRec)
		if err!=nil{
			log.Println(err)
			lResponseRec.Status="E"
			lResponseRec.ErrMsg="BillnoApi-02"+err.Error()
		}else{
			fmt.Fprintf(w,string(res))
			log.Println("bill_no  (-)")
		}
	}

}
func Billno_mtd()(int,error){
	//local variable
	var lbillno int
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		lrows,err:=db.Query(`select max(bill_no)
		from medapp_bill_details mbd ;`)
		if err!=nil{
			fmt.Println(err)
		}else{
			lrows.Next()
			err=lrows.Scan(&lbillno)
			if err!=nil{
				fmt.Println(err)
			}else{
				return lbillno,nil
			}
		}
	}
	return 0,nil
}
//------------------methods--------------------------------------------------------

func Quantitycheck_mtd(lmedicine_name string )(int , error) {
	//local variable to store the result rows
	var lquantity int 
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		lrows, err := db.Query(`select quantity
		from medapp_stock 
		where medicine_master_id =(select medicine_master_id 
		from medapp_medicine_master 
		where medicine_name =?);`,lmedicine_name)
		if err != nil {
			return 0, err
		} else {
			lrows.Next()
			err:=lrows.Scan(&lquantity)
			if err != nil {
				return 0, err
			}else{
				return lquantity, nil
			}
		}
	}
}
