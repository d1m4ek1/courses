<template>
  <div class="container">
    <h1>Добавить курс</h1>

    <form action="/add" method="POST">
      <div class="input-field col s6">
        <input
          v-model="courseData.title"
          id="title"
          name="title"
          type="text"
          class="validate"
          required
        />
        <label for="title">Название курса</label>
        <span class="helper-text" data-error="Введите название"></span>
      </div>

      <div class="input-field col s6">
        <input
          v-model="courseData.price"
          id="price"
          name="price"
          type="number"
          class="validate"
          required
          min="1"
        />
        <label for="price">Цена курса</label>
        <span class="helper-text" data-error="Введите цену"></span>
      </div>

      <div class="input-field col s6">
        <input
          v-model="courseData.banner"
          id="banner"
          name="banner"
          type="text"
          class="validate"
          required
        />
        <label for="banner">URL баннера</label>
        <span class="helper-text" data-error="Введите url баннера"></span>
      </div>

      <button @click="addCourse($event)" class="btn btn-primary">
        Добавить курс
      </button>
    </form>
  </div>
</template>

<script>
import User from "../models/user";

export default {
  name: "AddCourseVue",
  data() {
    return {
      courseData: {
        title: "",
        price: "",
        banner: "",
      },
    };
  },
  methods: {
    async addCourse(event) {
      event.preventDefault();

      await User.addCourse(this.courseData)
        .then((response) => response.json())
        .then((data) => console.log(data));
    },
  },
};
</script>