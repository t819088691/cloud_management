import request from "@/utils/request";

const methods = {
    // 用户登录
    login: ((formdata) => {
        return request.post("/auth/login", formdata)
    }),

    // 用户注册
    register: ((formdata) => {
        return request.post("auth/register", formdata)
    }),

    // 用户忘记密码
    forget: ((formdata) => {
        return request.post("auth/forget", formdata)
    }),

    // 列出用户列表
    listUser: (() => {
        return request.get("/user/list")
    }),

    // 禁用用户
    disableUser: (() => {
        return request.post()
    }),

    // 启用用户
    enableUser: (()=> {
        return request.post()
    })
}

export default {
    ...methods
}