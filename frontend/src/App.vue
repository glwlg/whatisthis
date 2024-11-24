<script setup>
import {RouterView} from 'vue-router';
import StatusBar from "./components/StatusBar.vue";
import {EventsOn} from "../wailsjs/runtime/runtime.js";
import router from "./router/index.js";

EventsOn("goto", function (path) {
  router.push(path);
})

</script>

<template>
  <v-app>
    <StatusBar></StatusBar>
    <RouterView/>
  </v-app>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

.v-overlay__scrim {
  width: 80vw !important;
  margin-left: 20vw !important;
}

.unselectable-text {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
