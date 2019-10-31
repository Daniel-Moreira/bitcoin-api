CREATE TABLE users_dev (
  UserId        VARCHAR(40) PRIMARY KEY        NOT NULL,
  Password      VARCHAR(40)                    NOT NULL,
  Name          VARCHAR(100)                   NOT NULL,
  Birth         DATE                           NOT NULL,
  BitCoinAmount DOUBLE                         DEFAULT 0
);

CREATE TABLE transactions_dev (
  TransactionId   INT PRIMARY KEY AUTO_INCREMENT            NOT NULL,
  Type            VARCHAR(10)                               NOT NULL,
  UserId          VARCHAR(40)                               NOT NULL,
  Amount          DOUBLE                                    NULL,
  Price           FLOAT                                     NULL,
  Date            DATETIME                                  NULL,
  FOREIGN KEY(UserId) REFERENCES users_dev(userId)
)
  ENGINE = INNODB;

INSERT INTO db.users_dev (UserId, Password, Name, Birth) VALUES ('admin', 'admin', 'Administrator', '1992-12-01');
INSERT INTO db.transactions_dev (UserId, Type, Amount, Price, Date) VALUES ('admin', 'buy', 2.3, 1000.00, '1992-12-01 00:00:00');

create user gouser@'%' identified by 'gouser@)@<_U2';
grant all privileges on db.* to gouser@'%';
ALTER USER 'gouser'@'%' IDENTIFIED WITH mysql_native_password BY 'gouser@)@<_U2';
