<template>
  <div class="p-6 bg-gray-100">
    <h1 class="text-3xl font-bold mb-6">Analytics Dashboard</h1>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div
        v-for="(metric, index) in summaryMetrics"
        :key="index"
        class="bg-white rounded-lg shadow p-6"
      >
        <h3 class="text-lg font-semibold text-gray-700">{{ metric.name }}</h3>
        <p class="text-3xl font-bold mt-2">{{ metric.value }}</p>
        <p
          :class="['text-sm mt-2', metric.change > 0 ? 'text-green-600' : 'text-red-600']"
        >
          {{ metric.change > 0 ? '↑' : '↓' }} {{ Math.abs(metric.change) }}% from last week
        </p>
      </div>
    </div>

    <!-- Chart -->
    <div class="bg-white rounded-lg shadow p-6 mb-8">
      <h2 class="text-xl font-semibold mb-4">Revenue Over Time</h2>
      <div class="h-64">
        <LineChart :chartData="chartData" :options="chartOptions" />
      </div>
    </div>

    <!-- Recent Activity Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <h2 class="text-xl font-semibold p-6 bg-gray-50 border-b">Recent Activity</h2>
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">User</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Action</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="(activity, index) in recentActivity" :key="index">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="flex-shrink-0 h-10 w-10">
                  <img class="h-10 w-10 rounded-full" :src="activity.userAvatar" alt="" />
                </div>
                <div class="ml-4">
                  <div class="text-sm font-medium text-gray-900">{{ activity.userName }}</div>
                  <div class="text-sm text-gray-500">{{ activity.userEmail }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                {{ activity.action }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ activity.date }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { ref, defineComponent } from 'vue';
import LineChart from './LineChart.vue'; // Import the LineChart component

export default defineComponent({
  name: 'Analytics',
  components: {
    LineChart // Register the LineChart component
  },
  setup() {
    const summaryMetrics = ref([
      { name: 'Total Users', value: '10,492', change: 12.5 },
      { name: 'Active Users', value: '8,123', change: 8.2 },
      { name: 'Revenue', value: '$284,391', change: 15.3 },
      { name: 'Conversion Rate', value: '3.6%', change: -2.1 }
    ]);

    const chartData = ref({
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
      datasets: [{
        label: 'Revenue',
        data: [30000, 40000, 35000, 50000, 49000, 60000],
        borderColor: '#3b82f6',
        backgroundColor: 'rgba(59, 130, 246, 0.5)',
      }]
    });

    const chartOptions = {
      responsive: true,
      maintainAspectRatio: false
    };

    const recentActivity = ref([
      { userName: 'John Doe', userEmail: 'john@example.com', userAvatar: 'https://randomuser.me/api/portraits/men/1.jpg', action: 'New Purchase', date: '2023-06-15' },
      { userName: 'Jane Smith', userEmail: 'jane@example.com', userAvatar: 'https://randomuser.me/api/portraits/women/2.jpg', action: 'Account Created', date: '2023-06-14' },
      { userName: 'Bob Johnson', userEmail: 'bob@example.com', userAvatar: 'https://randomuser.me/api/portraits/men/3.jpg', action: 'Subscription Renewed', date: '2023-06-13' },
    ]);

    return {
      summaryMetrics,
      chartData,
      chartOptions,
      recentActivity
    };
  }
});
</script>

<style scoped>
/* Add any specific styles you need here */
</style>

