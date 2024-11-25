<template>
  <n-config-provider>
    <n-layout>
      <n-space vertical>
        <n-card id="main-box" title="API 设置">
          <n-form ref="form" :model="openai" :rules="rules">
            <n-form-item label="API Key" path="api_key">
              <n-input v-model:value="openai.api_key" placeholder="请输入"/>
            </n-form-item>
            <n-form-item label="Base URL" path="base_url">
              <n-input v-model:value="openai.base_url" placeholder="请输入"/>
            </n-form-item>
            <n-form-item label="模型" path="model">
              <n-input v-model:value="openai.model" placeholder="请输入"/>
            </n-form-item>
            <n-form-item label="API Version">
              <n-input v-model:value="openai.api_version" placeholder="请输入"/>
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
import {GetConfig, SetConfig} from "../../wailsjs/go/main/App.js";
import router from "../router/index.js";
import {NButton, NCard, NConfigProvider, NForm, NFormItem, NInput, NLayout, NSpace, useMessage} from 'naive-ui';

const rules = ref({
  api_key: [{
    required: true,
    trigger: ['input', 'blur'],
    message: '请输入API Key'
  }],
  base_url: {
    required: true,
    trigger: ['input', 'blur'],
    message: '请输入Base URL'
  },
  model: {
    required: true,
    trigger: ['input', 'blur'],
    message: '请输入模型名称'
  }
});
const message = useMessage();

const openai = ref({
  api_key: '',
  base_url: '',
  model: '',
  api_version: ''
});

const submit = () => {
  if (openai.value.api_key && openai.value.base_url && openai.value.model) {
    SetConfig({openai: openai.value});
    message.success('保存成功！');
  }
};

const gotoWhatIsThisPage = () => {
  router.push('/');
};

onMounted(() => {
  GetConfig().then(config => {
    console.log(config);
    if (config.openai) {
      openai.value = config.openai;
    }
  });
});
</script>
