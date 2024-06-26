import { comments } from '@/types/Comment';
import HttpServer from '../http/index';

const post_comment = (data:{
	content:            string;
    article_id:         number;
    user_id:          number;
    username:        string;
    parent_id:          number,
    root_id:            number,
    parent_name:        string,
    user_avatar:      string,
}) => HttpServer.request<typeof data,null>({
    url:        `/api/comment/post`,
    method:     "POST",
    data
})

const reply_comment = (data:{
	content:            string,
    article_id:         number;
    user_id:          number,
    username:        string,
    parent_id:          number,
    root_id:            number,
    parent_name:        string,
    user_avatar:      string,
}) => HttpServer.request<typeof data,null>({
    url: `/api/comment/post`,
    method: "POST",
    data
})

const get_last_seven_comments = (data:{
	num: number;
    article_id:         number;
}) => HttpServer.request<typeof data, {comments: comments[]}>({
    url: `/api/comment/last_seven/${data.article_id}/${data.num}`,
    method: "GET",
})

const get_comments_by_article = (data:{
	num: number;
    article_id:         number;
}) => HttpServer.request<typeof data,{comments: comments[]}>({
    url: `/api/comment/${data.article_id}`,
    method: "GET",
})

const get_random_comment = (
    data:{
        article_id:         number; 
    }
) => HttpServer.request<null, Comment[]>({
    url: `/api/comment/random/${data.article_id}`,
    method: "GET",
})

export {
	post_comment,
    get_last_seven_comments,
    get_random_comment,
    reply_comment,
    get_comments_by_article,
}
