<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted, onUnmounted} from "vue";

const {t} = useI18n();

var timer: number | undefined;
onMounted(() => {
  timer = setInterval(() => {
    doimer();
  }, 1000);
});

onUnmounted(() => {
  clearInterval(timer);
});

function doimer() {
  const timestamp = new Date().getTime();
  let selectElement = getSelectElement({id: "fruits"});
  if ("ms" === selectElement) {
    setValue("now", timestamp + "");
  } else {
    setValue("now", Math.floor(timestamp / 1000) + "");
  }
}

// 根据id得到当前选择的值
function getSelectElement({ id }: { id: string }): string {
  var obj = document.getElementById(id) as HTMLSelectElement;
  if (obj == null) {
    return "";
  }
  for (var i = 0; i < obj.length; i++) {
    if ((obj[i] as HTMLOptionElement).selected) {
      var text = (obj[i] as HTMLOptionElement).value;
      return text;
    }
  }
  return "";
}

// 根据id设置为指定的值
function setValue(id: string, value: string) {
  var myDiv = document.getElementById(id);
  if (myDiv != null) {
    myDiv.innerHTML = value;
  }
}

function setInputValue(id: string, value: string) {
  var myDiv = document.getElementById(id) as HTMLInputElement;
  if (myDiv != null) {
    myDiv.value = value;
  }
}

// 根据id得到指定的值
function getValue(id: string) {
  var myDiv = document.getElementById(id);
  if (myDiv != null) {
    return myDiv.innerHTML;
  }
}

function cknow() {
  setInputValue("sjc1", getValue("now") + "")
}

function zhuanh1() {
  let elementById = document.getElementById("sjc1") as HTMLInputElement;
  if (elementById == null) {
    return;
  }
  var myDiv = elementById.value;
  if (myDiv == null || "" == myDiv) {
    return;
  }
  let intS = parseInt(myDiv.toString());
  let selectElement = getSelectElement({id: "fruits"});
  try {
    if ("s" === selectElement) {
      intS = intS * 1000;
    }
    setInputValue("times", formatTimestamp(intS + ""));
  } catch (errort) {

  }
}


function zhuanh2() {
  let elementById = document.getElementById("time2") as HTMLInputElement ;
  if (elementById == null) {
    return;
  }
  var myDiv = elementById.value;
  console.log(myDiv)
  let timestamp = Date.parse(myDiv);
  let selectElement = getSelectElement({id: "fruits2"});
  if ("s" === selectElement) {
    timestamp = Math.floor(timestamp / 1000);
  }
  setInputValue("sjc3", timestamp + "")
}

function formatTimestamp(timestamp: string) {
  const date = new Date(parseInt(timestamp));
  const year = date.getFullYear();
  const month = ("0" + (date.getMonth() + 1)).slice(-2);
  const day = ("0" + date.getDate()).slice(-2);
  const hours = ("0" + date.getHours()).slice(-2);
  const minutes = ("0" + date.getMinutes()).slice(-2);
  const seconds = ("0" + date.getSeconds()).slice(-2);
  let s = year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;
  if (s.includes("N")) {
    return "Invalid Date";
  }
  return s;
}
</script>

<template>
  <div class="all-T">
    <h1>{{ t("tool-timestamp.timestampP") }}</h1>
    <div class="first">
      <div class="all-T-F">{{ t("tool-timestamp.now") }}:</div>
      <div class="all-T-S">
        <span id="now" @click="cknow"></span>
      </div>
    </div>
    <div class="first">
      <div class="all-T-F">{{ t("tool-timestamp.timestamp") }}</div>
      <div>
        <input class="all-T-S all-T-input" id="sjc1" autocomplete='off'/>
      </div>
      <div>
        <select id="fruits" name="fruits">
          <option value="ms">{{ t("tool-timestamp.ms") }}</option>
          <option value="s">{{ t("tool-timestamp.s") }}</option>
        </select>
      </div>
      <div>
        <el-button plain @click="zhuanh1">{{ t("tool-timestamp.transform") }}</el-button>
      </div>
      <div>
        <input class="all-T-input" id="times" autocomplete='off'/>
      </div>
      <div>{{ t("tool-timestamp.beijing") }}</div>
    </div>

    <div class="first">
      <div class="all-T-F">{{ t("tool-timestamp.time") }}</div>
      <div>
        <input class="all-T-S all-T-input" id="time2" autocomplete='off'/></div>
      <div>{{ t("tool-timestamp.beijing") }}</div>
      <div>
        <el-button plain @click="zhuanh2">{{ t("tool-timestamp.transform") }}</el-button>
      </div>
      <div>
        <input class="all-T-input" id="sjc3" autocomplete='off'/></div>
      <div>
        <select id="fruits2" name="fruits">
          <option value="ms">{{ t("tool-timestamp.ms") }}</option>
          <option value="s">{{ t("tool-timestamp.s") }}</option>
        </select>
<!--        <el-select class="m-2">-->
<!--          <el-option value="ms">{{ t("tool-timestamp.ms") }}</el-option>-->
<!--          <el-option value="s">{{ t("tool-timestamp.s") }}</el-option>-->
<!--        </el-select>-->
      </div>
    </div>

  </div>


</template>

<style>
body {
  overflow-x: hidden;
}

.all-T {
  width: 80%;
  margin: 0 auto;
}

.all-T h1 {
  font-size: 30px;
}

input:focus {
  outline: none;
}

.first {
  display: flex;
  margin-bottom: 40px;
  margin-top: 40px;
  font-size: 18px;
}

.first div {
  margin-right: 10px;
  margin-left: 10px;
}

.all-T-F {
  width: 10%;
}

.all-T-S {
  width: 100%;
}

.all-T-input {
  border: 1px solid #64d390;
}

#sjc1 {
  maxlength: 13
}

@media (max-width: 2000px) {
  .first {
    font-size: 18px;
  }
}

/* 当屏幕宽度小于等于480px时，设置字体大小为12px */
@media (max-width: 1200px) {
  .first {
    font-size: 16px;
  }
}

/* 当屏幕宽度小于等于480px时，设置字体大小为12px */
@media (max-width: 1085px) {
  .first {
    font-size: 15px;
  }
}

</style>