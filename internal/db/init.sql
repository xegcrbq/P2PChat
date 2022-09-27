CREATE TABLE Users(
    UserId SERIAL PRIMARY KEY,
    UserName varchar(50) NOT NULL,
    Password varchar(50) NOT NULL,
    IsTrader bit DEFAULT 0::bit
);

create table Orders(
    OrderId SERIAL PRIMARY KEY,
    SellerId integer references Users(UserId) NOT NULL,
    SellerTicker varchar(10) NOT NULL,
    SellerAmount numeric(17,8) NOT NULL,
    SellerVoteComplete bit DEFAULT 0::bit,
    BuyerId integer references Users(UserId) NOT NULL,
    BuyerTicker varchar(10) NOT NULL,
    BuyerAmount numeric(17,8) NOT NULL,
    BuyerVoteComplete bit DEFAULT 0::bit,
    IsCompleted bit DEFAULT 0::bit
);

create table Attachments(
    AttachmentId SERIAL PRIMARY KEY,
    Path varchar(100)
);

create table Messages(
    MessageId SERIAL PRIMARY KEY,
    SenderId integer references Users(UserId) NOT NULL,
    ReaderId integer references Users(UserId) NOT NULL,
    OrderId integer references Orders(OrderId) NOT NULL,
    MessageText text,
    AttachmentId integer NULL,
    Foreign Key(AttachmentId) references Attachments(AttachmentId),
    SendTime timestamp WITH TIME ZONE NOT NULL,
    IsRead bit DEFAULT 0::bit,
    TalkMeId integer unique
);

INSERT INTO
    users (UserId, UserName, Password, IsTrader)
VALUES
    (1,'admin', 'admin', 1::bit),
    (2,'user', 'password', 0::bit),
    (3,'login', 'password', 0::bit);


insert into
    orders(OrderId, SellerId, SellerTicker, SellerAmount, BuyerId, BuyerTicker, BuyerAmount)
VALUES
    (1, 1, 'BTC', 0.0002, 2, 'RUB', 800);

insert into
    attachments(attachmentid, path)
values
    (1, './test.http');

select * from messages where SenderId=1 ORDER BY SendTime desc LIMIT 1;
