    create table student (
        Id int primary key auto_increment,
        name varchar(40) not null,
        age int not null,
        cgpa float not null,
        dob  date not null ,
        mobile_num varchar(10) not null,
        created_at datetime not null default current_timestamp
    );

ALTER TABLE student AUTO_INCREMENT = 1;

INSERT INTO student (name, age, cgpa, dob, mobile_num)  VALUES("krishna",22,8.9,"2002-02-20","7995805080");