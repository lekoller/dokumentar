import { ref, Ref } from "vue";
import { defineStore } from "pinia";
import EntryType from "@/types/entry";

export const useEntriesStore = defineStore("entries", () => {
	const count = ref(1)
	const initial: EntryType[] = [
		{
			id: count.value,
			entity: "",
			connection: "",
			method: "",
			endpoint: "",
			json: "",
			comment: "",
		}
	]
    const entries: Ref<EntryType[]> = ref(initial);
	
	function increment() {
        count.value++;
    }
	function reset() {
		count.value = 1;
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
	function clear() {
		reset()
		entries.value = initial
	}
	
    return { entries, add, clear };
});
