<template>
  <n-layout>
    <n-space vertical size="large" style="max-width: 90%; margin: 0 auto;">
      <n-card id="main-box">
        <template #default>
          <n-space vertical>
            <div class="result-container">
              <n-text v-if="data.resultText" class="result">
                {{ data.resultText }}
              </n-text>
              <n-text v-else>
                请选择文本
              </n-text>
            </div>
            <n-divider v-if="data.resultText"/>
            <n-space v-if="data.resultText" justify="center">
              <n-button type="primary" @click="copy">
                <template #icon>
                  <n-icon>
                    <CopyOutline/>
                  </n-icon>
                </template>
                复制
              </n-button>
            </n-space>
          </n-space>
        </template>
      </n-card>
    </n-space>
  </n-layout>
</template>

<style scoped>
.result-container {
  height: calc(90vh - 180px);
  overflow-y: auto;
  padding: 16px;
}

.result {
  font-size: 16px;
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-all;
}

#main-box {
  margin-top: 16px;
  margin-bottom: 16px;
  height: 90vh;
}
</style>

<script setup>
import {reactive, ref} from 'vue';
import {ClipboardSetText, EventsOn} from "../../wailsjs/runtime/runtime.js";
import {CopyOutline} from '@vicons/ionicons5';
import {NButton, NCard, NDivider, NIcon, NLayout, NSpace, NText, useMessage} from 'naive-ui';
import router from "../router/index.js";

const message = useMessage();
const data = reactive({
  resultText: ''
});

const copy = () => {
  if (data.resultText) {
    ClipboardSetText(data.resultText);
    message.success('复制成功！');
  }
};

EventsOn("onSearchResult", function (resultText) {
  data.resultText = resultText;
});

EventsOn("onSearchResultStream", function (resultText) {
  if (streamEnd.value) {
    data.resultText = "";
    streamEnd.value = false;
  }
  data.resultText += resultText;
})
EventsOn("onSearchResultStreamEnd", function (resultText) {
  streamEnd.value = true;
})

EventsOn("openSetting", function () {
  goto('config')
})
EventsOn("openApiConfig", function () {
  goto('apiConfig')
})
const goto = (path) => {
  router.push('/' + path);
};
const streamEnd = ref(true);
</script>
