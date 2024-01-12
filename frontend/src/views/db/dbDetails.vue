<template>
  <el-empty description="数据库无法连接" v-if="allIsEmpty" />
  <el-affix position="bottom" :offset="20" v-if="isHidenDdatabase">
    <el-icon class="arrow-icon" @click="changecollapse"
             style="opacity: 0.4;position: absolute;cursor: grab; top: 50%; left: 0; transform: translateY(-50%); width: 20px;z-index: 99">
      <ArrowRightBold/>
    </el-icon>
  </el-affix>
  <div class="common-layout">
    <el-container>
      <el-aside width="10px" v-if="isHidenDdatabase" @click="changecollapse">
        <div class="arrow"
             style="width: 20px; height: 90vh; background-color: #a3e1ff; position: relative;opacity: 0.2;"></div>
      </el-aside>
      <el-main>
        <div class="common-layout">
          <el-container v-if="!allIsEmpty">
            <el-aside :width="collapse">
              <!--        这里是库名列表-->
              <el-table :data="tableData" style="width: 100%" height="90vh" :row-class-name="tableRowClassName"
                        @row-contextmenu="rowContextmenu"
                        @mouseleave="elCardDatabaseRightHidPopup" @row-click="dataBaseRowClick"
                        @cell-mouse-enter="dataBaseRowClickMouse">
                <el-table-column prop="dataBaseName">
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
              <el-main>



              </el-main>

            </el-container>
          </el-container>
        </div>


      </el-main>
    </el-container>
  </div>
  <el-card class="custom-card" id="el-card-database-right" shadow="hover" v-if="elCardDatabaseRightHid"
           :style="{ top: elCardDatabaseRightPosition.y + 'px', left: elCardDatabaseRightPosition.x + 'px' }"
           @mouseleave="elCardDatabaseRightHid=false;elCardDatabaseRightNowMouseInMenu=false"
           @mouseover="elCardDatabaseRightHid=true;elCardDatabaseRightNowMouseInMenu=true">
    <div class="el-card-database-right-header" @mouseleave="ssssss('s1')" @mousemove="sssss2('s1')" id="s1">
      <div class="el-card-database-right">
        <el-container @click="createDataBaseOpen">
          <el-aside width="30px">
            <el-icon style="top: 2px">
              <FolderAdd/>
            </el-icon>
          </el-aside>
          <el-aside width="120px">
            <span class="el-card-database-right-span1">
              {{ dbType === "MySQL" ? "创建数据库" : "创建模式" }}
            </span>
          </el-aside>
          <el-aside>
            <span class="el-card-database-right-span2">Ctrl+N</span>
          </el-aside>
        </el-container>
      </div>
      <div class="databaseBr"></div>
    </div>
    <div class="el-card-database-right-header"
         @click="showDatabaseDDL"
         @mouseleave="ssssss('s2')" @mousemove="sssss2('s2')" id="s2">
      <div class="el-card-database-right">
        <el-container>
          <el-aside width="30px">
            <el-icon style="top: 2px">
              <FullScreen/>
            </el-icon>
          </el-aside>
          <el-aside width="120px">
            <span class="el-card-database-right-span1">查看DDL</span>
          </el-aside>
        </el-container>
      </div>
      <div class="databaseBr"></div>
    </div>

    <div class="el-card-database-right-header" @mouseleave="ssssss('s3')" @mousemove="sssss2('s3')" id="s3"
      @click="databaseReNameClick"
    >
      <div class="el-card-database-right">
        <el-container>
          <el-aside width="30px">
            <el-icon style="top: 2px">
              <Edit/>
            </el-icon>
          </el-aside>
          <el-aside width="120px">
            <span class="el-card-database-right-span1">重命名</span>
          </el-aside>
          <el-aside>
            <span class="el-card-database-right-span2">M</span>
          </el-aside>
        </el-container>
      </div>
      <div class="databaseBr"></div>
    </div>
    <div class="el-card-database-right-header" @mouseleave="ssssss('s4')" @mousemove="sssss2('s4')" id="s4">
      <div class="el-card-database-right">
        <el-container>
          <el-aside width="30px">
            <el-icon style="top: 2px">
              <Delete/>
            </el-icon>
          </el-aside>
          <el-aside width="120px">
            <span class="el-card-database-right-span1">删除</span>
          </el-aside>
          <el-aside>
            <span class="el-card-database-right-span2">Delete</span>
          </el-aside>
        </el-container>
      </div>
    </div>
  </el-card>

<!-- 创建数据库的弹窗-->
  <el-dialog
      v-model="createDataBase"
      title="创建"
      width="30%"
  >
