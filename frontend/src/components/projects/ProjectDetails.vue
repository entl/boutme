<template>
  <div class="text-center px-4 py-6">
    <p class="font-general-semibold text-xl text-ternary-dark dark:text-ternary-light font-semibold mb-2">{{ name }}</p>
    <span class="font-general-medium text-lg text-ternary-dark dark:text-ternary-light block overflow-hidden" :title="description">
      {{ truncateText(description, 100) }}
    </span>
  </div>
</template>

<script>
export default {
  props: ['name', 'description'],
  methods: {
    truncateText(text, length) {
      if (text.length <= length) return text;
      return text.substring(0, length) + '...';
    }
  }
};
</script>
