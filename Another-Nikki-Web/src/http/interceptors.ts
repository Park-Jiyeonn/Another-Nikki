import { getCookies } from "@/hooks/useCookies";
import axios, { AxiosError, AxiosInstance } from "axios";
import { ElMessage } from "element-plus";
import router from '@/router'

export class Interceptors {
  instance: AxiosInstance;
  constructor() {
    this.instance = axios.create({
      baseURL: import.meta.env.VITE_API_URL,
    });
  }
  // 初始化拦截器
  init() {
    // 请求接口拦截器
    this.instance.interceptors.request.use(
      (config) => {
				// const token = useTokenInfo()
				const token = getCookies('token')
				if(token) config.headers['Authorization'] = `Bearer ${token}`
        return config;
      },
      (err) => {
        console.error(err,'error');
      }
    );

    // 响应拦截器
    this.instance.interceptors.response.use(
      (response) => {

        return response;
      },
      (error:AxiosError) => {
        ElMessage.error("Something went wrong, please try again.")
				if (error.response?.status === 401){
					// TODO 把当前路由写进 redirect 参数
					router.push({path:'/auth/login'})
				}
        if (error.message === "Request failed with status code 500") {
          // console.error("系统错误，请检查API是否正常！");
          return;
        }
        // return Promise.resolve(error);
      }
    );
  }
  // 返回一下
  getInterceptors() {
    this.init()
    return this.instance;
  }
}
