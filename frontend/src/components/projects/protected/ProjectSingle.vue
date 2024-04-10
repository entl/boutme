

<script>
import ProjectImage from '../ProjectImage.vue';
import ProjectDetails from '../ProjectDetails.vue';
import ProjectTechStack from '../ProjectTechStack.vue';
import UpdateProject from "@/components/projects/protected/UpdateProject.vue";
import DeleteProject from "@/components/projects/protected/DeleteProject.vue";

export default {
  props: ['project'],
  components: {
    DeleteProject,
    UpdateProject,
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
    <div class="flex justify-center gap-2">
      <UpdateProject :project="project"/>
      <DeleteProject :project="project"/>
    </div>
  </div>

</template>

<style scoped>
/* Your component-specific styles */
</style>