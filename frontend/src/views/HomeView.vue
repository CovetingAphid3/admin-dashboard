<template>
  <div class="bg-gray-100 min-h-screen">
    <div class="container mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold text-gray-800 mb-8">Dashboard Overview</h1>

      <!-- Quick Stats -->
      <QuickStats :quickStats="quickStats" />

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Recent Activity -->
        <RecentActivity :recentActivities="recentActivities" />

        <!-- Quick Actions -->
        <QuickActions :quickActions="quickActions" />
      </div>

      <!-- Recent Announcements -->
      <RecentAnnouncements />

      <!-- Recent Transactions -->
      <RecentTransactions :recentTransactions="recentTransactions" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import QuickStats from '../components/QuickStats.vue'
import RecentActivity from '../components/RecentActivity.vue'
import QuickActions from '../components/QuickActions.vue'
import RecentAnnouncements from '../components/RecentAnnouncements.vue'
import RecentTransactions from '../components/RecentTransactions.vue'


const quickStats = ref([
  {
    title: 'Total Users',
    value: '1,234',
    change: '+5.2% from last week',
    icon: 'fas fa-users',
    iconColor: 'text-blue-500',
    valueColor: 'text-blue-600',
    changeColor: 'text-green-500',
    changeIcon: 'fas fa-arrow-up'
  },
  {
    title: 'Revenue',
    value: '$12,345',
    change: '+3.8% from last month',
    icon: 'fas fa-dollar-sign',
    iconColor: 'text-green-500',
    valueColor: 'text-green-600',
    changeColor: 'text-green-500',
    changeIcon: 'fas fa-arrow-up'
  },
  {
    title: 'Active Projects',
    value: '42',
    change: 'No change',
    icon: 'fas fa-project-diagram',
    iconColor: 'text-purple-500',
    valueColor: 'text-purple-600',
    changeColor: 'text-gray-500',
    changeIcon: 'fas fa-minus'
  },
  {
    title: 'Support Tickets',
    value: '18',
    change: '-2 from yesterday',
    icon: 'fas fa-ticket-alt',
    iconColor: 'text-red-500',
    valueColor: 'text-red-600',
    changeColor: 'text-red-500',
    changeIcon: 'fas fa-arrow-down'
  }
])

const recentActivities = ref([
  {
    id: 1,
    type: 'user',
    icon: 'fas fa-user',
    description: 'New user registered',
    time: '5 minutes ago'
  },
  {
    id: 2,
    type: 'order',
    icon: 'fas fa-shopping-cart',
    description: 'New order placed',
    time: '15 minutes ago'
  },
  {
    id: 3,
    type: 'user',
    icon: 'fas fa-user-edit',
    description: 'User profile updated',
    time: '1 hour ago'
  },
  {
    id: 4,
    type: 'order',
    icon: 'fas fa-truck',
    description: 'Order #1234 shipped',
    time: '2 hours ago'
  }
])

const quickActions = ref([
  { title: 'Add User', icon: 'fas fa-user-plus', iconColor: 'text-blue-500' },
  { title: 'Create Invoice', icon: 'fas fa-file-invoice-dollar', iconColor: 'text-green-500' },
  { title: 'View Reports', icon: 'fas fa-chart-bar', iconColor: 'text-purple-500' },
  { title: 'Manage Settings', icon: 'fas fa-cog', iconColor: 'text-gray-500' }
])

const recentTransactions = ref([
  { id: '1001', user: 'John Doe', amount: '$120.00', status: 'Completed', date: '2023-05-01' },
  { id: '1002', user: 'Jane Smith', amount: '$75.50', status: 'Pending', date: '2023-05-02' },
  { id: '1003', user: 'Bob Johnson', amount: '$200.00', status: 'Completed', date: '2023-05-03' },
  { id: '1004', user: 'Alice Brown', amount: '$50.00', status: 'Failed', date: '2023-05-04' },
  { id: '1005', user: 'Charlie Wilson', amount: '$180.00', status: 'Completed', date: '2023-05-05' }
])
</script>
