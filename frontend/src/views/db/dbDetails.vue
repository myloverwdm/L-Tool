<template>
  <el-empty description="数据库无法连接" v-if="allIsEmpty" />
  <el-affix position="bottom" :offset="20" v-if="isHidenDdatabase">
    <el-icon class="arrow-icon" @click="changecollapse"  style="opacity: 0.4;position: absolute;cursor: grab; top: 50%; left: 0; transform: translateY(-50%); width: 20px;z-index: 99">
      <ArrowRightBold />
    </el-icon>
  </el-affix>
  <div class="common-layout">
    <el-container>
      <el-aside width="10px" v-if="isHidenDdatabase" @click="changecollapse" >
        <div class="arrow" style="width: 20px; height: 90vh; background-color: #a3e1ff; position: relative;opacity: 0.2;"></div>
      </el-aside>
      <el-main>
        <div class="common-layout">
          <el-container v-if="!allIsEmpty">
            <el-aside :width="collapse">
              <!--        这里是库名列表-->
              <el-table :data="tableData" style="width: 100%" height="90vh"  :row-class-name="tableRowClassName" @row-contextmenu="rowContextmenu">
                <el-table-column prop="dataBaseName" >
                  <template #header>
                    <div style="white-space: nowrap; display: flex; align-items: center;">
                      <span>数据库名</span>
                      <el-icon @click="GetDataBaseListByRegNameD2" style="margin-left: auto; cursor: grab;">
                        <Refresh/>
                      </el-icon>
                      <el-icon @click="changecollapse" style="margin-left: auto; cursor: grab;">
                        <ArrowLeft/>
                      </el-icon>
                    </div>
                  </template>
                </el-table-column>
              </el-table>

            </el-aside>
            <el-container>
              <el-header>

              </el-header>
              <el-main>{{ dbName }}</el-main>
              <el-footer>{{ dbName }}</el-footer>
            </el-container>
          </el-container>
        </div>
      </el-main>
    </el-container>
  </div>

  <div style="display: none" id="dbName">{{ dbName }}</div>

</template>


<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {ArrowLeft, ArrowRightBold, Refresh} from "@element-plus/icons-vue";
import {PingDbWithName, GetDataBaseListByRegName} from "../../../wailsjs/go/main/App";
import {languages} from "monaco-editor";
import json = languages.json;
import {ElMessage} from "element-plus";


const collapse = ref("200px");
// 是否连接不上，若是连接不上则显示空状态
const allIsEmpty = ref(false);
// 当前数据库的内容是否隐藏
const isHidenDdatabase = ref(false);
// 当前数据库的内容隐藏后，显示的左侧条，是否显示icon
const isHidenDdatabaseIcon = ref(false);

// let dbName = ""l
// const dbName = ref("");

onMounted(() => {
  // 1. 判断是否能够连接数据库
  let htmlElement = document.getElementById("dbName");
  if (htmlElement == null) {
    return
  }
  // 以 db: 开头， 所以是从第三个开始
  var dbName = htmlElement.innerHTML.substring(3, htmlElement.innerHTML.length);
  if (dbName == "") {
    return;
  }
  console.log(htmlElement.innerHTML + "    " + dbName)
  PingDbWithName(dbName).then((s) => {
    allIsEmpty.value = !s;
    if (s) {
      // 得到所有数据库
      GetDataBaseListByRegNameD(false, dbName);
    } else {
      ElMessage({
        message: "数据库无法连接",
        type: "warning",
      });
    }
  });
  // updateTableDataInDb();
});

function GetDataBaseListByRegNameD(alertM : boolean, dbName: string) {
  console.log("=======================" + dbName)
  GetDataBaseListByRegName(dbName).then((s) => {
    tableData.value = [];
    if (s.success) {
      JSON.parse(s.data).forEach((value: string) => {
        tableData.value.push({ dataBaseName: value });
      });
      if (alertM) {
        ElMessage({
          message: "刷新成功",
          type: "success",
        });
      }
    } else {
      ElMessage({
        message: s.message,
        type: "warning",
      });
    }
  });
}

function GetDataBaseListByRegNameD2() {
  let htmlElement = document.getElementById("dbName");
  if (htmlElement == null) {
    return
  }
  // 以 db: 开头， 所以是从第三个开始
  var dbName = htmlElement.innerHTML.substring(3, htmlElement.innerHTML.length);
  GetDataBaseListByRegNameD(true, dbName);
}

async function changecollapse() {
  if (collapse.value === "200px") {
    let size = 200;
    while (size > 0) {
      size -= 10;
      collapse.value = size + "px";
      await sleep(1);
    }
    isHidenDdatabase.value = true;
  } else {
    isHidenDdatabase.value = false;
    let size = 0;
    while (size < 200) {
      size += 10;
      collapse.value = size + "px";
      await sleep(1);
    }
  }
}
const tableRowClassName = ({
                             row,
                             rowIndex,
                           }: {
  row: string
  rowIndex: number
}) => {
  if (rowIndex % 2 == 0) {
    return "lyb-row";
  } else {
    return "";
  }

}

function sleep(ms: number | undefined) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

const rowContextmenu = (row: JSON, column: any, event: any
) => {
  console.log("右键 + " + JSON.stringify(row))
}

const tableData = ref([{ dataBaseName: "test" }]);
</script>

<script lang="ts">
export default {
  props: ["dbName"], // 接收父组件传来的参数
}
</script>

<style>
.lyb-row {
  background-color: rgba(156, 189, 227, 0.1) !important;
}

</style>