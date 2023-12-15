<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, onUnmounted, ref} from "vue";
import {SearchBar} from "search-bar-vue3";
import {GetItemList, GetLanguageList, GetNowLanguage, SetNowLanguage} from "../wailsjs/go/main/App";
import type {configuration} from "../wailsjs/go/models";
import {Location, Setting} from "@element-plus/icons-vue";

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
const list = ref<configuration.Index[]>([]);

onUnmounted(() => {
  document.removeEventListener("keydown", keyboardDownBack, true);
});

onMounted(() => {
  document.addEventListener("keydown", keyboardDownBack, true);
  GetNowLanguage().then((s) => {
    console.log("得到的是: " + s);
    locale.value = s;
  });
  GetItemList().then((s) => {
    list.value = s;
  });
  GetLanguageList().then((s) => {
    listLanguages.value = s;
  });

  // todo 删除上下按钮
  // displayName("pre-match");
  // displayName("next-match");

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
const listLanguages = ref<configuration.Language[]>([]);

const onclickLanguageHandle = (item: string) => {
  SetNowLanguage(item);
  console.log("设置的语言是:" + item)
  locale.value = item;
};
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
                <el-sub-menu index="0">
                  <template #title>
                    <el-icon>
                      <location/>
                    </el-icon>
                    <span>{{ t("index.tool") }}</span>
                  </template>

                  <div
                      class="header-list"
                      v-for="(item, index) in list"
                      :key="index"
                  >
                    <router-link :to="item.route" tag="span">
                      <el-menu-item :index="'0-' + index"
                      >{{ t(item.name) }}
                      </el-menu-item>
                    </router-link>
                  </div>
                </el-sub-menu>
<!--                <el-sub-menu index="1">-->
<!--                  <template #title>-->
<!--                    <el-icon>-->
<!--                      <location/>-->
<!--                    </el-icon>-->
<!--                    <span>Navigator One</span>-->
<!--                  </template>-->

<!--                  <el-menu-item-group title="Group Two">-->
<!--                    <el-menu-item index="1-3">item three</el-menu-item>-->
<!--                  </el-menu-item-group>-->
<!--                  <el-sub-menu index="1-4">-->
<!--                    <template #title>item four</template>-->
<!--                    <el-menu-item index="1-4-1">item one</el-menu-item>-->
<!--                  </el-sub-menu>-->
<!--                </el-sub-menu>-->
                <el-menu-item index="4" @click="settingDrawer = true">
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
</style>
