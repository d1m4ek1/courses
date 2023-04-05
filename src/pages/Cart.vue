<template>
  <div class="container">
    <h1>Корзина</h1>
    <div id="card">
      <template v-if="products.length">
        <table>
          <thead>
            <tr>
              <th>Название</th>
              <th>Цена</th>
              <th>Количество</th>
              <th>Действие</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(product, idx) in products" :key="`product${idx}`">
              <td>
                <a :href="`/courses/${product.courseId}`" target="_blank">
                  {{ product.title }}</a
                >
              </td>
              <td><Priceui :price="product.price" /></td>
              <td>{{ product.count }}</td>
              <td>
                <button
                  @click="deleteFromCart(product.id)"
                  class="btn btn-small js-remove"
                >
                  Удалить
                </button>
              </td>
            </tr>
          </tbody>
        </table>

        <p>Общая цена: <Priceui :price="totalPrice" /></p>
        <button @click="addToOrder()" class="btn">Сделать заказ</button>
      </template>
      <h5 v-else>Корзина пуста</h5>
    </div>
  </div>
</template>

<script>
import Priceui from "@/components/Price.vue";
import Cart from "@/models/cart";
import Order from "@/models/order";

export default {
  components: { Priceui },
  name: "CartVue",
  data() {
    return {
      products: [],
    };
  },
  methods: {
    async deleteFromCart(id) {
      await Cart.deleteFromCart(id)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            this.products = this.products.filter((p) =>
              p.id === id ? (p.count > 1 ? p.count-- : null) : p
            );
          }
        });
    },
    async addToOrder() {
      await Order.AddToOrder()
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            this.$router.push("/orders");
            return;
          }
        });
    },
  },
  async created() {
    await Cart.getAllProducts()
      .then((response) => response.json())
      .then((response) => {
        this.products = response.data ? response.data : [];
      });
  },
  computed: {
    totalPrice() {
      return this.products.reduce((total, product) => {
        return (total += product.price * product.count);
      }, 0);
    },
  },
};
</script>