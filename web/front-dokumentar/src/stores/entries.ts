import { ref, Ref } from "vue";
import { defineStore } from "pinia";
import EntryType from "@/types/entry";

const initial: EntryType[] = [
	{
		id: 1,
		entity: "",
		connection: "",
		method: "",
		endpoint: "",
		json: "",
		comment: "",
	},
];

export const useEntriesStore = defineStore("entries", () => {
  const count = ref(1);
  
  const entries: Ref<EntryType[]> = ref(initial);

  function increment() {
    count.value++;
  }
  function reset() {
    count.value = 1;
  }

  function add() {
    increment();
    entries.value.push({
      id: count.value,
      entity: "",
      connection: "",
      method: "",
      endpoint: "",
      json: "",
      comment: "",
    });
		console.log(entries.value)
  }
  function clear() {
    reset();
    entries.value =[{
			id: 1,
			entity: "",
			connection: "",
			method: "",
			endpoint: "",
			json: "",
			comment: "",
		}]
		console.log(entries.value)
  }

  return { entries, add, clear };
});
