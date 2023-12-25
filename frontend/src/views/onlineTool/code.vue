<script lang="ts" setup>
import {onMounted, ref} from "vue";
import type {TabsPaneContext} from "element-plus";
import {ElMessage} from "element-plus";
import {GetCodeGroup} from "../../../wailsjs/go/main/App";
import type {code} from "../../../wailsjs/go/models";
import {useI18n} from "vue-i18n";
// 代码的类型
const languageTypeList = ref(new Set<string>());
const languageType = ref("ALL");
const codeGroup = ref(Array<code.CodeGroup>());
const {t, availableLocales: languages, locale} = useI18n();

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}

onMounted(() => {

  GetCodeGroup().then((value) => {
    codeGroup.value = value;
    languageTypeList.value.add("ALL");
    for (let nowCodeGroup of value) {
      for (let oneCode of nowCodeGroup.codeList) {
        // 添加代码类型到Set中
        languageTypeList.value.add(oneCode.type);
      }
    }
  });
})

function hideOneGroup(group: code.CodeGroup) {
  if (languageType.value === "ALL") {
    return true;
  }
  for (let oneCode of group.codeList) {
    if (oneCode.type === languageType.value) {
      return true;
    }
  }
  return false;
}

function copyData(code: string) {
  const el = document.createElement("textarea");
  el.value = code;
  document.body.appendChild(el);
  el.select();
  document.execCommand("copy");
  document.body.removeChild(el);
  ElMessage({
    message: t("globals.copy-success"),
    type: "success",
  });

}

</script>

<template>

  <el-tabs v-model="languageType" class="demo-tabs" @tab-click="handleClick">
    <div v-for="(value, key) in languageTypeList" :key="key">
      <el-scrollbar height="90vh" :sync="true">
        <el-tab-pane :label="value" :name="value" class="llllYYYYYY">
          <div v-for="group in codeGroup">
            <div v-if="hideOneGroup(group)">
              <h1 style="font-size: 20px" :ref="group.name">{{ group.name }}</h1>
              <!--          一个代码的Tab-->
              <el-tabs type="card" class="demo-tabs"
                       :model-value="languageType === 'ALL' ? group.codeList[0].type : languageType">
                <!--            遍历所有的类型-->
                <div v-for="code in group.codeList">
                  <el-tab-pane :label="code.type" :name="code.type"
                               v-if="languageType === 'ALL' || code.type === languageType">
                    <el-button
                        style="margin-top: 10px;margin-left: 95%;position: absolute;z-index: 9"
                        key="primary"
                        type="primary"
                        @click="copyData(code.code)"
                        text
                    >复制
                    </el-button>
                    <el-input
                        v-model="code.code"
                        autosize
                        resize="none"
                        readonly
                        type="textarea"
                        placeholder="Please input"
                    />
                  </el-tab-pane>
                </div>
              </el-tabs>
            </div>
          </div>

        </el-tab-pane>
      </el-scrollbar>
    </div>
  </el-tabs>
</template>

<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.el-tabs__content {
  padding: 0px !important;
}

.llllYYYYYY {
  display: block !important;
}

</style>
