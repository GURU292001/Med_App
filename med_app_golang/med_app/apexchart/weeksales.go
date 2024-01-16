package apexchart

import (
	"fmt"
	database "medapp/database"
)

type Weeksales_struct struct {
	Weekday string `json:"weekday"`
	Total float64 `json:"total"`
}


func Weeeksales_mtd() ([]Weeksales_struct, error) {
	fmt.Println("Weeeksales_mtd(+)")
	//local array
	var lweeksalesRec []Weeksales_struct
	db, err := database.LocalDbConnect()
	if err != nil {
		fmt.Println(err)
	} else {
		defer db.Close()
		lrows, err := db.Query(`
		select nvl(a.dayname,'-'), nvl(floor(b.net_price),0)
		from (select dayname(curdate()) as dayname
		union
		select dayname(subdate(curdate(),1))
		union
		select dayname(subdate(curdate(),2))
		union
		select dayname(subdate(curdate(),3))
		union
		select dayname(subdate(curdate(),4))
		union
		select dayname(subdate(curdate(),5))
		union
		select dayname(subdate(curdate(),6))) as a 
		left join 
		(select dayname(bill_date) dayname,sum(net_price) as net_price
		from  medapp_bill_master mbm 
		where mbm.bill_date in (select subdate(curdate(),1)
		union
		select subdate(curdate(),2)
		union
		select subdate(curdate(),3)
		union
		select subdate(curdate(),4)
		union
		select subdate(curdate(),5)
		union
		select subdate(curdate(),6))
		group by bill_date) as b
		on a.dayname = b.dayname;`)
		if err != nil {
			fmt.Println(err)
		} else {
			for lrows.Next(){
				var larr Weeksales_struct
				err:=lrows.Scan(&larr.Weekday,&larr.Total)
				if err!=nil{
					fmt.Println(err)
				}else{
					lweeksalesRec=append(lweeksalesRec,larr )
				}
			}
	fmt.Println("Weeeksales_mtd(-)")
			return lweeksalesRec,nil
		}
	}
	return nil, nil
}
/*

select nvl(a.dayname,'-'), nvl(b.net_price,0)
from (select dayname(curdate()) as dayname
union
select dayname(subdate(curdate(),1))
union
select dayname(subdate(curdate(),2))
union
select dayname(subdate(curdate(),3))
union
select dayname(subdate(curdate(),4))
union
select dayname(subdate(curdate(),5))
union
select dayname(subdate(curdate(),6))) as a 
left join 
(select dayname(bill_date) dayname,sum(net_price) as net_price
from  medapp_bill_master mbm 
where mbm.bill_date in (select subdate(curdate(),1)
union
select subdate(curdate(),2)
union
select subdate(curdate(),3)
union
select subdate(curdate(),4)
union
select subdate(curdate(),5)
union
select subdate(curdate(),6))
group by bill_date) as b
on a.dayname = b.dayname;
*/