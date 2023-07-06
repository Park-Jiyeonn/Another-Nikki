import { Problem } from '@/types/Problem';
import HttpServer from '../http/index';

const page_que = (data:{
	page: number;
}) => HttpServer.request<typeof data,Problem[]>({
    url: `/api/problem/problems`,
    method: "GET",
    data
})

const get_problem = (data:{
	ID: number;
}) => HttpServer.request<typeof data,Problem>({
    url: `/api/problem`,
    method: "GET",
    data
})

const update_problem = (data:{
    name: string;
	ID: number;
    content: string;
}) => HttpServer.request<typeof data,Problem>({
    url: `/api/problem/update`,
    method: "POST",
    data
})

const post_problem = (data:{
    name: string;
    content: string;
}) => HttpServer.request<typeof data,Problem>({
    url: `/api/problem`,
    method: "POST",
    data
}) 

export {
	page_que,
    get_problem,
    update_problem,
    post_problem,
}
