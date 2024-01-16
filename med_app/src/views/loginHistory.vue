<template>
  <v-card outlined color="primary" class="mx-16 mt-10">
    <butn
      class="d-flex justify-end white px-5 py-5"
      btname="Download"
      @buttonEvent="download"
    />
    <v-data-table
      :headers="headers"
      :items="loginhistory"
      class="elevation-1"
      :search="search"
      :custom-filter="filterOnlyCapsText"
    >
      <template v-slot:top>
        <v-text-field
          v-model="search"
          label="Search"
          class="mx-4"
        ></v-text-field>
      </template>

      <template v-slot:append>
        <tr>
          <td></td>
          <td>
            <v-text-field
              v-model="Brand"
              type="number"
              label="Less than"
            ></v-text-field>
          </td>
          <td colspan="4"></td>
        </tr>
      </template>
      <!-- <template v-slot:item.type="{ item }"> -->
        <!-- <v-chip :color="getColor(item.type)" dark> -->
          <!-- {{ item.type }} -->
        <!-- </v-chip> -->
      <!-- </template> -->
    </v-data-table>
    <!-- <butn
      class="d-flex justify-end white px-5 py-5"
      btname="Download"
      @buttonEvent="download"
    /> -->
    <v-snackbar v-model="snackbar" :timeout="timeout" color="green">
      {{ text }}
    </v-snackbar>
  </v-card>
</template>
<script>
import butn from "../components/Button/button.vue";
import Papa from "papaparse";
import Eventservices from "../services/Eventservices";

export default {
  name: "bttn",
  components: {
    butn,
  },
  data() {
    return {
      search: "",
      Brand: "",
      loginhistory:[] ,
      snackbar: false,
      text: "Start Downloaded..",
      timeout: 1500,
    };
  },
  computed: {
    headers() {
      return [
        {
          text: "UserId(E-mail)",
          align: "start",
          sortable: false,
          value: "user_id",
        },
        { text: "Login Date", value: "login_date"},
        { text: "Login Time", value: "login_Time" },
        { text: "Logout Date", value: "logout_date" },
        { text: "Logout Time", value: "logout_time" },
      ];
    },
  },
  methods: {
    getColor(calories) {
      if (calories == "login") {
        return "green";
      } else if (calories == "logout") {
        return "red";
      } else {
        return "green";
      }
    },
    filterOnlyCapsText(value, search) {
      return (
        value != null &&
        search != null &&
        typeof value === "string" &&
        value.toString().toLocaleUpperCase().indexOf(search.toUpperCase()) !==
          -1
      );
    },
    download() {
      // Your CSV data (replace this with your actual data)
      const csvData = this.login;

      // Convert the data array to a CSV string
      const csvString = Papa.unparse(csvData);

      // Create a Blob from the CSV string
      const blob = new Blob([csvString], { type: "text/csv" });

      // Create a temporary link element
      const link = document.createElement("a");
      link.href = window.URL.createObjectURL(blob);

      // Set the filename for the download
      link.download = "loginHistory.csv";

      // Append the link to the document
      document.body.appendChild(link);

      // Trigger the click event to initiate the download
      link.click();

      // Remove the link from the document
      document.body.removeChild(link);
      this.snackbar = true;
    },
  },
  mounted() {
    if (this.$store.state.currentUser == "") {
      this.$router.push("/");
    }
    Eventservices.loginHistory()
      .then((response) => {
        if (response.data.status == "S") {
          for (let i of response.data.login_Hsty_Array) {
            for (let ii of i) {
              var obj = {
                user_id: ii.user_id,
                login_date: ii.login_date,
                login_Time: ii.login_time,
                logout_date: ii.logout_date,
                logout_time: ii.logout_time,
              };
              this.loginhistory.push(obj);
            }
          }
        } else {
          console.log(response.data.errMsg);
        }
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>
