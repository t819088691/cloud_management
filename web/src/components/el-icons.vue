<template>
  <el-icon :size="size" :color="color">
    <component :is="name"></component>
  </el-icon>
</template>

<script>
import {defineComponent} from "vue";
import * as Icons from "@element-plus/icons-vue"

export default defineComponent({
  name: "el-icons",
  components: Icons,
  props: {
    name: {
      type: String,
      required: true,
    },
    size: {
      type: String,
      default: "",
    },
    color: {
      type: String,
      default: "",
    }
  }
})
</script>
