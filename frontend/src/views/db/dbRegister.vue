<template>

  <el-row :gutter="20">
    <el-col :span="10">
      <div class="grid-content ep-bg-purple"/>
    </el-col>
    <el-col :span="7">
      <div class="grid-content ep-bg-purple"/>
    </el-col>
    <el-col :span="4">
      <div class="grid-content ep-bg-purple"/>
      <el-input
          v-model="searchValue"
          :placeholder="t('globals.please-input')"
          class="input-with-select"
          :width="100"
          clearable
          @clear="updateTableDataInDb"
      >
        <template #append>
          <el-button @click="updateTableDataInDb" :icon="Search"/>
        </template>
      </el-input>
    </el-col>
    <el-col :span="1">
      <div class="grid-content ep-bg-purple"/>
      <el-button type="primary" @click="reg">{{ t("globals.reg") }}</el-button>
    </el-col>
    <el-col :span="2">
      <div class="grid-content ep-bg-purple"/>
      <el-button type="primary" :icon="Refresh" @click="updateTableDataInDb"/>
    </el-col>

  </el-row>

  <el-divider content-position="center">T^T</el-divider>
  <el-table :data="tableDataCopy" style="width: 100%" :empty-text="t('globals.empty-text')">

    <el-table-column fixed prop="name" :label="t('db.name')"/>
    <el-table-column prop="dbType" :label="t('db.dbType')"/>
    <el-table-column prop="address" :label="t('db.address')"/>
    <el-table-column prop="port" :label="t('db.port')"/>
    <el-table-column prop="regTimeStr" :label="t('db.regTime')"/>
    <el-table-column prop="updateTimeStr" :label="t('db.updateTime')"/>
    <el-table-column fixed="right" :label="t('globals.operations')">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="updateDb(scope.row)">{{ t("globals.edit") }}</el-button>
        <el-button link type="primary" size="small" @click="deleteClick(scope.row)">{{
            t("globals.delete")
          }}
        </el-button>
      </template>
    </el-table-column>
  </el-table>


  <!--  注册或编辑-->
  <el-dialog v-model="regOrUpdate" :title="editTitle">
    <el-form :model="nowDataBaseInfo" label-width="120px">
      <el-form-item :label="t('db.name')">
        <el-input v-model="nowDataBaseInfo.name" :disabled="isUpdate"/>
      </el-form-item>
      <el-form-item :label="t('db.dbType')">
        <el-select v-model="nowDataBaseInfo.dbType" @change="changeDbType">
          <el-option label="MySQL" value="MySQL"/>
          <el-option label="Oracle" value="Oracle"/>
        </el-select>
      </el-form-item>
      <el-form-item :label="t('db.ip')">

        <div class="common-layout">
          <el-container>
            <el-aside>
              <el-input v-model="nowDataBaseInfo.address"/>
            </el-aside>
            <span>&nbsp;&nbsp;&nbsp;{{ t("db.port") }}：&nbsp;</span>
            <el-aside width="200px">
              <el-input v-model="nowDataBaseInfo.port" @change="changePort" type="number" min="0" max="65535"/>
            </el-aside>

          </el-container>
        </div>

      </el-form-item>
      <el-form-item :label="t('db.username')">
        <el-input v-model="nowDataBaseInfo.userName"/>
      </el-form-item>

      <el-form-item :label="t('db.password')">
        <el-input
            v-model="nowDataBaseInfo.password"
            type="password"
            show-password
            :readonly="!canEditPassword"
        >
          <template #append v-if="!canEditPassword">
            <el-button :icon="Edit " @click="showPasswordT"/>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary">{{ t("db.ping") }}</el-button>
        <el-button type="success" @click="addOrUpdate">{{ t("globals.submit") }}</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

</template>

