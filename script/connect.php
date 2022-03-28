<?php
    $Pseudo = $_POST['pseudo'];
    $password = $_POST['mot_de_passe'];
    $prenom = $_POST['prenom'];
    $nom = $_POST['nom'];
    $adresse_mail = $_POST['adresse_mail'];

    $conn = new mysqli('localhost', 'root','', 'forum')
    if ($conn->connect_error){
        die('*** Erreur de connection *** :'.$conn->connect_error);
    }else{
        $stmt = $conn->prepare("insert into utilisateurs(pseudo, nom, prenom, adresse_mail, mot_de_passe) VALUES(?, ?, ?, ?, ?)")
        $stmt->bind_para("sssss", $Pseudo, $nom,$prenom, $adresse_mail, $password);
        $stmt->execute();
        echo("connection reussie")
        $stmt->close();
        $conn->close();
    }
?>