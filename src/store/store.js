import { createStore } from "vuex";

// Create a new store instance.
const store = createStore({
  state() {
    return {
      isAuth: false,
      X_CSRF_TOKEN: "",
    };
  },
  mutations: {
    userLogin(state) {
      state.isAuth = true;
    },
    userLogout(state) {
      state.isAuth = false;
    },
    setCSRFToken(state, token) {
      state.X_CSRF_TOKEN = token;
    },
  },
  getters: {
    userAuth(state) {
      return state.isAuth;
    },
    XCSRFToken(state) {
      return state.X_CSRF_TOKEN;
    },
  },
});

export default store;
