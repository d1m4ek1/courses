<template>
  <span class="price">{{ toCurrency }}</span>
</template>

<script>
export default {
  name: "PriceUI",
  props: {
    price: Number,
  },
  computed: {
    toCurrency() {
      return new Intl.NumberFormat("ru-RU", {
        currency: "rub",
        style: "currency",
      }).format(this.price);
    },
  },
};
</script>