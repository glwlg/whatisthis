<template>
  <v-app>
    <v-main>
      <v-container class="ma-0 pa-0">
        <v-row justify="center" align="center">
          <v-col cols="12" sm="8" md="6">
            <v-card id="main-box" class="elevation-12" outlined>
              <v-card-title v-if="showSetting"  style="--wails-draggable: drag">
                <span class="headline">API 设置</span>
              </v-card-title>
              <v-card-title v-else  style="--wails-draggable: drag">
                <span class="headline">这是什么?</span>
              </v-card-title>
              <v-card-text>
                <v-form v-if="showSetting" v-model="valid" ref="form" lazy-validation>
                  <v-text-field
                      v-model="openai.api_key"
                      :rules="apiKeyRules"
                      label="API Key"
                      required
                      outlined
                      dense
                  ></v-text-field>
                  <v-text-field
                      v-model="openai.base_url"
                      :rules="baseUrlRules"
                      label="Base URL"
                      required
                      outlined
                      dense
                  ></v-text-field>
                  <v-text-field
                      v-model="openai.model"
                      :rules="modelRules"
                      label="模型"
                      required
                      outlined
                      dense
                  ></v-text-field>
                  <v-text-field
                      v-model="openai.api_version"
                      :rules="apiVersionRules"
                      label="API Version"
                      required
                      outlined
                      dense
                  ></v-text-field>
                  <v-btn color="primary" @click="submit" class="mr-4">Submit</v-btn>
                </v-form>
                <div v-else>
                  <div v-if="data.resultText" class="pa-3">
                    <div id="result" class="result ma-3" v-text="data.resultText"></div>
                    <v-btn color="primary" @click="copy">Copy</v-btn>
                  </div>
                </div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<style scoped>
.result {
  max-height: 80vh;
  overflow-y: auto;
  background-color: #f1f1f1;
  border-radius: 8px;
  padding: 10px;
  font-size: 16px;
}
#main-box{
  height: 100vh;
}
</style>

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