import { AxiosInstance, AxiosRequestConfig, AxiosResponse,AxiosError ,AxiosPromise} from "axios";
import { Interceptors } from "./interceptors";
import { successType } from '../types/Api'

export class HttpServer {
  axios: AxiosInstance;
  constructor() {
    this.axios = new Interceptors().getInterceptors();
  }
  request<T = any,K = any>(config: AxiosRequestConfig<T>):AxiosPromise<successType<K>> {
    if(config.method === 'GET') config.params = config.data ?? ''
    return new Promise((resolve, reject) => {
      this.axios(config).then((res:AxiosResponse<successType<K>,T>) => {
        resolve(res);
      }).catch((err:AxiosError) => {
        reject(err)
      });
    });
  }
}

const http = new HttpServer()

export default http