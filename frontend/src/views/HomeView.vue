<script setup lang="ts">
import {onMounted, ref} from "vue";
import {GetOnlineToolsList} from "../../wailsjs/go/main/App";
import type {onlineTools} from "../../wailsjs/go/models";
import {useI18n} from "vue-i18n";

const {t, availableLocales: languages, locale} = useI18n();

const size = ref(20)

const onlineToolList = ref<onlineTools.OnlineTool[]>([]);

onMounted(() => {
  GetOnlineToolsList().then((s) => {
    onlineToolList.value = s;
  });
});

</script>

<template>
  <el-slider v-model="size"/>
  <el-scrollbar max-height="95vh" max="89">
    <el-space wrap :size="size">
      <el-card class="box-card" v-for="(item, index) in onlineToolList" :key="index" size="100px">
        <router-link :to="item.router" tag="span">
          <el-button >{{ t(item.name) }}</el-button>
        </router-link>
      </el-card>
    </el-space>
  </el-scrollbar>
</template>

<style>
.custom-button {
  background-color: #628a93; /* 默认背景色 */
  color: #000000; /* 默认字体颜色 */
}

.custom-button:hover {
  background-color: #18e80b; /* 悬停背景色 */
}
</style>