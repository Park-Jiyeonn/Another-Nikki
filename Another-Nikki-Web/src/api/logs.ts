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

export {
	page_que,
    count,
}
