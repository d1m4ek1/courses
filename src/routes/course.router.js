import Courses from "../pages/Courses.vue";
import Add from "../pages/Add.vue";
import Course from "../pages/Course.vue";

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
  {
    path: "/courses/:id",
    component: Course,
  },
];

export default router;
