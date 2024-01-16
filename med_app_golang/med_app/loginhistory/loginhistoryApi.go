package loginhistory

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	database "medapp/database"
	"net/http"
)

// loginHistory structure to recieve from the stock table
type Login_Hsty_struct struct {
	User_id     string `json:"user_id"`
	Login_Date  string `json:"login_date"`
	Login_Time  string `json:"login_time"`
	Logout_Date string `json:"logout_date"`
	Logout_Time string `json:"logout_time"`
}

// loginHistory response structure
type Login_Hsty_Resp_struct struct {
	Status           string                `json:"status"`
	ErrMsg           string                `json:"errMsg"`
	Login_Hsty_Array [][]Login_Hsty_struct `json:"login_Hsty_Array"`
}

func LoginHistoryApi(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHistoryApi Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTION")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "POST" {
		var lLoginhistoryArr [][]Login_Hsty_struct
		var lResponseRec Login_Hsty_Resp_struct
		lLoginhistoryArr, err := Loginhistory_mtd()
		if err != nil {
			log.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "LoginHistoryApi-01"+err.Error()
		} else {
			lResponseRec.Login_Hsty_Array = lLoginhistoryArr
			lResponseRec.Status = "S"
			lResponseRec.ErrMsg = "no Error"
			//marshall
			res, err := json.Marshal(lResponseRec)
			if err != nil {
				log.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg =  "LoginHistoryApi-02"+err.Error()
			} else {
				fmt.Fprintln(w, string(res))
				log.Println("LoginHistoryApi Api (-)")
			}
		}

	}
}

//------------------methods--------------------------------------------------------

func Loginhistory_mtd() ([][]Login_Hsty_struct, error) {
	//local variable to store the result rows
	var lLoginhistoryArr [][]Login_Hsty_struct
	var lLoginhistoryRec Login_Hsty_struct
	//LocalDbConnect functionCall..
	db, err := database.LocalDbConnect()
	if err != nil {
		return nil, errors.New("query not executed")
	} else {
		defer db.Close()// nvl(mlh.logout_time,'-')
		var lquery = `select nvl(ml.user_id,'-') ,nvl(mlh.login_date,'-') ,nvl(mlh.login_time ,'-'), nvl(mlh.logout_date  ,'-') ,nvl(mlh.logout_time,'-')
		from medapp_login ml ,medapp_login_history mlh 
		where ml.login_id =mlh.login_id ;`
		lrows, err := db.Query(lquery)
		if err != nil {
			return nil, errors.New("query not executed")
		} else {
			for lrows.Next() {
				err := lrows.Scan(&lLoginhistoryRec.User_id, &lLoginhistoryRec.Login_Date, &lLoginhistoryRec.Login_Time, &lLoginhistoryRec.Logout_Date, &lLoginhistoryRec.Logout_Time)
				if err != nil {
					log.Println(err)
				} else {
					var lArr []Login_Hsty_struct
					lArr = append(lArr, lLoginhistoryRec)
					lLoginhistoryArr = append(lLoginhistoryArr, lArr)
				}
			}
			return lLoginhistoryArr, nil
		}
	}
}
