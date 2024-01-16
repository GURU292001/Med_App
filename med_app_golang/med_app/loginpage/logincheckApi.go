package loginpage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	database "medapp/database"
	"net/http"
)

// login structure to recieve from the stock table
type Login_Req_struct struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

//login structure to recieve the login table
type Login_mtd_struct struct {
	Login_id int    `json:"login_id"`
	User_id  string `json:"user_id"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

//  login response structure
type Login_Resp_struct struct {
	Status   string `json:"status"`
	ErrMsg   string `json:"errMsg"`
	Login_id int    `json:"login_id"`
	Role     string `json:"role"`
}

func LoginCheckApi(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginCheckApi  (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		//local variable
		var lLoginReqRec Login_Req_struct
		var lLoginArr []Login_mtd_struct
		var lResponseRec Login_Resp_struct
		var exists = false
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "LoginCheckApi-01" + err.Error()
		} else {
			//unmarshal
			err := json.Unmarshal(body, &lLoginReqRec)
			//method calling
			lLoginArr, err = Login()
			// fmt.Println("lLoginArr:",lLoginArr)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "LoginCheckApi-01" + err.Error()
			} else {
				for _, value := range lLoginArr {
					// fmt.Println(i,"=",value)
					if lLoginReqRec.User_id == value.User_id {
						exists = true
						lResponseRec.Login_id = value.Login_id
						lResponseRec.Role = value.Role
					}
				}
				//checkin whether the user is present or not
				if exists {
					lResponseRec.Status = "S"
					lResponseRec.ErrMsg = "No error"
				} else {
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "User Not Exists !"
				}
				//marshall
				res, err := json.Marshal(lResponseRec)
				if err != nil {
					log.Println(err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg = "LoginCheckApi-03" + err.Error()
				} else {
					fmt.Fprintln(w, string(res))
					log.Println("login check Api (-)")
				}
			}
		}
	}
}

//------------------methods--------------------------------------------------------

func Login() ([]Login_mtd_struct, error) {
	fmt.Println("login mtd(+)")
	//local variable to store the result rows
	var lLoginRec Login_mtd_struct
	var lArr []Login_mtd_struct
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return nil, errors.New("query not executed")
	} else {
		defer db.Close()
		lrows, err := db.Query(`select nvl(login_id,'-') ,nvl(user_id,'-') ,nvl(password,'-') ,nvl(role,'-')
		from medapp_login ;`)
		if err != nil {
			return nil, err
		} else {
			for lrows.Next() {
				err := lrows.Scan(&lLoginRec.Login_id, &lLoginRec.User_id, &lLoginRec.Password, &lLoginRec.Role)
				if err != nil {
					log.Println(err)
					return nil, err
				} else {
					lArr = append(lArr, lLoginRec)
				}
			}
			fmt.Println("login mtd(-)")
			return lArr, nil
		}
	}
}
