<template>
  <v-container >
   <v-row class="d-flex justify-space-around">
    <v-col
      cols="12"
      sm="6"
      md="4"
    >
      <v-menu
        v-model="menu1"
        :close-on-content-click="false"
        :nudge-right="40"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date1"
            label="From"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="date1"
          @input="menu1 = false"
        ></v-date-picker>
      </v-menu>
    </v-col>
   <!-- </v-row> -->
   <!-- <v-row> -->
    <v-col
      cols="12"
      sm="6"
      md="4"
    >
      <v-menu
        v-model="menu2"
        :close-on-content-click="false"
        :nudge-right="40"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date2"
            label="To"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="date2"
          @input="menu2 = false"
        ></v-date-picker>
      </v-menu>
    </v-col>
    <v-col  class="d-flex justify-space-around" align-self="center">
        <searchbtn btname="search" @buttonEvent="search"/>
    </v-col>
   </v-row>
 
 
  
  </v-container>
</template>
<script>
import searchbtn from "../../Button/button.vue"
export default {
    data(){
        return{
            date1: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
            date2: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
      menu: false,
      modal: false,
      menu1: false,
      menu2: false,
        }
    },
    components:{
        searchbtn,
    },
   methods:{
    search(){
        this.$emit('searchevent',this.date1,this.date2)
    }
   }
}
</script>