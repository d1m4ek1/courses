<template>
  <div class="container">
    <h1>Добавить курс</h1>

    <form
      :class="{
        hide: indicators.isEditedData === 1 || indicators.isEditedData === -1,
      }"
      action="/add"
      method="POST"
    >
      <div class="input-field col s6">
        <input
          @input="(courseData.title = $event.target.value), validInput()"
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
          @input="
            (courseData.price = Number($event.target.value)), validInput()
          "
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
          @input="(courseData.banner = $event.target.value), validInput()"
          id="banner"
          name="banner"
          type="text"
          class="validate"
          required
        />
        <label for="banner">URL баннера</label>
        <span class="helper-text" data-error="Введите url баннера"></span>
      </div>

      <button
        :disabled="!validInput()"
        @click="addCourse($event)"
        class="btn btn-primary"
      >
        Добавить курс
      </button>
    </form>
    <span
      :class="{
        hide: indicators.isEditedData === 0 || indicators.isEditedData === -1,
      }"
      class="loader"
    ></span>
    <h5 v-if="indicators.isEditedData === -1 && indicators.successfully">
      Данные редактированны
    </h5>
    <h5 v-if="indicators.isEditedData === -1 && !indicators.successfully">
      Неизвестная ошибка
    </h5>
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
        price: 0,
        banner: "",
      },
      indicators: {
        isEditedData: 0,
        successfully: null,
      },
    };
  },
  methods: {
    validInput() {
      let countNotNull = 0;
      Object.keys(this.courseData).forEach((key) => {
        if (this.courseData[key] !== "") {
          countNotNull++;
        }
      });

      if (countNotNull === 3) {
        return true;
      }
      return false;
    },
    async addCourse(event) {
      event.preventDefault();

      this.indicators.isEditedData = 1;

      await User.addCourse(this.courseData)
        .then((response) => response.json())
        .then((response) => (this.indicators = { ...response }));

      this.indicators.isEditedData = -1;

      setTimeout(() => {
        this.indicators.isEditedData = 0;
      }, 2000);
    },
  },
};
</script>

<style scoped>
.hide {
  display: none;
}
</style>