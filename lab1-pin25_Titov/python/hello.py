import sys
import time
from datetime import datetime

COLORS = {
    "red": "\033[91m",
    "green": "\033[92m",
    "yellow": "\033[93m",
    "blue": "\033[94m",
    "magenta": "\033[95m",
    "cyan": "\033[96m",
    "reset": "\033[0m",
}


def greet(name: str) -> str:
    """Формирует приветствие с текущей датой и временем."""
    now = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    return f"Hello, {name}! Current time: {now}"


def get_color(name: str) -> str:
    """Возвращает ANSI-цвет в зависимости от первой буквы имени."""
    if not name:
        return COLORS["reset"]
    letter = name[0].lower()
    if letter <= "e":
        return COLORS["red"]
    elif letter <= "j":
        return COLORS["green"]
    elif letter <= "o":
        return COLORS["yellow"]
    elif letter <= "t":
        return COLORS["blue"]
    elif letter <= "y":
        return COLORS["magenta"]
    return COLORS["cyan"]


def main() -> None:
    if len(sys.argv) > 1:
        name = " ".join(sys.argv[1:])
    else:
        name = input("Enter your name: ")

    if not name:
        name = "World"

    # Замеряем только выполнение (ввод имени уже позади)
    start = time.perf_counter()
    message = greet(name)
    color = get_color(name)
    elapsed_ms = (time.perf_counter() - start) * 1000

    print(f"{color}{message}{COLORS['reset']}")
    print(f"[Python] Execution time: {elapsed_ms:.4f} ms")


if __name__ == "__main__":
    main()
