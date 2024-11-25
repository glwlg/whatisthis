<template>
  <n-layout-header bordered class="status-bar" style="--wails-draggable: drag">
    <n-space justify="space-between" align="center">
      <n-space align="center">
        <n-button quaternary circle @click="toggleMenu">
          <template #icon>
            <n-icon color="#fff"><MenuIcon /></n-icon>
          </template>
        </n-button>
        <span class="unselectable-text" @dblclick="toggleMaximizeWindow">这是什么？</span>
      </n-space>
      
      <n-space>
        <n-button quaternary circle @click="minimizeWindow">
          <template #icon>
            <n-icon color="#fff"><MinimizeIcon /></n-icon>
          </template>
        </n-button>
        <n-button quaternary circle @click="toggleMaximizeWindow">
          <template #icon>
            <n-icon color="#fff">
              <MaximizeIcon v-if="!isMaximized" />
              <RestoreIcon v-else />
            </n-icon>
          </template>
        </n-button>
        <n-button quaternary circle @click="closeWindow">
          <template #icon>
            <n-icon color="#fff"><CloseIcon /></n-icon>
          </template>
        </n-button>
      </n-space>
    </n-space>
  </n-layout-header>

  <n-drawer v-model:show="menuVisible" :width="200" placement="left">
    <n-space vertical style="padding-top: 48px">
      <n-menu :options="menuOptions" @update:value="handleMenuSelect" />
    </n-space>
  </n-drawer>
</template>

<script setup>
import { ref } from 'vue';
import { WindowMinimise, WindowToggleMaximise, WindowHide } from '../../wailsjs/runtime/runtime.js';
import router from "../router/index.js";
import {
  NLayoutHeader,
  NSpace,
  NButton,
  NIcon,
  NDrawer,
  NMenu
} from 'naive-ui';
import {
  MenuOutline as MenuIcon,
  RemoveOutline as MinimizeIcon,
  SquareOutline as MaximizeIcon,
  ContractOutline as RestoreIcon,
  CloseOutline as CloseIcon
} from '@vicons/ionicons5';

const menuVisible = ref(false);
const isMaximized = ref(false);

const menuOptions = [
  {
    label: '应用配置',
    key: 'config'
  },
  {
    label: 'API配置',
    key: 'apiConfig'
  }
];

const handleMenuSelect = (key) => {
  console.log('handleMenuSelect');
  goto(key);
  menuVisible.value = false;
};

const goto = (path) => {
  router.push('/' + path);
};

const toggleMenu = () => {
  menuVisible.value = !menuVisible.value;
};

const minimizeWindow = () => {
  WindowMinimise();
};

const toggleMaximizeWindow = () => {
  WindowToggleMaximise();
  isMaximized.value = !isMaximized.value;
};

const closeWindow = () => {
  WindowHide();
};
</script>

<style scoped>
.status-bar {
  height: 32px;
  padding: 0 10px;
  /* background: rgb(var(--primary-color)); */
  color: white;
}

.unselectable-text {
  user-select: none;
  cursor: default;
}
</style>
