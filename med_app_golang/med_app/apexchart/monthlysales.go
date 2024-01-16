package apexchart

import (
	"encoding/json"
	"fmt"
	"log"
	database "medapp/database"
	"net/http"
)

type Mnt_sales_struct struct {
	Monthname string  `json:"monthname"`
	Total     float64 `json:"total"`
}

type Saleschart_Resp_struct struct {
	Status          string                     `json:"status"`
	ErrMsg          string                     `json:"errMsg"`
	MntySales       []Mnt_sales_struct         `json:"mntySales"`
	Weeksales       []Weeksales_struct         `json:"weeksales"`
	MntySalmanperfm []Mntsalesman_pfm_struct   `json:"mntySalmanperfm"`
	TdySalmanperfm  []Todaysalesman_pmf_struct `json:"tdySalmanperfm"`
}

func ApexchartApi(w http.ResponseWriter, r *http.Request) {
	log.Println(" Api (+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "GET" {
		var lMntySales []Mnt_sales_struct
		var lweeksales []Weeksales_struct
		var lMntySalmanperfm []Mntsalesman_pfm_struct
		var lTdySalmanperfm []Todaysalesman_pmf_struct
		var lResponseRec Saleschart_Resp_struct
		//daily sales trend
		lweeksales, err := Weeeksales_mtd()
		if err != nil {
			fmt.Println(err)
			lResponseRec.Status = "E"
			lResponseRec.ErrMsg = "Weeeksales_mtd"+err.Error()
		} else {
			lResponseRec.Weeksales = lweeksales
			lResponseRec.Status = "S"
			lResponseRec.ErrMsg = "no Error"
			// monthly sales trend
			lMntySales, err = Monthlysales_mtd()
			if err != nil {
				fmt.Println(err)
				lResponseRec.Status = "E"
				lResponseRec.ErrMsg = "Monthlysales_mtd"+err.Error()
			} else {
				lResponseRec.Status = "S"
				lResponseRec.ErrMsg = "no Error"
				lResponseRec.MntySales = lMntySales
				//today salesman Performance
				lMntySalmanperfm, err = Monthlysalesprd_mtd()
				if err != nil {
					fmt.Println(err)
					lResponseRec.Status = "E"
					lResponseRec.ErrMsg =  "Monthlysalesperfm_mtd"+err.Error()
				} else {
					lResponseRec.Status = "S"
					lResponseRec.ErrMsg = "no Error"
					//current month salesman performance
					lResponseRec.MntySalmanperfm = lMntySalmanperfm
					lTdySalmanperfm, err = Todaysales_mtd()
					if err != nil {
						fmt.Println(err)
						lResponseRec.Status = "E"
						lResponseRec.ErrMsg =  "Todaysales_mtd"+err.Error()
						} else {
							lResponseRec.Status = "S"
							lResponseRec.ErrMsg = "no Error"
							lResponseRec.TdySalmanperfm = lTdySalmanperfm
							//marshall
							res, err := json.Marshal(lResponseRec)
							if err != nil {
								log.Println(err)
								lResponseRec.Status = "E"
								lResponseRec.ErrMsg = "Error"
								} else {
									fmt.Fprintln(w, string(res))
									log.Println("Api (-)")
								}
					}
				}
			}
		}
		
		
	}
}

func Monthlysales_mtd() ([]Mnt_sales_struct, error) {
	fmt.Println("MonthlysalesApi(+)")
	//local array
	var lMntsalesRecr []Mnt_sales_struct
	db, err := database.LocalDbConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		defer db.Close()
		lrows, err := db.Query(`select nvl(a.month_name,0) as month_name,nvl(floor(b.net_price),0) price
		from (select monthname(CURRENT_DATE) as month_name
		union
		select monthname(CURRENT_DATE-interval 1 month)
		union
		select monthname(CURRENT_DATE-interval 2 month)
		union
		select monthname(CURRENT_DATE-interval 3 month)
		union
		select monthname(CURRENT_DATE-interval 4 month)
		union
		select monthname(CURRENT_DATE-interval 5 month)
		union
		select monthname(CURRENT_DATE-interval 6 month)
		union
		select monthname(CURRENT_DATE-interval 7 month)
		union
		select monthname(CURRENT_DATE-interval 8 month)
		union
		select monthname(CURRENT_DATE-interval 9 month)
		union
		select monthname(CURRENT_DATE-interval 10 month)
		union
		select monthname(CURRENT_DATE-interval 11 month))as a left join (select monthname(bill_date) as month_name, sum(net_price) net_price
		from medapp_bill_master mbm
		where bill_date between concat(year(curdate()-interval 11 month),'-',month(curdate()-interval 11 month),'-01') and curdate()
		group by monthname(bill_date) ) as b
		on a.month_name=b.month_name;`)
		if err != nil {
			fmt.Println(err)
		} else {
			for lrows.Next() {
				var larr Mnt_sales_struct
				err := lrows.Scan(&larr.Monthname, &larr.Total)
				if err != nil {
					fmt.Println(err)
				} else {
					lMntsalesRecr = append(lMntsalesRecr, larr)
				}
			}
	fmt.Println("MonthlysalesApi(-)")
			return lMntsalesRecr, nil
		}
	}
	return nil, nil
}

/*
select nvl(a.month_name,0) as month_name,nvl(b.net_price,0) price
from (select monthname(CURRENT_DATE) as month_name
union
select monthname(CURRENT_DATE-interval 1 month)
union
select monthname(CURRENT_DATE-interval 2 month)
union
select monthname(CURRENT_DATE-interval 3 month)
union
select monthname(CURRENT_DATE-interval 4 month)
union
select monthname(CURRENT_DATE-interval 5 month)
union
select monthname(CURRENT_DATE-interval 6 month)
union
select monthname(CURRENT_DATE-interval 7 month)
union
select monthname(CURRENT_DATE-interval 8 month)
union
select monthname(CURRENT_DATE-interval 9 month)
union
select monthname(CURRENT_DATE-interval 10 month)
union
select monthname(CURRENT_DATE-interval 11 month))as a left join (select monthname(bill_date) as month_name, sum(net_price) net_price
from medapp_bill_master mbm
where bill_date between concat(year(curdate()-interval 11 month),'-',month(curdate()-interval 11 month),'-01') and curdate()
group by monthname(bill_date) ) as b
on a.month_name=b.month_name;
*/
