from enum import Enum


class Color(Enum):
    RED = "red"
    GREEN = "green"
    BLUE = "blue"


color = Color(input("Enter your choice of 'red', 'blue' or 'green'"))

match color:
    case Color.RED:
        print("I see Red")
    case Color.GREEN:
        print("I see Green")
    case Color.BLUE:
        print("I see Blue")
