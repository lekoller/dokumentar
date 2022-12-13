import { createApp } from "vue";
import { createPinia } from "pinia";
import { VueQueryPlugin } from "vue-query";

import App from "./App.vue";
import router from "./router";

// vuetify
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import "@mdi/font/css/materialdesignicons.css";

const vuetify = createVuetify({
    components,
    directives,
});

import "./assets/main.css";

const app = createApp(App);

app.use(vuetify);
app.use(createPinia());
app.use(router);
app.use(VueQueryPlugin);

app.mount("#app");