<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {h, onMounted, onUnmounted, ref} from "vue";
import {SearchBar} from "search-bar-vue3";
import type { Action } from 'element-plus'
import {
  GetCopyData,
  GetCopyHis,
  GetLanguageList,
  GetNowLanguage,
  SetNowLanguage
} from "../wailsjs/go/main/App";
import type {configuration, copy} from "../wailsjs/go/models";
import {DocumentCopy, Tools, Setting, ScaleToOriginal, DocumentRemove} from "@element-plus/icons-vue";
import {ElMessage, ElNotification, ElMessageBox} from "element-plus";

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

// import {GetItemList} from '../../wailsjs/go/main/APP'
const {t, availableLocales: languages, locale} = useI18n();

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
onUnmounted(() => {
  document.removeEventListener("keydown", keyboardDownBack, true);
});
var timer: number | undefined;
onMounted(() => {
  document.addEventListener("keydown", keyboardDownBack, true);
  GetNowLanguage().then((s) => {
    locale.value = s;
  });
  GetLanguageList().then((s) => {
    listLanguages.value = s;
  });
  // 获取剪切板
  timer = setInterval(() => {
    GetCopyData();
  }, 1000);
});

function displayName(name: string) {
  const elements = document.querySelectorAll("." + name);
  for (let i = 0; i < elements.length; i++) {
    const element = elements[i] as HTMLElement;
    element.style.display = "none";
  }
}

const keyboardDownBack = (event: KeyboardEvent) => {
  if (event.ctrlKey && event.key === "f") {
    showSearchBar.value = !showSearchBar.value;
    if (!showSearchBar.value) {
      // 0.2s后将焦点聚焦到搜索框
      setTimeout(() => {
        const inputWrapper = document.querySelector(".input-wrapper");
        if (inputWrapper) {
          const inputElement = inputWrapper.querySelector("input");
          if (inputElement) {
            inputElement.focus();
          }
        }
      }, 200);
    }
  }
  if (event.key === "Escape") {
    showSearchBar.value = true;
  }
};

//=========================================setting======================================
let settingDrawer = ref(false);
let showSearchBar = ref(true);
let copyhis = ref(false);
let copyData = ref(Array<copy.CopyHis>());
const listLanguages = ref<configuration.Language[]>([]);

const onclickLanguageHandle = (item: string) => {
  SetNowLanguage(item);
  locale.value = item;
};

function openCopyHis() {
  GetCopyHis().then((s) => {
    copyData.value = s.reverse();
  });
  copyhis.value = true;
}

function copyDataFun(row: copy.CopyHis) {
  const el = document.createElement("textarea");
  el.value = row.data;
  document.body.appendChild(el);
  el.select();
  document.execCommand("copy");
  document.body.removeChild(el);
  ElMessage({
    message: t("globals.copy-success"),
    type: "success",
  });
}

function clickData(row: string) {
  console.log(row)
  ElMessageBox.alert(row, t("index.full-content"), {
    confirmButtonText: t("globals.close"),
    callback: (action: Action) => {},
  })
}
</script>

<template>
  <div id="appDiff"></div>
  <!--  高亮框-->
  <search-bar
      :root="'#document'"
      id="search-bar"
      :highlightClass="'myHighLight'"
      :selectedClass="'selected-highlight'"
      v-model:hidden="showSearchBar"
  />

  <!--  setting  start-->
  <el-drawer v-model="settingDrawer">
    <template #header>
      <h4>{{ t("index.setting") }}</h4>
    </template>
    <div class="flex flex-wrap items-center">
      <el-dropdown>
        <el-button class="popperClass">
          {{ t("index.language") }}
          <el-icon class="el-icon--right">
            <!--          <setting />-->
          </el-icon>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu>
            <div
                class="header-list"
                v-for="(item, index) in listLanguages"
                :key="index"
            >
              <el-dropdown-item
                  @click="onclickLanguageHandle(item.languageCode)"
              >{{ item.name }}
              </el-dropdown-item>
            </div>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </el-drawer>
  <!--  setting end-->

  <!--  copy history start-->
  <el-drawer v-model="copyhis">
    <template #header>
      <h4>{{ t("index.copy-header") }}</h4>
    </template>
    <div class="flex flex-wrap items-center">

      <el-table :data="copyData" border style="width: 100%" @row-click="copyDataFun">
        <el-table-column prop="dataOmit" :label="t('index.copy-data')" width="350" show-overflow-tooltip	>
          <template v-slot="scope">
            <div class="table-cell-wrapper">
              <span class="text">{{ scope.row.dataOmit }}</span>
              <span class="text" v-if="scope.row.isOmit" @click.stop="clickData(scope.row.data)" style="color: #9494e7; cursor: pointer;">
                {{ t("index.view") }}
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="time" :label="t('index.copy-time')"/>
      </el-table>

    </div>
  </el-drawer>
  <!--  copy history end-->
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <el-scrollbar max-height="100vh">
          <el-row class="tac">
            <el-col :span="12">
              <el-menu
                  default-active="2"
                  class="el-menu-vertical-demo"
                  @open="handleOpen"
                  @close="handleClose"
                  :default-openeds="['0']"
              >
                <router-link to="/">
                  <el-menu-item index="1">
                    <el-icon><Tools /></el-icon>
                    <span>{{ t("index.tool") }}</span>
                  </el-menu-item>
                </router-link>
                <!--                </el-sub-menu>-->
                <el-menu-item index="3" @click="openCopyHis">
                  <el-icon>
                    <DocumentCopy/>
                  </el-icon>
                  <span>{{ t("index.copy-his") }}</span>
                </el-menu-item>

                <router-link to="/onlineTool/code">
                  <el-menu-item index="4">
                    <el-icon><ScaleToOriginal /></el-icon>
                    <span>{{ t("index.code") }}</span>
                  </el-menu-item>
                </router-link>

                <router-link to="/onlineTool/command">
                  <el-menu-item index="5">
                    <el-icon><DocumentRemove /></el-icon>
                    <span>{{ t("tool.command") }}</span>
                  </el-menu-item>
                </router-link>

                <el-menu-item index="6" @click="settingDrawer = true">
                  <el-icon>
                    <setting/>
                  </el-icon>
                  <span>{{ t("index.setting") }}</span>
                </el-menu-item>
              </el-menu>
            </el-col>
          </el-row>
        </el-scrollbar>
      </el-aside>
      <el-main>
        <!--        <el-scrollbar max-height="100vh">-->
        <!--          高亮-->
        <div id="document">
          <router-view></router-view>
        </div>
        <!--        </el-scrollbar>-->
      </el-main>
    </el-container>
  </div>
</template>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}

body {
  overflow: hidden;
}

button {
  color: #fff;
}

.myHighLight {
  background-color: yellow;
}

.selected-highlight {
  background-color: #e3ba6e;
}

/*去除路由的下划线*/
a {
  text-decoration: none;
}

.router-link-active {
  text-decoration: none;
}


.el-table .cell {
  white-space: pre-line !important;
}


</style>
