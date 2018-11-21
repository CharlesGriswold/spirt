# spirt
A toy that scans through text files to generate random text.

Usage: spirt textfile [n]

The program reads the text file and spirts out n number of randomly-generated pieces of text. n defaults to 1.

The source file format is very simple. It's composed of sections of text separated by lines containing only a single '␞' (U+241E SYMBOL FOR RECORD SEPARATOR) character. Lines starting with '#' (i.e. "comments") are ignored; this is, of course, to allow for shebang lines. :grin: The program chooses a random line from each section, concatenates them together, and prints the resulting bit of nonsense.

The name is a nod to an old random text generator named "spew".
