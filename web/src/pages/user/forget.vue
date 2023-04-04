<template>
  <div class="card_warp">
    <el-card>
      <el-form ref="ruleForm" :model="data.form" :rules="data.rules" status-icon>
        <el-form-item label="用户名" label-width="70px" prop="username">
          <el-input v-model="data.form.username" autocomplete="off" placeholder="请输入邮箱"></el-input>
        </el-form-item>
        <el-form-item label="新密码" label-width="70px" prop="password">
          <el-input v-model="data.form.password" placeholder="请输入密码" type="password"></el-input>
        </el-form-item>
        <div class="btns">
          <el-button @click="forget">确定</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>
// import axios from "axios"
import qs from "qs"
import router from "@/router";
import {reactive, ref} from "vue";
import userService from "@/service/userService";

export default {
  name: "user-forget",
  setup() {
    const ruleForm = ref(null)

    const methods = {
      forget: function () {
        ruleForm.value.validate((valid) => {
          if (valid) {
            var formdata = qs.stringify({
              "username": this.data.form.username,
              "password": this.data.form.password
            })

            userService.forget(formdata).then((res) => {
                  console.log(res.data)
                }
            )
            router.push("/user/login")
          }
        })


      },
      validateUsername: (rule, value, callback) => {
        if (value === "") {
          callback(new Error("用户名不能为空"))
        }
        callback()
      },
      validatePassword: (rule, value, callback) => {
        if (value === "") {
          callback(new Error("新密码不能为空"))
        }
        if (value.length < 8) {
          callback(new Error("新密码不得小于8位"))
        }
        callback()
      }
    }

    const data = reactive({
      form: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          {validator: methods.validateUsername, trigger: ["blur", "change"]}
        ],
        password: [
          {validator: methods.validatePassword, trigger: ["blur", "change"]}
        ]
      }
    })

    return {
      ...methods,
      ruleForm,
      data
    }
  },

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
  width: 350px;
}
</style>