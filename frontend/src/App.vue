<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {h, onMounted, onUnmounted, ref} from "vue";
import {SearchBar} from "search-bar-vue3";
import type { Action } from 'element-plus'
import {
  GetCopyData,
  GetCopyHis,
  GetLanguageList,
  GetOrDefaultAllSettings,
  UpdateSystemSettings
} from "../wailsjs/go/main/App";
import type {configuration, copy} from "../wailsjs/go/models";
import {DocumentCopy, Tools, Setting, ScaleToOriginal, DocumentRemove, QuestionFilled} from "@element-plus/icons-vue";
import {ElMessage, ElNotification, ElMessageBox} from "element-plus";
import {goFunc} from "../wailsjs/go/models";

const handleOpen = (key: string, keyPath: string[]) => {

};
const handleClose = (key: string, keyPath: string[]) => {

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
  GetLanguageList().then((s) => {
    listLanguages.value = s;
  });
  // 得到所有的配置
  GetOrDefaultAllSettings().then((s) => {
    systemSetting.value = s;
    systemSettingSnapshot.value = s;
    locale.value = s.language;
  });
  // 获取剪切板
  timer = setInterval(() => {
    GetCopyData();
  }, 1000);
});



function openSettingDrawer() {
  // 深拷贝
  systemSettingSnapshot.value = JSON.parse(JSON.stringify(systemSetting.value));
  settingDrawer.value = true;
}

function updateSetting() {
  // 更新
  UpdateSystemSettings(JSON.stringify(systemSettingSnapshot.value));
  systemSetting.value = JSON.parse(JSON.stringify(systemSettingSnapshot.value));
  // 修改文字
  locale.value = systemSetting.value.language;
  // 提示成功
  ElMessage({
    message: t("globals.copy-success"),
    type: "success",
  });
  // 关闭弹窗
  settingDrawer.value = false;
}

function handleInput1() {
  let number = parseInt(systemSettingSnapshot.value.copySetting.maxCount);
  if (number <= 0) {
    systemSettingSnapshot.value.copySetting.maxCount = "100";
  } else if (number > 10000) {
    systemSettingSnapshot.value.copySetting.maxCount = "10000";
  }
}

function handleInput2() {
  let number = parseInt(systemSettingSnapshot.value.copySetting.maxCountOneData);
  if (number <= 0) {
    systemSettingSnapshot.value.copySetting.maxCountOneData = "100";
  } else if (number > 10000) {
    systemSettingSnapshot.value.copySetting.maxCountOneData = "10000";
  }
}

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
let copyHisDataOmit = ref("");
let copyHisDataOmitHid = ref(false);
let systemSetting = ref(new goFunc.SystemSetting());
let systemSettingSnapshot = ref(new goFunc.SystemSetting());
const visible = ref(false);

const listLanguages = ref<configuration.Language[]>([]);

const onclickLanguageHandle = (item: string) => {
  systemSettingSnapshot.value.language = item;
}

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
  copyHisDataOmit.value = row;

  copyHisDataOmitHid.value = true;
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
      <el-form
          label-position="right"
          label-width="100px"
          :model="systemSettingSnapshot"
          style="max-width: 460px"
      >
        <el-form-item :label="t('index.language')">
          <el-select v-model="systemSettingSnapshot.language" class="m-2" placeholder="Select">
            <el-option
                v-for="item in listLanguages"
                :key="item.languageCode"
                :label="item.name"
                :value="item.languageCode"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-tag>{{ t("copy-his.setting") }}</el-tag>
        </el-form-item>

        <el-form-item :label="t('copy-his.maxCount')">
          <el-input v-model="systemSettingSnapshot.copySetting.maxCount"  @input="handleInput1" type="number" :min="1" :max="10000"/>
        </el-form-item>
        <el-form-item :label="t('copy-his.maxCountOneData')">
          <el-input
              v-model="systemSettingSnapshot.copySetting.maxCountOneData"
              @input="handleInput2"
              type="number"
              :min="40" :max="10000"
          >
            <template #append>
              <el-tooltip :visible="visible" effect="light">
                <template #content>
                  <span style="display: inline-block; word-wrap: break-word; word-break: break-all; max-width: 100ch;color: #ab8585">{{
                    t("copy-his.tip")
                  }}</span>
                </template>
                <el-icon @mouseenter="visible = true" @mouseleave="visible = false"><QuestionFilled/></el-icon>
              </el-tooltip>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="updateSetting">{{ t("globals.submit") }}</el-button>
        </el-form-item>
      </el-form>

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

<!--  copy history dataOmit start-->
  <el-dialog v-model="copyHisDataOmitHid" :title="t('index.full-content')">
    <el-input
        v-model="copyHisDataOmit"
        readonly
        :autosize="{ minRows: 2, maxRows: 18 }"
        type="textarea"
        class="no-wrap-textarea"
        placeholder="Please input"
        resize="none"
    />
  </el-dialog>
  <!--  copy history dataOmit end-->

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

                <el-menu-item index="6" @click="openSettingDrawer">
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
.no-wrap-textarea textarea {
  text-overflow: ellipsis;
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
