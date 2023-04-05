<template>
  <div class="container">
    <h1>Заказы</h1>
    <div v-if="orders" class="row">
      <div class="col s6">
        <div class="card" v-for="(order, idx) in orders" :key="`order${idx}`">
          <div class="card-content">
            <span class="card-title">
              Заказ
              <small>№{{ order._id }}</small>
            </span>
            <!-- <p class="date">{{ order.date }}</p> -->
            <p><em>name: Test</em> (email: test@test.com)</p>

            <ol
              v-for="(item, idxx) in order.items"
              :key="`order${idx}_item${idxx}`"
            >
              <li>
                <a :href="`/courses/${item.courseId}`" target="_blank">{{
                  item.title
                }}</a>
                (x<strong>{{ item.count }}</strong
                >)
              </li>
            </ol>

            <hr />

            <p>Общая цена: <Priceui :price="totalPrice(order)" /></p>
            <button
              @click="deleteOrder(order._id)"
              style="margin-top: 20px"
              class="btn red"
            >
              Удалить заказ
            </button>
          </div>
        </div>
      </div>
    </div>
    <h5 v-else>Заказов пока нет</h5>
  </div>
</template>

<script>
import Order from "@/models/order";
import Priceui from "@/components/Price.vue";

export default {
  name: "OrderVue",
  data() {
    return {
      orders: [],
    };
  },
  methods: {
    totalPrice(arr) {
      return arr.items.reduce((total, product) => {
        return (total += product.price * product.count);
      }, 0);
    },
    async deleteOrder(orderId) {
      await Order.DeleteOrder(orderId)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            this.orders = this.orders.filter((o) => o._id !== orderId);
          }
        });
    },
  },
  async created() {
    await Order.GetOrders()
      .then((response) => response.json())
      .then((response) => (this.orders = response.data));
  },
  components: {
    Priceui,
  },
};
</script>