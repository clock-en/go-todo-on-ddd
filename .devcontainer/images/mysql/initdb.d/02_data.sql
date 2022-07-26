USE go_todo;

/*
id: '2a7143fc-f229-49e7-8af1-d14ea2295874'
name: 'test1'
email: 'test1@example.com'
password: 'Test-1234'
*/
INSERT INTO users (id, name, email, password) VALUES (
    '2a7143fc-f229-49e7-8af1-d14ea2295874',
    'test1',
    'test1@example.com',
    '57f226e08ab8269072db12617e4aa0ac7091d909c11158a634ed9084f97a7681'
);

INSERT INTO tasks (id, title, content, user_id) VALUES
('18258ee0-c97d-410b-aef7-9eaea375549c', 'golangの勉強', 'ポインタを理解する','2a7143fc-f229-49e7-8af1-d14ea2295874'),
('0c6de6dc-18ec-4e60-a2d8-2d520935648c', '部屋の掃除', '','2a7143fc-f229-49e7-8af1-d14ea2295874');
