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
import { quickStats, recentActivities, quickActions, recentTransactions } from '../constants/index'
</script>
