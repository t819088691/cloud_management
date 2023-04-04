<template>
  <div class="gcp table">
    <monthly-table></monthly-table>
  </div>
</template>

<script>
import monthlyTable from "@/units/billing/monthlyTable";

export default {
  name: "gcp-billing",
  components: {
    monthlyTable
  },
}
</script>

<style scoped>
.gcp .table {
  height: 50vh;
}
</style>