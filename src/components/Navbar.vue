<template>
  <nav>
    <div class="nav-wrapper">
      <a href="#" class="brand-logo">Приложение курсов</a>
      <ul id="nav-mobile" class="right hide-on-med-and-down">
        <router-link to="/" v-slot="{ href, navigate, isActive }" custom>
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Главная</a>
          </li>
        </router-link>
        <router-link to="/courses" v-slot="{ href, navigate, isActive }" custom>
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Курсы</a>
          </li>
        </router-link>
        <template v-if="userAuth">
          <router-link
            to="/courses/add"
            v-slot="{ href, navigate, isActive }"
            custom
          >
            <li :class="[isActive && 'active']">
              <a :href="href" @click="navigate">Добавить курс</a>
            </li>
          </router-link>
          <router-link
            v-if="userAuth"
            to="/cart"
            v-slot="{ href, navigate, isActive }"
            custom
          >
            <li :class="[isActive && 'active']">
              <a :href="href" @click="navigate">Корзина</a>
            </li>
          </router-link>
          <router-link
            to="/orders"
            v-slot="{ href, navigate, isActive }"
            custom
          >
            <li :class="[isActive && 'active']">
              <a :href="href" @click="navigate">Заказы</a>
            </li>
          </router-link>
          <li>
            <a @click="logout()">Выйти</a>
          </li>
        </template>
        <router-link
          v-else
          to="/auth/login"
          v-slot="{ href, navigate, isActive }"
          custom
        >
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Войти</a>
          </li>
        </router-link>
      </ul>
    </div>
  </nav>
</template>

<script>
import Auth from "@/models/auth";
import { mapGetters, mapMutations } from "vuex";

export default {
  name: "NavBar",
  methods: {
    showOrderLink() {
      return Auth.getAnswerAllow;
    },
    async logout() {
      await Auth.Logout(this.XCSRFToken)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            this.userLogout();
            return this.$router.push("/");
          }
        });
    },
    ...mapMutations(["userLogout"]),
  },
  computed: {
    ...mapGetters(["userAuth", "XCSRFToken"]),
  },
};
</script>