USE go_todo;

/*
name: 'test1'
email: 'test1@example.com'
password: 'test-1234'
*/
INSERT INTO users (id, name, email, password) VALUES (
    '2a7143fc-f229-49e7-8af1-d14ea2295874',
    'test1',
    'test1@example.com',
    '$2b$12$9hO29GtfAY3xmv83yr3u7uMroqPStHQuPptNjm3Qu1Rb7UU9UB4Ga'
);

/*
name: '東京駅'
location: 35.68142354732969, 139.76709261114823
user_id: 1
*/
INSERT INTO tasks (id, title, content, user_id) VALUES
('18258ee0-c97d-410b-aef7-9eaea375549c', 'golangの勉強', 'ポインタを理解する','2a7143fc-f229-49e7-8af1-d14ea2295874'),
('0c6de6dc-18ec-4e60-a2d8-2d520935648c', '部屋の掃除', '','2a7143fc-f229-49e7-8af1-d14ea2295874');
