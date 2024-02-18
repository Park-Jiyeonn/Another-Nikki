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
    page_num : number,
    page_size: number,
    user_id: number,
}) => HttpServer.request<typeof data, Commits>({
    url: `api/user/profile/${data.user_id}/commit-record/${data.page_num}/${data.page_size}`,
    method: "GET",
})

const count = (data: {
    user_id: number,
}) => HttpServer.request<null,{sum:number}>({
    url: `/api/user/profile/${data.user_id}/commit-record/sum`,
    method: "GET",
})

const update = (data:{
	username: string,
    avatar: string,
    description: string,
}) => HttpServer.request<typeof data, UserType>({
    url: `/api/user/update`,
    method: "POST",
    data
})

const get_commit_by_id = (data: {
    judge_id : number,
}) => HttpServer.request<typeof data, Commits>({
    url: `/api/code/view-submission/${data.judge_id}`,
    method: "GET",
})

const get_user_by_id = (data: {
    user_id : number,
}) => HttpServer.request<typeof data, UserType>({
    url: `/api/user/profile/${data.user_id}`,
    method: "GET",
})

export {
	login,
    register,
    commit_records,
    count,
    update,
    get_commit_by_id,
    get_user_by_id,
}
