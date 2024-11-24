<template>
  <v-app>
    <v-main>
      <v-container class="ma-0 pa-0">
        <v-row justify="center" align="center">
          <v-col cols="12" md="6">
            <v-card id="main-box" class="elevation-12" outlined>
              <v-card-title style="--wails-draggable: drag">
              </v-card-title>
              <v-card-text>
                <div v-if="data.resultText" class="pa-3">
                  <div id="result" class="result ma-3" v-text="data.resultText"></div>
                  <v-btn color="primary" @click="copy">复制</v-btn>
                </div>
                <div v-else class="pa-3">
                  <div id="result" class="result ma-3">请选择文本</div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
        <v-snackbar ref="snackbar" v-model="snackbarShow" :timeout="3000" color="success">
          复制成功！
        </v-snackbar>
      </v-container>
    </v-main>
  </v-app>
</template>

<style scoped>
.result {
  max-height: 80vh;
  min-height: 40vh;
  overflow-y: auto;
  background-color: #f1f1f1;
  border-radius: 8px;
  padding: 10px;
  font-size: 16px;
}

#main-box {
  height: 100vh;
}
</style>

<script setup>
import {reactive, ref} from 'vue';
import {ClipboardSetText, EventsOn} from "../../wailsjs/runtime/runtime.js";
import router from "../router/index.js";

const data = reactive({
  resultText: "",
})
const streamEnd = ref(true);
const snackbarShow = ref(false);

const copy = () => {
  ClipboardSetText(data.resultText)
  snackbarShow.value = true;
};

EventsOn("onSearchResult", function (resultText) {
  data.resultText = resultText
})

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
  router.push(path);
};
</script>
