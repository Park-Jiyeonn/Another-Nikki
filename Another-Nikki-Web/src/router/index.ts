import { createRouter, createWebHistory,  } from 'vue-router'
import { useIsLoggedIn } from '@/hooks/userIsLogin'

const systemRoutes = [
	{
		path:'/',
		component: () => import('../views/HomeView.vue'),
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
	{
		path: '/auth',
		meta: { hideSidebar: true },
		children: [
			{
				name: 'Login',
				path: 'login',
				component: () => import('@/views/auth/LoginView.vue'),
			},
			{
				name: 'register',
				path: 'register',
				component: () => import('@/views/auth/Register.vue'),
			},
		],
	},
	{
		path:'/article',
		component: () => import('@/views/article/ArticleListView.vue'),
	}
]

const RunCodeRoutes = [
	{
		path:'/runcode',
		component: () => import('../views/runCode/RunCodeView.vue'),
	},
]

const privateRoutes = [
	{
		path:'/jelly',
		component: () => import('../views/jelly/BlogView.vue'),
		meta: {
			requiresAuth: true,
		},
	},
	{
		path:"/logs",
		component: () => import('@/views/logs_/LogView.vue'),
		meta: {
			requiresAuth: true,
		},
	}
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
		...RunCodeRoutes,
		...privateRoutes,
	],
})

router.beforeEach(to => {
	const isLoggedIn = useIsLoggedIn()
	if (to.meta.requiresAuth && !isLoggedIn.value) {
		// let query
		// 通过 redirect 参数保存当前所在的位置，以便登录后返回
		// 如果当前是首页，就不用存了，因为登录后默认回首页
		// if (to.fullPath !== '/') query = { redirect: to.fullPath }
		return {
			path: '/error/403',
			// query,
		}
	}
})

export default router
