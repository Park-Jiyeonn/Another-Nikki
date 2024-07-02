import { LogsType } from '@/types/Logs';
import HttpServer from '../http/index';

const page_que = (data:{
	page_size: number;
    page_num: number;
}) => HttpServer.request<typeof data,{logs:LogsType[], sum_log: number}>({
    url: `/api/log/${data.page_size}/${data.page_num}`,
    method: "GET",
    data
})

export {
	page_que,
}
