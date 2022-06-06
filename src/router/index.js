import { createWebHistory, createRouter } from "vue-router"
import Home from "../pages/Home.vue"
import Login from "../pages/Login.vue"
import Reg from "../pages/Reg.vue"
import NotFound from "../pages/NotFound.vue"
import UserHome from "../pages/UserHome.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
    },
    {
        path: "/reg",
        name: "Reg",
        component: Reg,
    },
    {
        path: "/login",
        name: "Login",
        component: Login,
    },
    {
        path: '/userHome',
        name: "UserHome",
        component: UserHome,
    },
    {
        path: '/:catchAll(.*)*',
        name: "PageNotFound",
        component: NotFound,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to, from) => {
    const logged = sessionStorage.getItem("loggedIn")
    if (logged !== "true" && to.name === 'UserHome') {
        return { name: 'Login' }
    }
    if(logged === "true" && (to.name === 'Reg' || to.name === 'Login')){
        return {name: 'UserHome'}
    }
})

export default router