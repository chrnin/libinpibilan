package libinpibilan

import (
	"encoding/xml"
	"strconv"
	"time"
)

// XMLBilans structure Bilans XML
type XMLBilans struct {
	XMLName xml.Name `xml:"bilans"`
	Bilan   XMLBilan `xml:"bilan"`
}

// XMLBilan structure Bilan XML
type XMLBilan struct {
	XMLName  xml.Name    `xml:"bilan"`
	Identite XMLIdentite `xml:"identite"`
	Detail   XMLDetail   `xml:"detail"`
}

// XMLIdentite structure Identite XML
type XMLIdentite struct {
	XMLName                      xml.Name  `xml:"identite"`
	Siren                        string    `xml:"siren"`
	DateClotureExercice          string    `xml:"date_cloture_exercice"`
	CodeGreffe                   string    `xml:"code_greffe"`
	NumDepot                     string    `xml:"num_depot"`
	NumGestion                   string    `xml:"num_gestion"`
	CodeActivite                 string    `xml:"code_activite"`
	DateClotureExercicePrecedent string    `xml:"date_cloture_exercice_n-1"`
	DureeExercice                string    `xml:"duree_exercice_n"`
	DureeExercicePrecedent       string    `xml:"duree_exercice_n-1"`
	DateDepot                    string    `xml:"date_depot"`
	CodeMotif                    string    `xml:"code_motif"`
	CodeTypeBilan                TypeBilan `xml:"code_type_bilan"`
	CodeDevise                   string    `xml:"code_devise"`
	CodeOrigineDevise            string    `xml:"code_origine_devise"`
	CodeConfidentialite          string    `xml:"code_confidentialite"`
	Denomination                 string    `xml:"denomination"`
	InfoTraitement               string    `xml:"info_traitement"`
	Adresse                      string    `xml:"adresse"`
}

// XMLDetail structure Detail XML
type XMLDetail struct {
	XMLName xml.Name  `xml:"detail"`
	Page    []XMLPage `xml:"page"`
}

// XMLPage structure Page XML
type XMLPage struct {
	XMLName xml.Name       `xml:"page"`
	Numero  Page           `xml:"numero,attr"`
	Liasse  []XMLLigneINPI `xml:"liasse"`
}

// XMLLigneINPI structure Liasse XML
type XMLLigneINPI struct {
	XMLName  xml.Name `xml:"liasse"`
	CodeINPI CodeINPI `xml:"code,attr"`
	M1       *int     `xml:"m1,string,attr"`
	M2       *int     `xml:"m2,string,attr"`
	M3       *int     `xml:"m3,string,attr"`
	M4       *int     `xml:"m4,string,attr"`
}

// Champ
type Champ struct {
	Page        Page
	CodeLiasse  CodeLiasse
	CodeINPI    CodeINPI
	ColonneINPI ColonneINPI
	Rubrique    Rubrique
	Libelle     Libelle
}

type Ligne map[Champ]*int

// Bilan objet restructurant les données d'un bilan au format XML rncs.
type Bilan struct {
	Siren                        string
	DateClotureExercice          time.Time
	CodeGreffe                   string
	NumDepot                     string
	NumGestion                   string
	CodeActivite                 string
	DateClotureExercicePrecedent time.Time
	JoursExercice                int
	JoursExercicePrecedent       int
	DateDepot                    time.Time
	CodeMotif                    string
	CodeTypeBilan                TypeBilan
	CodeDevise                   string
	CodeOrigineDevise            string
	CodeConfidentialite          string
	Denomination                 string
	Adresse                      string
	Source                       *XMLBilans
	NomFichier                   string
	RapportConversion            []error
	InfoTraitement               string
	Lignes                       map[Champ]int
}

