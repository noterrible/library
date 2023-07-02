create table library_v2.book_infos
(
    id                  int(11) unsigned           not null comment '书的id'
        primary key,
    book_name           varchar(200)               null comment '书名',
    author              varchar(50)                null comment '作者',
    publishing_house    varchar(50)                null comment '出版社',
    translator          varchar(50)                null comment '译者',
    publish_date        date                       null comment '出版时间',
    pages               int(10)        default 100 null comment '页数',
    ISBN                varchar(20)                null comment 'ISBN号码',
    price               double         default 1   null comment '价格',
    brief_introduction  varchar(15000) default ''  null comment '内容简介',
    author_introduction varchar(5000)  default ''  null comment '作者简介',
    img_url             varchar(200)               null comment '封面地址',
    del_flg             int(1)         default 0   null comment '删除标识',
    count               int            default 100 null,
    category_id         int            default 1   null,
    constraint book_infos_ISBN_book_name_uindex
        unique (ISBN, book_name)
);

create table library_v2.categories
(
    id   bigint auto_increment
        primary key,
    name varchar(100) null
);

create table library_v2.librarians
(
    id        bigint auto_increment
        primary key,
    user_name varchar(100) null,
    password  varchar(100) null,
    name      varchar(100) null,
    sex       varchar(100) null,
    phone     varchar(100) null
);

create table library_v2.messages
(
    id          bigint auto_increment,
    user_id     bigint       null,
    message     varchar(100) null,
    status      bigint       not null,
    create_time datetime(3)  null,
    primary key (id, status)
);

create table library_v2.records
(
    id          bigint auto_increment
        primary key,
    user_id     bigint      null,
    book_id     bigint      null,
    status      bigint      null,
    start_time  datetime(3) null,
    over_time   datetime(3) null,
    return_time datetime(3) null
);

create table library_v2.users
(
    id        bigint auto_increment
        primary key,
    user_name varchar(100) null,
    password  varchar(100) null,
    name      varchar(100) null,
    sex       varchar(100) null,
    phone     varchar(100) null,
    status    bigint       null,
    constraint users_user_name_uindex
        unique (user_name)
);


