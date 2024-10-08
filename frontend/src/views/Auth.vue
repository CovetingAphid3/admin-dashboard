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
import axios from 'axios'

const isLogin = ref(true)
const email = ref('')
const password = ref('')
const firstName = ref('')
const lastName = ref('')

const handleSubmit = async () => {
  const url = isLogin.value
    ? 'http://localhost:8000/users/login'
    : 'http://localhost:8000/users/signup'
  const body = isLogin.value
    ? { Email: email.value, Password: password.value }
    : {
        Email: email.value,
        Password: password.value,
        FirstName: firstName.value,
        LastName: lastName.value
      }

  try {
    const response = await axios.post(url, body)
    console.log(response.data) // Handle successful response (e.g., redirect or show message)
  } catch (error) {
    console.error('Error during login/signup:', error.response.data)
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
