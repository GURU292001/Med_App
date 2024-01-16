<template>
  <div>
      <!-- <v-card elevation="4"> -->
     <!-- <v-row  class="  my-14  pt-7 px-5 "> -->
       <!-- <v-col> -->
        <!-- <v-card elevation="4" > -->
          <apexChart type="donut" :options="chartOptions03" :series="series03"></apexChart>
         <!-- </v-card> -->
       <!-- </v-col> -->
     <!-- </v-row> -->
   <!-- </v-card> -->
  </div>
</template>
<script>
import VueApexCharts from "vue-apexcharts";
export default{

   components:{
       apexChart:VueApexCharts
   },
   props:[
       'series03',
       'chartOptions03']
}
</script>