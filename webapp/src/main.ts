import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import './style.css'
import App from './App.vue'

import 'primevue/resources/themes/bootstrap4-dark-purple/theme.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.css'

import Button from "primevue/button"
import ConfirmationService from 'primevue/confirmationservice';


const pinia = createPinia()


const app = createApp(App)
app.use(pinia)
app.use(PrimeVue, { ripple: true })
app.use(ConfirmationService);

app.component('Button', Button);

app.mount('#app')
