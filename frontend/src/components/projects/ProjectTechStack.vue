<template>
  <div class="px-4 py-1 flex justify-center flex-wrap gap-2 my-2">
    <span v-for="tech in techStack" :key="tech" class="bg-primary-light dark:bg-primary-dark text-primary-dark dark:text-primary-light font-medium py-1 px-3 rounded-full text-sm">
      {{ tech }}
    </span>
  </div>
</template>

<script>
export default {
  props: ['techStack']
};
</script>
