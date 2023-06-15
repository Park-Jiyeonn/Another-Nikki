import { createRouter, createWebHistory,  } from 'vue-router'

const systemRoutes = [
	{
		path:'/jelly',
		children: [
			{
				path: '',
				component: () => import('../views/jelly/BlogView.vue'),
			},
		],
	},
	{
		path: '/error',
		children: [
			{
				path: '403',
				component: () => import('../views/error/Error403.vue'),
				meta: { title: 'Error 403' },
			},
			{
				path: '404',
				component: () => import('../views/error/Error404.vue'),
				meta: { title: 'Error 404' },
			},
			{
				path: '500',
				component: () => import('../views/error/Error500.vue'),
				meta: { title: 'Error 500' },
			},
		],
	},
	// {
	// 	path: '/auth',
	// 	meta: { hideSidebar: true },
	// 	children: [
	// 		{
	// 			name: 'Login',
	// 			path: 'login',
	// 			component: () => import('@/views/auth/LoginView.vue'),
	// 		},
	// 		{
	// 			name: 'Logout',
	// 			path: 'logout',
	// 			component: () => import('@/views/auth/LogoutView.vue'),
	// 		},
	// 	],
	// },
]

const fallbackRoutes = [{
	path: '/:any+',
	redirect: '/error/404',
}]

const router = createRouter({
	history: createWebHistory(),
	routes: [
		...systemRoutes,
		...fallbackRoutes,
	],
})

export default router
