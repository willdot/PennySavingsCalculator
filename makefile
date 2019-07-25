buildFrontEnd:
	docker build -f "DailyPennySavingsCalculator/dockerfile" -t willdot/pennybudgetcalculatorapp ./DailyPennySavingsCalculator

pushFrontEnd:
	docker push willdot/pennybudgetcalculatorapp:latest

buildBackEnd:
	docker build -f "dockerfile" -t willdot/pennybudgetcalculator .

pushBackEnd:
	docker push willdot/pennybudgetcalculator:latest
