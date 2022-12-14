<template>
    <v-list-item v-for="entry in store.entries" :key="entry?.id">
        <EntryFormVue :data="entry" />
    </v-list-item>
    <v-divider></v-divider>
    <v-container class="container">
        <v-btn variant="outlined" color="#544D6B" @click="store.clear">
            <span>clear</span>
            <v-icon class="icon">mdi-delete</v-icon>
        </v-btn>
        <v-btn
            icon
            color="#4B5468"
            variant="outlined"
            size="large"
            @click="store.add"
        >
            <v-icon>mdi-plus</v-icon>
        </v-btn>
        <v-btn variant="outlined" color="#405D5D"> 
            <span>Write Documentation</span>
            <v-icon class="icon">mdi-pencil</v-icon>
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
    justify-content: space-between;
    width: 100%;
    padding-left: 30px;
    padding-right: 30px;
}
.icon {
    margin-left: 12px;
}
</style>
