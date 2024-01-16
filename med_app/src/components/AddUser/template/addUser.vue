<template>
  <v-container>
    <v-expansion-panels>
      <v-expansion-panel>
        <v-expansion-panel-header v-model="dialog">
          Add User
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <div class="d-flex justify-space-around">
            <v-text-field
              class="mx-5"
              v-model="obj.user_id"
              label="UserId"
              :rules="UserIdRules"
              @keydown.space.prevent
              outlined
            ></v-text-field>
            <v-text-field
              class="mx-5"
              v-model="obj.password"
              label="Password"
              type="password"
              :rules="[passwordrules.required,passwordrules.min]"
              @keydown.space.prevent
              outlined
            ></v-text-field>
            <v-select
              class="mx-5"
              v-model="obj.role"
              :items="roles"
              label="Role"
              outlined
            ></v-select>

            <div class="my-3">
              <btn btname="+add" @buttonEvent="adduser" />
            </div>
          </div>
          <!-- </v-card> -->
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-snackbar v-model="snackbar" :timeout=2000 :color="color">
      {{ text }}
      <template v-slot:action="{ attrs }" >
        <v-btn color="white" text v-bind="attrs"  @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>
<script>
import btn from "../../Button/button.vue";
import Eventservices from "../../../services/Eventservices";
export default {
  name: "comp",
  data() {
    return {
      obj: {
        current_User_id: "",
        user_id: "",
        password: "",
        role: "",
      },
      roles: ["manager", "salesman", "inventry", "admin"],
      dialog: false,
      UserIdRules: [
        (v)=>v.length!=0 || "Enter the valid data",
      (v) => !!v || "E-mail is required",
        (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
      ],
      passwordrules:{
      required: (value) => !!value || "Required.",
        min: (v) => v.length >= 4 || "Min 8 characters",
        // emailMatch: () => `The email and password you entered don't match`,
      },
      snackbar: false,
      text: "",
      
      color: "",
    };
  },
  components: {
    btn,
  },
  methods: {
    adduser() {
      console.log("inside add user");
      if (this.userid ==""  && this.password == "" && this.role != "") {
        console.log("first if block");
          //----------------------adduser-------------------------------
          this.obj.current_User_id = this.$store.state.currentUser;
          Eventservices.adduser(this.obj)
            .then((response) => {
              if (response.data.status == "S") {
                this.dialog = false;
                this.obj.user_id = "";
                this.obj.password = "";
                this.obj.role = "";
                this.text = "Added Successfully ";
                this.snackbar = true;
                this.color = "green";
              } else if (response.data.status == "Exist") {
                this.obj.user_id = "";
                this.obj.password = "";
                this.obj.role = "";
                this.text = "User Already Exists";
                this.snackbar = true;
                this.color = "red";
              } else {
                console.log(response.data.errMsg);
              }
            })
            .catch((error) => {
              console.log(error);
            });
        }else {
        this.text = "Enter the valid data";
        this.snackbar = true;
        this.color = "red";
        this.obj.user_id = "";
                this.obj.password = "";
                this.obj.role = "";
      }
    },
  },
};
</script>
