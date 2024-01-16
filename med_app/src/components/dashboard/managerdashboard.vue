<template>
  <v-container>
    <div class="d-flex mt-16 pt-16">
      <v-row >
        <v-col>
          <tdysales  title="today sales" :amount="todaySales" />
        </v-col>
        <v-col>
          <tdysales title="Current Inventry Value" :amount="curInvValue" />
        </v-col>
      </v-row>
    </div>
    <piecharts/>
   
  </v-container>
</template>
<script>
import tdysales from "../TodaySales/todaysalesTemplate.vue";
import EventServices from "../../services/Eventservices"
import piecharts from "../dashboard/managerapexchart.vue"
export default {
  data() {
    return {
      todaySales: 0,
      curInvValue: 0,
    };
  },
  components: {
    tdysales,
    piecharts,
  },
  name: "weeklychart",
  props: {
    managers: Boolean,
  },
  mounted(){
//  today sales
EventServices.mgrdashboard()
.then((response) => {
        if (response.data.status == "S") {
          console.log("response", response);
          this.todaySales = response.data.todaySales;
          this.curInvValue = Math.round(response.data.currentInventryValue);
        } else {
          console.log(response.data.errMsg);
        }
      })
      .catch((error) => {
        console.log(error);
      });

},
}
</script>
