<script lang="ts" setup>
import type {ElScrollbar} from "element-plus";
import {ElMessage, ElNotification} from "element-plus";
import {h, onMounted, ref} from 'vue'
import {GetCommand} from "../../../wailsjs/go/main/App";
import type {index} from "../../../wailsjs/go/models";
import {useI18n} from "vue-i18n";

const {t, availableLocales: languages, locale} = useI18n();

const commandData = ref(Array<index.CommandPojo>());

onMounted(() => {
  var topTable = document.getElementById("top-table");
  if (topTable) {
    var bodyWrapper = topTable.getElementsByClassName(
        "el-table__body-wrapper"
    )[0];
    if (bodyWrapper && bodyWrapper.parentNode) {
      bodyWrapper.parentNode.removeChild(bodyWrapper);
    }
  }


  GetCommand().then((value) => {
    console.log(value);
    commandData.value = value;
  })
})

const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>()

function copyData(row: index.OneCommand) {
  const el = document.createElement("textarea");
  el.value = row.command;
  document.body.appendChild(el);
  el.select();
  document.execCommand("copy");
  document.body.removeChild(el);
  if (row.url == null || row.url == "") {
    ElMessage({
      message: t("command.copy-success"),
      type: "success",
    });
  } else {
    ElNotification({
      title: t("command.copy-success"),
      message: h('a', {onClick: () => openWebsite(row.url), style: {cursor: 'pointer', color: "#5797eb"}},
          t('command.more')),
      type: "success",
    });
  }

}

function openWebsite(url: string) {
  window.open(url, "_blank");
}

function AAA(id: string) {
  let height = 0;
  for (let i = 0; i < commandData.value.length; i++) {
    const commandDataKey = commandData.value[i];
    if (commandDataKey.name === id) {
      break;
    }
    // 获取 a 元素
    let elementById = document.getElementById(commandDataKey.name);
    if (elementById == null) {
      continue;
    }
    height += elementById.offsetHeight;
  }
  scrollbarRef.value!.setScrollTop(height);
}
</script>

<template>
  <div class="common-layout" id="command-layout">
    <el-container>
      <el-main>

        <div class="common-layout">
          <el-container>
            <el-header height="30px">
              <el-table
                  :data="[]"
                  style="width: 100%"
                  id="top-table"
                  border
              >
                <el-table-column type="index" prop="index" :label="t('command.index')" width="100"/>
                <el-table-column prop="desc" :label="t('command.command')" width="500"/>
                <el-table-column prop="command" :label="t('command.description')"/>
              </el-table>
            </el-header>
            <el-main>
              <el-scrollbar max-height="100vh" ref="scrollbarRef" id="el-scrollbar">
                <div v-for="(value, key) in commandData" :key="key" :id="value.name">
                  <h1 style="font-size: 20px" :ref="value.name">{{ value.name }}</h1>
                  <br/>
                  <el-table
                      :data="value.commands"
                      style="width: 100%"
                      @row-click="copyData"

                      :show-header="false"
                      border
                  >
                    <el-table-column type="index" prop="index" :label="t('command.index')" width="100"/>
                    <el-table-column prop="desc" :label="t('command.command')" width="500"/>
                    <el-table-column prop="command" :label="t('command.description')"/>
                  </el-table>
                  <br/>
                </div>
              </el-scrollbar>
            </el-main>
          </el-container>
        </div>


      </el-main>
      <el-aside width="200px">
        <div style="height:30px"></div>
        <div v-for="(value, key) in commandData" :key="key" style="margin-top: 20px;margin-left: 20px">
          <el-button @click="AAA(value.name)">{{ value.name }}</el-button>
        </div>

      </el-aside>
    </el-container>
  </div>


</template>


<style>
.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}

.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}

#command-layout {
  margin-top: -39px;
}
</style>
