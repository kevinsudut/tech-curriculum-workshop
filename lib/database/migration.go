package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

var (
	DROP_TABLE_QUERY = []string{
		"DROP TABLE IF EXISTS users",
		"DROP TABLE IF EXISTS books",
		"DROP TABLE IF EXISTS reviews",
	}

	CREATE_TABLE_QUERY = []string{
		`CREATE TABLE IF NOT EXISTS users (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR NOT NULL,
			email VARCHAR NOT NULL,
			password VARCHAR NOT NULL,
			role VARCHAR NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS books (
			id BIGSERIAL PRIMARY KEY,
			title VARCHAR NOT NULL,
			subtitle VARCHAR NOT NULL,
			description VARCHAR NOT NULL,
			isbn VARCHAR NOT NULL,
			author VARCHAR NOT NULL,
			published DATE NOT NULL,
			publisher VARCHAR NOT NULL,
			pages INT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS reviews (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			book_id BIGINT NOT NULL,
			rate INT NOT NULL,
			content VARCHAR NOT NULL,
			status int NOT NULL,
			created_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	QUERY_SEEDER = []string{
		`
			INSERT INTO users (name,email,password,role) VALUES ('Super Admin','secret.admin@techcurri.com',MD5('adminadmin'),'admin');
			INSERT INTO users (name,email,password,role) VALUES ('Kevin SW','kevin.wahyudi@tokopedia.com',MD5('P@ssw0rd!!'),'general')
		`,
		`
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781593279509','Eloquent JavaScript, Third Edition','A Modern Introduction to Programming','Marijn Haverbeke','2018-12-04T00:00:00.000Z','No Starch Press',472,'JavaScript lies at the heart of almost every modern web application, from social apps like Twitter to browser-based game frameworks like Phaser and Babylon. Though simple for beginners to pick up and play with, JavaScript is a flexible, complex language that you can use to build full-scale applications.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781491943533','Practical Modern JavaScript','Dive into ES6 and the Future of JavaScript','Nicolás Bevacqua','2017-07-16T00:00:00.000Z','O''Reilly Media',334,'To get the most out of modern JavaScript, you need learn the latest features of its parent specification, ECMAScript 6 (ES6). This book provides a highly practical look at ES6, without getting lost in the specification or its implementation details.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781593277574','Understanding ECMAScript 6','The Definitive Guide for JavaScript Developers','Nicholas C. Zakas','2016-09-03T00:00:00.000Z','No Starch Press',352,'ECMAScript 6 represents the biggest update to the core of JavaScript in the history of the language. In Understanding ECMAScript 6, expert developer Nicholas C. Zakas provides a complete guide to the object types, syntax, and other exciting changes that ECMAScript 6 brings to JavaScript.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781449365035','Speaking JavaScript','An In-Depth Guide for Programmers','Axel Rauschmayer','2014-04-08T00:00:00.000Z','O''Reilly Media',460,'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781449331818','Learning JavaScript Design Patterns','A JavaScript and jQuery Developer''s Guide','Addy Osmani','2012-08-30T00:00:00.000Z','O''Reilly Media',254,'With Learning JavaScript Design Patterns, you''ll learn how to write beautiful, structured, and maintainable JavaScript by applying classical and modern design patterns to the language. If you want to keep your code efficient, more manageable, and up-to-date with the latest best practices, this book is for you.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9798602477429','You Don''t Know JS Yet','Get Started','Kyle Simpson','2020-01-28T00:00:00.000Z','Independently published',143,'The worldwide best selling You Don''t Know JS book series is back for a 2nd edition: You Don''t Know JS Yet. All 6 books are brand new, rewritten to cover all sides of JS for 2020 and beyond.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781484200766','Pro Git','Everything you neeed to know about Git','Scott Chacon and Ben Straub','2014-11-18T00:00:00.000Z','Apress; 2nd edition',458,'Pro Git (Second Edition) is your fully-updated guide to Git and its usage in the modern world. Git has come a long way since it was first developed by Linus Torvalds for Linux kernel development. It has taken the open source world by storm since its inception in 2005, and this book teaches you how to use it like a pro.');
			INSERT INTO books (isbn,title,subtitle,author,published,publisher,pages,description) VALUES ('9781484242216','Rethinking Productivity in Software Engineering','','Caitlin Sadowski, Thomas Zimmermann','2019-05-11T00:00:00.000Z','Apress',310,'Get the most out of this foundational reference and improve the productivity of your software teams. This open access book collects the wisdom of the 2017 "Dagstuhl" seminar on productivity in software engineering, a meeting of community leaders, who came together with the goal of rethinking traditional definitions and measures of productivity.');
		`,
		`
			INSERT INTO reviews (user_id,book_id,rate,"content",status) VALUES (2,1,5,'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam eu massa gravida, aliquet quam eget, cursus nulla. Vestibulum non risus id nulla posuere bibendum. Praesent eget sem velit. Vivamus efficitur nisi in libero sodales venenatis. Sed aliquet erat quis tristique ultricies. Phasellus rutrum sed quam ac facilisis. Vivamus at magna mauris. Sed ac nunc sed sem rutrum iaculis. Nunc et magna sed libero scelerisque tristique at nec ligula. Donec dui nulla, gravida et tortor a, ultricies dapibus nibh. Integer ornare at risus ac convallis. Duis vitae ultrices erat.',1);
		`,
	}

	JOBS = map[string][]string{
		"Dropped all tables successfully":      DROP_TABLE_QUERY,
		"Migration table created successfully": CREATE_TABLE_QUERY,
		"Seeding data successfully":            QUERY_SEEDER,
	}
)

func Migration() error {
	db, err := Connect()
	if err != nil {
		return errors.AddTrace(err)
	}

	err = runJobs(db)
	if err != nil {
		return errors.AddTrace(err)
	}

	return nil
}

func runJobs(db *sqlx.DB) error {
	for msg, job := range JOBS {
		tx := db.MustBegin()

		for _, query := range job {
			tx.MustExec(query)
		}

		err := tx.Commit()
		if err != nil {
			return errors.AddTrace(err)
		}

		fmt.Println(msg)

		time.Sleep(time.Duration(time.Millisecond * 500))
	}

	return nil
}
