import Login from "@/pages/Login";

const router = [
  {
    path: "/auth/login",
    component: Login,
    name: "Войти в аккаунт",
  },
  {
    path: "/auth/register",
    component: Login,
    name: "Создать аккаунт",
  },
];

export default router;
