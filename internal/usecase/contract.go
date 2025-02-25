package usecase

import (
	"fmt"
	"recisao/internal/controllers"
	"recisao/internal/entity"
	"strconv"
	"time"
)

type UsecaseContract struct {
}

func (u UsecaseContract) CalculateContractDetails(contractInfo entity.EmploymentContract, month time.Month, year int) (entity.ContractCalculationResults, error) {

	contractInfo.Validations()

	dailySalary, err := u.calculateDailySalary(contractInfo, month, year)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	years, months, days, err := u.CalculateTimeOfWork(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	Thirteenth, err := u.calculateProportionalThirteenthSalary(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	accruedVacation, err := u.calculateAccruedVacation(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}

	fgts, err := u.CalculoFGTS(contractInfo)
	if err != nil {
		return entity.ContractCalculationResults{}, err
	}
	results := entity.ContractCalculationResults{
		DailySalary:     dailySalary,
		Years:           years,
		Months:          months,
		Days:            days,
		Thirteenth:      Thirteenth,
		AccruedVacation: accruedVacation,
		Fgts:            fgts,
		Total:           Thirteenth + accruedVacation + fgts + contractInfo.Salary,
	}

	return results, nil
}

func (u UsecaseContract) calculateDailySalary(contractInfo entity.EmploymentContract, month time.Month, year int) (string, error) {
	daysInMonth := daysInMonth(month, year)

	salaryPerDay := contractInfo.Salary / float64(daysInMonth)
	formattedSalary := fmt.Sprintf("%.1f", salaryPerDay)
	return formattedSalary, nil
}

func daysInMonth(month time.Month, year int) int {

	longMonths := map[time.Month]bool{
		time.January: true, time.March: true, time.May: true, time.July: true,
		time.August: true, time.October: true, time.December: true,
	}

	if month == time.February {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29
		}
		return 28
	}

	if longMonths[month] {
		return 31
	}
	return 30

}
func (u UsecaseContract) CalculateTimeOfWork(contractInfo entity.EmploymentContract) (int, int, int, error) {
	start := contractInfo.AdmissionDate
	end := contractInfo.DismissalDate

	years := end.Year() - start.Year()
	if end.YearDay() < start.YearDay() {
		years--
	}

	months := int(end.Month()) - int(start.Month())
	if months < 0 {
		months += 12
		years--
	}

	days := end.Day() - start.Day()
	if days < 0 {
		months--
		if months < 0 {
			months += 12
			years--
		}
		days += 30
	}

	return years, months, days, nil
}

func (u UsecaseContract) calculateProportionalThirteenthSalary(contractInfo entity.EmploymentContract) (float64, error) {

	thirteenthSalary := (contractInfo.Salary / 12) * float64(contractInfo.DismissalDate.Month())

	formattedThirteenthSalary := fmt.Sprintf("%.1f", thirteenthSalary)

	thirteenthSalaryFloat, err := strconv.ParseFloat(formattedThirteenthSalary, 64)

	if err != nil {
		return 0, err
	}

	return thirteenthSalaryFloat, nil
}

func (u UsecaseContract) calculateAccruedVacation(contractInfo entity.EmploymentContract) (float64, error) {

	if contractInfo.TerminationType == entity.JustCause && contractInfo.AccumulatedVacationMonths < 12 {

		return 0, nil

	} else if contractInfo.AccumulatedVacationMonths < 12 {

		proportionalVacation := (contractInfo.AccumulatedVacationMonths * 30) / 12
		proportionalVacationValue := (contractInfo.Salary / 30) * float64(proportionalVacation)
		additional := proportionalVacationValue / 3
		total := proportionalVacationValue + additional
		formattedTotal := fmt.Sprintf("%.1f", total)

		formattedFloat, err := strconv.ParseFloat(formattedTotal, 64)

		if err != nil {
			return 0, err
		}

		return formattedFloat, nil

	} else if contractInfo.AccumulatedVacationMonths >= 12 {

		proportionalVacationValue := contractInfo.Salary
		additional := proportionalVacationValue / 3
		totalFor12Months := proportionalVacationValue + additional

		monthsWorked := contractInfo.AccumulatedVacationMonths % 12
		proportionalVacation := (contractInfo.Salary / 30) * float64((monthsWorked * 30 / 12))
		additionalProportional := proportionalVacation / 3
		totalProportional := proportionalVacation + additionalProportional

		TotalPayable := totalFor12Months + totalProportional

		formattedTotalPayable := fmt.Sprintf("%.1f", TotalPayable)

		formattedFloat, err := strconv.ParseFloat(formattedTotalPayable, 64)

		if err != nil {
			return 0, err
		}

		return formattedFloat, nil

	}
	return 0, fmt.Errorf("não foi possível calcular as férias")
}

func (u UsecaseContract) CalculoFGTS(contractInfo entity.EmploymentContract) (float64, error) {

	tempoDeTrabalho, _, _, _ := u.CalculateTimeOfWork(contractInfo)

	fgts := contractInfo.Salary * 8 / 100
	totalDeFgts := fgts * float64(tempoDeTrabalho)

	if contractInfo.TerminationType == entity.WithoutCause {
		finePercentage40 := totalDeFgts * 40 / 100
		totalDeFgts += finePercentage40
	}

	formattedTotal := fmt.Sprintf("%.1f", totalDeFgts)
	formattedFloat, err := strconv.ParseFloat(formattedTotal, 64)
	if err != nil {
		return 0, err
	}

	return formattedFloat, nil
}

func NewUsecaseContract() controllers.Usecase {
	return UsecaseContract{}
}
