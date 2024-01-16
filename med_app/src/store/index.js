import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    currentlogin:"",
    currentUser:"",
    currentUserRole:"",
    Login:[{email:"guru@gmail.com",password:"guru",role:"manager"},
    {email:"kumar@gmail.com",password:"guru",role:"salesman"},
    {email:"jaga@gmail.com",password:"guru",role:"admin"},
    {email:"vasanth@gmail.com",password:"guru",role:"inventry"}],
    
    LoginHistory:[{user:"guru@gmail.com",type:"login",date:"2023-12-16",time:"55"},
  ],
    
    MedicineMaster:[
      {medicineName:"Paracetamol",brand:"Calpol"},
      {medicineName:"Aspirin",brand:"Ecosprin"},
      {medicineName:"Metformin",brand:"Glucophage"},
      {medicineName:"Ampicillin",brand:"Roscillin"},
      {medicineName:"dolo5",brand:"parasitamal5"},
    ],

    Stock:[
      {medicineName:"Paracetamol",quantity:10,unitPrice:10},
      {medicineName:"Aspirin",quantity:10,unitPrice:10},
      {medicineName:"Metformin",quantity:10,unitPrice:10},
      {medicineName:"Ampicillin",quantity:10,unitPrice:10},
      {medicineName:"dolo5",quantity:10,unitPrice:10},
    ],

    BillMaster:[{billNo:1001,billDate:"2023-12-07",billAmount:1000,billGst:20,netPrice:120,userId:"guru@gmail.com"},
    {billNo:1002,billDate:"2023-12-07",billAmount:1500,billGst:20,netPrice:120,userId:"kumar@gmail.com"},
    {billNo:1006,billDate:"2023-12-07",billAmount:1000,billGst:20,netPrice:120,userId:"kumar@gmail.com"},
    {billNo:1003,billDate:"2023-12-15",billAmount:1000,billGst:20,netPrice:120,userId:"jaga@gmail.com"},
    {billNo:1004,billDate:"2023-12-20",billAmount:1000,billGst:20,netPrice:120,userId:"guru@gmail.com"},
    {billNo:1005,billDate:"2023-12-25",billAmount:1000,billGst:20,netPrice:120,userId:"mega@gmail.com"}
  ],


    BillDetails:[{billNo:1001,medicineName:'Paracetamol',quantity:15,unitPrice:20,amount:120},
      {billNo:1002,medicineName:'Aspirin',quantity:15,unitPrice:20,amount:120},
      {billNo:1003,medicineName:'Metformin',quantity:15,unitPrice:20,amount:120},
      {billNo:1004,medicineName:'Ampicillin',quantity:15,unitPrice:20,amount:120},
      {billNo:1005,medicineName:'dolo5',quantity:15,unitPrice:20,amount:120},
  ],
    
  },
  mutations: {
    addUser(state,obj){
      console.log(obj);
      state.Login.push(obj)
      console.log(state.Login);
    },
    // login(state,userId,role){
    //   state.currentUser=userId
    //   state.currentUserRole=role
    // },
    adddstock(state,obj){
      state.MedicineMaster.push(obj)
    },
    stockUpdate(state,obj){
      console.log(obj);
      const stock=state.Stock.find(medicine=>medicine.medicineName === obj.medicineName)
      if(stock)
      {
        console.log("if block");
        stock.quantity=obj.quantity
        stock.unitPrice=obj.unitPrice
        console.log(state.Stock);
      }
      else{
        console.log("else block");
      let stockobj={medicineName:obj.medicineName,quantity:obj.quantity,unitPrice:obj.unitPrice}
        state.Stock.push(stockobj)
        console.log(stockobj);
        console.log(state.Stock);
        state.MedicineMaster.push({medicineName:obj.medicineName,brand:obj.brand})
        console.log(state.MedicineMaster);

      }
    },
    billUpdate(state,obj){
          let obj1=[];
          for(let i of obj.preview)
          {
            let x={billNo:obj.billNo,medicineName:i.medicineName,quantity:i.quantity,unitPrice:i.unitPrice,amount:i.amount}
            obj1.push(x)
          }
          
          let obj2={billNo:obj.billNo,billDate:obj.billDate,billAmount:obj.billAmount,billGst:obj.billGst,netPrice:obj.netPrice,userId:obj.userId}
          state.BillDetails.push(obj1)
          state.BillMaster.push(obj2)
      },
      loginupdate(state,obj)
      {
        // console.log(obj);
        // console.log(obj.date);
        let value={user:obj.user,type:obj.type,date:obj.date,time:obj.time}
        // console.log(value);
        state.LoginHistory.push(value)
      }

  },
  actions: {},
  modules: {},
});
