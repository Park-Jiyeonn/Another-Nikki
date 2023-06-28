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

const update_article = (data:{
    title: string;
	ID: number;
    content: string;
    description:string,
}) => HttpServer.request<typeof data,Article>({
    url: `/api/article/update`,
    method: "POST",
    data
})

export {
	page_que,
    get_article,
    update_article,
}
