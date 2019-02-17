CREATE TABLE `categories` (
	`name` varchar(255),
	`description` varchar(255) DEFAULT 'nothing in here',
	PRIMARY KEY (`name`)
)

CREATE TABLE `emails` (
	`id` integer primary key AUTO_INCREMENT,
	`user_id` integer,
	`email` varchar(100),
	`subscribed` bool
)

CREATE TABLE `languages` (`id` integer primary key AUTO_INCREMENT,`name` varchar(255),`code` varchar(255) )
CREATE TABLE `origins` (`id` integer primary key AUTO_INCREMENT,`product_id` integer,`address1` varchar(255) NOT NULL UNIQUE,`address2` varchar(255) UNIQUE )
CREATE TABLE `products` (`id` integer primary key AUTO_INCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(255),`sid` integer,`score` real NOT NULL  DEFAULT 1.0,`description` varchar(255) DEFAULT 'nothing in here' )
CREATE TABLE `greek_alphabets` (`id` integer primary key AUTO_INCREMENT,`latin_name` varchar(255),`upper_code` integer,`lower_code` integer,`is_frequent` bool )
