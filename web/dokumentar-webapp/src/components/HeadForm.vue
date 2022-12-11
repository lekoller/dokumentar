<template>
    <v-container class="bg-surface-primary">
        <v-form v-model="valid">
            <v-container>
                <v-row>
                    <v-col cols="12" md="4">
                        <v-text-field
                            v-model="name"
                            :rules="nameRules"
                            :counter="20"
                            label="Your Name"
                            required
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="4">
                        <v-text-field
                            v-model="role"
                            :rules="optionalNameRules"
                            :counter="20"
                            label="Your Role"
                            required
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="4">
                        <v-text-field
                            v-model="email"
                            :rules="optionalNameRules"
                            label="E-mail"
                            required
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row>
                    <v-col cols="12" md="4">
                        <v-text-field
                            v-model="project"
                            :rules="nameRules"
                            :counter="20"
                            label="Project Name"
                            required
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="4">
                        <v-text-field
                            v-model="container"
                            :rules="nameRules"
                            :counter="20"
                            label="Container Name"
                            required
                        ></v-text-field>
                    </v-col>

                    <v-col cols="12" md="4">
                        <v-text-field
                            v-model="module"
                            :rules="nameRules"
                            :counter="20"
                            label="Module Name"
                            required
                        ></v-text-field>
                    </v-col>
                </v-row>
            </v-container>
        </v-form>
    </v-container>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    name: "Layout",
    data: () => ({
        valid: false,
        name: "",
        role: "",
        project: "",
        container: "",
        module: "",
        nameRules: [
            (v: string) => !!v || "Name is required",
            (v: string) =>
                v.length <= 20 || "Name must be less than 20 characters",
        ],
        optionalNameRules: [
            (v: string) =>
                v.length <= 20 || "Name must be less than 20 characters",
        ],
        email: "",
        emailRules: [
            (v: string) => !!v || "E-mail is required",
            (v: string) => /.+@.+/.test(v) || "E-mail must be valid",
        ],
    }),
});
</script>
