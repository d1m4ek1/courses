import { createApp } from "vue";
import router from "./routes/index.routes";
import App from "./App.vue";

import "./assets/Main.css";
import "./assets/utils";

createApp(App).use(router).mount("#app");
