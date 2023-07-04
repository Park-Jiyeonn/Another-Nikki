import { useCookies } from '@vueuse/integrations/useCookies'
import { CookieSetOptions } from 'universal-cookie'

const Cookies = useCookies()

export const setCookies = (data: string, key: any) => {
	const hostname = window.location.hostname
	// hostname 除了常规域名和 IP 地址以外，还有可能是 localhost。这里兼容一下 localhost。
	const domain = hostname.includes('.') ? '.' + hostname : hostname
	Cookies.set(data, key, {
		expires: new Date(new Date().getTime() + (30 * 24 * 60 * 60 * 1000)),	// 30 天
		domain,
		path: '/',
	})
}

export const removeCookies = (key: string , options?:CookieSetOptions) => Cookies.remove(key,options)

export const getCookies = (key: string) => Cookies.get(key)
