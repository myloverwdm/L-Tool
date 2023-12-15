<script setup lang="ts">
import {AddPath, DeleteZk, InitZk, ReadZkNode, SetForm} from "../../../wailsjs/go/main/App";
import {onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";
import {ElLoading, ElMessage} from 'element-plus'
import type {Ref} from 'vue/dist/vue'
import {index} from "../../../wailsjs/go/models";

const {t, availableLocales: languages, locale} = useI18n();

interface DataItem {
  children: Ref<DataItem[]>;
  id: any;
  label: any;
}

const data: Ref<DataItem[]> = ref([]);

let zkData = ref("");
let nowNode = ref("");
let tatData : Ref<index.Stat>  = ref(new index.Stat());
let drawer2 = ref(false);
const drawer3 = ref(false);
let textareaData = ref("");

let newNodeName = ref("");
let newNodeValue = ref("");

let sss = ref("")
let zkForm: index.ZkForm = {
  address: "",
  zkData: null as unknown as index.ZkData,
  convertValues: function (a: any, classs: any, asMap?: boolean) {
    throw new Error("Function not implemented.");
  }
};
let nowSelectPath = "";
let nowNodeObject: { children: { id: number; label: string }[] } | null = null;
onMounted(() => {
  initDo();
});

function initDo() {
  data.value = [];
  InitZk().then(function (s) {
    if (s.address === undefined || s.address === "") {
      drawer2.value = true;
      return;
    }
    zkForm.address = s.address;
    sss.value = s.address;
    if (s.zkData === null || s.zkData.nextPath === null) {
      return;
    }
    s.zkData.nextPath.forEach((value, index) => {
      data.value.push({
        id: index as never,
        label: value,
        children: null as unknown as Ref<DataItem[]>,
      });
    });
  });
}

function handleNodeClick(
    nodeData: { children: { id: number; label: string }[] } | null
) {
  nowNodeObject = nodeData;
  nowNode.value = JSON.parse(JSON.stringify(nodeData))["label"];
  const path = getNodePath(nodeData);
  nowSelectPath = path;
  let readZkNode1 = ReadZkNode(path);
  readZkNode1.then((s) => {
    zkData.value = s.data;
    // eslint-disable-next-line no-undef
    tatData.value = s.statData;
    let child: { id: number; label: string }[] = [];
    s.nextPath.forEach((value, index) => {
      child.push({
        id: index,
        label: value,
      });
    });
    if (nodeData != null) {
      nodeData.children = child;
    }

    textareaData.value = s.data;
  });
}

function getNodePath(node: any) {
  const path1: any[] = [];
  findNodePath(data, node, path1);
  return "/" + path1.join("/");
}

function findNodePath(nodes: Ref<DataItem[]>, targetNode: any, path: any[]) {
  if (Array.isArray(nodes)) {
    for (const node of nodes) {
      if (node === targetNode) {
        path.unshift(node.label);
        return true;
      }
      if (node.children && node.children.length > 0) {
        if (findNodePath(node.children, targetNode, path)) {
          path.unshift(node.label);
          return true;
        }
      }
    }
  } else if (nodes.value && Array.isArray(nodes.value)) {
    for (const node of nodes.value) {
      if (node === targetNode) {
        path.unshift(node.label);
        return true;
      }
      if (node.children && JSON.stringify(node.children).length > 2) {
        if (findNodePath(node.children, targetNode, path)) {
          path.unshift(node.label);
          return true;
        }
      }
    }
  }
  return false;
}

function setForm() {
  let loading = ElLoading.service({});
  zkForm.address = sss.value;
  SetForm(zkForm);
  ElMessage({
    message: "保存成功,开始初始化数据",
    type: "success",
  });
  initDo();
  drawer2.value = false;
  loading.close();
}

function addNewNode() {
  if (newNodeName.value === "") {
    ElMessage({
      message: t("zk.p-node-name"),
      type: "warning",
    });
    return;
  }
  let loading = ElLoading.service({});
  AddPath(nowSelectPath, newNodeName.value, newNodeValue.value);
  ElMessage({
    message: t("globals.add-success"),
    type: "success",
  });
  drawer3.value = false;
  newNodeName.value = "";
  newNodeValue.value = "";
  loading.close();
  handleNodeClick(nowNodeObject)
}

function jsonParser() {
  try {
    textareaData.value = JSON.stringify(JSON.parse(zkData.value), null, 4);
    // eslint-disable-next-line no-empty
  } catch (e) {}
}


function deleteZk() {
  if (nowSelectPath === "") {
    ElMessage({
      message: t("zk.p-node-delete"),
      type: "warning",
    });
    return;
  }
  let loading = ElLoading.service({});
  DeleteZk(nowSelectPath).then((s) => {
    if (s == null || s == "") {
      ElMessage({
        message: t("globals.delete-success"),
        type: "success",
      });
      // 当前节点删除
      deleteNowNode(data.value);
      nowSelectPath = "";
      nowNodeObject = null;

    } else {
      ElMessage.error(t("globals.delete-fail") + ":" + s);
    }
  });
  loading.close();
}

function deleteNowNode(data: any[], parentPath = '') {
  for (let i = 0; i < data.length; i++) {
    const node = data[i];
    const path = parentPath + node.label;
    if ("/" + path === nowSelectPath) {
      data.splice(i, 1);
      return;
    }
    if (node.children && node.children.length > 0) {
      deleteNowNode(node.children, path + '/');
    }
  }
}


function drawer3Open() {
  if (nowSelectPath === "") {
    ElMessage({
      message: t("zk.p-node-add"),
      type: "warning",
    });
    return;
  }
  drawer3.value = true;
}
</script>
<template>
  <el-drawer v-model="drawer2">
    <el-form label-width="100px" :model="zkForm" style="max-width: 460px">
      <el-form-item :label="t('zk.host')">
        <el-input v-model="sss" />
      </el-form-item>
      <!--      <el-form-item label="Activity zone">-->
      <!--        <el-input v-model="zkForm.region"/>-->
      <!--      </el-form-item>-->
      <!--      <el-form-item label="Activity form">-->
      <!--        <el-input v-model="zkForm.type"/>-->
      <!--      </el-form-item>-->
      <el-button :onclick="setForm">{{ t("globals.submit") }}</el-button>
    </el-form>
  </el-drawer>

  <el-drawer v-model="drawer3">
    <el-form label-width="100px" :model="zkForm" style="max-width: 460px">
      <el-form-item :label="t('zk.node-name')">
        <el-input v-model="newNodeName"/>
      </el-form-item>
      <el-form-item :label="t('zk.node-value')">
        <el-input v-model="newNodeValue"/>
      </el-form-item>
      <el-button :onclick="addNewNode">{{ t("globals.add") }}</el-button>
    </el-form>
  </el-drawer>

  <div class="common-layout">
    <el-container>
      <el-container>
        <el-aside width="300px">
          <el-tree
              :data="data"
              show-icon
              highlight-current
              @node-click="handleNodeClick"
          ></el-tree>
        </el-aside>
        <el-main>

          <div class="common-layout">

            <el-row :gutter="20">
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">{{ t('zk.node-name2') }}</el-text>
                <el-text class="mx-1" type="success">{{
                    nowNode
                  }}
                </el-text>
              </el-col>
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">{{ t('zk.create-time') }}</el-text>
                <el-text class="mx-1" type="success"
                  >{{
                    !tatData ? "" : tatData.ctime == null ? 0 : tatData.ctime
                  }}
                </el-text>
              </el-col>
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">{{ t('zk.data-length') }}</el-text>
                <el-text class="mx-1" type="success">
                  {{ !tatData ? "" : tatData.dataLength == null ? 0 : tatData.dataLength }}
                </el-text>
              </el-col>
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">{{ t('zk.version') }}</el-text>
                <el-text class="mx-1" type="success">{{
                    !tatData ? "" : tatData.version == null ? 0 : tatData.version
                  }}
                </el-text>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">{{ t('zk.modify-time') }}</el-text>
                <el-text class="mx-1" type="success">{{
                    !tatData ? "" : tatData.mtime == null ? 0 : tatData.mtime
                  }}
                </el-text>
              </el-col>
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">{{ t('zk.child-count') }}</el-text>
                <el-text class="mx-1" type="success">
                  {{ !tatData ? "" : tatData.numChildren == null ? 0 : tatData.numChildren }}
                </el-text>
              </el-col>
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">pzxid：</el-text>
                <el-text class="mx-1" type="success">{{
                    !tatData ? "" : tatData.pzxid == null ? 0 : tatData.pzxid
                  }}
                </el-text>
              </el-col>
              <el-col :span="6">
                <div class="grid-content ep-bg-purple"/>
                <el-text class="mx-1" type="success">cversion：</el-text>
                <el-text class="mx-1" type="success">{{
                    !tatData ? "" : tatData.cversion == null ? 0 : tatData.cversion
                  }}
                </el-text>
              </el-col>
            </el-row>
          </div>

          <el-input
              v-model="textareaData"
              rows="20"
              type="textarea"
              readonly
              resize="none"
              show-word-limit
              maxlength="10000000"
              style="white-space: pre-wrap;"
          />
          <el-header></el-header>
          <el-row :gutter="20">
            <el-col :span="6">
              <div class="grid-content ep-bg-purple"/>
              <el-button @click="drawer2 = true"
              >{{ t("globals.setting") }}
              </el-button>
            </el-col>
            <el-col :span="6">
              <div class="grid-content ep-bg-purple"/>
              <el-button @click="jsonParser"
              >{{ t("tool.json-parse") }}
              </el-button>
            </el-col>
            <el-col :span="6">
              <div class="grid-content ep-bg-purple"/>

              <el-popconfirm :title="t('tool-timestamp.delete-node-desc')" confirm-button-type=""
                             @confirm="deleteZk"
              >
                <template #reference>
                  <el-button>{{ t("tool-timestamp.delete-node") }}</el-button>
                </template>
              </el-popconfirm>
            </el-col>
            <el-col :span="6">
              <div class="grid-content ep-bg-purple"/>
              <el-button @click="drawer3Open"
              >{{ t("tool-timestamp.add-node") }}
              </el-button>
            </el-col>
          </el-row>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
