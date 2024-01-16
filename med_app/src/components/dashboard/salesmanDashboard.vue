<template>
  <v-container class="mt-16">
    <v-card class="primary" width="450" height="170">
      <v-card-title class="text-h5 white--text font-weight-black mb-5">
        Your Today Sales
      </v-card-title>
      <v-card-subtitle class="white--text text-h3">
        ₹ {{ todaysales }}
        <v-icon v-if="high" color="green" size="40">north</v-icon>
        <v-icon v-if="low" color="red" size="40">south</v-icon>
        <div class="text-h6 d-inline-block">({{ percentage }} %)</div>
      </v-card-subtitle>
    </v-card>
  </v-container>
</template>
<script>
import Eventservices from '../../services/Eventservices';
export default {
  data() {
    return {
      user: this.$store.state.currentUser,
      BillMaster: this.$store.state.BillMaster,
      high: false,
      low: false,
      yesterdaysales: 0,
      todaysales:0,
      percentage:0,
      todaydate: this.getTodayDate(),
      yesterdayDate: this.getYesterdayDate(),
      yesterdayEarn: 0,
    };
  },
  computed: {
  
  },
  methods: {
    getYesterdayDate() {
      const today = new Date();
      const yesterday = new Date(today);
      yesterday.setDate(today.getDate() - 1);
      // Format the date as needed
      const formattedDate = yesterday.toISOString().split("T")[0];
      return formattedDate;
    },
    getTodayDate() {
      let currentDate = new Date().toJSON().slice(0, 10);
      console.log(currentDate);
      return currentDate;
    },
  },
  mounted() {
    // console.log("this.$store.state.currentlogin: ",this.$store.state.currentlogin)
    Eventservices.salesmandashboard(this.$store.state.currentlogin)
    .then((response) => {
         if (response.data.status == "S") {
                // console.log("response=",response);
                this.todaysales=response.data.todaySales
                this.yesterdaysales=response.data.yesterdaysales
                // console.log("todaysales:",this.todaysales)
                // console.log("yesterdaysales:",this.yesterdaysales)

                let value = (this.todaysales * 100) / this.yesterdaysales;
    //no earnings in today and yesterday
    if (this.todaysales == 0 && this.yesterdaysales ==0) {
      value = 0;
    }
    //comparing  whether the earning is greater or lesser than yesterday earnings
    if (this.todaysales > this.yesterdaysales) {
      this.high = true;
      this.percentage = Math.round(value - 100);
    } else if (this.todaysales == this.yesterdaysales) {
      this.high = true;
      this.percentage = 0;
    } else {
      this.low = true;
      this.percentage = Math.round(100 - value);
    }
         }else{
                console.log(response.data.errMsg);
              }
            })
            .catch((error) => {
              console.log(error);
            });
  },
};
</script>
