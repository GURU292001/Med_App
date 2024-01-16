package adduser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

// Adduser structure to recieve the data from the front
type AddUser_Req_struct struct {
	Current_User_id string `json:"current_User_id"`
	User_id         string `json:"user_id"`
	Password        string `json:"password"`
	Role            string `json:"role"`
}

// Adduser response structure
type AddUser_Resp_struct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

func AddUserApi(w http.ResponseWriter, r *http.Request) {
	log.Println("AddUserApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var result int64
		var lAdduserRec AddUser_Req_struct
		var lResponseRec AddUser_Resp_struct
		//read the body and get the form and to date
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "Adduser-01" + err.Error()
		} else {
			//UNmarshall
			err := json.Unmarshal(body, &lAdduserRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "Adduser-02" + err.Error()
			} else {
				result, err = Adduser_mtd(lAdduserRec)
				fmt.Println("result", result)
				if err != nil {
					log.Println(err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "Adduser-03" + err.Error()
				} else {
					//checking whether the insert operation is done or not
					if result == 1 {
						lResponseRec.Status = "S"
						lResponseRec.ErrMsg = "no Error"
					} else if result == 0 {
						lResponseRec.Status = "Exist"
						lResponseRec.ErrMsg = " Error"
					}
					//marshall
					res, err := json.Marshal(lResponseRec)
					if err != nil {
						log.Println(err)
						lResponseRec.Status = "E"
						lResponseRec.ErrMsg = "Adduser-04" + err.Error()
					} else {
						fmt.Fprintln(w, string(res))
						log.Println("AddUserApi(-)")
					}
				}
			}
		}

	}
}

//------------------methods--------------------------------------------------------

func Adduser_mtd(adduserRec AddUser_Req_struct) (int64, error) {
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return 0, err
	} else {
		defer db.Close()

		lcheck := `select * from medapp_login where user_id=? and password=? `
		lrows, ler := db.Query(lcheck, adduserRec.User_id, adduserRec.Password)
		if ler != nil {
			log.Println(ler)
		} else {
			if lrows.Next() {
				fmt.Println("exist")
				return 0, nil
			} else {
				var lquery = `insert into medapp_login 
			(user_id,password, role,created_by,created_date,updated_by,updated_date)
			values(?,?,?,?,curdate(),?,curdate());`
				lrows, err := db.Exec(lquery, adduserRec.User_id, adduserRec.Password, adduserRec.Role, adduserRec.Current_User_id, adduserRec.Current_User_id)
				if err != nil {
					return 0, err
				} else {
					insert, err := lrows.RowsAffected()
					if err != nil {
						return 0, err
					} else {
						return insert, nil
					}
				}
			}
		}
	}
	return 0, err
}
