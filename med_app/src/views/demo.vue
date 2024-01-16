<template>
    <v-data-table
      :headers="headers"
      :items="stock"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar  flat >
          <v-dialog
            v-model="dialog2"
            max-width="500px">
            <v-card>
              <v-card-text>
                <v-container>
                    <v-row class="mt-5 ml-1">
                        <h1>Edit Item</h1>
                    </v-row>
                  <v-row>
                    <v-col
                      cols="12"
                      sm="6"
                      md="4"  >
                      <v-text-field
                        v-model="editedItem.medicineName"
                        label="medicineName"
                      ></v-text-field>
                    </v-col>
                    <v-col
                      cols="12"
                      sm="6"
                      md="4"  >
                      <v-text-field
                        v-model="editedItem.brand"
                        label="brand"
                      ></v-text-field>
                    </v-col>
                    <v-col
                      cols="12"
                      sm="6"
                      md="4"
                    >
                      <v-text-field
                        v-model="editedItem.quantity"
                        label="quantity"
                      ></v-text-field>
                    </v-col>
                    <v-col
                      cols="12"
                      sm="6"
                      md="4"
                    >
                      <v-text-field
                        v-model="editedItem.unitPrice"
                        label="unitPrice"
                      ></v-text-field>
                    </v-col>
                    <v-col
                      cols="12"
                      sm="6"
                      md="4"
                    >
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>
  
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="blue darken-1"
                  text
                  @click="cancel"
                >
                  Cancel
                </v-btn>
                <v-btn
                  color="blue darken-1"
                  text
                  @click="savechanges"
                >
                  Save
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
          <v-dialog v-model="dialogDelete" max-width="500px">
            <v-card>
              <v-card-title class="text-h5">Are you sure you want to delete this item?</v-card-title>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="closeDelete">Cancel</v-btn>
                <v-btn color="blue darken-1" text @click="deleteItemConfirm">OK</v-btn>
                <v-spacer></v-spacer>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-icon
          small
          class="mr-2"
          @click="editItem(item)"
        >
        edit
        </v-icon>
        <v-icon
          small
          @click="deleteItem(item)"
        >
          delete
        </v-icon>
      </template>
      <!-- <template v-slot:no-data>
        <v-btn
          color="primary"
          @click="initialize"
        >
          Reset
        </v-btn>
      </template> -->
    </v-data-table>
  </template>
  <script>
  export default {
    data: () => ({
      dialog2: false,
      dialogDelete: false,
      headers: [
        {
          text: 'medicineName',
          align: 'start',
          sortable: false,
          value: 'medicineName',
        },
        { text: 'brand', value: 'brand' },
        { text: 'quantity', value: 'quantity' },
        { text: 'unitPrice', value: 'unitPrice' },
        { text: 'Actions', value: 'actions', sortable: false },
      ],
      stock:  [
        {medicineName:"Paracetamol",brand:"Calpol",
        quantity:10,unitPrice:10}, {medicineName:"Paracetamol",brand:"Calpol",
        quantity:10,unitPrice:10}, {medicineName:"Paracetamol",brand:"Calpol",
        quantity:4,unitPrice:10}, {medicineName:"Paracetamol",brand:"Calpol",
        quantity:10,unitPrice:10}, {medicineName:"Paracetamol",brand:"Calpol",
        quantity:10,unitPrice:10},
        ],
      editedIndex: -1,
      editedItem: {
        medicineName:"",brand:"",
        quantity:0,unitPrice:0
      },
      defaultItem: {
        medicineName:"",brand:"",
        quantity:0,unitPrice:0
      },
    }),

    computed: {
    //   formTitle () {
    //     return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    //   },
    },

    watch: {
      dialog2 (val) {
        val || this.cancel()
      },
      dialogDelete (val) {
        val || this.closeDelete()
      },
    },

    // created () {
    //   this.initialize()
    // },

    methods: {
    //   initialize () {
    //     this.stock =
    //   },

      editItem (item) {
        this.editedIndex = this.stock.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog2 = true
        //props
      },

      deleteItem (item) {
        this.editedIndex = this.stock.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialogDelete = true
        //emits
      },

      deleteItemConfirm () {
        this.stock.splice(this.editedIndex, 1)
        this.closeDelete()
      },

      cancel () {
        this.dialog2 = false
        this.$nextTick(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        })
      },

      closeDelete () {
        this.dialogDelete = false
        this.$nextTick(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        })
      },

      savechanges () {
        if (this.editedIndex > -1) {
          Object.assign(this.stock[this.editedIndex], this.editedItem)
        } else {
          this.stock.push(this.editedItem)
        }
        //emits
        this.cancel()
      },
    },
  }
</script>