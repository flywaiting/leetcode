import sys
import os

path = "./rust/src/solution/"


def touch():
    args = sys.argv
    if len(args) < 2:
        print("need title parameter")
        return
    title = "_".join([s.strip().replace(" ", "_") for s in args[1].lower().split(".")][::-1])
    # print(title)

    if not os.path.exists(path):
        # print(path)
        os.makedirs(path)

    file_path = os.path.join(path, title+".rs")
    if os.path.exists(file_path):
        print("file is exist")
        return

    with open(file_path, "w") as _:
        pass


touch()
