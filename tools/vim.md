# VIM

## Visual Mode

Visual model allows us to select a range of text and then operate upon it. However intuitive this might seem, Vim's perspective on selecting text is different from other text editors.

Many of the commands that you are familiar with from Normal mode work just the same in Visual mode. We can still use h, j, k, and l as cursor keys. We can use f{char} to jump to a character on the current line and then repeat or reverse the jump with the ; and , commands respectively. We can even use the search command (and n/N) to jump to pattern matches. Each time we move our cursor in Visual model, we change the bounds of the selection.

|Command|Effect|
|-|-|
|v/V/<C-v>|Switch to Normal model|
|v|Switch to character-wise Visual mode|
|V|Switch to line-wise Visual mode|
|<C-v>|Switch to block-wise Visual mode|
|o|Go the other end of highlight text|


