import { createApp } from "vue";
import router from "./routes/index.routes";
import App from "./App.vue";
import store from "./store/store";

import "./assets/Main.css";

createApp(App).use(store).use(router).mount("#app");
