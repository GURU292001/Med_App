<template>
    <div>
          <apexChart class="py-13 "
          type="bar"
          max-height="80%"
          :options="chartOptions04"
          :series="series04"
        ></apexChart>
         </div>
</template>
<script>
import VueApexCharts from "vue-apexcharts";
export default{

   components:{
       apexChart:VueApexCharts
   },
   props:[
       'series04',
       'chartOptions04']
}
</script>