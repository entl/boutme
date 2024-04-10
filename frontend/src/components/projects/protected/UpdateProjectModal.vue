<script>
import feather from 'feather-icons';
import Button from '../../reusable/Button.vue';
import FormInput from '../../reusable/FormInput.vue';
import FormTextarea from '../../reusable/FormTextarea.vue';
import axios from "axios";
export default {
	props: ['showModal', 'modal', 'project_in'],
	components: { Button, FormInput, FormTextarea },
	data() {
		return {
      project : {
        name: '',
        description: '',
        image_url: '',
        github_url: '',
        tech_stack: '',
      }
    };
	},
	mounted() {
		feather.replace();
    this.initializeForm();
	},
	updated() {
		feather.replace();
	},
  watch: {
    project_in: {
      deep: true,
      immediate: true,
      handler(newVal) {
        this.initializeForm(newVal);
      }
    }
  },
	methods: {
    initializeForm(projectData = this.project_in) {
      console.log('Project data:', projectData);
      const techStack = projectData.tech_stack.join(',')
      if (projectData) {
        this.project = {
          ...this.project,
          ...projectData,
          tech_stack: techStack
        };
      }
      console.log('Project:', this.project);
    },
    async submitForm() {
      this.$refs.form.requestSubmit();
    },
    async submit() {
      try {
        this.project.tech_stack = this.project.tech_stack.split(',');
        console.log('Form submitted:', this.project);
        await axios.put(`http://localhost:8080/admin/projects/${this.project.id}`, this.project, {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwt') // Add the token to the headers
          }
        });
      } catch (error) {
        console.error('Adding project failed:', error);
      }
    }
  },
};
</script>

<template>
	<transition name="fade">
		<div v-show="modal" class="font-general-regular fixed inset-0 z-30">
			<!-- Modal body background as backdrop -->
			<div
				v-show="modal"
				@click="showModal()"
				class="bg-filter bg-black bg-opacity-50 fixed inset-0 w-full h-full z-20"
			></div>
			<!-- Modal content -->
			<main
				class="flex flex-col items-center justify-center h-full w-full"
			>
				<transition name="fade-up-down">
					<div
						v-show="modal"
						class="modal-wrapper flex items-center z-30"
					>
						<div
							class="modal max-w-md mx-5 md:max-w-xl bg-secondary-light dark:bg-primary-dark max-h-screen shadow-lg flex-row rounded-lg relative"
						>
							<div
								class="modal-header flex justify-between gap-10 p-5 border-b border-ternary-light dark:border-ternary-dark"
							>
								<h5
									class="text-primary-dark dark:text-primary-light text-xl"
								>
									Woooow!! 0_0 You've made changes!
								</h5>
								<button
									class="px-4 text-primary-dark dark:text-primary-light"
									@click="showModal()"
								>
									<i data-feather="x"></i>
								</button>
							</div>
                <div class="modal-body p-5 w-full h-full">
                  <form ref="form" class="max-w-xl m-4 text-left"  @submit="submit">
                    <FormInput
                      label="Project Name"
                      inputIdentifier="name"
                      v-model:value="project.name"
                      class="mb-2"
                    />
                    <FormTextarea
                      label="Description"
                      inputIdentifier="description"
                      v-model:value="project.description"
                      class="mb-2"
                    />
                    <FormInput
                      label="Image URL"
                      inputIdentifier="image_url"
                      v-model:value="project.image_url"
                      class="mb-2"
                    />
                    <FormInput
                      label="Github URL"
                      inputIdentifier="github_url"
                      v-model:value="project.github_url"
                      class="mb-2"
                    />
                    <FormInput
                      label="Tech Stack"
                      inputIdentifier="tech_stack"
                      v-model:value="project.tech_stack"
                      class="mb-2"
                    />
                  </form>
              </div>
              <div class="modal-footer mt-2 sm:mt-0 py-5 px-8 border-t">
                <div class="flex justify-end">
                    <Button
                        title="Update Project"
                        class="px-4 sm:px-6 py-2 sm:py-2.5 text-white bg-indigo-500 hover:bg-indigo-600 rounded-md focus:ring-1 focus:ring-indigo-900 duration-500"
                        type="submit"
                        @click="submitForm"
                        aria-label="Add Project" />
                  <Button
                      class="px-4 sm:px-6 py-2 bg-gray-600 text-primary-light hover:bg-ternary-dark dark:bg-gray-200 dark:text-secondary-dark dark:hover:bg-primary-light rounded-md focus:ring-1 focus:ring-indigo-900 duration-500 mx-2"
                      title="Close"
                      @click="showModal()"
                      aria-label="Close Modal" />
                </div>
              </div>
						</div>
					</div>
				</transition>
			</main>
		</div>
	</transition>
</template>

<style scoped>
.modal-body {
	max-height: 570px;
  overflow-y: scroll;
}
.bg-gray-800-opacity {
	background-color: #2d374850;
}
@media screen and (max-width: 768px) {
	.modal-body {
		max-height: 400px;
    overflow-y: scroll;
	}
}
.fade-up-down-enter-active {
	transition: all 0.3s ease;
}
.fade-up-down-leave-active {
	transition: all 0.3s ease;
}
.fade-up-down-enter {
	transform: translateY(10%);
	opacity: 0;
}
.fade-up-down-leave-to {
	transform: translateY(10%);
	opacity: 0;
}

.fade-enter-active {
	-webkit-transition: opacity 2s;
	transition: opacity 0.3s;
}
.fade-leave-active {
	transition: opacity 0.3s;
}

.fade-enter,
.fade-leave-to {
	opacity: 0;
}
</style>
