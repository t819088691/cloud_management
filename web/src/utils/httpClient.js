import axios from "axios"
import Qs from "qs"

export default {
    get (uri, params, next) {
        axios.get(uri, {params}).then(res => {
            next(res.data.code, res.data.data, res.data.msg)
        })
    },
    post (uri, data, next) {
        axios.post(uri, Qs.stringify(data)).then(res => {
            next(res.data.code, res.data.data, res.data.msg)
        })
    }
}