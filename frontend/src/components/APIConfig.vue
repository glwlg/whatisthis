<template>
  <v-app>
    <v-main>
      <v-container class="ma-0 pa-0">
        <v-row justify="center" align="center">
          <v-col cols="12" md="6">
            <v-card id="main-box" class="elevation-12" outlined>
              <v-card-title style="--wails-draggable: drag">
                <span class="headline">API 设置</span>
              </v-card-title>
              <v-card-text>
                <v-form v-model="valid" ref="form" lazy-validation>
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
                      label="API Version"
                      required
                      outlined
                      dense
                  ></v-text-field>
                  <v-btn color="cancel" @click="gotoWhatIsThisPage" class="mr-4">取消</v-btn>
                  <v-btn color="primary" @click="submit" class="mr-4">保存</v-btn>
                </v-form>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<style scoped>
#main-box {
  height: 100vh;
}
</style>

<script setup>
import {onMounted, ref} from 'vue';
import {GetConfig, SetConfig} from "../../wailsjs/go/main/App.js";
import router from "../router/index.js"; // 引入路由实例

const valid = ref(false);
const openai = ref({api_key: '', base_url: '', model: '', api_version: ''});
const apiKeyRules = [v => !!v || 'API Key is required'];
const baseUrlRules = [v => !!v || 'Base URL is required'];
const modelRules = [v => !!v || 'Model is required'];


const submit = () => {
  if (valid.value && openai.value.api_key && openai.value.base_url && openai.value.model) {
    console.log({openai: openai.value});
    SetConfig({openai: openai.value});
    console.log('配置保存成功！');
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
