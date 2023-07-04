import { Comment } from '@/types/Comment';
import HttpServer from '../http/index';

const post_comment = (data:{
	content:            string;
    author_id:          number;
    author_name:        string;
    parent_id:          number,
    root_id:            number,
    parent_name:        string,
    author_avatar:      string,
}) => HttpServer.request<typeof data,null>({
    url:        `/api/blog/create_blog`,
    method:     "POST",
    data
})

const reply_comment = (data:{
	content:            string,
    author_id:          number,
    author_name:        string,
    parent_id:          number,
    root_id:            number,
    parent_name:        string,
    author_avatar:      string,
}) => HttpServer.request<typeof data,null>({
    url: `/api/blog/create_blog`,
    method: "POST",
    data
})

const get_last_seven_comments = (data:{
	num: number;
}) => HttpServer.request<typeof data,Comment[]>({
    url: `/api/blog/get_last_seven_blogs`,
    method: "GET",
    data
})

const get_random_comment = () => HttpServer.request<null, Comment[]>({
    url: `/api/blog/get_random_blog`,
    method: "GET",
})

export {
	post_comment,
    get_last_seven_comments,
    get_random_comment,
    reply_comment,
}
