<template>
  <el-container>
    <el-container>
      <!--   介绍   -->
      <div></div>

      <!--  导航栏  -->
      <el-menu :collapse="isCollapse" :default-active="activeIndex" :router="true"
               background-color="#0F3E7BFF" class="el-menu-vertical"
               text-color="#FFFFFF" @select="handleSelect">
        <template v-for="(item, index) in menuItem" :key="index">
          <!--     多级菜单     -->
          <el-sub-menu v-if="item.childMenu && item.childMenu.length > 0" :index="item.index">
            <template #title>
              <el-icons :name="item.icon"></el-icons>
              <span>{{ item.name }}</span>
            </template>
            <el-menu-item v-for="(child, childIndex) in item.childMenu" :key="childIndex" :index="child.url">
              <span>{{ child.name }}</span>
            </el-menu-item>
          </el-sub-menu>

          <!--     一级菜单     -->
          <el-menu-item v-if="!item.childMenu" :index="item.url" default-active="$route.path">
            <el-icons :name="item.icon"/>
            <span>{{ item.name }}</span>
          </el-menu-item>
        </template>
      </el-menu>

      <!--  主窗口  -->
      <el-container>
        <el-header>
          <el-icons class="collapse-btn" :name="collapseIcon" size="26px" @click="collapse"></el-icons>
          <div class="user">
            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
            <span>{{ userInfo.username }}</span>
          </div>
        </el-header>
        <el-main>
          <router-view name="dashboard"></router-view>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>

  </el-container>
</template>


<script>
import elIcons from "@/components/el-icons";
import router from "@/router";
import userModule from "@/store/module/user";

export default {
  name: "nav-menu",
  data() {
    return {
      collapseIcon: "expand",
      isCollapse: true,
      activeIndex: "/",
      menuItem: [{
        name: "结算",
        icon: "menu",
        index: "1-1",
        childMenu: [{
          name: "汇总",
          url: "/dashboard/billing/index"
        }, {
          name: "GCP",
          url: "/dashboard/billing/gcp"
        }, {
          name: "AWS",
          url: "/dashboard/billing/aws"
        }, {
          name: "Azure",
          url: "/dashboard/billing/azure"
        }]
      }, {
        name: "选项二",
        icon: "menu",
        url: "/dashboard/billing/gcp"
      }, {
        name: "选项三",
        icon: "menu",
        url: "/dashboard/billing/aws"
      }, {
        name: "选项四",
        icon: "menu",
        url: "/dashboard/billing/azure"
      }, {
        name: "平台管理",
        icon: "setting",
        index: "2-1",
        childMenu: [{
          name: "用户列表",
          is_admin: true,
          url: "/dashboard/user/list"
        }]
      }],
      userInfo: {
        avatar: "",
        username: userModule.state.username,
      }
    };
  },
  components: {elIcons},
  methods: {
    collapse() {
      if (!this.isCollapse) {
        this.isCollapse = true;
        this.collapseIcon = "expand"
      } else {
        this.isCollapse = false;
        this.collapseIcon = "fold"
      }
    },

    // eslint-disable-next-line no-unused-vars
    handleSelect(routePath) {
      this.activeIndex = router.path
    }
  }

}
</script>

<style scoped>
.el-header, el-footer {
  height: 50px;
  background-color: #ffffff;
  color: #333;
  line-height: 50px;
}

.el-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.el-header > .collapse-btn {
  cursor: pointer;
}

.el-header > .user {
//margin-right: 20px; cursor: pointer; display: flex;
  justify-items: center;
  align-items: center;
}

.el-menu-vertical {
  height: 100vh;
}

.el-menu-vertical:not(.el-menu--collapse) {
  width: 250px;

}


.el-main {
  background-color: #E9EEF3;
  color: #333;
}

body > .el-container {
  margin-bottom: 40px;
}

</style>