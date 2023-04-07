import { createStore } from "vuex";

// Create a new store instance.
const store = createStore({
  state() {
    return {
      isAuth: false,
    };
  },
  mutations: {
    userLogin(state) {
      state.isAuth = true;
    },
    userLogout(state) {
      state.isAuth = false;
    },
  },
  getters: {
    userAuth(state) {
      return state.isAuth;
    },
  },
});

export default store;
