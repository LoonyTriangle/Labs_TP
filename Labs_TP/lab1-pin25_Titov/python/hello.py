def greet(name: str) -> str:
    """Формирует приветственное сообщение для пользователя."""
    return f"Hello, {name}!"


def main() -> None:
    name = input("Enter your name: ")
    print(greet(name))


if __name__ == "__main__":
    main()
