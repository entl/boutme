<script>
import AddProjectModal from "@/components/projects/protected/AddProjectModal.vue";
import Button from '@/components/reusable/Button.vue';
import feather from "feather-icons";
import ProjectSingle from "@/components/projects/protected/ProjectSingle.vue";
import {getProjects} from "@/data/projects";

export default {
  components: {
    ProjectSingle,
    AddProjectModal,
    Button
  },
  data() {
    return {
      isOpen: false,
      modal: false,
      projects: null,
    };
  },
  created() {
    this.theme = localStorage.getItem('theme') || 'light';
  },
  async mounted() {
    feather.replace();
    this.theme = localStorage.getItem('theme') || 'light';
    this.projects = await getProjects();
    feather.replace();
  },
  methods: {
    updateTheme(theme) {
      this.theme = theme;
    },
    showModal() {
      if (this.modal) {
        // Stop screen scrolling
        document
            .getElementsByTagName('html')[0]
            .classList.remove('overflow-y-hidden');
        this.modal = false;
      } else {
        document
            .getElementsByTagName('html')[0]
            .classList.add('overflow-y-hidden');
        this.modal = true;
      }
    },
  },
  updated() {
    feather.replace();
  },
};
</script>

<template>
  <!-- Header section with title and add button -->
  <div class="container mx-auto p-4 flex justify-end items-center">


    <Button
        title="+"
        class="flex items-center justify-center bg-indigo-500 hover:bg-indigo-600 text-white shadow-sm rounded-md p-2 duration-300 font-bold text-2xl sm:text-3xl w-8 h-8 sm:w-10 sm:h-10"
        @click="showModal"
        aria-label="Add New Project"
    >
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 5l7 7-7 7-7-7"/></svg>
    </Button>
  </div>

  <!-- Projects grid -->
  <div class="container mx-auto">
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 mt-6 sm:gap-10">
      <ProjectSingle
          v-for="project in projects"
          :key="project.id"
          :project="project"
      />
    </div>
  </div>

  <AddProjectModal :showModal="showModal" :modal="modal"/>
</template>


<style scoped>

</style>