func (xmlBilans XMLBilans) BuildBilan() Bilan {
	var bilan Bilan
	bilan.Adresse = xmlBilans.Bilan.Identite.Adresse
	bilan.Siren = xmlBilans.Bilan.Identite.Siren
	dateClotureExercice, err := time.Parse("20060102", xmlBilans.Bilan.Identite.DateClotureExercice)
	if err != nil {
		bilan.RapportConversion = append(bilan.RapportConversion, ErreurDateClotureExerciceInvalide{wrap(err)})
	}
	bilan.DateClotureExercice = dateClotureExercice
	bilan.CodeGreffe = xmlBilans.Bilan.Identite.CodeGreffe
	bilan.NumDepot = xmlBilans.Bilan.Identite.NumDepot
	bilan.NumGestion = xmlBilans.Bilan.Identite.NumGestion
	bilan.CodeActivite = xmlBilans.Bilan.Identite.CodeActivite
	dateClotureExercicePrecedent, err := time.Parse("20060102", xmlBilans.Bilan.Identite.DateClotureExercicePrecedent)
	if err != nil {
		bilan.RapportConversion = append(bilan.RapportConversion, ErreurDateClotureExercicePrecedentInvalide{wrap(err)})
	}
	bilan.DateClotureExercicePrecedent = dateClotureExercicePrecedent
	if dureeExercice, err := strconv.Atoi(xmlBilans.Bilan.Identite.DureeExercice); err != nil {
		bilan.RapportConversion = append(bilan.RapportConversion, ErreurDureeExerciceInvalide{wrap(err)})
	} else {
		bilan.JoursExercice = dureeExercice
	}
	if dureeExercicePrecedent, err := strconv.Atoi(xmlBilans.Bilan.Identite.DureeExercice); err != nil {
		bilan.RapportConversion = append(bilan.RapportConversion, ErreurDureeExercicePrecedentInvalide{wrap(err)})
	} else {
		bilan.JoursExercice = dureeExercicePrecedent
	}

	dateDepot, err := time.Parse("20060102", xmlBilans.Bilan.Identite.DateDepot)
	if err != nil {
		bilan.RapportConversion = append(bilan.RapportConversion, ErreurDateDepotInvalide{wrap(err)})
	}
	bilan.DateDepot = dateDepot
	bilan.CodeMotif = xmlBilans.Bilan.Identite.CodeMotif
	bilan.CodeTypeBilan = xmlBilans.Bilan.Identite.CodeTypeBilan
	bilan.CodeDevise = xmlBilans.Bilan.Identite.CodeDevise
	bilan.CodeOrigineDevise = xmlBilans.Bilan.Identite.CodeOrigineDevise
	bilan.CodeConfidentialite = xmlBilans.Bilan.Identite.CodeConfidentialite
	bilan.Denomination = xmlBilans.Bilan.Identite.Denomination
	bilan.Source = &xmlBilans
	bilan.Lignes = make(map[Champ]int)
	for _, page := range xmlBilans.Bilan.Detail.Page {
		for _, liasse := range page.Liasse {
			ligne := SCHEMA[liasse.CodeINPI][xmlBilans.Bilan.Identite.CodeTypeBilan]
			if err != nil {
				bilan.RapportConversion = append(bilan.RapportConversion, ErreurCodeINPIInconnu{liasse.CodeINPI})
			}
			if err == nil {
				if ligne.M1 != "" && liasse.M1 != nil {
					key := Champ{
						Page:        page.Numero,
						CodeLiasse:  "",
						CodeINPI:    liasse.CodeINPI,
						ColonneINPI: "M1",
						Rubrique:    ligne.PartieBilan,
						Libelle:     ligne.M1,
					}
					bilan.Lignes[key] = *liasse.M1
				}
				if ligne.M2 != "" && liasse.M2 != nil {
					key := Champ{
						Page:        page.Numero,
						CodeLiasse:  "",
						CodeINPI:    liasse.CodeINPI,
						ColonneINPI: "M2",
						Rubrique:    ligne.PartieBilan,
						Libelle:     ligne.M2,
					}
					bilan.Lignes[key] = *liasse.M2
				}
				if ligne.M3 != "" && liasse.M3 != nil {
					key := Champ{
						Page:        page.Numero,
						CodeLiasse:  "",
						CodeINPI:    liasse.CodeINPI,
						ColonneINPI: "M3",
						Rubrique:    ligne.PartieBilan,
						Libelle:     ligne.M3,
					}
					bilan.Lignes[key] = *liasse.M3
				}
				if ligne.M4 != "" && liasse.M4 != nil {
					key := Champ{
						Page:        page.Numero,
						CodeLiasse:  "",
						CodeINPI:    liasse.CodeINPI,
						ColonneINPI: "M4",
						Rubrique:    ligne.PartieBilan,
						Libelle:     ligne.M4,
					}
					bilan.Lignes[key] = *liasse.M4
				}
			}
		}
	}
	return bilan
}
