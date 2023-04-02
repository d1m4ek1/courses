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
        <router-link
          to="/courses/add"
          v-slot="{ href, navigate, isActive }"
          custom
        >
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Добавить курс</a>
          </li>
        </router-link>
        <router-link to="/cart" v-slot="{ href, navigate, isActive }" custom>
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Корзина</a>
          </li>
        </router-link>
        <router-link to="/orders" v-slot="{ href, navigate, isActive }" custom>
          <li :class="[isActive && 'active']">
            <a :href="href" @click="navigate">Заказы</a>
          </li>
        </router-link>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavBar",
};
</script>