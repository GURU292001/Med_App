<template>
  <div>
    <div>
      <div>
        <addbutton @add="dropdown" class="flex-wrap" />
      </div>
      <v-container>
        <v-card flat class="d-flex justify-space-between px-5 py-5 flex-wrap">
          <v-card flat width="200" class="justify-center px-5 pt-2 flex-wrap">
            <v-row>
              <v-col>
                <v-select
                  v-model="reqbody.medicine_Name"
                  :items="MedicineMaster"
                  label="Medicine Name"
                  item-text="medicine_Name"
                  outlined
                ></v-select>
              </v-col>
            </v-row>
          </v-card>

          <v-card flat width="200" class="justify-center px-5 pt-2 flex-wrap">
            <v-row>
              <v-col>
                <v-card-text
                  outlined
                  lighten-4
                  class="black--text text-h6"
                  v-model="brand"
                >
                  {{ brand }}
                </v-card-text>
              </v-col>
            </v-row>
          </v-card>

          <v-card flat>
            <v-text-field
              v-model.number="reqbody.quantity"
              label="quantity"
              outlined
              type="number"
            ></v-text-field>
          </v-card>
          <v-card flat class="d-flex justify-center">
            <v-text-field
              v-model.number="reqbody.unit_Price"
              label="Unit Price"
              outlined
              type="number"
            ></v-text-field>
          </v-card>
          <v-snackbar v-model="snackbar" :color="color">
            {{ text }}

            <template v-slot:action="{ attrs }">
              <v-btn color="pink" text v-bind="attrs" @click="snackbar = false">
                Close
              </v-btn>
            </template>
          </v-snackbar>
        </v-card>
      </v-container>
      <v-card flat class="d-flex justify-center flex-wrap">
        <v-btn width="1100" @click="update" class="primary flex-wrap"
          >Update</v-btn
        >
      </v-card>
    </div>
  </div>
</template>
<script>
// import addbtn from "../../Button/button.vue";
import addbutton from "../template/newaddstock.vue";
import EventServices from "../../../services/Eventservices";
export default {
  name: "comp",
  components: {
    // addbtn,
    addbutton,
  },
  data() {
    return {
      dialog: false,
      stock: this.$store.state.Stock,
      MedicineMaster: [],
      reqbody: {
        user_id: "",
        medicine_Name: "",
        unit_Price: 0,
        quantity: "",
        brand: "",
      },
      medicine_Name: "",
      snackbar: false,
      text: "",
      color: "",
    };
  },
  methods: {
    update() {
      //  console.log( this.medicine_Name);
      //  console.log( this.brand);
      //  console.log( this.quantity);
      //  console.log( this.unit_Price);
      if (this.reqbody.quantity > 0 && this.reqbody.unit_Price > 0) {
        if (this.reqbody.quantity % 1 == 0) {
          this.reqbody.brand = this.brand;
          this.reqbody.user_id = this.$store.state.currentUser;
          console.log("before event", this.reqbody);
          EventServices.updatestock(this.reqbody)
            .then((response) => {
              if (response.data.status == "S") {
                console.log("response", response);
                // this.$store.commit('stockUpdate',{medicineName:this.medicine_Name,brand:this.brand,quantity:this.quantity,unit_Price:this.unit_Price})
                this.snackbar = true;
                this.reqbody.medicine_Name = "";
                this.reqbody.quantity = 0;
                this.reqbody.unit_Price = 0;
                this.text = "Stock Updated";
                this.color = "success";
              } else {
                this.reqbody.medicine_Name = "";
                this.reqbody.quantity = 0;
                this.reqbody.unit_Price = 0;
                this.color = "failure";
                this.text = response.data.errMsg;
                this.snackbar = true;
              }
            })
            .catch((error) => {
              console.log(error);
            });
        } else {
          this.reqbody.medicine_Name = "";
          this.reqbody.quantity = 0;
          this.reqbody.unit_Price = 0;
          this.color = "failure";
          this.text = `Quantity can't be decimal`;
          this.snackbar = true;
          // this.text='Stock Updated'
        }
      } else {
        this.reqbody.medicine_Name = "";
        this.reqbody.quantity = 0;
        this.reqbody.unit_Price = 0;
        this.text = `Enter the Valid data!`;
        this.snackbar = true;
      }
    },
    getDropdownData() {
      EventServices.showdorpdown()
        .then((response) => {
          if (response.data.status == "S") {
            console.log(
              "esponse.data.med_mst_tbl_Array",
              response.data.med_mst_tbl_Arrayponse
            );
            for (let i of response.data.med_mst_tbl_Array) {
              var obj = {
                medicine_Name: i.medicine_Name,
                brand: i.brand,
              };
              this.MedicineMaster.push(obj);
            }
          } else {
            this.color = "failure";
            this.text = response.data.errMsg;
            this.snackbar = true;
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    dropdown() {
      console.log("getdd");
      this.getDropdownData();
    },
  },
  mounted() {
    // EventServices.showdorpdown()
    //   .then((response) => {
    //     if (response.data.status == "S") {
    //       console.log("esponse.data.med_mst_tbl_Array",response.data.med_mst_tbl_Arrayponse)
    //       for(let i of response.data.med_mst_tbl_Array)
    //       {
    //        var obj={
    //         medicine_Name:i.medicine_Name,
    //         brand:i.brand
    //        }
    //        this.MedicineMaster.push(obj)
    //       }
    //     } else {
    //   this.color='failure'
    //   this.text=response.data.errMsg
    //   this.snackbar=true
    //     }
    //   })
    //   .catch((error) => {
    //     console.log(error);
    //   });
    this.getDropdownData();
    // let neww=this.stock
    //   console.log("InsideBrandfun");
    //   console.log(this.medicine_Name);
    //   for (let i in neww) {
    //   for (let j in this.$store.state.MedicineMaster) {
    //     if (neww[i].medicineName === this.$store.state.MedicineMaster[j].medicineName) {
    //       neww[i]['brand'] = this.$store.state.MedicineMaster[j].brand;
    //     }
    //   }
    // }
    // this.stock = neww;
  },
  computed: {
    brand() {
      let exsist = this.MedicineMaster.find(
        (element) => element.medicine_Name == this.reqbody.medicine_Name
      );
      return exsist ? exsist.brand : "NoBrand";
    },
  },
};
</script>
