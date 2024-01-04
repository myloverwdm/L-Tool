<script setup lang="ts">
import {onMounted, ref} from "vue";
import {ElMessage} from "element-plus";
import {JsonToXml,} from "../../../wailsjs/go/main/App";
import Codemirror from "codemirror-editor-vue3";
// import "codemirror-editor-vue3/node_modules/codemirror/lib/codemirror.css";
// import "codemirror-editor-vue3/node_modules/codemirror/theme/material.css";
import "codemirror-editor-vue3/node_modules/codemirror/mode/javascript/javascript.js";
import "codemirror-editor-vue3/node_modules/codemirror/mode/css/css.js";
import "codemirror-editor-vue3/node_modules/codemirror/mode/xml/xml.js";
import "codemirror-editor-vue3/node_modules/codemirror/mode/shell/shell.js";
import "codemirror-editor-vue3/node_modules/codemirror/mode/sql/sql.js";
//代码补全提示
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/anyword-hint.js";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/css-hint.js";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/html-hint.js";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/javascript-hint.js";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/show-hint.css";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/show-hint.js";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/sql-hint.js";
import "codemirror-editor-vue3/node_modules/codemirror/addon/hint/xml-hint.js";
import {
  UpdateFileCache,
  GetFileCache,
} from "../../../wailsjs/go/main/App";
import {useI18n} from "vue-i18n";
import {Close} from "@element-plus/icons-vue";
const {t, availableLocales: languages, locale} = useI18n();

const dialogTableVisible = ref(false);
const xmlData = ref("");
let drawerDiff = ref(false);
let drawerRtlDiff = ref(false);
let emptyHideen = ref(false);
function foramtJson() {
  if (code.value == "" || code.value.trim() == "") {
    return;
  }
  try {
    code.value = JSON.stringify(JSON.parse(code.value), null, 4);
  } catch (e) {
    ElMessage({
      message: t("tool-json.error-json"),
      type: "warning",
    });
  }
}

function foramtJsonTo1() {
  if (code.value == "" || code.value.trim() == "") {
    return;
  }
  try {
    code.value = JSON.stringify(JSON.parse(code.value), null).replace(
        /\s+/g,
        ""
    );
  } catch (e) {
    ElMessage({
      message: t("tool-json.error-json"),
      type: "warning",
    });
  }
}

