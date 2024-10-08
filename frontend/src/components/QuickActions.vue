<template>
  <div class="bg-white rounded-xl shadow-md p-6">
    <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
      <i class="fas fa-bolt mr-2 text-yellow-500"></i> Quick Actions
    </h2>
    <div class="grid grid-cols-2 gap-4">
      <button
        v-for="action in quickActions"
        :key="action.title"
        class="flex flex-col items-center justify-center p-4 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors duration-200"
      >
        <i :class="[action.icon, 'text-2xl mb-2', action.iconColor]"></i>
        <span class="text-sm font-medium text-gray-700">{{ action.title }}</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, PropType } from 'vue'

const props = defineProps({
  quickActions: {
    type: Array as PropType<Array<{
      title: string
      icon: string
      iconColor: string
    }>>,
    required: true
  }
})
</script>

