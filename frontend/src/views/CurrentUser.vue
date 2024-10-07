<template>
  <div v-if="user">
    <h2>Welcome, {{ user.first_name }}!</h2>
    <p>Email: {{ user.email }}</p>
  </div>
  <div v-else>
    <p>Loading user data...</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: null, // To store the user data
      error: null, // To store any error message
    };
  },
  created() {
    this.fetchCurrentUser();
  },
  methods: {
    async fetchCurrentUser() {
      try {
        const response = await fetch("http://localhost:8000/users/me", {
          credentials: "include", // Send the cookie with the request
        
        });
        
        if (!response.ok) {
          throw new Error("Failed to fetch user data.");
        }

        const data = await response.json();
        this.user = data; // Store the fetched user data
      } catch (error) {
        this.error = error.message;
      }
    },
  },
};
</script>

<style scoped>
h2 {
  color: #2c3e50;
}
</style>

