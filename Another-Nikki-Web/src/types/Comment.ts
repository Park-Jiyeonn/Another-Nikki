export interface Comment {
    ID:             number,
    CreatedAt:      string,
    content:        string,
    author_id:      number,
    author_name:    string,
    root_id:        number,
    parent_id:      number,
    parent_name:    string,
    children:       Comment[],
    author_avatar:  string,

    // 
    replyIsVisible: boolean
    replyText:      string
}