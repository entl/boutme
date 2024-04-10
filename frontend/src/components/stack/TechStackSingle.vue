<script>
export default {
	props: ['technology'],
};
</script>

<template>
	<div>
		<img
			:src="technology.img"
			:alt="technology.title"
			class="w-60 h-[110px] py-5 px-10 border border-ternary-light dark:border-ternary-dark shadow-sm rounded-lg mb-8 cursor-pointer dark:bg-secondary-light"
    />
	</div>
<!--  <div class="h-20 w-64 flex justify-center items-center border border-ternary-light dark:border-ternary-dark shadow-sm rounded-lg mb-8 cursor-pointer dark:bg-secondary-light">-->
<!--    <img-->
<!--        :src="technology.img"-->
<!--        :alt="technology.title"-->
<!--        class="max-h-fit max-w-fit object-fill"-->
<!--    />-->
<!--  </div>-->
</template>

<style lang="scss" scoped></style>
