import {createStore} from 'vuex'
import userModules from "@/store/module/user";

export default createStore({
    strict: process.env.NODE_ENV !== "production",
    state() {

    },
    mutations: {

    },
    getters: {

    },
    modules: {
        userModules
    }
})