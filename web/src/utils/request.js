import axios from 'axios'
import userModule from "@/store/module/user";

const service = axios.create({
    baseURL: "http://192.168.1.107/api/",
    timeout: 1000*5,
})

service.interceptors.request.use((config) => {
    Object.assign(config.headers, {Authorization: userModule.state.token})
    return config;
}, (error) => {
    return Promise.reject(error);
});

export default service;