<!--<script>-->
<!--export default {-->
<!--	props: ['project'],-->
<!--  mounted() {-->
<!--    console.log("ProjectSingle.vue mounted()")-->
<!--    console.log(this.project);-->
<!--  },-->
<!--  methods: {-->
<!--    truncateText(text, length) {-->
<!--      if (text.length <= length) return text;-->
<!--      return text.substring(0, length) + '...';-->
<!--    }-->
<!--  }-->
<!--};-->
<!--</script>-->

<!--<template>-->
<!--	<a :href="project.github_url"-->
<!--     target="_blank"-->
<!--		class="rounded-xl shadow-lg hover:shadow-xl cursor-pointer mb-10 sm:mb-0 bg-secondary-light dark:bg-ternary-dark"-->
<!--		aria-label="Single Project"-->
<!--	>-->
<!--		<div>-->
<!--			<img-->
<!--				:src="project.image_url"-->
<!--				:alt="project.name"-->
<!--				class="rounded-t-xl border-none h-60 w-full object-cover"-->
<!--			/>-->
<!--		</div>-->
<!--		<div class="text-center px-4 py-6">-->
<!--			<p-->
<!--				class="font-general-semibold text-xl text-ternary-dark dark:text-ternary-light font-semibold mb-2"-->
<!--			>-->
<!--				{{ project.name }}-->
<!--			</p>-->
<!--      <span-->
<!--          class="font-general-medium text-lg text-ternary-dark dark:text-ternary-light block overflow-hidden"-->
<!--          :title="project.description"-->
<!--      >-->
<!--        {{ truncateText(project.description, 100) }}-->
<!--      </span>-->
<!--		</div>-->
<!--    <div class="px-4 py-1 flex justify-center flex-wrap gap-2">-->
<!--      <span-->
<!--          v-for="tech in project.tech_stack"-->
<!--          :key="tech"-->
<!--          class="bg-primary-light dark:bg-primary-dark text-primary-dark dark:text-primary-light font-medium py-1 px-3 rounded-full text-sm"-->
<!--      >-->
<!--        {{ tech }}-->
<!--      </span>-->
<!--    </div>-->

<!--	</a>-->
<!--</template>-->

<!--<style lang="scss" scoped></style>-->

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