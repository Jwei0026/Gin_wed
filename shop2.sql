--创建相关表结构和插入数据

--根据看到的效果先拟（后续补充完整）

--品牌
create table brand(
    id int primary key auto_increment,
    name varchar(20) not null unique,
    story text not null
)charset utf8;

--插入数据
insert into brand values(null,"瑞士爱宝时(EPOS)","品牌故事1"),
insert into brand values(null,"百达翡丽","品牌故事2"),
insert into brand values(null,"浪琴longines","品牌故事3");

--商品
create table goods(
    id int primary key auto_increment,
    name varchar(100) not null,
    normal_price int unsigned nto null comment '无符号整数：没有负数',
    current_price int unsigned nto null default 0,

    biaoke varchar(20) not null,
    baodi varchar(20) not null,
    color varchar(20) not null,
    func varchar(20) not null,
    biaojing varchar(20) not null,
    biaokou varchar(20) not null,
    kuanshi enum("男","女") default 1 comment"enum是美剧类型，从1开始",
    `year` year int not null,
    heart varchar(20) not null, 
    houdu decimal(5,2) not null comment"5代表有效长度，2代表小数有效位",
    fangshui tinyint not null default 100,

    brand_id int not null,
    serias_id int not null,
);