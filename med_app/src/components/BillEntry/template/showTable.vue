<template>
  <v-container flat  class="d-flex flex-column  mt-16 ">
    <v-card  elevation="0" class=" d-flex justify-space-between ">
      <div class=" d-flex justify-space-between ">
      <div class="">
    <v-dialog
      v-model="dialog"
      width="800"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
        elevation="0"
          v-bind="attrs"
          v-on="on"
          class="primary mx-4"
        >
        preview
        <!-- <previewbtn btname="preview"/> -->
        </v-btn>
      </template>

      <v-card>
        <v-card-title class="text-h5 grey lighten-2">
        Preview
        </v-card-title>

        <v-card-text>
            <previewevent  :gst="GST" :netpayable="netpayable" :total="total" :previewtable="previewtable" @closeEvent="closeEvent" />
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
        </v-card-text>
      </v-card>
    </v-dialog>
   </div> 
        <savebtn  btname="save" @buttonEvent="save"/>
      </div>
        <div class=" d-flex justify-start ">
            <div  class="black--text font-weight-regular text-body-2 mx-7">BILLNO :{{ billno }}</div>
            <div  class="black--text font-weight-regular text-body-2 mx-7">DATE :
              {{date.split('-').reverse().join('-')}}</div>
            <div  class="black--text font-weight-regular text-body-2 mx-7">TOTAL(₹):{{total}}</div>
            <div  class="black--text font-weight-regular text-body-2 mx-7">GST(₹) :{{GST}}</div>
            <div  class="black--text font-weight-regular text-body-2 mx-7">NET PAYABLE(₹) :{{netpayable}}</div>
        </div>
        </v-card>
    <v-card outlined color="primary" class=" mt-10">

    <v-data-table
      :headers="headers"
      :items="stock"
      item-key="medicinename"
      :search="search"
      :custom-filter="filterOnlyCapsText"
    >
      <template v-slot:top>
       <div class="d-flex">
       <v-col cols="10">
        <v-text-field
        
          v-model="search"
          label="Search"
          class="mx-4 "
        ></v-text-field>
       </v-col>
       <v-col >
        <v-btn color="red" class="white--text" @click="deletebill">delete</v-btn>
       </v-col>
       </div>
      </template>

      <template v-slot:append>
        <tr>
          <td></td>
          <!-- <td> -->
            <!-- <v-text-field
              v-model="Brand"
              type="number"
              label="Less than"
            ></v-text-field> -->
          <!-- </td> -->
          <!-- <td colspan="4"></td> -->
        </tr>
      </template>
    </v-data-table>
  </v-card>
  </v-container>
</template>
<script>

import savebtn from "../../Button/button.vue";
import previewevent from "../template/preview.vue"
export default {
    // name:"tablecomp",
  components: {
    savebtn,
    previewevent,
  },
  props:{
    previewtable:Array,
    billno:Number,
        date:String,
        total:Number,
        GST:Number,
        netpayable:Number,
        headers:Array,
        stock:Array,
  },
  data() {
    return {
      dialog:false,
        // netpayable:0,
      search: "",
      // Brand: "",
      // stock:this.$store.state.Stock,
      MedicineMaster:this.$store.state.MedicineMaster,
      snackbar: false,
      text: 'start printed..',
      timeout: 2000,
    };
  },
  computed: {
    // Brand(){
    //   // for(let i in this.MedicineMaster)

    //   return this.Brand;
    // },
  
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
    save(){
      this.$emit('savebutton')
    },
    closeEvent()
    {
      this.dialog=false
    },
    deletebill(){
      this.$emit('deletebill')
    }
  },

 
};
</script>
