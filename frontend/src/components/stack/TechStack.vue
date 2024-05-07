<script>
import TechStackSingle from "@/components/stack/TechStackSingle.vue";

export default {
	components: {
		TechStackSingle,
	},
	data: () => {
		return {
			techStackHeading: 'Some of the technologies I have worked with:',
			technologies: [
				{
					id: 1,
					title: 'Python',
					img: "https://img.shields.io/badge/python-white?style=for-the-badge&logo=python",
				},
				{
					id: 2,
					title: 'Go',
					img: "https://img.shields.io/badge/go-white.svg?style=for-the-badge&logo=go&logoColor=blue",
				},
				{
					id: 3,
					title: 'Java',
					img: "https://img.shields.io/badge/java-white?style=for-the-badge&logo=openjdk&logoColor=red",
				},
				{
					id: 4,
					title: 'SQL',
					img: "https://img.shields.io/badge/sql-white?style=for-the-badge&logo=mysql&logoColor=blue",
				},
				{
					id: 5,
					title: 'TypeScript',
					img: "https://img.shields.io/badge/TypeScript-white?style=for-the-badge&logo=typescript&logoColor=black",
				},
				{
					id: 6,
					title: 'FastAPI',
					img: "https://img.shields.io/badge/FastAPI-white?style=for-the-badge&logo=fastapi",
				},
        {
          id: 12,
          title: 'Spring MVC',
          img: "https://img.shields.io/badge/Spring-white?style=for-the-badge&logo=spring",
        },
				{
					id: 7,
					title: 'PostrgeSQL',
					img: "https://img.shields.io/badge/postgres-white?style=for-the-badge&logo=postgresql",
				},
				{
					id: 8,
					title: 'Redis',
					img: "https://img.shields.io/badge/redis-white.svg?&style=for-the-badge&logo=redis",
				},
        {
          id: 9,
          title: 'MongoDB',
          img: "https://img.shields.io/badge/MongoDB-white?style=for-the-badge&logo=mongodb",
        },
        {
          id: 10,
          title: 'Git',
          img: "https://img.shields.io/badge/GIT-white?style=for-the-badge&logo=git",
        },
        {
          id: 11,
          title: 'Docker',
          img: "https://img.shields.io/badge/docker-white.svg?style=for-the-badge&logo=docker",
        },
			],
		};
	},
};
</script>

<template>
	<div class="mt-10 sm:mt-20">
		<p
			class="font-general-medium text-2xl sm:text-3xl text-primary-dark dark:text-primary-light"
		>
			{{ techStackHeading }}
		</p>
    <div class="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-4 mt-10 sm:mt-14 gap-2">
    <TechStackSingle
				v-for="technology in technologies"
				:key="technology.id"
				:technology="technology"
			/>
		</div>
	</div>
</template>