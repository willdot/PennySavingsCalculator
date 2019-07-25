buildFrontEnd:
	docker build -f "DailyPennySavingsCalculator/dockerfile" -t willdot/pennybudgetcalculatorapp ./DailyPennySavingsCalculator

pushFrontEnd:
	docker push willdot/pennybudgetcalculatorapp:latest
