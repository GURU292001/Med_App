<template>
    <div >
        <v-tab outlined class=" primary--text text-decoration-none px-6 py-4 " :to="navigator">{{ pagename }}</v-tab>
    </div>
</template>
<script>
export default {
    props:{
        navigator:String,
        pagename:String,
    }
}
</script>