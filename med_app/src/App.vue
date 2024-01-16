<template>
  <v-app>
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
 <navbar style="postion:sticky" v-if="this.$store.state.currentUser!=''"/>
 <!-- <navbar/> -->
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import navbar from "../src/components/navigation/navigation.vue"
import Eventservices from "../src/services/Eventservices"
export default {
  components:{
    navbar,
  },
  name: "App",


  created() {
    window.addEventListener('beforeunload', this.logout);
  },
  beforeDestroy() {
    window.removeEventListener('beforeunload', this.logout);
  },
  methods:{
   logout(){
    Eventservices.logoutdata(this.$store.state.currentlogin)
  .then((response) => {
    if (response.data.status == "S") {
      console.log(response.data)
      this.$store.state.currentlogin=""
      this.$store.state.currentUser=""
      this.$store.state.currentUserRole=""
          this.$router.push("/");
        } else {
           console.log(response.data.errMsg)
        }
      })
      .catch((error) => {
        console.log(error);
      });
   }
  }
};
</script>
