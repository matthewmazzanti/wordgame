create table game (
    id int not null primary key auto_increment,
    letters varchar(32) not null
);

create table word (
    id int not null primary key auto_increment,
    word varchar(32) not null,
    bitmap int
);
