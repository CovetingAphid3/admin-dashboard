<!-- LineChart.vue -->
<template>
  <div>
    <Line :data="chartData" :options="options" />
  </div>
</template>

<script>
import { defineComponent } from 'vue';
import { Line } from 'vue-chartjs';
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, LinearScale, CategoryScale, PointElement } from 'chart.js';

ChartJS.register(Title, Tooltip, Legend, LineElement, LinearScale, CategoryScale, PointElement);

export default defineComponent({
  name: 'LineChart',
  components: {
    Line
  },
  props: {
    chartData: {
      type: Object,
      required: true
    },
    options: {
      type: Object,
      default: () => ({})
    }
  }
});
</script>

<style scoped>
/* Add any specific styles you need for the LineChart here */
</style>

