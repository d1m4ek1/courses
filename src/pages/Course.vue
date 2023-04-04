<template>
  <div class="container">
    <div class="banner_title">
      <img :src="courseData.banner" :alt="courseData.title" />
      <h1>{{ courseData.title }}</h1>
    </div>
    <p class="price">{{ courseData.price }}</p>
  </div>
</template>

<script>
import { DOMToCurrency } from "@/assets/utils";
import User from "../models/user";
export default {
  name: "CourseVue",
  data() {
    return {
      courseData: {
        title: "",
        price: "",
        banner: "",
      },
    };
  },
  async created() {
    const ref = this;

    await User.getCourseById(this.$route.params.id)
      .then((response) => response.json())
      .then((data) => (ref.courseData = { ...data.data }));

    window.document.title = this.courseData.title;
    DOMToCurrency();
  },
};
</script>