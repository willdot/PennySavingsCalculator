buildFrontEnd:
	docker build -f "DailyPennySavingsCalculator/dockerfile" -t willdot/pennybudgetcalculatorapp ./DailyPennySavingsCalculator

pushFrontEnd:
	docker push willdot/pennybudgetcalculatorapp:latest

buildBackEnd:
	docker build -f "backend/dockerfile" -t willdot/pennybudgetcalculator ./backend

pushBackEnd:
	docker push willdot/pennybudgetcalculator:latest
