<script>
import UpdateProjectModal from "@/components/projects/protected/UpdateProjectModal.vue";
import feather from "feather-icons";
export default {
  components: {
    UpdateProjectModal,
  },
  props: {
    project: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      isOpen: false,
      modal: false,
      projects: null,
    };
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
  <div class="my-3">
    <button
        @click="showModal"
        class="px-4 py-2.5 mx-2 text-ternary-dark dark:text-gray-300 tracking-wider hover:-translate-y-1 focus:ring-1 focus:ring-indigo-900 rounded-lg duration-500"
        aria-label="Edit"
    >
      <i data-feather="edit-2" class="inline-block"></i>
    </button>
  </div>
  <UpdateProjectModal :showModal="showModal" :modal="modal" :project_in="project"/>

</template>


<style scoped>
</style>
