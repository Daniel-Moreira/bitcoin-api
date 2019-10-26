CREATE TABLE Users (
  UserId       VARCHAR(40) PRIMARY KEY        NOT NULL,
  Password     VARCHAR(40)                    NOT NULL,
  Name         VARCHAR(100)                   NOT NULL,
  Birth        DATE                           NOT NULL,
)

CREATE TABLE Transactions (
  TransactionId   INT PRIMARY KEY AUTO_INCREMENT       NOT NULL,     
  UserId          FOREIGN KEY REFERENCES Users(UserId) NOT NULL,
  Type            VARCHAR(10)                          NOT NULL,
  Amount          DOUBLE                               NULL,
  Price           FLOAT                                NULL,
  Date            DATETIME                             NULL,
)
  ENGINE = INNODB

INSERT INTO db.Users (UserId, Password, Name, Birth) VALUES ('admin', 'admin', 'Administrator', '1992-12-01');
INSERT INTO db.Users (UserId, Type, Amount, Price, Date) VALUES ('admin', 'buy', 2.3, 1000.00, '1992-12-01 00:00:00');

create user gouser@'%' identified by 'gouser@)@<_U2';
grant all privileges on db.* to gouser@'%';
ALTER USER 'gouser'@'%' IDENTIFIED WITH mysql_native_password BY 'gouser@)@<_U2';