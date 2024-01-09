<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, onUnmounted, ref} from "vue";
import {SearchBar} from "search-bar-vue3";

import type {TabPaneName, TabsPaneContext} from "element-plus";
import {ElMessage} from "element-plus";
import {
  GetCopyData,
  GetCopyHis,
  GetFileCache,
  GetLanguageList,
  GetOrDefaultAllSettings,
  UpdateFileCache,
  UpdateSystemSettings,
} from "../wailsjs/go/main/App";
import type {configuration, copy} from "../wailsjs/go/models";
import {goFunc} from "../wailsjs/go/models";
import {DocumentCopy, DocumentRemove, QuestionFilled, ScaleToOriginal, Setting, Tools,Wallet} from "@element-plus/icons-vue";
import commandVue from "../src/views/onlineTool/command.vue";
import HomeView from "../src/views/HomeView.vue";
import codeVue from "../src/views/onlineTool/code.vue";
import timeStampView from "../src/views/onlineTool/Timestamp.vue";
import zkView from "../src/views/onlineTool/zk.vue";
import dbRegisterVue from "../src/views/db/dbRegister.vue";
import dbDetails from "../src/views/db/dbDetails.vue";
import JsonView from "../src/views/onlineTool/json.vue";

// page start
const currentComponent = ref(commandVue);

// page end

const handleOpen = (key: string, keyPath: string[]) => {
};
const handleClose = (key: string, keyPath: string[]) => {
};

const editableTabsValue = ref("1");
const editableTabs = ref([
  {
    title: "index.tool",
    name: "1",
    page: "/",
    suffix : ""
  },
]);

// import {GetItemList} from '../../wailsjs/go/main/APP'
const {t, te,availableLocales: languages, locale} = useI18n();
let tabIndex = 2;

function handleTabsEdit(targetName: string, action: string, name: string, suffix="") {
  for (let i = 0; i < editableTabs.value.length; i++) {
    if (editableTabs.value[i].name == name) {
      editableTabsValue.value = name;
      return;
    }
  }
  editableTabs.value.push({
    title: targetName,
    name: name,
    page: action,
    suffix: suffix,
  });
  editableTabsValue.value = name;
  updateTabCache();
}

const handleTabsEdit2 = (
    targetName: TabPaneName | undefined,
    action: "remove" | "add"
) => {
  const tabs = editableTabs.value;
  let activeName = editableTabsValue.value;
  if (activeName === targetName) {
    tabs.forEach((tab, index) => {
      if (tab.name === targetName) {
        const nextTab = tabs[index + 1] || tabs[index - 1];
        if (nextTab) {
          activeName = nextTab.name;
        }
      }
    });
  }
  editableTabsValue.value = activeName;
  editableTabs.value = tabs.filter((tab) => tab.name !== targetName);
  updateTabCache();
};



function tabClick(pane: TabsPaneContext, ev: Event) {
  UpdateFileCache("cache/tabSelectCache", pane.props.name + "");
}
function tabClick2(name: TabPaneName) {
  setTimeout(() => {
    // 得到上次缓存的json
    UpdateFileCache("cache/tabSelectCache", editableTabsValue.value);
  }, 100);
}

function GetNowTabName() {
  return editableTabsValue.value;
}

