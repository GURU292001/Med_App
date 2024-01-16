package apexchart

import (
	"fmt"
	database "medapp/database"
)

type  Todaysalesman_pmf_struct  struct {
	User_id string `json:"user_id"`
	Total int `json:"total"`
}

func Todaysales_mtd() ([]Todaysalesman_pmf_struct, error) {
	fmt.Println("Todaysales_mtd(+)")
	//local array
	var lTdysalespmfArr []Todaysalesman_pmf_struct
	db, err := database.LocalDbConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		defer db.Close()
		lrows, err := db.Query(`select (select b.user_id
			from medapp_login b
			where b.login_id =a.login_id) user_id,floor(sum(net_price))
				from medapp_bill_master a 
				where bill_date =curdate()
				group by login_id  ;`)
		if err != nil {
			fmt.Println(err)
		} else {
			for lrows.Next(){
				var larr Todaysalesman_pmf_struct
				err:=lrows.Scan(&larr.User_id,&larr.Total)
				if err!=nil{
					fmt.Println(err)
				}else{
					lTdysalespmfArr=append(lTdysalespmfArr,larr )
				}
			}
	fmt.Println("Todaysales_mtd(-)")
			return lTdysalespmfArr,nil
		}
	}
	return nil, nil
}

/*
select (select b.user_id
	from medapp_login b
	where b.login_id =a.login_id) user_id ,sum(net_price)
		from medapp_bill_master a 
		where bill_date =curdate()
		group by login_id  ;
*/
