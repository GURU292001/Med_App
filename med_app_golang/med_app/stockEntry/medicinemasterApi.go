package stockEntry

import (
	"encoding/json"
	"fmt"
	"log"
	database "medapp/database"
	"net/http"
)

type Med_mst_tbl_struct struct {
	Medicine_Name string `json:"medicine_Name"`
	Brand         string `json:"brand"`
}

type Med_mst_tbl_Resp_struct struct {
	Status            string               `json:"status"`
	ErrMsg            string               `json:"errMsg"`
	Med_mst_tbl_Array []Med_mst_tbl_struct `json:"med_mst_tbl_Array"`
}

func MedicineMasterApi(w http.ResponseWriter, r *http.Request) {
	log.Println("MedicineMasterApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTION")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var medmstArr []Med_mst_tbl_struct
		var lResponseRec Med_mst_tbl_Resp_struct
		medmstArr, err := Med_mst_mtd()
		fmt.Println("medmstArr",medmstArr)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "MedicineMasterApi-01"+err.Error()
		} else {
			lResponseRec.Med_mst_tbl_Array = medmstArr
			lResponseRec.Status = "S"
			lResponseRec.ErrMsg ="No Error"
		}
		//marshall
		fmt.Println("lResponseRec:",lResponseRec)
		res, err := json.Marshal(lResponseRec)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg ="MedicineMasterApi-02"+err.Error()
		} else {
			fmt.Fprintln(w, string(res))
			log.Println("MedicineMasterApi(-)")
		}
	}
}

//------------------methods--------------------------------------------------------

func Med_mst_mtd() ([]Med_mst_tbl_struct, error) {
	//local variable to store the result
	var lmed_mst Med_mst_tbl_struct
	var lmed_mstRec []Med_mst_tbl_struct
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return nil, err
	} else {
		defer db.Close()
		var lquery = `select medicine_name ,brand 
		from medapp_medicine_master mmm ;`
		lrows, err := db.Query(lquery)
		if err != nil {
			return nil, err
		} else {
			for lrows.Next() {
				err := lrows.Scan(&lmed_mst.Medicine_Name, &lmed_mst.Brand)
				if err != nil {
					log.Println(err)
				} else {
					lmed_mstRec = append(lmed_mstRec, lmed_mst)
				}
			}
			return lmed_mstRec, nil
		}
	}
}
