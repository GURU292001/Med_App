<template>
    <v-card outlined color="primary" class="mt-6">
      <butn class=" d-flex justify-end white px-5 py-5" btname="Download" @buttonEvent="download"/>
      <v-data-table
        :headers="headers"
        :items="details"
        item-key="billNo"
        class="elevation-1"
        :search="search"
        hide-default-footer
      >
     
        <template v-slot:append>
          <tr>
            <td></td>
            <td>
          
            </td>
            <td colspan="4"></td>
          </tr>
        </template>
      </v-data-table>

    </v-card>
  </template>
  <script>
import butn from "../../Button/button.vue"

  export default {
     name:"bttn",
  components:{
    butn,
  },
    data () {
      return {
        search: '',
        Brand: '',
      
      }
    },
  props:{
    headers:Array,
    details:Array,
  },
    methods: {
     
      download(){
        this.$emit('downloadevent')
      }
    },
  }
</script>