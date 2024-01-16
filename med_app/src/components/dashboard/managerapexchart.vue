<template>
        <div>
      <!-- chart-1 today sales trend -->  
      <dailytrends :series01="series01" :chartOptions01="chartOptions01"/>
      <!-- chart-2 today sales trend -->
      <mntlytrend :series02="series02" :chartOptions02="chartOptions02" />
      <!-- chart-3 piechart trend -->
     
      
      <v-row class="  my-14  pt-7 px-5 ">
        <v-col class="col-6">
         <v-card elevation="4" >
          <piechart :series03="series03" :chartOptions03="chartOptions03"/>
         </v-card>
        </v-col>
      <!-- chart-4 piechart trend -->
        <v-col class="col-6">
          <v-card elevation="4">
            <barchart :series04="series04" :chartOptions04="chartOptions04"/>
         </v-card>
        </v-col>
      </v-row>
    </div>

</template>
<script>
import EventServices from "../../services/Eventservices";
// import VueApexCharts from "vue-apexcharts";
import dailytrends from "../apexchart/dailysales.vue"
import mntlytrend from "../apexchart/monthlytreand.vue"
import piechart from "../apexchart/piechart.vue"
import barchart from "../apexchart/barchart.vue"
    export default{
        data(){
            return {
                weekdayxaxis:[],
      weekdayyaxis:[],
      monthxaxis:[],
      monthyaxis:[],
     series01: [
        {
          name: "Annnual",
          //y-axis
          data: [],
        },
      ],
      chartOptions01: {
        chart: {
          height: 350,
          type: "line",
          zoom: {
            enabled: false,
          },
        },
        dataLabels: {
          enabled: true,
        },
        stroke: {
          curve: "straight",
        },
        title: {
          text: "Daily sales trend",
          align: "left",
        },
        grid: {
          row: {
            colors: ["#f3f3f3", "transparent"], // takes an array which will be repeated on columns
            opacity: 0.5,
          },
        },
        xaxis: {
          //x-axis
          categories: ['sun','mon','tue','wed','thur','fri','satur'],
        },
      },
      series02: [
        {
          name: "Annnual",
          //y-axis
          data: [],
        },
      ],
      chartOptions02: {
        chart: {
          height: 350,
          type: "line",
          zoom: {
            enabled: false,
          },
        },
        dataLabels: {
          enabled: true,
        },
        stroke: {
          curve: "straight",
        },
        title: {
          text: "Monthly sales trend",
          align: "left",
        },
        grid: {
          row: {
            colors: ["#f3f3f3", "transparent"], // takes an array which will be repeated on columns
            opacity: 0.5,
          },
        },
        xaxis: {
          //x-axis
          categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
        },
      },
      series03: [],
      chartOptions03: {
        chart: {
          type: "donut",
        },
        labels: [],
        title: {
          text: "Today Biller's performance",
          align: "left",
        },
        plotOptions: {
          pie: {
            donut: {
              labels: {
                show: true,
                total: {
                  showAlways: true,
                  show: true,
                },
              },
            },
          },
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: 100,
              },
              legend: {
                position: "bottom",
              },
            },
          },
        ],
      },
       series04: [
          {
            name: "Biller",
            data: [],
          },
        ],
    chartOptions04: {
          chart: {
            type: "bar",
            height: 400,
          },
          title: {
                text: "Current Month Biller's performance",
                align: 'left'
              },
          plotOptions: {
            bar: {
              horizontal: false,
              columnWidth: "55%",
              endingShape: "rounded",
            },
          },
          dataLabels: {
            enabled: false,
          },
          stroke: {
            show: true,
            width: 5,
            colors: ["transparent"],
          },
          xaxis: {
            categories: [],
          },
          yaxis: {
            title: {
              text: "Rs. (thousands)",
            },
          },
          fill: {
            opacity: 1,
          },
          tooltip: {
            y: {
              formatter: function (val) {
                return "Rs. " + val + " thousands";
              },
            },
          },
        },
            }
        },
        components:{
            dailytrends,
            mntlytrend,
            piechart,
            barchart
        },

        mounted() {
    // //today sales
    // EventServices.mgrdashboard()
    //   .then((response) => {
    //     if (response.data.status == "S") {
    //       console.log("response", response);
    //       this.todaySales = response.data.todaySales;
    //       this.curInvValue = Math.round(response.data.currentInventryValue);
    //     } else {
    //       console.log(response.data.errMsg);
    //     }
    //   })
    //   .catch((error) => {
    //     console.log(error);
    //   });

    //apex chart
    EventServices.getchartvalues()
      .then((response) => {
        if (response.data.status == "S") {
          console.log(response.data)
          let z=0
          for(let i=response.data.weeksales.length-1; i>=0; i--){
            this.chartOptions01.xaxis.categories[z++] =response.data.weeksales[i].weekday;
            this.weekdayyaxis.push(response.data.weeksales[i].total)
          }
          this.series01 = [
            {
              name: "Annual",
              data: this.weekdayyaxis,
            },
          ];
          //Monthly sales trends mntySales
          let j=0
          for(let i=response.data.mntySales.length-1; i>=0; i--){
            this.chartOptions02.xaxis.categories[j++] =response.data.mntySales[i].monthname
            this.monthyaxis.push(response.data.mntySales[i].total)
          }
          this.series02 = [
            {
              name: "Annual",
              data: this.monthyaxis,
            },
          ];
          //pie chart
         if(response.data.tdySalmanperfm!=null){
          for(let i of response.data.tdySalmanperfm){
            this.chartOptions03.labels.push(i.user_id)
             this.series03.push(i.total)
          }
         }else{
          this.chartOptions03={labels:["Not Yet","Not Yet","Not Yet"]}
          this.series03=[0,0,0]
         }
          //current month salesman performance
          let user=[]
          let monthlysales=[]
          console.log("response.data.mntySalmanperfm",response.data.mntySalmanperfm);
          if (response.data.mntySalmanperfm!=null){
            for(let i of response.data.mntySalmanperfm){
           user.push(i.user_id)
           monthlysales.push(i.total)
          }
          this.series04 = [{
        name: "Billers",
        data: monthlysales
      }];
      this.chartOptions04={
          xaxis:{
              categories:user
          }
          }
      }else{
        this.series04 = [{
        name: "Billers",
        data: [0,0,0]
      }];
      this.chartOptions04={
          xaxis:{
              categories:["Not Yet","Not Yet","Not Yet"]
          }
      }
        }} else {
          console.log(response.data.errMsg);
        }
      })
      .catch((error) => {
        console.log(error);
      });
  },
    }

</script>