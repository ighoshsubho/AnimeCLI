# AnimeCLI 📰🎌



https://github.com/ighoshsubho/AnimeCLI/assets/93722719/b5053148-6848-478e-b9fc-a144c24a8a7e



This is a command-line interface (CLI) application written in Go that fetches and displays information about recent anime news from Anime News Network (ANN) RSS feed. The app allows you to specify the number of news items you want to display as a command-line argument.

## Features ✨

- Fetches anime news from the [Anime News Network RSS feed](https://www.animenewsnetwork.com/all/rss.xml?ann-edition=in).
- Displays news items including title, link, description, publication date, and category.
- Customizable number of news items to display.

## Prerequisites 🛠️

Before you begin, ensure you have the following:

🐧 [Go](https://go.dev/) installed on your system.

## Installation ⚙️

1. Clone this repository to your local machine:

```bash
https://github.com/ighoshsubho/AnimeCLI
```

2. Change directory to the project folder:

```bash
cd AnimeCLI
```

3. Build the application:

```bash
go build main.go
```

## Usage 🚀

Run the CLI app by executing the compiled binary along with the desired number of news items to display:

```bash
./AnimeCLI [number_of_items]
```

Replace `[number_of_items]` with the number of news items you want to display (default is 5 if no argument is provided).

Example:

```bash
./AnimeCLI 10
```

This will display the latest 10 anime news items from the Anime News Network RSS feed.

## Acknowledgments 🙏

This CLI app utilizes the Anime News Network RSS feed to fetch and display anime news. 
The [Anime News Network](https://www.animenewsnetwork.com/) is a valuable resource for anime-related information.
