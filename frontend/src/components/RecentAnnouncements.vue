<template>
  <div class="bg-white rounded-xl shadow-md p-6">
    <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center">
      <i class="fas fa-bullhorn mr-2 text-yellow-500"></i> Recent Announcements
    </h2>
    <div class="overflow-hidden">
      <ul class="divide-y divide-gray-200">
        <li
          v-for="announcement in announcements"
          :key="announcement.id"
          class="py-4 flex items-center space-x-4 hover:bg-gray-50 transition duration-150 ease-in-out"
        >
          <div class="flex-grow">
            <p class="text-sm font-medium text-gray-900">{{ announcement.title }}</p>
            <p class="text-xs text-gray-500">{{ announcement.body }}</p>
            <p class="text-xs text-gray-400">{{ formatDate(announcement.start_date) }}</p>
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
      View All Announcements
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'

// Declare announcements as a ref so Vue can track reactivity
const announcements = ref([])

// Function to format the date
const formatDate = (dateString: string) => {
  const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' }
  return new Date(dateString).toLocaleDateString(undefined, options)
}

// Function to fetch announcements from the API
const fetchAnnouncements = async () => {
  try {
    const response = await axios.get('http://localhost:8000/announcements')
    // Assign the response data to announcements.value to maintain reactivity
    announcements.value = response.data
    console.log(announcements.value)
  } catch (error) {
    console.error('Error fetching announcements:', error)
  }
}

// Fetch announcements when the component is created
fetchAnnouncements()
</script>

<style scoped>
/* Add any scoped styles here */
</style>
