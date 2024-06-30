<template>
  <v-app>
    <v-main style="--wails-draggable: drag">
      <v-container>
        <!-- 判断data中的字段是否有值，如果没有则显示表单 -->
        <v-form v-if="showSetting" v-model="valid" ref="form" lazy-validation>
          <v-text-field
              v-model="openai.api_key"
              :rules="apiKeyRules"
              label="API Key"
              required
          ></v-text-field>
          <v-text-field
              v-model="openai.base_url"
              :rules="baseUrlRules"
              label="Base URL"
              required
          ></v-text-field>
          <v-text-field
              v-model="openai.model"
              :rules="modelRules"
              label="模型"
              required
          ></v-text-field>
          <v-text-field
              v-model="openai.api_version"
              :rules="apiVersionRules"
              label="API Version"
              required
          ></v-text-field>
          <v-btn @click="submit">Submit</v-btn>
        </v-form>
        <div v-else>
          <div v-if="data.resultText">
            <div id="result" class="result">{{ data.resultText }}</div>
            <v-btn @click="copy">copy</v-btn>
          </div>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import {onMounted, reactive, ref} from 'vue';
import {GetConfig, SetConfig, WriteToClipboard} from "../../wailsjs/go/main/App.js";
import {EventsOn} from "../../wailsjs/runtime/runtime.js";

const data = reactive({
  resultText: "",
})

const valid = ref(false);
const openai = ref({api_key: '', base_url: '', model: '', api_version: ''});
const apiKeyRules = [v => !!v || 'API Key is required'];
const baseUrlRules = [v => !!v || 'Base URL is required'];
const apiVersionRules = [v => !!v || 'API Version is required'];
const modelRules = [v => !!v || 'Model is required'];

const showSetting = ref(true);

const submit = () => {
  if (valid.value && openai.value.api_key && openai.value.base_url && openai.value.model && openai.value.api_version) {
    showSetting.value = false;
    SetConfig({openai: openai.value})
  }
};

const copy = () => {
  WriteToClipboard(data.resultText)
};

onMounted(() => {
  GetConfig().then(config => {
    console.log(config);
    if (config.openai) {
      openai.value = config.openai;
      showSetting.value = false;
    }
  });

});


EventsOn("onSearchResult", function (resultText) {
  data.resultText = resultText
})

EventsOn("openSetting", function () {
  showSetting.value = true;
})



</script>