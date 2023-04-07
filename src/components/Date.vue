<template>
  <span class="date">{{ toDate }}</span>
</template>

<script>
export default {
  name: "DateUi",
  props: {
    date: String,
  },
  computed: {
    toDate() {
      return new Intl.DateTimeFormat("ru-RU", {
        day: "2-digit",
        month: "long",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      }).format(this.date === "" ? "" : new Date(this.date));
    },
  },
};
</script>