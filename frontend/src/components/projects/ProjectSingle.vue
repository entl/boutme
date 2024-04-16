<script>
import ProjectImage from './ProjectImage.vue';
import ProjectDetails from './ProjectDetails.vue';
import ProjectTechStack from './ProjectTechStack.vue';

export default {
  props: {
    project: {
      type: Object,
      required: true,
    },
  },
  components: {
    ProjectImage,
    ProjectDetails,
    ProjectTechStack,
  }
};
</script>

<template>
  <div
      class="flex flex-col justify-between rounded-xl shadow-lg hover:shadow-xl cursor-pointer mb-10 sm:mb-0 bg-secondary-light dark:bg-ternary-dark"
      aria-label="Single Project"
  >
    <div class="flex-grow">
      <a target="_blank"
         :href="project.github_url"
         class="flex flex-col"
      >
        <ProjectImage :imageUrl="project.image_url" :name="project.name"/>
        <ProjectDetails :name="project.name" :description="project.description"/>
      </a>
    </div>
    <ProjectTechStack :techStack="project.tech_stack" class="p-4"/>
  </div>
</template>


<style scoped>
/* Your component-specific styles */
</style>