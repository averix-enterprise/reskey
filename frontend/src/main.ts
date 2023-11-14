import {createApp} from 'vue'
import { VueQueryPlugin } from "@tanstack/vue-query";
import App from './App.vue'
import './style.css';

const app = createApp(App)

// add any needed plugins
app.use(VueQueryPlugin)

// mount app
app.mount('#app')