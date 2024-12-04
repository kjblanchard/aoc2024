def is_safe(report):
    diffs = [report[i + 1] - report[i] for i in range(len(report) - 1)]
    return (
        all(1 <= d <= 3 for d in diffs) or
        all(-3 <= d <= -1 for d in diffs)
    )


def safe_with_one_removal(report):
    n = len(report)
    for i in range(n):
        modified_report = report[:i] + report[i + 1:]
        if is_safe(modified_report):
            return True
    return False


def count_safe_reports_with_dampener(reports):
    safe_count = 0
    for report in reports:
        if is_safe(report) or safe_with_one_removal(report):
            safe_count += 1
    return safe_count


reportsList = []
with open('day2/input.txt', 'r') as reports:
    for report in reports:
        thing = [int(i) for i in report.split()]
        reportsList.append(thing)
count = count_safe_reports_with_dampener(reportsList)
print(count)
