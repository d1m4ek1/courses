<template>
  <Navbar />
  <router-view></router-view>
</template>

<script>
import { mapMutations } from "vuex";
import Navbar from "./components/Navbar.vue";
import Auth from "./models/auth";

export default {
  name: "App",
  components: {
    Navbar,
  },
  methods: {
    ...mapMutations(["userLogin", "userLogout"]),
  },
  async created() {
    await Auth.CheckUserAuth()
      .then((response) => response.json())
      .then((response) => {
        if (response.successfully) {
          if (response.isAuth) this.userLogin();
          else this.userLogout();
        }
      });
  },
};
</script>