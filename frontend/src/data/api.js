import axios from "axios";

export const getProjects = async () => {
	try {
		const response = await axios.get(`${process.env.VUE_APP_API_URL}/projects`)
		return response.data;
	} catch (error) {
		console.error("Get Projects Error", error);
	}
}

export const getProject = async (id) => {
	try {
		const response = await axios.get(`${process.env.VUE_APP_API_URL}/${id}`)
		return response.data;
	} catch (error) {
		console.error("Get Project Error", error);
	}
}

export const createProject = async (project, jwt) => {
	try {
		await axios.post(`${process.env.VUE_APP_API_URL}/admin/projects`, project, {
			headers: {
				'Content-Type': 'application/json',
				'Authorization': 'Bearer ' + jwt// Add the token to the headers
			}
		});
	}
	catch (error) {
		console.error("Error creating project", error);
	}
}

export const updateProject = async (project, jwt) => {
	try {
		await axios.put(`${process.env.VUE_APP_API_URL}/admin/projects/${project.id}`, project, {
			headers: {
				'Content-Type': 'application/json',
				'Authorization': 'Bearer ' + jwt // Add the token to the headers
			}
		});
	}
	catch (error) {
		console.error("Error updating project", error);
	}
}

export const deleteProject = async (id, jwt) => {
	try {
		await axios.delete(`${process.env.VUE_APP_API_URL}/admin/projects/${id}`, {
			headers: {
				'Content-Type': 'application/json',
				'Authorization': 'Bearer ' + jwt // Add the token to the headers
			}
		});
	}
	catch (error) {
		console.error("Error deleting project", error);
	}
}

export const login = async (formData) => {
	try {
		const response = await axios.post(`${process.env.VUE_APP_API_URL}/login`, formData, {
			headers: {
				'Content-Type': 'multipart/form-data'
			}
		});
		console.log('After login:', response.data);
		console.log('Response Headers:', response.headers);
		console.log('Request:', response.request);
		return response.data.token
	} catch (error) {
		console.error('Login failed:', error);
	}
}