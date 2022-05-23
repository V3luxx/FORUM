package main

type Utilisateurs struct {
	Id_utilisateurs int
	Pseudo          string
	Nom             string
	Prenom          string
	Adresse_mail    string
	Mot_de_passe    string
}

type Stock_forum struct {
	Id_user      int
	Commentaires string
}

type Adresses struct {
	Ville       string
	Id_data     int
	Code_postal string
}

type Services struct {
	Logs    string
	Adress  string
	Id_logs int
}

type Gestion struct {
	Admin_id   int
	Admin_name string
	Modo_name  string
}
