<p align="center">
    <img src="https://cdn.discordapp.com/attachments/902632707975168041/912424484143775754/unknown.png" alt="game of life logo image">
</p>

- Platform supported **Windows, Linux, Mac**
# What is?
**Wikipedia definition:** The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves. It is Turing complete and can simulate a universal constructor or any other Turing machine.<p/>
![Img](https://user-images.githubusercontent.com/57921215/95711012-7d471380-0c27-11eb-9b8f-1ea399542ecd.png)

 **Rules:**
- Rule #1: A cell is "dead" if have 3 neighbors alive around.
- Rule #2: A cell is "alive" with less of 2 of more than 3 neighbors alive.
---
# How to install?
**1.** Open your console and clone the repo:
```
git clone https://github.com/jumorap/ConwayGameOfLife
```
**2.** Change the working directory to opera-extension-generator:
```
cd ConwayGameOfLife
```
**3.** Install the requirements (You need to have installed Golang):
```
go get -v -u github.com/gen2brain/raylib-go/raylib
```

And will appear a window with a started funny "game".
**TO CLOSE THE WINDOW** in your console press (Ctrl + C) 1 time, and several times if don't works (don't worry).

# How to use?
Execute main.exe to run the game and select one of the 4 options to deploy the game
- [0] Random game (80 x 80)
- [1] Cyclic game (11 x 11)
- [2] Static game (13 x 13)
- [3] Infinite game (11 x 11)

If you want compile the game with your modifications, use the commands at `ConwayGameOfLife` directory
```
go mod init
```
```
go mod tidy
```
```
go build main.go
```
```
go run main.go
```

*Enjoy this game, you will love how works the automata.*

# License
MIT © Game of Life

-> Feel free to modify this game
