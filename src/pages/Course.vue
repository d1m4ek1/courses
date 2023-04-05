<template>
  <div class="container">
    <div class="banner_title">
      <img :src="courseData.banner" :alt="courseData.title" />
      <h1>{{ courseData.title }}</h1>
    </div>
    <Priceui :price="courseData.price" />
  </div>
</template>

<script>
import User from "../models/user";
import Priceui from "@/components/Price.vue";

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
  },
  components: {
    Priceui,
  },
};
</script>