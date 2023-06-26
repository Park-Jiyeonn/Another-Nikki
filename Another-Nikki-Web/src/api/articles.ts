import { Article } from '@/types/Article';
import HttpServer from '../http/index';

const page_que = (data:{
	page: number;
}) => HttpServer.request<typeof data,Article[]>({
    url: `/api/article/articles`,
    method: "GET",
    data
})

const get_article = (data:{
	ID: number;
}) => HttpServer.request<typeof data,Article>({
    url: `/api/article`,
    method: "GET",
    data
})

export {
	page_que,
    get_article,
}
