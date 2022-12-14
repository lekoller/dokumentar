<script setup lang="ts">
import HeadFormVue from "@/components/HeadForm.vue";
import EntriesListVue from "@/components/EntriesList.vue";
// import { findAll } from "@/libs/api";
// import { useQuery } from "vue-query";

// const query = useQuery("findAll", () => findAll());
</script>

<template>
    <main>
        <h1>D O K U M E N T A R</h1>
        <HeadFormVue />
        <v-divider></v-divider>
        <EntriesListVue />
    </main>
</template>
<TheWelcome />
