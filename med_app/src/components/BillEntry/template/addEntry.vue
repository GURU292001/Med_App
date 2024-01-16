
   <template>
    <v-container outlined class=" d-flex justify-space-between ">
     <v-expansion-panels>
      <v-expansion-panel>
        <v-expansion-panel-header>
         Add Item
        </v-expansion-panel-header>
        <v-expansion-panel-content >
       <v-form>
          <div class="d-flex justify-space-around">
          <v-card flat width="300" >
        <v-row >
          <v-col>
            <v-select
              v-model="selectedItem"
              :items="stock"
              label="Medicine Name"
              item-text="medicine_Name"
            
            ></v-select>
          </v-col>
        </v-row>
      </v-card>
  
      <!-- <v-card width="300" class="justify-center px-5 pt-2"> -->
        <!-- <v-row> -->
         <v-card flat width="300">
          <v-col>
            <v-text-field
            v-model.number="quantity"
           label="Quantity"
            outlined
            type="number"
          ></v-text-field>
          </v-col>
         </v-card>
        <!-- </v-row> -->
      <!-- </v-card> -->
      <div   elevation="0" class="mt-4" >
        <addbtn btname="+ Add" @buttonEvent="add"/>
      </div>
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
         </div>
       </v-form>
        </v-expansion-panel-content>
      </v-expansion-panel>
     </v-expansion-panels>
    </v-container>
  </template>
  
  <script>
import addbtn from "../../Button/button.vue"
import EventServices from "../../../services/Eventservices"
  export default {
    // name:"addcompt",
    components: {addbtn},
    data() {
      return {
        // brand:"",
        snackbar: false,
      text: `Qunatity not be a decimal`,
        stock:[],
        selectedItem: "",
        // MedicineMaster:this.$store.state.MedicineMaster,
        quantity:'',
        
        // billstock:[],
      };
    },
  
   
   methods:{
        // brandfun(){
        //     console.log("InsideBrandfun");
        //     console.log(this.selectedItem);
        //     let findstock=this.stock.find(element=>element.value==this.selectedItem)
        //     console.log(findstock);
        //     this.brand = findstock.brand;
        // },
        add(){
         if(this.quantity%1==0)
         {
          this.$emit('addbutton',this.selectedItem,this.quantity)
          this.selectedItem=''
          this.quantity=0
         }
         else{
            console.log("else block");
            this.snackbar=true
         }
        }
   },
   mounted(){
    EventServices.showdorpdown()
      .then((response) => {
        if (response.data.status == "S") {
          console.log("response",response)
          for(let i of response.data.med_mst_tbl_Array)
          {
           var obj={
            medicine_Name:i.medicine_Name,
            brand:i.brand
           }
           this.stock.push(obj)
          }
        } else {
          this.result = response.data.errMsg;
        }
      })
      .catch((error) => {
        console.log(error);
      });
   }



  };
  </script>
    