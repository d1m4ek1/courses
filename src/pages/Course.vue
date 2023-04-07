<template>
  <div class="container">
    <div class="banner_title">
      <img :src="courseData.banner" :alt="courseData.title" />
      <h1>{{ courseData.title }}</h1>
    </div>
    <div class="course_content">
      <p>Цена: <Priceui :price="courseData.price" /></p>
      <p>
        Основатель
        <i
          >({{
            `${courseData.secondName} ${courseData.firstName} ${courseData.thirdName}`
          }})</i
        >
      </p>
      <p>Создан <Dateui :date="courseData.addDate" /></p>
      <p v-if="courseData.editDate">
        В последний раз обновлен
        <Dateui :date="courseData.editDate" />
      </p>
    </div>
  </div>
</template>

<script>
import User from "../models/user";
import Priceui from "@/components/Price.vue";
import Dateui from "@/components/Date.vue";

export default {
  name: "CourseVue",
  data() {
    return {
      courseData: {
        title: "",
        price: 0,
        banner: "",
        addDate: "",
        editDate: "",
        firstName: "",
        secondName: "",
        thirdName: "",
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
    Dateui,
  },
};
</script>