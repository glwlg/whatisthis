<template>
  <n-config-provider>
    <n-layout>
      <n-space vertical>
        <n-card id="main-box" title="应用设置">
          <n-switch v-model:value="robotEnable" @update:value="robotToggle">
            <template #checked>
              划词已启用
            </template>
            <template #unchecked>
              划词已禁用
            </template>
          </n-switch>
          <n-form ref="form" :model="formValue">
            <n-form-item label="划词延迟">
              <n-input-number v-model:value="delay"
                              :min="200"
                              :max="10000"
                              placeholder="请输入"
                              style="width: 100%">
                <template #suffix>
                  ms
                </template>
              </n-input-number>
            </n-form-item>
            <n-space justify="center">
              <n-button @click="gotoWhatIsThisPage">取消</n-button>
              <n-button type="primary" @click="submit">保存</n-button>
            </n-space>
          </n-form>
        </n-card>
      </n-space>
    </n-layout>
  </n-config-provider>
</template>

<style scoped>
#main-box {
  height: 100vh;
}
</style>

<script setup>
import {onMounted, ref} from 'vue';
import {GetConfig, RobotEnable, RobotToggle, SetConfig} from "../../wailsjs/go/main/App.js";
import router from "../router/index.js";
import {
  NButton,
  NCard,
  NConfigProvider,
  NForm,
  NFormItem,
  NInputNumber,
  NLayout,
  NSpace,
  NSwitch,
  useMessage
} from 'naive-ui';

const delay = ref(1000);
const robotEnable = ref(false);
const formValue = ref({});
const message = useMessage();

const submit = () => {
  if (delay.value) {
    console.log({delay: delay.value});
    SetConfig({delay: delay.value.toString()}).then(() => {
      console.log('配置保存成功！');
      message.success('保存成功！');
    })
  }
};

const robotToggle = () => {
  RobotToggle(false);
  RobotEnable().then(enable => {
    console.log(enable);
    robotEnable.value = enable;
  });
};

const gotoWhatIsThisPage = () => {
  router.push('/');
};

onMounted(() => {
  RobotEnable().then(enable => {
    console.log(enable);
    robotEnable.value = enable;
  });
  GetConfig().then(config => {
    console.log('当前延迟配置：', config.delay);
    if (config.delay) {
      delay.value = parseInt(config.delay);
    }
  });
});
</script>
