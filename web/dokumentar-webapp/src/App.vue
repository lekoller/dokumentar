<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useQueryProvider } from "vue-query";
import { ref } from "vue";

// import { useTheme } from "vuetify";

useQueryProvider();

const theme = ref("dark");

function onClick() {
    theme.value = theme.value === "light" ? "dark" : "light";
}
</script>

<template>
    <v-app :theme="theme">
        <v-app-bar>
            <v-spacer></v-spacer>

            <v-btn
                :prepend-icon="
                    theme === 'light'
                        ? 'mdi-weather-sunny'
                        : 'mdi-weather-night'
                "
                @click="onClick"
                >Toggle Theme</v-btn
            >
        </v-app-bar>

        <v-main>
            <v-container>
                <RouterView />
            </v-container>
        </v-main>
    </v-app>
</template>

<style scoped>
/* header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
} */
</style>
