package data

import (
	"log"
	"time"
)

type Thread struct {
	Id         int
	Uuid       string
	Topic      string
	UserId     int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type Post struct {
	Id         int
	Uuid       string
	Body       string
	UserId     int
	ThreadId   int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// thread: return format Date string
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 1, 2000 at 5:00pm")
}

// post: return format Date string
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 1, 2000 at 5:00pm")
}

// thread: return total nums of replies
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		log.Fatal("Db query failed")
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			log.Fatal("Db query rows iterator failed")
			return
		}
	}
	return
}

// thread: return post along a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query(
		"SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId,
			&post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// thread: Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	crud := "INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1,$2,$3,$4) RETURNING id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(crud)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryPow to return id
	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(
		&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// thread: Query all threads
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// thread: Query thread by UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// thread: return user
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// post: return user
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
