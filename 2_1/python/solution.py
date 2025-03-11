import csv
from pathlib import Path
from itertools import pairwise

input_file_path = Path("input.txt")

def read_csv_file():
    with input_file_path.open(mode="r", encoding="utf-8", newline='') as infile:
        reader = csv.reader(infile, delimiter=" ")
        return list(reader)

def is_report_safe(report_line: list[str]) -> int:
    report_line = [int(i) for i in report_line]
    if len(report_line) < 2:
        return 1
    
    diffs = [int(a) - int(b) for a, b in pairwise(report_line)]

    for x in diffs:
        if abs(x) > 3 or abs(x) < 1:
            return 0

    sorted_report_line = sorted(report_line)
    reverse_sorted_report_line = sorted(report_line, reverse=True)
    asc_desc_check = (report_line == sorted_report_line) or (report_line == reverse_sorted_report_line) 

    return int(asc_desc_check)

if __name__ == "__main__":
    report_lines = read_csv_file()

    total = sum(map(is_report_safe, report_lines))
    print(total)
