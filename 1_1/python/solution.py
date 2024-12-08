from pathlib import Path
from functools import reduce

input_file_path = Path("input")

def read__columns_from_input() -> tuple[list[int], list[int]]:
    col_a = []
    col_b = []

    with input_file_path.open(mode="r", encoding="utf-8") as input_file:
        for line in input_file.readlines():
            line_items = [int(item) for item in line.split(" ") if item]
            col_a.append(line_items[0])
            col_b.append(line_items[1])

    return col_a, col_b


def calculate_similarity_score(col_a: list[int], col_b: list[int]) -> int:
    return sum(map(lambda x: x * col_b.count(x), col_a))


def main():
    col_a, col_b = read__columns_from_input()
    print(calculate_similarity_score(col_a, col_b))

if __name__ == "__main__":
    main()