<script lang="ts" setup>
import {Edit, Refresh, Search} from "@element-plus/icons-vue";
import {onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";
import {AddOrUpdateDataBaseInfo, GetAllDataBaseInfo, DeleteOneDataBaseInfo} from "../../../wailsjs/go/main/App";
import {db} from "../../../wailsjs/go/models";
import {ElMessage, ElMessageBox} from "element-plus";

const {t, availableLocales: languages, locale} = useI18n();
const searchValue = ref("");
const regOrUpdate = ref(false);
const canEditPassword = ref(false);
const nowDataBaseInfo = ref(new db.DataBaseInfo());
const editTitle = ref(t("globals.edit"));
const isUpdate = ref(false);

onMounted(() => {
  updateTableDataInDb();
});

const tableData = ref(Array<db.DataBaseInfo>());
const tableDataCopy = ref(Array<db.DataBaseInfo>());


function updateTableDataInDb() {
  GetAllDataBaseInfo().then((s) => {
    tableData.value = s;
    updateTableData();
  });
}

function updateTableData() {
  tableDataCopy.value = []
  for (let i = 0; i < tableData.value.length; i++) {
    if (hiden(tableData.value[i])) {
      let parse = JSON.parse(JSON.stringify(tableData.value[i]));
      console.log(parse)
      // 时间格式化
      parse.regTimeStr = dateFormat(parse.regTime);
      parse.updateTimeStr = dateFormat(parse.updateTime);
      tableDataCopy.value.push(parse);
    }
  }
}

function dateFormat(time: number) {
  var currentDate = new Date(time);
  var year = currentDate.getFullYear();
  var month = formatTwoDigits(currentDate.getMonth() + 1);
  var day = formatTwoDigits(currentDate.getDate());
  var hours = formatTwoDigits(currentDate.getHours());
  var minutes = formatTwoDigits(currentDate.getMinutes());
  var seconds = formatTwoDigits(currentDate.getSeconds());
  var formattedDate = year + '-' + month + '-' + day + ' ' + hours + ':' + minutes + ':' + seconds;
  return formattedDate;
}

function formatTwoDigits(number) {
  return number < 10 ? '0' + number : number;
}


function hiden(row: db.DataBaseInfo) {
  if (searchValue.value === '') {
    return true;
  }
  return (
      row.name.includes(searchValue.value) ||
      row.address.includes(searchValue.value) ||
      row.dbType.includes(searchValue.value) ||
      row.port.includes(searchValue.value));
}

function reg() {
  nowDataBaseInfo.value = new db.DataBaseInfo();
  nowDataBaseInfo.value.dbType = "MySQL";
  nowDataBaseInfo.value.port = "3306";
  canEditPassword.value = true;
  editTitle.value = t("globals.reg");
  isUpdate.value = false;
  regOrUpdate.value = true;
}

const updateDb = (dataBaseInfo: db.DataBaseInfo) => {
  nowDataBaseInfo.value = JSON.parse(JSON.stringify(dataBaseInfo));
  canEditPassword.value = false;
  editTitle.value = t("globals.edit");
  isUpdate.value = true;
  regOrUpdate.value = true;
}

function showPasswordT() {
  nowDataBaseInfo.value.password = "";
  canEditPassword.value = true;
}

function changeDbType(value: any) {
  if (nowDataBaseInfo.value.dbType === "MySQL") {
    nowDataBaseInfo.value.port = "3306";
  } else if (nowDataBaseInfo.value.dbType === "Oracle") {
    nowDataBaseInfo.value.port = "1521";
  }
  // todo 修改其它类型的数据库的默认端口号
}

function changePort() {
  let port = parseInt(nowDataBaseInfo.value.port);
  if (port < 0 || port > 65535) {
    changeDbType(null);
  }
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
function addOrUpdate() {
  console.log(JSON.stringify(nowDataBaseInfo.value))
  AddOrUpdateDataBaseInfo(nowDataBaseInfo.value, !isUpdate.value).then((s) => {
    if (s != "") {
      ElMessage({
        message: t(s),
        type: "warning",
      });
      return;
    }
    ElMessage({
      message: t("globals.operations-success"),
      type: "success",
    });
    updateTableDataInDb();
    regOrUpdate.value = false;
  });
}

const deleteClick = (dataBaseInfo: db.DataBaseInfo) => {
  ElMessageBox.confirm(
      t('globals.sure-delete'),
      t('globals.warning'),
      {
        confirmButtonText: t('globals.confirm'),
        cancelButtonText: t('globals.cancel'),
        type: 'warning',
      }
  )
      .then(() => {
        DeleteOneDataBaseInfo(dataBaseInfo.name).then((s) => {
          if (s != "") {
            ElMessage({
              type: 'warning',
              message: t(s),
            })
            return false
          }
          updateTableDataInDb();
          ElMessage({
            type: 'success',
            message: t('globals.delete-success'),
          })
        })

      })
      .catch(() => {

      })
}

</script>
