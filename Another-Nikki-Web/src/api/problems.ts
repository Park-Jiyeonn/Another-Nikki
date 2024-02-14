import { GetProblemByPageResp, GetProblemByIdResp, Problem } from '@/types/Problem';
import HttpServer from '../http/index';

const page_que = (data:{
	page_size: number,
    page_num: number,
}) => HttpServer.request<typeof data,GetProblemByPageResp[]>({
    url: `/api/problem/${data.page_num}/${data.page_size}`,
    method: "GET",
})

const get_problem = (data:{
	problem_id: number;
}) => HttpServer.request<typeof data,GetProblemByIdResp>({
    url: `/api/problem/${data.problem_id}`,
    method: "GET",
})

const update_problem = (data:{
    problem_title: string;
	problem_id: number;
    problem_content: string;
}) => HttpServer.request<typeof data,Problem>({
    url: `/api/problem/post`,
    method: "POST",
    data
})

const post_problem = (data:{
    problem_title: string;
    problem_content: string;
}) => HttpServer.request<typeof data,Problem>({
    url: `/api/problem/post`,
    method: "POST",
    data
}) 

export {
	page_que,
    get_problem,
    update_problem,
    post_problem,
}
