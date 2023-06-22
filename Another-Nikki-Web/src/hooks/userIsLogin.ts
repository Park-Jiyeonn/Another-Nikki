import { ref } from 'vue'
import { getCookies } from './useCookies'

const $isLoggedIn = ref(false)

// 判断是否已登录（通过本地存储中的 token 来判断，但 token 可能会过期）
export function useIsLoggedIn() {
	// TODO: useTokenInfo() 需要改造一下，改为可读可写。这里应该避免手工读取本地存储
	let authToken = getCookies('token');
	$isLoggedIn.value = !!authToken
	return $isLoggedIn
}
