package main

import (
	"fmt"
	adduser "medapp/adduser"
	logdata "medapp/logdata"
	loginHistory "medapp/loginhistory"
	login "medapp/loginpage"
	salesreport "medapp/salesReport"
	stockEntry "medapp/stockEntry"
	stockview "medapp/stockview"
	billEntry "medapp/billEntry"
	apexchart "medapp/apexchart"
	"net/http"
)

func main() {
	fmt.Println("Server started..(+)")
	http.HandleFunc("/stockview", stockview.StockviewApi)
	//---------stock entry--------
	//addstock
	http.HandleFunc("/addstock", stockEntry.AddStockIApi)
	//medicine_name dropdown
	http.HandleFunc("/dropdown", stockEntry.MedicineMasterApi)
	//Update stock in the medicine master
	http.HandleFunc("/updatestock", stockEntry.UpdateStockApi)
		//---------stock entry--------
	//adduser
	http.HandleFunc("/adduser", adduser.AddUserApi)
	//salesReport
	http.HandleFunc("/salesReport", salesreport.SalesReportApi)
	//loginHistory
	http.HandleFunc("/loginhistory", loginHistory.LoginHistoryApi)
	//log data
	http.HandleFunc("/logindata", logdata.LoginDataApi)
	http.HandleFunc("/logoutdata", logdata.LogoutDataApi)
	//login check
	http.HandleFunc("/logincheck", login.LoginCheckApi)
	//manager_dashboard
	http.HandleFunc("/mgrdashboard", login.MgrDashboardApi)
	//salesman _dashboard
	http.HandleFunc("/salesmandashboard", login.SalesmanDashboardApi)
	//---------------bill_Entry----------------------------------
	//quantitycheck
	http.HandleFunc("/quantitycheck",billEntry.QuantityCheckApi )
	//bill no update
	http.HandleFunc("/billno",billEntry.BillnoApi)
	//bill_master
	http.HandleFunc("/billmasterApi",billEntry.BillMasterApi )
	//bill_details
		http.HandleFunc("/billdetailsApi",billEntry.BillDetailsApi )
		//Apexchart
		http.HandleFunc("/apexchartApi",apexchart.ApexchartApi )
		http.ListenAndServe(":2904", nil)
		fmt.Println("Server ended..(-)")
	}
