import { Blog } from '@/types/Blog';
import HttpServer from '../http/index';

const PostBlog = (data:{
	content: string;
}) => HttpServer.request<typeof data,null>({
    url: `/api/blog/create_blog`,
    method: "POST",
    data
})

const get_last_seven_blogs = (data:{
	num: number;
}) => HttpServer.request<typeof data,Blog[]>({
    url: `/api/blog/get_last_seven_blogs`,
    method: "GET",
    data
})

const get_random_blogs = () => HttpServer.request<null, Blog[]>({
    url: `/api/blog/get_random_blog`,
    method: "GET",
})

export {
	PostBlog,
    get_last_seven_blogs,
    get_random_blogs,
}
