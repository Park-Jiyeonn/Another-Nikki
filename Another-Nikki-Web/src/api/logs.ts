import { LogsType } from '@/types/Logs';
import HttpServer from '../http/index';

const page_que = (data:{
	page: number;
}) => HttpServer.request<typeof data,{logs:LogsType[]}>({
    url: `/api/log/get_page_que`,
    method: "GET",
    data
})

const count = () => HttpServer.request<null,{sum:number}>({
    url: `/api/log/count`,
    method: "GET",
})

const get_visit_time = (data:{
	user_id: string;
}) => HttpServer.request<typeof data,{sum_visit_time:number,today_visit_time:number}>({
    url: `/api/log/visit_time`,
    method: "GET",
    data
})

export {
	page_que,
    count,
    get_visit_time,
}
