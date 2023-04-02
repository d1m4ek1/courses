import Courses from "../pages/Courses.vue";
import Add from "../pages/Add.vue";

const router = [
  {
    path: "/courses",
    component: Courses,
    name: "Курсы",
  },
  {
    path: "/courses/add",
    component: Add,
    name: "Добавить курс",
  },
];

export default router;
