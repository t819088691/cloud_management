<template>
  <el-card>
    <el-auto-resizer>
      <template #default="{ height, width }">
        <el-table-v2
            :columns="columns"
            :data="data"
            :height="height"
            :width="width"
            fixed
        />
      </template>
    </el-auto-resizer>
  </el-card>
</template>

<script>
export default {
  name: "monthlyTable",
  data() {
    return {
      year: "",
      month: "",
      columns: [
        {dataKey: "column-1", key: "column-1", title: "column 1", width: 150},
        {dataKey: "column-2", key: "column-2", title: "column 2", width: 150},
        {dataKey: "column-3", key: "column-3", title: "column 3", width: 150},
        {dataKey: "column-4", key: "column-4", title: "column 4", width: 150},
        {dataKey: "column-5", key: "column-5", title: "column 5", width: 150},
      ],
      data: [
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
        {"column-1": "test1", "column-2": "test2", "column-3": "test3", "column-4": "test4", "column-5": "test5"},
      ]
    }
  },
  created() {

  },
  methods: {}
}
</script>

<style scoped>
.el-table-v2 {
  height: 400px;
}
</style>