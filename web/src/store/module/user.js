import storageService from "@/service/storageService";

const userModule = {
    namespaced: true, state: {
        token: storageService.get(storageService.USER_TOKEN),
        username: storageService.get(storageService.USER_INFO),
    }, mutations: {
        SET_TOKEN(state, token) {
            // 更新本地缓存
            storageService.set(storageService.USER_TOKEN, token)
            // 更新state
            state.token = token
        }, SET_USERINFO(state, userInfo) {
            // 更新本地缓存
            storageService.set(storageService.USER_INFO, userInfo)
            // 更新state
            state.username = userInfo
        },
    }, getters: {}
}
export default userModule;