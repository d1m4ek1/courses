<template>
  <div class="container">
    <div class="auth">
      <div class="row">
        <div class="col s12">
          <ul class="tabs">
            <li class="tab col s6">
              <a class="active" href="/auth/login#login">Войти</a>
            </li>
            <li class="tab col s6">
              <a href="/auth/register#register">Регистрация</a>
            </li>
          </ul>
        </div>
        <div id="login" class="col s6 offset-s3">
          <!-- {{#if loginError}}
          <p class="alert">{{ loginError }}</p>
          {{/if}} -->
          <p class="alert" v-if="errorAuth">Не верный логин или пароль</p>
          <h4>Войти в магазин</h4>
          <form action="/login" method="POST">
            <div class="input-field">
              <input
                id="email"
                name="email"
                type="email"
                v-model="loginData.email"
                class="validate"
                required
              />
              <label for="email">Email</label>
              <span class="helper-text" data-error="Введите email"></span>
            </div>

            <div class="input-field">
              <input
                id="password"
                name="password"
                type="password"
                v-model="loginData.password"
                class="validate"
                required
              />
              <label for="password">Пароль</label>
              <span class="helper-text" data-error="Введите пароль"></span>
            </div>

            <!-- <input type="hidden" name="_csrf" value="{{csrf}}" /> -->

            <button class="btn btn-primary" @click="login($event)">
              Войти
            </button>
          </form>
        </div>
        <div id="register" class="col s6 offset-s3">
          <!-- {{#if registerError}}
          <p class="alert">{{ registerError }}</p>
          {{/if}} -->
          <h4>Создать аккаунт</h4>
          <form action="/register" method="POST">
            <div class="input-field">
              <input
                v-model="registerData.email"
                id="remail"
                name="email"
                type="email"
                class="validate"
                required
              />
              <label for="remail">Email</label>
              <span class="helper-text" data-error="Введите email"></span>
            </div>

            <div class="input-field">
              <input
                v-model="registerData.password"
                id="rpassword"
                name="password"
                type="password"
                class="validate"
                required
              />
              <label for="rpassword">Пароль</label>
              <span class="helper-text" data-error="Введите пароль"></span>
            </div>

            <div class="input-field">
              <input
                v-model="registerData.confirmPassword"
                id="confirm"
                name="confirm"
                type="password"
                class="validate"
                required
              />
              <label for="confirm">Повторите пароль</label>
              <span class="helper-text" data-error="Введите пароль"></span>
            </div>

            <div class="input-field">
              <input
                v-model="registerData.firstName"
                id="firstName"
                name="firstName"
                type="text"
                class="validate"
                required
              />
              <label for="firstName">Ваше имя</label>
              <span class="helper-text" data-error="Введите имя"></span>
            </div>

            <div class="input-field">
              <input
                v-model="registerData.secondName"
                id="secondName"
                name="secondName"
                type="text"
                class="validate"
                required
              />
              <label for="secondName">Ваше фамилия</label>
              <span class="helper-text" data-error="Введите фамилию"></span>
            </div>

            <div class="input-field">
              <input
                v-model="registerData.thirdName"
                id="thirdName"
                name="thirdName"
                type="text"
                class="validate"
                required
              />
              <label for="thirdName">Ваше отчество</label>
              <span class="helper-text" data-error="Введите отчество"></span>
            </div>

            <!-- <input type="hidden" name="_csrf" value="{{csrf}}" /> -->

            <button class="btn btn-primary" @click="register($event)">
              Зарегистрироваться
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Auth from "@/models/auth";
import { mapMutations } from "vuex";

export default {
  name: "LoginVue",
  data() {
    return {
      loginData: {
        email: "",
        password: "",
      },
      registerData: {
        email: "",
        confirmPassword: "",
        password: "",
        firstName: "",
        secondName: "",
        thirdName: "",
      },
      errorAuth: false,
    };
  },
  methods: {
    async login(e) {
      e.preventDefault();

      await Auth.Login(this.loginData)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            if (response.isLogin) {
              this.userLogin();
              this.errorAuth = false;
              return this.$router.push("/");
            } else {
              this.errorAuth = true;
            }
          }
        });
    },
    async register(e) {
      e.preventDefault();

      await Auth.Register(this.registerData)
        .then((response) => response.json())
        .then((response) => {
          if (response.successfully) {
            if (response.isLogin) {
              this.userLogin();
              this.errorAuth = false;
              return this.$router.push("/");
            } else {
              this.errorAuth = true;
            }
          }
        });
    },
    ...mapMutations(["userLogin"]),
  },
  mounted() {
    // eslint-disable-next-line
    M.Tabs.init(document.querySelectorAll(".tabs"));
  },
};
</script>