package main

type Utilisateurs struct {
	Id_utilisateurs int
	Nom             string
	Prenom          string
	Adresse_mail    string
	Pseudo          string
}

type Stock_forum struct {
	Id_user      int
	Commentaires string
	User_level   string
	Upvotes      int
	Downvotes    int
}

type Adresses struct {
	Ville       string
	Id_data     int
	Code_postal string
}

type services struct {
	Logs    string
	Adress  string
	Id_logs int
}

type gestion struct {
	Admin_id   int
	Admin_name string
	Modo_name  string
}
