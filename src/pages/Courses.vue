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
            <Priceui :price="course.price" />
            <p>Создан <Dateui :date="course.addDate" /></p>
          </div>
          <div class="card-action actions">
            <router-link :to="`/courses/${course.id}`" target="_blank"
              >Открыть курс</router-link
            >
            <template v-if="userAuth">
              <router-link :to="`/courses/${course.id}/edit`"
                >Редактировать</router-link
              >
              <button @click="addToCart(course.id)" class="btn btn-primary">
                Купить
              </button>
            </template>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import User from "../models/user";
import Auth from "@/models/auth";
import Priceui from "@/components/Price.vue";
import Dateui from "@/components/Date.vue";
import { mapGetters } from "vuex";

export default {
  name: "CoursesVue",
  data() {
    return {
      courses: [],
    };
  },
  methods: {
    async addToCart(courseId) {
      const data = {
        courseId,
        userId: Auth.getPersonID,
      };

      await User.AddCourseToCart(this.XCSRFToken, data)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) this.$router.push("/cart");
        });
    },
  },
  async created() {
    const ref = this;
    await User.getCourses()
      .then((response) => response.json())
      .then((data) => (ref.courses = data.data));
  },
  components: {
    Priceui,
    Dateui,
  },
  computed: {
    ...mapGetters(["userAuth", "XCSRFToken"]),
  },
};
</script>