<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, ref, computed} from "vue";

import {GetItemList} from "../../wailsjs/go/main/App";
import type {configuration} from "../../wailsjs/go/models";

// import {GetItemList} from '../../wailsjs/go/main/APP'
const {t, availableLocales: languages, locale} = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};

const onclickMinimise = () => {
};

const onclickQuit = () => {
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
function currentView(pageStr: any) {
  computed(() => {
    return pageStr;
  });
}

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
const list = ref<configuration.Index[]>([]);
onMounted(() => {
  // 在初始化时设置list变量的值
  GetItemList().then((s) => {
    list.value = s;
  });
});
</script>

<template>
  <div class="list-container">
    <div class="rounded-rectangle" v-for="(item, index) in list" :key="index">
      <h3 class="element-In" onclick="currentView(item.route)">{{ t(item.name) }} AAAAAAAAAAAA</h3>
    </div>
  </div>
</template>

<style lang="scss">
@import url("src/assets/css/reset.css");
@import url("src/assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
}

#app {

}

.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: rgba(171, 126, 220, 0.9);

  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #ab7edc;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;

      &:hover,
      &.router-link-exact-active {
        background-color: #d7a8d8;
        color: #ffffff;
      }
    }
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;

    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #c3c3c3;
      overflow: hidden;

      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #ff050542;
          cursor: pointer;
        }

        &.active {
          background-color: #ff050542;
          color: #ffffff;
          cursor: not-allowed;
        }
      }
    }

    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;

      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        margin-left: 8px;
        background-color: #ab7edc;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #d7a8d8;
          color: #ffffff;
          cursor: pointer;
        }
      }
    }
  }
}

.list-container {
  display: flex;
  flex-wrap: wrap;
}

.rounded-rectangle {
  background-color: #ffd1dc;
  size: 300px;
}

.rounded-rectangle {
  margin-right: 40px;
  margin-bottom: 40px;
}

.element-In {
  font-size: 30px;
}

</style>
