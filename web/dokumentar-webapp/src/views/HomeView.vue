<script setup lang="ts">
import { findAll } from "@/libs/api";
import { useQuery } from "vue-query";
import HeadFormVue from "../components/HeadForm.vue";

const query = useQuery("findAll", () => findAll());
</script>

<template>
    <main>
        <h1>D O K U M E N T A R</h1>
        <HeadFormVue />
    </main>
</template>
<TheWelcome />
