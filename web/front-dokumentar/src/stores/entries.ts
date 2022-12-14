import { ref, Ref } from "vue";
import { defineStore } from "pinia";
import EntryType from "@/types/entry";

export const useEntriesStore = defineStore("entries", () => {
	const count = ref(1)
    const entries: Ref<EntryType[]> = ref([
		{
			id: count.value,
			entity: "",
			connection: "",
			method: "",
			endpoint: "",
			json: "",
			comment: "",
		}
	]);
	function increment() {
        count.value++;
    }
	function add() {
		increment()
		entries.value.push({
			id: count.value,
			entity: "",
			connection: "",
			method: "",
			endpoint: "",
			json: "",
			comment: "",
		})
	}

    return { entries, add };
});
