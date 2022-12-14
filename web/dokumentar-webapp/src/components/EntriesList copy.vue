<template>
    <div :key="count">
        <v-list-item v-for="entry in entries" :key="entry?.id">
            <EntryFormVue :data="entry" />
        </v-list-item>
    </div>
    <v-divider></v-divider>
    <v-container class="container">
        <v-btn
            icon
            color="#4B5468"
            variant="outlined"
            size="large"
            @click="addEntry"
        >
            <v-icon>mdi-plus</v-icon>
        </v-btn>
    </v-container>
</template>

<script lang="ts">
import type EntryType from "@/types/entry";
import EntryFormVue from "./EntryForm.vue";
import { defineComponent } from "vue";

const entries: EntryType[] = [
    {
        id: 1,
        entity: "foo",
        connection: "",
        method: "",
        endpoint: "",
        json: "",
        comments: "",
    },
];
let count = 1;
const addEntry = () => {
    entries.push({
        id: 2,
        entity: "foo",
        connection: "",
        method: "",
        endpoint: "",
        json: "",
        comments: "",
    });
    count++;
};

export default defineComponent({
    name: "EntriesList",
    data: () => ({
        entries,
        count,
    }),
    components: {
        EntryFormVue,
    },
    methods: {
        addEntry,
    },
});
</script>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    width: 100%;
}
</style>
