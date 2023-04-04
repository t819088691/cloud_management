<template>
  <div class="card_warp">
    <el-card>
      <el-form ref="ruleForm" :model="data.form" :rules="data.rules" status-icon>
        <el-form-item label="用户名" label-width="70px" prop="username">
          <el-input v-model="data.form.username" placeholder="请输入邮箱"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="70px" prop="password">
          <el-input v-model="data.form.password" placeholder="请输入密码" type="password"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" label-width="70px" prop="checkPass">
          <el-input v-model="data.form.checkPass" placeholder="请再次输入密码" type="password"></el-input>
        </el-form-item>
        <div class="btns">
          <el-button @click="toLogin">登录</el-button>
          <el-button @click="registry">注册</el-button>
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

export default {
  name: "user-registry",
  setup() {
    const ruleForm = ref(null)

    const methods = {
      registry: function () {
        ruleForm.value.validate((valid) => {
          if (valid) {
            var formdata = qs.stringify({
              "username": data.form.username,
              "password": data.form.password,
              "checkPassword": data.form.checkPass
            })
            userService.register(formdata).then((res) => {
              if (res.data.code !== 200){
                ElMessage.error(res.data.message)
                return
                // console.log(res.data)
              }
              ElMessage.success(res.data.message)
              router.push("/user/login")
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
        if (value !== data.form.checkPass) {
          callback(new Error("密码不匹配"))
        }
        if (value.length < 8) {
          callback(new Error("密码不得小于8位"))
        }
        callback()
      },
      validateCheckPass: (ref, value, callback) => {
        if (value === "") {
          callback(new Error("密码不能为空"))
        }
        if (value !== data.form.password) {
          callback(new Error("密码不匹配"))
        }
        if (value.length < 8) {
          callback(new Error("密码不得小于8位"))
        }
        callback()
      },
    }

    const data = reactive({
      form: {
        username: "",
        password: "",
        checkPass: "",
        code: "",
      },
      rules: {
        username: [{validator: methods.validateUsername, trigger: ["blur", "change"]}],
        password: [{validator: methods.validatePassword, trigger: ["blur", "change"]}],
        checkPass: [{validator: methods.validateCheckPass, trigger: ["blur", "change"]}],
      }
    })
    return {
      ...methods,
      ruleForm,
      data,
    }
  },
  methods: {
    toLogin: function () {
      router.push("/user/login")
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