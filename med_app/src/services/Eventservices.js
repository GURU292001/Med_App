import axios from 'axios'
//this is the client server localhost API
const baseApiClient = axios.create({
    baseURL:`http://localhost:2904`,
    withCredentials:false,//this is default
    headers:{
        Accept:'application/json',
        'Content-Type':'application/json',
    },
})

export default{
    //this is for add api
    // getvalue(value){
    //     console.log("value=",value)
    //     // const hdr={
    //     //     headers:{"VALUE":JSON.stringify(value)}
    //     // }
    //     return baseApiClient.post('/api1',value)
    // },
    //this for get details by using IFSC
    // getDetails(head,body){
    //     console.log("inside get details")
    //     console.log("head:",head)
    //     console.log("body:",body)

    //     const hdr={
    //         headers:{"HEADER":JSON.stringify(head)}
    //     }
    //     return ifscClient.put('/ifsc',body,hdr)
    // },
    //This for send csv from front to API
    // sendCsvFile(file){
    //     const hdr={
    //         headers:{
    //             'Content-Type':'multipart/form-data',
    //         }
    //     }
    //     console.log(file);
    //     console.log("inside sendcsv file",file);
    //     return baseApiClient.post('/csvcalc',file,hdr)
    // },
    //Chocolate name and quantity  send to  API
    //this is for calculate the price 
    // TotalPrice(body){
    //     console.log("inside totalPrice:",body)
    //     return baseApiClient.post('/fruitApi',body)
    // },
//---------------------------------------stockview-------------------------------------------
    stockview(){
        return baseApiClient.post('/stockview')
    },
//-----------------------------------stockEntry------------------------
//add stock
addstock(body){
    return baseApiClient.post('/addstock',body)
},
//select option in updatestock
showdorpdown(){
    return baseApiClient.post('/dropdown')
},
//update qty and price
updatestock(body){
    console.log("body:",body);
    return baseApiClient.post('/updatestock',body)
},

//-----------------------------------stockEntry------------------------
//adduser
adduser(body){
    return baseApiClient.post('/adduser',body)
},
//salesReport
searchSalesReport(body){
    return baseApiClient.post('/salesReport',body)
},
//login history
loginHistory(){
    return baseApiClient.post('/loginhistory')
},
//log_Historyd
logindata(head){
    console.log(typeof(head))
    const hdr={
                headers:{"HEADER":head}
            }
    return baseApiClient.post('/logindata',"",hdr)
},
logoutdata(head){
    const hdr={
                headers:{"HEADER":head}
            }
    return baseApiClient.post('/logoutdata',"",hdr)
},
//login check api
logincheck(body){
    return baseApiClient.post('/logincheck',body)
},
//managerdashboard
mgrdashboard(){
    return baseApiClient.post('/mgrdashboard')
},
//salesman dashboard
salesmandashboard(head){
    console.log("head",head)
    const hdr={
        headers:{"HEADER":head}
    }
    return baseApiClient.post('/salesmandashboard',"",hdr)
},
//billEntry
//quantity check
qtycheck(body){
    return baseApiClient.post('/quantitycheck',body)
},
//billmaster
billmasterinsert(body){
    return baseApiClient.post('/billmasterApi',body)
},
//billdetails
billdetailsinsert(body){
    return baseApiClient.post('/billdetailsApi',body)
},
//quantity update
getbillno()
{
    return baseApiClient.get('/billno')
},
//apex chart
getchartvalues(){
    return baseApiClient.get('/apexchartApi')
}
 
}
