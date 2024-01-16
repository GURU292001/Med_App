package apexchart

import (
	"fmt"
	database "medapp/database"
)
type Mntsalesman_pfm_struct struct {
	User_id string `json:"user_id"`
	Total int `json:"total"`
}

func Monthlysalesprd_mtd() ([]Mntsalesman_pfm_struct, error) {
	fmt.Println("Monthlysales performance_mtd(+)")
	//local array
	var lMonthlysalesArr []Mntsalesman_pfm_struct
	db, err := database.LocalDbConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		defer db.Close()
		lrows, err := db.Query(`select nvl((select b.user_id
			from medapp_login b
			where b.login_id =a.login_id),'-') user_id,floor(sum(net_price))
			from medapp_bill_master a 
			where bill_date between DATE_FORMAT(NOW() ,'%Y-%m-01') AND NOW() 
			group by login_id;	`)
		if err != nil {
			fmt.Println(err)
		} else {
			for lrows.Next(){
				var larr Mntsalesman_pfm_struct
				err:=lrows.Scan(&larr.User_id,&larr.Total)
				if err!=nil{
					fmt.Println(err)
				}else{
					lMonthlysalesArr=append(lMonthlysalesArr,larr )
				}
			}
			fmt.Println("Monthlysales performance_mtd(-)")
			return lMonthlysalesArr,nil

		}
	}
	return nil, nil
}
/*
select (select b.user_id
	from medapp_login b
	where b.login_id =a.login_id) user_id,nvl(sum(net_price),0)
		from medapp_bill_master a 
		where bill_date between DATE_FORMAT(NOW() ,'%Y-%m-01') AND NOW() 
		group by login_id;
*/
