import HttpServer from '../http/index';
import { UserType, Commits } from '@/types/User';
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
    password: string;
    confirm_password: string;
}) => HttpServer.request<typeof data,UserType>({
    url: `/api/user/register`,
    method: "POST",
    data
})

const commit_records = (data: {
    user_id : number,
}) => HttpServer.request<typeof data, Commits>({
    url: `api/user/profile/${data.user_id}/commit-record`,
    method: "GET",
})

export {
	login,
    register,
    commit_records,
}
