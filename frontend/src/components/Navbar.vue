<template>
  <nav class="bg-gray-800 text-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex items-center">
          <div class="flex-shrink-0 flex items-center">
            <img class="h-8 w-auto" src="@/assets/logo.svg" alt="Logo" />
            <span class="ml-2 text-xl font-semibold">Admin Panel</span>
          </div>
          <div class="hidden md:flex md:ml-6 md:space-x-4">
            <RouterLink
              v-for="item in navItems"
              :key="item.name"
              :to="item.to"
              class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700"
              :class="{ 'bg-gray-900': item.current }"
            >
              {{ item.name }}
            </RouterLink>
          </div>
        </div>
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <button
              @click="toggleSearch"
              class="p-1 rounded-full hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
            >
              <span class="sr-only">Search</span>
              <svg
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
              </svg>
            </button>
          </div>
          <div class="ml-3 relative">
            <div
              v-if="isUserMenuOpen"
              class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5"
            >
              <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >Your Profile</a
              >
              <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >Settings</a
              >
              <a href="#" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >Sign out</a
              >
            </div>
          </div>
          <div class="ml-3 md:hidden">
            <button
              @click="toggleMenu"
              class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
            >
              <span class="sr-only">Open main menu</span>
              <svg
                v-if="!isMenuOpen"
                class="block h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
              <svg
                v-else
                class="block h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile menu -->
    <div v-show="isMenuOpen" class="md:hidden">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <RouterLink
          v-for="item in navItems"
          :key="item.name"
          :to="item.to"
          class="block px-3 py-2 rounded-md text-base font-medium text-white hover:bg-gray-700"
          :class="{ 'bg-gray-900': item.current }"
        >
          {{ item.name }}
        </RouterLink>
      </div>
    </div>
    <!-- Search overlay -->
    <div
      v-if="isSearchOpen"
      class="fixed inset-0 z-10 bg-gray-900 bg-opacity-50 flex items-center justify-center"
    >
      <div class="bg-white p-4 rounded-lg shadow-xl w-full max-w-md">
        <div class="flex items-center border-b border-gray-300 pb-2">
          <input type="text" placeholder="Search..." class="flex-grow outline-none text-gray-700" />
          <button @click="toggleSearch" class="ml-2 text-gray-500 hover:text-gray-700">
            <svg
              class="h-5 w-5"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

const isMenuOpen = ref(false)
const isUserMenuOpen = ref(false)
const isSearchOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const toggleUserMenu = () => {
  isUserMenuOpen.value = !isUserMenuOpen.value
}

const toggleSearch = () => {
  isSearchOpen.value = !isSearchOpen.value
}

const navItems = computed(() => [
  { name: 'Dashboard', to: '/', current: route.path === '/' },
  { name: 'Analytics', to: '/analytics', current: route.path === '/analytics' },
  { name: 'Roles & Permissions', to: '/users', current: route.path === '/users' },
  { name: 'Activity Log', to: '/activity-log', current: route.path === '/activity-log' },
  { name: 'Reports', to: '/reports', current: route.path === '/reports' },
  { name: 'System Health', to: '/system-health', current: route.path === '/system-health' },
  { name: 'Settings', to: '/settings', current: route.path === '/settings' },
  { name: 'Log In', to: '/auth', current: route.path === '/auth' }
])
</script>
