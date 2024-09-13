

f = open("input2.txt", "r")

inputs = f.read().splitlines()
print(inputs[0])
for line in inputs:
    game = line.split(':')
    gameNo = game[0].replace('Game ','')
    GameData = game[1].split(';')