<template>
  <div class="bg-gray-100 min-h-screen p-6 flex justify-center items-center">
    <div class="bg-white rounded-lg shadow-md p-8 w-96">
      <h2 class="text-2xl font-bold text-gray-800 mb-6">{{ isLogin ? 'Login' : 'Signup' }}</h2>
      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <input
            v-model="email"
            type="email"
            placeholder="Email"
            class="border rounded-lg p-2 w-full shadow-md"
            required
          />
        </div>

        <div class="mb-4">
          <input
            v-model="password"
            type="password"
            placeholder="Password"
            class="border rounded-lg p-2 w-full shadow-md"
            required
          />
        </div>

        <div v-if="!isLogin">
          <div class="mb-4">
            <input
              v-model="firstName"
              type="text"
              placeholder="First Name"
              class="border rounded-lg p-2 w-full shadow-md"
              required
            />
          </div>
          <div class="mb-4">
            <input
              v-model="lastName"
              type="text"
              placeholder="Last Name"
              class="border rounded-lg p-2 w-full shadow-md"
              required
            />
          </div>
        </div>

        <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded-lg w-full">
          {{ isLogin ? 'Login' : 'Signup' }}
        </button>
      </form>

      <p class="mt-4 text-gray-600">
        {{ isLogin ? "Don't have an account?" : 'Already have an account?' }}
        <button @click="toggleAuth" class="text-blue-500 underline">
          {{ isLogin ? 'Signup' : 'Login' }}
        </button>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const isLogin = ref(true)
const email = ref('')
const password = ref('')
const firstName = ref('')
const lastName = ref('')

// Mock user data for login/signup
const mockUsers = ref([
  { email: 'testuser@example.com', password: 'password123', firstName: 'Test', lastName: 'User' }
])

const handleSubmit = () => {
  if (isLogin.value) {
    // Mock login
    const user = mockUsers.value.find(
      (user) => user.email === email.value && user.password === password.value
    )
    if (user) {
      console.log('Login successful:', user)
      // Simulate redirect or other post-login actions
    } else {
      console.error('Invalid credentials')
    }
  } else {
    // Mock signup (ensure user doesn't already exist)
    const existingUser = mockUsers.value.find((user) => user.email === email.value)
    if (existingUser) {
      console.error('User already exists')
    } else {
      const newUser = {
        email: email.value,
        password: password.value,
        firstName: firstName.value,
        lastName: lastName.value
      }
      mockUsers.value.push(newUser)
      console.log('Signup successful:', newUser)
      // Simulate redirect or other post-signup actions
    }
  }
}

const toggleAuth = () => {
  isLogin.value = !isLogin.value
}
</script>

<style scoped>
input:focus {
  outline: none;
  border-color: #3182ce;
}
</style>

