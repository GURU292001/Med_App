<template>
    
    <v-container  class="d-flex align-content-center flex-wrap mt-16 " >
     <v-card outlined  class="px-16 pb-16 d-flex justify-center ">
           <addstock/>
     </v-card>
      

    </v-container>

</template>
<script>
import addstock from "../components/StockEntry/template/updateStock.vue"
export default{
    name:"comp",
    components:{
        addstock
    },
    mounted(){
        if(this.$store.state.currentUser=='')
        {
            this.$router.push('/')
        }
    },
}
</script>