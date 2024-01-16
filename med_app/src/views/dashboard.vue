<template>
   <div>
    <admin v-if="role=='admin'|| role=='inventry'"/>
    <manager v-if="role=='manager'"/>
    <salesman v-if="role=='salesman'"/>
   </div>
</template>
<script>
// import links from "../components/navigation/links.vue"
import admin from "../components/dashboard/adminInventrydashboard.vue"
import manager from "../components/dashboard/managerdashboard.vue"
import salesman from "../components/dashboard/salesmanDashboard.vue"
export default{
    components:{
        // links,
        admin,
        manager,
        salesman,
    },
    data(){
        return{
            role:this.$store.state.currentUserRole,
        }
    },
    mounted(){
        if(this.$store.state.currentUser=='')
        {
            this.$router.push('/')
        }
    }
}
</script>