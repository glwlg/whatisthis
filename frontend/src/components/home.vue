<template>
  <v-app>
    <v-main>
      <v-container>
        <!-- 判断data中的字段是否有值，如果没有则显示表单 -->
        <v-form v-if="!isConfigured" v-model="valid" ref="form" lazy-validation>
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
              v-model="openai.api_version"
              :rules="apiVersionRules"
              label="API Version"
              required
          ></v-text-field>
          <v-btn @click="submit">Submit</v-btn>
        </v-form>
        <div v-else>
          <div id="result" class="result">{{ data.resultText }}</div>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import {onMounted, reactive, ref, watch} from 'vue';
import {GetConfig, SetConfig} from "../../wailsjs/go/main/App.js";

const data = reactive({
  resultText: "",
})

const valid = ref(false);
const openai = ref({api_key: '', base_url: '', api_version: ''});
const apiKeyRules = [v => !!v || 'API Key is required'];
const baseUrlRules = [v => !!v || 'Base URL is required'];
const apiVersionRules = [v => !!v || 'API Version is required'];

const isConfigured = ref(false);

const submit = () => {
  if (valid.value && openai.value.api_key && openai.value.base_url && openai.value.api_version) {
    isConfigured.value = true;
    SetConfig({openai: openai.value})
  }
};

onMounted(() => {
  GetConfig().then(config => {
    console.log(config);
    if (config.openai) {
      openai.value = config.openai;
      isConfigured.value = true;
    }
  });

});


</script>