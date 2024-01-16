<template>
     <v-card elevation="4">
      <v-row  class="  my-14  pt-7 px-5 ">
        <v-col>
          <apexChart
            type="line"
            height="298"
            :options="chartOptions01"
            :series="series01"
          ></apexChart>
        </v-col>
      </v-row>
    </v-card>
</template>
<script>
import VueApexCharts from "vue-apexcharts";
export default{
 
    components:{
        apexChart:VueApexCharts
    },
    props:[
        'series01',
        'chartOptions01']
}
</script>