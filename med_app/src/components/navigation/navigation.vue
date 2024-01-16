<template>
  <v-card class="mt-3 mx-0">
    <v-tabs class="my-3 mx-1 d-flex justify-start">
      <v-img
        class="d-flex mx-10"
        width="150"
        src="../../assets/medlogo.png"
      ></v-img>
      <v-tab disabled> {{ role }}</v-tab>
      <links
        v-if="role == 'admin' || role == 'inventry'"
        pagename="Dashboard"
        navigator="/dashboard"
      />
      <links
        v-if="role == 'manager'"
        pagename="Dashboard"
        navigator="/dashboard"
      />
      <links
        v-if="role == 'salesman'"
        pagename="Dashboard"
        navigator="/dashboard"
      />
      <links
        v-if="role == 'admin'"
        pagename="Login History"
        navigator="/loginhistory"
      />
      <links v-if="role == 'admin'" pagename="Add User" navigator="/adduser" />
      <links
        v-if="role == 'manager' || role == 'salesman' || role == 'inventry'"
        pagename="Stock View"
        navigator="/stockview"
      />
      <links
        v-if="role == 'inventry' || role == 'manager'"
        pagename="Stock Entry"
        navigator="/stockentry"
      />
      <links
        v-if="role == 'salesman'"
        pagename="Bill Entry"
        navigator="/billentry"
      />
      <links
        v-if="role == 'manager'"
        pagename="Sales Report"
        navigator="/salesreport"
      />
      <v-btn elevation="0"  class="red--text mt-1" @click="log" >logout</v-btn>
    </v-tabs>
    <!-- ------------------------------------------------------------- -->

    <v-dialog
      v-model="dialog"
      max-width="300"
     
    >
      <v-card>
        <v-card-title class="text-h5">
          Are you sure to logout?
        </v-card-title>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="black "
            text
            @click="dialog = false"
          >
           Cancel
          </v-btn>

          <v-btn
            color="red darken-1"
            text
            @click="logout"
          >
            Logout
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
     <!-- ---------------------------- -->


    <v-overlay :value="overlay">
      <v-progress-circular
        indeterminate
        size="60"
      ></v-progress-circular>
    </v-overlay>
  </v-card>
</template>
<script>
import links from "./links.vue";
import Eventservices from "../../services/Eventservices";
export default {
  components: {
    links,
  },
  data() {
    return {
      role: this.$store.state.currentUserRole,
      overlay:false,
      dialog:false,
    };
  },
  methods: {
    log(){
      this.dialog=true
    },
    logout() {
      //for log out dialogu
      this.dialog=false
  // Set the overlay to true
  this.overlay = true;
  // Use setTimeout to simulate a delay before completing the logout
  setTimeout(() => {
  //   // Commit the logout update to the store
  //   this.$store.commit("loginupdate", obj);
  //   // After completing the logout actions, set the overlay back to false
    this.overlay = false;
  // // Clear user information in the store
  // this.$store.state.currentUser = "";
  //   this.$store.state.currentUserRole = "";
  //   // Navigate to the home page ("/")
  
  Eventservices.logoutdata(this.$store.state.currentlogin)
  .then((response) => {
    if (response.data.status == "S"){
      console.log(response.data)
      this.$store.state.currentlogin=""
      this.$store.state.currentUser=""
      this.$store.state.currentUserRole=""
          this.$router.push("/");
        }else{
           console.log(response.data.errMsg)
        }
      })
      .catch((error) => {
        console.log(error);
      });
    }, 2000); // Adjust the delay as needed
    
},

  },
  computed: {
    date() {
      const currentDate = new Date();
      const year = currentDate.getFullYear();
      const month = (currentDate.getMonth() + 1).toString().padStart(2, "0");
      const day = currentDate.getDate().toString().padStart(2, "0");
      const formattedDate = `${year}-${month}-${day}`;
      console.log(formattedDate);
      return formattedDate;
    },
  },
  watch: {
      overlay (val) {
        val && setTimeout(() => {
          this.overlay = false
        }, 3000)
      },
    },
};
</script>
