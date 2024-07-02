export interface comments {
    comment_id: number,
    created_time: string,
    content: string,
    user_avatar: number,
    username: string,
    root_id: number,
    parent_id: number,
    parent_name: string,
    children: comments[],
    author_avatar: string,

    // 
    replyIsVisible: boolean
    replyText: string
}