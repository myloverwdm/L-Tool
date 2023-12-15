import {createApp} from "vue";
import {createPinia} from "pinia";

import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import "./style.scss";

// @ts-ignore
import SearchBar from "search-bar-vue3";
import CodeDiff from "v-code-diff";


// import "./assets/main.css";

// @ts-ignore
import HighlightableInput from "vue-highlightable-input";

const app = createApp(App);

app.component('highlightable-input', HighlightableInput)
app.use(SearchBar);
app.use(createPinia());
app.use(router);
app.use(i18n);
app.use(CodeDiff);
// @ts-ignore
app.use(ElementPlus);
app.mount("#app");

