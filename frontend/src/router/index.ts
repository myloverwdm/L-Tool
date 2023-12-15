import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

// @ts-ignore
// @ts-ignore
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/about",
      name: "about",
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/onlineTool/timestamp",
      name: "onlineTool-timestamp",
      component: () => import("../views/onlineTool/Timestamp.vue"),
    },
    {
      path: "/onlineTool/json",
      name: "onlineTool-json",
      component: () => import("../views/onlineTool/json.vue"),
    },
    {
      path: "/onlineTool/zk",
      name: "onlineTool-zk",
      component: () => import("../views/onlineTool/zk.vue"),
    },
    {
      path: "/onlineTool/command",
      name: "command",
      component: () => import("../views/onlineTool/command.vue"),
    },
    {
      path: "/onlineTool/kafka",
      name: "kafka",
      component: () => import("../views/onlineTool/kafka.vue"),
    }
  ],
});

export default router;
