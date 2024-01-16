<template>
    <v-container class="mt-16">
            <adduser/>
    </v-container>
</template>
<script>
import adduser from "../components/AddUser/template/addUser.vue"
export default{
    name:"comp",
    components:{
        adduser,
    },
    mounted(){
        if(this.$store.state.currentUser=='')
        {
            this.$router.push('/')
        }
    }
}
</script>