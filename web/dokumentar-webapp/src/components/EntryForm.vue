<template>
    <v-form v-model="valid">
        <v-container class="container">
            <v-row>
                <v-col cols="9" md="3">
                    <v-text-field
                        v-model="entity"
                        :rules="nameRules"
                        :counter="20"
                        label="Entity"
                        required
                    ></v-text-field>
                </v-col>

                <v-col cols="9" md="3">
                    <v-text-field
                        v-model="connection"
                        :rules="nameRules"
                        :counter="20"
                        label="Connection Type"
                        required
                    ></v-text-field>
                </v-col>

                <v-col cols="9" md="3">
                    <v-text-field
                        v-model="method"
                        :rules="nameRules"
                        :counter="20"
                        label="HTTP Method"
                        required
                    ></v-text-field>
                </v-col>

                <v-col cols="9" md="3">
                    <v-text-field
                        v-model="endpoint"
                        :rules="nameRules"
                        :counter="20"
                        label="/path/to"
                        required
                    ></v-text-field>
                </v-col>
            </v-row>

            <v-row>
                <v-col cols="12" md="6">
                    <v-textarea
                        v-model="json"
                        label="JSON"
                        required
                    ></v-textarea>
                </v-col>

                <v-col cols="12" md="6">
                    <v-textarea
                        v-model="comment"
                        :rules="commentRules"
                        :counter="500"
                        label="Commentary"
                    ></v-textarea>
                </v-col>
            </v-row>
        </v-container>
    </v-form>
</template>

<script lang="ts">
import { defineComponent, type PropType } from "vue";
import type EntryType from "../types/entry";

export default defineComponent({
    name: "EntryForm",
    props: {
        data: {
            type: Object as PropType<EntryType>,
            required: true,
        },
    },
    data: () => ({
        valid: false,
        entity: "",
        connection: "",
        method: "",
        endpoint: "",
        module: "",
        json: "",
        comment: "",
        nameRules: [
            (v: string) => !!v || "Name is required",
            (v: string) =>
                v.length <= 20 || "Name must be less than 20 characters",
        ],
        commentRules: [
            (v: string) =>
                v.length <= 20 || "Name must be less than 500 characters",
        ],
    }),
});
</script>

<style scoped>
.container {
    margin: 32px 0 12px 0;
}
</style>
