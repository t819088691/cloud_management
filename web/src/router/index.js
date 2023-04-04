import {createRouter, createWebHashHistory} from "vue-router";
import userModule from "@/store/module/user";

const routes = [{
    path: "/dashboard", component: () => import("@/components/nav-menu"), children: [   // 添加子路由
        {
            path: '/dashboard/billing/index',
            meta: {requiresAuth: true},
            components: {dashboard: () => import('@/pages/billing/summary')}
        }, {
            path: '/dashboard/billing/gcp',
            meta: {requiresAuth: true},
            components: {
                dashboard: () => import('@/pages/billing/gcp')
            }
        }, {
            path: '/dashboard/billing/aws', components: {
                dashboard: () => import('@/pages/billing/aws')
            }
        }, {
            path: '/dashboard/billing/azure', components: {
                dashboard: () => import('@/pages/billing/azure')
            }
        }, {
            path: "/dashboard/user/list",
            meta: {requiresAuth: true},
            components: {
                dashboard: () => import("@/pages/user/list")
            }
        }]
}, {
    path: "/user/forget", component: () => import("@/pages/user/forget"),
}, {
    path: "/user/registry", component: () => import("@/pages/user/registry")
}, {
    path: "/user/login", component: () => import("@/pages/user/login")
}, {
    path: "/", component: () => import("@/components/nav-menu")
}]

const router = createRouter({
    routes, history: createWebHashHistory()
})

// eslint-disable-next-line no-unused-vars
router.beforeEach((to, from, next) => {
    // 判断是否登录
    if (to.meta.requiresAuth) {
        if (!userModule.state.token) {
            router.push("/user/login")
        }
        // 验证token有效性

    }
    next()
})

export default router