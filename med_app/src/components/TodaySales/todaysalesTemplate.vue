<template>
  <v-container>
    <v-card class="primary" width="500" height="150">
      <v-card-title class="text-h5 white--text font-weight-black mb-5">
        {{ title }}
      </v-card-title>
      <v-card-subtitle class="white--text text-h3">
        <span class="green--text">₹</span>
        {{ amount }}
      </v-card-subtitle>
    </v-card>
  </v-container>
</template>
<script>
export default {
  props: ["amount", "title"],
};
</script>
