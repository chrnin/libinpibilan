package libinpibilan

import "fmt"

type wrappedError struct {
	err error
}

func (w wrappedError) Unwrap() error {
	return w.err
}
func wrap(err error) wrappedError {
	return wrappedError{err}
}

type ErreurCodeINPIInconnu struct {
	codeINPI CodeINPI
}

func (err ErreurCodeINPIInconnu) Error() string {
	return fmt.Sprintf("code INPI inconnu: %s", err.codeINPI)
}

// ErreurCodeLiasseInconnu survient lorsque le champ du fichier XML ne peut être interprêté comme une date
type ErreurCodeLiasseInconnu struct {
	codeLiasse CodeLiasse
}

func (err ErreurCodeLiasseInconnu) Error() string {
	return fmt.Sprintf("code liasse inconnu: %s", err.codeLiasse)
}

type ErreurDateClotureExerciceInvalide struct {
	wrappedError
}

func (err ErreurDateClotureExerciceInvalide) Error() string {
	return "date de cloture de l'exercice invalide"
}

type ErreurDateClotureExercicePrecedentInvalide struct {
	wrappedError
}

func (err ErreurDateClotureExercicePrecedentInvalide) Error() string {
	return "date de cloture de l'exercice précédent invalide"
}

type ErreurDateDepotInvalide struct {
	wrappedError
}

func (err ErreurDateDepotInvalide) Error() string { return "date de dépôt invalide" }

type ErreurDureeExerciceInvalide struct {
	wrappedError
}

func (err ErreurDureeExerciceInvalide) Error() string { return "durée invalide" }

type ErreurDureeExercicePrecedentInvalide struct {
	wrappedError
}

func (err ErreurDureeExercicePrecedentInvalide) Error() string { return "durée invalide" }

type ErreurConversionImparfaite struct {
	nbErreur int
}

func (err ErreurConversionImparfaite) Error() string {
	return fmt.Sprintf("problèmes survenus pendant la conversion: %d", err.nbErreur)
}

type ErreurConversionImpossible struct {
	wrappedError
}

func (err ErreurConversionImpossible) Error() string {
	return "conversion impossible"
}

type ErreurLectureSourceImpossible struct {
	wrappedError
}

func (err ErreurLectureSourceImpossible) Error() string {
	return "problème d'accès à la source de données"
}
