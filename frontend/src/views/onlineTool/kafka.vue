<script setup lang="ts">
import type { Ref } from "vue/dist/vue";
import {ElMessage} from "element-plus";
import {onMounted, ref} from 'vue'
import { Calendar, Search } from '@element-plus/icons-vue'
import {GetTopicAndPartition, GetKafkaForm, UpdateKafkaForm} from "../../../wailsjs/go/main/App";
import {index} from "../../../wailsjs/go/models";
import { Location } from "@element-plus/icons-vue";
import {useI18n} from "vue-i18n";
const textarea = ref('')
const {t, availableLocales: languages, locale} = useI18n();

const kafkaData = ref(Array<index.TopicAndPartition>());
const input2 = ref('')
const nowClickTopic = ref('')
const dialogFormVisible = ref(false);
let kafkaForm : Ref<index.KafkaForm>  = ref(new index.KafkaForm());
onMounted(() => {
  GetKafkaForm().then(function (s) {
    kafkaForm.value = s;
    if (kafkaForm.value.connection === "") {
      dialogFormVisible.value = true;
    }
  });
  GetTopicAndPartition().then(function (s) {
    kafkaData.value = s;
  });
});

function hide(topicName: string) {
  let value = input2.value;
  if (value == null || value == "") {
    return true;
  }
  return topicName.toLowerCase().indexOf(value.toLowerCase()) != -1;
}

function iconView(topicName: string) {
  let value = nowClickTopic.value;
  return value != null && value === topicName && hide(topicName);
}

function clickTopic(topicNameAndPartition: index.TopicAndPartition) {
  nowClickTopic.value = topicNameAndPartition.topicName;
}

function changeS() {
}

function submitForm() {
  let form = kafkaForm.value;
  if (form.connection === "") {
    ElMessage({
      message: "Kafka地址必填",
      type: "warning",
    });
    return;
  }
  UpdateKafkaForm(form).then(function (s) {
     if (s != "") {
       ElMessage({
         message: s,
         type: "warning",
       });
     } else {
       dialogFormVisible.value = false;
     }
  });
}


</script>
<template>

  <el-dialog v-model="dialogFormVisible" title="Kafka连接配置">

    <el-form
        label-position="right"
        label-width="100px"
        :model="kafkaForm"
        style="max-width: 600px"
    >
      <el-form-item label="Kafka地址" required="required" style="width: 600px">
        <el-input v-model="kafkaForm.connection" style="width: 600px"  />
      </el-form-item>
      <el-form-item>
        <el-button type="success" plain @click="submitForm">确定</el-button
        >
<!--        <el-button @click="resetForm(ruleFormRef)">校验连接</el-button>-->
      </el-form-item>
    </el-form>


  </el-dialog>


  <div class="common-layout" style="height: 100vh">
    <el-container>
      <el-aside width="200px">
        <el-input
            v-model="input2"
            class="w-50 m-2"
            :placeholder="t('globals.search')"
            :prefix-icon="Search"
            style="width: 175px"
            clearable
            @input="changeS"
        />


        <div v-for="(value, key) in kafkaData" :key="key" :id="value.topicName" style="margin-top: 10px;margin-left: 30px">
          <el-button  plain v-if="hide(value.topicName)" @click="clickTopic(value)">{{ value.topicName }}</el-button>
          <el-icon v-if="iconView(value.topicName)"><Location /></el-icon>
        </div>

      </el-aside>
      <el-main style="background-color: springgreen">Main</el-main>
    </el-container>
  </div>
</template>