<!--    创建数据库  dataBaseCreateForm-->
    <el-form
        label-position="right"
        label-width="100px"
        :model="dataBaseCreateForm"
        style="max-width: 460px"
    >
      <el-form-item label="数据库名">
        <el-input v-model="dataBaseCreateForm.dataBaseName" />
      </el-form-item>
      <el-form-item label="字符集">
        <el-select
            v-model="dataBaseCreateForm.characterSet"
            class="m-2"
            placeholder="Select"
            style="width: 240px"
            @change="dbcharacterSetOPtChange"
        >
          <el-option
              v-for="item in Object.keys(mysqlChartSet)"
              :key="item"
              :label="item"
              :value="item"
          />
        </el-select>

      </el-form-item>
      <el-form-item label="排序规则">
        <el-select
            v-model="dataBaseCreateForm.collation"
            class="m-2"
            placeholder="Select"
            style="width: 240px"
            default="utf8"
        >
          <el-option
              v-for="item in dbCollationOPt"
              :key="item"
              :label="item"
              :value="item"
          />
        </el-select>
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="createDataBase = false">取消</el-button>
        <el-button type="primary" @click="CreateDatabaseFun">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>

  <!--SQL的弹窗-->
  <el-dialog
      v-model="SqlDialog"
      :title="SqlDialogTitle"
      width="30%"
  >
    <span>{{ ddlData }}</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="SqlDialog = false">取消</el-button>
      </span>
    </template>
  </el-dialog>

  <!--数据库重命名-->
  <el-dialog
      v-model="databaseReName"
      title="重命名"
      width="30%"
  >
    <el-input
        v-model="databaseReNameNowName"
        class="w-50 m-2"
        id="databaseReName"
        :prefix-icon="Edit"
    />
    <template #footer>
      <span class="dialog-footer">
       <el-button @click="databaseReName = false">取消</el-button>
        <el-button type="primary" @click="databaseReName = false">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>


</template>


<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {ArrowLeft, ArrowRightBold, Delete, Edit, FolderAdd, FullScreen, Refresh} from "@element-plus/icons-vue";
import {GetDataBaseListByRegName, PingDbWithName, GetDataBaseType, GetAllCharset, CreateDatabase, GetDataBaseDDL} from "../../../wailsjs/go/main/App";
import {info} from "../../../wailsjs/go/models";
import {ElMessage} from "element-plus";

const collapse = ref("200px");
// 是否连接不上，若是连接不上则显示空状态
const allIsEmpty = ref(true);
// 当前数据库的内容是否隐藏
const isHidenDdatabase = ref(false);
// 当前数据库的内容隐藏后，显示的左侧条，是否显示icon
const isHidenDdatabaseIcon = ref(false);
// 创建数据库的弹窗是否显示
const createDataBase = ref(false);
function createDataBaseOpen() {
  createDataBase.value = true;
  if (dbType.value == "MySQL") {
    // 默认的字符集和排序规则
    dataBaseCreateForm.value.characterSet = "utf8";
    dataBaseCreateForm.value.collation = "utf8_general_ci";
    dbCollationOPt.value = mysqlChartSet.value[dataBaseCreateForm.value.characterSet];
  }
}

const dbName = ref("");
const dbType = ref("");
onMounted(() => {
  if (props == undefined || props.GetNowTabName == undefined) {
    return;
  }
  let getNowTabName = props.GetNowTabName();
  dbName.value = getNowTabName.substring(3, getNowTabName.length);
  // 以 db: 开头， 所以是从第三个开始
  PingDbWithName(dbName.value).then((s) => {
    allIsEmpty.value = !s;
    if (s) {
      // 得到当前的数据库类型
      GetDataBaseType(dbName.value).then((s) => {
        dbType.value = s;
        // todo 根据不同的数据库类型，做不同的操作
        if (dbType.value == "MySQL") {
          // 得到所有的字符集和排序规则
          GetAllCharset(dbName.value, "MySQL").then((res) => {
            mysqlChartSet.value = res;
          });
        }
      });
      // 得到所有数据库
      GetDataBaseListByRegNameD(false, dbName.value);
    } else {
      ElMessage({
        message: "数据库无法连接",
        type: "warning",
      });
    }
  });
  // updateTableDataInDb();
});

// eslint-disable-next-line vue/valid-define-props
const props = defineProps({
  // 这是父组件传递过来的函数
  GetNowTabName: {
    type: Function
  }
})

