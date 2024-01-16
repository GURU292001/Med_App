<template>
  <div class="flex-wrap">
    <div class="text-h6 d-flex justify-end my-5 flex-wrap">
      To Add new stock
    </div>
    <div class="d-flex justify-end my-5 flex-wrap">
      <v-dialog v-model="dialog" width="400">
        <template v-slot:activator="{ on, attrs }">
          <v-btn elevation="0" class="primary" v-bind="attrs" v-on="on"
            >+ADD
          </v-btn>
        </template>

        <v-form>
          <v-card elevation="0">
            <v-card-title class="text-h5 primary flex-wrap">
              Add Stock
            </v-card-title>

            <v-card-text
              class="d-flex flex-column justify-center mt-10 flex-wrap"
            >
              <v-text-field
                v-model="new_medicine.Medicine_Name"
                label="Medicine Name"
                :rules="[noLeadingTrailingSpaces, noDots, noNegatives]"
                outlined
              ></v-text-field>
              <v-text-field
                v-model="new_medicine.brand"
                :rules="[noLeadingTrailingSpaces, noDots, noNegatives]"
                label="Brand"
                outlined
              ></v-text-field>
              <v-card flat class="d-flex justify-space-between my-3">
                <v-btn class="red white--text" @click="close">CLOSE</v-btn>
                <addbtn btname="+add" @buttonEvent="add" />
              </v-card>
            </v-card-text>
            <v-divider></v-divider>
          </v-card>
        </v-form>
        <v-snackbar v-model="snackbar" :timeout="timeout" :color="color">
          {{ text }}
        </v-snackbar>
      </v-dialog>
    </div>
  </div>
</template>
<script>
import addbtn from "../../Button/button.vue";
import EventServices from "../../../services/Eventservices";
export default {
  components: {
    addbtn,
  },
  data() {
    return {
      dialog: false,
      new_medicine: {
        Medicine_Name: "",
        brand: "",
        user_id: "",
      },
      // MedicineMaster: this.$store.state.MedicineMaster,
      track: null,
      snackbar: false,
      text: "",
      timeout: 2000,
      color: "",
    };
  },
  methods: {
    add() {
      this.new_medicine.user_id = this.$store.state.currentUser;
      if (
        this.new_medicine.Medicine_Name.length > 2 &&
        this.new_medicine.brand.length > 2 &&
        this.noLeadingTrailingSpaces(this.new_medicine.Medicine_Name) &&
        this.noDots(this.new_medicine.Medicine_Name) &&
        this.noNegatives(this.new_medicine.Medicine_Name)
      ) {
        //------------------------------------------------------------------------------
        EventServices.addstock(this.new_medicine)
          .then((response) => {
            this.$emit("add");
            console.log(response);
            if (response.data.status == "S") {
              console.log("response.data.status", response.data.status);
              this.color = "success";
              this.text = "Added Successfully";
              this.snackbar = true;
              this.new_medicine.Medicine_Name = "";
              this.new_medicine.brand = "";
              // this.dialog = false;
            } else {
              this.result = response.data.errMsg;
            }
          })
          .catch((error) => {
            console.log(error);
          });
        //-------------------------------------
      } else {
        this.color = "red";
        this.text = "space,dot,(-) not allowed";
        this.snackbar = true;
      }
    },
    close() {
      this.new_medicine.Medicine_Name = "";
      this.new_medicine.brand = "";
      this.dialog = false;
    },
    noLeadingTrailingSpaces(value) {
      return value.trim() === value ? true : false;
    },
    noDots(value) {
      return value.indexOf(".") === -1 ? true : false;
    },
    noNegatives(value) {
      return value.indexOf("-") === -1 ? true : false;
    },
  },
};
</script>
