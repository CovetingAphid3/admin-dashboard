<template>
  <div class="bg-gray-100 min-h-screen">
    <div class="container mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold text-gray-800 mb-8">Dashboard Overview</h1>

      <!-- Quick Stats -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div v-for="stat in quickStats" :key="stat.title" 
             class="bg-white rounded-xl shadow-md p-6 transition duration-300 ease-in-out transform hover:scale-105">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-lg font-semibold text-gray-700">{{ stat.title }}</h2>
            <span :class="['text-2xl', stat.iconColor]"><i :class="stat.icon"></i></span>
          </div>
          <p class="text-3xl font-bold" :class="stat.valueColor">{{ stat.value }}</p>
          <p :class="['text-sm mt-2', stat.changeColor]">
            <i :class="[stat.changeIcon, 'mr-1']"></i>
            {{ stat.change }}
          </p>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Recent Activity -->
        <div class="lg:col-span-2 bg-white rounded-xl shadow-md p-6">
          <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
            <i class="fas fa-history mr-2 text-blue-600"></i> Recent Activity
          </h2>
          <div class="overflow-hidden">
            <ul class="divide-y divide-gray-200">
              <li v-for="activity in recentActivities" :key="activity.id" class="py-4 flex items-center space-x-4 hover:bg-gray-50 transition duration-150 ease-in-out">
                <div :class="[
                  'flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center',
                  activity.type === 'user' ? 'bg-blue-100 text-blue-600' : 'bg-green-100 text-green-600'
                ]">
                  <i :class="activity.icon"></i>
                </div>
                <div class="flex-grow">
                  <p class="text-sm font-medium text-gray-900">{{ activity.description }}</p>
                  <p class="text-xs text-gray-500">{{ activity.time }}</p>
                </div>
                <div class="flex-shrink-0">
                  <button class="text-sm text-blue-600 hover:text-blue-800">View</button>
                </div>
              </li>
            </ul>
          </div>
          <button class="mt-4 w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition duration-150 ease-in-out">
            View All Activity
          </button>
        </div>

        <!-- Quick Actions -->
        <div class="bg-white rounded-xl shadow-md p-6">
          <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
            <i class="fas fa-bolt mr-2 text-yellow-500"></i> Quick Actions
          </h2>
          <div class="grid grid-cols-2 gap-4">
            <button v-for="action in quickActions" :key="action.title" 
                    class="flex flex-col items-center justify-center p-4 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors duration-200">
              <i :class="[action.icon, 'text-2xl mb-2', action.iconColor]"></i>
              <span class="text-sm font-medium text-gray-700">{{ action.title }}</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Recent Transactions -->
      <div class="mt-8 bg-white rounded-xl shadow-md p-6">
        <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
          <i class="fas fa-exchange-alt mr-2 text-purple-600"></i> Recent Transactions
        </h2>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">User</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Amount</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="transaction in recentTransactions" :key="transaction.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#{{ transaction.id }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ transaction.user }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ transaction.amount }}</td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[
                    'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                    transaction.status === 'Completed' ? 'bg-green-100 text-green-800' : 
                    transaction.status === 'Pending' ? 'bg-yellow-100 text-yellow-800' : 'bg-red-100 text-red-800'
                  ]">
                    {{ transaction.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ transaction.date }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const quickStats = ref([
  { title: 'Total Users', value: '1,234', change: '+5.2% from last week', icon: 'fas fa-users', iconColor: 'text-blue-500', valueColor: 'text-blue-600', changeColor: 'text-green-500', changeIcon: 'fas fa-arrow-up' },
  { title: 'Revenue', value: '$12,345', change: '+3.8% from last month', icon: 'fas fa-dollar-sign', iconColor: 'text-green-500', valueColor: 'text-green-600', changeColor: 'text-green-500', changeIcon: 'fas fa-arrow-up' },
  { title: 'Active Projects', value: '42', change: 'No change', icon: 'fas fa-project-diagram', iconColor: 'text-purple-500', valueColor: 'text-purple-600', changeColor: 'text-gray-500', changeIcon: 'fas fa-minus' },
  { title: 'Support Tickets', value: '18', change: '-2 from yesterday', icon: 'fas fa-ticket-alt', iconColor: 'text-red-500', valueColor: 'text-red-600', changeColor: 'text-red-500', changeIcon: 'fas fa-arrow-down' },
]);

const recentActivities = ref([
  { id: 1, type: 'user', icon: 'fas fa-user', description: 'New user registered', time: '5 minutes ago' },
  { id: 2, type: 'order', icon: 'fas fa-shopping-cart', description: 'New order placed', time: '15 minutes ago' },
  { id: 3, type: 'user', icon: 'fas fa-user-edit', description: 'User profile updated', time: '1 hour ago' },
  { id: 4, type: 'order', icon: 'fas fa-truck', description: 'Order #1234 shipped', time: '2 hours ago' },
]);

const quickActions = ref([
  { title: 'Add User', icon: 'fas fa-user-plus', iconColor: 'text-blue-500' },
  { title: 'Create Invoice', icon: 'fas fa-file-invoice-dollar', iconColor: 'text-green-500' },
  { title: 'View Reports', icon: 'fas fa-chart-bar', iconColor: 'text-purple-500' },
  { title: 'Manage Settings', icon: 'fas fa-cog', iconColor: 'text-gray-500' },
]);

const recentTransactions = ref([
  { id: '1001', user: 'John Doe', amount: '$120.00', status: 'Completed', date: '2023-05-01' },
  { id: '1002', user: 'Jane Smith', amount: '$75.50', status: 'Pending', date: '2023-05-02' },
  { id: '1003', user: 'Bob Johnson', amount: '$200.00', status: 'Completed', date: '2023-05-03' },
  { id: '1004', user: 'Alice Brown', amount: '$50.00', status: 'Failed', date: '2023-05-04' },
  { id: '1005', user: 'Charlie Wilson', amount: '$180.00', status: 'Completed', date: '2023-05-05' },
]);
</script>
