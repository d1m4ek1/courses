import { createApp } from "vue";
import router from "./routes/index.routes";
import App from "./App.vue";

import "./assets/Main.css";

createApp(App).use(router).mount("#app");
