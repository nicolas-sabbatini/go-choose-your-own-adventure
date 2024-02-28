# Choose your own adventure

## Description

The third installment of the [Gophercises](https://courses.calhoun.io/lessons/les_goph_06) series.

In this installment we must build a `choose your own adventure` game. The game is a simple
text-based adventure game. The user will be given a story and then have to make a choice
of what they want to do next. Depending on the choice they make, the story will unfold
in different ways. The story is written in a JSON file and the user's choices will dictate
which part of the story they see next.

## Usage

Use the make file to run the application.

```bash
make run
```

If you want to build the application, you must copy the asset folder to the correct location, or give
the program a path to the assets folder. The assets folder contains the JSON file with the story.
