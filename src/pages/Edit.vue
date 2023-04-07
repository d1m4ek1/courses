<template>
  <div class="container">
    <h1>Редактировать курс</h1>
    <form
      :class="{
        hide: indicators.isEditedData === 1 || indicators.isEditedData === -1,
      }"
      action="/courses/edit"
      method="POST"
      class="course-form"
    >
      <div class="input-field col s6">
        <input
          id="title"
          name="title"
          type="text"
          class="validate"
          required
          @input="(editedData.title = $event.target.value), validInputs()"
          :value="editedData.title"
        />
        <label :class="{ active: course.title }" for="title"
          >Название курса</label
        >
        <span class="helper-text" data-error="Введите название"></span>
      </div>

      <div class="input-field col s6">
        <input
          id="price"
          name="price"
          type="number"
          class="validate"
          required
          min="1"
          @input="(editedData.price = +$event.target.value), validInputs()"
          :value="editedData.price"
        />
        <label :class="{ active: course.price }" for="price">Цена курса</label>
        <span class="helper-text" data-error="Введите цену"></span>
      </div>

      <div class="input-field col s6">
        <input
          id="banner"
          name="banner"
          type="text"
          class="validate"
          required
          @input="(editedData.banner = $event.target.value), validInputs()"
          :value="editedData.banner"
        />
        <label :class="{ active: course.banner }" for="banner"
          >URL баннера</label
        >
        <span class="helper-text" data-error="Введите url баннера"></span>
      </div>

      <button
        :disabled="!validInputs()"
        @click="editCourse($event)"
        class="btn btn-primary"
      >
        Редактировать
      </button>

      <button @click="deleteCourse($event)" class="btn red">
        Удалить курс
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
import { mapGetters } from "vuex";
import User from "../models/user";

export default {
  name: "EditCourse",
  data() {
    return {
      course: {
        id: "",
        title: "",
        price: 0,
        banner: "",
      },
      editedData: {
        id: "",
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
    validInputs() {
      let flagEdit = false;

      Object.keys(this.course).forEach((key) => {
        if (key === "id") return;

        if (this.course[key] !== this.editedData[key]) {
          flagEdit = true;
        }

        if (!this.editedData[key]) {
          flagEdit = false;
        }
      });

      return flagEdit;
    },
    async editCourse(e) {
      e.preventDefault();

      if (this.validInputs()) {
        const ref = this;
        this.indicators.isEditedData = 1;
        await User.editCourse(
          this.XCSRFToken,
          this.$route.params.id,
          this.editedData
        )
          .then((response) => response.json())
          .then((response) => (ref.indicators = { ...response }));

        this.indicators.isEditedData = -1;

        setTimeout(() => {
          this.indicators.isEditedData = 0;
        }, 2000);
      }
    },
    async deleteCourse(e) {
      e.preventDefault();

      User.deleteCourse(this.XCSRFToken, this.$route.params.id)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            this.$router.push("/courses");
          }
        });
    },
  },
  async created() {
    const ref = this;

    await User.getCourseById(this.$route.params.id)
      .then((response) => response.json())
      .then((response) => {
        [
          "userId",
          "id",
          "addDate",
          "editDate",
          "firstName",
          "secondName",
          "thirdName",
        ].forEach((key) => {
          delete response.data[key];
        });

        [ref.course, ref.editedData] = [
          { ...response.data },
          { ...response.data },
        ];
      });
  },
  computed: {
    ...mapGetters(["XCSRFToken"]),
  },
};
</script>

<style scoped>
.hide {
  display: none;
}
</style>