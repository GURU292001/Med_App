<template>
  <v-container outlined elevation="2" class="mt-16 pt-16">
    <searchdate @searchevent="searchevent" />
    <tablecomp :details="filter" :headers="headers" @downloadevent="downloadevent"/>
    <!-- {{ tabledata }} -->
    <!-- {{ details }} -->
  </v-container>
</template>

<script>
import searchdate from "../components/SalesReport/template/search.vue";
import tablecomp from "../components/SalesReport/template/table.vue";
import Eventservices from "../services/Eventservices";
import Papa from 'papaparse';
export default {
  data() {
    return {
      dates:{
        from:"",
        to:""
      },
      tabledata:[],
      BillDetails: this.$store.state.BillDetails,
      BillMaster: this.$store.state.BillMaster,
     //filter used for data shown in the table
      filter: [],
      //temp- used for custom array object
      temp: [],
    };
  },
  name: "comp",
  components: {
    searchdate,
    tablecomp,
  },
  mounted() {
    if (this.$store.state.currentUser == "") {
      this.$router.push("/");
    }
  },
  methods: {
    searchevent(from, to) {
      this.dates.from=from
      this.dates.to=to
      Eventservices.searchSalesReport(this.dates)
        .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          this.tabledata = response.data.salesReport_Array;
          console.log("tabledata",this.tabledata)
          for (let i of this.tabledata) {
           for (let ii of i){
            var obj = {
              bill_No:ii.bill_No,
              bill_Date:ii.bill_Date,
              medicine_Name:ii.medicine_Name,
              quantity:ii.quantity,
              amount:ii.amount
            };
            this.filter.push(obj);
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
    
    downloadevent(){
       // Your CSV data (replace this with your actual data)
       const csvData =this.filter

      // Convert the data array to a CSV string
      const csvString = Papa.unparse(csvData);

      // Create a Blob from the CSV string
      const blob = new Blob([csvString], { type: 'text/csv' });

      // Create a temporary link element
      const link = document.createElement('a');
      link.href = window.URL.createObjectURL(blob);

      // Set the filename for the download
      link.download = 'salesReport.csv';

      // Append the link to the document
      document.body.appendChild(link);

      // Trigger the click event to initiate the download
      link.click();

      // Remove the link from the document
      document.body.removeChild(link);
    }
  },
  computed: {
    headers() {
      return [
        {
          text: "Bill No",
          align: "start",
          sortable: false,
          value: "bill_No",
        },
        {
          text: "Bill Date",
          value: "bill_Date",
          // filter: value => {
          //   if (!this.Brand) return true

          //   return value < parseInt(this.Brand)
          // },
        },
        { text: "Medicine Name", value: "medicine_Name" },
        { text: "Quantity", value: "quantity" },
        { text: "Amount(₹)", value: "amount" },
      ];
    },
    details(){
      let datas = []
      for(let i of this.BillDetails)
      {
        datas.push({billNo : i.billNo , medicineName: i.medicineName, quantity:i.quantity, amount:i.amount})
      }
      for(let i of this.BillMaster)
      {
        for(let j of datas){
          if(j.billNo==i.billNo)
        {
          j.billDate=(i.billDate)
        }
        }
      }
      console.log(datas);
      return datas
    }
  },
};
</script>
