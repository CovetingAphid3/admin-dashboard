<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <h1 class="text-3xl font-bold text-gray-800 mb-6">Settings</h1>

    <div class="bg-white rounded-lg shadow-md p-6">
      <h2 class="text-xl font-semibold text-gray-700 mb-4">Profile Settings</h2>
      <form @submit.prevent="updateProfile">
        <div class="mb-4">
          <label for="username" class="block text-sm font-medium text-gray-600 mb-1"
            >Username</label
          >
          <input
            type="text"
            id="username"
            v-model="user.username"
            class="border rounded-lg p-2 w-full shadow-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label for="email" class="block text-sm font-medium text-gray-600 mb-1">Email</label>
          <input
            type="email"
            id="email"
            v-model="user.email"
            class="border rounded-lg p-2 w-full shadow-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label for="password" class="block text-sm font-medium text-gray-600 mb-1"
            >Password</label
          >
          <input
            type="password"
            id="password"
            v-model="user.password"
            class="border rounded-lg p-2 w-full shadow-sm"
            required
          />
        </div>

        <button
          type="submit"
          class="bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition duration-150 ease-in-out"
        >
          Update Profile
        </button>
      </form>
    </div>

    <div class="bg-white rounded-lg shadow-md p-6 mt-6">
      <h2 class="text-xl font-semibold text-gray-700 mb-4">Notification Preferences</h2>
      <div class="flex items-center mb-4">
        <input type="checkbox" id="emailNotifications" v-model="notifications.email" class="mr-2" />
        <label for="emailNotifications" class="text-gray-600">Email Notifications</label>
      </div>
      <div class="flex items-center mb-4">
        <input type="checkbox" id="smsNotifications" v-model="notifications.sms" class="mr-2" />
        <label for="smsNotifications" class="text-gray-600">SMS Notifications</label>
      </div>
      <div class="flex items-center mb-4">
        <input type="checkbox" id="pushNotifications" v-model="notifications.push" class="mr-2" />
        <label for="pushNotifications" class="text-gray-600">Push Notifications</label>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow-md p-6 mt-6">
      <h2 class="text-xl font-semibold text-gray-700 mb-4">Account Management</h2>
      <button
        @click="deleteAccount"
        class="bg-red-600 text-white py-2 px-4 rounded-lg hover:bg-red-700 transition duration-150 ease-in-out"
      >
        <i class="fas fa-user-slash mr-2"></i> Delete Account
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const user = ref({
  username: 'JohnDoe',
  email: 'john.doe@example.com',
  password: ''
})

const notifications = ref({
  email: true,
  sms: false,
  push: true
})

const updateProfile = () => {
  // Logic to update the user profile
  console.log('Profile updated:', user.value, notifications.value)
}

const deleteAccount = () => {
  // Logic to delete the user account
  console.log('Account deleted')
}
</script>

<style scoped>
/* Additional styling can be added here if needed */
</style>
