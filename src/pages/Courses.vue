<template>
  <div class="container">
    <h1>Курсы</h1>
    <div class="row">
      <div
        v-for="(course, idx) in courses"
        :key="`course${idx}`"
        class="col s6"
      >
        <div class="card">
          <div class="card-image">
            <img :src="course.banner" :alt="course.title" />
            <span class="card-title">{{ course.title }}</span>
          </div>
          <div class="card-content">
            <p class="price">{{ course.price }}</p>
          </div>
          <div class="card-action actions">
            <router-link :to="`/courses/${course.id}`" target="_blank"
              >Открыть курс</router-link
            >
            <router-link :to="`/courses/${course.id}/edit`"
              >Редактировать</router-link
            >
            <button type="submit" class="btn btn-primary">Купить</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { DOMToCurrency } from "@/assets/utils";
import User from "../models/user";

export default {
  name: "CoursesVue",
  data() {
    return {
      courses: [],
    };
  },
  async created() {
    const ref = this;
    await User.getCourses()
      .then((response) => response.json())
      .then((data) => (ref.courses = data.data));

    DOMToCurrency();
  },
};
</script>