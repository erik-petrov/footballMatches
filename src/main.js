import { createApp } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons';
import VueCookies from 'vue-cookies'
import mitt from 'mitt';
const emitter = mitt();
import App from './App.vue'
import router from './router'
import vfmPlugin from 'vue-final-modal'

library.add(fas);

const app = createApp(App)
    .use(vfmPlugin)
    .use(VueCookies, {expire: '12h'})
    .use(router, {mode: 'history'})
    .component('fa', FontAwesomeIcon);

app.config.globalProperties.emitter = emitter
app.mount('#app')