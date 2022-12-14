<template>
    <v-list-item v-for="entry in store.entries" :key="entry?.id">
        <EntryFormVue :data="entry" />
    </v-list-item>
    <v-divider></v-divider>
    <v-container class="container">
        <v-btn
            icon
            color="#4B5468"
            variant="outlined"
            size="large"
            @click="store.add"
        >
            <v-icon>mdi-plus</v-icon>
        </v-btn>
    </v-container>
</template>

<script lang="ts">
import EntryFormVue from "./EntryForm.vue";
import { defineComponent } from "vue";
import { useEntriesStore } from "@/stores/entries";

export default defineComponent({
    name: "EntriesList",
    setup() {
        const store = useEntriesStore();
        console.log(store.entries);

        return {
            store,
        };
    },
    components: {
        EntryFormVue,
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
