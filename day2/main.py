safeReports = 0
with open('day2/input.txt', 'r') as reports:
    for report in reports.readlines():
        lastDigit = 0
        safe = True
        increasing = None
        for i, level in enumerate(report.split()):
            levelInt = int(level)
            if i == 0:
                lastDigit = levelInt
                continue
            if increasing == None:
                if levelInt < lastDigit:
                    increasing = False
                elif levelInt > lastDigit:
                    increasing = True
                else:
                    safe = False
                    break
            if increasing and levelInt > lastDigit and levelInt < lastDigit + 4:
                lastDigit = levelInt
                continue
            if not increasing and levelInt < lastDigit and levelInt > lastDigit - 4:
                lastDigit = levelInt
                continue
            safe = False
            break
        if safe:
            safeReports += 1
print(safeReports)
