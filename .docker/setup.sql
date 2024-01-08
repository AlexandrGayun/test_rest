CREATE DATABASE IF NOT EXISTS test_task_database;

GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `test_task_database`.* TO 'test_task_user'@'%';

FLUSH PRIVILEGES;

# schema

CREATE TABLE `auth` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `api_key` varchar(32) NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `username` varchar(64) NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_profile` (
                            `user_id` bigint(20) NOT NULL,
                            `first_name` varchar(32) NOT NULL,
                            `last_name` varchar(64) NOT NULL,
                            `phone` varchar(64) NOT NULL,
                            `address` varchar(64) NOT NULL,
                            `city` varchar(64) NOT NULL,
                            PRIMARY KEY (`user_id`),
                            FOREIGN KEY(`user_id`) REFERENCES user(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_data` (
                             `user_id` bigint(20) NOT NULL,
                             `school` varchar(32) NOT NULL,
                             PRIMARY KEY (`user_id`),
                             FOREIGN KEY(`user_id`) REFERENCES user(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# seeds
INSERT INTO `auth` VALUES (1,'www-dfq92-sqfwf'),(2,'ffff-2918-xcas');
INSERT INTO `user` VALUES (1,'test'),(2,'admin'),(3,'guest');
INSERT INTO `user_data` VALUES (1,'гімназія №179 міста Києва'),(2,'ліцей №227'),(3,'Медична гімназія №33 міста Києва');
INSERT INTO `user_profile` VALUES (1,'Александр','Школьный','+38050123455','ул. Сибирская 2','Киев'),(2,'Дмитрий','Арбузов','+38065133223','ул. Белая 4','Харьков'),(3,'Василий','Шпак','+38055221166','ул. Северная 5','Житомир');