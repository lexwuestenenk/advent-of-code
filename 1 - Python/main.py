import argparse
from collections import Counter

def main():
    parser = argparse.ArgumentParser(description="Process input for debugging and parts.")
    parser.add_argument(
        "--debug",
        type=bool,
        default=False,
        help="Enable or disable debug mode (default: False)"
    )
    parser.add_argument(
        "--part",
        type=int,
        choices=[1, 2],
        default=1,
        help="Specify which part to execute (1 or 2, default: 1)"
    )
    args = parser.parse_args()

    DEBUG = args.debug
    PART = args.part
    file = ""

    if DEBUG == True:
        file_path = f"1 - Python/input/t_input_{"one" if PART == 1 else "two"}.txt"
        file = open(file_path)
    else: 
        file_path = f"1 - Python/input/input.txt"
        file = open(file_path)

    if PART == 1:
        part_1(file)
    else:
        part_2(file)

def parse_input(line):
    # Replace the first space with a comma and remove the rest of the spaces
    # Strip is called to remove the newlines (\n)
    return line.replace(" ", ",", 1).replace(" ", "").strip().split(",")

def part_1(file):
    left_list = []
    right_list = []
    diff = 0

    for line in file.readlines():
        split_line = parse_input(line)
        left_list.append(int(split_line[0]))
        right_list.append(int(split_line[1]))

    left_list.sort()
    right_list.sort()

    for index, number in enumerate(left_list):
        if number > right_list[index]:
            diff += (number - right_list[index])
            continue

        diff += (right_list[index] - number)

    print(diff)

def part_2(file):
    left_list = []
    right_list = []
    total = 0

    for line in file.readlines():
        split_line = parse_input(line)
        left_list.append(int(split_line[0]))
        right_list.append(int(split_line[1]))

    # Counts the amount of times a value appears in a list
    item_count = Counter(right_list)

    # Loop through the left numbers, multiple by times they appear.
    for number in left_list:
        total += number * item_count[number]

    print(total)

if __name__ == "__main__":
    main()