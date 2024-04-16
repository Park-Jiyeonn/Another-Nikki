import { createRouter, createWebHistory,  } from 'vue-router'
import { useIsLoggedIn } from '@/hooks/userIsLogin'

const systemRoutes = [
	{
		path:'/',
		component: () => import('@/views/HomeView.vue'),
	},
	{
		name: "auth",
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
		name:'articles',
		path:'/articles',
		component: () => import('@/views/article/ArticleListView.vue'),
	},
	{
		path: '/article/:id',
		component: () => import('@/views/article/ArticleView.vue'),
	},
	{
		path: '/article/update/:id',
		component: () => import('@/views/article/ArticleUpdate.vue'),
		meta: {
			requiresAuth: true,
		},
	},
	{
		path:'/article/post',
		component: () => import('@/views/article/ArticlePost.vue'),
		meta: {
			requiresAuth: true,
		},
	},	

	{
		name:'problems',
		path:'/problems',
		component: () => import('@/views/problem/ProblemList.vue'),
	},
	{
		path: '/problem/:id',
		component: () => import('@/views/problem/ProblemDetail.vue'),
	},
	{
		path:'/problem/post',
		component: () => import('@/views/problem/ProblemPost.vue'),
		meta: {
			requiresAuth: true,
		},
	},{
		path: '/problem/update/:id',
		component: () => import('@/views/problem/ProblemUpdate.vue'),
		meta: {
			requiresAuth: true,
		},
	},	{
		path: '/error',
		children: [
			{
				path: '403',
				component: () => import('@/views/error/Error403.vue'),
				meta: { title: 'Error 403' },
			},
			{
				path: '404',
				component: () => import('@/views/error/Error404.vue'),
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
		path:'/message-board',
		component: () => import('@/views/jelly/OtherView.vue'),
		meta: {
			requiresAuth: true,
		},
	},
]

const RunCodeRoutes = [
	{
		path:'/profile/:user_id',
		component: () => import('@/views/runCode/CommitRecord.vue'),
	},
	{
		path: '/user/code/:judge_id',
		component: () => import('@/views/runCode/CodeDetail.vue'),
		meta: {
			requiresAuth: true,
		},
	},
	{
		path: '/wrong_answer/:user_id',
		component: () => import('@/views/runCode/WrongAns.vue'),
		meta: {
			requiresAuth: true,
		},
	}
]

const privateRoutes = [
	{
		path:'/jelly',
		component: () => import('@/views/jelly/JellyView.vue'),
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
		let query
		// 通过 redirect 参数保存当前所在的位置，以便登录后返回
		// 如果当前是首页，就不用存了，因为登录后默认回首页
		if (to.fullPath !== '/') query = { redirect: to.fullPath }
		return {
			path: '/auth/login',
			query,
		}
	}
})

export default router
