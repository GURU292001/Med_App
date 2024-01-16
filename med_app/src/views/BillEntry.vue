<template>
  <v-container class="mt-16 pt-12">
    <v-form lazy-validation>
      <Addentry class="primary" @addbutton="addbutton" />
      <showtable
        :previewtable="previewtable"
        :billno="billno"
        :date="date"
        :GST="gst"
        :total="total"
        :netpayable="netpayable"
        :headers="headers"
        :stock="EntrdqtyTemp"
        @savebutton="savebtn"
        @deletebill="deletebill"
        @editItem="editItem"
        @deleteItem="deleteItem"
      />
    </v-form>
    <v-snackbar v-model="snackbar" :timeout="2000" :color="color">
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn   color="white"
          text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>
<script>
import Addentry from "../components/BillEntry/template/addEntry.vue";
import showtable from "../components/BillEntry/template/showTable.vue";
import EventServices from "../services/Eventservices";
export default {
  data() {
    return {
      dialoge: false,
      // billno: 1001,
      BillMaster: this.$store.state.BillMaster,
      MedicineMaster: this.$store.state.MedicineMaster,
      stock: this.$store.state.Stock,
      snackbar: false,
      text: "",
      color: "",

      date: "",
      previewtable: [],
      //table to show in front
      table: [],
      //user entered qty in added list
      EntrdqtyTemp: [],
      //avalableqty
      AvaildqtyTemp: [],
      brand: 0,
      unitPrice: 0,
      amount: 0,
      selecteditem: "",
      quantity: 0,
      //bill from api
      billnoFrmApi:0
    };
  },
  name: "addcomp",
  components: {
    Addentry,
    showtable,
  },

  computed: {
    headers() {
      return [
        {
          text: "Medicine Name",
          align: "start",
          sortable: false,
          value: "medicineName",
        },
        {
          text: "Brand",
          value: "brand",
          filter: (value) => {
            if (!this.Brand) return true;
            return value < parseInt(this.Brand);
          },
        },
        { text: "Quantity", value: "quantity" },
        { text: "Amount", value: "amount" },
      ];
    },
    gst() {
      return Math.round(this.total * (18 / 100));
    },
    total() {
      let sum = 0;
      for (let i of this.EntrdqtyTemp) {
        console.log("i.quantity",i.quantity);
        console.log("i.unitprice",i.unitprice);
        sum += i.amount;
        console.log("sum:",sum);
      }
      return sum;
    },
    netpayable() {
      return this.total + this.gst;
    },
    preview() {
      return this.total;
    },
    billno() {
      let x 
      x=this.billnoFrmApi
    return ++x
    },
  },
  mounted() {
    let currentDate = new Date().toJSON().slice(0, 10);
    this.date = currentDate;
    if (this.$store.state.currentUser == "") {
      this.$router.push("/");
    }
    EventServices.stockview()
      .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          this.tabledata = response.data.stockview_array;
          console.log("tabledata", this.tabledata);
          for (let i of this.tabledata) {
            for (let ii of i) {
              var obj = {
                medicineName: ii.medicine_Name,
                brand: ii.brand,
                quantity: ii.quantity,
                unitPrice: ii.unit_Price,
              };
              this.AvaildqtyTemp.push(obj);
            }
          }
        } else {
          this.result = response.data.errMsg;
        }
      })
      .catch((error) => {
        console.log(error);
      });
    
  },
  methods: {
    savebtn() {
     if(this.total!=0){
      var body={
        User_id:this.$store.state.currentUser,
        login_id:this.$store.state.currentlogin,
        bill_No:this.billno,
        bill_Amount:this.total,
        bill_Gst:this.gst,
        net_Price:this.netpayable
      }
      EventServices.billmasterinsert(body)
        .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          console.log("response.data.status",response.data.status);
        } else {
          console.log("response.data.errMsg",response.data.errMsg);
          this.text = response.data.errMsg;
            this.color = "red";
            this.snackbar = true;
        }
      })
      .catch((error) => {
        console.log(error);
      });
     }
     let lmeddetailArr =[]
     for(let i of this.EntrdqtyTemp){
      lmeddetailArr.push({
        medicine_Name: i.medicineName, 
        brand: i.brand, 
        quantity: i.quantity, 
        unit_Price: i.unitPrice,
         amount: i.amount 
      })
     }
     var data={
      bill_no:this.billno,
      user_id:this.$store.state.currentUser,
      bill_detailsArr:lmeddetailArr
     }
     EventServices.billdetailsinsert(data)
        .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          console.log("response.data.status",response.data);
          this.EntrdqtyTemp=[],
      this.text = "Saved Successfully";
        this.snackbar = true;
        this.color = "success";
        } else {
          console.log("response.data.errMsg",response.data.errMsg);
          this.EntrdqtyTemp=[],
      this.text = response.data.errMsg;
        this.snackbar = true;
        this.color = "failure";
        }
      })
      .catch((error) => {
        console.log(error);
      });
      this.EntrdqtyTemp=[],
      this.text = "Saved Successfully";
        this.snackbar = true;
        this.color = "success";
    },
    addbutton(selecteditem, quantity) {
      if (selecteditem == "" || quantity == 0) {
        this.text = "Enter the Valid Data....";
        this.color = "red";
        this.snackbar = true;
      } else{
         //checking whether the table having medicine already
      let check = this.EntrdqtyTemp.find(
        (ele) => ele.medicineName == selecteditem
      );
      //If present below will executed
      if (check) {
        for (let i of this.AvaildqtyTemp) {
          if (check.medicineName == i.medicineName) {
            if (check.quantity + quantity <= i.quantity) {
              check.quantity = check.quantity + quantity;
              check.amount=(check.quantity*check.unitPrice)
              this.text = `Item Added Successfully`;
              this.color = "success";
              this.snackbar = true;
              //update the preview table
              var prevwExist =this.previewtable.find(ele=>ele.medicineName=check.medicineName)
              if(prevwExist){
                prevwExist.amount=check.amount
              }
            } else {
              console.log( "i.quantity-check.quantity",  i.quantity - check.quantity  );
              this.text = `only ${i.quantity - check.quantity} left`;
              this.color = "red";
              this.snackbar = true;
            }
          }
        }
      } 
      //if medicine not already exist below will execute
      else {
        for (let i of this.AvaildqtyTemp) {
          if ((i.medicineName == selecteditem) & (i.quantity >= quantity)) {
            var object = {
              medicineName: selecteditem,
              brand: i.brand,
              quantity: quantity,
              unitPrice: i.unitPrice,
              amount:(quantity*i.unitPrice)
            };
            this.EntrdqtyTemp.push(object);
            var prvobj = {
              medicineName: object.medicineName,
              quantity: object.quantity,
             amount:object.amount,
            };
            this.previewtable.push(prvobj)
            this.text = `Item Added Successfully`;
            this.color = "success";
            this.snackbar = true;
          } else if (
            (i.medicineName == selecteditem) &
            (i.quantity < quantity)
          ) {
            console.log("i.quantity", i.quantity);
            this.text = `only ${i.quantity} left`;
            this.color = "red";
            this.snackbar = true;
          }
        }
      }
        //to get bill number
        EventServices.getbillno()
        .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          console.log("response.data.status",response.data.status);
          this.billnoFrmApi=response.data.bill_no
        } else {
          console.log("response.data.errMsg",response.data.errMsg);
          this.text = response.data.errMsg;
            this.color = "red";
            this.snackbar = true;
        }
      })
      .catch((error) => {
        console.log(error);
      });
      }
     
      
    },
    deletebill() {
      this.EntrdqtyTemp = [];
      this.previewtable = [];
      this.text = `bill deletd successfully`;
      this.color = "red";
      this.snackbar = true;
    },
    editItem(medicineName, quantity) {
      console.log(medicineName);
      console.log(quantity);
    },
    deleteItem(medicineName) {
      console.log(medicineName);
    },
  },
};
</script>
