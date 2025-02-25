package usecase

import (
	"fmt"
	"recisao/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Validations(t *testing.T) {
	admissionDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	dismissionDate := time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC)
	salary := 3750.00
	terminationType := entity.WithoutCause

	contractInfo := entity.EmploymentContract{
		AdmissionDate:   admissionDate,
		DismissalDate:   dismissionDate,
		Salary:          salary,
		TerminationType: terminationType,
	}

	t.Run("missing admission date", func(t *testing.T) {
		contractInfo.AdmissionDate = time.Time{}
		uc := contractInfo.Validations()
		assert.ErrorContains(t, uc, "data de admissão não pode ser vazio")
	})

	t.Run("missing dismissal date", func(t *testing.T) {
		contractInfo.AdmissionDate = admissionDate
		contractInfo.DismissalDate = time.Time{}
		uc := contractInfo.Validations()
		assert.ErrorContains(t, uc, "data de demissão não pode ser vazio")
	})

	t.Run("dismissal date before admission date", func(t *testing.T) {
		contractInfo.AdmissionDate = admissionDate
		contractInfo.DismissalDate = dismissionDate.AddDate(-6, 0, 0)
		uc := contractInfo.Validations()
		assert.ErrorContains(t, uc, "a data de demissão não pode ser anterior à data de admissão")

	})

	t.Run("missing salary", func(t *testing.T) {
		contractInfo.AdmissionDate = admissionDate
		contractInfo.DismissalDate = dismissionDate
		contractInfo.Salary = 0
		uc := contractInfo.Validations()
		assert.ErrorContains(t, uc, "salario deve ser maior que zero")
	})

	t.Run("not just cause or without cause", func(t *testing.T) {
		contractInfo.AdmissionDate = admissionDate
		contractInfo.DismissalDate = dismissionDate
		contractInfo.Salary = salary
		contractInfo.TerminationType = "InvalidTerminationType"
		uc := contractInfo.Validations()
		assert.ErrorContains(t, uc, "tipo de recisão inválido, expectativa: justa causa, sem justa causa ou demissão")
	})

	t.Run("success", func(t *testing.T) {
		contractInfo.TerminationType = entity.WithoutCause
		uc := contractInfo.Validations()
		assert.Nil(t, uc)
	})

}

func TestCalculateDailySalary(t *testing.T) {
	contractInfo := entity.EmploymentContract{
		Salary: 1520,
	}

	month := time.February
	year := 2025

	usecase := UsecaseContract{}
	result, err := usecase.calculateDailySalary(contractInfo, month, year)

	assert.NoError(t, err)
	assert.Equal(t, "54.3", result)

	contractInfo.Salary = 1520
	month = time.January
	year = 2025

	result, err = usecase.calculateDailySalary(contractInfo, month, year)

	assert.NoError(t, err)
	assert.Equal(t, "49.0", result)

}

func TestCalculateTimeOfWork(t *testing.T) {

	admissionDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	dismissalDate := time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC)

	contractInfo := entity.EmploymentContract{
		AdmissionDate: admissionDate,
		DismissalDate: dismissalDate,
	}

	usecase := UsecaseContract{}

	years, months, days, err := usecase.CalculateTimeOfWork(contractInfo)

	assert.NoError(t, err, "não deveria haver erro")

	assert.Equal(t, 3, years, "O número de anos calculado deveria ser 3")
	assert.Equal(t, 2, months, "O número de meses calculado deveria ser 2")
	assert.Equal(t, 14, days, "O número de dias calculado deveria ser 14")
}

func TestCalculateProportionalThirteenthSalary(t *testing.T) {

	admissionDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	dismissalDate := time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC)

	contractInfo := entity.EmploymentContract{
		AdmissionDate: admissionDate,
		DismissalDate: dismissalDate,
		Salary:        3750,
	}

	usecase := UsecaseContract{}

	result, err := usecase.calculateProportionalThirteenthSalary(contractInfo)

	assert.NoError(t, err, "não deveria haver erro")

	expectedResult := (3750.0 / 12.0) * 3.0
	expectedResultFormatted := fmt.Sprintf("%.1f", expectedResult)

	assert.Equal(t, expectedResultFormatted, fmt.Sprintf("%.1f", result), "O valor calculado do 13º salário proporcional está incorreto")
}

func TestCalculateAccruedVacation(t *testing.T) {

	contractInfo := entity.EmploymentContract{
		TerminationType:           entity.WithoutCause,
		AccumulatedVacationMonths: 11,
		Salary:                    3750,
	}

	usecase := UsecaseContract{}

	result, err := usecase.calculateAccruedVacation(contractInfo)

	assert.NoError(t, err, "não deveria haver erro")
	assert.Equal(t, 4500.0, result, "O valor calculado das férias acrescidas deveria ser 4500.0")
}

func TestCalculoFGTS(t *testing.T) {
	admissionDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	dismissalDate := time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC)

	contractInfo := entity.EmploymentContract{
		AdmissionDate:   admissionDate,
		DismissalDate:   dismissalDate,
		Salary:          3750,
		TerminationType: entity.JustCause,
	}

	usecase := UsecaseContract{}
	t.Run("Com justa causa", func(t *testing.T) {

		result, err := usecase.CalculoFGTS(contractInfo)

		assert.NoError(t, err, "não deveria haver erro")

		tempoDeTrabalho, _, _, _ := usecase.CalculateTimeOfWork(contractInfo)
		fgts := contractInfo.Salary * 8 / 100
		totalDeFgts := fgts * float64(tempoDeTrabalho)

		expectedResult := fmt.Sprintf("%.1f", totalDeFgts)

		assert.Equal(t, expectedResult, fmt.Sprintf("%.1f", result), "O valor calculado do FGTS está incorreto")
	})

	t.Run("Sem justa causa", func(t *testing.T) {
		contractInfo.TerminationType = entity.WithoutCause

		usecase := UsecaseContract{}

		result, err := usecase.CalculoFGTS(contractInfo)

		assert.NoError(t, err, "não deveria haver erro")

		tempoDeTrabalho, _, _, _ := usecase.CalculateTimeOfWork(contractInfo)
		fgts := contractInfo.Salary * 8 / 100
		totalDeFgts := fgts * float64(tempoDeTrabalho)
		finePercentage40 := totalDeFgts * 40 / 100
		totalReceive := totalDeFgts + finePercentage40

		expectedResult := fmt.Sprintf("%.1f", totalReceive)

		assert.Equal(t, expectedResult, fmt.Sprintf("%.1f", result), "O valor calculado do FGTS está incorreto")
	})

}
