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

type Med_mst_Req_struct struct {
	User_id       string  `json:"user_id"`
	Medicine_Name string  `json:"medicine_Name"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Unit_Price    float64 `json:"unit_Price"`
}
type Med_mst_Resp_struct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func UpdateStockApi(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateStockApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "POST" {
		//local variable
		var lresult int64
		var lMed_mstRec Med_mst_Req_struct
		var lResponseRec Med_mst_Resp_struct
		//read the body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "UpdateStockApi-01"+err.Error()
		} else {
			err = json.Unmarshal(body, &lMed_mstRec)
			fmt.Println("lMed_mstRec", lMed_mstRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "Error"
				lResponseRec.ErrMsg = "UpdateStockApi-02"+err.Error()
			} else {
				//method calling (Updatestock_mtd)
				lresult, err = Updatestock_mtd(lMed_mstRec)
				fmt.Println("lresult:", lresult)
				if err != nil {
					fmt.Println("error mtd:", err)
				}
				//checking whether the lInsert operation is done or not
				if lresult == 1 {
					lResponseRec.Status = "S"
					lResponseRec.ErrMsg = "no Error"
				} else {
					lResponseRec.Status = "Error"
					lResponseRec.ErrMsg ="UpdateStockApi-03"+err.Error()
				}
				//marshall
				res, err := json.Marshal(lResponseRec)
				if err != nil {
					log.Println(err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg ="UpdateStockApi-04"+err.Error()
				} else {
					fmt.Fprintln(w, string(res))
					log.Println("UpdateStockApi(-)")
				}
			}
		}
	}
}

//------------------methods--------------------------------------------------------

func Updatestock_mtd(med_mst_Rec Med_mst_Req_struct) (int64, error) {
	fmt.Println("med_mst_Rec:", med_mst_Rec.Medicine_Name)
	fmt.Println("med_mst_Rec:", med_mst_Rec.Brand)
	//local variable to store the result
	var lmaster_id int
	var lmaster_id_Arr []int
	var lflag = false
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		var lquery = `select nvl(medicine_master_id,0)
		from medapp_medicine_master
		where medicine_name =? and brand =?;`
		lrow, err := db.Query(lquery, med_mst_Rec.Medicine_Name, med_mst_Rec.Brand)
		if err != nil {
			return 0, err
		} else {
			lrow.Next()
			err=lrow.Scan(&lmaster_id)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println("lmaster_id", lmaster_id)
			lrow, err = db.Query(`select nvl(medicine_master_id,0)
			from medapp_stock ;`)
			if err != nil {
				return 0, err
			} else {
				for lrow.Next(){
					var lmst_lst int
				err:=lrow.Scan(&lmst_lst)
				if err!=nil{
					fmt.Println(err)
				}else{
					lmaster_id_Arr=append(lmaster_id_Arr,lmst_lst )
				}
			}
				fmt.Println("lmaster_id_Arr:", lmaster_id_Arr)
				for _, value := range lmaster_id_Arr {
					if value == lmaster_id {
						lflag = true
					}
				}
				if lflag {
					lupdate_query := `update medapp_stock 
					set quantity =?,
					unit_price =?,
					updated_by =?,
					updated_date =curdate() 
					where medicine_master_id =?;`
					lrows, err := db.Exec(lupdate_query, med_mst_Rec.Quantity, med_mst_Rec.Unit_Price, med_mst_Rec.User_id, lmaster_id)
					if err != nil {
						return 0, errors.New("query not executed")
					} else {
						lUpdate, err := lrows.RowsAffected()
						if err != nil {
							log.Println(err)
						} else {
							return lUpdate, nil
						}

					}
				} else {
					linsert_query := `insert into medapp_stock 
					(medicine_master_id,quantity,unit_price,created_by,created_date,updated_by,updated_date)
					values(?,?,?,?,curdate(),?,curdate());`
					lrows, err := db.Exec(linsert_query, lmaster_id, med_mst_Rec.Quantity, med_mst_Rec.Unit_Price, med_mst_Rec.User_id, med_mst_Rec.User_id)
					if err != nil {
						return 0, errors.New("query not executed")
					} else {
						lInsert, err := lrows.RowsAffected()
						if err != nil {
							log.Println(err)
						} else {
							return lInsert, nil
						}
					}
				}
			}
			return 0, nil
		}
	}
}
