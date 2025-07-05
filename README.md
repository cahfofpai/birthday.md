# birthday.md

Organize the birthdays of your friends using a Markdown-like file format! Convert your birthday.md file to .ics and import it into your calendar to never miss a birthday again!

## CLI tool
Reads your birthday.md file line by line and converts the entries to .ics calendar events.

Syntax: `birthday-md <input file> <output file>`

## File format

```md
# <heading>
dd.mm. <name>
## <heading>
dd.mm.yyyy <name>
<!-- <comment> -->

```

* birthday entries
    * contain the date of the birthday
    * optionally contain the year of birth
    * everything after the date is treated as the name of the person
* headings
    * for structuring your birthday.md file
    * are ignored
* comments
    * for adding comments or ignoring birthdays
    * are ignored
* blank lines
    * are ignored
* invalid lines
    * everything else
    * for every invalid line an error message is printed

## Usage of AI tools

This project uses AI tools for generating source code. All commits containing such code have a co-author with the AI tool string: `<ai tool>:<llm router>:<llm identifier>`. Everything is in lower-case, blanks are replaced by hyphen.

Example: `roo-code:requesty:vertex/anthropic/claude-3-7-sonnet-latest@europe-west1`

You can specify a co-author like this: `git commit -m "<commit message>" -m "Co-authored-by: <ai tool string>"`. For multiple different tools / routers / llms, a separate co-author is specified for every combination.

## License

This project is licensed under the [GNU GPLv3](LICENSE). AI generated content may be licensed differently.