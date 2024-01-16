<template>
  <div elevation="0">
    <v-btn depressed color="primary" @click="EmitEvent">
      {{ btname }}
    </v-btn>
  </div>
</template>
<script>
export default {
  name:"btname",
  props: {
    btname: String,
  },
  methods:{
    EmitEvent(){
      console.log("****************");
        this.$emit("buttonEvent")
    }
  }
};
</script>
