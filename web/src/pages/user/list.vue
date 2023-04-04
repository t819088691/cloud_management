<template>
  <el-card>
    <el-table :data="userList" max-height="500">
      <el-table-column label="用户id" prop="id" width="80"/>
      <el-table-column label="用户名" prop="username" width="200"/>
      <el-table-column label="用户状态" prop="status" width="80">
        <template #default="scope">
          <el-tag :type="scope.row.status ? 'success' : 'danger' ">{{ scope.row.status ? "已启用" : "已禁止" }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="最后一次登录时间" prop="created_at" width="300"/>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button @click="enableUser(scope.$index, scope.row)" type="success" text bg :disabled="scope.row.status ? true : false">启用</el-button>
          <el-button @click="disableUser(scope.$index, scope.row)" type="danger" text bg :disabled="!scope.row.status ? true : false">禁用</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import userService from "@/service/userService";

export default {
  name: "user-list",
  data() {
    return {
      timer: 5000,
      userList: []
    }
  },
  methods: {
    enableUser: function (index, row){
      console.log(index, row)
    },
    disableUser: function (index, row){
      console.log(index, row)
    },
  },
  mounted() {
    setInterval(() => {
      userService.listUser().then((res) => {
        this.userList = res.data.data
      })
    }, this.timer)
  },
  beforeUnmount() {
    clearInterval(this.timer)
    this.timer = null
  }
}
</script>

<style scoped>
.el-card {
  width: 854px;
}
</style>