<script setup>
import {RouterView} from 'vue-router';
import StatusBar from "./components/StatusBar.vue";
import {EventsOn} from "../wailsjs/runtime/runtime.js";
import router from "./router/index.js";
import { NConfigProvider, NMessageProvider } from 'naive-ui'
import { darkTheme, useOsTheme } from 'naive-ui';
import { computed } from 'vue';

const themeOverrides = {
  common: {
    primaryColor: '#FF9671',
    primaryColorHover: '#FFA88C',
    primaryColorPressed: '#E88664',
    primaryColorSuppl: '#FFB1A0'
  },
  Card: {
    borderRadius: '12px',
    padding: '24px'
  },
  Button: {
    borderRadius: '8px',
    textColor: '#333',
    colorPrimary: '#FF9671',
    colorHoverPrimary: '#FFA88C',
    colorPressedPrimary: '#E88664',
    colorFocusPrimary: '#FFB1A0'
  }
};

EventsOn("goto", function (path) {
  router.push(path);
})

</script>

<template>
  <n-config-provider :theme="darkTheme" :theme-overrides="themeOverrides">
    <n-message-provider>
      <StatusBar></StatusBar>
      <RouterView/>
      <n-global-style />
    </n-message-provider>
  </n-config-provider>
</template>

<style>
body {
  margin: 0;
}
.unselectable-text {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
