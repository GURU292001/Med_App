<template>
  <v-container class="mt-16 ">
    <v-img max-height="650"  src="../assets/loginbg.jpg">
      <v-row justify="space-around">
        <v-card tile width="350" class="mt-16 px-4 pt-5 rounded-xxl rounded-b-0">
          <v-img  src="../assets/medlogo.png"></v-img>
          <v-card-title
            class=" green--text  font-weight-black text-h6 d-flex justify-center"
          >
            LOGIN
          </v-card-title>
        </v-card>
      </v-row>
    <!-- </v-container>
    <v-container> -->
      <v-row justify="space-around">
        <v-card elevation="9" tile width="350" class="py-10 mb-16 rounded-xxl rounded-t-0" >
          <v-form  class="px-8">
            <v-text-field label="E-Mail" v-model="user.user_id" 
            @keydown.space.prevent
            :rules="emailRules">
            </v-text-field>
            <v-text-field
              v-model="user.password"
              :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[rules.required, rules.min]"
              :type="show1 ? 'text' : 'password'"
              label="Password"
              @keydown.space.prevent
              v-on:keyup.enter="login"
              counter
              @click:append="show1 = !show1"
            ></v-text-field>

            <butn
              elevation="0"
              class="d-flex justify-center mt-10"
              btname="Login"
              @buttonEvent="login"
              @click="overlay = !overlay"
            />
          </v-form>
         
        </v-card>
        <v-snackbar
      v-model="snackbar"
    >
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn
          color="pink"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
    <!-- ---------------------------------- -->
    <v-overlay :value="overlay">
      <v-progress-circular
        indeterminate
        size="50"
      ></v-progress-circular>
    </v-overlay>
   
  <!-- </div> -->
    <!-- ---------------------------------- -->
      </v-row>
      <!-- {{ date }} -->
    </v-img>
  </v-container>
</template>
<script>
import butn from "../components/Button/button.vue";
import EventServices from "../services/Eventservices";

export default {
  name: "btname",
  components: { 
    butn,
  },
  data() {
    return {
      user:{
        user_id: "",
      password: "",
      },
      // date:'',
      snackbar: false,
      text: '',
      show1: false,
      overlay:false,
      userlogin: this.$store.state.Login,
      role: this.$store.state.currentUserRole,
      username: this.$store.state.currentUser,
      rules: {
        required: (value) => !!value || "Required.",
        min: (v) => v.length >= 4 || "Min 8 characters",
        emailMatch: () => `The email and password you entered don't match`,
      },
      emailRules: [
        (v) => !!v || "E-mail is required",
        (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
      ],
    };
  },
  methods: {
    login() {
      //----------------------------logincheck--------------------------------------
      EventServices.logincheck(this.user)
      .then((response) => {
        console.log(response);
        if (response.data.status == "S") {
          this.$store.state.currentlogin=response.data.login_id
          this.$store.state.currentUserRole=response.data.role
          console.log("1st reseponse recieved ")
          //-------------------------login data-----------
          EventServices.logindata(response.data.login_id)
          .then((response) => {
            console.log(response.data.status);
            if (response.data.status == "S") {
              this.overlay = true;
              setTimeout(() => {
                this.$store.state.currentUser=this.user.user_id
                this.$router.push("/dashboard");
                this.overlay = false;
              }, 1500);
          
          console.log("logindata -reseponse recieved ")
         } else {
          console.log(response.data.errMsg); 
          this.snackbar=true;
          this.text=response.date.errMsg
          this.email=''
          this.password=''
       }
      })
      .catch((error) => {
        console.log(error);
      });
          //--------------------------------
        } else {
          this.snackbar=true;
          this.text=response.data.errMsg
       this.email=''
       this.password=''
        }
      })
      .catch((error) => {
        console.log(error);
         
      });
      //------------------------------------logincheck----------------------------
      // alert("logged in")
      // let check = false;
      // let role = "";
      // this.userlogin.forEach((element) => {
      //   if (element.email == this.email && element.password == this.password) {
      //     role = element.role;
      //     check = true;
      //   }
      // });
      // if (check) {
      //   this.overlay = true;
      
        // this.$router.push("/dashboard");
        // let obj = { user: this.email, type: "login", date: this.date};
        
        
  //     } else {
  //      this.snackbar=true;
  //      this.email=''
  //      this.password=''
  //     }
      //------------------login data-----------------------------
      //------------------login data-----------------------------

     
    },
  },
  computed: {
    date() {
      // const currentDate = new Date();
      // const year = currentDate.getFullYear();
      // const month = (currentDate.getMonth() + 1).toString().padStart(2, "0");
      // const day = currentDate.getDate().toString().padStart(2, "0");
      // const x = `${day}-${month}-${year}`;
      const formattedDate = new Date().toJSON().slice(0,10);
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
