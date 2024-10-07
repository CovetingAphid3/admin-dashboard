<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <h1 class="text-3xl font-bold text-gray-800 mb-6">User Management</h1>

    <!-- Search Bar -->
    <div class="mb-4">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search users..."
        class="border rounded-lg p-2 w-full shadow-md"
      />
    </div>

    <!-- User Table -->
    <div class="overflow-hidden rounded-lg shadow-md">
      <table class="min-w-full bg-white rounded-lg">
        <thead class="bg-gray-200">
          <tr>
            <th
              class="py-3 px-4 text-left text-xs font-medium text-gray-600 uppercase tracking-wider"
            >
              ID
            </th>
            <th
              class="py-3 px-4 text-left text-xs font-medium text-gray-600 uppercase tracking-wider"
            >
              Name
            </th>
            <th
              class="py-3 px-4 text-left text-xs font-medium text-gray-600 uppercase tracking-wider"
            >
              Email
            </th>
            <th
              class="py-3 px-4 text-left text-xs font-medium text-gray-600 uppercase tracking-wider"
            >
              Role
            </th>
            <th
              class="py-3 px-4 text-left text-xs font-medium text-gray-600 uppercase tracking-wider"
            >
              Status
            </th>
            <th
              class="py-3 px-4 text-left text-xs font-medium text-gray-600 uppercase tracking-wider"
            >
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="user in filteredUsers"
            :key="user.ID"
            class="hover:bg-gray-100 transition duration-150 ease-in-out"
          >
            <td class="py-3 px-4">{{ user.ID }}</td>
            <td class="py-3 px-4">{{ user.first_name }}</td>
            <td class="py-3 px-4">{{ user.email }}</td>
            <td class="py-3 px-4">{{ user.role }}</td>
            <td class="py-3 px-4">
              <span :class="user.is_active ? 'text-green-500' : 'text-red-500'">
                {{ user.is_active ? 'Active' : 'Inactive' }}
              </span>
            </td>
            <td class="py-3 px-4 flex space-x-4">
              <button
                @click="editUser(user.ID)"
                class="text-blue-500 hover:text-blue-700 flex items-center"
              >
                <i class="fas fa-edit mr-1"></i> Edit
              </button>
              <button
                @click="deleteUser(user.ID)"
                class="text-red-500 hover:text-red-700 flex items-center"
              >
                <i class="fas fa-trash-alt mr-1"></i> Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="flex justify-between items-center mt-6">
      <span class="text-gray-600">
        Showing {{ (currentPage - 1) * pageSize + 1 }} to
        {{ Math.min(currentPage * pageSize, users.length) }} of {{ users.length }} users
      </span>
      <div>
        <button
          @click="prevPage"
          :disabled="currentPage === 1"
          class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-300 transition duration-150 ease-in-out"
        >
          Previous
        </button>
        <button
          @click="nextPage"
          :disabled="currentPage * pageSize >= users.length"
          class="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg hover:bg-gray-300 transition duration-150 ease-in-out ml-2"
        >
          Next
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const users = ref([]) // Initialize as an empty array
const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = 10 // Adjust the number of items displayed per page

// Fetch users from the backend when the component is mounted
const fetchUsers = async () => {
  try {
    const response = await axios.get('http://localhost:8000/users') // Adjust the URL if necessary
    console.log(response.data)
    users.value = response.data // Assign the fetched data to the users array
  } catch (error) {
    console.error('Error fetching users:', error)
  }
}

// Call fetchUsers when the component is mounted
onMounted(fetchUsers)

const filteredUsers = computed(() => {
  const query = searchQuery.value.toLowerCase()
  return users.value
    .filter(
      (user) =>
        user.first_name.toLowerCase().includes(query) || user.email.toLowerCase().includes(query)
    )
    .slice((currentPage.value - 1) * pageSize, currentPage.value * pageSize)
})

const editUser = (userId: number) => {
  // Logic to edit the user, e.g., navigate to an edit page
  console.log(`Edit user with ID: ${userId}`)
}

const deleteUser = async (userId: number) => {
  try {
    const response = await axios.delete(`http://localhost:8000/users/${userId}`);
    console.log('Delete response:', response); // Log the response
    if (response.status === 200) {
      users.value = users.value.filter((user) => user.id !== userId);
      console.log(`Delete user with ID: ${userId}`);
    }
  } catch (error) {
    console.error('Error deleting user:', error);
  }
}


const nextPage = () => {
  if (currentPage.value * pageSize < users.value.length) {
    currentPage.value++
  }
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border: 1px solid #e2e8f0;
}
</style>
