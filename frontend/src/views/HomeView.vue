<script setup lang="ts">
import {onMounted, ref} from "vue";
import {GetOnlineToolsList} from "../../wailsjs/go/main/App";
import type {onlineTools} from "../../wailsjs/go/models";
import {useI18n} from "vue-i18n";
import app from "../App.vue";

const {t, availableLocales: languages, locale} = useI18n();

const size = ref(20)

const onlineToolList = ref<onlineTools.OnlineTool[]>([]);

onMounted(() => {
  GetOnlineToolsList().then((s) => {
    onlineToolList.value = s;
  });
});


const props = defineProps({
  // 这是父组件传递过来的函数
  handleTabsEdit: {
    type: Function
  }
});

function addTab(argetName: string, action: string, name: string) {
  if (props !== undefined && props.handleTabsEdit !== undefined) {
    props.handleTabsEdit(argetName, action, name);
  }
}

</script>

<template>
  <br/>
<!--  <el-scrollbar max-height="95vh" max="89">-->
    <el-space wrap :size="size">
      <el-card class="box-card" v-for="(item, index) in onlineToolList" :key="index" size="100px">
        <el-button
          @click="addTab(item.name, item.router, 'home' + index)">
          {{ t(item.name) }}</el-button>
      </el-card>
    </el-space>
<!--  </el-scrollbar>-->
</template>

<style scoped>
.custom-button {
  background-color: #628a93; /* 默认背景色 */
  color: #000000; /* 默认字体颜色 */
}

.custom-button:hover {
  background-color: #18e80b; /* 悬停背景色 */
}
</style>