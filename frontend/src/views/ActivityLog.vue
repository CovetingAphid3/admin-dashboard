<!-- ActivityLog.vue -->
<template>
  <div class="p-6 bg-gray-100">
    <h1 class="text-3xl font-bold mb-6">Activity Log</h1>

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
          <tr v-for="(activity, index) in activities" :key="index">
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

export default defineComponent({
  name: 'ActivityLog',
  setup() {
    const activities = ref([
      { userName: 'John Doe', userEmail: 'john@example.com', userAvatar: 'https://randomuser.me/api/portraits/men/1.jpg', action: 'New Purchase', date: '2024-10-17' },
      { userName: 'Jane Smith', userEmail: 'jane@example.com', userAvatar: 'https://randomuser.me/api/portraits/women/2.jpg', action: 'Account Created', date: '2024-10-16' },
      { userName: 'Bob Johnson', userEmail: 'bob@example.com', userAvatar: 'https://randomuser.me/api/portraits/men/3.jpg', action: 'Subscription Renewed', date: '2024-10-15' },
      { userName: 'Alice Brown', userEmail: 'alice@example.com', userAvatar: 'https://randomuser.me/api/portraits/women/4.jpg', action: 'Profile Updated', date: '2024-10-14' },
      { userName: 'Charlie Black', userEmail: 'charlie@example.com', userAvatar: 'https://randomuser.me/api/portraits/men/5.jpg', action: 'Password Changed', date: '2024-10-13' },
    ]);

    return {
      activities
    };
  }
});
</script>

<style scoped>
/* Add any specific styles you need for the ActivityLog here */
</style>

