import { Article } from '@/types/Article';
import HttpServer from '../http/index';

const page_que = (data:{
	page_num: number;
    page_size: number;
}) => HttpServer.request<typeof data,Article[]>({
    url: `/api/article/${data.page_num}/${data.page_size}`,
    method: "GET",
})

const get_article = (data:{
	article_id: number;
}) => HttpServer.request<typeof data,Article>({
    url: `/api/article/${data.article_id}`,
    method: "GET",
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

const post_article = (data:{
    article_id: number,
    article_title: string;
    article_content: string;
    article_description:string, 
}) => HttpServer.request<typeof data,Article>({
    url: `/api/article/post`,
    method: "POST",
    data
}) 

export {
	page_que,
    get_article,
    update_article,
    post_article,
}
