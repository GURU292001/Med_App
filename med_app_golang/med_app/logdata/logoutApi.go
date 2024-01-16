package logdata

import (
	"encoding/json"
	"fmt"
	"log"
	database "medapp/database"
	struc "medapp/loginhistory"
	"net/http"
	"strconv"
)

func LogoutDataApi(w http.ResponseWriter, r *http.Request) {
	log.Println("LogoutDataApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "HEADER,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "POST" {
		//local variable
		var lLoginid int
		var lResponseRec struc.Login_Hsty_Resp_struct
		//read the head
		header := r.Header.Get("HEADER")
		lLoginid, err := strconv.Atoi(header)
		fmt.Println("logout header", lLoginid)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "LogoutDataApi-01"+err.Error()
		}
		lresult, err := Logoutdata(lLoginid)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg ="LogoutDataApi-02"+err.Error()
		} else {
			//checking whether the insert operation is done or not
			if lresult == 1 {
				lResponseRec.Status = "S"
				lResponseRec.ErrMsg = "no Error"
			} else {
				lResponseRec.Status = "Error"
				lResponseRec.ErrMsg = "LogoutDataApi-03"+err.Error()
			}
			//marshall
			res, err := json.Marshal(lResponseRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "LogoutDataApi-04"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("LogoutDataApi(-)")
			}
		}
	}
}

//------------------methods--------------------------------------------------------

func Logoutdata(login_id int) (int64, error) {
	//local variable to store the result
	var linsert int64
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		//Query for update
		var lquery = `update medapp_login_history 
		set logout_date =curdate(),
		logout_time =curtime()
		where login_id =? and login_history_id =(select max(login_history_id)
		from medapp_login_history
		where login_id =?);`
		lrows, err := db.Exec(lquery, login_id, login_id)
		if err != nil {
			return 0, err
		} else {
			linsert, err = lrows.RowsAffected()
			if err != nil {
				return 0, err
			}
			return linsert, nil
		}
	}
}
