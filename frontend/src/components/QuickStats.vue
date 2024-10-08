<template>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
    <div
      v-for="stat in quickStats"
      :key="stat.title"
      class="bg-white rounded-xl shadow-md p-6 transition duration-300 ease-in-out transform hover:scale-105"
    >
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
</template>

<script setup lang="ts">
import { ref, PropType } from 'vue'

const props = defineProps({
  quickStats: {
    type: Array as PropType<
      Array<{
        title: string
        value: string
        change: string
        icon: string
        iconColor: string
        valueColor: string
        changeColor: string
        changeIcon: string
      }>
    >,
    required: true
  }
})
</script>
