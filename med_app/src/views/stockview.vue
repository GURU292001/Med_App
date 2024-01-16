<template>
    <v-card outlined color="primary" class="mx-16 mt-10">
    <v-card tile flat class="d-flex justify-end px-2 py-2">
      <downloadbtn btname="download" @buttonEvent="download" />
    </v-card>
    <v-data-table
      :headers="headers"
      :items="stock"
      item-key="medicineName"
      class="elevation-1"
      :search="search"
      :custom-filter="filterOnlyCapsText"
    >
      <template v-slot:top>
        <v-text-field
          v-model="search"
          label="Search"
          class="mx-4"
        ></v-text-field>
      </template>

      <template v-slot:append>
        <tr>
          <td></td>
          <td>
            <v-text-field v-model="brand" type="string"></v-text-field>
          </td>
          <td colspan="4"></td>
        </tr>
      </template>
    </v-data-table>
    <!-- <v-card tile flat class="d-flex justify-end px-2 py-2">
      <downloadbtn btname="download" @buttonEvent="download" />
    </v-card> -->
    <v-snackbar v-model="snackbar">
      {{ text }}

      <template v-slot:action="{ attrs }" >
        <v-btn color="pink" text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-card>
</template>
<script>
import downloadbtn from "../components/Button/button.vue";
import EventServices from "../services/Eventservices";
import Papa from "papaparse";
export default {
  data() {
    return {
      search: "",
      snackbar: false,
      text: '',
      stock: [],
    };
  },
  components: { downloadbtn },
  computed: {
    headers() {
      console.log("mounted");
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
        },
        { text: "Quantity", value: "quantity" },
        { text: "Unit Price(₹)", value: "unitPrice" },
      ];
    },
  },
  methods: {
    filterOnlyCapsText(value, search) {
      return (
        value != null &&
        search != null &&
        typeof value === "string" &&
        value.toString().toLocaleUpperCase().indexOf(search.toUpperCase()) !==
          -1
      );
    },
    download() {
      // Your CSV data (replace this with your actual data)
      const csvData = this.stock;

      // Convert the data array to a CSV string
      const csvString = Papa.unparse(csvData);

      // Create a Blob from the CSV string
      const blob = new Blob([csvString], { type: "text/csv" });

      // Create a temporary link element
      const link = document.createElement("a");
      link.href = window.URL.createObjectURL(blob);

      // Set the filename for the download
      link.download = "stock.csv";

      // Append the link to the document
      document.body.appendChild(link);

      // Trigger the click event to initiate the download
      link.click();

      // Remove the link from the document
      document.body.removeChild(link);
      this.text=`download Completed`
      this.snackbar = true;
    },
  },
  mounted() {
    if(this.$store.state.currentUser=='')
        {
            this.$router.push('/')
        }
    //------------------------------stock view---------------------------------------------------
    EventServices.stockview()
      .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          this.tabledata = response.data.stockview_array;
          console.log("tabledata",this.tabledata)
          for (let i of this.tabledata) {
           for (let ii of i){
            var obj = {
              medicineName: ii.medicine_Name,
              brand: ii.brand,
              quantity: ii.quantity,
              unitPrice: ii.unit_Price,
            };
            this.stock.push(obj);
           }
          }
        } else {
          this.text=response.data.errMsg
      this.snackbar = true;
        }
      })
      .catch((error) => {
        console.log(error);
      });
    
  },
};
</script>
