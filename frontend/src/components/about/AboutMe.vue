<script>
export default {
	setup() {
		return {
			bios: [
				{
					id: 1,
					bio:
            "I'm a highly motivated Software Engineering student at Anglia Ruskin University (88/100 GPA) with a proven track record of translating real-world experience into measurable results."
				},
				{
					id: 2,
					bio:
						"I possess a robust foundation in Python, GoLang, Java, JavaScript/TypeScript, various databases (PostgreSQL, MongoDB, Redis), and frameworks such as FastAPI, Echo, Spring MVC, React, and Vue. My passion for continuous learning is evident in certifications and my consistent academic excellence.",
				},
                  {
          id: 3,
          bio:
            "I thrive in collaborative, fast-paced settings and I'm seeking opportunities to apply my skills, adaptability, and results-driven mindset to contribute to impactful software projects. Let's connect!"
        },
			],
		};
	},
};
</script>

<template>
	<div class="block sm:flex sm:gap-10 mt-10 sm:mt-20">
		<!-- About profile image -->
		<div class="w-full sm:w-1/4 mb-7 sm:mb-0">
			<img
				src="@/assets/images/profile.jpg"
				class="rounded-xl w-96"
				alt=""
			/>
		</div>

		<!-- About details -->
		<div class="w-full sm:w-3/4 text-left">
			<p
				v-for="bio in bios"
				:key="bio.id"
				class="font-general-regular mb-4 text-ternary-dark dark:text-ternary-light text-lg"
			>
				{{ bio.bio }}
			</p>
		</div>
	</div>
</template>
