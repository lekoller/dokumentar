/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from "./App.vue";

// Composables
import { createApp } from "vue";

// Plugins
import { VueQueryPlugin } from "vue-query";
import { registerPlugins } from "@/plugins";
import { createPinia } from "pinia";
import router from "./router";

import "./assets/main.css";

const app = createApp(App);

app.use(router);
app.use(VueQueryPlugin);
app.use(createPinia());

registerPlugins(app);

app.mount("#app");
