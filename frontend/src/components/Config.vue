<template>
  <v-app>
    <v-main>
      <v-container class="ma-0 pa-0">
        <v-row justify="center" align="center">
          <v-col cols="12" md="6">
            <v-card id="main-box" class="elevation-12" outlined>
              <v-card-title style="--wails-draggable: drag">
                <span class="headline">应用设置</span>
              </v-card-title>
              <v-card-text>
                <v-form v-model="valid" ref="form" lazy-validation>
                  <v-switch
                      :model-value="robotEnable"
                      color="primary"
                      label="开启划词"
                      @change="robotToggle"
                  ></v-switch>
                  <v-text-field
                      v-model="delay"
                      :rules="delayRules"
                      label="划词延迟"
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
import {GetConfig, RobotEnable, RobotToggle, SetConfig} from "../../wailsjs/go/main/App.js";
import router from "../router/index.js";
import {EventsOn} from "../../wailsjs/runtime/runtime.js"; // 引入路由实例

const valid = ref(false);
const delay = ref(1000);
const delayRules = [v => !!v || v < 200 || '延迟必须大于200ms'];
const robotEnable = ref(false);

const submit = () => {
  if (delay.value) {
    console.log({delay: delay.value});
    SetConfig({delay: delay.value});
    console.log('配置保存成功！');
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

EventsOn("robotToggle", function () {
  console.log("robotToggle");
  RobotEnable().then(enable => {
    console.log(enable);
    robotEnable.value = enable;
  });
})

onMounted(() => {
  RobotEnable().then(enable => {
    console.log(enable);
    robotEnable.value = enable;
  });
  GetConfig().then(config => {
    console.log(config);
    if (delay.value) {
      delay.value = config.delay;
    }
  });

});
</script>
