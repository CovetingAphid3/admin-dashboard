<template>
  <div class="lg:col-span-2 bg-white rounded-xl shadow-md p-6">
    <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
      <i class="fas fa-history mr-2 text-blue-600"></i> Recent Activity
    </h2>
    <div class="overflow-hidden">
      <ul class="divide-y divide-gray-200">
        <li
          v-for="activity in recentActivities"
          :key="activity.id"
          class="py-4 flex items-center space-x-4 hover:bg-gray-50 transition duration-150 ease-in-out"
        >
          <div
            :class="[
              'flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center',
              activity.type === 'user'
                ? 'bg-blue-100 text-blue-600'
                : 'bg-green-100 text-green-600'
            ]"
          >
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
    <button
      class="mt-4 w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition duration-150 ease-in-out"
    >
      View All Activity
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, PropType } from 'vue'

const props = defineProps({
  recentActivities: {
    type: Array as PropType<Array<{
      id: number
      type: string
      icon: string
      description: string
      time: string
    }>>,
    required: true
  }
})
</script>

