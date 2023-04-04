<template>
  <div class="card_warp">
    <el-card>
      <el-form ref="ruleForm" :model="data.form" :rules="data.form.rule">
        <el-form-item label="用户名" label-width="70px" prop="username">
          <el-input v-model="data.form.username" placeholder="请输入邮箱"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="70px" prop="password">
          <el-input v-model="data.form.password" placeholder="请输入密码" type="password"></el-input>
        </el-form-item>
        <el-form-item v-show="0" label="验证码" label-width="70px">
          <el-input v-model="data.form.code" placeholder="请输入验证码"></el-input>
        </el-form-item>
        <router-link class="forget" to="/user/forget">忘记密码</router-link>
        <div class="btns">
          <el-button @click="login">登录</el-button>
          <el-button @click="toRegistry">注册</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import qs from "qs"
import router from "@/router";
import userService from "@/service/userService";
import {reactive, ref} from "vue";
import {ElMessage} from "element-plus";
import storageService from "@/service/storageService";
import store from "@/store";

export default {
  name: "user-login",
  setup() {
    const ruleForm = ref(null)

    const methods = {
      login: function () {
        ruleForm.value.validate((valid) => {
          if (valid) {
            var formdata = qs.stringify({
              "username": data.form.username,
              "password": data.form.password
            })
            userService.login(formdata).then(res => {
              if (res.data.code !== 200) {
                ElMessage.error(res.data.message)
                return
              }
              storageService.set(storageService.USER_TOKEN, res.data.data.token)
              storageService.set(storageService.USER_INFO, res.data.data.username)

              store.commit("userModules/SET_TOKEN", res.data.data.token)
              store.commit("userModules/SET_USERINFO", res.data.data.username)
              ElMessage.success(res.data.message)

              router.push("/")
            })
          }
        })
      },
      validateUsername: (ref, value, callback) => {
        if (value === "") {
          callback(new Error("用户名不能为空"))
        }
        const re = /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/
        if (!re.test(value)) {
          callback(new Error("邮箱格式不正确"))
        }
        callback()
      },
      validatePassword: (ref, value, callback) => {
        if (value === "") {
          callback(new Error("密码不能为空"))
        }
      }
    }

    const data = reactive({
      form: {
        username: "",
        password: "",
        code: "",
      },
      rules: {
        username: [{validator: methods.validateUsername, trigger: ["blur"]}],
        password: [{validator: methods.validatePassword, trigger: ["blur"]}],
      }
    })

    return {
      ...methods,
      ruleForm,
      data,
    }
  },
  methods: {
    toRegistry: function () {
      router.push("/user/registry")
    },
  }
}
</script>

<style scoped>
.card_warp {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.el-card {
  width: 350px;
}

.forget {
  font-size: 12px;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}

.btns {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.el-button {
  margin-top: 10px;
  width: 150px;
}
</style>