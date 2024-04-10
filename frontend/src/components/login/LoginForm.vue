<script>
import FormInput from "@/components/reusable/FormInput.vue";
import Button from '@/components/reusable/Button.vue';
import axios from "axios";

export default {
  components: {
    FormInput,
    Button
  },
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async submit() {
      try {
        console.log('Form submitted:', this.username, this.password);
        const formData = new FormData();
        formData.append('username', this.username);
        formData.append('password', this.password);

        const response = await axios.post('http://localhost:8080/login', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });
        localStorage.setItem('jwt', response.data.token); // Store the token or auth status
        this.$router.push('/admin/dashboard');
      } catch (error) {
        console.error('Login failed:', error);
      }
    }
  }
};
</script>

<template>
  <div class="max-w-md mx-auto mt-10 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
    <form class="space-y-6" @submit.prevent="submit">
      <FormInput
          label="Username"
          inputIdentifier="username"
          inputType="username"
          v-model:value="username"
          placeholder="Enter your username"
      />
      <FormInput
          label="Password"
          inputIdentifier="password"
          v-model:value="password"
          inputType="password"
          placeholder="Enter your password"
      />
      <Button
          title="Log In"
          type="submit"
          class="w-full bg-primary-dark text-white font-bold py-2 px-4 rounded hover:bg-secondary-dark transition-colors"
      />
    </form>
  </div>
</template>


<style scoped>

</style>