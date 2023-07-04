import HttpServer from '../http/index';
import { UserType } from '@/types/User';
const login = (data:{
	username: string;
    password: string;
}) => HttpServer.request<typeof data,UserType>({
    url: `/api/user/login`,
    method: "POST",
    data
})

const register = (data:{
	username: string;
    password1: string;
    password2: string;
}) => HttpServer.request<typeof data,UserType>({
    url: `/api/user/register`,
    method: "POST",
    data
})

export {
	login,
    register,
}
