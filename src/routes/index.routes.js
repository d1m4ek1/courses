import { createWebHistory, createRouter } from "vue-router";

const routes = [];

import Home from "./home.router";
routes.push(...Home);

import Courses from "./course.router";
routes.push(...Courses);

import Cart from "./cart.router";
routes.push(...Cart);

import Order from "./order.router";
routes.push(...Order);

const router = createRouter({
  history: createWebHistory(),
  routes,
  linkActiveClass: "active",
  linkExactActiveClass: "exact-active",
});

router.beforeEach((to) => {
  document.title = to.name;
});

export default router;
