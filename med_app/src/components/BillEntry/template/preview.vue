<template>
  <v-card width="800" class=" ms py-0 px-0">
    <v-simple-table >
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Medicine Name</th>
            <th class="text-left">Quantity</th>
            <th class="text-left">Amount</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in previewtable" :key="item.medicineName">
            <td>{{ item.medicineName }}</td>
            <td>{{ item.quantity }}</td>
            <td>{{ item.amount }}</td>
          </tr>
          <tr>
            <td></td>
            <td>Total :</td>
            <td>{{ total }}</td>
          </tr>
          <tr>
            <td></td>
            <td>GST:</td>
            <td>{{ gst }}</td>
          </tr>
          <tr>
            <td></td>
            <td>Net Payable:</td>
            <td>{{ netpayable }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <v-card tile class="d-flex justify-space-between">
        <v-btn class="red white--text mx-2 my-2" @click="close"> close</v-btn>
    <print class="mx-2 my-2"  :btname="btnname" @buttonEvent="printfun" />
    </v-card>
    <v-snackbar
      v-model="snackbar"
      :timeout="timeout"
    >
      {{ text }}

      <template v-slot:action="{ attrs }">
        <v-btn
          color="green"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-card>
</template>
<script>
import print from "../../Button/button.vue";
import Papa from "papaparse"
export default {
  data() {
    return {
        btnname:'print',
      // desserts: [{}],
      snackbar: false,
      text: 'start printed..',
      timeout: 2000,
    };
  },
  name: "printbtn",
  components: {
    print,
  },
  props: {
    gst: Number,
    netpayable: Number,
    total: Number,
    previewtable:Array,
  },
  methods: {
    printfun() {
      // Your CSV data (replace this with your actual data)
      const csvData = this.previewtable;

      // Convert the data array to a CSV string
      const csvString = Papa.unparse(csvData);

      // Create a Blob from the CSV string
      const blob = new Blob([csvString], { type: 'text/csv' });

      // Create a temporary link element
      const link = document.createElement('a');
      link.href = window.URL.createObjectURL(blob);

      // Set the filename for the download
      link.download = 'bill.csv';

      // Append the link to the document
      document.body.appendChild(link);

      // Trigger the click event to initiate the download
      link.click();

      // Remove the link from the document
      document.body.removeChild(link);
      
      this.snackbar=true
    },
    close() {
      this.$emit('closeEvent')
    },
  },    
};
</script>