function quzhuanyi() {
  let a = '"' + code.value + '"';
  try {
    code.value = code.value
        .replace(/\\"/g, '"')
        .replace(/\\n/g, "n")
        .replace(/\\t/g, "t")
        .replace(/\\r/g, "r")
        .replace(/\\\\/g, "\\");
  } catch (e) {
    ElMessage({
      message: t("tool-json.cant-parse"),
      type: "warning",
    });
  }
}

function zhuanyi() {
  code.value = code.value.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
}

function copyXmlContent() {
  const input = document.createElement("textarea");
  input.value = xmlData.value;
  document.body.appendChild(input);
  input.select();
  document.execCommand("copy");
  document.body.removeChild(input);
  ElMessage({
    message: t("globals.copy-success"),
    type: "success",
  });
}

function jsonToXmls() {
  try {
    JsonToXml(code.value).then((value) => {
      if (value.success) {
        xmlData.value = value.data;
        dialogTableVisible.value = true;
      } else {
        ElMessage({
          message: value.errorMsg,
          type: "warning",
        });
      }
    });
  } catch (e) {
    ElMessage({
      message: t("tool-json.error-json"),
      type: "warning",
    });
  }
}

let code = ref("");
let codeChange = ref(false);
const jsonFileName = "cache/jsonCache";
onMounted(() => {
  setTimeout(() => {
    // 得到上次缓存的json
    GetFileCache(jsonFileName).then((s) => {
      code.value = s;
      // 需要延时执行
    });
  }, 100);
});

function inputS() {
  if (!codeChange.value) {
    codeChange.value = true;
    setTimeout(() => {
      UpdateFileCache(jsonFileName, code.value);
    }, 1000);
    codeChange.value = false;
  }
}

let diffRight = ref("");
let cmOptions = ref({
  // mode: ["text/javascript", "text/xml"], // Language mode
  mode: "text/javascript",
  autoRefresh: true,
  theme: "material", // Theme
  lineNumbers: true, // Show line number
  smartIndent: true, // Smart indent
  indentUnit: 2, // The smart indent unit is 2 spaces in length
  foldGutter: true, // Code folding
  styleActiveLine: true, // Display the style of the selected row
});

function hiddenS() {
  emptyHideen.value = code.value == "" && diffRight.value == "";
  drawerDiff.value = true;
}

</script>
<template>

  <el-drawer
      v-model="drawerDiff"
      direction="ttb"
      :with-header="false"
      class="LLLLL"
  >
    <el-affix :offset="15" >
      <el-icon class="el-iconLYB2" @click="drawerDiff = false"><Close/></el-icon>
    </el-affix>
    <el-empty v-if="emptyHideen" :description="t('tool-json.isEmpty')" />
    <code-diff
        :old-string="code"
        :new-string="diffRight"
        language="javascript"
        output-format="side-by-side"
        :editable="true"
    ></code-diff>
  </el-drawer>

  <el-drawer
      v-model="drawerRtlDiff"
      direction="rtl"
      :with-header="false"
  >

    <el-row :gutter="340">
      <el-col :span="340"><div class="grid-content ep-bg-purple" />
        <el-button @click="hiddenS" plain>{{ t("tool-json.duibi") }}</el-button>
      </el-col>
      <el-col :span="340"><div class="grid-content ep-bg-purple" />
        <el-icon class="el-iconLYB" @click="drawerRtlDiff = false"><Close/></el-icon>
      </el-col>
    </el-row>


    <br>
    <el-input
        v-model="diffRight"
        :rows="32"
        type="textarea"
        :placeholder="t('tool-json.pleaInput')"
    />
  </el-drawer>


  <div class="common-layout">
    <el-container>
      <el-header>
        <el-button @click="foramtJson" plain>{{
            t("tool-json.json-parse")
          }}
        </el-button>
        <el-button @click="foramtJsonTo1" plain>{{
            t("tool-json.json-one")
          }}
        </el-button>
        <el-button @click="quzhuanyi" plain>{{
            t("tool-json.quzhuan")
          }}
        </el-button>
        <el-button @click="zhuanyi" plain>{{ t("tool-json.zhuan") }}</el-button>

        <el-button @click="drawerRtlDiff = true" plain>{{ t("tool-json.duibi") }}</el-button>
      </el-header>
    </el-container>
  </div>

  <Codemirror
      v-model:value="code"
      :options="cmOptions"
      border
      placeholder="test placeholder"
      :height="760"
      @input="inputS"
  />
  <div style="font-size: 15px">
    <span>{{ t("tool-json.str-length") }}</span><span >{{ code.length }}</span>
  </div>
</template>

<style>

center {
  /*// 解决vue-code-diff对不齐和显示下拉标志问题*/
  max-height: 600px;
  overflow-y: auto;
  overflow-x: hidden;
}
/* 样式穿透-起始行左右对齐，*/
.d2h-code-side-line {
  height: 15px;
}

.LLLLL {
  height: 85vh !important;;
}

.el-iconLYB {
  cursor: pointer;
  top: 5px;
}
.el-iconLYB2 {
  cursor: pointer;
  /*top: 5px;*/
}


code.hljs {
  padding: 0;
}

/*// 删除行统计显示*/
.d2h-code-side-linenumber {
  display: none;
}

.d2h-code-side-line {
  padding: unset;
}

.d2h-code-line-ctn {
  width: unset;
}

/*// 删除第一行的统计结果*/
.file-info {
  display: none !important;
}
.file-header {
  display: none !important;;
}

/*
  Name:       material
  Author:     Mattia Astorino (http://github.com/equinusocio)
  Website:    https://material-theme.site/
*/
.CodeMirror,
.CodeMirror-wrap {
  background-color: white;
}

.cm-s-material.CodeMirror {
  background-color: #ffffff;
  /*color: #EEFFFF;*/
}

.cm-s-material .CodeMirror-gutters {
  background: #ffffff;
  color: #546e7a;
  border: none;
}

.cm-s-material .CodeMirror-guttermarker,
.cm-s-material .CodeMirror-guttermarker-subtle,
.cm-s-material .CodeMirror-linenumber {
  color: #546e7a;
}

.cm-s-material .CodeMirror-cursor {
  border-left: 1px solid #ffcc00;
}

.cm-s-material.cm-fat-cursor .CodeMirror-cursor {
  background-color: #5d6d5c80 !important;
}

.cm-s-material .cm-animate-fat-cursor {
  background-color: #5d6d5c80 !important;
}

.cm-s-material div.CodeMirror-selected {
  background: rgba(128, 203, 196, 0.2);
}

.cm-s-material.CodeMirror-focused div.CodeMirror-selected {
  background: rgba(128, 203, 196, 0.2);
}

.cm-s-material .CodeMirror-line::selection,
.cm-s-material .CodeMirror-line > span::selection,
.cm-s-material .CodeMirror-line > span > span::selection {
  background: rgba(128, 203, 196, 0.2);
}

.cm-s-material .CodeMirror-line::-moz-selection,
.cm-s-material .CodeMirror-line > span::-moz-selection,
.cm-s-material .CodeMirror-line > span > span::-moz-selection {
  background: rgba(128, 203, 196, 0.2);
}

.cm-s-material .CodeMirror-activeline-background {
  background: rgba(0, 0, 0, 0.5);
}

.cm-s-material .cm-keyword {
  color: #c792ea;
}

.cm-s-material .cm-operator {
  color: #89ddff;
}

.cm-s-material .cm-variable-2 {
  color: #eeffff;
}

.cm-s-material .cm-variable-3,
.cm-s-material .cm-type {
  color: #f07178;
}

.cm-s-material .cm-builtin {
  color: #ffcb6b;
}

.cm-s-material .cm-atom {
  color: #f78c6c;
}

.cm-s-material .cm-number {
  color: #ff5370;
}

.cm-s-material .cm-def {
  color: #82aaff;
}

.cm-s-material .cm-string {
  color: #4cc7e7;
}

.cm-s-material .cm-string-2 {
  color: #f07178;
}

.cm-s-material .cm-comment {
  color: #b6b6b6;
}

.cm-s-material .cm-variable {
  color: #ef19d2;
}

.cm-s-material .cm-tag {
  color: #ff5370;
}

.cm-s-material .cm-meta {
  color: #ffcb6b;
}

.cm-s-material .cm-attribute {
  color: #c792ea;
}

.cm-s-material .cm-property {
  color: #c792ea;
}

.cm-s-material .cm-qualifier {
  color: #decb6b;
}

.cm-s-material .cm-variable-3,
.cm-s-material .cm-type {
  color: #decb6b;
}

.cm-s-material .cm-error {
  color: rgba(255, 255, 255, 1);
  background-color: #ff5370;
}

.cm-s-material .CodeMirror-matchingbracket {
  text-decoration: underline;
  color: white !important;
}

/* BASICS */

.CodeMirror {
  /* Set height, width, borders, and global font properties here */
  font-family: monospace;
  height: 300px;
  color: #5e5656;
  direction: ltr;
}

/* PADDING */

.CodeMirror-lines {
  padding: 4px 0; /* Vertical padding around content */
}
.CodeMirror-lines {
  cursor: text;
  min-height: 1px; /* prevents collapsing before first draw */
}


.CodeMirror pre.CodeMirror-line,
.CodeMirror pre.CodeMirror-line-like {
  padding: 0 4px; /* Horizontal padding of content */
}

.CodeMirror-scrollbar-filler,
.CodeMirror-gutter-filler {
  background-color: white; /* The little square between H and V scrollbars */
}

/* GUTTER */

.CodeMirror-gutters {
  border-right: 1px solid #ddd;
  background-color: #f7f7f7;
  white-space: nowrap;
}

.CodeMirror-linenumbers {
}

.CodeMirror-linenumber {
  padding: 0 3px 0 5px;
  min-width: 20px;
  text-align: right;
  color: #999;
  white-space: nowrap;
}

.CodeMirror-guttermarker {
  color: black;
}

.CodeMirror-guttermarker-subtle {
  color: #999;
}

/* CURSOR */

.CodeMirror-cursor {
  border-left: 1px solid black;
  border-right: none;
  width: 0;
}

/* Shown when moving in bi-directional text */
.CodeMirror div.CodeMirror-secondarycursor {
  border-left: 1px solid silver;
}

.cm-fat-cursor .CodeMirror-cursor {
  width: auto;
  border: 0 !important;
  background: #7e7;
}

.cm-fat-cursor div.CodeMirror-cursors {
  z-index: 1;
}

.cm-fat-cursor .CodeMirror-line::selection,
.cm-fat-cursor .CodeMirror-line > span::selection,
.cm-fat-cursor .CodeMirror-line > span > span::selection {
  background: transparent;
}

.cm-fat-cursor .CodeMirror-line::-moz-selection,
.cm-fat-cursor .CodeMirror-line > span::-moz-selection,
.cm-fat-cursor .CodeMirror-line > span > span::-moz-selection {
  background: transparent;
}

.cm-fat-cursor {
  caret-color: transparent;
}

@-moz-keyframes blink {
  0% {
  }
  50% {
    background-color: transparent;
  }
  100% {
  }
}

@-webkit-keyframes blink {
  0% {
  }
  50% {
    background-color: transparent;
  }
  100% {
  }
}

@keyframes blink {
  0% {
  }
  50% {
    background-color: transparent;
  }
  100% {
  }
}

/* Can style cursor different in overwrite (non-insert) mode */
.CodeMirror-overwrite .CodeMirror-cursor {
}

.cm-tab {
  display: inline-block;
  text-decoration: inherit;
}

.CodeMirror-rulers {
  position: absolute;
  left: 0;
  right: 0;
  top: -50px;
  bottom: 0;
  overflow: hidden;
}

.CodeMirror-ruler {
  border-left: 1px solid #ccc;
  top: 0;
  bottom: 0;
  position: absolute;
}

/* DEFAULT THEME */

.cm-s-default .cm-header {
  color: blue;
}

.cm-s-default .cm-quote {
  color: #090;
}

.cm-negative {
  color: #d44;
}

.cm-positive {
  color: #292;
}

.cm-header,
.cm-strong {
  font-weight: normal;
}

.cm-em {
  font-style: italic;
}

.cm-link {
  text-decoration: underline;
}

.cm-strikethrough {
  text-decoration: line-through;
}

.cm-s-default .cm-keyword {
  color: #708;
}

.cm-s-default .cm-atom {
  color: #219;
}

.cm-s-default .cm-number {
  color: #164;
}

.cm-s-default .cm-def {
  color: #00f;
}

.cm-s-default .cm-variable,
.cm-s-default .cm-punctuation,
.cm-s-default .cm-property,
.cm-s-default .cm-operator {
}

.cm-s-default .cm-variable-2 {
  color: #05a;
}

.cm-s-default .cm-variable-3,
.cm-s-default .cm-type {
  color: #085;
}

.cm-s-default .cm-comment {
  color: #a50;
}

.cm-s-default .cm-string {
  color: #a11;
}

.cm-s-default .cm-string-2 {
  color: #f50;
}

.cm-s-default .cm-meta {
  color: #555;
}

.cm-s-default .cm-qualifier {
  color: #555;
}

.cm-s-default .cm-builtin {
  color: #30a;
}

.cm-s-default .cm-bracket {
  color: #997;
}

.cm-s-default .cm-tag {
  color: #170;
}

.cm-s-default .cm-attribute {
  color: #00c;
}

.cm-s-default .cm-hr {
  color: #999;
}

.cm-s-default .cm-link {
  color: #00c;
}

.cm-s-default .cm-error {
  color: #f00;
}

.cm-invalidchar {
  color: #f00;
}

.CodeMirror-composing {
  border-bottom: 2px solid;
}

/* Default styles for common addons */

div.CodeMirror span.CodeMirror-matchingbracket {
  color: #0b0;
}

div.CodeMirror span.CodeMirror-nonmatchingbracket {
  color: #a22;
}

.CodeMirror-matchingtag {
  background: rgba(255, 150, 0, 0.3);
}

.CodeMirror-activeline-background {
  background: #e8f2ff;
}

/* STOP */

/* The rest of this file contains styles related to the mechanics of
   the editor. You probably shouldn't touch them. */

.CodeMirror {
  position: relative;
  overflow: hidden;
  background: white;
}

.CodeMirror-scroll {
  overflow: scroll !important; /* Things will break if this is overridden */
  /* 50px is the magic margin used to hide the element's real scrollbars */
  /* See overflow: hidden in .CodeMirror */
  margin-bottom: -50px;
  margin-right: -50px;
  padding-bottom: 50px;
  height: 100%;
  outline: none; /* Prevent dragging from highlighting the element */
  position: relative;
  z-index: 0;
}

.CodeMirror-sizer {
  position: relative;
  border-right: 50px solid transparent;
}

/* The fake, visible scrollbars. Used to force redraw during scrolling
   before actual scrolling happens, thus preventing shaking and
   flickering artifacts. */
.CodeMirror-vscrollbar,
.CodeMirror-hscrollbar,
.CodeMirror-scrollbar-filler,
.CodeMirror-gutter-filler {
  position: absolute;
  z-index: 6;
  display: none;
  outline: none;
}

.CodeMirror-vscrollbar {
  right: 0;
  top: 0;
  overflow-x: hidden;
  overflow-y: scroll;
}

.CodeMirror-hscrollbar {
  bottom: 0;
  left: 0;
  overflow-y: hidden;
  overflow-x: scroll;
}

.CodeMirror-scrollbar-filler {
  right: 0;
  bottom: 0;
}

.CodeMirror-gutter-filler {
  left: 0;
  bottom: 0;
}

.CodeMirror-gutters {
  position: absolute;
  left: 0;
  top: 0;
  min-height: 100%;
  z-index: 3;
}

.CodeMirror-gutter {
  white-space: normal;
  height: 100%;
  display: inline-block;
  vertical-align: top;
  margin-bottom: -50px;
}

.CodeMirror-gutter-wrapper {
  position: absolute;
  z-index: 4;
  background: none !important;
  border: none !important;
}

.CodeMirror-gutter-background {
  position: absolute;
  top: 0;
  bottom: 0;
  z-index: 4;
}

.CodeMirror-gutter-elt {
  position: absolute;
  cursor: default;
  z-index: 4;
}

.CodeMirror-gutter-wrapper ::selection {
  background-color: transparent;
}

.CodeMirror-gutter-wrapper ::-moz-selection {
  background-color: transparent;
}


.CodeMirror pre.CodeMirror-line,
.CodeMirror pre.CodeMirror-line-like {
  /* Reset some styles that the rest of the page might have set */
  -moz-border-radius: 0;
  -webkit-border-radius: 0;
  border-radius: 0;
  border-width: 0;
  background: transparent;
  font-family: inherit;
  font-size: 20px;
  margin: 0;
  white-space: pre;
  word-wrap: normal;
  line-height: inherit;
  color: inherit;
  z-index: 2;
  position: relative;
  overflow: visible;
  -webkit-tap-highlight-color: transparent;
  -webkit-font-variant-ligatures: contextual;
  font-variant-ligatures: contextual;
}

.CodeMirror-wrap pre.CodeMirror-line,
.CodeMirror-wrap pre.CodeMirror-line-like {
  word-wrap: break-word;
  white-space: pre-wrap;
  word-break: normal;
}

.CodeMirror-linebackground {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  z-index: 0;
}

.CodeMirror-linewidget {
  position: relative;
  z-index: 2;
  padding: 0.1px; /* Force widget margins to stay inside of the container */
}

.CodeMirror-widget {
}

.CodeMirror-rtl pre {
  direction: rtl;
}

.CodeMirror-code {
  outline: none;
}

/* Force content-box sizing for the elements where we expect it */
.CodeMirror-scroll,
.CodeMirror-sizer,
.CodeMirror-gutter,
.CodeMirror-gutters,
.CodeMirror-linenumber {
  -moz-box-sizing: content-box;
  box-sizing: content-box;
}

.CodeMirror-linenumber {
  /*position: relative;*/
  top: 6px;
}

.CodeMirror-measure {
  position: absolute;
  width: 100%;
  height: 0;
  overflow: hidden;
  visibility: hidden;
}

.CodeMirror-cursor {
  position: absolute;
  pointer-events: none;
}

.CodeMirror-measure pre {
  position: static;
}

div.CodeMirror-cursors {
  visibility: hidden;
  position: relative;
  z-index: 3;
}

div.CodeMirror-dragcursors {
  visibility: visible;
}

.CodeMirror-focused div.CodeMirror-cursors {
  visibility: visible;
}

.CodeMirror-selected {
  background: #d9d9d9;
}

.CodeMirror-focused .CodeMirror-selected {
  background: #d7d4f0;
}

.CodeMirror-crosshair {
  cursor: crosshair;
}

.CodeMirror-line::selection,
.CodeMirror-line > span::selection,
.CodeMirror-line > span > span::selection {
  background: #d7d4f0;
}

.CodeMirror-line::-moz-selection,
.CodeMirror-line > span::-moz-selection,
.CodeMirror-line > span > span::-moz-selection {
  background: #d7d4f0;
}

.cm-searching {
  background-color: #ffa;
  background-color: rgba(255, 255, 0, 0.4);
}

/* Used to force a border model for a node */
.cm-force-border {
  padding-right: 0.1px;
}

@media print {
  /* Hide the cursor when printing */
  .CodeMirror div.CodeMirror-cursors {
    visibility: hidden;
  }
}

/* See issue #2901 */
.cm-tab-wrap-hack:after {
  content: "";
}

/* Help users use markselection to safely style text background */
span.CodeMirror-selectedtext {
  background: none;
}

highlightable-input {
  width: 300px;
  height: 100px;
}

pre {
  font-weight: 10 !important;
}
</style>
