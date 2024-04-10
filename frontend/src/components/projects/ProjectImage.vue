<template>
  <div>
    <img :src="imageUrl" :alt="name" class="rounded-t-xl border-none h-60 w-full object-cover" />
  </div>
</template>

<script>
export default {
  props: ['imageUrl', 'name']
};
</script>
