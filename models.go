package main

import "github.com/gocraft/dbr/v2"

type Mumble struct {
	ID         int    `db:"mumble_id"`
	UserName   string `db:"user_name"`
	Mumble     string `db:"mumble"`
	Likes      int    `db:"likes"`
	Mumbletime int    `db:"mumbletime"`
}

func getMumbles(sess *dbr.Session) []Mumble {
	var mumbles []Mumble
	sess.Select("*").From("mumbles").Load(&mumbles)
	return mumbles
}

func getMumble(sess *dbr.Session, mumbleID int) *Mumble {
	var mumble *Mumble
	sess.Select("*").From("mumbles").
		Where("mumble_id = ?", mumbleID).
		LoadOne(&mumble)
	return mumble
}

func addLike(sess *dbr.Session, mumbleID int) {
	likes := getMumble(sess, mumbleID).Likes
	sess.Update("mumbles").
		Set("likes", likes+1).
		Where("mumble_id = ?", 5).
		Exec()
}

func createMumbles(sess *dbr.Session, userName string, mumble string, likes int, mumbletime int) error {
	_, err := sess.InsertInto("mumbles").
		Pair("user_name", userName).
		Pair("mumble", mumble).
		Pair("likes", likes).
		Pair("mumbletime", mumbletime).
		Exec()
	return err
}

func deleteMumble(sess *dbr.Session, mumbleID int) {
	sess.DeleteFrom("mumbles").
		Where("mumble_id = ?", mumbleID).
		Exec()
}
