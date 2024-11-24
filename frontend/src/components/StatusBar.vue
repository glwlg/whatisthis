<template>
  <v-app-bar
      app
      color="primary"
      dense
      elevation="0"
      class="status-bar"
      style="--wails-draggable: drag"
  >
    <template v-slot:prepend>
      <v-btn
          icon
          class="menu-icon"
          @click="toggleMenu"
      >
        <v-app-bar-nav-icon></v-app-bar-nav-icon>
      </v-btn>
    </template>
    <v-toolbar-title @dblclick="toggleMaximizeWindow" class="unselectable-text">这是什么？</v-toolbar-title>
    <template v-slot:append>
      <v-btn
          icon
          @click="minimizeWindow"
      >
        <v-icon icon="mdi-window-minimize"></v-icon>
      </v-btn>
      <v-btn
          icon
          @click="toggleMaximizeWindow"
      >
        <v-icon :icon="isMaximized? 'mdi-window-restore' : 'mdi-window-maximize'"></v-icon>
      </v-btn>
      <v-btn
          icon
          @click="closeWindow"
      >
        <v-icon>mdi-window-close</v-icon>
      </v-btn>
    </template>
  </v-app-bar>

  <!-- 在这里添加菜单的显示区域 -->
  <div v-if="menuVisible" class="menu-container" transition="scroll-x-transition">
    <v-list>
      <v-list-item>
        <v-list-item-title @click="goto('config')" v-text="'应用配置'"></v-list-item-title>
      </v-list-item>
      <v-list-item>
        <v-list-item-title @click="goto('apiConfig')" v-text="'API配置'"></v-list-item-title>
      </v-list-item>
    </v-list>
  </div>
  <v-overlay v-model="menuVisible" @click="toggleMenu" />
</template>

<script>
import { defineComponent, ref } from 'vue';
import router from "../router/index.js";
import {WindowHide, WindowMinimise, WindowToggleMaximise} from "../../wailsjs/runtime/runtime.js";

export default defineComponent({
  setup() {
    const isMaximized = ref(false);
    const menuVisible = ref(false); // 新增：用于控制菜单的显示状态

    const toggleMenu = () => {
      console.log('toggleMenu');
      menuVisible.value = !menuVisible.value;
    };
    const goto = (path) => {
      toggleMenu();
      router.push(path);
    };

    const minimizeWindow = () => {
      console.log('窗口最小化');
      WindowMinimise();
    };

    const toggleMaximizeWindow = () => {
      WindowToggleMaximise();
      if (isMaximized.value) {
        console.log('窗口恢复');
      } else {
        console.log('窗口最大化');
      }
      isMaximized.value =!isMaximized.value;
    };

    const closeWindow = () => {
      console.log('窗口关闭');
      WindowHide();
    };

    return {
      goto,
      toggleMenu,
      minimizeWindow,
      toggleMaximizeWindow,
      closeWindow,
      isMaximized,
      menuVisible // 新增：将菜单显示状态暴露给模板
    };
  }
});
</script>

<style scoped>
.status-bar {
  position: relative!important;
  z-index: 1000;
}

.menu-container {
  position: absolute;
  padding-top: 8vh;
  top: 0;
  left: 0;
  width: 20vw;
  min-width: 100px;
  height: 100vh;
  background-color: #fff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  z-index: 999;
}
</style>