// 缓存当前页面的数据内容
function updateTabCache() {
  UpdateFileCache("cache/tabCache", JSON.stringify(editableTabs.value));
}

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
  // 得到上次的标签页
  GetFileCache("cache/tabCache").then((s) => {
    try {
      if (s !== "") {
        editableTabs.value = JSON.parse(s);
      }
    } catch (error) {
      console.error("Invalid JSON string:", error);
    }
  });
  // 得到上次选择的标签页
  GetFileCache("cache/tabSelectCache").then((s) => {
    console.log("上次的结果: " + s);
    if (s !== "") {
      editableTabsValue.value = s;
    }
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
    message: t("globals.operations-success"),
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
  let number = parseInt(
      systemSettingSnapshot.value.copySetting.maxCountOneData
  );
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

function labelValue(title : string, suffix : string) {
 return t(title) + (suffix == null ? "" : suffix);
}

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
  copyHisDataOmit.value = row;

  copyHisDataOmitHid.value = true;
}
// 屏蔽系统右键菜单
document.oncontextmenu = function (e) {
  return false;
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
          <el-select
              v-model="systemSettingSnapshot.language"
              class="m-2"
              placeholder="Select"
          >
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
          <el-input
              v-model="systemSettingSnapshot.copySetting.maxCount"
              @input="handleInput1"
              type="number"
              :min="1"
              :max="10000"
          />
        </el-form-item>
        <el-form-item :label="t('copy-his.maxCountOneData')">
          <el-input
              v-model="systemSettingSnapshot.copySetting.maxCountOneData"
              @input="handleInput2"
              type="number"
              :min="40"
              :max="10000"
          >
            <template #append>
              <el-tooltip :visible="visible" effect="light">
                <template #content>
                  <span
                      style="
                      display: inline-block;
                      word-wrap: break-word;
                      word-break: break-all;
                      max-width: 100ch;
                      color: #ab8585;
                    "
                  >{{ t("copy-his.tip") }}</span
                  >
                </template>
                <el-icon
                    @mouseenter="visible = true"
                    @mouseleave="visible = false"
                >
                  <QuestionFilled
                  />
                </el-icon>
              </el-tooltip>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="updateSetting">{{
              t("globals.submit")
            }}
          </el-button>
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
      <el-table
          :data="copyData"
          border
          style="width: 100%"
          @row-click="copyDataFun"
      >
        <el-table-column
            prop="dataOmit"
            :label="t('index.copy-data')"
            width="340"
        >
          <template v-slot="scope">
            <div class="table-cell-wrapper">
              <span class="text">{{ scope.row.dataOmit }}</span>
              <span
                  class="text"
                  v-if="scope.row.isOmit"
                  @click.stop="clickData(scope.row.data)"
                  style="color: #9494e7; cursor: pointer"
              >
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
                <el-menu-item
                    index="1"
                    @click="handleTabsEdit('index.tool', '/', '1')"
                >
                  <el-icon>
                    <Tools/>
                  </el-icon>
                  <span>{{ t("index.tool") }}</span>
                </el-menu-item>
                <!--                </el-sub-menu>-->
                <el-menu-item index="3" @click="openCopyHis">
                  <el-icon>
                    <DocumentCopy/>
                  </el-icon>
                  <span>{{ t("index.copy-his") }}</span>
                </el-menu-item>

                <el-menu-item
                    index="4"
                    @click="handleTabsEdit('index.code', '/onlineTool/code', '4')"
                >
                  <el-icon>
                    <ScaleToOriginal/>
                  </el-icon>
                  <span>{{ t("index.code") }}</span>
                </el-menu-item>

                <el-menu-item
                    index="5"
                    @click="
                    handleTabsEdit('tool.command', '/onlineTool/command', '5')
                  "
                >
                  <el-icon>
                    <DocumentRemove/>
                  </el-icon>
                  <span>{{ t("tool.command") }}</span>
                </el-menu-item>

                <el-menu-item
                    index="6"
                    @click="handleTabsEdit('index.db', '/db/dbRegister', '6')"
                >
                  <el-icon>
                    <Wallet/>
                  </el-icon>
                  <span>{{ t("index.db") }}</span>
                </el-menu-item>

                <el-menu-item index="7" @click="openSettingDrawer">
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
          <el-tabs
              v-model="editableTabsValue"
              type="card"
              editable
              closable
              @edit="handleTabsEdit2"
              @tab-click="tabClick"
              @tab-remove="tabClick2"
          >
<!--            "te(item.title) ? t(item.title) : item.title" 是否国际化的判断-->
            <el-tab-pane
                v-for="item in editableTabs"
                :key="item.name"
                :label="labelValue(item.title, item.suffix)"
                :name="item.name"
                :lazy="true"
            >
<!--               :lazy="true" 是必须的, 解决一些页面一开始加载,导致js未生效问题,如字符串工具-->
              <template v-if="item.page === '/onlineTool/code'">
                <codeVue></codeVue>
              </template>
              <template v-else-if="item.page === '/onlineTool/command'">
                <commandVue></commandVue>
                <!--                <component :is="currentComponent"></component>-->
              </template>
              <template v-else-if="item.page === '/db/dbRegister'">
                <dbRegisterVue :handleTabsEdit="handleTabsEdit"></dbRegisterVue>
              </template>
              <template v-else-if="item.page === '/'">
                <HomeView :handleTabsEdit="handleTabsEdit"></HomeView>
              </template>
              <template v-else-if="item.page === '/onlineTool/json'">
                <JsonView></JsonView>
              </template>
              <template v-else-if="item.page === '/onlineTool/timestamp'">
                <timeStampView></timeStampView>
              </template>
              <template v-else-if="item.page === '/onlineTool/zk'">
                <zkView></zkView>
              </template>
              <template v-else-if="item.page === '/db/dbDetails'">
                <dbDetails :dbName="item.name" :GetNowTabName="GetNowTabName"></dbDetails>
              </template>

            </el-tab-pane>
          </el-tabs>
        </div>
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
body {
}
.no-wrap-textarea textarea {
  text-overflow: ellipsis;
}

body {
  overflow-y: hidden;
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

.el-tabs__new-tab {
  display: none !important;
}
</style>