function GetDataBaseListByRegNameD(alertM: boolean, dbName: string) {
  if (props == undefined || props.GetNowTabName == null) {
    return;
  }
  GetDataBaseListByRegName(dbName).then((s) => {
    tableData.value = [];
    if (s.success) {
      JSON.parse(s.data).forEach((value: string) => {
        tableData.value.push({dataBaseName: value});
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
  GetDataBaseListByRegNameD(true, dbName.value);
}

// 隐藏或者回显数据库栏
async function changecollapse() {
  // 右键菜单隐藏
  elCardDatabaseRightHid.value = false;
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
  row: JSON
  rowIndex: number
}) => {
  if ((row as any)["dataBaseName"] == nowSelectDataBaseName.value) {
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
  nowSelectDataBaseName.value = (row as any)["dataBaseName"];
  console.log((row as any)["dataBaseName"]);
  elCardDatabaseRightHid.value = false;
  elCardDatabaseRightNowMouseInMenu.value = true;
  elCardDatabaseRightPosition.value.x = event.clientX;
  elCardDatabaseRightPosition.value.y = event.clientY;
  elCardDatabaseRightHid.value = true;
}

const tableData = ref([{dataBaseName: "test"}]);

// 当前选择的数据库
const nowSelectDataBaseName = ref('');
const showPopover = ref(false);
// 数据库右键菜单的位置
let elCardDatabaseRightPosition = ref({x: 300, y: 300});
// 数据库右键菜单的显示与隐藏
let elCardDatabaseRightHid = ref(false);
// 当前鼠标是否在右键菜单上
let elCardDatabaseRightNowMouseInMenu = ref(false);

function elCardDatabaseRightHidPopup() {
  // 这个操作必不可少
  elCardDatabaseRightNowMouseInMenu.value = false;
  setTimeout(() => {
    if (elCardDatabaseRightNowMouseInMenu.value) {
      return;
    }
    elCardDatabaseRightHid.value = false;
  }, 100);
}

// 单击一个数据库
function dataBaseRowClick(row: JSON, column: any, event: any) {
  nowSelectDataBaseName.value = (row as any)["dataBaseName"];

}

// 注册新的数据库的表单
const dataBaseCreateForm = ref(new info.DataBaseCreateForm());
const mysqlChartSet = ref<{ [key: string]: Array<string> }>({});
const dbCollationOPt = ref(new Array<string>());
// 新建数据库,选择字符集后的操作,需要修改
function dbcharacterSetOPtChange() {
  dbCollationOPt.value = mysqlChartSet.value[dataBaseCreateForm.value.characterSet];
  dataBaseCreateForm.value.collation = "";
}
// 新建数据库,选择字符集后的操作,需要修改
function CreateDatabaseFun() {
  //  todo 表单校验
  CreateDatabase(dbName.value, dbType.value, dataBaseCreateForm.value).then((s) => {
    if (s.success) {
        ElMessage({
        message: "创建成功",
        type: "success",
      });
      GetDataBaseListByRegNameD2();
      createDataBase.value = false
    } else {
      ElMessage({
        message: s.message,
        type: "warning",
      });
    }
  });
}

let ddlData = ref("");
// 查看数据库DDL
function showDatabaseDDL() {
  GetDataBaseDDL(dbType.value, dbName.value, nowSelectDataBaseName.value).then((s) => {
    if (s != '') {
      ddlData.value = s;
      SqlDialogTitle.value = "DDL"
      SqlDialog.value = true;
    } else {
      ElMessage({
        message: "获取DDL失败",
        type: "warning",
      });
    }
  });
}

let SqlDialog = ref(false);
let SqlDialogTitle = ref("");
// 数据库重命名的弹窗
let databaseReName = ref(false);
let databaseReNameNowName = ref("");
function databaseReNameClick() {
  databaseReName.value = true;
  databaseReNameNowName.value = nowSelectDataBaseName.value;
  // 鼠标光标定位到输入框
  // 通过元素ID获取输入框
  // 将鼠标光标定位到输入框
  setTimeout(() => {
    const inputElement = document.getElementById("databaseReName");
    if (inputElement != undefined) {
      inputElement.focus();
    }
  }, 100);
}

// 鼠标进入一个数据库行
function dataBaseRowClickMouse(row: JSON, column: any, cell: any, event: any) {

}

function sssss2(idStr: string) {
  const element = document.getElementById(idStr);
  if (element == undefined) {
    return;
  }
  element.style.backgroundColor = "#C7DEF7";
}

function ssssss(idStr: string) {
  const element = document.getElementById(idStr);
  if (element == undefined) {
    return;
  }
  element.style.backgroundColor = '';
}

</script>


<style>
.lyb-row {
  background-color: #f5f7fa !important;
  /*pointer-events: none !important;*/
}

.el-card-database-right {
  padding: 10px;
  /*text-align: right;*/
}

.el-card {
  --el-card-padding: 0px !important;
}

.el-card-database-icon {
  top: 2px !important;
}

.el-card-database-right-span1 {
  font-size: 12px;
}

.el-card-database-right-span2 {
  font-size: 11px;
  color: #909399;
}

#el-card-database-right {
  width: 250px;
  position: fixed;
  z-index: 999;
  transition: none
}

#el-card-database-right :hover {
  cursor: pointer;
}

.databaseBr {
  height: 1px;
  background-color: #b6b6b6
}

</style>