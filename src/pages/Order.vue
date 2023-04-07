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
            <Dateui :date="order.addDateOrder" />
            <p>
              <em>{{
                `${order.secondName} ${order.firstName} ${order.thirdName}`
              }}</em>
              (email: {{ order.email }})
            </p>

            <ol
              v-for="(item, idxx) in order.items"
              :key="`order${idx}_item${idxx}`"
            >
              <li>
                <a :href="`/courses/${item.courseId}`" target="_blank">{{
                  item.title
                }}</a>
                (<Priceui :price="item.price" /> x<strong>{{
                  order.count[idxx]
                }}</strong
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
import Dateui from "@/components/Date.vue";
import { mapGetters } from "vuex";

export default {
  name: "OrderVue",
  data() {
    return {
      orders: [],
    };
  },
  methods: {
    totalPrice(arr) {
      return arr.items.reduce((total, product, index) => {
        return (total += product.price * arr.count[index]);
      }, 0);
    },
    async deleteOrder(orderId) {
      await Order.DeleteOrder(this.XCSRFToken, orderId)
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
      .then((response) => {
        for (let i = 0; i < response.data.length; i++) {
          const course = response.data[i];
          course.count = course.count.reverse();
        }
        this.orders = response.data;
      });
  },
  components: {
    Priceui,
    Dateui,
  },
  computed: {
    ...mapGetters(["XCSRFToken"]),
  },
};
</script>