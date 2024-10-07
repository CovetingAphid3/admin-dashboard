<template>
  <div class="bg-gray-100 min-h-screen p-6">
    <h1 class="text-3xl font-bold text-gray-800 mb-6">Reports</h1>

    <div class="bg-white rounded-lg shadow-md p-6 mb-6">
      <h2 class="text-xl font-semibold text-gray-700 mb-4">Filter Reports</h2>
      <div class="flex mb-4">
        <input
          type="text"
          v-model="filterQuery"
          placeholder="Search reports..."
          class="border rounded-lg p-2 w-full shadow-sm"
        />
        <button
          @click="applyFilter"
          class="ml-2 bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition duration-150 ease-in-out"
        >
          Filter
        </button>
      </div>
    </div>

    <div class="bg-white rounded-lg shadow-md p-6">
      <h2 class="text-xl font-semibold text-gray-700 mb-4">Report List</h2>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">User</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Report Type</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in filteredReports" :key="report.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">#{{ report.id }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ report.user }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ report.type }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ report.date }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <button
                  @click="viewReport(report.id)"
                  class="text-blue-600 hover:text-blue-800"
                >
                  <i class="fas fa-eye mr-1"></i> View
                </button>
                <button
                  @click="deleteReport(report.id)"
                  class="text-red-600 hover:text-red-800 ml-2"
                >
                  <i class="fas fa-trash mr-1"></i> Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

const reports = ref([
  { id: 1, user: 'John Doe', type: 'Sales Report', date: '2024-10-01' },
  { id: 2, user: 'Jane Smith', type: 'Inventory Report', date: '2024-10-02' },
  { id: 3, user: 'Alice Brown', type: 'Expense Report', date: '2024-10-03' },
  { id: 4, user: 'Bob Johnson', type: 'User Activity Report', date: '2024-10-04' },
]);

const filterQuery = ref('');

const filteredReports = computed(() => {
  return reports.value.filter(report => 
    report.user.toLowerCase().includes(filterQuery.value.toLowerCase()) || 
    report.type.toLowerCase().includes(filterQuery.value.toLowerCase())
  );
});

const applyFilter = () => {
  // Logic to apply filter if needed
  console.log('Filter applied:', filterQuery.value);
};

const viewReport = (id: number) => {
  // Logic to view the report details
  console.log('View report with ID:', id);
};

const deleteReport = (id: number) => {
  // Logic to delete the report
  console.log('Delete report with ID:', id);
  reports.value = reports.value.filter(report => report.id !== id);
};
</script>

<style scoped>
/* Additional styling can be added here if needed */
</style>

