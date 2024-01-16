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



func LoginDataApi(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginDataApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "HEADER ,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var lresult int64 
		var lLoginid int 
		var lResponseRec struc.Login_Hsty_Resp_struct
		//read the head
				header:=r.Header.Get("HEADER")
				log.Println("head:",header)
				lLoginid,err:=strconv.Atoi(header)
				if err != nil{
					log.Println( err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg =  "LoginDataApi-01"+err.Error()
				}
				lresult, err = Logindata(lLoginid)
				if err != nil {
					log.Println( err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg =  "LoginDataApi-02"+err.Error()
				} else {
					//checking whether the insert operation is done or not
					if lresult==1 {
						lResponseRec.Status = "S"
					lResponseRec.ErrMsg = "no Error"
					}else{
						log.Println( err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg =  "LoginDataApi-03"+err.Error()
					}
					//marshall
					res, err := json.Marshal(lResponseRec)
					if err != nil {
						log.Println(err)
						lResponseRec.Status = "E"
						lResponseRec.ErrMsg = "LoginDataApi-04"+err.Error()
					} else {
						fmt.Fprintln(w, string(res))
						log.Println("LoginDataApi(-)")
					}
				}
			}
		}
//------------------methods--------------------------------------------------------

func Logindata(login_id int) (int64,error) {
	//local variable to store the result
	var linsert int64
	var luser_id string
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()
		var user_id_query=`select user_id
		from medapp_login 
		where login_id =?;`
		lrows,err:=db.Query(user_id_query,login_id)
		if err != nil {
			return 0, err
		} else {
			lrows.Next()
			err=lrows.Scan(&luser_id)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println("user_id:",luser_id)
		var lquery = `insert into
		medapp_login_history (login_id,login_date,login_time, created_by,created_date,updated_by,updated_date)
	   values(?,curdate(),curtime(),?,curdate(),?,curdate());`
		lrows, err := db.Exec(lquery,login_id,luser_id,luser_id)
		if err != nil {
			return 0,err
		} else {
			linsert,err=lrows.RowsAffected()
			if err != nil{
				return 0,err
			}
			return linsert, nil
		}
	}
}
}
