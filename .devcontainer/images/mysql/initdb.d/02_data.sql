USE go_todo;

/*
name: 'test1'
email: 'test1@example.com'
password: 'test-1234'
*/
INSERT INTO users (name, email, password) VALUES (
    'test1',
    'test1@example.com',
    '$2b$12$9hO29GtfAY3xmv83yr3u7uMroqPStHQuPptNjm3Qu1Rb7UU9UB4Ga'
);

/*
name: '東京駅'
location: 35.68142354732969, 139.76709261114823
user_id: 1
*/
INSERT INTO todos (title, task, user_id) VALUES
('golangの勉強', 'ポインタを理解する',1),
('部屋の掃除', '',1);
