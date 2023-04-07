import { createWebHistory, createRouter } from "vue-router";

const routes = [];

import Login from "./auth.router";
routes.push(...Login);

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
