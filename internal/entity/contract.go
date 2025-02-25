package entity

import (
	"errors"
	"time"
)

type TerminationType string

const (
	JustCause    TerminationType = "justa causa"
	WithoutCause TerminationType = "sem justa causa"
	Layoff       TerminationType = "demissão"
)

type EmploymentContract struct {
	AdmissionDate             time.Time       `json:"admission_date"`
	DismissalDate             time.Time       `json:"dismissal_date"`
	Salary                    float64         `json:"salary"`
	TerminationType           TerminationType `json:"termination_type"`
	AccumulatedVacationMonths int             `json:"accumulate_vacation_Months"` 

}

type ContractCalculationResults struct {
	DailySalary     string  `json:"daily_salary"`
	Years           int     `json:"years"`
	Months          int     `json:"months"`
	Days            int     `json:"days"`
	Thirteenth      float64 `json:"thirteenth"`
	AccruedVacation float64 `json:"accrued_vacation"`
	Fgts            float64 `json:"fgts"`
	Total           float64 `json:"total"`
}

func (e EmploymentContract) Validations() error {

	if e.AdmissionDate.IsZero() {
		return errors.New("data de admissão não pode ser vazio")
	}

	if e.DismissalDate.IsZero() {
		return errors.New("data de demissão não pode ser vazio")
	}

	if e.DismissalDate.Before(e.AdmissionDate) {
		return errors.New("a data de demissão não pode ser anterior à data de admissão")
	}

	if e.Salary <= 0 {
		return errors.New("salario deve ser maior que zero")
	}

	if e.TerminationType != JustCause && e.TerminationType != WithoutCause && e.TerminationType != Layoff {
		return errors.New("tipo de recisão inválido, expectativa: justa causa, sem justa causa ou demissão")
	}


	return nil
}
