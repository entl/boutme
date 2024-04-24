import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import axios from "axios";

const routes = [
	{
		path: '/',
		name: 'Home',
		component: Home,
		meta: {
			title: 'Maksym - Home',
		},
	},
	{
		path: '/about',
		name: 'About',
		// route level code-splitting
		// this generates a separate chunk (about.[hash].js) for this route
		// which is lazy-loaded when the route is visited.
		component: () =>
			import(/* webpackChunkName: "about" */ '../views/About.vue'),
		meta: {
			title: 'Maksym - About',
		},
	},
	{
		path: '/projects',
		name: 'Projects',
		// route level code-splitting
		// this generates a separate chunk (projects.[hash].js) for this route
		// which is lazy-loaded when the route is visited.
		component: () =>
			import(/* webpackChunkName: "projects" */ '../views/Projects.vue'),
		meta: {
			title: 'Maksym - Projects',
		},
	},
	{
		path: '/contact',
		name: 'Contact',
		// route level code-splitting
		// this generates a separate chunk (projects.[hash].js) for this route
		// which is lazy-loaded when the route is visited.
		component: () =>
			import(/* webpackChunkName: "projects" */ '../views/Contact.vue'),
		meta: {
			title: 'Maksym - Contact',
		},
	},
	{
		path: '/login',
		name: 'Login',
		// route level code-splitting
		// this generates a separate chunk (projects.[hash].js) for this route
		// which is lazy-loaded when the route is visited.
		component: () =>
			import(/* webpackChunkName: "projects" */ '../views/Login.vue'),
		meta: {
			title: 'Maksym - Login',
		},
	},
	{
		path: '/admin',
		name: 'Admin',
		component: () =>
			import(/* webpackChunkName: "admin" */ '../protected/Admin.vue'),
		meta: {
			requiresAuth: true,
			title: 'Maksym - Admin',
		},
		children: [
			{
				path: 'dashboard', // Note: This is a relative path
				name: 'Dashboard',
				component: () => import(/* webpackChunkName: "dashboard" */ '../protected/Dashboard.vue'),
				meta: {
					requiresAuth: true,
					title: 'Maksym - Dashboard',
				},
			},
		],
	},
	{
		path: '/:catchAll(.*)',
		name: '404',
		component: () => import(/* webpackChunkName: "404" */ '../views/404.vue'),
	}
];

const router = createRouter({
	history: createWebHistory("/"),
	routes: [...routes],
	scrollBehavior() {
		document.getElementById('app').scrollIntoView();
	},
});
export default router;

router.beforeEach(async (to, from, next) => {
	const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
	const jwtTokenLocal = localStorage.getItem('jwt'); // Or your authentication check logic

	if (requiresAuth && !jwtTokenLocal) {
		next({ name: 'Login' });
	}
	else if (requiresAuth && jwtTokenLocal) {
		const response = await axios.get(`${process.env.VUE_APP_API_URL}/auth`,
			{
				headers: {
					Authorization: `Bearer ${jwtTokenLocal}`,
				},
			})
			.catch(() => {
				next({ name: 'Login' });
			});
		if (response.status === 200) {
			next();
		}
	}
	else {
		next();
	}
});
