import { ref, Ref } from "vue";
import { defineStore } from "pinia";
import HeadType from "@/types/head";

export const useHeadStore = defineStore("head", () => {
    const head: Ref<HeadType> = ref({
        name: "",
        role: "",
        email: "",
        project: "",
        container: "",
        module: "",
    });

    return { head };
});
